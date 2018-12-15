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

// DataStreamDeleteReader is a Reader for the DataStreamDelete structure.
type DataStreamDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataStreamDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDataStreamDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDataStreamDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDataStreamDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataStreamDeleteOK creates a DataStreamDeleteOK with default headers values
func NewDataStreamDeleteOK() *DataStreamDeleteOK {
	return &DataStreamDeleteOK{}
}

/*DataStreamDeleteOK handles this case with default header values.

Ok
*/
type DataStreamDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *DataStreamDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /datastreams/{id}][%d] dataStreamDeleteOK  %+v", 200, o.Payload)
}

func (o *DataStreamDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataStreamDeleteBadRequest creates a DataStreamDeleteBadRequest with default headers values
func NewDataStreamDeleteBadRequest() *DataStreamDeleteBadRequest {
	return &DataStreamDeleteBadRequest{}
}

/*DataStreamDeleteBadRequest handles this case with default header values.

Bad request
*/
type DataStreamDeleteBadRequest struct {
}

func (o *DataStreamDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /datastreams/{id}][%d] dataStreamDeleteBadRequest ", 400)
}

func (o *DataStreamDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataStreamDeleteDefault creates a DataStreamDeleteDefault with default headers values
func NewDataStreamDeleteDefault(code int) *DataStreamDeleteDefault {
	return &DataStreamDeleteDefault{
		_statusCode: code,
	}
}

/*DataStreamDeleteDefault handles this case with default header values.

Unexpected error
*/
type DataStreamDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the data stream delete default response
func (o *DataStreamDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DataStreamDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /datastreams/{id}][%d] DataStreamDelete default ", o._statusCode)
}

func (o *DataStreamDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}