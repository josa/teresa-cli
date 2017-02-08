package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa-api/models"
)

// GetTeamDetailReader is a Reader for the GetTeamDetail structure.
type GetTeamDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetTeamDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTeamDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetTeamDetailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTeamDetailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetTeamDetailDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetTeamDetailOK creates a GetTeamDetailOK with default headers values
func NewGetTeamDetailOK() *GetTeamDetailOK {
	return &GetTeamDetailOK{}
}

/*GetTeamDetailOK handles this case with default header values.

Team details
*/
type GetTeamDetailOK struct {
	Payload *models.Team
}

func (o *GetTeamDetailOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamDetailOK  %+v", 200, o.Payload)
}

func (o *GetTeamDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamDetailUnauthorized creates a GetTeamDetailUnauthorized with default headers values
func NewGetTeamDetailUnauthorized() *GetTeamDetailUnauthorized {
	return &GetTeamDetailUnauthorized{}
}

/*GetTeamDetailUnauthorized handles this case with default header values.

User not authorized
*/
type GetTeamDetailUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetTeamDetailUnauthorized) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamDetailUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamDetailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamDetailForbidden creates a GetTeamDetailForbidden with default headers values
func NewGetTeamDetailForbidden() *GetTeamDetailForbidden {
	return &GetTeamDetailForbidden{}
}

/*GetTeamDetailForbidden handles this case with default header values.

User does not have the credentials to access this resource

*/
type GetTeamDetailForbidden struct {
	Payload *models.Unauthorized
}

func (o *GetTeamDetailForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamDetailForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamDetailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamDetailDefault creates a GetTeamDetailDefault with default headers values
func NewGetTeamDetailDefault(code int) *GetTeamDetailDefault {
	return &GetTeamDetailDefault{
		_statusCode: code,
	}
}

/*GetTeamDetailDefault handles this case with default header values.

Error
*/
type GetTeamDetailDefault struct {
	_statusCode int

	Payload *models.GenericError
}

// Code gets the status code for the get team detail default response
func (o *GetTeamDetailDefault) Code() int {
	return o._statusCode
}

func (o *GetTeamDetailDefault) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamDetail default  %+v", o._statusCode, o.Payload)
}

func (o *GetTeamDetailDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
