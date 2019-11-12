// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/runtime/virtcontainers/pkg/firecracker/client/models"
)

// PutLoggerReader is a Reader for the PutLogger structure.
type PutLoggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLoggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutLoggerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutLoggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutLoggerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLoggerNoContent creates a PutLoggerNoContent with default headers values
func NewPutLoggerNoContent() *PutLoggerNoContent {
	return &PutLoggerNoContent{}
}

/*PutLoggerNoContent handles this case with default header values.

Logger created.
*/
type PutLoggerNoContent struct {
}

func (o *PutLoggerNoContent) Error() string {
	return fmt.Sprintf("[PUT /logger][%d] putLoggerNoContent ", 204)
}

func (o *PutLoggerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLoggerBadRequest creates a PutLoggerBadRequest with default headers values
func NewPutLoggerBadRequest() *PutLoggerBadRequest {
	return &PutLoggerBadRequest{}
}

/*PutLoggerBadRequest handles this case with default header values.

Logger cannot be initialized due to bad input.
*/
type PutLoggerBadRequest struct {
	Payload *models.Error
}

func (o *PutLoggerBadRequest) Error() string {
	return fmt.Sprintf("[PUT /logger][%d] putLoggerBadRequest  %+v", 400, o.Payload)
}

func (o *PutLoggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLoggerDefault creates a PutLoggerDefault with default headers values
func NewPutLoggerDefault(code int) *PutLoggerDefault {
	return &PutLoggerDefault{
		_statusCode: code,
	}
}

/*PutLoggerDefault handles this case with default header values.

Internal server error.
*/
type PutLoggerDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put logger default response
func (o *PutLoggerDefault) Code() int {
	return o._statusCode
}

func (o *PutLoggerDefault) Error() string {
	return fmt.Sprintf("[PUT /logger][%d] putLogger default  %+v", o._statusCode, o.Payload)
}

func (o *PutLoggerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
