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

// ScriptRuntimeListReader is a Reader for the ScriptRuntimeList structure.
type ScriptRuntimeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScriptRuntimeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewScriptRuntimeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewScriptRuntimeListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewScriptRuntimeListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScriptRuntimeListOK creates a ScriptRuntimeListOK with default headers values
func NewScriptRuntimeListOK() *ScriptRuntimeListOK {
	return &ScriptRuntimeListOK{}
}

/*ScriptRuntimeListOK handles this case with default header values.

Ok
*/
type ScriptRuntimeListOK struct {
	Payload []*models.ScriptRuntime
}

func (o *ScriptRuntimeListOK) Error() string {
	return fmt.Sprintf("[GET /scriptruntimes][%d] scriptRuntimeListOK  %+v", 200, o.Payload)
}

func (o *ScriptRuntimeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScriptRuntimeListBadRequest creates a ScriptRuntimeListBadRequest with default headers values
func NewScriptRuntimeListBadRequest() *ScriptRuntimeListBadRequest {
	return &ScriptRuntimeListBadRequest{}
}

/*ScriptRuntimeListBadRequest handles this case with default header values.

Bad request
*/
type ScriptRuntimeListBadRequest struct {
}

func (o *ScriptRuntimeListBadRequest) Error() string {
	return fmt.Sprintf("[GET /scriptruntimes][%d] scriptRuntimeListBadRequest ", 400)
}

func (o *ScriptRuntimeListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScriptRuntimeListDefault creates a ScriptRuntimeListDefault with default headers values
func NewScriptRuntimeListDefault(code int) *ScriptRuntimeListDefault {
	return &ScriptRuntimeListDefault{
		_statusCode: code,
	}
}

/*ScriptRuntimeListDefault handles this case with default header values.

Unexpected error
*/
type ScriptRuntimeListDefault struct {
	_statusCode int
}

// Code gets the status code for the script runtime list default response
func (o *ScriptRuntimeListDefault) Code() int {
	return o._statusCode
}

func (o *ScriptRuntimeListDefault) Error() string {
	return fmt.Sprintf("[GET /scriptruntimes][%d] ScriptRuntimeList default ", o._statusCode)
}

func (o *ScriptRuntimeListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
