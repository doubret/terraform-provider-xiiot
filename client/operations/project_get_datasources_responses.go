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

// ProjectGetDatasourcesReader is a Reader for the ProjectGetDatasources structure.
type ProjectGetDatasourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetDatasourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectGetDatasourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectGetDatasourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewProjectGetDatasourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetDatasourcesOK creates a ProjectGetDatasourcesOK with default headers values
func NewProjectGetDatasourcesOK() *ProjectGetDatasourcesOK {
	return &ProjectGetDatasourcesOK{}
}

/*ProjectGetDatasourcesOK handles this case with default header values.

Ok
*/
type ProjectGetDatasourcesOK struct {
	Payload []*models.DataSource
}

func (o *ProjectGetDatasourcesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/datasources][%d] projectGetDatasourcesOK  %+v", 200, o.Payload)
}

func (o *ProjectGetDatasourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetDatasourcesBadRequest creates a ProjectGetDatasourcesBadRequest with default headers values
func NewProjectGetDatasourcesBadRequest() *ProjectGetDatasourcesBadRequest {
	return &ProjectGetDatasourcesBadRequest{}
}

/*ProjectGetDatasourcesBadRequest handles this case with default header values.

Bad request
*/
type ProjectGetDatasourcesBadRequest struct {
}

func (o *ProjectGetDatasourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/datasources][%d] projectGetDatasourcesBadRequest ", 400)
}

func (o *ProjectGetDatasourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectGetDatasourcesDefault creates a ProjectGetDatasourcesDefault with default headers values
func NewProjectGetDatasourcesDefault(code int) *ProjectGetDatasourcesDefault {
	return &ProjectGetDatasourcesDefault{
		_statusCode: code,
	}
}

/*ProjectGetDatasourcesDefault handles this case with default header values.

Unexpected error
*/
type ProjectGetDatasourcesDefault struct {
	_statusCode int
}

// Code gets the status code for the project get datasources default response
func (o *ProjectGetDatasourcesDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetDatasourcesDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/datasources][%d] ProjectGetDatasources default ", o._statusCode)
}

func (o *ProjectGetDatasourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
