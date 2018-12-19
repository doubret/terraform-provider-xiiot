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

// EdgeDeleteReader is a Reader for the EdgeDelete structure.
type EdgeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEdgeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEdgeDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEdgeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeDeleteOK creates a EdgeDeleteOK with default headers values
func NewEdgeDeleteOK() *EdgeDeleteOK {
	return &EdgeDeleteOK{}
}

/*EdgeDeleteOK handles this case with default header values.

Ok
*/
type EdgeDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *EdgeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /edges/{edgeId}][%d] edgeDeleteOK  %+v", 200, o.Payload)
}

func (o *EdgeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDeleteBadRequest creates a EdgeDeleteBadRequest with default headers values
func NewEdgeDeleteBadRequest() *EdgeDeleteBadRequest {
	return &EdgeDeleteBadRequest{}
}

/*EdgeDeleteBadRequest handles this case with default header values.

Bad request
*/
type EdgeDeleteBadRequest struct {
}

func (o *EdgeDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /edges/{edgeId}][%d] edgeDeleteBadRequest ", 400)
}

func (o *EdgeDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeDeleteDefault creates a EdgeDeleteDefault with default headers values
func NewEdgeDeleteDefault(code int) *EdgeDeleteDefault {
	return &EdgeDeleteDefault{
		_statusCode: code,
	}
}

/*EdgeDeleteDefault handles this case with default header values.

Unexpected error
*/
type EdgeDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the edge delete default response
func (o *EdgeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *EdgeDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /edges/{edgeId}][%d] EdgeDelete default ", o._statusCode)
}

func (o *EdgeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
