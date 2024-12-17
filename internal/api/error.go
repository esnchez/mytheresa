package api

import "net/http"

func (a *App) internalServerError(w http.ResponseWriter) {
	writeJSON(w, http.StatusInternalServerError, map[string]string{"error" : "server had an internal problem"})
}

func (a *App) badRequest(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusBadRequest, map[string]string{"error" : message})
}

