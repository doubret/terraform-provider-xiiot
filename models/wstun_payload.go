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

// WstunPayload wstun payload
// swagger:model WstunPayload
type WstunPayload struct {

	// edge ID
	// Required: true
	EdgeID *string `json:"edgeId"`

	// expiration
	// Required: true
	Expiration *int64 `json:"expiration"`

	// port
	// Required: true
	Port *uint32 `json:"port"`

	// tenant ID
	// Required: true
	TenantID *string `json:"tenantId"`
}

// Validate validates this wstun payload
func (m *WstunPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
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

func (m *WstunPayload) validateEdgeID(formats strfmt.Registry) error {

	if err := validate.Required("edgeId", "body", m.EdgeID); err != nil {
		return err
	}

	return nil
}

func (m *WstunPayload) validateExpiration(formats strfmt.Registry) error {

	if err := validate.Required("expiration", "body", m.Expiration); err != nil {
		return err
	}

	return nil
}

func (m *WstunPayload) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *WstunPayload) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WstunPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WstunPayload) UnmarshalBinary(b []byte) error {
	var res WstunPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
