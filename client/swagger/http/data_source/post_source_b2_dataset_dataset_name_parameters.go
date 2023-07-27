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

// NewPostSourceB2DatasetDatasetNameParams creates a new PostSourceB2DatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceB2DatasetDatasetNameParams() *PostSourceB2DatasetDatasetNameParams {
	return &PostSourceB2DatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceB2DatasetDatasetNameParamsWithTimeout creates a new PostSourceB2DatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceB2DatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceB2DatasetDatasetNameParams {
	return &PostSourceB2DatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceB2DatasetDatasetNameParamsWithContext creates a new PostSourceB2DatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceB2DatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceB2DatasetDatasetNameParams {
	return &PostSourceB2DatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceB2DatasetDatasetNameParamsWithHTTPClient creates a new PostSourceB2DatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceB2DatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceB2DatasetDatasetNameParams {
	return &PostSourceB2DatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceB2DatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source b2 dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceB2DatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceB2Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source b2 dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceB2DatasetDatasetNameParams) WithDefaults() *PostSourceB2DatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source b2 dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceB2DatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceB2DatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceB2DatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceB2DatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceB2DatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) WithRequest(request *models.DatasourceB2Request) *PostSourceB2DatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source b2 dataset dataset name params
func (o *PostSourceB2DatasetDatasetNameParams) SetRequest(request *models.DatasourceB2Request) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceB2DatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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