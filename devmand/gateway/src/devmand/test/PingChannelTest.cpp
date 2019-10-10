// Copyright (c) 2016-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#include <devmand/channels/ping/Channel.h>
#include <devmand/test/EventBaseTest.h>

namespace devmand {
namespace test {

class PingChannelTest : public EventBaseTest {
 public:
  PingChannelTest() = default;
  ~PingChannelTest() override = default;
  PingChannelTest(const PingChannelTest&) = delete;
  PingChannelTest& operator=(const PingChannelTest&) = delete;
  PingChannelTest(PingChannelTest&&) = delete;
  PingChannelTest& operator=(PingChannelTest&&) = delete;
};

TEST_F(PingChannelTest, checkPing) {
  folly::IPAddress address("127.0.0.1");
  auto channel = std::make_shared<channels::ping::Channel>(eventBase, address);
  channel->ping();
}

} // namespace test
} // namespace devmand
