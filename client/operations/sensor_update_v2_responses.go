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

// SensorUpdateV2Reader is a Reader for the SensorUpdateV2 structure.
type SensorUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSensorUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSensorUpdateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSensorUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSensorUpdateV2OK creates a SensorUpdateV2OK with default headers values
func NewSensorUpdateV2OK() *SensorUpdateV2OK {
	return &SensorUpdateV2OK{}
}

/*SensorUpdateV2OK handles this case with default header values.

Ok
*/
type SensorUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *SensorUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /sensors/{id}][%d] sensorUpdateV2OK  %+v", 200, o.Payload)
}

func (o *SensorUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSensorUpdateV2BadRequest creates a SensorUpdateV2BadRequest with default headers values
func NewSensorUpdateV2BadRequest() *SensorUpdateV2BadRequest {
	return &SensorUpdateV2BadRequest{}
}

/*SensorUpdateV2BadRequest handles this case with default header values.

Bad request
*/
type SensorUpdateV2BadRequest struct {
}

func (o *SensorUpdateV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /sensors/{id}][%d] sensorUpdateV2BadRequest ", 400)
}

func (o *SensorUpdateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSensorUpdateV2Default creates a SensorUpdateV2Default with default headers values
func NewSensorUpdateV2Default(code int) *SensorUpdateV2Default {
	return &SensorUpdateV2Default{
		_statusCode: code,
	}
}

/*SensorUpdateV2Default handles this case with default header values.

Unexpected error
*/
type SensorUpdateV2Default struct {
	_statusCode int
}

// Code gets the status code for the sensor update v2 default response
func (o *SensorUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *SensorUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /sensors/{id}][%d] SensorUpdateV2 default ", o._statusCode)
}

func (o *SensorUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
