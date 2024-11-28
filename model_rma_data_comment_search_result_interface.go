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

// checks if the RmaDataCommentSearchResultInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaDataCommentSearchResultInterface{}

// RmaDataCommentSearchResultInterface Interface CommentSearchResultInterface
type RmaDataCommentSearchResultInterface struct {
	// Rma Status History list
	Items          []RmaDataCommentInterface        `json:"items"`
	SearchCriteria FrameworkSearchCriteriaInterface `json:"search_criteria"`
	// Total count.
	TotalCount           int32 `json:"total_count"`
	AdditionalProperties map[string]interface{}
}

type _RmaDataCommentSearchResultInterface RmaDataCommentSearchResultInterface

// NewRmaDataCommentSearchResultInterface instantiates a new RmaDataCommentSearchResultInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaDataCommentSearchResultInterface(items []RmaDataCommentInterface, searchCriteria FrameworkSearchCriteriaInterface, totalCount int32) *RmaDataCommentSearchResultInterface {
	this := RmaDataCommentSearchResultInterface{}
	this.Items = items
	this.SearchCriteria = searchCriteria
	this.TotalCount = totalCount
	return &this
}

// NewRmaDataCommentSearchResultInterfaceWithDefaults instantiates a new RmaDataCommentSearchResultInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaDataCommentSearchResultInterfaceWithDefaults() *RmaDataCommentSearchResultInterface {
	this := RmaDataCommentSearchResultInterface{}
	return &this
}

// GetItems returns the Items field value
func (o *RmaDataCommentSearchResultInterface) GetItems() []RmaDataCommentInterface {
	if o == nil {
		var ret []RmaDataCommentInterface
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RmaDataCommentSearchResultInterface) GetItemsOk() ([]RmaDataCommentInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RmaDataCommentSearchResultInterface) SetItems(v []RmaDataCommentInterface) {
	o.Items = v
}

// GetSearchCriteria returns the SearchCriteria field value
func (o *RmaDataCommentSearchResultInterface) GetSearchCriteria() FrameworkSearchCriteriaInterface {
	if o == nil {
		var ret FrameworkSearchCriteriaInterface
		return ret
	}

	return o.SearchCriteria
}

// GetSearchCriteriaOk returns a tuple with the SearchCriteria field value
// and a boolean to check if the value has been set.
func (o *RmaDataCommentSearchResultInterface) GetSearchCriteriaOk() (*FrameworkSearchCriteriaInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchCriteria, true
}

// SetSearchCriteria sets field value
func (o *RmaDataCommentSearchResultInterface) SetSearchCriteria(v FrameworkSearchCriteriaInterface) {
	o.SearchCriteria = v
}

// GetTotalCount returns the TotalCount field value
func (o *RmaDataCommentSearchResultInterface) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *RmaDataCommentSearchResultInterface) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *RmaDataCommentSearchResultInterface) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o RmaDataCommentSearchResultInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaDataCommentSearchResultInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["search_criteria"] = o.SearchCriteria
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaDataCommentSearchResultInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varRmaDataCommentSearchResultInterface := _RmaDataCommentSearchResultInterface{}

	err = json.Unmarshal(data, &varRmaDataCommentSearchResultInterface)

	if err != nil {
		return err
	}

	*o = RmaDataCommentSearchResultInterface(varRmaDataCommentSearchResultInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "search_criteria")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaDataCommentSearchResultInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaDataCommentSearchResultInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaDataCommentSearchResultInterface struct {
	value *RmaDataCommentSearchResultInterface
	isSet bool
}

func (v NullableRmaDataCommentSearchResultInterface) Get() *RmaDataCommentSearchResultInterface {
	return v.value
}

func (v *NullableRmaDataCommentSearchResultInterface) Set(val *RmaDataCommentSearchResultInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaDataCommentSearchResultInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaDataCommentSearchResultInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaDataCommentSearchResultInterface(val *RmaDataCommentSearchResultInterface) *NullableRmaDataCommentSearchResultInterface {
	return &NullableRmaDataCommentSearchResultInterface{value: val, isSet: true}
}

func (v NullableRmaDataCommentSearchResultInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaDataCommentSearchResultInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}