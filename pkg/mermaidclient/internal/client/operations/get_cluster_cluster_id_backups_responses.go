// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/pkg/mermaidclient/internal/models"
)

// GetClusterClusterIDBackupsReader is a Reader for the GetClusterClusterIDBackups structure.
type GetClusterClusterIDBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterClusterIDBackupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClusterClusterIDBackupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterClusterIDBackupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterClusterIDBackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDBackupsOK creates a GetClusterClusterIDBackupsOK with default headers values
func NewGetClusterClusterIDBackupsOK() *GetClusterClusterIDBackupsOK {
	return &GetClusterClusterIDBackupsOK{}
}

/*GetClusterClusterIDBackupsOK handles this case with default header values.

Backup list
*/
type GetClusterClusterIDBackupsOK struct {
	Payload []*models.BackupListItem
}

func (o *GetClusterClusterIDBackupsOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/backups][%d] getClusterClusterIdBackupsOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDBackupsOK) GetPayload() []*models.BackupListItem {
	return o.Payload
}

func (o *GetClusterClusterIDBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDBackupsBadRequest creates a GetClusterClusterIDBackupsBadRequest with default headers values
func NewGetClusterClusterIDBackupsBadRequest() *GetClusterClusterIDBackupsBadRequest {
	return &GetClusterClusterIDBackupsBadRequest{}
}

/*GetClusterClusterIDBackupsBadRequest handles this case with default header values.

Bad Request
*/
type GetClusterClusterIDBackupsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDBackupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/backups][%d] getClusterClusterIdBackupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterClusterIDBackupsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDBackupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDBackupsNotFound creates a GetClusterClusterIDBackupsNotFound with default headers values
func NewGetClusterClusterIDBackupsNotFound() *GetClusterClusterIDBackupsNotFound {
	return &GetClusterClusterIDBackupsNotFound{}
}

/*GetClusterClusterIDBackupsNotFound handles this case with default header values.

Not found
*/
type GetClusterClusterIDBackupsNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDBackupsNotFound) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/backups][%d] getClusterClusterIdBackupsNotFound  %+v", 404, o.Payload)
}

func (o *GetClusterClusterIDBackupsNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDBackupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDBackupsInternalServerError creates a GetClusterClusterIDBackupsInternalServerError with default headers values
func NewGetClusterClusterIDBackupsInternalServerError() *GetClusterClusterIDBackupsInternalServerError {
	return &GetClusterClusterIDBackupsInternalServerError{}
}

/*GetClusterClusterIDBackupsInternalServerError handles this case with default header values.

Server error
*/
type GetClusterClusterIDBackupsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDBackupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/backups][%d] getClusterClusterIdBackupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterClusterIDBackupsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDBackupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDBackupsDefault creates a GetClusterClusterIDBackupsDefault with default headers values
func NewGetClusterClusterIDBackupsDefault(code int) *GetClusterClusterIDBackupsDefault {
	return &GetClusterClusterIDBackupsDefault{
		_statusCode: code,
	}
}

/*GetClusterClusterIDBackupsDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDBackupsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID backups default response
func (o *GetClusterClusterIDBackupsDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDBackupsDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/backups][%d] GetClusterClusterIDBackups default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDBackupsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDBackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}