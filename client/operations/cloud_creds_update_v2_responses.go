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

// CloudCredsUpdateV2Reader is a Reader for the CloudCredsUpdateV2 structure.
type CloudCredsUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredsUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCloudCredsUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCloudCredsUpdateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCloudCredsUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudCredsUpdateV2OK creates a CloudCredsUpdateV2OK with default headers values
func NewCloudCredsUpdateV2OK() *CloudCredsUpdateV2OK {
	return &CloudCredsUpdateV2OK{}
}

/*CloudCredsUpdateV2OK handles this case with default header values.

Ok
*/
type CloudCredsUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *CloudCredsUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /cloudcreds/{id}][%d] cloudCredsUpdateV2OK  %+v", 200, o.Payload)
}

func (o *CloudCredsUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredsUpdateV2BadRequest creates a CloudCredsUpdateV2BadRequest with default headers values
func NewCloudCredsUpdateV2BadRequest() *CloudCredsUpdateV2BadRequest {
	return &CloudCredsUpdateV2BadRequest{}
}

/*CloudCredsUpdateV2BadRequest handles this case with default header values.

Bad request
*/
type CloudCredsUpdateV2BadRequest struct {
}

func (o *CloudCredsUpdateV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /cloudcreds/{id}][%d] cloudCredsUpdateV2BadRequest ", 400)
}

func (o *CloudCredsUpdateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloudCredsUpdateV2Default creates a CloudCredsUpdateV2Default with default headers values
func NewCloudCredsUpdateV2Default(code int) *CloudCredsUpdateV2Default {
	return &CloudCredsUpdateV2Default{
		_statusCode: code,
	}
}

/*CloudCredsUpdateV2Default handles this case with default header values.

Unexpected error
*/
type CloudCredsUpdateV2Default struct {
	_statusCode int
}

// Code gets the status code for the cloud creds update v2 default response
func (o *CloudCredsUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *CloudCredsUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /cloudcreds/{id}][%d] CloudCredsUpdateV2 default ", o._statusCode)
}

func (o *CloudCredsUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
