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

// checks if the PostV1ProductsCostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1ProductsCostRequest{}

// PostV1ProductsCostRequest struct for PostV1ProductsCostRequest
type PostV1ProductsCostRequest struct {
	Prices               []CatalogDataCostInterface `json:"prices"`
	AdditionalProperties map[string]interface{}
}

type _PostV1ProductsCostRequest PostV1ProductsCostRequest

// NewPostV1ProductsCostRequest instantiates a new PostV1ProductsCostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1ProductsCostRequest(prices []CatalogDataCostInterface) *PostV1ProductsCostRequest {
	this := PostV1ProductsCostRequest{}
	this.Prices = prices
	return &this
}

// NewPostV1ProductsCostRequestWithDefaults instantiates a new PostV1ProductsCostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1ProductsCostRequestWithDefaults() *PostV1ProductsCostRequest {
	this := PostV1ProductsCostRequest{}
	return &this
}

// GetPrices returns the Prices field value
func (o *PostV1ProductsCostRequest) GetPrices() []CatalogDataCostInterface {
	if o == nil {
		var ret []CatalogDataCostInterface
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *PostV1ProductsCostRequest) GetPricesOk() ([]CatalogDataCostInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *PostV1ProductsCostRequest) SetPrices(v []CatalogDataCostInterface) {
	o.Prices = v
}

func (o PostV1ProductsCostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1ProductsCostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prices"] = o.Prices

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1ProductsCostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prices",
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

	varPostV1ProductsCostRequest := _PostV1ProductsCostRequest{}

	err = json.Unmarshal(data, &varPostV1ProductsCostRequest)

	if err != nil {
		return err
	}

	*o = PostV1ProductsCostRequest(varPostV1ProductsCostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "prices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1ProductsCostRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1ProductsCostRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1ProductsCostRequest struct {
	value *PostV1ProductsCostRequest
	isSet bool
}

func (v NullablePostV1ProductsCostRequest) Get() *PostV1ProductsCostRequest {
	return v.value
}

func (v *NullablePostV1ProductsCostRequest) Set(val *PostV1ProductsCostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1ProductsCostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1ProductsCostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1ProductsCostRequest(val *PostV1ProductsCostRequest) *NullablePostV1ProductsCostRequest {
	return &NullablePostV1ProductsCostRequest{value: val, isSet: true}
}

func (v NullablePostV1ProductsCostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1ProductsCostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}