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

// ExecuteEdgeUpgradeReader is a Reader for the ExecuteEdgeUpgrade structure.
type ExecuteEdgeUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteEdgeUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExecuteEdgeUpgradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewExecuteEdgeUpgradeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewExecuteEdgeUpgradeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExecuteEdgeUpgradeOK creates a ExecuteEdgeUpgradeOK with default headers values
func NewExecuteEdgeUpgradeOK() *ExecuteEdgeUpgradeOK {
	return &ExecuteEdgeUpgradeOK{}
}

/*ExecuteEdgeUpgradeOK handles this case with default header values.

Ok
*/
type ExecuteEdgeUpgradeOK struct {
	Payload *models.CreateDocumentResponse
}

func (o *ExecuteEdgeUpgradeOK) Error() string {
	return fmt.Sprintf("[POST /edges/upgrade][%d] executeEdgeUpgradeOK  %+v", 200, o.Payload)
}

func (o *ExecuteEdgeUpgradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteEdgeUpgradeBadRequest creates a ExecuteEdgeUpgradeBadRequest with default headers values
func NewExecuteEdgeUpgradeBadRequest() *ExecuteEdgeUpgradeBadRequest {
	return &ExecuteEdgeUpgradeBadRequest{}
}

/*ExecuteEdgeUpgradeBadRequest handles this case with default header values.

Bad request
*/
type ExecuteEdgeUpgradeBadRequest struct {
}

func (o *ExecuteEdgeUpgradeBadRequest) Error() string {
	return fmt.Sprintf("[POST /edges/upgrade][%d] executeEdgeUpgradeBadRequest ", 400)
}

func (o *ExecuteEdgeUpgradeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecuteEdgeUpgradeDefault creates a ExecuteEdgeUpgradeDefault with default headers values
func NewExecuteEdgeUpgradeDefault(code int) *ExecuteEdgeUpgradeDefault {
	return &ExecuteEdgeUpgradeDefault{
		_statusCode: code,
	}
}

/*ExecuteEdgeUpgradeDefault handles this case with default header values.

Unexpected error
*/
type ExecuteEdgeUpgradeDefault struct {
	_statusCode int
}

// Code gets the status code for the execute edge upgrade default response
func (o *ExecuteEdgeUpgradeDefault) Code() int {
	return o._statusCode
}

func (o *ExecuteEdgeUpgradeDefault) Error() string {
	return fmt.Sprintf("[POST /edges/upgrade][%d] ExecuteEdgeUpgrade default ", o._statusCode)
}

func (o *ExecuteEdgeUpgradeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
