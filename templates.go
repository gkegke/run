/*
   Helpful functions for templating & caching templates
   TODO
     - Implement template caching to speed up performance.
*/

package main

import (
	"html/template"
)

/*

  Helper functions related to templating

  InitTemplate
  InitTemplates

*/

// func map for templates -- functions too be used in templates
var Fm template.FuncMap = template.FuncMap{}


// initializes and returns a template ready for easy execution
func InitTemplate(name string, pathToTemplate string) *template.Template {
	t := template.New(name).Funcs(Fm)
	t = template.Must(t.ParseFiles(pathToTemplate))
	return t
}

// initializes and returns templates ready for easy execution
func InitTemplates(name string, ts ...string) *template.Template {
	t := template.New(name).Funcs(Fm)
	for _, pathToT := range ts {
		t = template.Must(t.ParseFiles(pathToT))
	}
	return t
}

