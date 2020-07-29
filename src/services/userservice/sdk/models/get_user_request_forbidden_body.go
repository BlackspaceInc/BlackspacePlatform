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

// GetUserRequestForbiddenBody GetUserRequestForbiddenBody get user request forbidden body
// swagger:model GetUserRequestForbiddenBody
type GetUserRequestForbiddenBody struct {

	// description of the error
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get user request forbidden body
func (m *GetUserRequestForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUserRequestForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUserRequestForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserRequestForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetUserRequestForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
