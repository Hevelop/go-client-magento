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

// checks if the PostV1CustomersResetPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1CustomersResetPasswordRequest{}

// PostV1CustomersResetPasswordRequest struct for PostV1CustomersResetPasswordRequest
type PostV1CustomersResetPasswordRequest struct {
	// If empty value given then the customer will be matched by the RP token.
	Email                string `json:"email"`
	ResetToken           string `json:"resetToken"`
	NewPassword          string `json:"newPassword"`
	AdditionalProperties map[string]interface{}
}

type _PostV1CustomersResetPasswordRequest PostV1CustomersResetPasswordRequest

// NewPostV1CustomersResetPasswordRequest instantiates a new PostV1CustomersResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1CustomersResetPasswordRequest(email string, resetToken string, newPassword string) *PostV1CustomersResetPasswordRequest {
	this := PostV1CustomersResetPasswordRequest{}
	this.Email = email
	this.ResetToken = resetToken
	this.NewPassword = newPassword
	return &this
}

// NewPostV1CustomersResetPasswordRequestWithDefaults instantiates a new PostV1CustomersResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1CustomersResetPasswordRequestWithDefaults() *PostV1CustomersResetPasswordRequest {
	this := PostV1CustomersResetPasswordRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *PostV1CustomersResetPasswordRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PostV1CustomersResetPasswordRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PostV1CustomersResetPasswordRequest) SetEmail(v string) {
	o.Email = v
}

// GetResetToken returns the ResetToken field value
func (o *PostV1CustomersResetPasswordRequest) GetResetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetToken
}

// GetResetTokenOk returns a tuple with the ResetToken field value
// and a boolean to check if the value has been set.
func (o *PostV1CustomersResetPasswordRequest) GetResetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetToken, true
}

// SetResetToken sets field value
func (o *PostV1CustomersResetPasswordRequest) SetResetToken(v string) {
	o.ResetToken = v
}

// GetNewPassword returns the NewPassword field value
func (o *PostV1CustomersResetPasswordRequest) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *PostV1CustomersResetPasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *PostV1CustomersResetPasswordRequest) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o PostV1CustomersResetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1CustomersResetPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["resetToken"] = o.ResetToken
	toSerialize["newPassword"] = o.NewPassword

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1CustomersResetPasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"resetToken",
		"newPassword",
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

	varPostV1CustomersResetPasswordRequest := _PostV1CustomersResetPasswordRequest{}

	err = json.Unmarshal(data, &varPostV1CustomersResetPasswordRequest)

	if err != nil {
		return err
	}

	*o = PostV1CustomersResetPasswordRequest(varPostV1CustomersResetPasswordRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "resetToken")
		delete(additionalProperties, "newPassword")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1CustomersResetPasswordRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1CustomersResetPasswordRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1CustomersResetPasswordRequest struct {
	value *PostV1CustomersResetPasswordRequest
	isSet bool
}

func (v NullablePostV1CustomersResetPasswordRequest) Get() *PostV1CustomersResetPasswordRequest {
	return v.value
}

func (v *NullablePostV1CustomersResetPasswordRequest) Set(val *PostV1CustomersResetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1CustomersResetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1CustomersResetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1CustomersResetPasswordRequest(val *PostV1CustomersResetPasswordRequest) *NullablePostV1CustomersResetPasswordRequest {
	return &NullablePostV1CustomersResetPasswordRequest{value: val, isSet: true}
}

func (v NullablePostV1CustomersResetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1CustomersResetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
