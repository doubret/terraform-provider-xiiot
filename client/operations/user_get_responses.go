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

// UserGetReader is a Reader for the UserGet structure.
type UserGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUserGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserGetOK creates a UserGetOK with default headers values
func NewUserGetOK() *UserGetOK {
	return &UserGetOK{}
}

/*UserGetOK handles this case with default header values.

Ok
*/
type UserGetOK struct {
	Payload *models.User
}

func (o *UserGetOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userGetOK  %+v", 200, o.Payload)
}

func (o *UserGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetBadRequest creates a UserGetBadRequest with default headers values
func NewUserGetBadRequest() *UserGetBadRequest {
	return &UserGetBadRequest{}
}

/*UserGetBadRequest handles this case with default header values.

Bad request
*/
type UserGetBadRequest struct {
}

func (o *UserGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userGetBadRequest ", 400)
}

func (o *UserGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetDefault creates a UserGetDefault with default headers values
func NewUserGetDefault(code int) *UserGetDefault {
	return &UserGetDefault{
		_statusCode: code,
	}
}

/*UserGetDefault handles this case with default header values.

Unexpected error
*/
type UserGetDefault struct {
	_statusCode int
}

// Code gets the status code for the user get default response
func (o *UserGetDefault) Code() int {
	return o._statusCode
}

func (o *UserGetDefault) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] UserGet default ", o._statusCode)
}

func (o *UserGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
