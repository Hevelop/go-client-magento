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

// checks if the PutV1ProductsTierpricesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1ProductsTierpricesRequest{}

// PutV1ProductsTierpricesRequest struct for PutV1ProductsTierpricesRequest
type PutV1ProductsTierpricesRequest struct {
	Prices               []CatalogDataTierPriceInterface `json:"prices"`
	AdditionalProperties map[string]interface{}
}

type _PutV1ProductsTierpricesRequest PutV1ProductsTierpricesRequest

// NewPutV1ProductsTierpricesRequest instantiates a new PutV1ProductsTierpricesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1ProductsTierpricesRequest(prices []CatalogDataTierPriceInterface) *PutV1ProductsTierpricesRequest {
	this := PutV1ProductsTierpricesRequest{}
	this.Prices = prices
	return &this
}

// NewPutV1ProductsTierpricesRequestWithDefaults instantiates a new PutV1ProductsTierpricesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1ProductsTierpricesRequestWithDefaults() *PutV1ProductsTierpricesRequest {
	this := PutV1ProductsTierpricesRequest{}
	return &this
}

// GetPrices returns the Prices field value
func (o *PutV1ProductsTierpricesRequest) GetPrices() []CatalogDataTierPriceInterface {
	if o == nil {
		var ret []CatalogDataTierPriceInterface
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *PutV1ProductsTierpricesRequest) GetPricesOk() ([]CatalogDataTierPriceInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *PutV1ProductsTierpricesRequest) SetPrices(v []CatalogDataTierPriceInterface) {
	o.Prices = v
}

func (o PutV1ProductsTierpricesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1ProductsTierpricesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prices"] = o.Prices

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1ProductsTierpricesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPutV1ProductsTierpricesRequest := _PutV1ProductsTierpricesRequest{}

	err = json.Unmarshal(data, &varPutV1ProductsTierpricesRequest)

	if err != nil {
		return err
	}

	*o = PutV1ProductsTierpricesRequest(varPutV1ProductsTierpricesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "prices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1ProductsTierpricesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1ProductsTierpricesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1ProductsTierpricesRequest struct {
	value *PutV1ProductsTierpricesRequest
	isSet bool
}

func (v NullablePutV1ProductsTierpricesRequest) Get() *PutV1ProductsTierpricesRequest {
	return v.value
}

func (v *NullablePutV1ProductsTierpricesRequest) Set(val *PutV1ProductsTierpricesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1ProductsTierpricesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1ProductsTierpricesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1ProductsTierpricesRequest(val *PutV1ProductsTierpricesRequest) *NullablePutV1ProductsTierpricesRequest {
	return &NullablePutV1ProductsTierpricesRequest{value: val, isSet: true}
}

func (v NullablePutV1ProductsTierpricesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1ProductsTierpricesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
