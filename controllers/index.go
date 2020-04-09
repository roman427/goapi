package controllers

import "github.com/nafisfaysal/goapi/views"

func NewIndex() *Index {
	return &Index{
		Homepage: views.NewView("bootstrap", "index/homepage"),
	}
}

type Index struct {
	Homepage *views.View
}
