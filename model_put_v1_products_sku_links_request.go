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

// checks if the PutV1ProductsSkuLinksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1ProductsSkuLinksRequest{}

// PutV1ProductsSkuLinksRequest struct for PutV1ProductsSkuLinksRequest
type PutV1ProductsSkuLinksRequest struct {
	Entity               CatalogDataProductLinkInterface `json:"entity"`
	AdditionalProperties map[string]interface{}
}

type _PutV1ProductsSkuLinksRequest PutV1ProductsSkuLinksRequest

// NewPutV1ProductsSkuLinksRequest instantiates a new PutV1ProductsSkuLinksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1ProductsSkuLinksRequest(entity CatalogDataProductLinkInterface) *PutV1ProductsSkuLinksRequest {
	this := PutV1ProductsSkuLinksRequest{}
	this.Entity = entity
	return &this
}

// NewPutV1ProductsSkuLinksRequestWithDefaults instantiates a new PutV1ProductsSkuLinksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1ProductsSkuLinksRequestWithDefaults() *PutV1ProductsSkuLinksRequest {
	this := PutV1ProductsSkuLinksRequest{}
	return &this
}

// GetEntity returns the Entity field value
func (o *PutV1ProductsSkuLinksRequest) GetEntity() CatalogDataProductLinkInterface {
	if o == nil {
		var ret CatalogDataProductLinkInterface
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *PutV1ProductsSkuLinksRequest) GetEntityOk() (*CatalogDataProductLinkInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *PutV1ProductsSkuLinksRequest) SetEntity(v CatalogDataProductLinkInterface) {
	o.Entity = v
}

func (o PutV1ProductsSkuLinksRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1ProductsSkuLinksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity"] = o.Entity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1ProductsSkuLinksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity",
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

	varPutV1ProductsSkuLinksRequest := _PutV1ProductsSkuLinksRequest{}

	err = json.Unmarshal(data, &varPutV1ProductsSkuLinksRequest)

	if err != nil {
		return err
	}

	*o = PutV1ProductsSkuLinksRequest(varPutV1ProductsSkuLinksRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1ProductsSkuLinksRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1ProductsSkuLinksRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1ProductsSkuLinksRequest struct {
	value *PutV1ProductsSkuLinksRequest
	isSet bool
}

func (v NullablePutV1ProductsSkuLinksRequest) Get() *PutV1ProductsSkuLinksRequest {
	return v.value
}

func (v *NullablePutV1ProductsSkuLinksRequest) Set(val *PutV1ProductsSkuLinksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1ProductsSkuLinksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1ProductsSkuLinksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1ProductsSkuLinksRequest(val *PutV1ProductsSkuLinksRequest) *NullablePutV1ProductsSkuLinksRequest {
	return &NullablePutV1ProductsSkuLinksRequest{value: val, isSet: true}
}

func (v NullablePutV1ProductsSkuLinksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1ProductsSkuLinksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}