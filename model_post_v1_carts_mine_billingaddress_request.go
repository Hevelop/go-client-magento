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

// checks if the PostV1CartsMineBillingaddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1CartsMineBillingaddressRequest{}

// PostV1CartsMineBillingaddressRequest struct for PostV1CartsMineBillingaddressRequest
type PostV1CartsMineBillingaddressRequest struct {
	Address              QuoteDataAddressInterface `json:"address"`
	UseForShipping       *bool                     `json:"useForShipping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostV1CartsMineBillingaddressRequest PostV1CartsMineBillingaddressRequest

// NewPostV1CartsMineBillingaddressRequest instantiates a new PostV1CartsMineBillingaddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1CartsMineBillingaddressRequest(address QuoteDataAddressInterface) *PostV1CartsMineBillingaddressRequest {
	this := PostV1CartsMineBillingaddressRequest{}
	this.Address = address
	return &this
}

// NewPostV1CartsMineBillingaddressRequestWithDefaults instantiates a new PostV1CartsMineBillingaddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1CartsMineBillingaddressRequestWithDefaults() *PostV1CartsMineBillingaddressRequest {
	this := PostV1CartsMineBillingaddressRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *PostV1CartsMineBillingaddressRequest) GetAddress() QuoteDataAddressInterface {
	if o == nil {
		var ret QuoteDataAddressInterface
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PostV1CartsMineBillingaddressRequest) GetAddressOk() (*QuoteDataAddressInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PostV1CartsMineBillingaddressRequest) SetAddress(v QuoteDataAddressInterface) {
	o.Address = v
}

// GetUseForShipping returns the UseForShipping field value if set, zero value otherwise.
func (o *PostV1CartsMineBillingaddressRequest) GetUseForShipping() bool {
	if o == nil || IsNil(o.UseForShipping) {
		var ret bool
		return ret
	}
	return *o.UseForShipping
}

// GetUseForShippingOk returns a tuple with the UseForShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1CartsMineBillingaddressRequest) GetUseForShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseForShipping) {
		return nil, false
	}
	return o.UseForShipping, true
}

// HasUseForShipping returns a boolean if a field has been set.
func (o *PostV1CartsMineBillingaddressRequest) HasUseForShipping() bool {
	if o != nil && !IsNil(o.UseForShipping) {
		return true
	}

	return false
}

// SetUseForShipping gets a reference to the given bool and assigns it to the UseForShipping field.
func (o *PostV1CartsMineBillingaddressRequest) SetUseForShipping(v bool) {
	o.UseForShipping = &v
}

func (o PostV1CartsMineBillingaddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1CartsMineBillingaddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.UseForShipping) {
		toSerialize["useForShipping"] = o.UseForShipping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1CartsMineBillingaddressRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
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

	varPostV1CartsMineBillingaddressRequest := _PostV1CartsMineBillingaddressRequest{}

	err = json.Unmarshal(data, &varPostV1CartsMineBillingaddressRequest)

	if err != nil {
		return err
	}

	*o = PostV1CartsMineBillingaddressRequest(varPostV1CartsMineBillingaddressRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "useForShipping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1CartsMineBillingaddressRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1CartsMineBillingaddressRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1CartsMineBillingaddressRequest struct {
	value *PostV1CartsMineBillingaddressRequest
	isSet bool
}

func (v NullablePostV1CartsMineBillingaddressRequest) Get() *PostV1CartsMineBillingaddressRequest {
	return v.value
}

func (v *NullablePostV1CartsMineBillingaddressRequest) Set(val *PostV1CartsMineBillingaddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1CartsMineBillingaddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1CartsMineBillingaddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1CartsMineBillingaddressRequest(val *PostV1CartsMineBillingaddressRequest) *NullablePostV1CartsMineBillingaddressRequest {
	return &NullablePostV1CartsMineBillingaddressRequest{value: val, isSet: true}
}

func (v NullablePostV1CartsMineBillingaddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1CartsMineBillingaddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
