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

// DataSourceListReader is a Reader for the DataSourceList structure.
type DataSourceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataSourceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDataSourceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDataSourceListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDataSourceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataSourceListOK creates a DataSourceListOK with default headers values
func NewDataSourceListOK() *DataSourceListOK {
	return &DataSourceListOK{}
}

/*DataSourceListOK handles this case with default header values.

Ok
*/
type DataSourceListOK struct {
	Payload []*models.DataSource
}

func (o *DataSourceListOK) Error() string {
	return fmt.Sprintf("[GET /datasources][%d] dataSourceListOK  %+v", 200, o.Payload)
}

func (o *DataSourceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataSourceListBadRequest creates a DataSourceListBadRequest with default headers values
func NewDataSourceListBadRequest() *DataSourceListBadRequest {
	return &DataSourceListBadRequest{}
}

/*DataSourceListBadRequest handles this case with default header values.

Bad request
*/
type DataSourceListBadRequest struct {
}

func (o *DataSourceListBadRequest) Error() string {
	return fmt.Sprintf("[GET /datasources][%d] dataSourceListBadRequest ", 400)
}

func (o *DataSourceListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataSourceListDefault creates a DataSourceListDefault with default headers values
func NewDataSourceListDefault(code int) *DataSourceListDefault {
	return &DataSourceListDefault{
		_statusCode: code,
	}
}

/*DataSourceListDefault handles this case with default header values.

Unexpected error
*/
type DataSourceListDefault struct {
	_statusCode int
}

// Code gets the status code for the data source list default response
func (o *DataSourceListDefault) Code() int {
	return o._statusCode
}

func (o *DataSourceListDefault) Error() string {
	return fmt.Sprintf("[GET /datasources][%d] DataSourceList default ", o._statusCode)
}

func (o *DataSourceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
