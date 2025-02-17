// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HealthCheckResponse health check response
//
// swagger:model HealthCheckResponse
type HealthCheckResponse struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this health check response
func (m *HealthCheckResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health check response based on context it is used
func (m *HealthCheckResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthCheckResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthCheckResponse) UnmarshalBinary(b []byte) error {
	var res HealthCheckResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
