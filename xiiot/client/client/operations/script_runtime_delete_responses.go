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

// ScriptRuntimeDeleteReader is a Reader for the ScriptRuntimeDelete structure.
type ScriptRuntimeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScriptRuntimeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewScriptRuntimeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewScriptRuntimeDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewScriptRuntimeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScriptRuntimeDeleteOK creates a ScriptRuntimeDeleteOK with default headers values
func NewScriptRuntimeDeleteOK() *ScriptRuntimeDeleteOK {
	return &ScriptRuntimeDeleteOK{}
}

/*ScriptRuntimeDeleteOK handles this case with default header values.

Ok
*/
type ScriptRuntimeDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *ScriptRuntimeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /scriptruntimes/{id}][%d] scriptRuntimeDeleteOK  %+v", 200, o.Payload)
}

func (o *ScriptRuntimeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScriptRuntimeDeleteBadRequest creates a ScriptRuntimeDeleteBadRequest with default headers values
func NewScriptRuntimeDeleteBadRequest() *ScriptRuntimeDeleteBadRequest {
	return &ScriptRuntimeDeleteBadRequest{}
}

/*ScriptRuntimeDeleteBadRequest handles this case with default header values.

Bad request
*/
type ScriptRuntimeDeleteBadRequest struct {
}

func (o *ScriptRuntimeDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /scriptruntimes/{id}][%d] scriptRuntimeDeleteBadRequest ", 400)
}

func (o *ScriptRuntimeDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScriptRuntimeDeleteDefault creates a ScriptRuntimeDeleteDefault with default headers values
func NewScriptRuntimeDeleteDefault(code int) *ScriptRuntimeDeleteDefault {
	return &ScriptRuntimeDeleteDefault{
		_statusCode: code,
	}
}

/*ScriptRuntimeDeleteDefault handles this case with default header values.

Unexpected error
*/
type ScriptRuntimeDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the script runtime delete default response
func (o *ScriptRuntimeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ScriptRuntimeDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /scriptruntimes/{id}][%d] ScriptRuntimeDelete default ", o._statusCode)
}

func (o *ScriptRuntimeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
