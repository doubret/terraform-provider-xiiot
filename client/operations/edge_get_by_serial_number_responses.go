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

// EdgeGetBySerialNumberReader is a Reader for the EdgeGetBySerialNumber structure.
type EdgeGetBySerialNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeGetBySerialNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEdgeGetBySerialNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEdgeGetBySerialNumberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEdgeGetBySerialNumberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEdgeGetBySerialNumberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeGetBySerialNumberOK creates a EdgeGetBySerialNumberOK with default headers values
func NewEdgeGetBySerialNumberOK() *EdgeGetBySerialNumberOK {
	return &EdgeGetBySerialNumberOK{}
}

/*EdgeGetBySerialNumberOK handles this case with default header values.

Ok
*/
type EdgeGetBySerialNumberOK struct {
	Payload *models.Edge
}

func (o *EdgeGetBySerialNumberOK) Error() string {
	return fmt.Sprintf("[POST /edgebyserialnumber][%d] edgeGetBySerialNumberOK  %+v", 200, o.Payload)
}

func (o *EdgeGetBySerialNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Edge)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeGetBySerialNumberBadRequest creates a EdgeGetBySerialNumberBadRequest with default headers values
func NewEdgeGetBySerialNumberBadRequest() *EdgeGetBySerialNumberBadRequest {
	return &EdgeGetBySerialNumberBadRequest{}
}

/*EdgeGetBySerialNumberBadRequest handles this case with default header values.

Bad request
*/
type EdgeGetBySerialNumberBadRequest struct {
}

func (o *EdgeGetBySerialNumberBadRequest) Error() string {
	return fmt.Sprintf("[POST /edgebyserialnumber][%d] edgeGetBySerialNumberBadRequest ", 400)
}

func (o *EdgeGetBySerialNumberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeGetBySerialNumberNotFound creates a EdgeGetBySerialNumberNotFound with default headers values
func NewEdgeGetBySerialNumberNotFound() *EdgeGetBySerialNumberNotFound {
	return &EdgeGetBySerialNumberNotFound{}
}

/*EdgeGetBySerialNumberNotFound handles this case with default header values.

Edge not found
*/
type EdgeGetBySerialNumberNotFound struct {
}

func (o *EdgeGetBySerialNumberNotFound) Error() string {
	return fmt.Sprintf("[POST /edgebyserialnumber][%d] edgeGetBySerialNumberNotFound ", 404)
}

func (o *EdgeGetBySerialNumberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeGetBySerialNumberDefault creates a EdgeGetBySerialNumberDefault with default headers values
func NewEdgeGetBySerialNumberDefault(code int) *EdgeGetBySerialNumberDefault {
	return &EdgeGetBySerialNumberDefault{
		_statusCode: code,
	}
}

/*EdgeGetBySerialNumberDefault handles this case with default header values.

Unexpected error
*/
type EdgeGetBySerialNumberDefault struct {
	_statusCode int
}

// Code gets the status code for the edge get by serial number default response
func (o *EdgeGetBySerialNumberDefault) Code() int {
	return o._statusCode
}

func (o *EdgeGetBySerialNumberDefault) Error() string {
	return fmt.Sprintf("[POST /edgebyserialnumber][%d] EdgeGetBySerialNumber default ", o._statusCode)
}

func (o *EdgeGetBySerialNumberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}