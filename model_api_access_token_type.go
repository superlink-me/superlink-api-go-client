/*
Superlink

API for Superlink

API version: 1.0
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SuperlinkAPI

import (
	"encoding/json"
	"fmt"
)

// ApiAccessTokenType the model 'ApiAccessTokenType'
type ApiAccessTokenType string

// List of api.AccessTokenType
const (
	AccessTokenTypeAdmin ApiAccessTokenType = "ADMIN"
	AccessTokenTypeCustomer ApiAccessTokenType = "CUSTOMER"
)

// All allowed values of ApiAccessTokenType enum
var AllowedApiAccessTokenTypeEnumValues = []ApiAccessTokenType{
	"ADMIN",
	"CUSTOMER",
}

func (v *ApiAccessTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiAccessTokenType(value)
	for _, existing := range AllowedApiAccessTokenTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiAccessTokenType", value)
}

// NewApiAccessTokenTypeFromValue returns a pointer to a valid ApiAccessTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiAccessTokenTypeFromValue(v string) (*ApiAccessTokenType, error) {
	ev := ApiAccessTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiAccessTokenType: valid values are %v", v, AllowedApiAccessTokenTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiAccessTokenType) IsValid() bool {
	for _, existing := range AllowedApiAccessTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to api.AccessTokenType value
func (v ApiAccessTokenType) Ptr() *ApiAccessTokenType {
	return &v
}

type NullableApiAccessTokenType struct {
	value *ApiAccessTokenType
	isSet bool
}

func (v NullableApiAccessTokenType) Get() *ApiAccessTokenType {
	return v.value
}

func (v *NullableApiAccessTokenType) Set(val *ApiAccessTokenType) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccessTokenType) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccessTokenType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccessTokenType(val *ApiAccessTokenType) *NullableApiAccessTokenType {
	return &NullableApiAccessTokenType{value: val, isSet: true}
}

func (v NullableApiAccessTokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccessTokenType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

