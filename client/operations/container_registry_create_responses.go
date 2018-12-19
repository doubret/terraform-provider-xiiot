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

// ContainerRegistryCreateReader is a Reader for the ContainerRegistryCreate structure.
type ContainerRegistryCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRegistryCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerRegistryCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewContainerRegistryCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewContainerRegistryCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContainerRegistryCreateOK creates a ContainerRegistryCreateOK with default headers values
func NewContainerRegistryCreateOK() *ContainerRegistryCreateOK {
	return &ContainerRegistryCreateOK{}
}

/*ContainerRegistryCreateOK handles this case with default header values.

Ok
*/
type ContainerRegistryCreateOK struct {
	Payload *models.CreateDocumentResponse
}

func (o *ContainerRegistryCreateOK) Error() string {
	return fmt.Sprintf("[POST /containerregistries][%d] containerRegistryCreateOK  %+v", 200, o.Payload)
}

func (o *ContainerRegistryCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRegistryCreateBadRequest creates a ContainerRegistryCreateBadRequest with default headers values
func NewContainerRegistryCreateBadRequest() *ContainerRegistryCreateBadRequest {
	return &ContainerRegistryCreateBadRequest{}
}

/*ContainerRegistryCreateBadRequest handles this case with default header values.

Bad request
*/
type ContainerRegistryCreateBadRequest struct {
}

func (o *ContainerRegistryCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /containerregistries][%d] containerRegistryCreateBadRequest ", 400)
}

func (o *ContainerRegistryCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRegistryCreateDefault creates a ContainerRegistryCreateDefault with default headers values
func NewContainerRegistryCreateDefault(code int) *ContainerRegistryCreateDefault {
	return &ContainerRegistryCreateDefault{
		_statusCode: code,
	}
}

/*ContainerRegistryCreateDefault handles this case with default header values.

Unexpected error
*/
type ContainerRegistryCreateDefault struct {
	_statusCode int
}

// Code gets the status code for the container registry create default response
func (o *ContainerRegistryCreateDefault) Code() int {
	return o._statusCode
}

func (o *ContainerRegistryCreateDefault) Error() string {
	return fmt.Sprintf("[POST /containerregistries][%d] ContainerRegistryCreate default ", o._statusCode)
}

func (o *ContainerRegistryCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
