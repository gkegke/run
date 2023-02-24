/*

Core server functions

  1) Runs app on given port.
  
  2) Opens a webview instance, and navigates to that port in the webview.

  3) Opens some golang -> js functions that are called on given triggers.

  wv

  main()

  returnError
  returnSuccess

  app()

*/

package main

import (
	"fmt"
	"log"
	"net/http"

    "github.com/gorilla/mux"

    "github.com/webview/webview"
)

var wv webview.WebView

func main() {

    debug := true

    go app()

    wv = webview.New(debug)

    defer wv.Destroy()

    wv.SetTitle("Run")
    wv.SetSize(600,500, webview.HintNone)
    wv.Navigate("http://localhost:1111/")

    /*
    wv.Dispatch(func() {
        // Inject JS to disable right click and inspect element
        wv.Eval(fmt.Sprintf(`
           $(document).on("contextmenu",function(e){
                 e.preventDefault();
           });
        `))
    })
    */

    wv.Run()

}

func returnError(wv webview.WebView, reason string) {

    wv.Dispatch(func() {
        wv.Eval(fmt.Sprintf(`
            $("#errors").text("` + reason + `");
            $("#errors").show(300);
        `))
    })

}

func returnSuccess(wv webview.WebView, data string) {

    wv.Dispatch(func() {
        wv.Eval(fmt.Sprintf(`
            $("#errors").text("");
            $("#errors").hide();
            update("` + data + `");
            success();
        `))
    })

}



func app() {

    /* Dynamic Config */

    env := &Env{
        Port : HOST_PORT,
    }

	r := mux.NewRouter().StrictSlash(true)

	for url, handler := range HANDLERS {
		r.Handle(url, Handler{env, handler})
	}

	http.Handle("/", r)

	http.Handle("/s/", http.StripPrefix("/s", http.FileServer(http.Dir("./static/"))))

	fmt.Println("serving from ", HOST_PORT)
    err := http.ListenAndServe(HOST_PORT, nil)

	if err != nil {
		log.Fatal(err)
	}

}

