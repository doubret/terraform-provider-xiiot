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

// UserListReader is a Reader for the UserList structure.
type UserListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUserListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserListOK creates a UserListOK with default headers values
func NewUserListOK() *UserListOK {
	return &UserListOK{}
}

/*UserListOK handles this case with default header values.

Ok
*/
type UserListOK struct {
	Payload []*models.User
}

func (o *UserListOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] userListOK  %+v", 200, o.Payload)
}

func (o *UserListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserListBadRequest creates a UserListBadRequest with default headers values
func NewUserListBadRequest() *UserListBadRequest {
	return &UserListBadRequest{}
}

/*UserListBadRequest handles this case with default header values.

Bad request
*/
type UserListBadRequest struct {
}

func (o *UserListBadRequest) Error() string {
	return fmt.Sprintf("[GET /users][%d] userListBadRequest ", 400)
}

func (o *UserListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserListDefault creates a UserListDefault with default headers values
func NewUserListDefault(code int) *UserListDefault {
	return &UserListDefault{
		_statusCode: code,
	}
}

/*UserListDefault handles this case with default header values.

Unexpected error
*/
type UserListDefault struct {
	_statusCode int
}

// Code gets the status code for the user list default response
func (o *UserListDefault) Code() int {
	return o._statusCode
}

func (o *UserListDefault) Error() string {
	return fmt.Sprintf("[GET /users][%d] UserList default ", o._statusCode)
}

func (o *UserListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
