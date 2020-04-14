// Code generated by go-swagger; DO NOT EDIT.

package stage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProjectProjectNameStageStageNameParams creates a new GetProjectProjectNameStageStageNameParams object
// with the default values initialized.
func NewGetProjectProjectNameStageStageNameParams() GetProjectProjectNameStageStageNameParams {

	var (
		// initialize parameters with default values

		disableUpstreamSyncDefault = bool(false)
	)

	return GetProjectProjectNameStageStageNameParams{
		DisableUpstreamSync: &disableUpstreamSyncDefault,
	}
}

// GetProjectProjectNameStageStageNameParams contains all the bound params for the get project project name stage stage name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetProjectProjectNameStageStageName
type GetProjectProjectNameStageStageNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Disable sync of upstream repo before reading content
	  In: query
	  Default: false
	*/
	DisableUpstreamSync *bool
	/*Name of the project
	  Required: true
	  In: path
	*/
	ProjectName string
	/*Name of the stage
	  Required: true
	  In: path
	*/
	StageName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetProjectProjectNameStageStageNameParams() beforehand.
func (o *GetProjectProjectNameStageStageNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDisableUpstreamSync, qhkDisableUpstreamSync, _ := qs.GetOK("disableUpstreamSync")
	if err := o.bindDisableUpstreamSync(qDisableUpstreamSync, qhkDisableUpstreamSync, route.Formats); err != nil {
		res = append(res, err)
	}

	rProjectName, rhkProjectName, _ := route.Params.GetOK("projectName")
	if err := o.bindProjectName(rProjectName, rhkProjectName, route.Formats); err != nil {
		res = append(res, err)
	}

	rStageName, rhkStageName, _ := route.Params.GetOK("stageName")
	if err := o.bindStageName(rStageName, rhkStageName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDisableUpstreamSync binds and validates parameter DisableUpstreamSync from query.
func (o *GetProjectProjectNameStageStageNameParams) bindDisableUpstreamSync(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetProjectProjectNameStageStageNameParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("disableUpstreamSync", "query", "bool", raw)
	}
	o.DisableUpstreamSync = &value

	return nil
}

// bindProjectName binds and validates parameter ProjectName from path.
func (o *GetProjectProjectNameStageStageNameParams) bindProjectName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectName = raw

	return nil
}

// bindStageName binds and validates parameter StageName from path.
func (o *GetProjectProjectNameStageStageNameParams) bindStageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.StageName = raw

	return nil
}
