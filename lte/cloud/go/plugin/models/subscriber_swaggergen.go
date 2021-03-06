// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subscriber subscriber
// swagger:model subscriber
type Subscriber struct {

	// id
	// Required: true
	// Pattern: ^(IMSI\d{10,15})$
	ID string `json:"id"`

	// lte
	// Required: true
	Lte *LteSubscription `json:"lte"`
}

// Validate validates this subscriber
func (m *Subscriber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLte(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscriber) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	if err := validate.Pattern("id", "body", string(m.ID), `^(IMSI\d{10,15})$`); err != nil {
		return err
	}

	return nil
}

func (m *Subscriber) validateLte(formats strfmt.Registry) error {

	if err := validate.Required("lte", "body", m.Lte); err != nil {
		return err
	}

	if m.Lte != nil {
		if err := m.Lte.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lte")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscriber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscriber) UnmarshalBinary(b []byte) error {
	var res Subscriber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
