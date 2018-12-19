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

// DataSourceUpdateV2Reader is a Reader for the DataSourceUpdateV2 structure.
type DataSourceUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataSourceUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDataSourceUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDataSourceUpdateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDataSourceUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataSourceUpdateV2OK creates a DataSourceUpdateV2OK with default headers values
func NewDataSourceUpdateV2OK() *DataSourceUpdateV2OK {
	return &DataSourceUpdateV2OK{}
}

/*DataSourceUpdateV2OK handles this case with default header values.

Ok
*/
type DataSourceUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *DataSourceUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /datasources/{id}][%d] dataSourceUpdateV2OK  %+v", 200, o.Payload)
}

func (o *DataSourceUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataSourceUpdateV2BadRequest creates a DataSourceUpdateV2BadRequest with default headers values
func NewDataSourceUpdateV2BadRequest() *DataSourceUpdateV2BadRequest {
	return &DataSourceUpdateV2BadRequest{}
}

/*DataSourceUpdateV2BadRequest handles this case with default header values.

Bad request
*/
type DataSourceUpdateV2BadRequest struct {
}

func (o *DataSourceUpdateV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /datasources/{id}][%d] dataSourceUpdateV2BadRequest ", 400)
}

func (o *DataSourceUpdateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataSourceUpdateV2Default creates a DataSourceUpdateV2Default with default headers values
func NewDataSourceUpdateV2Default(code int) *DataSourceUpdateV2Default {
	return &DataSourceUpdateV2Default{
		_statusCode: code,
	}
}

/*DataSourceUpdateV2Default handles this case with default header values.

Unexpected error
*/
type DataSourceUpdateV2Default struct {
	_statusCode int
}

// Code gets the status code for the data source update v2 default response
func (o *DataSourceUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *DataSourceUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /datasources/{id}][%d] DataSourceUpdateV2 default ", o._statusCode)
}

func (o *DataSourceUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
