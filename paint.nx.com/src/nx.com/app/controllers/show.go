package controllers

import "github.com/revel/revel"

type Show struct {
	*revel.Controller
}

func (c Show) Index() revel.Result {
	return c.Render()
}
