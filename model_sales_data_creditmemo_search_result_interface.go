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

// checks if the SalesDataCreditmemoSearchResultInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataCreditmemoSearchResultInterface{}

// SalesDataCreditmemoSearchResultInterface Credit memo search result interface. After a customer places and pays for an order and an invoice has been issued, the merchant can create a credit memo to refund all or part of the amount paid for any returned or undelivered items. The memo restores funds to the customer account so that the customer can make future purchases.
type SalesDataCreditmemoSearchResultInterface struct {
	// Array of collection items.
	Items          []SalesDataCreditmemoInterface   `json:"items"`
	SearchCriteria FrameworkSearchCriteriaInterface `json:"search_criteria"`
	// Total count.
	TotalCount           int32 `json:"total_count"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataCreditmemoSearchResultInterface SalesDataCreditmemoSearchResultInterface

// NewSalesDataCreditmemoSearchResultInterface instantiates a new SalesDataCreditmemoSearchResultInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataCreditmemoSearchResultInterface(items []SalesDataCreditmemoInterface, searchCriteria FrameworkSearchCriteriaInterface, totalCount int32) *SalesDataCreditmemoSearchResultInterface {
	this := SalesDataCreditmemoSearchResultInterface{}
	this.Items = items
	this.SearchCriteria = searchCriteria
	this.TotalCount = totalCount
	return &this
}

// NewSalesDataCreditmemoSearchResultInterfaceWithDefaults instantiates a new SalesDataCreditmemoSearchResultInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataCreditmemoSearchResultInterfaceWithDefaults() *SalesDataCreditmemoSearchResultInterface {
	this := SalesDataCreditmemoSearchResultInterface{}
	return &this
}

// GetItems returns the Items field value
func (o *SalesDataCreditmemoSearchResultInterface) GetItems() []SalesDataCreditmemoInterface {
	if o == nil {
		var ret []SalesDataCreditmemoInterface
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoSearchResultInterface) GetItemsOk() ([]SalesDataCreditmemoInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SalesDataCreditmemoSearchResultInterface) SetItems(v []SalesDataCreditmemoInterface) {
	o.Items = v
}

// GetSearchCriteria returns the SearchCriteria field value
func (o *SalesDataCreditmemoSearchResultInterface) GetSearchCriteria() FrameworkSearchCriteriaInterface {
	if o == nil {
		var ret FrameworkSearchCriteriaInterface
		return ret
	}

	return o.SearchCriteria
}

// GetSearchCriteriaOk returns a tuple with the SearchCriteria field value
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoSearchResultInterface) GetSearchCriteriaOk() (*FrameworkSearchCriteriaInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchCriteria, true
}

// SetSearchCriteria sets field value
func (o *SalesDataCreditmemoSearchResultInterface) SetSearchCriteria(v FrameworkSearchCriteriaInterface) {
	o.SearchCriteria = v
}

// GetTotalCount returns the TotalCount field value
func (o *SalesDataCreditmemoSearchResultInterface) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoSearchResultInterface) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SalesDataCreditmemoSearchResultInterface) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o SalesDataCreditmemoSearchResultInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataCreditmemoSearchResultInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["search_criteria"] = o.SearchCriteria
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataCreditmemoSearchResultInterface) UnmarshalJSON(data []byte) (err error) {
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

	varSalesDataCreditmemoSearchResultInterface := _SalesDataCreditmemoSearchResultInterface{}

	err = json.Unmarshal(data, &varSalesDataCreditmemoSearchResultInterface)

	if err != nil {
		return err
	}

	*o = SalesDataCreditmemoSearchResultInterface(varSalesDataCreditmemoSearchResultInterface)

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
func (o *SalesDataCreditmemoSearchResultInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataCreditmemoSearchResultInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataCreditmemoSearchResultInterface struct {
	value *SalesDataCreditmemoSearchResultInterface
	isSet bool
}

func (v NullableSalesDataCreditmemoSearchResultInterface) Get() *SalesDataCreditmemoSearchResultInterface {
	return v.value
}

func (v *NullableSalesDataCreditmemoSearchResultInterface) Set(val *SalesDataCreditmemoSearchResultInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataCreditmemoSearchResultInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataCreditmemoSearchResultInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataCreditmemoSearchResultInterface(val *SalesDataCreditmemoSearchResultInterface) *NullableSalesDataCreditmemoSearchResultInterface {
	return &NullableSalesDataCreditmemoSearchResultInterface{value: val, isSet: true}
}

func (v NullableSalesDataCreditmemoSearchResultInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataCreditmemoSearchResultInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}