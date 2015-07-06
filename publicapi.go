package main

import (
    "fmt"
    "sort"
    "strings"
    "net/http"
    "github.com/gorilla/mux"
    // "html/template"
)

// Public API
func TemplateListHandler(w http.ResponseWriter, r *http.Request) {
  list := Datastore.ignoreList
  sort.Strings(list)
  templateList := strings.Join(list[:],",")
  fmt.Fprintln(w, templateList)
}
func IgnoreWebFileHandler(w http.ResponseWriter, r *http.Request) {
  ignoreString := mux.Vars(r)["ignore"]
  arr := strings.Split(ignoreString, ",")
  fmt.Println(arr)



  fmt.Fprintln(w, "IgnoreWebFileHandler")
}

func IgnoreDownloadFileHandler(w http.ResponseWriter, r *http.Request) {
  // ignoreString := mux.Vars(r)["ignore"]
  fmt.Fprintln(w, "IgnoreDownloadFileHandler")
}

func CommandLineHelpHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "gitignore.io help:\n  list    - lists the operating systems, programming languages and IDE input types\n  :types: - creates .gitignore files for types of operating systems, programming languages or IDEs\n")
}

func generateFile(w []string) string {

}
