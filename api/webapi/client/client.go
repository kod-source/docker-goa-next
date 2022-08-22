// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": Client
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=/Users/horikoudai/Documents/ProgrammingLearning/docker-goa-next/api/webapi
// --version=v1.5.13

package client

import (
	goa "github.com/shogo82148/goa-v1"
	goaclient "github.com/shogo82148/goa-v1/client"
)

// Client is the docker_goa_next service client.
type Client struct {
	*goaclient.Client
	JWTSigner goaclient.Signer
	Encoder   *goa.HTTPEncoder
	Decoder   *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")
	client.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetJWTSigner sets the request signer for the jwt security scheme.
func (c *Client) SetJWTSigner(signer goaclient.Signer) {
	c.JWTSigner = signer
}
