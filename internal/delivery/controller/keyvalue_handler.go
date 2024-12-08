package controller

import (
	"encoding/json"
	"kv-store/internal/usecase"
	"net/http"
)

type KeyValueHandler struct {
	service *usecase.KeyValueService
}

func NewKeyValueHandler(service *usecase.KeyValueService) *KeyValueHandler {
	return &KeyValueHandler{service: service}
}

func (h *KeyValueHandler) Set(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, value := data["key"], data["value"]
	if err := h.service.Set(key, value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *KeyValueHandler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, err := h.service.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte(value))
}

func (h *KeyValueHandler) Delete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if err := h.service.Delete(key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *KeyValueHandler) List(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
