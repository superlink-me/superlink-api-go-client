/*
Superlink

API for Superlink

API version: v0.5.0
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
	"fmt"
)

// ApiSubdomainStatus the model 'ApiSubdomainStatus'
type ApiSubdomainStatus string

// List of api.SubdomainStatus
const (
	SubdomainStatusAvailable ApiSubdomainStatus = "available"
	SubdomainStatusUnavailable ApiSubdomainStatus = "unavailable"
	SubdomainStatusRevoked ApiSubdomainStatus = "revoked"
)

// All allowed values of ApiSubdomainStatus enum
var AllowedApiSubdomainStatusEnumValues = []ApiSubdomainStatus{
	"available",
	"unavailable",
	"revoked",
}

func (v *ApiSubdomainStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiSubdomainStatus(value)
	for _, existing := range AllowedApiSubdomainStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiSubdomainStatus", value)
}

// NewApiSubdomainStatusFromValue returns a pointer to a valid ApiSubdomainStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiSubdomainStatusFromValue(v string) (*ApiSubdomainStatus, error) {
	ev := ApiSubdomainStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiSubdomainStatus: valid values are %v", v, AllowedApiSubdomainStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiSubdomainStatus) IsValid() bool {
	for _, existing := range AllowedApiSubdomainStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to api.SubdomainStatus value
func (v ApiSubdomainStatus) Ptr() *ApiSubdomainStatus {
	return &v
}

type NullableApiSubdomainStatus struct {
	value *ApiSubdomainStatus
	isSet bool
}

func (v NullableApiSubdomainStatus) Get() *ApiSubdomainStatus {
	return v.value
}

func (v *NullableApiSubdomainStatus) Set(val *ApiSubdomainStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubdomainStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubdomainStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubdomainStatus(val *ApiSubdomainStatus) *NullableApiSubdomainStatus {
	return &NullableApiSubdomainStatus{value: val, isSet: true}
}

func (v NullableApiSubdomainStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubdomainStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
