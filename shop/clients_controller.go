package main

import "github.com/vtg/flash"

type ClContr struct {
	flash.Controller
}

func (c *ClContr) Index() {
	c.RenderJSON(200, flash.JSON{"clients": allCls})
}
func (c *ClContr) Show() {
	client := allCls.FindByID(c.ID64())
	if client == nil {
		c.RenderJSONError(404, "record not found")
		return
	}
	c.RenderJSON(200, flash.JSON{"client": client})
}

func (c *ClContr) Create() {
	cl := Client{}
	if cl.Name == "" {
		c.LoadJSONRequest("client", &cl)
		c.RenderJSONError(422, "name required")
	} else {
		allCls.Append(cl)
		c.RenderJSON(200, flash.JSON{"client": cl})
	}
}
func (c *ClContr) Update() {
	id := c.ID64()
	client := allCls.FindByID(id)

	if client == nil {
		c.RenderJSONError(404, "record not found")
		return
	}

	cl := Client{}
	c.LoadJSONRequest("client", &cl)
	allCls.Edit(int(id), cl)
	c.RenderJSON(200, flash.JSON{"client": client})
}
func (c *ClContr) Destroy() {
	id := c.ID64()
	client := allCls.FindByID(id)

	if client == nil {
		c.RenderJSONError(404, "record not found")
		return
	}

	allCls.Remove(int(id))
	c.RenderJSON(203, flash.JSON{})
}
