// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMetricsFlowcountHandlerFunc turns a function with the right signature into a get metrics flowcount handler
type GetMetricsFlowcountHandlerFunc func(GetMetricsFlowcountParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMetricsFlowcountHandlerFunc) Handle(params GetMetricsFlowcountParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetMetricsFlowcountHandler interface for that can handle valid get metrics flowcount params
type GetMetricsFlowcountHandler interface {
	Handle(GetMetricsFlowcountParams, interface{}) middleware.Responder
}

// NewGetMetricsFlowcount creates a new http.Handler for the get metrics flowcount operation
func NewGetMetricsFlowcount(ctx *middleware.Context, handler GetMetricsFlowcountHandler) *GetMetricsFlowcount {
	return &GetMetricsFlowcount{Context: ctx, Handler: handler}
}

/*
	GetMetricsFlowcount swagger:route GET /metrics/flowcount getMetricsFlowcount

# Get flow count metrics

Get metrics related to flow counts.
*/
type GetMetricsFlowcount struct {
	Context *middleware.Context
	Handler GetMetricsFlowcountHandler
}

func (o *GetMetricsFlowcount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMetricsFlowcountParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
