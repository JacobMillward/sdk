// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra-client-go/models"
)

// PatchOAuth2ClientReader is a Reader for the PatchOAuth2Client structure.
type PatchOAuth2ClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOAuth2ClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOAuth2ClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPatchOAuth2ClientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchOAuth2ClientOK creates a PatchOAuth2ClientOK with default headers values
func NewPatchOAuth2ClientOK() *PatchOAuth2ClientOK {
	return &PatchOAuth2ClientOK{}
}

/*PatchOAuth2ClientOK handles this case with default header values.

oAuth2Client
*/
type PatchOAuth2ClientOK struct {
	Payload *models.OAuth2Client
}

func (o *PatchOAuth2ClientOK) Error() string {
	return fmt.Sprintf("[PATCH /clients/{id}][%d] patchOAuth2ClientOK  %+v", 200, o.Payload)
}

func (o *PatchOAuth2ClientOK) GetPayload() *models.OAuth2Client {
	return o.Payload
}

func (o *PatchOAuth2ClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOAuth2ClientInternalServerError creates a PatchOAuth2ClientInternalServerError with default headers values
func NewPatchOAuth2ClientInternalServerError() *PatchOAuth2ClientInternalServerError {
	return &PatchOAuth2ClientInternalServerError{}
}

/*PatchOAuth2ClientInternalServerError handles this case with default header values.

genericError
*/
type PatchOAuth2ClientInternalServerError struct {
	Payload *models.GenericError
}

func (o *PatchOAuth2ClientInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /clients/{id}][%d] patchOAuth2ClientInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchOAuth2ClientInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *PatchOAuth2ClientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}