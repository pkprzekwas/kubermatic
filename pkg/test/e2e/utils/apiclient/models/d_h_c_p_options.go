// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DHCPOptions Extra DHCP options to use in the interface.
//
// swagger:model DHCPOptions
type DHCPOptions struct {

	// If specified will pass option 67 to interface's DHCP server
	// +optional
	BootFileName string `json:"bootFileName,omitempty"`

	// If specified will pass the configured NTP server to the VM via DHCP option 042.
	// +optional
	NTPServers []string `json:"ntpServers"`

	// If specified will pass extra DHCP options for private use, range: 224-254
	// +optional
	PrivateOptions []*DHCPPrivateOptions `json:"privateOptions"`

	// If specified will pass option 66 to interface's DHCP server
	// +optional
	TFTPServerName string `json:"tftpServerName,omitempty"`
}

// Validate validates this d h c p options
func (m *DHCPOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DHCPOptions) validatePrivateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.PrivateOptions); i++ {
		if swag.IsZero(m.PrivateOptions[i]) { // not required
			continue
		}

		if m.PrivateOptions[i] != nil {
			if err := m.PrivateOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privateOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this d h c p options based on the context it is used
func (m *DHCPOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DHCPOptions) contextValidatePrivateOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrivateOptions); i++ {

		if m.PrivateOptions[i] != nil {
			if err := m.PrivateOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privateOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DHCPOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DHCPOptions) UnmarshalBinary(b []byte) error {
	var res DHCPOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
