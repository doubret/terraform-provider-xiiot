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

// ProjectGetContainerRegistriesReader is a Reader for the ProjectGetContainerRegistries structure.
type ProjectGetContainerRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetContainerRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectGetContainerRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectGetContainerRegistriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewProjectGetContainerRegistriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetContainerRegistriesOK creates a ProjectGetContainerRegistriesOK with default headers values
func NewProjectGetContainerRegistriesOK() *ProjectGetContainerRegistriesOK {
	return &ProjectGetContainerRegistriesOK{}
}

/*ProjectGetContainerRegistriesOK handles this case with default header values.

Ok
*/
type ProjectGetContainerRegistriesOK struct {
	Payload []*models.ContainerRegistry
}

func (o *ProjectGetContainerRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/containerregistries][%d] projectGetContainerRegistriesOK  %+v", 200, o.Payload)
}

func (o *ProjectGetContainerRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetContainerRegistriesBadRequest creates a ProjectGetContainerRegistriesBadRequest with default headers values
func NewProjectGetContainerRegistriesBadRequest() *ProjectGetContainerRegistriesBadRequest {
	return &ProjectGetContainerRegistriesBadRequest{}
}

/*ProjectGetContainerRegistriesBadRequest handles this case with default header values.

Bad request
*/
type ProjectGetContainerRegistriesBadRequest struct {
}

func (o *ProjectGetContainerRegistriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/containerregistries][%d] projectGetContainerRegistriesBadRequest ", 400)
}

func (o *ProjectGetContainerRegistriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectGetContainerRegistriesDefault creates a ProjectGetContainerRegistriesDefault with default headers values
func NewProjectGetContainerRegistriesDefault(code int) *ProjectGetContainerRegistriesDefault {
	return &ProjectGetContainerRegistriesDefault{
		_statusCode: code,
	}
}

/*ProjectGetContainerRegistriesDefault handles this case with default header values.

Unexpected error
*/
type ProjectGetContainerRegistriesDefault struct {
	_statusCode int
}

// Code gets the status code for the project get container registries default response
func (o *ProjectGetContainerRegistriesDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetContainerRegistriesDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/containerregistries][%d] ProjectGetContainerRegistries default ", o._statusCode)
}

func (o *ProjectGetContainerRegistriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
