package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouteAdapter struct {
	chi *chi.Mux
}

func NewChiRouteAdapter() *chiRouteAdapter {
	return &chiRouteAdapter{
		chi: chi.NewMux(),
	}
}

func (a *chiRouteAdapter) Use(middlewares ...func(http.Handler) http.Handler) {
	a.chi.Use(middlewares...)
}

func (a *chiRouteAdapter) With(middlewares ...func(http.Handler) http.Handler) Router {
	a.chi.With(middlewares...)

	return a
}

func (a *chiRouteAdapter) Route(pattern string, fn func(r Router)) Router {
	chiFn := func(r chi.Router) {
		fn(a)
	}

	a.chi.Route(pattern, chiFn)
	return a
}

// HTTP-method routing along `pattern`
func (a *chiRouteAdapter) Connect(pattern string, h http.HandlerFunc) {
	a.chi.Connect(pattern, h)
}

func (a *chiRouteAdapter) Delete(pattern string, h http.HandlerFunc) {
	a.chi.Delete(pattern, h)
}

func (a *chiRouteAdapter) Get(pattern string, h http.HandlerFunc) {
	a.chi.Get(pattern, h)
}

func (a *chiRouteAdapter) Head(pattern string, h http.HandlerFunc) {
	a.chi.Head(pattern, h)
}

func (a *chiRouteAdapter) Options(pattern string, h http.HandlerFunc) {
	a.chi.Get(pattern, h)
}

func (a *chiRouteAdapter) Patch(pattern string, h http.HandlerFunc) {
	a.chi.Options(pattern, h)
}

func (a *chiRouteAdapter) Post(pattern string, h http.HandlerFunc) {
	a.chi.Post(pattern, h)
}

func (a *chiRouteAdapter) Put(pattern string, h http.HandlerFunc) {
	a.chi.Put(pattern, h)
}

func (a *chiRouteAdapter) Trace(pattern string, h http.HandlerFunc) {
	a.chi.Trace(pattern, h)
}

func (a *chiRouteAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.chi.ServeHTTP(w, r)
}
