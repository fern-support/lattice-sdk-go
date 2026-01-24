// Code generated from our API definition. DO NOT EDIT.

package entity

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/anduril/lattice-sdk-go/v5/internal"
	big "math/big"
)

var (
	getTokenRequestFieldCode         = big.NewInt(1 << 0)
	getTokenRequestFieldRedirectURI  = big.NewInt(1 << 1)
	getTokenRequestFieldClientID     = big.NewInt(1 << 2)
	getTokenRequestFieldClientSecret = big.NewInt(1 << 3)
	getTokenRequestFieldRefreshToken = big.NewInt(1 << 4)
	getTokenRequestFieldUsername     = big.NewInt(1 << 5)
	getTokenRequestFieldPassword     = big.NewInt(1 << 6)
	getTokenRequestFieldScope        = big.NewInt(1 << 7)
)

type GetTokenRequest struct {
	// The type of grant being requested
	// The authorization code (required for authorization_code grant type)
	Code *string `json:"code,omitempty" url:"-"`
	// The redirect URI (required for authorization_code grant type)
	RedirectURI *string `json:"redirect_uri,omitempty" url:"-"`
	// The client identifier
	ClientID *string `json:"client_id,omitempty" url:"-"`
	// The client secret
	ClientSecret *string `json:"client_secret,omitempty" url:"-"`
	// The refresh token (required for refresh_token grant type)
	RefreshToken *string `json:"refresh_token,omitempty" url:"-"`
	// The resource owner username (required for password grant type)
	Username *string `json:"username,omitempty" url:"-"`
	// The resource owner password (required for password grant type)
	Password *string `json:"password,omitempty" url:"-"`
	// The scope of the access request
	Scope     *string `json:"scope,omitempty" url:"-"`
	grantType string

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetTokenRequest) GrantType() string {
	return g.grantType
}

func (g *GetTokenRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetCode sets the Code field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetCode(code *string) {
	g.Code = code
	g.require(getTokenRequestFieldCode)
}

// SetRedirectURI sets the RedirectURI field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetRedirectURI(redirectURI *string) {
	g.RedirectURI = redirectURI
	g.require(getTokenRequestFieldRedirectURI)
}

// SetClientID sets the ClientID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetClientID(clientID *string) {
	g.ClientID = clientID
	g.require(getTokenRequestFieldClientID)
}

// SetClientSecret sets the ClientSecret field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetClientSecret(clientSecret *string) {
	g.ClientSecret = clientSecret
	g.require(getTokenRequestFieldClientSecret)
}

// SetRefreshToken sets the RefreshToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetRefreshToken(refreshToken *string) {
	g.RefreshToken = refreshToken
	g.require(getTokenRequestFieldRefreshToken)
}

// SetUsername sets the Username field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetUsername(username *string) {
	g.Username = username
	g.require(getTokenRequestFieldUsername)
}

// SetPassword sets the Password field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetPassword(password *string) {
	g.Password = password
	g.require(getTokenRequestFieldPassword)
}

// SetScope sets the Scope field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenRequest) SetScope(scope *string) {
	g.Scope = scope
	g.require(getTokenRequestFieldScope)
}

func (g *GetTokenRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler GetTokenRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*g = GetTokenRequest(body)
	g.grantType = "client_credentials"
	return nil
}

func (g *GetTokenRequest) MarshalJSON() ([]byte, error) {
	type embed GetTokenRequest
	var marshaler = struct {
		embed
		GrantType string `json:"grant_type"`
	}{
		embed:     embed(*g),
		GrantType: "client_credentials",
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	getTokenResponseFieldAccessToken  = big.NewInt(1 << 0)
	getTokenResponseFieldTokenType    = big.NewInt(1 << 1)
	getTokenResponseFieldExpiresIn    = big.NewInt(1 << 2)
	getTokenResponseFieldRefreshToken = big.NewInt(1 << 3)
	getTokenResponseFieldScope        = big.NewInt(1 << 4)
)

type GetTokenResponse struct {
	// The access token
	AccessToken string `json:"access_token" url:"access_token"`
	// The type of token (typically "Bearer")
	TokenType string `json:"token_type" url:"token_type"`
	// Lifetime of the access token in seconds
	ExpiresIn *int `json:"expires_in,omitempty" url:"expires_in,omitempty"`
	// The refresh token
	RefreshToken *string `json:"refresh_token,omitempty" url:"refresh_token,omitempty"`
	// The scope of the access token
	Scope *string `json:"scope,omitempty" url:"scope,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetTokenResponse) GetAccessToken() string {
	if g == nil {
		return ""
	}
	return g.AccessToken
}

func (g *GetTokenResponse) GetTokenType() string {
	if g == nil {
		return ""
	}
	return g.TokenType
}

func (g *GetTokenResponse) GetExpiresIn() *int {
	if g == nil {
		return nil
	}
	return g.ExpiresIn
}

func (g *GetTokenResponse) GetRefreshToken() *string {
	if g == nil {
		return nil
	}
	return g.RefreshToken
}

func (g *GetTokenResponse) GetScope() *string {
	if g == nil {
		return nil
	}
	return g.Scope
}

func (g *GetTokenResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetTokenResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAccessToken sets the AccessToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenResponse) SetAccessToken(accessToken string) {
	g.AccessToken = accessToken
	g.require(getTokenResponseFieldAccessToken)
}

// SetTokenType sets the TokenType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenResponse) SetTokenType(tokenType string) {
	g.TokenType = tokenType
	g.require(getTokenResponseFieldTokenType)
}

// SetExpiresIn sets the ExpiresIn field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenResponse) SetExpiresIn(expiresIn *int) {
	g.ExpiresIn = expiresIn
	g.require(getTokenResponseFieldExpiresIn)
}

// SetRefreshToken sets the RefreshToken field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenResponse) SetRefreshToken(refreshToken *string) {
	g.RefreshToken = refreshToken
	g.require(getTokenResponseFieldRefreshToken)
}

// SetScope sets the Scope field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTokenResponse) SetScope(scope *string) {
	g.Scope = scope
	g.require(getTokenResponseFieldScope)
}

func (g *GetTokenResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetTokenResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetTokenResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetTokenResponse) MarshalJSON() ([]byte, error) {
	type embed GetTokenResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetTokenResponse) String() string {
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}
