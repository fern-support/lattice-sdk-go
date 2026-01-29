// Code generated from our API definition. DO NOT EDIT.

package common

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/anduril/lattice-sdk-go/v4/internal"
	big "math/big"
)

// Represents an agent capable of processing tasks.
var (
	agentFieldEntityID = big.NewInt(1 << 0)
)

type Agent struct {
	// Entity ID of the agent.
	EntityID *string `json:"entityId,omitempty" url:"entityId,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *Agent) GetEntityID() *string {
	if a == nil {
		return nil
	}
	return a.EntityID
}

func (a *Agent) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *Agent) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetEntityID sets the EntityID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *Agent) SetEntityID(entityID *string) {
	a.EntityID = entityID
	a.require(agentFieldEntityID)
}

func (a *Agent) UnmarshalJSON(data []byte) error {
	type unmarshaler Agent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = Agent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *Agent) MarshalJSON() ([]byte, error) {
	type embed Agent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *Agent) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}
