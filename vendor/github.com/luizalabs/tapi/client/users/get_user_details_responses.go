package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// GetUserDetailsReader is a Reader for the GetUserDetails structure.
type GetUserDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetUserDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetUserDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetUserDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetUserDetailsOK creates a GetUserDetailsOK with default headers values
func NewGetUserDetailsOK() *GetUserDetailsOK {
	return &GetUserDetailsOK{}
}

/*GetUserDetailsOK handles this case with default header values.

User details
*/
type GetUserDetailsOK struct {
	Payload *models.User
}

func (o *GetUserDetailsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserDetailsOK  %+v", 200, o.Payload)
}

func (o *GetUserDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDetailsUnauthorized creates a GetUserDetailsUnauthorized with default headers values
func NewGetUserDetailsUnauthorized() *GetUserDetailsUnauthorized {
	return &GetUserDetailsUnauthorized{}
}

/*GetUserDetailsUnauthorized handles this case with default header values.

User not authorized
*/
type GetUserDetailsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetUserDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDetailsForbidden creates a GetUserDetailsForbidden with default headers values
func NewGetUserDetailsForbidden() *GetUserDetailsForbidden {
	return &GetUserDetailsForbidden{}
}

/*GetUserDetailsForbidden handles this case with default header values.

User does not have the credentials to access this resource

*/
type GetUserDetailsForbidden struct {
	Payload *models.Unauthorized
}

func (o *GetUserDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDetailsNotFound creates a GetUserDetailsNotFound with default headers values
func NewGetUserDetailsNotFound() *GetUserDetailsNotFound {
	return &GetUserDetailsNotFound{}
}

/*GetUserDetailsNotFound handles this case with default header values.

Resource not found
*/
type GetUserDetailsNotFound struct {
	Payload *models.NotFound
}

func (o *GetUserDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDetailsDefault creates a GetUserDetailsDefault with default headers values
func NewGetUserDetailsDefault(code int) *GetUserDetailsDefault {
	return &GetUserDetailsDefault{
		_statusCode: code,
	}
}

/*GetUserDetailsDefault handles this case with default header values.

Error
*/
type GetUserDetailsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get user details default response
func (o *GetUserDetailsDefault) Code() int {
	return o._statusCode
}

func (o *GetUserDetailsDefault) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
