// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/doubret/terraform-provider-xiiot/models"
)

// DockerProfileGetReader is a Reader for the DockerProfileGet structure.
type DockerProfileGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DockerProfileGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDockerProfileGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDockerProfileGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDockerProfileGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDockerProfileGetOK creates a DockerProfileGetOK with default headers values
func NewDockerProfileGetOK() *DockerProfileGetOK {
	return &DockerProfileGetOK{}
}

/*DockerProfileGetOK handles this case with default header values.

Ok
*/
type DockerProfileGetOK struct {
	Payload *models.DockerProfile
}

func (o *DockerProfileGetOK) Error() string {
	return fmt.Sprintf("[GET /dockerprofiles/{id}][%d] dockerProfileGetOK  %+v", 200, o.Payload)
}

func (o *DockerProfileGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DockerProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDockerProfileGetBadRequest creates a DockerProfileGetBadRequest with default headers values
func NewDockerProfileGetBadRequest() *DockerProfileGetBadRequest {
	return &DockerProfileGetBadRequest{}
}

/*DockerProfileGetBadRequest handles this case with default header values.

Bad request
*/
type DockerProfileGetBadRequest struct {
}

func (o *DockerProfileGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /dockerprofiles/{id}][%d] dockerProfileGetBadRequest ", 400)
}

func (o *DockerProfileGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDockerProfileGetDefault creates a DockerProfileGetDefault with default headers values
func NewDockerProfileGetDefault(code int) *DockerProfileGetDefault {
	return &DockerProfileGetDefault{
		_statusCode: code,
	}
}

/*DockerProfileGetDefault handles this case with default header values.

Unexpected error
*/
type DockerProfileGetDefault struct {
	_statusCode int
}

// Code gets the status code for the docker profile get default response
func (o *DockerProfileGetDefault) Code() int {
	return o._statusCode
}

func (o *DockerProfileGetDefault) Error() string {
	return fmt.Sprintf("[GET /dockerprofiles/{id}][%d] DockerProfileGet default ", o._statusCode)
}

func (o *DockerProfileGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
