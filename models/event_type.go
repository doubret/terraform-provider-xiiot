// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// EventType event type
// swagger:model EventType
type EventType string

const (

	// EventTypeALERT captures enum value "ALERT"
	EventTypeALERT EventType = "ALERT"

	// EventTypeSTATUS captures enum value "STATUS"
	EventTypeSTATUS EventType = "STATUS"

	// EventTypeMETRIC captures enum value "METRIC"
	EventTypeMETRIC EventType = "METRIC"
)

// for schema
var eventTypeEnum []interface{}

func init() {
	var res []EventType
	if err := json.Unmarshal([]byte(`["ALERT","STATUS","METRIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypeEnum = append(eventTypeEnum, v)
	}
}

func (m EventType) validateEventTypeEnum(path, location string, value EventType) error {
	if err := validate.Enum(path, location, value, eventTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this event type
func (m EventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}