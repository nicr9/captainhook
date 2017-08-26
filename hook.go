package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type HookManager struct {
	All    []Hook
	Create chan Hook
	Delete chan Hook
}

func NewHookManager() *HookManager {
	hooks := make([]Hook, 0)

	create := make(chan Hook)
	delete := make(chan Hook)

	return &HookManager{hooks, create, delete}
}

func (h *HookManager) Run() {
	for {
		select {
		case hook := <-h.Create:
			h.All = append(h.All, hook)
		case hook := <-h.Delete:
			for i, v := range h.All {
				if v == hook {
					h.All = append(h.All[:i], h.All[i+1:]...)
					break
				}
			}
		}
	}
}

type Hook struct {
	Id   uuid.UUID
	Path string
}

func (h Hook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hook ID: %s\n", h.Id)
}

func NewHook() Hook {
	id := uuid.New()
	url := fmt.Sprintf("/hook/%s", id)

	hook := Hook{id, url}
	http.Handle(url, hook)

	return hook
}
