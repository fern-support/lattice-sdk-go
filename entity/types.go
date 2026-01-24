// Code generated from our API definition. DO NOT EDIT.

package entity

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/anduril/lattice-sdk-go/v5/internal"
	big "math/big"
)

var (
	badRequestErrorBodyFieldError            = big.NewInt(1 << 0)
	badRequestErrorBodyFieldErrorDescription = big.NewInt(1 << 1)
)

type BadRequestErrorBody struct {
	Error            string  `json:"error" url:"error"`
	ErrorDescription *string `json:"error_description,omitempty" url:"error_description,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BadRequestErrorBody) GetError() string {
	if b == nil {
		return ""
	}
	return b.Error
}

func (b *BadRequestErrorBody) GetErrorDescription() *string {
	if b == nil {
		return nil
	}
	return b.ErrorDescription
}

func (b *BadRequestErrorBody) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BadRequestErrorBody) require(field *big.Int) {
	if b.explicitFields == nil {
		b.explicitFields = big.NewInt(0)
	}
	b.explicitFields.Or(b.explicitFields, field)
}

// SetError sets the Error field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BadRequestErrorBody) SetError(error_ string) {
	b.Error = error_
	b.require(badRequestErrorBodyFieldError)
}

// SetErrorDescription sets the ErrorDescription field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BadRequestErrorBody) SetErrorDescription(errorDescription *string) {
	b.ErrorDescription = errorDescription
	b.require(badRequestErrorBodyFieldErrorDescription)
}

func (b *BadRequestErrorBody) UnmarshalJSON(data []byte) error {
	type unmarshaler BadRequestErrorBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BadRequestErrorBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BadRequestErrorBody) MarshalJSON() ([]byte, error) {
	type embed BadRequestErrorBody
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, b.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (b *BadRequestErrorBody) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

var (
	errorFieldErrorCode = big.NewInt(1 << 0)
	errorFieldMessage   = big.NewInt(1 << 1)
)

type Error struct {
	ErrorCode string `json:"errorCode" url:"errorCode"`
	Message   string `json:"message" url:"message"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *Error) GetErrorCode() string {
	if e == nil {
		return ""
	}
	return e.ErrorCode
}

func (e *Error) GetMessage() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func (e *Error) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *Error) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetErrorCode sets the ErrorCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *Error) SetErrorCode(errorCode string) {
	e.ErrorCode = errorCode
	e.require(errorFieldErrorCode)
}

// SetMessage sets the Message field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *Error) SetMessage(message string) {
	e.Message = message
	e.require(errorFieldMessage)
}

func (e *Error) UnmarshalJSON(data []byte) error {
	type unmarshaler Error
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = Error(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *Error) MarshalJSON() ([]byte, error) {
	type embed Error
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *Error) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// Datetime string in ISO 8601 format.
type Timestamp = string

var (
	unauthorizedErrorBodyFieldError            = big.NewInt(1 << 0)
	unauthorizedErrorBodyFieldErrorDescription = big.NewInt(1 << 1)
)

type UnauthorizedErrorBody struct {
	Error            string  `json:"error" url:"error"`
	ErrorDescription *string `json:"error_description,omitempty" url:"error_description,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UnauthorizedErrorBody) GetError() string {
	if u == nil {
		return ""
	}
	return u.Error
}

func (u *UnauthorizedErrorBody) GetErrorDescription() *string {
	if u == nil {
		return nil
	}
	return u.ErrorDescription
}

func (u *UnauthorizedErrorBody) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UnauthorizedErrorBody) require(field *big.Int) {
	if u.explicitFields == nil {
		u.explicitFields = big.NewInt(0)
	}
	u.explicitFields.Or(u.explicitFields, field)
}

// SetError sets the Error field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UnauthorizedErrorBody) SetError(error_ string) {
	u.Error = error_
	u.require(unauthorizedErrorBodyFieldError)
}

// SetErrorDescription sets the ErrorDescription field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UnauthorizedErrorBody) SetErrorDescription(errorDescription *string) {
	u.ErrorDescription = errorDescription
	u.require(unauthorizedErrorBodyFieldErrorDescription)
}

func (u *UnauthorizedErrorBody) UnmarshalJSON(data []byte) error {
	type unmarshaler UnauthorizedErrorBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnauthorizedErrorBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnauthorizedErrorBody) MarshalJSON() ([]byte, error) {
	type embed UnauthorizedErrorBody
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, u.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (u *UnauthorizedErrorBody) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
