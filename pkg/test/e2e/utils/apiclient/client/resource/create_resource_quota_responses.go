// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateResourceQuotaReader is a Reader for the CreateResourceQuota structure.
type CreateResourceQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourceQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateResourceQuotaCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateResourceQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateResourceQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateResourceQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateResourceQuotaCreated creates a CreateResourceQuotaCreated with default headers values
func NewCreateResourceQuotaCreated() *CreateResourceQuotaCreated {
	return &CreateResourceQuotaCreated{}
}

/* CreateResourceQuotaCreated describes a response with status code 201, with default header values.

EmptyResponse is a empty response
*/
type CreateResourceQuotaCreated struct {
}

func (o *CreateResourceQuotaCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/quotas][%d] createResourceQuotaCreated ", 201)
}

func (o *CreateResourceQuotaCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateResourceQuotaUnauthorized creates a CreateResourceQuotaUnauthorized with default headers values
func NewCreateResourceQuotaUnauthorized() *CreateResourceQuotaUnauthorized {
	return &CreateResourceQuotaUnauthorized{}
}

/* CreateResourceQuotaUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateResourceQuotaUnauthorized struct {
}

func (o *CreateResourceQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/quotas][%d] createResourceQuotaUnauthorized ", 401)
}

func (o *CreateResourceQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateResourceQuotaForbidden creates a CreateResourceQuotaForbidden with default headers values
func NewCreateResourceQuotaForbidden() *CreateResourceQuotaForbidden {
	return &CreateResourceQuotaForbidden{}
}

/* CreateResourceQuotaForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateResourceQuotaForbidden struct {
}

func (o *CreateResourceQuotaForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/quotas][%d] createResourceQuotaForbidden ", 403)
}

func (o *CreateResourceQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateResourceQuotaDefault creates a CreateResourceQuotaDefault with default headers values
func NewCreateResourceQuotaDefault(code int) *CreateResourceQuotaDefault {
	return &CreateResourceQuotaDefault{
		_statusCode: code,
	}
}

/* CreateResourceQuotaDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateResourceQuotaDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create resource quota default response
func (o *CreateResourceQuotaDefault) Code() int {
	return o._statusCode
}

func (o *CreateResourceQuotaDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/quotas][%d] createResourceQuota default  %+v", o._statusCode, o.Payload)
}
func (o *CreateResourceQuotaDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateResourceQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
