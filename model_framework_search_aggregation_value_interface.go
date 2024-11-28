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

// checks if the FrameworkSearchAggregationValueInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameworkSearchAggregationValueInterface{}

// FrameworkSearchAggregationValueInterface Interface Aggregation Value
type FrameworkSearchAggregationValueInterface struct {
	// Aggregation
	Value string `json:"value"`
	// Metrics
	Metrics              []string `json:"metrics"`
	AdditionalProperties map[string]interface{}
}

type _FrameworkSearchAggregationValueInterface FrameworkSearchAggregationValueInterface

// NewFrameworkSearchAggregationValueInterface instantiates a new FrameworkSearchAggregationValueInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameworkSearchAggregationValueInterface(value string, metrics []string) *FrameworkSearchAggregationValueInterface {
	this := FrameworkSearchAggregationValueInterface{}
	this.Value = value
	this.Metrics = metrics
	return &this
}

// NewFrameworkSearchAggregationValueInterfaceWithDefaults instantiates a new FrameworkSearchAggregationValueInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkSearchAggregationValueInterfaceWithDefaults() *FrameworkSearchAggregationValueInterface {
	this := FrameworkSearchAggregationValueInterface{}
	return &this
}

// GetValue returns the Value field value
func (o *FrameworkSearchAggregationValueInterface) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchAggregationValueInterface) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FrameworkSearchAggregationValueInterface) SetValue(v string) {
	o.Value = v
}

// GetMetrics returns the Metrics field value
func (o *FrameworkSearchAggregationValueInterface) GetMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchAggregationValueInterface) GetMetricsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *FrameworkSearchAggregationValueInterface) SetMetrics(v []string) {
	o.Metrics = v
}

func (o FrameworkSearchAggregationValueInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameworkSearchAggregationValueInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["metrics"] = o.Metrics

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrameworkSearchAggregationValueInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"metrics",
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

	varFrameworkSearchAggregationValueInterface := _FrameworkSearchAggregationValueInterface{}

	err = json.Unmarshal(data, &varFrameworkSearchAggregationValueInterface)

	if err != nil {
		return err
	}

	*o = FrameworkSearchAggregationValueInterface(varFrameworkSearchAggregationValueInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "metrics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *FrameworkSearchAggregationValueInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *FrameworkSearchAggregationValueInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableFrameworkSearchAggregationValueInterface struct {
	value *FrameworkSearchAggregationValueInterface
	isSet bool
}

func (v NullableFrameworkSearchAggregationValueInterface) Get() *FrameworkSearchAggregationValueInterface {
	return v.value
}

func (v *NullableFrameworkSearchAggregationValueInterface) Set(val *FrameworkSearchAggregationValueInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameworkSearchAggregationValueInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameworkSearchAggregationValueInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameworkSearchAggregationValueInterface(val *FrameworkSearchAggregationValueInterface) *NullableFrameworkSearchAggregationValueInterface {
	return &NullableFrameworkSearchAggregationValueInterface{value: val, isSet: true}
}

func (v NullableFrameworkSearchAggregationValueInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameworkSearchAggregationValueInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}