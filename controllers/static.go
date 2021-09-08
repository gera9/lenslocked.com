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
			"bootswatch", "views/static/home.gohtml"),
		Contact: views.NewView(
			"bootswatch", "views/static/contact.gohtml"),
		FAQ: views.NewView(
			"bootswatch", "views/static/faq.gohtml"),
	}
}
