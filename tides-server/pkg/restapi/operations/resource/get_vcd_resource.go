// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVcdResourceHandlerFunc turns a function with the right signature into a get vcd resource handler
type GetVcdResourceHandlerFunc func(GetVcdResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVcdResourceHandlerFunc) Handle(params GetVcdResourceParams) middleware.Responder {
	return fn(params)
}

// GetVcdResourceHandler interface for that can handle valid get vcd resource params
type GetVcdResourceHandler interface {
	Handle(GetVcdResourceParams) middleware.Responder
}

// NewGetVcdResource creates a new http.Handler for the get vcd resource operation
func NewGetVcdResource(ctx *middleware.Context, handler GetVcdResourceHandler) *GetVcdResource {
	return &GetVcdResource{Context: ctx, Handler: handler}
}

/* GetVcdResource swagger:route GET /resource/vcd/{id} resource getVcdResource

get vcd resource

*/
type GetVcdResource struct {
	Context *middleware.Context
	Handler GetVcdResourceHandler
}

func (o *GetVcdResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVcdResourceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetVcdResourceOKBody get vcd resource o k body
//
// swagger:model GetVcdResourceOKBody
type GetVcdResourceOKBody struct {

	// allocation model
	AllocationModel string `json:"allocationModel,omitempty"`

	// current CPU
	CurrentCPU float64 `json:"currentCPU,omitempty"`

	// current disk
	CurrentDisk float64 `json:"currentDisk,omitempty"`

	// current RAM
	CurrentRAM float64 `json:"currentRAM,omitempty"`

	// datacenter
	Datacenter string `json:"datacenter,omitempty"`

	// href
	Href string `json:"href,omitempty"`

	// is active
	IsActive bool `json:"isActive,omitempty"`

	// job completed
	JobCompleted int64 `json:"jobCompleted,omitempty"`

	// monitored
	Monitored bool `json:"monitored,omitempty"`

	// organization
	Organization string `json:"organization,omitempty"`

	// policy
	Policy int64 `json:"policy,omitempty"`

	// setup status
	SetupStatus string `json:"setupStatus,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// total CPU
	TotalCPU float64 `json:"totalCPU,omitempty"`

	// total disk
	TotalDisk float64 `json:"totalDisk,omitempty"`

	// total jobs
	TotalJobs int64 `json:"totalJobs,omitempty"`

	// total RAM
	TotalRAM float64 `json:"totalRAM,omitempty"`

	// total v ms
	TotalVMs int64 `json:"totalVMs,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`
}

// Validate validates this get vcd resource o k body
func (o *GetVcdResourceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get vcd resource o k body based on context it is used
func (o *GetVcdResourceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetVcdResourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVcdResourceOKBody) UnmarshalBinary(b []byte) error {
	var res GetVcdResourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
