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

// WstunRequest wstun request
// swagger:model WstunRequest
type WstunRequest struct {

	// edge ID
	// Required: true
	EdgeID *string `json:"edgeId"`

	// tenant ID
	// Required: true
	TenantID *string `json:"tenantId"`
}

// Validate validates this wstun request
func (m *WstunRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WstunRequest) validateEdgeID(formats strfmt.Registry) error {

	if err := validate.Required("edgeId", "body", m.EdgeID); err != nil {
		return err
	}

	return nil
}

func (m *WstunRequest) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WstunRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WstunRequest) UnmarshalBinary(b []byte) error {
	var res WstunRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
