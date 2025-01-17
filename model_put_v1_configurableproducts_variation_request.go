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

// checks if the PutV1ConfigurableproductsVariationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV1ConfigurableproductsVariationRequest{}

// PutV1ConfigurableproductsVariationRequest struct for PutV1ConfigurableproductsVariationRequest
type PutV1ConfigurableproductsVariationRequest struct {
	Product              CatalogDataProductInterface              `json:"product"`
	Options              []ConfigurableProductDataOptionInterface `json:"options"`
	AdditionalProperties map[string]interface{}
}

type _PutV1ConfigurableproductsVariationRequest PutV1ConfigurableproductsVariationRequest

// NewPutV1ConfigurableproductsVariationRequest instantiates a new PutV1ConfigurableproductsVariationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV1ConfigurableproductsVariationRequest(product CatalogDataProductInterface, options []ConfigurableProductDataOptionInterface) *PutV1ConfigurableproductsVariationRequest {
	this := PutV1ConfigurableproductsVariationRequest{}
	this.Product = product
	this.Options = options
	return &this
}

// NewPutV1ConfigurableproductsVariationRequestWithDefaults instantiates a new PutV1ConfigurableproductsVariationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV1ConfigurableproductsVariationRequestWithDefaults() *PutV1ConfigurableproductsVariationRequest {
	this := PutV1ConfigurableproductsVariationRequest{}
	return &this
}

// GetProduct returns the Product field value
func (o *PutV1ConfigurableproductsVariationRequest) GetProduct() CatalogDataProductInterface {
	if o == nil {
		var ret CatalogDataProductInterface
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *PutV1ConfigurableproductsVariationRequest) GetProductOk() (*CatalogDataProductInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *PutV1ConfigurableproductsVariationRequest) SetProduct(v CatalogDataProductInterface) {
	o.Product = v
}

// GetOptions returns the Options field value
func (o *PutV1ConfigurableproductsVariationRequest) GetOptions() []ConfigurableProductDataOptionInterface {
	if o == nil {
		var ret []ConfigurableProductDataOptionInterface
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *PutV1ConfigurableproductsVariationRequest) GetOptionsOk() ([]ConfigurableProductDataOptionInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *PutV1ConfigurableproductsVariationRequest) SetOptions(v []ConfigurableProductDataOptionInterface) {
	o.Options = v
}

func (o PutV1ConfigurableproductsVariationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV1ConfigurableproductsVariationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutV1ConfigurableproductsVariationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product",
		"options",
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

	varPutV1ConfigurableproductsVariationRequest := _PutV1ConfigurableproductsVariationRequest{}

	err = json.Unmarshal(data, &varPutV1ConfigurableproductsVariationRequest)

	if err != nil {
		return err
	}

	*o = PutV1ConfigurableproductsVariationRequest(varPutV1ConfigurableproductsVariationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PutV1ConfigurableproductsVariationRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PutV1ConfigurableproductsVariationRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePutV1ConfigurableproductsVariationRequest struct {
	value *PutV1ConfigurableproductsVariationRequest
	isSet bool
}

func (v NullablePutV1ConfigurableproductsVariationRequest) Get() *PutV1ConfigurableproductsVariationRequest {
	return v.value
}

func (v *NullablePutV1ConfigurableproductsVariationRequest) Set(val *PutV1ConfigurableproductsVariationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV1ConfigurableproductsVariationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV1ConfigurableproductsVariationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV1ConfigurableproductsVariationRequest(val *PutV1ConfigurableproductsVariationRequest) *NullablePutV1ConfigurableproductsVariationRequest {
	return &NullablePutV1ConfigurableproductsVariationRequest{value: val, isSet: true}
}

func (v NullablePutV1ConfigurableproductsVariationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV1ConfigurableproductsVariationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
