package main

import "net/http"

type HandlerBasedOnMap struct {
	//key应该是method+url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
}
func (h *HandlerBasedOnMap) key(mthod string, pattern string) string {
	return mthod + "#" + pattern
}
