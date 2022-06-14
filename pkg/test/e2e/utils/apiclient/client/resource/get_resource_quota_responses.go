// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetResourceQuotaReader is a Reader for the GetResourceQuota structure.
type GetResourceQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetResourceQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResourceQuotaOK creates a GetResourceQuotaOK with default headers values
func NewGetResourceQuotaOK() *GetResourceQuotaOK {
	return &GetResourceQuotaOK{}
}

/* GetResourceQuotaOK describes a response with status code 200, with default header values.

ResourceQuota
*/
type GetResourceQuotaOK struct {
	Payload *models.ResourceQuota
}

func (o *GetResourceQuotaOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/quotas/{name}][%d] getResourceQuotaOK  %+v", 200, o.Payload)
}
func (o *GetResourceQuotaOK) GetPayload() *models.ResourceQuota {
	return o.Payload
}

func (o *GetResourceQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceQuota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceQuotaUnauthorized creates a GetResourceQuotaUnauthorized with default headers values
func NewGetResourceQuotaUnauthorized() *GetResourceQuotaUnauthorized {
	return &GetResourceQuotaUnauthorized{}
}

/* GetResourceQuotaUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetResourceQuotaUnauthorized struct {
}

func (o *GetResourceQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/quotas/{name}][%d] getResourceQuotaUnauthorized ", 401)
}

func (o *GetResourceQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceQuotaForbidden creates a GetResourceQuotaForbidden with default headers values
func NewGetResourceQuotaForbidden() *GetResourceQuotaForbidden {
	return &GetResourceQuotaForbidden{}
}

/* GetResourceQuotaForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetResourceQuotaForbidden struct {
}

func (o *GetResourceQuotaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/quotas/{name}][%d] getResourceQuotaForbidden ", 403)
}

func (o *GetResourceQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceQuotaDefault creates a GetResourceQuotaDefault with default headers values
func NewGetResourceQuotaDefault(code int) *GetResourceQuotaDefault {
	return &GetResourceQuotaDefault{
		_statusCode: code,
	}
}

/* GetResourceQuotaDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetResourceQuotaDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get resource quota default response
func (o *GetResourceQuotaDefault) Code() int {
	return o._statusCode
}

func (o *GetResourceQuotaDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/quotas/{name}][%d] getResourceQuota default  %+v", o._statusCode, o.Payload)
}
func (o *GetResourceQuotaDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourceQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetResourceQuotaBody get resource quota body
swagger:model GetResourceQuotaBody
*/
type GetResourceQuotaBody struct {

	// quota
	Quota *models.ResourceDetails `json:"Quota,omitempty"`

	// subject
	Subject *models.Subject `json:"Subject,omitempty"`
}

// Validate validates this get resource quota body
func (o *GetResourceQuotaBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourceQuotaBody) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(o.Quota) { // not required
		return nil
	}

	if o.Quota != nil {
		if err := o.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "Quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "Quota")
			}
			return err
		}
	}

	return nil
}

func (o *GetResourceQuotaBody) validateSubject(formats strfmt.Registry) error {
	if swag.IsZero(o.Subject) { // not required
		return nil
	}

	if o.Subject != nil {
		if err := o.Subject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "Subject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "Subject")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get resource quota body based on the context it is used
func (o *GetResourceQuotaBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSubject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourceQuotaBody) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if o.Quota != nil {
		if err := o.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "Quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "Quota")
			}
			return err
		}
	}

	return nil
}

func (o *GetResourceQuotaBody) contextValidateSubject(ctx context.Context, formats strfmt.Registry) error {

	if o.Subject != nil {
		if err := o.Subject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "Subject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "Subject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceQuotaBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceQuotaBody) UnmarshalBinary(b []byte) error {
	var res GetResourceQuotaBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
