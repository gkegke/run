package main

import (
)

/* 

  Generic helper functions

*/

func contains(s []string, t string) bool {
    for _, v := range s {
        if v == t {
            return true
        }
    }
    return false
}
