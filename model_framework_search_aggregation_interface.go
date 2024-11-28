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

// checks if the FrameworkSearchAggregationInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameworkSearchAggregationInterface{}

// FrameworkSearchAggregationInterface Interface Aggregation to get faceted data
type FrameworkSearchAggregationInterface struct {
	// All Document fields
	Buckets []FrameworkSearchBucketInterface `json:"buckets"`
	// Document field names
	BucketNames          []string `json:"bucket_names"`
	AdditionalProperties map[string]interface{}
}

type _FrameworkSearchAggregationInterface FrameworkSearchAggregationInterface

// NewFrameworkSearchAggregationInterface instantiates a new FrameworkSearchAggregationInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameworkSearchAggregationInterface(buckets []FrameworkSearchBucketInterface, bucketNames []string) *FrameworkSearchAggregationInterface {
	this := FrameworkSearchAggregationInterface{}
	this.Buckets = buckets
	this.BucketNames = bucketNames
	return &this
}

// NewFrameworkSearchAggregationInterfaceWithDefaults instantiates a new FrameworkSearchAggregationInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkSearchAggregationInterfaceWithDefaults() *FrameworkSearchAggregationInterface {
	this := FrameworkSearchAggregationInterface{}
	return &this
}

// GetBuckets returns the Buckets field value
func (o *FrameworkSearchAggregationInterface) GetBuckets() []FrameworkSearchBucketInterface {
	if o == nil {
		var ret []FrameworkSearchBucketInterface
		return ret
	}

	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchAggregationInterface) GetBucketsOk() ([]FrameworkSearchBucketInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buckets, true
}

// SetBuckets sets field value
func (o *FrameworkSearchAggregationInterface) SetBuckets(v []FrameworkSearchBucketInterface) {
	o.Buckets = v
}

// GetBucketNames returns the BucketNames field value
func (o *FrameworkSearchAggregationInterface) GetBucketNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BucketNames
}

// GetBucketNamesOk returns a tuple with the BucketNames field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchAggregationInterface) GetBucketNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BucketNames, true
}

// SetBucketNames sets field value
func (o *FrameworkSearchAggregationInterface) SetBucketNames(v []string) {
	o.BucketNames = v
}

func (o FrameworkSearchAggregationInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameworkSearchAggregationInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buckets"] = o.Buckets
	toSerialize["bucket_names"] = o.BucketNames

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrameworkSearchAggregationInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buckets",
		"bucket_names",
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

	varFrameworkSearchAggregationInterface := _FrameworkSearchAggregationInterface{}

	err = json.Unmarshal(data, &varFrameworkSearchAggregationInterface)

	if err != nil {
		return err
	}

	*o = FrameworkSearchAggregationInterface(varFrameworkSearchAggregationInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "buckets")
		delete(additionalProperties, "bucket_names")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *FrameworkSearchAggregationInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *FrameworkSearchAggregationInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableFrameworkSearchAggregationInterface struct {
	value *FrameworkSearchAggregationInterface
	isSet bool
}

func (v NullableFrameworkSearchAggregationInterface) Get() *FrameworkSearchAggregationInterface {
	return v.value
}

func (v *NullableFrameworkSearchAggregationInterface) Set(val *FrameworkSearchAggregationInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameworkSearchAggregationInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameworkSearchAggregationInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameworkSearchAggregationInterface(val *FrameworkSearchAggregationInterface) *NullableFrameworkSearchAggregationInterface {
	return &NullableFrameworkSearchAggregationInterface{value: val, isSet: true}
}

func (v NullableFrameworkSearchAggregationInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameworkSearchAggregationInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
