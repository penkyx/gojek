/*
goid.gojekapi.com

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goid_gojek_go

import (
	"encoding/json"
)

// GenerateTokenResponse struct for GenerateTokenResponse
type GenerateTokenResponse struct {
	AccessToken *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	DblEnabled *bool `json:"dbl_enabled,omitempty"`
	Flags *GenerateTokenResponseFlags `json:"flags,omitempty"`
}

// NewGenerateTokenResponse instantiates a new GenerateTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateTokenResponse() *GenerateTokenResponse {
	this := GenerateTokenResponse{}
	return &this
}

// NewGenerateTokenResponseWithDefaults instantiates a new GenerateTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateTokenResponseWithDefaults() *GenerateTokenResponse {
	this := GenerateTokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *GenerateTokenResponse) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *GenerateTokenResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *GenerateTokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *GenerateTokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *GenerateTokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *GenerateTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetDblEnabled returns the DblEnabled field value if set, zero value otherwise.
func (o *GenerateTokenResponse) GetDblEnabled() bool {
	if o == nil || o.DblEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DblEnabled
}

// GetDblEnabledOk returns a tuple with the DblEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateTokenResponse) GetDblEnabledOk() (*bool, bool) {
	if o == nil || o.DblEnabled == nil {
		return nil, false
	}
	return o.DblEnabled, true
}

// HasDblEnabled returns a boolean if a field has been set.
func (o *GenerateTokenResponse) HasDblEnabled() bool {
	if o != nil && o.DblEnabled != nil {
		return true
	}

	return false
}

// SetDblEnabled gets a reference to the given bool and assigns it to the DblEnabled field.
func (o *GenerateTokenResponse) SetDblEnabled(v bool) {
	o.DblEnabled = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *GenerateTokenResponse) GetFlags() GenerateTokenResponseFlags {
	if o == nil || o.Flags == nil {
		var ret GenerateTokenResponseFlags
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateTokenResponse) GetFlagsOk() (*GenerateTokenResponseFlags, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *GenerateTokenResponse) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given GenerateTokenResponseFlags and assigns it to the Flags field.
func (o *GenerateTokenResponse) SetFlags(v GenerateTokenResponseFlags) {
	o.Flags = &v
}

func (o GenerateTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if o.DblEnabled != nil {
		toSerialize["dbl_enabled"] = o.DblEnabled
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateTokenResponse struct {
	value *GenerateTokenResponse
	isSet bool
}

func (v NullableGenerateTokenResponse) Get() *GenerateTokenResponse {
	return v.value
}

func (v *NullableGenerateTokenResponse) Set(val *GenerateTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateTokenResponse(val *GenerateTokenResponse) *NullableGenerateTokenResponse {
	return &NullableGenerateTokenResponse{value: val, isSet: true}
}

func (v NullableGenerateTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


