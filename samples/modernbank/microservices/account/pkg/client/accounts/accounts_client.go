// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateAccount creates a new account for a customer

Creates a new account for a given customer
*/
func (a *Client) CreateAccount(params *CreateAccountParams) (*CreateAccountCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAccount",
		Method:             "POST",
		PathPattern:        "/users/{owner}/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAccountCreated), nil

}

/*
DeleteAccount deletes account by account number

Delete account by account number.
*/
func (a *Client) DeleteAccount(params *DeleteAccountParams) (*DeleteAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAccount",
		Method:             "DELETE",
		PathPattern:        "/users/{owner}/accounts/{number}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAccountOK), nil

}

/*
GetAccountByNumber gets account details
*/
func (a *Client) GetAccountByNumber(params *GetAccountByNumberParams) (*GetAccountByNumberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountByNumberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccountByNumber",
		Method:             "GET",
		PathPattern:        "/users/{owner}/accounts/{number}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAccountByNumberReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountByNumberOK), nil

}

/*
ListAccounts lists all accounts for a given customer

Lists all accounts for a given customer
*/
func (a *Client) ListAccounts(params *ListAccountsParams) (*ListAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAccounts",
		Method:             "GET",
		PathPattern:        "/users/{owner}/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAccountsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
