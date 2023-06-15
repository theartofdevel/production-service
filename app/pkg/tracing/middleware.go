package tracing

import (
	"net/http"
	"strings"
	"sync"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type traceHandlers struct {
	mu   sync.RWMutex
	data map[string]http.Handler
}

func (h *traceHandlers) Get(path string) (http.Handler, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	handler, ok := h.data[path]

	return handler, ok
}

func (h *traceHandlers) Set(path string, handler http.Handler) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.data[path] = handler
}

func Middleware(next http.Handler) http.Handler {
	pathHandlers := traceHandlers{
		data: map[string]http.Handler{},
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		buf := strings.Builder{}
		buf.WriteString(r.Method)
		buf.WriteString(" ")

		var uri string
		if r.URL != nil {
			uri = r.URL.Path
		} else {
			uri = r.RequestURI
		}
		buf.WriteString(uri)

		path := buf.String()

		var h http.Handler

		h, ok := pathHandlers.Get(path)
		if !ok {
			h = otelhttp.NewHandler(
				next,
				path,
				otelhttp.WithPropagators(otel.GetTextMapPropagator()),
			)
			pathHandlers.Set(path, h)
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
