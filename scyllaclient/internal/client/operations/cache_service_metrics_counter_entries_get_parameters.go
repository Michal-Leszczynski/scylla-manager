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

// NewCacheServiceMetricsCounterEntriesGetParams creates a new CacheServiceMetricsCounterEntriesGetParams object
// with the default values initialized.
func NewCacheServiceMetricsCounterEntriesGetParams() *CacheServiceMetricsCounterEntriesGetParams {

	return &CacheServiceMetricsCounterEntriesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsCounterEntriesGetParamsWithTimeout creates a new CacheServiceMetricsCounterEntriesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsCounterEntriesGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsCounterEntriesGetParams {

	return &CacheServiceMetricsCounterEntriesGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsCounterEntriesGetParamsWithContext creates a new CacheServiceMetricsCounterEntriesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsCounterEntriesGetParamsWithContext(ctx context.Context) *CacheServiceMetricsCounterEntriesGetParams {

	return &CacheServiceMetricsCounterEntriesGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsCounterEntriesGetParamsWithHTTPClient creates a new CacheServiceMetricsCounterEntriesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsCounterEntriesGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsCounterEntriesGetParams {

	return &CacheServiceMetricsCounterEntriesGetParams{
		HTTPClient: client,
	}
}

/*CacheServiceMetricsCounterEntriesGetParams contains all the parameters to send to the API endpoint
for the cache service metrics counter entries get operation typically these are written to a http.Request
*/
type CacheServiceMetricsCounterEntriesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics counter entries get params
func (o *CacheServiceMetricsCounterEntriesGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsCounterEntriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics counter entries get params
func (o *CacheServiceMetricsCounterEntriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics counter entries get params
func (o *CacheServiceMetricsCounterEntriesGetParams) WithContext(ctx context.Context) *CacheServiceMetricsCounterEntriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics counter entries get params
func (o *CacheServiceMetricsCounterEntriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics counter entries get params
func (o *CacheServiceMetricsCounterEntriesGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsCounterEntriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics counter entries get params
func (o *CacheServiceMetricsCounterEntriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsCounterEntriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}