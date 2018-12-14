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

// ApplicationDeleteReader is a Reader for the ApplicationDelete structure.
type ApplicationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewApplicationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewApplicationDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewApplicationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationDeleteOK creates a ApplicationDeleteOK with default headers values
func NewApplicationDeleteOK() *ApplicationDeleteOK {
	return &ApplicationDeleteOK{}
}

/*ApplicationDeleteOK handles this case with default header values.

Ok
*/
type ApplicationDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *ApplicationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /application/{id}][%d] applicationDeleteOK  %+v", 200, o.Payload)
}

func (o *ApplicationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationDeleteBadRequest creates a ApplicationDeleteBadRequest with default headers values
func NewApplicationDeleteBadRequest() *ApplicationDeleteBadRequest {
	return &ApplicationDeleteBadRequest{}
}

/*ApplicationDeleteBadRequest handles this case with default header values.

Bad request
*/
type ApplicationDeleteBadRequest struct {
}

func (o *ApplicationDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /application/{id}][%d] applicationDeleteBadRequest ", 400)
}

func (o *ApplicationDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplicationDeleteDefault creates a ApplicationDeleteDefault with default headers values
func NewApplicationDeleteDefault(code int) *ApplicationDeleteDefault {
	return &ApplicationDeleteDefault{
		_statusCode: code,
	}
}

/*ApplicationDeleteDefault handles this case with default header values.

Unexpected error
*/
type ApplicationDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the application delete default response
func (o *ApplicationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /application/{id}][%d] ApplicationDelete default ", o._statusCode)
}

func (o *ApplicationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
