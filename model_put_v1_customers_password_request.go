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

// checks if the PutV1CustomersPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1CustomersPasswordRequest{}

// PutV1CustomersPasswordRequest struct for PutV1CustomersPasswordRequest
type PutV1CustomersPasswordRequest struct {
	Email                string `json:"email"`
	Template             string `json:"template"`
	WebsiteId            *int32 `json:"websiteId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PutV1CustomersPasswordRequest PutV1CustomersPasswordRequest

// NewPutV1CustomersPasswordRequest instantiates a new PutV1CustomersPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1CustomersPasswordRequest(email string, template string) *PutV1CustomersPasswordRequest {
	this := PutV1CustomersPasswordRequest{}
	this.Email = email
	this.Template = template
	return &this
}

// NewPutV1CustomersPasswordRequestWithDefaults instantiates a new PutV1CustomersPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1CustomersPasswordRequestWithDefaults() *PutV1CustomersPasswordRequest {
	this := PutV1CustomersPasswordRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *PutV1CustomersPasswordRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PutV1CustomersPasswordRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PutV1CustomersPasswordRequest) SetEmail(v string) {
	o.Email = v
}

// GetTemplate returns the Template field value
func (o *PutV1CustomersPasswordRequest) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *PutV1CustomersPasswordRequest) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *PutV1CustomersPasswordRequest) SetTemplate(v string) {
	o.Template = v
}

// GetWebsiteId returns the WebsiteId field value if set, zero value otherwise.
func (o *PutV1CustomersPasswordRequest) GetWebsiteId() int32 {
	if o == nil || IsNil(o.WebsiteId) {
		var ret int32
		return ret
	}
	return *o.WebsiteId
}

// GetWebsiteIdOk returns a tuple with the WebsiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutV1CustomersPasswordRequest) GetWebsiteIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebsiteId) {
		return nil, false
	}
	return o.WebsiteId, true
}

// HasWebsiteId returns a boolean if a field has been set.
func (o *PutV1CustomersPasswordRequest) HasWebsiteId() bool {
	if o != nil && !IsNil(o.WebsiteId) {
		return true
	}

	return false
}

// SetWebsiteId gets a reference to the given int32 and assigns it to the WebsiteId field.
func (o *PutV1CustomersPasswordRequest) SetWebsiteId(v int32) {
	o.WebsiteId = &v
}

func (o PutV1CustomersPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1CustomersPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["template"] = o.Template
	if !IsNil(o.WebsiteId) {
		toSerialize["websiteId"] = o.WebsiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1CustomersPasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"template",
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

	varPutV1CustomersPasswordRequest := _PutV1CustomersPasswordRequest{}

	err = json.Unmarshal(data, &varPutV1CustomersPasswordRequest)

	if err != nil {
		return err
	}

	*o = PutV1CustomersPasswordRequest(varPutV1CustomersPasswordRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "template")
		delete(additionalProperties, "websiteId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1CustomersPasswordRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1CustomersPasswordRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1CustomersPasswordRequest struct {
	value *PutV1CustomersPasswordRequest
	isSet bool
}

func (v NullablePutV1CustomersPasswordRequest) Get() *PutV1CustomersPasswordRequest {
	return v.value
}

func (v *NullablePutV1CustomersPasswordRequest) Set(val *PutV1CustomersPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1CustomersPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1CustomersPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1CustomersPasswordRequest(val *PutV1CustomersPasswordRequest) *NullablePutV1CustomersPasswordRequest {
	return &NullablePutV1CustomersPasswordRequest{value: val, isSet: true}
}

func (v NullablePutV1CustomersPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1CustomersPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
