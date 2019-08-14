// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigOverrideDecommissionReader is a Reader for the FindConfigOverrideDecommission structure.
type FindConfigOverrideDecommissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigOverrideDecommissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigOverrideDecommissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigOverrideDecommissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigOverrideDecommissionOK creates a FindConfigOverrideDecommissionOK with default headers values
func NewFindConfigOverrideDecommissionOK() *FindConfigOverrideDecommissionOK {
	return &FindConfigOverrideDecommissionOK{}
}

/*FindConfigOverrideDecommissionOK handles this case with default header values.

Config value
*/
type FindConfigOverrideDecommissionOK struct {
	Payload bool
}

func (o *FindConfigOverrideDecommissionOK) Error() string {
	return fmt.Sprintf("[GET /config/override_decommission][%d] findConfigOverrideDecommissionOK  %+v", 200, o.Payload)
}

func (o *FindConfigOverrideDecommissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigOverrideDecommissionDefault creates a FindConfigOverrideDecommissionDefault with default headers values
func NewFindConfigOverrideDecommissionDefault(code int) *FindConfigOverrideDecommissionDefault {
	return &FindConfigOverrideDecommissionDefault{
		_statusCode: code,
	}
}

/*FindConfigOverrideDecommissionDefault handles this case with default header values.

unexpected error
*/
type FindConfigOverrideDecommissionDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config override decommission default response
func (o *FindConfigOverrideDecommissionDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigOverrideDecommissionDefault) Error() string {
	return fmt.Sprintf("[GET /config/override_decommission][%d] find_config_override_decommission default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigOverrideDecommissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}