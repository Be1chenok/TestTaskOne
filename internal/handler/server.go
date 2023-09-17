package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", h.ServerHealthChek)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			users, err := h.FindUsers(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Users: %v", string(users))))

		case http.MethodPost:
			user, err := h.NewUser(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("User created: %v", string(user))))
		}
	})

	return mux
}

func (h *Handler) ServerHealthChek(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Correctly"))
	default:
		http.Error(w, "Another method", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) GetURLParams(r *http.Request) map[string]string {
	urlParams := make(map[string]string)
	urlParams["size"] = r.URL.Query().Get("size")
	urlParams["age"] = r.URL.Query().Get("age")
	urlParams["country"] = r.URL.Query().Get("country")
	urlParams["city"] = r.URL.Query().Get("city")
	urlParams["street"] = r.URL.Query().Get("street")

	return urlParams
}
