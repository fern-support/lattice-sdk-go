// Code generated from our API definition. DO NOT EDIT.

package task

import (
	json "encoding/json"
	core "github.com/anduril/lattice-sdk-go/v5/core"
)

// Bad request
type BadRequestError struct {
	*core.APIError
	Body interface{}
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

// The specified resource was not found
type NotFoundError struct {
	*core.APIError
	Body interface{}
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

// Unauthorized to access resource
type UnauthorizedError struct {
	*core.APIError
	Body interface{}
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}
