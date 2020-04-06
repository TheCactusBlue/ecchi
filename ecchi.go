package ecchi

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

var _ http.Handler = (*Router)(nil)

// Handler is the handler function
type Handler func(*Ctx) *Ctx

// Ctx is the context for ecchi
type Ctx struct {
	W http.ResponseWriter
	R *http.Request
}

// ReadJSON reads JSON Payload
func (c *Ctx) ReadJSON(dst interface{}) error {
	dec := json.NewDecoder(c.R.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		return err
	}

	err = validate.Struct(dst)
	if err != nil {
		return NewError("ValidationFailed", 400, err.Error())
	}

	return nil
}

// Code sets the status code. Defaults to 200.
func (c *Ctx) Code(code int) *Ctx {
	c.W.WriteHeader(code)
	return c
}

// URLParam gets URL param
func (c *Ctx) URLParam(key string) string {
	return chi.URLParam(c.R, key)
}

// JSON writes a JSON response.
func (c *Ctx) JSON(v interface{}) *Ctx {
	c.W.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(c.W).Encode(v)
	if err != nil {
		c.W.WriteHeader(500)
		c.W.Write(marshalErr)
	}
	return c
}

// Router routes ecchi
type Router struct {
	chi.Router
}

// NewRouter creates an ecchi router.
func NewRouter() *Router {
	return &Router{chi.NewRouter()}
}

// Wrap converts a ecchi handler into http handler.
func Wrap(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(&Ctx{
			W: w,
			R: r,
		})
	}
}
