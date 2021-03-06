#!/usr/bin/env python3

"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""


import asyncio
import logging
import shlex
import subprocess
import re
from typing import List
from magma.common.misc_utils import (
    IpPreference,
    get_ip_from_if,
    get_if_ip_with_netmask
)
from magma.configuration.service_configs import load_service_config

IPTABLES_RULE_FMT = """sudo iptables -t nat
    -{add} PREROUTING
    -d {public_ip}
    -p tcp
    --dport {port}
    -j DNAT --to-destination {private_ip}"""

EXPECTED_IP4 = ('192.168.60.142', '10.0.2.1')
EXPECTED_MASK = '255.255.255.0'


def get_iptables_rule(port, enodebd_public_ip, private_ip, add=True):
    return IPTABLES_RULE_FMT.format(
        add='A' if add else 'D',
        public_ip=enodebd_public_ip,
        port=port,
        private_ip=private_ip,
    )


def does_iface_config_match_expected(ip: str, netmask: str) -> bool:
    return ip in EXPECTED_IP4 and netmask == EXPECTED_MASK


def _get_prerouting_rules(output: str) -> List[str]:
    prerouting_rules = output.split('\n\n')[0]
    prerouting_rules = prerouting_rules.split('\n')
    # Skipping the first two lines since it contains only column names
    prerouting_rules = prerouting_rules[2:]
    return prerouting_rules


async def check_and_apply_iptables_rules(port: str,
                                         enodebd_public_ip: str,
                                         enodebd_ip: str) -> None:
    command = 'sudo iptables -t nat -L'
    output = subprocess.run(command, shell=True, stdout=subprocess.PIPE, check=True)
    command_output = output.stdout.decode('utf-8').strip()
    prerouting_rules = _get_prerouting_rules(command_output)
    if not prerouting_rules:
        logging.info('Configuring Iptables rule')
        await run(
            get_iptables_rule(
                port,
                enodebd_public_ip,
                enodebd_ip,
                add=True,
            )
        )
    else:
        # Checks each rule in PREROUTING Chain
        check_rules(prerouting_rules, port, enodebd_public_ip, enodebd_ip)


def check_rules(prerouting_rules: List[str],
                port: str,
                enodebd_public_ip: str,
                private_ip: str) -> None:
    unexpected_rules = []
    pattern = r'DNAT\s+tcp\s+--\s+anywhere\s+{pub_ip}\s+tcp\s+dpt:{dport} to:{ip}'.format(
                pub_ip=enodebd_public_ip,
                dport=port,
                ip=private_ip,
    )
    for rule in prerouting_rules:
        match = re.search(pattern, rule)
        if not match:
            unexpected_rules.append(rule)
    if unexpected_rules:
        logging.warning('The following Prerouting rule(s) are unexpected')
        for rule in unexpected_rules:
            logging.warning(rule)


async def run(cmd):
    """Fork shell and run command NOTE: Popen is non-blocking"""
    cmd = shlex.split(cmd)
    proc = await asyncio.create_subprocess_shell(" ".join(cmd))
    await proc.communicate()
    if proc.returncode != 0:
        # This can happen because the NAT prerouting rule didn't exist
        logging.info('Possible error running async subprocess: %s exited with '
                     'return code [%d].', cmd, proc.returncode)
    return proc.returncode


async def set_enodebd_iptables_rule():
    """
    Remove & Set iptable rules for exposing public IP
    for enobeb instead of private IP..
    """
    # Remove & Set iptable rules for exposing public ip
    # for enobeb instead of private
    cfg = load_service_config('enodebd')
    port, interface = cfg['tr069']['port'], cfg['tr069']['interface']
    enodebd_public_ip = cfg['tr069']['public_ip']
    # IPv4 only as iptables only works for IPv4. TODO: Investigate ip6tables?
    enodebd_ip = get_ip_from_if(interface, preference=IpPreference.IPV4_ONLY)
    # Incoming data from 192.88.99.142 -> enodebd address (eg 192.168.60.142)
    enodebd_netmask = get_if_ip_with_netmask(
        interface,
        preference=IpPreference.IPV4_ONLY,
    )[1]
    verify_config = does_iface_config_match_expected(
        enodebd_ip,
        enodebd_netmask,
    )
    if not verify_config:
        logging.warning(
            'The IP address of the %s interface is %s. The '
            'expected IP addresses are %s',
            interface, enodebd_ip, str(EXPECTED_IP4)
        )
    await check_and_apply_iptables_rules(
        port,
        enodebd_public_ip,
        enodebd_ip,
    )


if __name__ == '__main__':
    set_enodebd_iptables_rule()
