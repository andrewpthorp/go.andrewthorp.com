package main

import (
  "html/template"
  "log"
  "net/http"
)

func main() {
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("app/assets/css"))))
  http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("app/statics"))))
  http.HandleFunc("/", serveSite)
  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}

func serveSite(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.ParseFiles("app/tmpl/index.tmpl")
  if err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil {
    log.Println("Error:")
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}
