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

// checks if the FrameworkSearchBucketInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameworkSearchBucketInterface{}

// FrameworkSearchBucketInterface Interface for facet Bucket
type FrameworkSearchBucketInterface struct {
	// Field name
	Name string `json:"name"`
	// Field values
	Values               []FrameworkSearchAggregationValueInterface `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _FrameworkSearchBucketInterface FrameworkSearchBucketInterface

// NewFrameworkSearchBucketInterface instantiates a new FrameworkSearchBucketInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameworkSearchBucketInterface(name string, values []FrameworkSearchAggregationValueInterface) *FrameworkSearchBucketInterface {
	this := FrameworkSearchBucketInterface{}
	this.Name = name
	this.Values = values
	return &this
}

// NewFrameworkSearchBucketInterfaceWithDefaults instantiates a new FrameworkSearchBucketInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkSearchBucketInterfaceWithDefaults() *FrameworkSearchBucketInterface {
	this := FrameworkSearchBucketInterface{}
	return &this
}

// GetName returns the Name field value
func (o *FrameworkSearchBucketInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchBucketInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrameworkSearchBucketInterface) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *FrameworkSearchBucketInterface) GetValues() []FrameworkSearchAggregationValueInterface {
	if o == nil {
		var ret []FrameworkSearchAggregationValueInterface
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchBucketInterface) GetValuesOk() ([]FrameworkSearchAggregationValueInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *FrameworkSearchBucketInterface) SetValues(v []FrameworkSearchAggregationValueInterface) {
	o.Values = v
}

func (o FrameworkSearchBucketInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameworkSearchBucketInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrameworkSearchBucketInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"values",
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

	varFrameworkSearchBucketInterface := _FrameworkSearchBucketInterface{}

	err = json.Unmarshal(data, &varFrameworkSearchBucketInterface)

	if err != nil {
		return err
	}

	*o = FrameworkSearchBucketInterface(varFrameworkSearchBucketInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *FrameworkSearchBucketInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *FrameworkSearchBucketInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableFrameworkSearchBucketInterface struct {
	value *FrameworkSearchBucketInterface
	isSet bool
}

func (v NullableFrameworkSearchBucketInterface) Get() *FrameworkSearchBucketInterface {
	return v.value
}

func (v *NullableFrameworkSearchBucketInterface) Set(val *FrameworkSearchBucketInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameworkSearchBucketInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameworkSearchBucketInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameworkSearchBucketInterface(val *FrameworkSearchBucketInterface) *NullableFrameworkSearchBucketInterface {
	return &NullableFrameworkSearchBucketInterface{value: val, isSet: true}
}

func (v NullableFrameworkSearchBucketInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameworkSearchBucketInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
