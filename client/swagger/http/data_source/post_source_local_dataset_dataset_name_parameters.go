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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewPostSourceLocalDatasetDatasetNameParams creates a new PostSourceLocalDatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceLocalDatasetDatasetNameParams() *PostSourceLocalDatasetDatasetNameParams {
	return &PostSourceLocalDatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceLocalDatasetDatasetNameParamsWithTimeout creates a new PostSourceLocalDatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceLocalDatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceLocalDatasetDatasetNameParams {
	return &PostSourceLocalDatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceLocalDatasetDatasetNameParamsWithContext creates a new PostSourceLocalDatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceLocalDatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceLocalDatasetDatasetNameParams {
	return &PostSourceLocalDatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceLocalDatasetDatasetNameParamsWithHTTPClient creates a new PostSourceLocalDatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceLocalDatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceLocalDatasetDatasetNameParams {
	return &PostSourceLocalDatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceLocalDatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source local dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceLocalDatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceLocalRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source local dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceLocalDatasetDatasetNameParams) WithDefaults() *PostSourceLocalDatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source local dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceLocalDatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceLocalDatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceLocalDatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceLocalDatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceLocalDatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) WithRequest(request *models.DatasourceLocalRequest) *PostSourceLocalDatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source local dataset dataset name params
func (o *PostSourceLocalDatasetDatasetNameParams) SetRequest(request *models.DatasourceLocalRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceLocalDatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasetName
	if err := r.SetPathParam("datasetName", o.DatasetName); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}