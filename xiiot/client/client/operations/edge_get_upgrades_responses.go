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

// EdgeGetUpgradesReader is a Reader for the EdgeGetUpgrades structure.
type EdgeGetUpgradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeGetUpgradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEdgeGetUpgradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEdgeGetUpgradesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEdgeGetUpgradesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeGetUpgradesOK creates a EdgeGetUpgradesOK with default headers values
func NewEdgeGetUpgradesOK() *EdgeGetUpgradesOK {
	return &EdgeGetUpgradesOK{}
}

/*EdgeGetUpgradesOK handles this case with default header values.

Ok
*/
type EdgeGetUpgradesOK struct {
	Payload []*models.EdgeUpgradeCore
}

func (o *EdgeGetUpgradesOK) Error() string {
	return fmt.Sprintf("[GET /edges/{edgeId}/upgradecompatible][%d] edgeGetUpgradesOK  %+v", 200, o.Payload)
}

func (o *EdgeGetUpgradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeGetUpgradesBadRequest creates a EdgeGetUpgradesBadRequest with default headers values
func NewEdgeGetUpgradesBadRequest() *EdgeGetUpgradesBadRequest {
	return &EdgeGetUpgradesBadRequest{}
}

/*EdgeGetUpgradesBadRequest handles this case with default header values.

Bad request
*/
type EdgeGetUpgradesBadRequest struct {
}

func (o *EdgeGetUpgradesBadRequest) Error() string {
	return fmt.Sprintf("[GET /edges/{edgeId}/upgradecompatible][%d] edgeGetUpgradesBadRequest ", 400)
}

func (o *EdgeGetUpgradesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeGetUpgradesDefault creates a EdgeGetUpgradesDefault with default headers values
func NewEdgeGetUpgradesDefault(code int) *EdgeGetUpgradesDefault {
	return &EdgeGetUpgradesDefault{
		_statusCode: code,
	}
}

/*EdgeGetUpgradesDefault handles this case with default header values.

Unexpected error
*/
type EdgeGetUpgradesDefault struct {
	_statusCode int
}

// Code gets the status code for the edge get upgrades default response
func (o *EdgeGetUpgradesDefault) Code() int {
	return o._statusCode
}

func (o *EdgeGetUpgradesDefault) Error() string {
	return fmt.Sprintf("[GET /edges/{edgeId}/upgradecompatible][%d] EdgeGetUpgrades default ", o._statusCode)
}

func (o *EdgeGetUpgradesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
