// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogEntry LogEntry - a log entry describes the metadata for a log bundle
// from an edge collected in part of a batch for a given tenant.
// swagger:model LogEntry
type LogEntry struct {

	// id that identifies logs from different edge as the same batch.
	BatchID string `json:"batchId,omitempty"`

	// edge ID
	// Required: true
	EdgeID *string `json:"edgeId"`

	// Error message - optional, should be populated when status == 'FAILED'
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// Location or object key for the log in the bucket.
	Location string `json:"location,omitempty"`

	// Tags carry the properties of the log
	Tags []*LogTag `json:"tags"`

	// status
	Status LogUploadStatus `json:"status,omitempty"`
}

// Validate validates this log entry
func (m *LogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntry) validateEdgeID(formats strfmt.Registry) error {

	if err := validate.Required("edgeId", "body", m.EdgeID); err != nil {
		return err
	}

	return nil
}

func (m *LogEntry) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogEntry) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntry) UnmarshalBinary(b []byte) error {
	var res LogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
