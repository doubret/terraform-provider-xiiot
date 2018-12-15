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

// UserPropsUpdateReader is a Reader for the UserPropsUpdate structure.
type UserPropsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserPropsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserPropsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserPropsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUserPropsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserPropsUpdateOK creates a UserPropsUpdateOK with default headers values
func NewUserPropsUpdateOK() *UserPropsUpdateOK {
	return &UserPropsUpdateOK{}
}

/*UserPropsUpdateOK handles this case with default header values.

Ok
*/
type UserPropsUpdateOK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *UserPropsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /userprops/{id}][%d] userPropsUpdateOK  %+v", 200, o.Payload)
}

func (o *UserPropsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPropsUpdateBadRequest creates a UserPropsUpdateBadRequest with default headers values
func NewUserPropsUpdateBadRequest() *UserPropsUpdateBadRequest {
	return &UserPropsUpdateBadRequest{}
}

/*UserPropsUpdateBadRequest handles this case with default header values.

Bad request
*/
type UserPropsUpdateBadRequest struct {
}

func (o *UserPropsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /userprops/{id}][%d] userPropsUpdateBadRequest ", 400)
}

func (o *UserPropsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserPropsUpdateDefault creates a UserPropsUpdateDefault with default headers values
func NewUserPropsUpdateDefault(code int) *UserPropsUpdateDefault {
	return &UserPropsUpdateDefault{
		_statusCode: code,
	}
}

/*UserPropsUpdateDefault handles this case with default header values.

Unexpected error
*/
type UserPropsUpdateDefault struct {
	_statusCode int
}

// Code gets the status code for the user props update default response
func (o *UserPropsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UserPropsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /userprops/{id}][%d] UserPropsUpdate default ", o._statusCode)
}

func (o *UserPropsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}