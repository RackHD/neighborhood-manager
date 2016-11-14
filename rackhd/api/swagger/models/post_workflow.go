package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// PostWorkflow post workflow
// swagger:model post_workflow
type PostWorkflow struct {

	// name
	Name string `json:"name,omitempty"`

	// options
	Options *PostWorkflowOptions `json:"options,omitempty"`
}

// Validate validates this post workflow
func (m *PostWorkflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostWorkflow) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {

		if err := m.Options.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// PostWorkflowOptions post workflow options
// swagger:model PostWorkflowOptions
type PostWorkflowOptions struct {

	// defaults
	Defaults *PostWorkflowOptionsDefaults `json:"defaults,omitempty"`
}

// Validate validates this post workflow options
func (m *PostWorkflowOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaults(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostWorkflowOptions) validateDefaults(formats strfmt.Registry) error {

	if swag.IsZero(m.Defaults) { // not required
		return nil
	}

	if m.Defaults != nil {

		if err := m.Defaults.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// PostWorkflowOptionsDefaults post workflow options defaults
// swagger:model PostWorkflowOptionsDefaults
type PostWorkflowOptionsDefaults struct {

	// graph options
	GraphOptions interface{} `json:"graphOptions,omitempty"`

	// node Id
	NodeID string `json:"nodeId,omitempty"`
}

// Validate validates this post workflow options defaults
func (m *PostWorkflowOptionsDefaults) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
