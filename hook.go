package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type HookManager struct {
	All []Hook
}

func NewHookManager() *HookManager {
	hooks := make([]Hook, 0)
	return &HookManager{hooks}
}

func (h *HookManager) AddHook(hook Hook) {
	h.All = append(h.All, hook)
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
