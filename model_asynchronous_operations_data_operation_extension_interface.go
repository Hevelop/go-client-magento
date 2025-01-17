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

// checks if the AsynchronousOperationsDataOperationExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsynchronousOperationsDataOperationExtensionInterface{}

// AsynchronousOperationsDataOperationExtensionInterface ExtensionInterface class for @see \\Magento\\AsynchronousOperations\\Api\\Data\\OperationInterface
type AsynchronousOperationsDataOperationExtensionInterface struct {
	StartTime            *string `json:"start_time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AsynchronousOperationsDataOperationExtensionInterface AsynchronousOperationsDataOperationExtensionInterface

// NewAsynchronousOperationsDataOperationExtensionInterface instantiates a new AsynchronousOperationsDataOperationExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsynchronousOperationsDataOperationExtensionInterface() *AsynchronousOperationsDataOperationExtensionInterface {
	this := AsynchronousOperationsDataOperationExtensionInterface{}
	return &this
}

// NewAsynchronousOperationsDataOperationExtensionInterfaceWithDefaults instantiates a new AsynchronousOperationsDataOperationExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsynchronousOperationsDataOperationExtensionInterfaceWithDefaults() *AsynchronousOperationsDataOperationExtensionInterface {
	this := AsynchronousOperationsDataOperationExtensionInterface{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AsynchronousOperationsDataOperationExtensionInterface) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationExtensionInterface) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AsynchronousOperationsDataOperationExtensionInterface) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *AsynchronousOperationsDataOperationExtensionInterface) SetStartTime(v string) {
	o.StartTime = &v
}

func (o AsynchronousOperationsDataOperationExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsynchronousOperationsDataOperationExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AsynchronousOperationsDataOperationExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varAsynchronousOperationsDataOperationExtensionInterface := _AsynchronousOperationsDataOperationExtensionInterface{}

	err = json.Unmarshal(data, &varAsynchronousOperationsDataOperationExtensionInterface)

	if err != nil {
		return err
	}

	*o = AsynchronousOperationsDataOperationExtensionInterface(varAsynchronousOperationsDataOperationExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "start_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AsynchronousOperationsDataOperationExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AsynchronousOperationsDataOperationExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAsynchronousOperationsDataOperationExtensionInterface struct {
	value *AsynchronousOperationsDataOperationExtensionInterface
	isSet bool
}

func (v NullableAsynchronousOperationsDataOperationExtensionInterface) Get() *AsynchronousOperationsDataOperationExtensionInterface {
	return v.value
}

func (v *NullableAsynchronousOperationsDataOperationExtensionInterface) Set(val *AsynchronousOperationsDataOperationExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableAsynchronousOperationsDataOperationExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableAsynchronousOperationsDataOperationExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsynchronousOperationsDataOperationExtensionInterface(val *AsynchronousOperationsDataOperationExtensionInterface) *NullableAsynchronousOperationsDataOperationExtensionInterface {
	return &NullableAsynchronousOperationsDataOperationExtensionInterface{value: val, isSet: true}
}

func (v NullableAsynchronousOperationsDataOperationExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsynchronousOperationsDataOperationExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
