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

// DataStreamListReader is a Reader for the DataStreamList structure.
type DataStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDataStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDataStreamListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDataStreamListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataStreamListOK creates a DataStreamListOK with default headers values
func NewDataStreamListOK() *DataStreamListOK {
	return &DataStreamListOK{}
}

/*DataStreamListOK handles this case with default header values.

Ok
*/
type DataStreamListOK struct {
	Payload []*models.DataStream
}

func (o *DataStreamListOK) Error() string {
	return fmt.Sprintf("[GET /datastreams][%d] dataStreamListOK  %+v", 200, o.Payload)
}

func (o *DataStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataStreamListBadRequest creates a DataStreamListBadRequest with default headers values
func NewDataStreamListBadRequest() *DataStreamListBadRequest {
	return &DataStreamListBadRequest{}
}

/*DataStreamListBadRequest handles this case with default header values.

Bad request
*/
type DataStreamListBadRequest struct {
}

func (o *DataStreamListBadRequest) Error() string {
	return fmt.Sprintf("[GET /datastreams][%d] dataStreamListBadRequest ", 400)
}

func (o *DataStreamListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataStreamListDefault creates a DataStreamListDefault with default headers values
func NewDataStreamListDefault(code int) *DataStreamListDefault {
	return &DataStreamListDefault{
		_statusCode: code,
	}
}

/*DataStreamListDefault handles this case with default header values.

Unexpected error
*/
type DataStreamListDefault struct {
	_statusCode int
}

// Code gets the status code for the data stream list default response
func (o *DataStreamListDefault) Code() int {
	return o._statusCode
}

func (o *DataStreamListDefault) Error() string {
	return fmt.Sprintf("[GET /datastreams][%d] DataStreamList default ", o._statusCode)
}

func (o *DataStreamListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
