package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteIPAMIPParams creates a new DeleteIPAMIPParams object
// with the default values initialized.
func NewDeleteIPAMIPParams() DeleteIPAMIPParams {
	var ()
	return DeleteIPAMIPParams{}
}

// DeleteIPAMIPParams contains all the bound params for the delete IP a m IP operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteIPAMIP
type DeleteIPAMIPParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*IP address
	  Required: true
	  In: path
	*/
	IP string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteIPAMIPParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rIP, rhkIP, _ := route.Params.GetOK("ip")
	if err := o.bindIP(rIP, rhkIP, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteIPAMIPParams) bindIP(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.IP = raw

	return nil
}
