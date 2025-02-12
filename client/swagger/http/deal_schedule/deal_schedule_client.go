// Code generated by go-swagger; DO NOT EDIT.

package deal_schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deal schedule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deal schedule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSchedule(params *CreateScheduleParams, opts ...ClientOption) (*CreateScheduleOK, error)

	ListPreparationSchedules(params *ListPreparationSchedulesParams, opts ...ClientOption) (*ListPreparationSchedulesOK, error)

	ListSchedules(params *ListSchedulesParams, opts ...ClientOption) (*ListSchedulesOK, error)

	PauseSchedule(params *PauseScheduleParams, opts ...ClientOption) (*PauseScheduleOK, error)

	ResumeSchedule(params *ResumeScheduleParams, opts ...ClientOption) (*ResumeScheduleOK, error)

	UpdateSchedule(params *UpdateScheduleParams, opts ...ClientOption) (*UpdateScheduleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSchedule creates a new schedule

Create a new schedule
*/
func (a *Client) CreateSchedule(params *CreateScheduleParams, opts ...ClientOption) (*CreateScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSchedule",
		Method:             "POST",
		PathPattern:        "/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateScheduleReader{formats: a.formats},
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
	success, ok := result.(*CreateScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListPreparationSchedules lists all schedules for a preparation
*/
func (a *Client) ListPreparationSchedules(params *ListPreparationSchedulesParams, opts ...ClientOption) (*ListPreparationSchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPreparationSchedulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListPreparationSchedules",
		Method:             "GET",
		PathPattern:        "/preparation/{id}/schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPreparationSchedulesReader{formats: a.formats},
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
	success, ok := result.(*ListPreparationSchedulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListPreparationSchedules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSchedules lists all deal making schedules
*/
func (a *Client) ListSchedules(params *ListSchedulesParams, opts ...ClientOption) (*ListSchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSchedulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListSchedules",
		Method:             "GET",
		PathPattern:        "/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSchedulesReader{formats: a.formats},
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
	success, ok := result.(*ListSchedulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListSchedules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PauseSchedule pauses a specific schedule
*/
func (a *Client) PauseSchedule(params *PauseScheduleParams, opts ...ClientOption) (*PauseScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPauseScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PauseSchedule",
		Method:             "POST",
		PathPattern:        "/schedule/{id}/pause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PauseScheduleReader{formats: a.formats},
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
	success, ok := result.(*PauseScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PauseSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResumeSchedule resumes a specific schedule
*/
func (a *Client) ResumeSchedule(params *ResumeScheduleParams, opts ...ClientOption) (*ResumeScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResumeScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResumeSchedule",
		Method:             "POST",
		PathPattern:        "/schedule/{id}/resume",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResumeScheduleReader{formats: a.formats},
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
	success, ok := result.(*ResumeScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResumeSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSchedule updates a schedule

Update a schedule
*/
func (a *Client) UpdateSchedule(params *UpdateScheduleParams, opts ...ClientOption) (*UpdateScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSchedule",
		Method:             "PATCH",
		PathPattern:        "/schedule/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateScheduleReader{formats: a.formats},
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
	success, ok := result.(*UpdateScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
