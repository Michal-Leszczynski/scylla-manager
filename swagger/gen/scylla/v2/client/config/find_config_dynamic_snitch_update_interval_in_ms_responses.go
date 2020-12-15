// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigDynamicSnitchUpdateIntervalInMsReader is a Reader for the FindConfigDynamicSnitchUpdateIntervalInMs structure.
type FindConfigDynamicSnitchUpdateIntervalInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigDynamicSnitchUpdateIntervalInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigDynamicSnitchUpdateIntervalInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigDynamicSnitchUpdateIntervalInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigDynamicSnitchUpdateIntervalInMsOK creates a FindConfigDynamicSnitchUpdateIntervalInMsOK with default headers values
func NewFindConfigDynamicSnitchUpdateIntervalInMsOK() *FindConfigDynamicSnitchUpdateIntervalInMsOK {
	return &FindConfigDynamicSnitchUpdateIntervalInMsOK{}
}

/*FindConfigDynamicSnitchUpdateIntervalInMsOK handles this case with default header values.

Config value
*/
type FindConfigDynamicSnitchUpdateIntervalInMsOK struct {
	Payload int64
}

func (o *FindConfigDynamicSnitchUpdateIntervalInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigDynamicSnitchUpdateIntervalInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigDynamicSnitchUpdateIntervalInMsDefault creates a FindConfigDynamicSnitchUpdateIntervalInMsDefault with default headers values
func NewFindConfigDynamicSnitchUpdateIntervalInMsDefault(code int) *FindConfigDynamicSnitchUpdateIntervalInMsDefault {
	return &FindConfigDynamicSnitchUpdateIntervalInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigDynamicSnitchUpdateIntervalInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigDynamicSnitchUpdateIntervalInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config dynamic snitch update interval in ms default response
func (o *FindConfigDynamicSnitchUpdateIntervalInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigDynamicSnitchUpdateIntervalInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigDynamicSnitchUpdateIntervalInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigDynamicSnitchUpdateIntervalInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}