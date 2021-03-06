// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// JSONText JSONText is a json.RawMessage, which is a []byte underneath.
//
// Value() validates the json format in the source, and returns an error if
// the json is not valid.  Scan does no validation.  JSONText additionally
// implements `Unmarshal`, which unmarshals the json within to an interface{}
// swagger:model JSONText
type JSONText struct {
	RawMessage
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *JSONText) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RawMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RawMessage = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m JSONText) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.RawMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this JSON text
func (m *JSONText) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JSONText) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONText) UnmarshalBinary(b []byte) error {
	var res JSONText
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
