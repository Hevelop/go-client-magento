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

// checks if the FrameworkSearchSearchResultInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameworkSearchSearchResultInterface{}

// FrameworkSearchSearchResultInterface Interface SearchResultInterface
type FrameworkSearchSearchResultInterface struct {
	Items          []FrameworkSearchDocumentInterface     `json:"items"`
	Aggregations   FrameworkSearchAggregationInterface    `json:"aggregations"`
	SearchCriteria FrameworkSearchSearchCriteriaInterface `json:"search_criteria"`
	// Total count.
	TotalCount           int32 `json:"total_count"`
	AdditionalProperties map[string]interface{}
}

type _FrameworkSearchSearchResultInterface FrameworkSearchSearchResultInterface

// NewFrameworkSearchSearchResultInterface instantiates a new FrameworkSearchSearchResultInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameworkSearchSearchResultInterface(items []FrameworkSearchDocumentInterface, aggregations FrameworkSearchAggregationInterface, searchCriteria FrameworkSearchSearchCriteriaInterface, totalCount int32) *FrameworkSearchSearchResultInterface {
	this := FrameworkSearchSearchResultInterface{}
	this.Items = items
	this.Aggregations = aggregations
	this.SearchCriteria = searchCriteria
	this.TotalCount = totalCount
	return &this
}

// NewFrameworkSearchSearchResultInterfaceWithDefaults instantiates a new FrameworkSearchSearchResultInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkSearchSearchResultInterfaceWithDefaults() *FrameworkSearchSearchResultInterface {
	this := FrameworkSearchSearchResultInterface{}
	return &this
}

// GetItems returns the Items field value
func (o *FrameworkSearchSearchResultInterface) GetItems() []FrameworkSearchDocumentInterface {
	if o == nil {
		var ret []FrameworkSearchDocumentInterface
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchSearchResultInterface) GetItemsOk() ([]FrameworkSearchDocumentInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *FrameworkSearchSearchResultInterface) SetItems(v []FrameworkSearchDocumentInterface) {
	o.Items = v
}

// GetAggregations returns the Aggregations field value
func (o *FrameworkSearchSearchResultInterface) GetAggregations() FrameworkSearchAggregationInterface {
	if o == nil {
		var ret FrameworkSearchAggregationInterface
		return ret
	}

	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchSearchResultInterface) GetAggregationsOk() (*FrameworkSearchAggregationInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregations, true
}

// SetAggregations sets field value
func (o *FrameworkSearchSearchResultInterface) SetAggregations(v FrameworkSearchAggregationInterface) {
	o.Aggregations = v
}

// GetSearchCriteria returns the SearchCriteria field value
func (o *FrameworkSearchSearchResultInterface) GetSearchCriteria() FrameworkSearchSearchCriteriaInterface {
	if o == nil {
		var ret FrameworkSearchSearchCriteriaInterface
		return ret
	}

	return o.SearchCriteria
}

// GetSearchCriteriaOk returns a tuple with the SearchCriteria field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchSearchResultInterface) GetSearchCriteriaOk() (*FrameworkSearchSearchCriteriaInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchCriteria, true
}

// SetSearchCriteria sets field value
func (o *FrameworkSearchSearchResultInterface) SetSearchCriteria(v FrameworkSearchSearchCriteriaInterface) {
	o.SearchCriteria = v
}

// GetTotalCount returns the TotalCount field value
func (o *FrameworkSearchSearchResultInterface) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *FrameworkSearchSearchResultInterface) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *FrameworkSearchSearchResultInterface) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o FrameworkSearchSearchResultInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameworkSearchSearchResultInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["aggregations"] = o.Aggregations
	toSerialize["search_criteria"] = o.SearchCriteria
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrameworkSearchSearchResultInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"aggregations",
		"search_criteria",
		"total_count",
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

	varFrameworkSearchSearchResultInterface := _FrameworkSearchSearchResultInterface{}

	err = json.Unmarshal(data, &varFrameworkSearchSearchResultInterface)

	if err != nil {
		return err
	}

	*o = FrameworkSearchSearchResultInterface(varFrameworkSearchSearchResultInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "search_criteria")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *FrameworkSearchSearchResultInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *FrameworkSearchSearchResultInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableFrameworkSearchSearchResultInterface struct {
	value *FrameworkSearchSearchResultInterface
	isSet bool
}

func (v NullableFrameworkSearchSearchResultInterface) Get() *FrameworkSearchSearchResultInterface {
	return v.value
}

func (v *NullableFrameworkSearchSearchResultInterface) Set(val *FrameworkSearchSearchResultInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameworkSearchSearchResultInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameworkSearchSearchResultInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameworkSearchSearchResultInterface(val *FrameworkSearchSearchResultInterface) *NullableFrameworkSearchSearchResultInterface {
	return &NullableFrameworkSearchSearchResultInterface{value: val, isSet: true}
}

func (v NullableFrameworkSearchSearchResultInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameworkSearchSearchResultInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
