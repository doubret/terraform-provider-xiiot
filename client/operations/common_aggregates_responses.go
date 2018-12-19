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

// CommonAggregatesReader is a Reader for the CommonAggregates structure.
type CommonAggregatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommonAggregatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCommonAggregatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCommonAggregatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCommonAggregatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCommonAggregatesOK creates a CommonAggregatesOK with default headers values
func NewCommonAggregatesOK() *CommonAggregatesOK {
	return &CommonAggregatesOK{}
}

/*CommonAggregatesOK handles this case with default header values.

Ok
*/
type CommonAggregatesOK struct {
	Payload []*models.AggregateInfo
}

func (o *CommonAggregatesOK) Error() string {
	return fmt.Sprintf("[POST /common/aggregates][%d] commonAggregatesOK  %+v", 200, o.Payload)
}

func (o *CommonAggregatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonAggregatesBadRequest creates a CommonAggregatesBadRequest with default headers values
func NewCommonAggregatesBadRequest() *CommonAggregatesBadRequest {
	return &CommonAggregatesBadRequest{}
}

/*CommonAggregatesBadRequest handles this case with default header values.

Bad request
*/
type CommonAggregatesBadRequest struct {
}

func (o *CommonAggregatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /common/aggregates][%d] commonAggregatesBadRequest ", 400)
}

func (o *CommonAggregatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommonAggregatesDefault creates a CommonAggregatesDefault with default headers values
func NewCommonAggregatesDefault(code int) *CommonAggregatesDefault {
	return &CommonAggregatesDefault{
		_statusCode: code,
	}
}

/*CommonAggregatesDefault handles this case with default header values.

Unexpected error
*/
type CommonAggregatesDefault struct {
	_statusCode int
}

// Code gets the status code for the common aggregates default response
func (o *CommonAggregatesDefault) Code() int {
	return o._statusCode
}

func (o *CommonAggregatesDefault) Error() string {
	return fmt.Sprintf("[POST /common/aggregates][%d] CommonAggregates default ", o._statusCode)
}

func (o *CommonAggregatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
