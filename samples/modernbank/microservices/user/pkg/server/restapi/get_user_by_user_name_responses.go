// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/user/pkg/model"
)

// GetUserByUserNameOKCode is the HTTP code returned for type GetUserByUserNameOK
const GetUserByUserNameOKCode int = 200

/*GetUserByUserNameOK Success!

swagger:response getUserByUserNameOK
*/
type GetUserByUserNameOK struct {

	/*
	  In: Body
	*/
	Payload *model.User `json:"body,omitempty"`
}

// NewGetUserByUserNameOK creates GetUserByUserNameOK with default headers values
func NewGetUserByUserNameOK() *GetUserByUserNameOK {

	return &GetUserByUserNameOK{}
}

// WithPayload adds the payload to the get user by user name o k response
func (o *GetUserByUserNameOK) WithPayload(payload *model.User) *GetUserByUserNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by user name o k response
func (o *GetUserByUserNameOK) SetPayload(payload *model.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByUserNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByUserNameNotFoundCode is the HTTP code returned for type GetUserByUserNameNotFound
const GetUserByUserNameNotFoundCode int = 404

/*GetUserByUserNameNotFound User not found

swagger:response getUserByUserNameNotFound
*/
type GetUserByUserNameNotFound struct {
}

// NewGetUserByUserNameNotFound creates GetUserByUserNameNotFound with default headers values
func NewGetUserByUserNameNotFound() *GetUserByUserNameNotFound {

	return &GetUserByUserNameNotFound{}
}

// WriteResponse to the client
func (o *GetUserByUserNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetUserByUserNameInternalServerErrorCode is the HTTP code returned for type GetUserByUserNameInternalServerError
const GetUserByUserNameInternalServerErrorCode int = 500

/*GetUserByUserNameInternalServerError Internal server error

swagger:response getUserByUserNameInternalServerError
*/
type GetUserByUserNameInternalServerError struct {
}

// NewGetUserByUserNameInternalServerError creates GetUserByUserNameInternalServerError with default headers values
func NewGetUserByUserNameInternalServerError() *GetUserByUserNameInternalServerError {

	return &GetUserByUserNameInternalServerError{}
}

// WriteResponse to the client
func (o *GetUserByUserNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
