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

// DataStreamGetReader is a Reader for the DataStreamGet structure.
type DataStreamGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataStreamGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDataStreamGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDataStreamGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDataStreamGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataStreamGetOK creates a DataStreamGetOK with default headers values
func NewDataStreamGetOK() *DataStreamGetOK {
	return &DataStreamGetOK{}
}

/*DataStreamGetOK handles this case with default header values.

Ok
*/
type DataStreamGetOK struct {
	Payload *models.DataStream
}

func (o *DataStreamGetOK) Error() string {
	return fmt.Sprintf("[GET /datastreams/{id}][%d] dataStreamGetOK  %+v", 200, o.Payload)
}

func (o *DataStreamGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataStream)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataStreamGetBadRequest creates a DataStreamGetBadRequest with default headers values
func NewDataStreamGetBadRequest() *DataStreamGetBadRequest {
	return &DataStreamGetBadRequest{}
}

/*DataStreamGetBadRequest handles this case with default header values.

Bad request
*/
type DataStreamGetBadRequest struct {
}

func (o *DataStreamGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /datastreams/{id}][%d] dataStreamGetBadRequest ", 400)
}

func (o *DataStreamGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataStreamGetDefault creates a DataStreamGetDefault with default headers values
func NewDataStreamGetDefault(code int) *DataStreamGetDefault {
	return &DataStreamGetDefault{
		_statusCode: code,
	}
}

/*DataStreamGetDefault handles this case with default header values.

Unexpected error
*/
type DataStreamGetDefault struct {
	_statusCode int
}

// Code gets the status code for the data stream get default response
func (o *DataStreamGetDefault) Code() int {
	return o._statusCode
}

func (o *DataStreamGetDefault) Error() string {
	return fmt.Sprintf("[GET /datastreams/{id}][%d] DataStreamGet default ", o._statusCode)
}

func (o *DataStreamGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
