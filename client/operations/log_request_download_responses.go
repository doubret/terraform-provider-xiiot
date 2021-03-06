// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// LogRequestDownloadReader is a Reader for the LogRequestDownload structure.
type LogRequestDownloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogRequestDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLogRequestDownloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLogRequestDownloadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewLogRequestDownloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogRequestDownloadOK creates a LogRequestDownloadOK with default headers values
func NewLogRequestDownloadOK() *LogRequestDownloadOK {
	return &LogRequestDownloadOK{}
}

/*LogRequestDownloadOK handles this case with default header values.

Ok
*/
type LogRequestDownloadOK struct {
}

func (o *LogRequestDownloadOK) Error() string {
	return fmt.Sprintf("[POST /logs/requestDownload][%d] logRequestDownloadOK ", 200)
}

func (o *LogRequestDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogRequestDownloadBadRequest creates a LogRequestDownloadBadRequest with default headers values
func NewLogRequestDownloadBadRequest() *LogRequestDownloadBadRequest {
	return &LogRequestDownloadBadRequest{}
}

/*LogRequestDownloadBadRequest handles this case with default header values.

Bad request
*/
type LogRequestDownloadBadRequest struct {
}

func (o *LogRequestDownloadBadRequest) Error() string {
	return fmt.Sprintf("[POST /logs/requestDownload][%d] logRequestDownloadBadRequest ", 400)
}

func (o *LogRequestDownloadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogRequestDownloadDefault creates a LogRequestDownloadDefault with default headers values
func NewLogRequestDownloadDefault(code int) *LogRequestDownloadDefault {
	return &LogRequestDownloadDefault{
		_statusCode: code,
	}
}

/*LogRequestDownloadDefault handles this case with default header values.

Unexpected error
*/
type LogRequestDownloadDefault struct {
	_statusCode int
}

// Code gets the status code for the log request download default response
func (o *LogRequestDownloadDefault) Code() int {
	return o._statusCode
}

func (o *LogRequestDownloadDefault) Error() string {
	return fmt.Sprintf("[POST /logs/requestDownload][%d] LogRequestDownload default ", o._statusCode)
}

func (o *LogRequestDownloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
