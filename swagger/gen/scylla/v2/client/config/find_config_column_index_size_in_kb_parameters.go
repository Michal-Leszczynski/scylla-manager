// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigColumnIndexSizeInKbParams creates a new FindConfigColumnIndexSizeInKbParams object
// with the default values initialized.
func NewFindConfigColumnIndexSizeInKbParams() *FindConfigColumnIndexSizeInKbParams {

	return &FindConfigColumnIndexSizeInKbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigColumnIndexSizeInKbParamsWithTimeout creates a new FindConfigColumnIndexSizeInKbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigColumnIndexSizeInKbParamsWithTimeout(timeout time.Duration) *FindConfigColumnIndexSizeInKbParams {

	return &FindConfigColumnIndexSizeInKbParams{

		timeout: timeout,
	}
}

// NewFindConfigColumnIndexSizeInKbParamsWithContext creates a new FindConfigColumnIndexSizeInKbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigColumnIndexSizeInKbParamsWithContext(ctx context.Context) *FindConfigColumnIndexSizeInKbParams {

	return &FindConfigColumnIndexSizeInKbParams{

		Context: ctx,
	}
}

// NewFindConfigColumnIndexSizeInKbParamsWithHTTPClient creates a new FindConfigColumnIndexSizeInKbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigColumnIndexSizeInKbParamsWithHTTPClient(client *http.Client) *FindConfigColumnIndexSizeInKbParams {

	return &FindConfigColumnIndexSizeInKbParams{
		HTTPClient: client,
	}
}

/*FindConfigColumnIndexSizeInKbParams contains all the parameters to send to the API endpoint
for the find config column index size in kb operation typically these are written to a http.Request
*/
type FindConfigColumnIndexSizeInKbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config column index size in kb params
func (o *FindConfigColumnIndexSizeInKbParams) WithTimeout(timeout time.Duration) *FindConfigColumnIndexSizeInKbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config column index size in kb params
func (o *FindConfigColumnIndexSizeInKbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config column index size in kb params
func (o *FindConfigColumnIndexSizeInKbParams) WithContext(ctx context.Context) *FindConfigColumnIndexSizeInKbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config column index size in kb params
func (o *FindConfigColumnIndexSizeInKbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config column index size in kb params
func (o *FindConfigColumnIndexSizeInKbParams) WithHTTPClient(client *http.Client) *FindConfigColumnIndexSizeInKbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config column index size in kb params
func (o *FindConfigColumnIndexSizeInKbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigColumnIndexSizeInKbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}