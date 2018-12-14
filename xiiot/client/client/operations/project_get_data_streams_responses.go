// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
)

// ProjectGetDataStreamsReader is a Reader for the ProjectGetDataStreams structure.
type ProjectGetDataStreamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetDataStreamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectGetDataStreamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectGetDataStreamsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewProjectGetDataStreamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetDataStreamsOK creates a ProjectGetDataStreamsOK with default headers values
func NewProjectGetDataStreamsOK() *ProjectGetDataStreamsOK {
	return &ProjectGetDataStreamsOK{}
}

/*ProjectGetDataStreamsOK handles this case with default header values.

Ok
*/
type ProjectGetDataStreamsOK struct {
	Payload []*models.DataStream
}

func (o *ProjectGetDataStreamsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/datastreams][%d] projectGetDataStreamsOK  %+v", 200, o.Payload)
}

func (o *ProjectGetDataStreamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetDataStreamsBadRequest creates a ProjectGetDataStreamsBadRequest with default headers values
func NewProjectGetDataStreamsBadRequest() *ProjectGetDataStreamsBadRequest {
	return &ProjectGetDataStreamsBadRequest{}
}

/*ProjectGetDataStreamsBadRequest handles this case with default header values.

Bad request
*/
type ProjectGetDataStreamsBadRequest struct {
}

func (o *ProjectGetDataStreamsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/datastreams][%d] projectGetDataStreamsBadRequest ", 400)
}

func (o *ProjectGetDataStreamsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectGetDataStreamsDefault creates a ProjectGetDataStreamsDefault with default headers values
func NewProjectGetDataStreamsDefault(code int) *ProjectGetDataStreamsDefault {
	return &ProjectGetDataStreamsDefault{
		_statusCode: code,
	}
}

/*ProjectGetDataStreamsDefault handles this case with default header values.

Unexpected error
*/
type ProjectGetDataStreamsDefault struct {
	_statusCode int
}

// Code gets the status code for the project get data streams default response
func (o *ProjectGetDataStreamsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetDataStreamsDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/datastreams][%d] ProjectGetDataStreams default ", o._statusCode)
}

func (o *ProjectGetDataStreamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
