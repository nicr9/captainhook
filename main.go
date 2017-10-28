package main

import (
	"fmt"
	"html/template"
	"k8s.io/client-go/rest"
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
	c.Manager.Create <- hook

	fmt.Fprintf(w, "Created a new hook: %+v\n", hook.Id)
}

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	
	manager := NewHookManager()
	controller := NewHookController(config)
	http.Handle("/", Homepage{manager})
	http.Handle("/create/", Creator{manager})
	http.Handle("/hookjob/", controller)

	log.Println("Starting server...")
	go manager.Run()
	http.ListenAndServe(":8080", nil)
}
