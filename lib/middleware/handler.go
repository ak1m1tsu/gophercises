package middleware

import "net/http"

type IHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(w, r, next)
}

type middleware struct {
	handler IHandler

	nextfn func(rw http.ResponseWriter, r *http.Request)
}

func newMiddleware(handler IHandler, next *middleware) middleware {
	return middleware{
		handler: handler,
		nextfn:  next.ServeHTTP,
	}
}

func (m middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(w, r, m.nextfn)
}

func Wrap(handler http.Handler) IHandler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(w, r)
		next(w, r)
	})
}

type Handler struct {
	middleware middleware
	handlers   []IHandler
}

func New(handlers ...IHandler) *Handler {
	return &Handler{
		handlers:   handlers,
		middleware: build(handlers),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.middleware.ServeHTTP(NewResponseWriter(w), r)
}

func (h *Handler) Use(handler IHandler) {
	if handler == nil {
		panic("handler cannot be nil")
	}

	h.handlers = append(h.handlers, handler)
	h.middleware = build(h.handlers)
}

func (h *Handler) UseHandler(handler http.Handler) {
	h.Use(Wrap(handler))
}

func (h *Handler) Handlers() []IHandler {
	return h.handlers
}

func build(handlers []IHandler) middleware {
	var next middleware
	switch {
	case len(handlers) == 0:
		return voidMiddleware()
	case len(handlers) > 1:
		next = build(handlers[1:])
	default:
		next = voidMiddleware()
	}
	return newMiddleware(handlers[0], &next)
}

func voidMiddleware() middleware {
	return newMiddleware(
		HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}),
		&middleware{},
	)
}
