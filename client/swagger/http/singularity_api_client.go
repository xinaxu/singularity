// Code generated by go-swagger; DO NOT EDIT.

package http

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/http/deal"
	"github.com/data-preservation-programs/singularity/client/swagger/http/deal_schedule"
	"github.com/data-preservation-programs/singularity/client/swagger/http/file"
	"github.com/data-preservation-programs/singularity/client/swagger/http/job"
	"github.com/data-preservation-programs/singularity/client/swagger/http/piece"
	"github.com/data-preservation-programs/singularity/client/swagger/http/preparation"
	"github.com/data-preservation-programs/singularity/client/swagger/http/storage"
	"github.com/data-preservation-programs/singularity/client/swagger/http/wallet"
	"github.com/data-preservation-programs/singularity/client/swagger/http/wallet_association"
)

// Default singularity API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:9090"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new singularity API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *SingularityAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new singularity API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *SingularityAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new singularity API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *SingularityAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(SingularityAPI)
	cli.Transport = transport
	cli.Deal = deal.New(transport, formats)
	cli.DealSchedule = deal_schedule.New(transport, formats)
	cli.File = file.New(transport, formats)
	cli.Job = job.New(transport, formats)
	cli.Piece = piece.New(transport, formats)
	cli.Preparation = preparation.New(transport, formats)
	cli.Storage = storage.New(transport, formats)
	cli.Wallet = wallet.New(transport, formats)
	cli.WalletAssociation = wallet_association.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// SingularityAPI is a client for singularity API
type SingularityAPI struct {
	Deal deal.ClientService

	DealSchedule deal_schedule.ClientService

	File file.ClientService

	Job job.ClientService

	Piece piece.ClientService

	Preparation preparation.ClientService

	Storage storage.ClientService

	Wallet wallet.ClientService

	WalletAssociation wallet_association.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *SingularityAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Deal.SetTransport(transport)
	c.DealSchedule.SetTransport(transport)
	c.File.SetTransport(transport)
	c.Job.SetTransport(transport)
	c.Piece.SetTransport(transport)
	c.Preparation.SetTransport(transport)
	c.Storage.SetTransport(transport)
	c.Wallet.SetTransport(transport)
	c.WalletAssociation.SetTransport(transport)
}
