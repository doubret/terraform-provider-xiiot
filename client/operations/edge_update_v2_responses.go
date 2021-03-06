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

// EdgeUpdateV2Reader is a Reader for the EdgeUpdateV2 structure.
type EdgeUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEdgeUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEdgeUpdateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEdgeUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeUpdateV2OK creates a EdgeUpdateV2OK with default headers values
func NewEdgeUpdateV2OK() *EdgeUpdateV2OK {
	return &EdgeUpdateV2OK{}
}

/*EdgeUpdateV2OK handles this case with default header values.

Ok
*/
type EdgeUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *EdgeUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /edges/{id}][%d] edgeUpdateV2OK  %+v", 200, o.Payload)
}

func (o *EdgeUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeUpdateV2BadRequest creates a EdgeUpdateV2BadRequest with default headers values
func NewEdgeUpdateV2BadRequest() *EdgeUpdateV2BadRequest {
	return &EdgeUpdateV2BadRequest{}
}

/*EdgeUpdateV2BadRequest handles this case with default header values.

Bad request
*/
type EdgeUpdateV2BadRequest struct {
}

func (o *EdgeUpdateV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /edges/{id}][%d] edgeUpdateV2BadRequest ", 400)
}

func (o *EdgeUpdateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeUpdateV2Default creates a EdgeUpdateV2Default with default headers values
func NewEdgeUpdateV2Default(code int) *EdgeUpdateV2Default {
	return &EdgeUpdateV2Default{
		_statusCode: code,
	}
}

/*EdgeUpdateV2Default handles this case with default header values.

Unexpected error
*/
type EdgeUpdateV2Default struct {
	_statusCode int
}

// Code gets the status code for the edge update v2 default response
func (o *EdgeUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *EdgeUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /edges/{id}][%d] EdgeUpdateV2 default ", o._statusCode)
}

func (o *EdgeUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
