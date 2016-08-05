package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// CreateAppReader is a Reader for the CreateApp structure.
type CreateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAppCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCreateAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewCreateAppCreated creates a CreateAppCreated with default headers values
func NewCreateAppCreated() *CreateAppCreated {
	return &CreateAppCreated{}
}

/*CreateAppCreated handles this case with default header values.

Newly created app
*/
type CreateAppCreated struct {
	Payload *models.App
}

func (o *CreateAppCreated) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/apps][%d] createAppCreated  %+v", 201, o.Payload)
}

func (o *CreateAppCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppUnauthorized creates a CreateAppUnauthorized with default headers values
func NewCreateAppUnauthorized() *CreateAppUnauthorized {
	return &CreateAppUnauthorized{}
}

/*CreateAppUnauthorized handles this case with default header values.

User not authorized
*/
type CreateAppUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *CreateAppUnauthorized) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/apps][%d] createAppUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppForbidden creates a CreateAppForbidden with default headers values
func NewCreateAppForbidden() *CreateAppForbidden {
	return &CreateAppForbidden{}
}

/*CreateAppForbidden handles this case with default header values.

User does not have the credentials to access this resource

*/
type CreateAppForbidden struct {
	Payload *models.Unauthorized
}

func (o *CreateAppForbidden) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/apps][%d] createAppForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppDefault creates a CreateAppDefault with default headers values
func NewCreateAppDefault(code int) *CreateAppDefault {
	return &CreateAppDefault{
		_statusCode: code,
	}
}

/*CreateAppDefault handles this case with default header values.

Error
*/
type CreateAppDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create app default response
func (o *CreateAppDefault) Code() int {
	return o._statusCode
}

func (o *CreateAppDefault) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/apps][%d] createApp default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
