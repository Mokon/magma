// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrometheusLabelset prometheus labelset
// swagger:model prometheus_labelset
type PrometheusLabelset struct {

	// name
	// Required: true
	Name *string `json:"__name__"`

	// prometheus labelset
	PrometheusLabelset map[string]string `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *PrometheusLabelset) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// name
		// Required: true
		Name *string `json:"__name__"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv PrometheusLabelset

	rcv.Name = stage1.Name

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "__name__")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]string)
		for k, v := range stage2 {
			var toadd string
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.PrometheusLabelset = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m PrometheusLabelset) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// name
		// Required: true
		Name *string `json:"__name__"`
	}

	stage1.Name = m.Name

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.PrometheusLabelset) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.PrometheusLabelset)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this prometheus labelset
func (m *PrometheusLabelset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrometheusLabelset) validateName(formats strfmt.Registry) error {

	if err := validate.Required("__name__", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusLabelset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusLabelset) UnmarshalBinary(b []byte) error {
	var res PrometheusLabelset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
