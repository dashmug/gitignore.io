package main

import (
    "os"
    "fmt"
    "regexp"
    "strings"
    "net/http"
    "io/ioutil"
    "path/filepath"
    "github.com/gorilla/mux"
)

var Datastore = struct {
  ignoreList []string
  ignoreTemplates map[string]IgnoreTemplate
  dropdownItems []DropdownItem
}{}

type DropdownItem struct {
  Id string `json:"id"`
  Text string `json:"text"`
}

type IgnoreTemplate struct {
  name string
  fileName string
  contents string
}

var Metrics = struct {
  templateCount int
}{ 0 }

func main() {

  Datastore.ignoreTemplates = make(map[string]IgnoreTemplate)
  filepath.Walk("data", LoadAllTemplates)
  // files,_ := ioutil.ReadDir("data")

  // Metrics.templateCount = len(files)
  // fmt.Println(Metrics.templateCount)
  // fmt.Println(Datastore.ignoreTemplates)

  r := mux.NewRouter()
  // Private API
  r.HandleFunc("/dropdown/templates.json", DropdownHandler)

  // Home
  r.HandleFunc("/", HomeHandler)
  r.HandleFunc("/docs", DocumentationHandler)

  // Public API
  r.HandleFunc("/list", TemplateListHandler)
  r.HandleFunc("/help", CommandLineHelpHandler)
  r.HandleFunc("/api/{ignore}", IgnoreWebFileHandler)
  r.HandleFunc("/api/f/{ignore}", IgnoreDownloadFileHandler)

  // Static files
  r.HandleFunc("/README.md", ReadmeHandler)
  s := http.StripPrefix("/", http.FileServer(http.Dir("./public/")))
  r.PathPrefix("/").Handler(s)

  fmt.Println("Starting server on :8080")
  http.ListenAndServe(":8080", r)
}

func LoadAllTemplates(fp string, fi os.FileInfo, err error) error {
    if err != nil {
        // fmt.Println(err) // can't walk here,
        return nil       // but continue walking elsewhere
    }
    if !!fi.IsDir() {
        // fmt.Println("directory:"+fp)
        return nil // not a file.
    }

    rp := regexp.MustCompile("[a-zA-Z1-9].gitignore$")
    foundBool := rp.MatchString(fp)

    if foundBool && !fi.IsDir() {
      templateName := strings.Split(fi.Name(), ".")[0]
      lowercaseTemplateName := strings.ToLower(templateName)
      templateContents,_ := ioutil.ReadFile(fp)

      // Create ignore string
      Datastore.ignoreList = append(Datastore.ignoreList, lowercaseTemplateName)

      // Create dropdown item
      dropdownItem := DropdownItem{lowercaseTemplateName, templateName}

      // Create ignore template
      ignoreTemplate := IgnoreTemplate{templateName, fi.Name(), string(templateContents)}

      Datastore.ignoreTemplates[lowercaseTemplateName] = ignoreTemplate
      Datastore.dropdownItems = append(Datastore.dropdownItems, dropdownItem)

      Metrics.templateCount += 1
    }
    return nil
}
