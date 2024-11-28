/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
	"fmt"
)

// checks if the PostV1TfaProviderU2fkeyVerifyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1TfaProviderU2fkeyVerifyRequest{}

// PostV1TfaProviderU2fkeyVerifyRequest struct for PostV1TfaProviderU2fkeyVerifyRequest
type PostV1TfaProviderU2fkeyVerifyRequest struct {
	Username                string `json:"username"`
	Password                string `json:"password"`
	PublicKeyCredentialJson string `json:"publicKeyCredentialJson"`
	AdditionalProperties    map[string]interface{}
}

type _PostV1TfaProviderU2fkeyVerifyRequest PostV1TfaProviderU2fkeyVerifyRequest

// NewPostV1TfaProviderU2fkeyVerifyRequest instantiates a new PostV1TfaProviderU2fkeyVerifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1TfaProviderU2fkeyVerifyRequest(username string, password string, publicKeyCredentialJson string) *PostV1TfaProviderU2fkeyVerifyRequest {
	this := PostV1TfaProviderU2fkeyVerifyRequest{}
	this.Username = username
	this.Password = password
	this.PublicKeyCredentialJson = publicKeyCredentialJson
	return &this
}

// NewPostV1TfaProviderU2fkeyVerifyRequestWithDefaults instantiates a new PostV1TfaProviderU2fkeyVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1TfaProviderU2fkeyVerifyRequestWithDefaults() *PostV1TfaProviderU2fkeyVerifyRequest {
	this := PostV1TfaProviderU2fkeyVerifyRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *PostV1TfaProviderU2fkeyVerifyRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *PostV1TfaProviderU2fkeyVerifyRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *PostV1TfaProviderU2fkeyVerifyRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *PostV1TfaProviderU2fkeyVerifyRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PostV1TfaProviderU2fkeyVerifyRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PostV1TfaProviderU2fkeyVerifyRequest) SetPassword(v string) {
	o.Password = v
}

// GetPublicKeyCredentialJson returns the PublicKeyCredentialJson field value
func (o *PostV1TfaProviderU2fkeyVerifyRequest) GetPublicKeyCredentialJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyCredentialJson
}

// GetPublicKeyCredentialJsonOk returns a tuple with the PublicKeyCredentialJson field value
// and a boolean to check if the value has been set.
func (o *PostV1TfaProviderU2fkeyVerifyRequest) GetPublicKeyCredentialJsonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyCredentialJson, true
}

// SetPublicKeyCredentialJson sets field value
func (o *PostV1TfaProviderU2fkeyVerifyRequest) SetPublicKeyCredentialJson(v string) {
	o.PublicKeyCredentialJson = v
}

func (o PostV1TfaProviderU2fkeyVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1TfaProviderU2fkeyVerifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["publicKeyCredentialJson"] = o.PublicKeyCredentialJson

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1TfaProviderU2fkeyVerifyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
		"publicKeyCredentialJson",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostV1TfaProviderU2fkeyVerifyRequest := _PostV1TfaProviderU2fkeyVerifyRequest{}

	err = json.Unmarshal(data, &varPostV1TfaProviderU2fkeyVerifyRequest)

	if err != nil {
		return err
	}

	*o = PostV1TfaProviderU2fkeyVerifyRequest(varPostV1TfaProviderU2fkeyVerifyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "publicKeyCredentialJson")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1TfaProviderU2fkeyVerifyRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1TfaProviderU2fkeyVerifyRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1TfaProviderU2fkeyVerifyRequest struct {
	value *PostV1TfaProviderU2fkeyVerifyRequest
	isSet bool
}

func (v NullablePostV1TfaProviderU2fkeyVerifyRequest) Get() *PostV1TfaProviderU2fkeyVerifyRequest {
	return v.value
}

func (v *NullablePostV1TfaProviderU2fkeyVerifyRequest) Set(val *PostV1TfaProviderU2fkeyVerifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1TfaProviderU2fkeyVerifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1TfaProviderU2fkeyVerifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1TfaProviderU2fkeyVerifyRequest(val *PostV1TfaProviderU2fkeyVerifyRequest) *NullablePostV1TfaProviderU2fkeyVerifyRequest {
	return &NullablePostV1TfaProviderU2fkeyVerifyRequest{value: val, isSet: true}
}

func (v NullablePostV1TfaProviderU2fkeyVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1TfaProviderU2fkeyVerifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}