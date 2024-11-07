// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new pet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new pet API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new pet API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for pet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV1ProfilesPets(params *GetAPIV1ProfilesPetsParams, opts ...ClientOption) (*GetAPIV1ProfilesPetsOK, error)

	GetAPIV1ProfilesPetsID(params *GetAPIV1ProfilesPetsIDParams, opts ...ClientOption) (*GetAPIV1ProfilesPetsIDOK, error)

	PostAPIV1ProfilesPets(params *PostAPIV1ProfilesPetsParams, opts ...ClientOption) (*PostAPIV1ProfilesPetsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIV1ProfilesPets gets all pet profiles

Get all pet profiles
*/
func (a *Client) GetAPIV1ProfilesPets(params *GetAPIV1ProfilesPetsParams, opts ...ClientOption) (*GetAPIV1ProfilesPetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV1ProfilesPetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV1ProfilesPets",
		Method:             "GET",
		PathPattern:        "/api/v1/profiles/pets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV1ProfilesPetsReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV1ProfilesPetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV1ProfilesPets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV1ProfilesPetsID gets pet profile by ID

Get pet profile by ID
*/
func (a *Client) GetAPIV1ProfilesPetsID(params *GetAPIV1ProfilesPetsIDParams, opts ...ClientOption) (*GetAPIV1ProfilesPetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV1ProfilesPetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV1ProfilesPetsID",
		Method:             "GET",
		PathPattern:        "/api/v1/profiles/pets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV1ProfilesPetsIDReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV1ProfilesPetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV1ProfilesPetsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIV1ProfilesPets creates a new pet profile

Creates a new pet profile
*/
func (a *Client) PostAPIV1ProfilesPets(params *PostAPIV1ProfilesPetsParams, opts ...ClientOption) (*PostAPIV1ProfilesPetsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV1ProfilesPetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV1ProfilesPets",
		Method:             "POST",
		PathPattern:        "/api/v1/profiles/pets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPIV1ProfilesPetsReader{formats: a.formats},
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
	success, ok := result.(*PostAPIV1ProfilesPetsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIV1ProfilesPets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
