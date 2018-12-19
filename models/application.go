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

// Application Application - the contents of an Application
// swagger:model Application
type Application struct {

	// The description of the application
	Description string `json:"description,omitempty"`

	// List of ids of edges this application should be deployed to.
	// Only relevant if parent project EdgeSelectorType = 'Explicit'
	EdgeIds []string `json:"edgeIds"`

	// Edge selectors - CategoryInfo list.
	// Only relevant when parent project EdgeSelectorType = 'Category'
	EdgeSelectors []*CategoryInfo `json:"edgeSelectors"`

	// ID
	ID string `json:"id,omitempty"`

	// The name of the application
	// E.g., FaceFeed
	// Required: true
	Name *string `json:"name"`

	// ID of parent project.
	// This should be required, but is not marked as such due to backward compatibility.
	ProjectID string `json:"projectId,omitempty"`

	// The yaml content for the app
	// Required: true
	YamlData *string `json:"yamlData"`
}

// Validate validates this application
func (m *Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYamlData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Application) validateEdgeSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.EdgeSelectors); i++ {
		if swag.IsZero(m.EdgeSelectors[i]) { // not required
			continue
		}

		if m.EdgeSelectors[i] != nil {
			if err := m.EdgeSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateYamlData(formats strfmt.Registry) error {

	if err := validate.Required("yamlData", "body", m.YamlData); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Application) UnmarshalBinary(b []byte) error {
	var res Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
