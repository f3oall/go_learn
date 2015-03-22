package main

import "github.com/vtg/flash"

type OrdContr struct {
	flash.Controller
}

func (o *OrdContr) Index() {
	o.RenderJSON(200, flash.JSON{"orders": allOrds})
}
func (o *OrdContr) Show() {
	order := allOrds.FindByID(o.ID64())
	if order == nil {
		o.RenderJSONError(404, "record not found")
		return
	}
	o.RenderJSON(200, flash.JSON{"order": order})
}

func (o *OrdContr) Create() {
	ord := Order{}
	if ord.Customer == "" {
		o.LoadJSONRequest("order", &ord)
		o.RenderJSONError(422, "customer required")
	} else {
		allOrds.Append(ord)
		o.RenderJSON(200, flash.JSON{"order": ord})
	}
}
func (o *OrdContr) Update() {
	id := o.ID64()
	order := allOrds.FindByID(id)

	if order == nil {
		o.RenderJSONError(404, "record not found")
		return
	}

	ord := Order{}
	o.LoadJSONRequest("order", &ord)
	allOrds.Edit(int(id), ord)
	o.RenderJSON(200, flash.JSON{"order": order})
}
func (o *OrdContr) Destroy() {
	id := o.ID64()
	order := allOrds.FindByID(id)

	if order == nil {
		o.RenderJSONError(404, "record not found")
		return
	}

	allOrds.Remove(int(id))
	o.RenderJSON(203, flash.JSON{})
}
