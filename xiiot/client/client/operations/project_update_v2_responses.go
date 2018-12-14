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

// ProjectUpdateV2Reader is a Reader for the ProjectUpdateV2 structure.
type ProjectUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectUpdateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewProjectUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectUpdateV2OK creates a ProjectUpdateV2OK with default headers values
func NewProjectUpdateV2OK() *ProjectUpdateV2OK {
	return &ProjectUpdateV2OK{}
}

/*ProjectUpdateV2OK handles this case with default header values.

Ok
*/
type ProjectUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *ProjectUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] projectUpdateV2OK  %+v", 200, o.Payload)
}

func (o *ProjectUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectUpdateV2BadRequest creates a ProjectUpdateV2BadRequest with default headers values
func NewProjectUpdateV2BadRequest() *ProjectUpdateV2BadRequest {
	return &ProjectUpdateV2BadRequest{}
}

/*ProjectUpdateV2BadRequest handles this case with default header values.

Bad request
*/
type ProjectUpdateV2BadRequest struct {
}

func (o *ProjectUpdateV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] projectUpdateV2BadRequest ", 400)
}

func (o *ProjectUpdateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectUpdateV2Default creates a ProjectUpdateV2Default with default headers values
func NewProjectUpdateV2Default(code int) *ProjectUpdateV2Default {
	return &ProjectUpdateV2Default{
		_statusCode: code,
	}
}

/*ProjectUpdateV2Default handles this case with default header values.

Unexpected error
*/
type ProjectUpdateV2Default struct {
	_statusCode int
}

// Code gets the status code for the project update v2 default response
func (o *ProjectUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *ProjectUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] ProjectUpdateV2 default ", o._statusCode)
}

func (o *ProjectUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
