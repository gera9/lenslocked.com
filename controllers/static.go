package controllers

import "github.com/gera9/lenslocked.com/views"

type Static struct {
	Home    *views.View
	Contact *views.View
	FAQ     *views.View
}

func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootswatch", "static/home",
		),
		Contact: views.NewView(
			"bootswatch", "static/contact",
		),
		FAQ: views.NewView(
			"bootswatch", "static/faq",
		),
	}
}
