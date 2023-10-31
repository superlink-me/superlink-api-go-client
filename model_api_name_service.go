/*
Superlink

API for Superlink

API version: v0.3.30
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
	"fmt"
)

// ApiNameService the model 'ApiNameService'
type ApiNameService string

// List of api.NameService
const (
	UD ApiNameService = "ud"
	ENS ApiNameService = "ens"
	SUPERLINK ApiNameService = "superlink"
	ICANN ApiNameService = "icann"
)

// All allowed values of ApiNameService enum
var AllowedApiNameServiceEnumValues = []ApiNameService{
	"ud",
	"ens",
	"superlink",
	"icann",
}

func (v *ApiNameService) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiNameService(value)
	for _, existing := range AllowedApiNameServiceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiNameService", value)
}

// NewApiNameServiceFromValue returns a pointer to a valid ApiNameService
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiNameServiceFromValue(v string) (*ApiNameService, error) {
	ev := ApiNameService(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiNameService: valid values are %v", v, AllowedApiNameServiceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiNameService) IsValid() bool {
	for _, existing := range AllowedApiNameServiceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to api.NameService value
func (v ApiNameService) Ptr() *ApiNameService {
	return &v
}

type NullableApiNameService struct {
	value *ApiNameService
	isSet bool
}

func (v NullableApiNameService) Get() *ApiNameService {
	return v.value
}

func (v *NullableApiNameService) Set(val *ApiNameService) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNameService) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNameService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNameService(val *ApiNameService) *NullableApiNameService {
	return &NullableApiNameService{value: val, isSet: true}
}

func (v NullableApiNameService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNameService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

