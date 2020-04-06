package ecchi

import "encoding/json"

// RestError is an error for REST endpoints.
type RestError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var _ error = (*RestError)(nil)

func (r *RestError) Error() string {
	return r.Message
}

// NewError creates a new error.
func NewError(name string, code int, message string) *RestError {
	return &RestError{
		Name:    name,
		Message: message,
		Code:    code,
	}
}

// ErrorJSON returns a JSON, as error.
func (c *Ctx) ErrorJSON(r error) *Ctx {
	v, ok := r.(*RestError)
	if ok {
		return c.Code(v.Code).JSON(v)
	}
	return c.Code(500).JSON(RestError{
		Name:    "ServerFailure",
		Message: r.Error(),
		Code:    500,
	})
}

var marshalErr, _ = json.Marshal(RestError{
	Name:    "MarshalFailure",
	Message: "Failed to marshal JSON",
	Code:    500,
})
