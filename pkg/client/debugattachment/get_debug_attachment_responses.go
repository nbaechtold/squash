// Code generated by go-swagger; DO NOT EDIT.

package debugattachment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// GetDebugAttachmentReader is a Reader for the GetDebugAttachment structure.
type GetDebugAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebugAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDebugAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDebugAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDebugAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetDebugAttachmentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDebugAttachmentOK creates a GetDebugAttachmentOK with default headers values
func NewGetDebugAttachmentOK() *GetDebugAttachmentOK {
	return &GetDebugAttachmentOK{}
}

/*GetDebugAttachmentOK handles this case with default header values.

OK
*/
type GetDebugAttachmentOK struct {
	Payload *models.DebugAttachment
}

func (o *GetDebugAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /debugattachment/{debugAttachmentId}][%d] getDebugAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetDebugAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDebugAttachmentBadRequest creates a GetDebugAttachmentBadRequest with default headers values
func NewGetDebugAttachmentBadRequest() *GetDebugAttachmentBadRequest {
	return &GetDebugAttachmentBadRequest{}
}

/*GetDebugAttachmentBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetDebugAttachmentBadRequest struct {
}

func (o *GetDebugAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /debugattachment/{debugAttachmentId}][%d] getDebugAttachmentBadRequest ", 400)
}

func (o *GetDebugAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDebugAttachmentNotFound creates a GetDebugAttachmentNotFound with default headers values
func NewGetDebugAttachmentNotFound() *GetDebugAttachmentNotFound {
	return &GetDebugAttachmentNotFound{}
}

/*GetDebugAttachmentNotFound handles this case with default header values.

Debug config not found
*/
type GetDebugAttachmentNotFound struct {
}

func (o *GetDebugAttachmentNotFound) Error() string {
	return fmt.Sprintf("[GET /debugattachment/{debugAttachmentId}][%d] getDebugAttachmentNotFound ", 404)
}

func (o *GetDebugAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDebugAttachmentUnprocessableEntity creates a GetDebugAttachmentUnprocessableEntity with default headers values
func NewGetDebugAttachmentUnprocessableEntity() *GetDebugAttachmentUnprocessableEntity {
	return &GetDebugAttachmentUnprocessableEntity{}
}

/*GetDebugAttachmentUnprocessableEntity handles this case with default header values.

Validation exception
*/
type GetDebugAttachmentUnprocessableEntity struct {
}

func (o *GetDebugAttachmentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /debugattachment/{debugAttachmentId}][%d] getDebugAttachmentUnprocessableEntity ", 422)
}

func (o *GetDebugAttachmentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
