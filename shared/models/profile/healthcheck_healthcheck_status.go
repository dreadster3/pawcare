// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HealthcheckHealthcheckStatus healthcheck healthcheck status
//
// swagger:model healthcheck.HealthcheckStatus
type HealthcheckHealthcheckStatus string

func NewHealthcheckHealthcheckStatus(value HealthcheckHealthcheckStatus) *HealthcheckHealthcheckStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HealthcheckHealthcheckStatus.
func (m HealthcheckHealthcheckStatus) Pointer() *HealthcheckHealthcheckStatus {
	return &m
}

const (

	// HealthcheckHealthcheckStatusHealthy captures enum value "healthy"
	HealthcheckHealthcheckStatusHealthy HealthcheckHealthcheckStatus = "healthy"

	// HealthcheckHealthcheckStatusUnhealthy captures enum value "unhealthy"
	HealthcheckHealthcheckStatusUnhealthy HealthcheckHealthcheckStatus = "unhealthy"

	// HealthcheckHealthcheckStatusDegraded captures enum value "degraded"
	HealthcheckHealthcheckStatusDegraded HealthcheckHealthcheckStatus = "degraded"
)

// for schema
var healthcheckHealthcheckStatusEnum []interface{}

func init() {
	var res []HealthcheckHealthcheckStatus
	if err := json.Unmarshal([]byte(`["healthy","unhealthy","degraded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		healthcheckHealthcheckStatusEnum = append(healthcheckHealthcheckStatusEnum, v)
	}
}

func (m HealthcheckHealthcheckStatus) validateHealthcheckHealthcheckStatusEnum(path, location string, value HealthcheckHealthcheckStatus) error {
	if err := validate.EnumCase(path, location, value, healthcheckHealthcheckStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this healthcheck healthcheck status
func (m HealthcheckHealthcheckStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHealthcheckHealthcheckStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this healthcheck healthcheck status based on context it is used
func (m HealthcheckHealthcheckStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}