// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetReader is a Reader for the ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGet structure.
type ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK creates a ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK with default headers values
func NewColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK() *ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK {
	return &ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK{}
}

/*ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK column family metrics row cache hit out of range by name get o k
*/
type ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/row_cache_hit_out_of_range/{name}][%d] columnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsRowCacheHitOutOfRangeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}