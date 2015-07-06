package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

// Private API
func DropdownHandler(w http.ResponseWriter, r *http.Request) {
  jsonPrep := &Datastore.dropdownItems
  jsonResponse, _ := json.Marshal(jsonPrep)
  // fmt.Println(string(jsonResponse))

  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintln(w, string(jsonResponse))
}
