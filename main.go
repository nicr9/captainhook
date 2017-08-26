package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Homepage struct {
	Manager *HookManager
}

func (h Homepage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, h.Manager)
}

type Creator struct {
	Manager *HookManager
}

func (c Creator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hook := NewHook()
	c.Manager.AddHook(hook)

	fmt.Fprintf(w, "Created a new hook: %+v\n", hook.Id)
}

func main() {
	hooks := NewHookManager()
	http.Handle("/", Homepage{hooks})
	http.Handle("/create/", Creator{hooks})

	log.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
