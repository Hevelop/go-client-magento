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

// checks if the PutV1BundleproductsSkuLinksIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1BundleproductsSkuLinksIdRequest{}

// PutV1BundleproductsSkuLinksIdRequest struct for PutV1BundleproductsSkuLinksIdRequest
type PutV1BundleproductsSkuLinksIdRequest struct {
	LinkedProduct        BundleDataLinkInterface `json:"linkedProduct"`
	AdditionalProperties map[string]interface{}
}

type _PutV1BundleproductsSkuLinksIdRequest PutV1BundleproductsSkuLinksIdRequest

// NewPutV1BundleproductsSkuLinksIdRequest instantiates a new PutV1BundleproductsSkuLinksIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1BundleproductsSkuLinksIdRequest(linkedProduct BundleDataLinkInterface) *PutV1BundleproductsSkuLinksIdRequest {
	this := PutV1BundleproductsSkuLinksIdRequest{}
	this.LinkedProduct = linkedProduct
	return &this
}

// NewPutV1BundleproductsSkuLinksIdRequestWithDefaults instantiates a new PutV1BundleproductsSkuLinksIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1BundleproductsSkuLinksIdRequestWithDefaults() *PutV1BundleproductsSkuLinksIdRequest {
	this := PutV1BundleproductsSkuLinksIdRequest{}
	return &this
}

// GetLinkedProduct returns the LinkedProduct field value
func (o *PutV1BundleproductsSkuLinksIdRequest) GetLinkedProduct() BundleDataLinkInterface {
	if o == nil {
		var ret BundleDataLinkInterface
		return ret
	}

	return o.LinkedProduct
}

// GetLinkedProductOk returns a tuple with the LinkedProduct field value
// and a boolean to check if the value has been set.
func (o *PutV1BundleproductsSkuLinksIdRequest) GetLinkedProductOk() (*BundleDataLinkInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkedProduct, true
}

// SetLinkedProduct sets field value
func (o *PutV1BundleproductsSkuLinksIdRequest) SetLinkedProduct(v BundleDataLinkInterface) {
	o.LinkedProduct = v
}

func (o PutV1BundleproductsSkuLinksIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1BundleproductsSkuLinksIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["linkedProduct"] = o.LinkedProduct

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1BundleproductsSkuLinksIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"linkedProduct",
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

	varPutV1BundleproductsSkuLinksIdRequest := _PutV1BundleproductsSkuLinksIdRequest{}

	err = json.Unmarshal(data, &varPutV1BundleproductsSkuLinksIdRequest)

	if err != nil {
		return err
	}

	*o = PutV1BundleproductsSkuLinksIdRequest(varPutV1BundleproductsSkuLinksIdRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "linkedProduct")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1BundleproductsSkuLinksIdRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1BundleproductsSkuLinksIdRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1BundleproductsSkuLinksIdRequest struct {
	value *PutV1BundleproductsSkuLinksIdRequest
	isSet bool
}

func (v NullablePutV1BundleproductsSkuLinksIdRequest) Get() *PutV1BundleproductsSkuLinksIdRequest {
	return v.value
}

func (v *NullablePutV1BundleproductsSkuLinksIdRequest) Set(val *PutV1BundleproductsSkuLinksIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1BundleproductsSkuLinksIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1BundleproductsSkuLinksIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1BundleproductsSkuLinksIdRequest(val *PutV1BundleproductsSkuLinksIdRequest) *NullablePutV1BundleproductsSkuLinksIdRequest {
	return &NullablePutV1BundleproductsSkuLinksIdRequest{value: val, isSet: true}
}

func (v NullablePutV1BundleproductsSkuLinksIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1BundleproductsSkuLinksIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
