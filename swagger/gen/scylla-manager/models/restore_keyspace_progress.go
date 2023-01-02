// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestoreKeyspaceProgress restore keyspace progress
//
// swagger:model RestoreKeyspaceProgress
type RestoreKeyspaceProgress struct {

	// completed at
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	// downloaded
	Downloaded int64 `json:"downloaded,omitempty"`

	// failed
	Failed int64 `json:"failed,omitempty"`

	// keyspace
	Keyspace string `json:"keyspace,omitempty"`

	// restored
	Restored int64 `json:"restored,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// skipped
	Skipped int64 `json:"skipped,omitempty"`

	// started at
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	// tables
	Tables []*RestoreTableProgress `json:"tables"`
}

// Validate validates this restore keyspace progress
func (m *RestoreKeyspaceProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreKeyspaceProgress) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RestoreKeyspaceProgress) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RestoreKeyspaceProgress) validateTables(formats strfmt.Registry) error {

	if swag.IsZero(m.Tables) { // not required
		return nil
	}

	for i := 0; i < len(m.Tables); i++ {
		if swag.IsZero(m.Tables[i]) { // not required
			continue
		}

		if m.Tables[i] != nil {
			if err := m.Tables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreKeyspaceProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreKeyspaceProgress) UnmarshalBinary(b []byte) error {
	var res RestoreKeyspaceProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}