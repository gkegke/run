/*

Handlers for main routes utilised by run

# Structs

 Response - helper struct for creating responses

# Handlers

 home
   - 

 commandHandler
   -

*/

package main

import (
    "fmt"
    "os/exec"
    "strings"

    "net/http"

    "encoding/json"
)

type Response map[string]interface{}

func (r Response) String() (string) {
    b, err := json.Marshal(r)
    if err != nil {
        return ""
    }
    return string(b)
}

// route : /
func homeHandler(env *Env, w http.ResponseWriter, r *http.Request) error {

	t := InitTemplates("app", ROOT_DIRECTORY + "/templates/app.html")
	t.ExecuteTemplate(w, "app", map[string]interface{}{
        "Commands" : COMMANDS,
    })

    return nil
}

// route : /command/
func commandHandler(env *Env, w http.ResponseWriter, r *http.Request) error {

    if r.Method == "POST" {

        w.Header().Set("Content-Type", "application/json")

        data := r.FormValue("command")

        fmt.Println(data)

        args := strings.Split(data, " ")


        // if first command not in valid COMMANDS

        if !contains(COMMANDS, args[0]) {

            fmt.Fprint(w, map[string]interface{}{
                "success" : false,
                "error" : "empty command ??",
            })

            return nil
        }

        // generate command
        var cmd *exec.Cmd

        if (len(args) > 1) {

            cmd = exec.Command(COMMAND_DIRECTORY + args[0], args[1:]...)

        } else {

            cmd = exec.Command(COMMAND_DIRECTORY + args[0])

        }

        // attempt execution and return relevant response
        if err := cmd.Run(); err != nil {

            fmt.Println("error")
            fmt.Fprint(w, Response{
                "success" : false,
                "error" : err.Error(),
            })

        } else {

            fmt.Println("success")
            fmt.Fprint(w, Response{
                "success" : true,
            })
        }

    }

    return nil

}
