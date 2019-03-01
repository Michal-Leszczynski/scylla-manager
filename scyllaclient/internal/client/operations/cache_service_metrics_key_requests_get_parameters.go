// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCacheServiceMetricsKeyRequestsGetParams creates a new CacheServiceMetricsKeyRequestsGetParams object
// with the default values initialized.
func NewCacheServiceMetricsKeyRequestsGetParams() *CacheServiceMetricsKeyRequestsGetParams {

	return &CacheServiceMetricsKeyRequestsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsKeyRequestsGetParamsWithTimeout creates a new CacheServiceMetricsKeyRequestsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsKeyRequestsGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsKeyRequestsGetParams {

	return &CacheServiceMetricsKeyRequestsGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsKeyRequestsGetParamsWithContext creates a new CacheServiceMetricsKeyRequestsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsKeyRequestsGetParamsWithContext(ctx context.Context) *CacheServiceMetricsKeyRequestsGetParams {

	return &CacheServiceMetricsKeyRequestsGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsKeyRequestsGetParamsWithHTTPClient creates a new CacheServiceMetricsKeyRequestsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsKeyRequestsGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsKeyRequestsGetParams {

	return &CacheServiceMetricsKeyRequestsGetParams{
		HTTPClient: client,
	}
}

/*CacheServiceMetricsKeyRequestsGetParams contains all the parameters to send to the API endpoint
for the cache service metrics key requests get operation typically these are written to a http.Request
*/
type CacheServiceMetricsKeyRequestsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics key requests get params
func (o *CacheServiceMetricsKeyRequestsGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsKeyRequestsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics key requests get params
func (o *CacheServiceMetricsKeyRequestsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics key requests get params
func (o *CacheServiceMetricsKeyRequestsGetParams) WithContext(ctx context.Context) *CacheServiceMetricsKeyRequestsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics key requests get params
func (o *CacheServiceMetricsKeyRequestsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics key requests get params
func (o *CacheServiceMetricsKeyRequestsGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsKeyRequestsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics key requests get params
func (o *CacheServiceMetricsKeyRequestsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsKeyRequestsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}