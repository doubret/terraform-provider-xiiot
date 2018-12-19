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

// ScriptUpdateV2Reader is a Reader for the ScriptUpdateV2 structure.
type ScriptUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScriptUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewScriptUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewScriptUpdateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewScriptUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScriptUpdateV2OK creates a ScriptUpdateV2OK with default headers values
func NewScriptUpdateV2OK() *ScriptUpdateV2OK {
	return &ScriptUpdateV2OK{}
}

/*ScriptUpdateV2OK handles this case with default header values.

Ok
*/
type ScriptUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *ScriptUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /scripts/{id}][%d] scriptUpdateV2OK  %+v", 200, o.Payload)
}

func (o *ScriptUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScriptUpdateV2BadRequest creates a ScriptUpdateV2BadRequest with default headers values
func NewScriptUpdateV2BadRequest() *ScriptUpdateV2BadRequest {
	return &ScriptUpdateV2BadRequest{}
}

/*ScriptUpdateV2BadRequest handles this case with default header values.

Bad request
*/
type ScriptUpdateV2BadRequest struct {
}

func (o *ScriptUpdateV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /scripts/{id}][%d] scriptUpdateV2BadRequest ", 400)
}

func (o *ScriptUpdateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScriptUpdateV2Default creates a ScriptUpdateV2Default with default headers values
func NewScriptUpdateV2Default(code int) *ScriptUpdateV2Default {
	return &ScriptUpdateV2Default{
		_statusCode: code,
	}
}

/*ScriptUpdateV2Default handles this case with default header values.

Unexpected error
*/
type ScriptUpdateV2Default struct {
	_statusCode int
}

// Code gets the status code for the script update v2 default response
func (o *ScriptUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *ScriptUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /scripts/{id}][%d] ScriptUpdateV2 default ", o._statusCode)
}

func (o *ScriptUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}