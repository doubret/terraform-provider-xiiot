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

// CloudCredsGetReader is a Reader for the CloudCredsGet structure.
type CloudCredsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCloudCredsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCloudCredsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCloudCredsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudCredsGetOK creates a CloudCredsGetOK with default headers values
func NewCloudCredsGetOK() *CloudCredsGetOK {
	return &CloudCredsGetOK{}
}

/*CloudCredsGetOK handles this case with default header values.

Ok
*/
type CloudCredsGetOK struct {
	Payload *models.CloudCreds
}

func (o *CloudCredsGetOK) Error() string {
	return fmt.Sprintf("[GET /cloudcreds/{id}][%d] cloudCredsGetOK  %+v", 200, o.Payload)
}

func (o *CloudCredsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudCreds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredsGetBadRequest creates a CloudCredsGetBadRequest with default headers values
func NewCloudCredsGetBadRequest() *CloudCredsGetBadRequest {
	return &CloudCredsGetBadRequest{}
}

/*CloudCredsGetBadRequest handles this case with default header values.

Bad request
*/
type CloudCredsGetBadRequest struct {
}

func (o *CloudCredsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudcreds/{id}][%d] cloudCredsGetBadRequest ", 400)
}

func (o *CloudCredsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloudCredsGetDefault creates a CloudCredsGetDefault with default headers values
func NewCloudCredsGetDefault(code int) *CloudCredsGetDefault {
	return &CloudCredsGetDefault{
		_statusCode: code,
	}
}

/*CloudCredsGetDefault handles this case with default header values.

Unexpected error
*/
type CloudCredsGetDefault struct {
	_statusCode int
}

// Code gets the status code for the cloud creds get default response
func (o *CloudCredsGetDefault) Code() int {
	return o._statusCode
}

func (o *CloudCredsGetDefault) Error() string {
	return fmt.Sprintf("[GET /cloudcreds/{id}][%d] CloudCredsGet default ", o._statusCode)
}

func (o *CloudCredsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
