// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/pkg/scyllaclient/internal/scylla/models"
)

// StorageServiceNativeTransportDeleteReader is a Reader for the StorageServiceNativeTransportDelete structure.
type StorageServiceNativeTransportDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceNativeTransportDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceNativeTransportDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceNativeTransportDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceNativeTransportDeleteOK creates a StorageServiceNativeTransportDeleteOK with default headers values
func NewStorageServiceNativeTransportDeleteOK() *StorageServiceNativeTransportDeleteOK {
	return &StorageServiceNativeTransportDeleteOK{}
}

/*StorageServiceNativeTransportDeleteOK handles this case with default header values.

StorageServiceNativeTransportDeleteOK storage service native transport delete o k
*/
type StorageServiceNativeTransportDeleteOK struct {
}

func (o *StorageServiceNativeTransportDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceNativeTransportDeleteDefault creates a StorageServiceNativeTransportDeleteDefault with default headers values
func NewStorageServiceNativeTransportDeleteDefault(code int) *StorageServiceNativeTransportDeleteDefault {
	return &StorageServiceNativeTransportDeleteDefault{
		_statusCode: code,
	}
}

/*StorageServiceNativeTransportDeleteDefault handles this case with default header values.

internal server error
*/
type StorageServiceNativeTransportDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service native transport delete default response
func (o *StorageServiceNativeTransportDeleteDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceNativeTransportDeleteDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceNativeTransportDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceNativeTransportDeleteDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}