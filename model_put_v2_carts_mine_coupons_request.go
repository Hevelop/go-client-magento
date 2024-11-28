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

// checks if the PutV2CartsMineCouponsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV2CartsMineCouponsRequest{}

// PutV2CartsMineCouponsRequest struct for PutV2CartsMineCouponsRequest
type PutV2CartsMineCouponsRequest struct {
	CouponCodes          []string `json:"couponCodes"`
	AdditionalProperties map[string]interface{}
}

type _PutV2CartsMineCouponsRequest PutV2CartsMineCouponsRequest

// NewPutV2CartsMineCouponsRequest instantiates a new PutV2CartsMineCouponsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV2CartsMineCouponsRequest(couponCodes []string) *PutV2CartsMineCouponsRequest {
	this := PutV2CartsMineCouponsRequest{}
	this.CouponCodes = couponCodes
	return &this
}

// NewPutV2CartsMineCouponsRequestWithDefaults instantiates a new PutV2CartsMineCouponsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV2CartsMineCouponsRequestWithDefaults() *PutV2CartsMineCouponsRequest {
	this := PutV2CartsMineCouponsRequest{}
	return &this
}

// GetCouponCodes returns the CouponCodes field value
func (o *PutV2CartsMineCouponsRequest) GetCouponCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CouponCodes
}

// GetCouponCodesOk returns a tuple with the CouponCodes field value
// and a boolean to check if the value has been set.
func (o *PutV2CartsMineCouponsRequest) GetCouponCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CouponCodes, true
}

// SetCouponCodes sets field value
func (o *PutV2CartsMineCouponsRequest) SetCouponCodes(v []string) {
	o.CouponCodes = v
}

func (o PutV2CartsMineCouponsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV2CartsMineCouponsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["couponCodes"] = o.CouponCodes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV2CartsMineCouponsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"couponCodes",
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

	varPutV2CartsMineCouponsRequest := _PutV2CartsMineCouponsRequest{}

	err = json.Unmarshal(data, &varPutV2CartsMineCouponsRequest)

	if err != nil {
		return err
	}

	*o = PutV2CartsMineCouponsRequest(varPutV2CartsMineCouponsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "couponCodes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV2CartsMineCouponsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV2CartsMineCouponsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV2CartsMineCouponsRequest struct {
	value *PutV2CartsMineCouponsRequest
	isSet bool
}

func (v NullablePutV2CartsMineCouponsRequest) Get() *PutV2CartsMineCouponsRequest {
	return v.value
}

func (v *NullablePutV2CartsMineCouponsRequest) Set(val *PutV2CartsMineCouponsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV2CartsMineCouponsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV2CartsMineCouponsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV2CartsMineCouponsRequest(val *PutV2CartsMineCouponsRequest) *NullablePutV2CartsMineCouponsRequest {
	return &NullablePutV2CartsMineCouponsRequest{value: val, isSet: true}
}

func (v NullablePutV2CartsMineCouponsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV2CartsMineCouponsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
