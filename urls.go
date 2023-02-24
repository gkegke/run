/*

   Url handler mapping

   main domains

   TODO

*/

package main

import (
	"net/http"
)

///////////// Url : Handler
var (

	// Main domain handlers
	HANDLERS = map[string]func(*Env, http.ResponseWriter, *http.Request) error {

		// General
		"/": homeHandler,
        "/command/" : commandHandler,

	}

)

