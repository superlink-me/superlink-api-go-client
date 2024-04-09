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

// DataAccessTokenType the model 'DataAccessTokenType'
type DataAccessTokenType string

// List of data.AccessTokenType
const (
	AccessTokenTypeAdmin DataAccessTokenType = "ADMIN"
	AccessTokenTypeCustomer DataAccessTokenType = "CUSTOMER"
	AccessTokenTypeUser DataAccessTokenType = "USER"
	AccessTokenTypeWallet DataAccessTokenType = "WALLET"
	AccessTokenTypeInvalid DataAccessTokenType = "INVALID"
)

// All allowed values of DataAccessTokenType enum
var AllowedDataAccessTokenTypeEnumValues = []DataAccessTokenType{
	"ADMIN",
	"CUSTOMER",
	"USER",
	"WALLET",
	"INVALID",
}

func (v *DataAccessTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataAccessTokenType(value)
	for _, existing := range AllowedDataAccessTokenTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataAccessTokenType", value)
}

// NewDataAccessTokenTypeFromValue returns a pointer to a valid DataAccessTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataAccessTokenTypeFromValue(v string) (*DataAccessTokenType, error) {
	ev := DataAccessTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataAccessTokenType: valid values are %v", v, AllowedDataAccessTokenTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataAccessTokenType) IsValid() bool {
	for _, existing := range AllowedDataAccessTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to data.AccessTokenType value
func (v DataAccessTokenType) Ptr() *DataAccessTokenType {
	return &v
}

type NullableDataAccessTokenType struct {
	value *DataAccessTokenType
	isSet bool
}

func (v NullableDataAccessTokenType) Get() *DataAccessTokenType {
	return v.value
}

func (v *NullableDataAccessTokenType) Set(val *DataAccessTokenType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccessTokenType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccessTokenType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccessTokenType(val *DataAccessTokenType) *NullableDataAccessTokenType {
	return &NullableDataAccessTokenType{value: val, isSet: true}
}

func (v NullableDataAccessTokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccessTokenType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

