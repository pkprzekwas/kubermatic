// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetMeteringReportConfigurationParams creates a new GetMeteringReportConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMeteringReportConfigurationParams() *GetMeteringReportConfigurationParams {
	return &GetMeteringReportConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeteringReportConfigurationParamsWithTimeout creates a new GetMeteringReportConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetMeteringReportConfigurationParamsWithTimeout(timeout time.Duration) *GetMeteringReportConfigurationParams {
	return &GetMeteringReportConfigurationParams{
		timeout: timeout,
	}
}

// NewGetMeteringReportConfigurationParamsWithContext creates a new GetMeteringReportConfigurationParams object
// with the ability to set a context for a request.
func NewGetMeteringReportConfigurationParamsWithContext(ctx context.Context) *GetMeteringReportConfigurationParams {
	return &GetMeteringReportConfigurationParams{
		Context: ctx,
	}
}

// NewGetMeteringReportConfigurationParamsWithHTTPClient creates a new GetMeteringReportConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMeteringReportConfigurationParamsWithHTTPClient(client *http.Client) *GetMeteringReportConfigurationParams {
	return &GetMeteringReportConfigurationParams{
		HTTPClient: client,
	}
}

/* GetMeteringReportConfigurationParams contains all the parameters to send to the API endpoint
   for the get metering report configuration operation.

   Typically these are written to a http.Request.
*/
type GetMeteringReportConfigurationParams struct {

	// ReportConfigurationName.
	ReportConfigurationName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get metering report configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMeteringReportConfigurationParams) WithDefaults() *GetMeteringReportConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get metering report configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMeteringReportConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) WithTimeout(timeout time.Duration) *GetMeteringReportConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) WithContext(ctx context.Context) *GetMeteringReportConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) WithHTTPClient(client *http.Client) *GetMeteringReportConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportConfigurationName adds the reportConfigurationName to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) WithReportConfigurationName(reportConfigurationName string) *GetMeteringReportConfigurationParams {
	o.SetReportConfigurationName(reportConfigurationName)
	return o
}

// SetReportConfigurationName adds the reportConfigurationName to the get metering report configuration params
func (o *GetMeteringReportConfigurationParams) SetReportConfigurationName(reportConfigurationName string) {
	o.ReportConfigurationName = reportConfigurationName
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeteringReportConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param report_configuration_name
	if err := r.SetPathParam("report_configuration_name", o.ReportConfigurationName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
