/*
Superlink

API for Superlink

API version: v0.3.3
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiDomainMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDomainMetadataResponse{}

// ApiDomainMetadataResponse struct for ApiDomainMetadataResponse
type ApiDomainMetadataResponse struct {
	AnimationUrl *string `json:"animation_url,omitempty"`
	Attributes []ApiDomainMetadataAttribute `json:"attributes,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
	BackgroundColor *string `json:"background_color,omitempty"`
	Description *string `json:"description,omitempty"`
	ExternalUrl *string `json:"external_url,omitempty"`
	Image *string `json:"image,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	Namehash *string `json:"namehash,omitempty"`
	TokenId *string `json:"tokenId,omitempty"`
}

// NewApiDomainMetadataResponse instantiates a new ApiDomainMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDomainMetadataResponse() *ApiDomainMetadataResponse {
	this := ApiDomainMetadataResponse{}
	return &this
}

// NewApiDomainMetadataResponseWithDefaults instantiates a new ApiDomainMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDomainMetadataResponseWithDefaults() *ApiDomainMetadataResponse {
	this := ApiDomainMetadataResponse{}
	return &this
}

// GetAnimationUrl returns the AnimationUrl field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetAnimationUrl() string {
	if o == nil || IsNil(o.AnimationUrl) {
		var ret string
		return ret
	}
	return *o.AnimationUrl
}

// GetAnimationUrlOk returns a tuple with the AnimationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetAnimationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AnimationUrl) {
		return nil, false
	}
	return o.AnimationUrl, true
}

// HasAnimationUrl returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasAnimationUrl() bool {
	if o != nil && !IsNil(o.AnimationUrl) {
		return true
	}

	return false
}

// SetAnimationUrl gets a reference to the given string and assigns it to the AnimationUrl field.
func (o *ApiDomainMetadataResponse) SetAnimationUrl(v string) {
	o.AnimationUrl = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetAttributes() []ApiDomainMetadataAttribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []ApiDomainMetadataAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetAttributesOk() ([]ApiDomainMetadataAttribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []ApiDomainMetadataAttribute and assigns it to the Attributes field.
func (o *ApiDomainMetadataResponse) SetAttributes(v []ApiDomainMetadataAttribute) {
	o.Attributes = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *ApiDomainMetadataResponse) SetAvatar(v string) {
	o.Avatar = &v
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetBackgroundColor() string {
	if o == nil || IsNil(o.BackgroundColor) {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetBackgroundColorOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundColor) {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasBackgroundColor() bool {
	if o != nil && !IsNil(o.BackgroundColor) {
		return true
	}

	return false
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *ApiDomainMetadataResponse) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiDomainMetadataResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetExternalUrl() string {
	if o == nil || IsNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetExternalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUrl) {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasExternalUrl() bool {
	if o != nil && !IsNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *ApiDomainMetadataResponse) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ApiDomainMetadataResponse) SetImage(v string) {
	o.Image = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ApiDomainMetadataResponse) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiDomainMetadataResponse) SetName(v string) {
	o.Name = &v
}

// GetNamehash returns the Namehash field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetNamehash() string {
	if o == nil || IsNil(o.Namehash) {
		var ret string
		return ret
	}
	return *o.Namehash
}

// GetNamehashOk returns a tuple with the Namehash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetNamehashOk() (*string, bool) {
	if o == nil || IsNil(o.Namehash) {
		return nil, false
	}
	return o.Namehash, true
}

// HasNamehash returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasNamehash() bool {
	if o != nil && !IsNil(o.Namehash) {
		return true
	}

	return false
}

// SetNamehash gets a reference to the given string and assigns it to the Namehash field.
func (o *ApiDomainMetadataResponse) SetNamehash(v string) {
	o.Namehash = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ApiDomainMetadataResponse) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataResponse) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ApiDomainMetadataResponse) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ApiDomainMetadataResponse) SetTokenId(v string) {
	o.TokenId = &v
}

func (o ApiDomainMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDomainMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnimationUrl) {
		toSerialize["animation_url"] = o.AnimationUrl
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.BackgroundColor) {
		toSerialize["background_color"] = o.BackgroundColor
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalUrl) {
		toSerialize["external_url"] = o.ExternalUrl
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namehash) {
		toSerialize["namehash"] = o.Namehash
	}
	if !IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	return toSerialize, nil
}

type NullableApiDomainMetadataResponse struct {
	value *ApiDomainMetadataResponse
	isSet bool
}

func (v NullableApiDomainMetadataResponse) Get() *ApiDomainMetadataResponse {
	return v.value
}

func (v *NullableApiDomainMetadataResponse) Set(val *ApiDomainMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDomainMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDomainMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDomainMetadataResponse(val *ApiDomainMetadataResponse) *NullableApiDomainMetadataResponse {
	return &NullableApiDomainMetadataResponse{value: val, isSet: true}
}

func (v NullableApiDomainMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDomainMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


