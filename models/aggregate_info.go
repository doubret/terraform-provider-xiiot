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

// AggregateInfo AggregateInfo is aggregate query response item
// swagger:model AggregateInfo
type AggregateInfo struct {

	// doc count
	// Required: true
	DocCount *int64 `json:"doc_count"`

	// key
	// Required: true
	Key *string `json:"key"`
}

// Validate validates this aggregate info
func (m *AggregateInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateInfo) validateDocCount(formats strfmt.Registry) error {

	if err := validate.Required("doc_count", "body", m.DocCount); err != nil {
		return err
	}

	return nil
}

func (m *AggregateInfo) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AggregateInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateInfo) UnmarshalBinary(b []byte) error {
	var res AggregateInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
