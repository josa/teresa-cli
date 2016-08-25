package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// NewPartialUpdateAppParams creates a new PartialUpdateAppParams object
// with the default values initialized.
func NewPartialUpdateAppParams() *PartialUpdateAppParams {
	var ()
	return &PartialUpdateAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPartialUpdateAppParamsWithTimeout creates a new PartialUpdateAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPartialUpdateAppParamsWithTimeout(timeout time.Duration) *PartialUpdateAppParams {
	var ()
	return &PartialUpdateAppParams{

		timeout: timeout,
	}
}

/*PartialUpdateAppParams contains all the parameters to send to the API endpoint
for the partial update app operation typically these are written to a http.Request
*/
type PartialUpdateAppParams struct {

	/*AppID
	  App ID

	*/
	AppID int64
	/*Body*/
	Body []*models.PatchAppRequest
	/*TeamID
	  Team ID

	*/
	TeamID int64

	timeout time.Duration
}

// WithAppID adds the appID to the partial update app params
func (o *PartialUpdateAppParams) WithAppID(appID int64) *PartialUpdateAppParams {
	o.AppID = appID
	return o
}

// WithBody adds the body to the partial update app params
func (o *PartialUpdateAppParams) WithBody(body []*models.PatchAppRequest) *PartialUpdateAppParams {
	o.Body = body
	return o
}

// WithTeamID adds the teamID to the partial update app params
func (o *PartialUpdateAppParams) WithTeamID(teamID int64) *PartialUpdateAppParams {
	o.TeamID = teamID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PartialUpdateAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", swag.FormatInt64(o.AppID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
