package ecchi

import "encoding/json"

// RestError is an error for REST endpoints.
type RestError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrorJSON returns a JSON, as error.
func (c *Ctx) ErrorJSON(r *RestError) *Ctx {
	return c.JSON(r).Code(r.Code)
}

// ServerError is for unexpected server errors.
func (c *Ctx) ServerError(e error) *Ctx {
	c.JSON(RestError{
		Name:    "ServerFailure",
		Message: e.Error(),
		Code:    500,
	})
	return c
}

var marshalErr, _ = json.Marshal(RestError{
	Name:    "MarshalFailure",
	Message: "Failed to marshal JSON",
	Code:    500,
})
