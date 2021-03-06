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

// Certificates Certificates is DB and object model for a device's certificates
// swagger:model Certificates
type Certificates struct {

	// Root CA certificate that signed the device certificate.
	// Required: true
	CACertificate *string `json:"CACertificate"`

	// Certificate for a device.
	// Required: true
	Certificate *string `json:"certificate"`

	// Encrypted private key corresponding to the certificate.
	// Required: true
	PrivateKey *string `json:"privateKey"`
}

// Validate validates this certificates
func (m *Certificates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCACertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Certificates) validateCACertificate(formats strfmt.Registry) error {

	if err := validate.Required("CACertificate", "body", m.CACertificate); err != nil {
		return err
	}

	return nil
}

func (m *Certificates) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

func (m *Certificates) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Certificates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Certificates) UnmarshalBinary(b []byte) error {
	var res Certificates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
