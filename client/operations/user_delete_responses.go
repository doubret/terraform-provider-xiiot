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

// UserDeleteReader is a Reader for the UserDelete structure.
type UserDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUserDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserDeleteOK creates a UserDeleteOK with default headers values
func NewUserDeleteOK() *UserDeleteOK {
	return &UserDeleteOK{}
}

/*UserDeleteOK handles this case with default header values.

Ok
*/
type UserDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *UserDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}][%d] userDeleteOK  %+v", 200, o.Payload)
}

func (o *UserDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserDeleteBadRequest creates a UserDeleteBadRequest with default headers values
func NewUserDeleteBadRequest() *UserDeleteBadRequest {
	return &UserDeleteBadRequest{}
}

/*UserDeleteBadRequest handles this case with default header values.

Bad request
*/
type UserDeleteBadRequest struct {
}

func (o *UserDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}][%d] userDeleteBadRequest ", 400)
}

func (o *UserDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserDeleteDefault creates a UserDeleteDefault with default headers values
func NewUserDeleteDefault(code int) *UserDeleteDefault {
	return &UserDeleteDefault{
		_statusCode: code,
	}
}

/*UserDeleteDefault handles this case with default header values.

Unexpected error
*/
type UserDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the user delete default response
func (o *UserDeleteDefault) Code() int {
	return o._statusCode
}

func (o *UserDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/{id}][%d] UserDelete default ", o._statusCode)
}

func (o *UserDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
