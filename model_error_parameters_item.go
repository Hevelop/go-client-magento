/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
)

// checks if the ErrorParametersItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorParametersItem{}

// ErrorParametersItem Error parameters item
type ErrorParametersItem struct {
	// ACL resource
	Resources *string `json:"resources,omitempty"`
	// Missing or invalid field name
	FieldName *string `json:"fieldName,omitempty"`
	// Incorrect field value
	FieldValue           *string `json:"fieldValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorParametersItem ErrorParametersItem

// NewErrorParametersItem instantiates a new ErrorParametersItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorParametersItem() *ErrorParametersItem {
	this := ErrorParametersItem{}
	return &this
}

// NewErrorParametersItemWithDefaults instantiates a new ErrorParametersItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorParametersItemWithDefaults() *ErrorParametersItem {
	this := ErrorParametersItem{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ErrorParametersItem) GetResources() string {
	if o == nil || IsNil(o.Resources) {
		var ret string
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorParametersItem) GetResourcesOk() (*string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ErrorParametersItem) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given string and assigns it to the Resources field.
func (o *ErrorParametersItem) SetResources(v string) {
	o.Resources = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *ErrorParametersItem) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorParametersItem) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *ErrorParametersItem) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *ErrorParametersItem) SetFieldName(v string) {
	o.FieldName = &v
}

// GetFieldValue returns the FieldValue field value if set, zero value otherwise.
func (o *ErrorParametersItem) GetFieldValue() string {
	if o == nil || IsNil(o.FieldValue) {
		var ret string
		return ret
	}
	return *o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorParametersItem) GetFieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.FieldValue) {
		return nil, false
	}
	return o.FieldValue, true
}

// HasFieldValue returns a boolean if a field has been set.
func (o *ErrorParametersItem) HasFieldValue() bool {
	if o != nil && !IsNil(o.FieldValue) {
		return true
	}

	return false
}

// SetFieldValue gets a reference to the given string and assigns it to the FieldValue field.
func (o *ErrorParametersItem) SetFieldValue(v string) {
	o.FieldValue = &v
}

func (o ErrorParametersItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorParametersItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.FieldName) {
		toSerialize["fieldName"] = o.FieldName
	}
	if !IsNil(o.FieldValue) {
		toSerialize["fieldValue"] = o.FieldValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorParametersItem) UnmarshalJSON(data []byte) (err error) {
	varErrorParametersItem := _ErrorParametersItem{}

	err = json.Unmarshal(data, &varErrorParametersItem)

	if err != nil {
		return err
	}

	*o = ErrorParametersItem(varErrorParametersItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resources")
		delete(additionalProperties, "fieldName")
		delete(additionalProperties, "fieldValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ErrorParametersItem) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ErrorParametersItem) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableErrorParametersItem struct {
	value *ErrorParametersItem
	isSet bool
}

func (v NullableErrorParametersItem) Get() *ErrorParametersItem {
	return v.value
}

func (v *NullableErrorParametersItem) Set(val *ErrorParametersItem) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorParametersItem) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorParametersItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorParametersItem(val *ErrorParametersItem) *NullableErrorParametersItem {
	return &NullableErrorParametersItem{value: val, isSet: true}
}

func (v NullableErrorParametersItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorParametersItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
