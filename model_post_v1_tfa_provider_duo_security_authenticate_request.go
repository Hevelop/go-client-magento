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

// checks if the PostV1TfaProviderDuoSecurityAuthenticateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1TfaProviderDuoSecurityAuthenticateRequest{}

// PostV1TfaProviderDuoSecurityAuthenticateRequest struct for PostV1TfaProviderDuoSecurityAuthenticateRequest
type PostV1TfaProviderDuoSecurityAuthenticateRequest struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	SignatureResponse    string `json:"signatureResponse"`
	AdditionalProperties map[string]interface{}
}

type _PostV1TfaProviderDuoSecurityAuthenticateRequest PostV1TfaProviderDuoSecurityAuthenticateRequest

// NewPostV1TfaProviderDuoSecurityAuthenticateRequest instantiates a new PostV1TfaProviderDuoSecurityAuthenticateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1TfaProviderDuoSecurityAuthenticateRequest(username string, password string, signatureResponse string) *PostV1TfaProviderDuoSecurityAuthenticateRequest {
	this := PostV1TfaProviderDuoSecurityAuthenticateRequest{}
	this.Username = username
	this.Password = password
	this.SignatureResponse = signatureResponse
	return &this
}

// NewPostV1TfaProviderDuoSecurityAuthenticateRequestWithDefaults instantiates a new PostV1TfaProviderDuoSecurityAuthenticateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1TfaProviderDuoSecurityAuthenticateRequestWithDefaults() *PostV1TfaProviderDuoSecurityAuthenticateRequest {
	this := PostV1TfaProviderDuoSecurityAuthenticateRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) SetPassword(v string) {
	o.Password = v
}

// GetSignatureResponse returns the SignatureResponse field value
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) GetSignatureResponse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureResponse
}

// GetSignatureResponseOk returns a tuple with the SignatureResponse field value
// and a boolean to check if the value has been set.
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) GetSignatureResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureResponse, true
}

// SetSignatureResponse sets field value
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) SetSignatureResponse(v string) {
	o.SignatureResponse = v
}

func (o PostV1TfaProviderDuoSecurityAuthenticateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1TfaProviderDuoSecurityAuthenticateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["signatureResponse"] = o.SignatureResponse

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
		"signatureResponse",
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

	varPostV1TfaProviderDuoSecurityAuthenticateRequest := _PostV1TfaProviderDuoSecurityAuthenticateRequest{}

	err = json.Unmarshal(data, &varPostV1TfaProviderDuoSecurityAuthenticateRequest)

	if err != nil {
		return err
	}

	*o = PostV1TfaProviderDuoSecurityAuthenticateRequest(varPostV1TfaProviderDuoSecurityAuthenticateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "signatureResponse")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1TfaProviderDuoSecurityAuthenticateRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1TfaProviderDuoSecurityAuthenticateRequest struct {
	value *PostV1TfaProviderDuoSecurityAuthenticateRequest
	isSet bool
}

func (v NullablePostV1TfaProviderDuoSecurityAuthenticateRequest) Get() *PostV1TfaProviderDuoSecurityAuthenticateRequest {
	return v.value
}

func (v *NullablePostV1TfaProviderDuoSecurityAuthenticateRequest) Set(val *PostV1TfaProviderDuoSecurityAuthenticateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1TfaProviderDuoSecurityAuthenticateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1TfaProviderDuoSecurityAuthenticateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1TfaProviderDuoSecurityAuthenticateRequest(val *PostV1TfaProviderDuoSecurityAuthenticateRequest) *NullablePostV1TfaProviderDuoSecurityAuthenticateRequest {
	return &NullablePostV1TfaProviderDuoSecurityAuthenticateRequest{value: val, isSet: true}
}

func (v NullablePostV1TfaProviderDuoSecurityAuthenticateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1TfaProviderDuoSecurityAuthenticateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
