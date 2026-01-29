// Code generated from our API definition. DO NOT EDIT.

package task

import (
	json "encoding/json"
	fmt "fmt"
	common "github.com/anduril/lattice-sdk-go/v4/common"
	internal "github.com/anduril/lattice-sdk-go/v4/internal"
	big "math/big"
)

// Represents a team of agents
var (
	teamFieldEntityID = big.NewInt(1 << 0)
	teamFieldMembers  = big.NewInt(1 << 1)
)

type Team struct {
	// Entity ID of the team
	EntityID *string         `json:"entityId,omitempty" url:"entityId,omitempty"`
	Members  []*common.Agent `json:"members,omitempty" url:"members,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *Team) GetEntityID() *string {
	if t == nil {
		return nil
	}
	return t.EntityID
}

func (t *Team) GetMembers() []*common.Agent {
	if t == nil {
		return nil
	}
	return t.Members
}

func (t *Team) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *Team) require(field *big.Int) {
	if t.explicitFields == nil {
		t.explicitFields = big.NewInt(0)
	}
	t.explicitFields.Or(t.explicitFields, field)
}

// SetEntityID sets the EntityID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *Team) SetEntityID(entityID *string) {
	t.EntityID = entityID
	t.require(teamFieldEntityID)
}

// SetMembers sets the Members field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *Team) SetMembers(members []*common.Agent) {
	t.Members = members
	t.require(teamFieldMembers)
}

func (t *Team) UnmarshalJSON(data []byte) error {
	type unmarshaler Team
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = Team(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *Team) MarshalJSON() ([]byte, error) {
	type embed Team
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *Team) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}
