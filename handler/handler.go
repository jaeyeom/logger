package handler

import (
	"net/http"
)

type LoggerHandler struct {
	logger func(content string)
}

// ServeHTTP calls logger with content parameter and redirect to url
// parameter.
func (h LoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")
	h.logger(content)

	url := r.FormValue("url")
	if url == "" {
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// New returns a new LoggerHandler with the provided logger. When the
// handler is called, it calls the logger with the content parameter.
func New(logger func(content string)) *LoggerHandler {
	return &LoggerHandler{logger}
}

// HandleLogger registers the handler for the given pattern in the
// http.DefaultServeMux. The documentation for http.ServeMux explains
// how patterns are matched.
func HandleLogger(pattern string, logger func(content string)) {
	http.Handle(pattern, New(logger))
}
