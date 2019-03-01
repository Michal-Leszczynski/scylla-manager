// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageProxyMetricsWriteEstimatedHistogramGetReader is a Reader for the StorageProxyMetricsWriteEstimatedHistogramGet structure.
type StorageProxyMetricsWriteEstimatedHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsWriteEstimatedHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageProxyMetricsWriteEstimatedHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyMetricsWriteEstimatedHistogramGetOK creates a StorageProxyMetricsWriteEstimatedHistogramGetOK with default headers values
func NewStorageProxyMetricsWriteEstimatedHistogramGetOK() *StorageProxyMetricsWriteEstimatedHistogramGetOK {
	return &StorageProxyMetricsWriteEstimatedHistogramGetOK{}
}

/*StorageProxyMetricsWriteEstimatedHistogramGetOK handles this case with default header values.

StorageProxyMetricsWriteEstimatedHistogramGetOK storage proxy metrics write estimated histogram get o k
*/
type StorageProxyMetricsWriteEstimatedHistogramGetOK struct {
}

func (o *StorageProxyMetricsWriteEstimatedHistogramGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/write/estimated_histogram/][%d] storageProxyMetricsWriteEstimatedHistogramGetOK ", 200)
}

func (o *StorageProxyMetricsWriteEstimatedHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}