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

// Tenant Tenant is the DB object and object model for tenant
//
// A tenant represents a customer account.
// A tenant may have multiple edges.
// Every object in DB belonging to a tenant
// will have a tenantId field.
// Tenant object, like every other object
// in DB will have id and version fields.
// The id and version fields are marked as optional
// because they are not required in create operation.
//
// Use Float b/c convert to map will change int to float64
// swagger:model Tenant
type Tenant struct {

	// The description of the tenant
	Description string `json:"description,omitempty"`

	// Unique tenant ID returned by my.nutanix.com
	// Required: true
	ExternalID *string `json:"externalId"`

	// Unique id to identify the tenant.
	// This could be supplied during create or DB generated.
	// For Nice we will have fixed tenant id such as
	// tenant-id-waldot
	// tenant-id-rocket-blue
	// Required: true
	ID *string `json:"id"`

	// Name of the tenant.
	// E.g., WalDot or Rocket Blue, etc.
	// Required: true
	Name *string `json:"name"`

	// Unique token for tenant.
	// Used in authentication.
	// Required: true
	Token *string `json:"token"`

	// Version number of object maintained by DB.
	// Not currently used.
	Version float64 `json:"version,omitempty"`
}

// Validate validates this tenant
func (m *Tenant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tenant) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("externalId", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *Tenant) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Tenant) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Tenant) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tenant) UnmarshalBinary(b []byte) error {
	var res Tenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
