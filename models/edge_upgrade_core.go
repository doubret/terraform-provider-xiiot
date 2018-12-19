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

// EdgeUpgradeCore EdgeUpgradeCore is object model for EdgeUpgradeCore
// swagger:model EdgeUpgradeCore
type EdgeUpgradeCore struct {

	// The changes that were made in this release from the previous release
	// Required: true
	Changelog *string `json:"changelog"`

	// This is the release that is avaliable
	// Required: true
	Release *string `json:"release"`
}

// Validate validates this edge upgrade core
func (m *EdgeUpgradeCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangelog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeUpgradeCore) validateChangelog(formats strfmt.Registry) error {

	if err := validate.Required("changelog", "body", m.Changelog); err != nil {
		return err
	}

	return nil
}

func (m *EdgeUpgradeCore) validateRelease(formats strfmt.Registry) error {

	if err := validate.Required("release", "body", m.Release); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeUpgradeCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeUpgradeCore) UnmarshalBinary(b []byte) error {
	var res EdgeUpgradeCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
