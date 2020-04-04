package ecchi

import "github.com/go-chi/chi"

// Route creates sub-routes.
func (r *Router) Route(route string, c func(*Router)) {
	r.Router.Route(route, func(r chi.Router) {
		c(&Router{chi.NewRouter()})
	})
}

// Connect is a wrapper method.
func (r *Router) Connect(route string, hf Handler) {
	r.Router.Connect(route, Wrap(hf))
}

// Delete is a wrapper method.
func (r *Router) Delete(route string, hf Handler) {
	r.Router.Delete(route, Wrap(hf))
}

// Get is a wrapper method.
func (r *Router) Get(route string, hf Handler) {
	r.Router.Get(route, Wrap(hf))
}

// Head is a wrapper method.
func (r *Router) Head(route string, hf Handler) {
	r.Router.Head(route, Wrap(hf))
}

// Options is a wrapper method.
func (r *Router) Options(route string, hf Handler) {
	r.Router.Options(route, Wrap(hf))
}

// Patch is a wrapper method.
func (r *Router) Patch(route string, hf Handler) {
	r.Router.Patch(route, Wrap(hf))
}

// Post is a wrapper method.
func (r *Router) Post(route string, hf Handler) {
	r.Router.Post(route, Wrap(hf))
}

// Put is a wrapper method.
func (r *Router) Put(route string, hf Handler) {
	r.Router.Put(route, Wrap(hf))
}

// Trace is a wrapper method.
func (r *Router) Trace(route string, hf Handler) {
	r.Router.Trace(route, Wrap(hf))
}
