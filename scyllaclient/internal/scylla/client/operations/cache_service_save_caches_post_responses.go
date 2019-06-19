// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CacheServiceSaveCachesPostReader is a Reader for the CacheServiceSaveCachesPost structure.
type CacheServiceSaveCachesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceSaveCachesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCacheServiceSaveCachesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceSaveCachesPostOK creates a CacheServiceSaveCachesPostOK with default headers values
func NewCacheServiceSaveCachesPostOK() *CacheServiceSaveCachesPostOK {
	return &CacheServiceSaveCachesPostOK{}
}

/*CacheServiceSaveCachesPostOK handles this case with default header values.

CacheServiceSaveCachesPostOK cache service save caches post o k
*/
type CacheServiceSaveCachesPostOK struct {
}

func (o *CacheServiceSaveCachesPostOK) Error() string {
	return fmt.Sprintf("[POST /cache_service/save_caches][%d] cacheServiceSaveCachesPostOK ", 200)
}

func (o *CacheServiceSaveCachesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}