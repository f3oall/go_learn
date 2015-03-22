package main

import "github.com/vtg/flash"

type PrdContr struct {
	flash.Controller
}

func (p *PrdContr) Index() {
	p.RenderJSON(200, flash.JSON{"products": allPrds})
}
func (p *PrdContr) Show() {
	product := allPrds.FindByID(p.ID64())
	if product == nil {
		p.RenderJSONError(404, "record not found")
		return
	}
	p.RenderJSON(200, flash.JSON{"product": product})
}

func (p *PrdContr) Create() {
	prd := Product{}
	if prd.Name == "" {
		p.LoadJSONRequest("product", &prd)
		p.RenderJSONError(422, "name required")
	} else {
		allPrds.Append(prd)
		p.RenderJSON(200, flash.JSON{"product": prd})
	}
}
func (p *PrdContr) Update() {
	id := p.ID64()
	product := allPrds.FindByID(id)

	if product == nil {
		p.RenderJSONError(404, "record not found")
		return
	}

	prd := Product{}
	p.LoadJSONRequest("product", &prd)
	allPrds.Edit(int(id), prd)
	p.RenderJSON(200, flash.JSON{"product": product})
}
func (p *PrdContr) Destroy() {
	id := p.ID64()
	product := allPrds.FindByID(id)

	if product == nil {
		p.RenderJSONError(404, "record not found")
		return
	}

	allPrds.Remove(int(id))
	p.RenderJSON(203, flash.JSON{})
}
