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

// EdgeCreateReader is a Reader for the EdgeCreate structure.
type EdgeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEdgeCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEdgeCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEdgeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeCreateOK creates a EdgeCreateOK with default headers values
func NewEdgeCreateOK() *EdgeCreateOK {
	return &EdgeCreateOK{}
}

/*EdgeCreateOK handles this case with default header values.

Ok
*/
type EdgeCreateOK struct {
	Payload *models.CreateDocumentResponse
}

func (o *EdgeCreateOK) Error() string {
	return fmt.Sprintf("[POST /edges][%d] edgeCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeCreateBadRequest creates a EdgeCreateBadRequest with default headers values
func NewEdgeCreateBadRequest() *EdgeCreateBadRequest {
	return &EdgeCreateBadRequest{}
}

/*EdgeCreateBadRequest handles this case with default header values.

Bad request
*/
type EdgeCreateBadRequest struct {
}

func (o *EdgeCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /edges][%d] edgeCreateBadRequest ", 400)
}

func (o *EdgeCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeCreateDefault creates a EdgeCreateDefault with default headers values
func NewEdgeCreateDefault(code int) *EdgeCreateDefault {
	return &EdgeCreateDefault{
		_statusCode: code,
	}
}

/*EdgeCreateDefault handles this case with default header values.

Unexpected error
*/
type EdgeCreateDefault struct {
	_statusCode int
}

// Code gets the status code for the edge create default response
func (o *EdgeCreateDefault) Code() int {
	return o._statusCode
}

func (o *EdgeCreateDefault) Error() string {
	return fmt.Sprintf("[POST /edges][%d] EdgeCreate default ", o._statusCode)
}

func (o *EdgeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}