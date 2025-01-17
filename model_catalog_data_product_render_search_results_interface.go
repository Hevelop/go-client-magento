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

// checks if the CatalogDataProductRenderSearchResultsInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDataProductRenderSearchResultsInterface{}

// CatalogDataProductRenderSearchResultsInterface Dto that holds render information about products
type CatalogDataProductRenderSearchResultsInterface struct {
	// List of products rendered information
	Items                []CatalogDataProductRenderInterface `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _CatalogDataProductRenderSearchResultsInterface CatalogDataProductRenderSearchResultsInterface

// NewCatalogDataProductRenderSearchResultsInterface instantiates a new CatalogDataProductRenderSearchResultsInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDataProductRenderSearchResultsInterface(items []CatalogDataProductRenderInterface) *CatalogDataProductRenderSearchResultsInterface {
	this := CatalogDataProductRenderSearchResultsInterface{}
	this.Items = items
	return &this
}

// NewCatalogDataProductRenderSearchResultsInterfaceWithDefaults instantiates a new CatalogDataProductRenderSearchResultsInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDataProductRenderSearchResultsInterfaceWithDefaults() *CatalogDataProductRenderSearchResultsInterface {
	this := CatalogDataProductRenderSearchResultsInterface{}
	return &this
}

// GetItems returns the Items field value
func (o *CatalogDataProductRenderSearchResultsInterface) GetItems() []CatalogDataProductRenderInterface {
	if o == nil {
		var ret []CatalogDataProductRenderInterface
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderSearchResultsInterface) GetItemsOk() ([]CatalogDataProductRenderInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CatalogDataProductRenderSearchResultsInterface) SetItems(v []CatalogDataProductRenderInterface) {
	o.Items = v
}

func (o CatalogDataProductRenderSearchResultsInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDataProductRenderSearchResultsInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogDataProductRenderSearchResultsInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varCatalogDataProductRenderSearchResultsInterface := _CatalogDataProductRenderSearchResultsInterface{}

	err = json.Unmarshal(data, &varCatalogDataProductRenderSearchResultsInterface)

	if err != nil {
		return err
	}

	*o = CatalogDataProductRenderSearchResultsInterface(varCatalogDataProductRenderSearchResultsInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogDataProductRenderSearchResultsInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogDataProductRenderSearchResultsInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogDataProductRenderSearchResultsInterface struct {
	value *CatalogDataProductRenderSearchResultsInterface
	isSet bool
}

func (v NullableCatalogDataProductRenderSearchResultsInterface) Get() *CatalogDataProductRenderSearchResultsInterface {
	return v.value
}

func (v *NullableCatalogDataProductRenderSearchResultsInterface) Set(val *CatalogDataProductRenderSearchResultsInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDataProductRenderSearchResultsInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDataProductRenderSearchResultsInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDataProductRenderSearchResultsInterface(val *CatalogDataProductRenderSearchResultsInterface) *NullableCatalogDataProductRenderSearchResultsInterface {
	return &NullableCatalogDataProductRenderSearchResultsInterface{value: val, isSet: true}
}

func (v NullableCatalogDataProductRenderSearchResultsInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDataProductRenderSearchResultsInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
