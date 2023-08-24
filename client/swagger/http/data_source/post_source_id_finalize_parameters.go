// Code generated by go-swagger; DO NOT EDIT.

package data_source

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
	"github.com/go-openapi/swag"
)

// NewPostSourceIDFinalizeParams creates a new PostSourceIDFinalizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceIDFinalizeParams() *PostSourceIDFinalizeParams {
	return &PostSourceIDFinalizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceIDFinalizeParamsWithTimeout creates a new PostSourceIDFinalizeParams object
// with the ability to set a timeout on a request.
func NewPostSourceIDFinalizeParamsWithTimeout(timeout time.Duration) *PostSourceIDFinalizeParams {
	return &PostSourceIDFinalizeParams{
		timeout: timeout,
	}
}

// NewPostSourceIDFinalizeParamsWithContext creates a new PostSourceIDFinalizeParams object
// with the ability to set a context for a request.
func NewPostSourceIDFinalizeParamsWithContext(ctx context.Context) *PostSourceIDFinalizeParams {
	return &PostSourceIDFinalizeParams{
		Context: ctx,
	}
}

// NewPostSourceIDFinalizeParamsWithHTTPClient creates a new PostSourceIDFinalizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceIDFinalizeParamsWithHTTPClient(client *http.Client) *PostSourceIDFinalizeParams {
	return &PostSourceIDFinalizeParams{
		HTTPClient: client,
	}
}

/*
PostSourceIDFinalizeParams contains all the parameters to send to the API endpoint

	for the post source ID finalize operation.

	Typically these are written to a http.Request.
*/
type PostSourceIDFinalizeParams struct {

	/* ID.

	   Source ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source ID finalize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceIDFinalizeParams) WithDefaults() *PostSourceIDFinalizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source ID finalize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceIDFinalizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) WithTimeout(timeout time.Duration) *PostSourceIDFinalizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) WithContext(ctx context.Context) *PostSourceIDFinalizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) WithHTTPClient(client *http.Client) *PostSourceIDFinalizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) WithID(id int64) *PostSourceIDFinalizeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post source ID finalize params
func (o *PostSourceIDFinalizeParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceIDFinalizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}