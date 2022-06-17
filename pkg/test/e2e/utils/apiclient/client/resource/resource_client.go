// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resource API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateResourceQuota(params *CreateResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateResourceQuotaCreated, error)

	DeleteResourceQuota(params *DeleteResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteResourceQuotaOK, error)

	GetResourceQuota(params *GetResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceQuotaOK, error)

	ListResourceQuotas(params *ListResourceQuotasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListResourceQuotasOK, error)

	UpdateResourceQuota(params *UpdateResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateResourceQuotaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateResourceQuota creates a new resource quota
*/
func (a *Client) CreateResourceQuota(params *CreateResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateResourceQuotaCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResourceQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createResourceQuota",
		Method:             "POST",
		PathPattern:        "/api/v1/admin/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateResourceQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateResourceQuotaCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateResourceQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteResourceQuota removes an existing resource quota
*/
func (a *Client) DeleteResourceQuota(params *DeleteResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteResourceQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteResourceQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteResourceQuota",
		Method:             "DELETE",
		PathPattern:        "/api/v1/admin/quotas/{quota_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteResourceQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteResourceQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteResourceQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetResourceQuota gets a specific resource quota
*/
func (a *Client) GetResourceQuota(params *GetResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceQuota",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/quotas/{quota_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetResourceQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListResourceQuotas gets a resource quota list
*/
func (a *Client) ListResourceQuotas(params *ListResourceQuotasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListResourceQuotasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListResourceQuotasParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listResourceQuotas",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListResourceQuotasReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListResourceQuotasOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListResourceQuotasDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateResourceQuota updates an existing resource quota
*/
func (a *Client) UpdateResourceQuota(params *UpdateResourceQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateResourceQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateResourceQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateResourceQuota",
		Method:             "PUT",
		PathPattern:        "/api/v1/admin/quotas/{quota_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateResourceQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateResourceQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateResourceQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
