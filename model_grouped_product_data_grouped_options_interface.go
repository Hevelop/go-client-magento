/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
)

// checks if the GroupedProductDataGroupedOptionsInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedProductDataGroupedOptionsInterface{}

// GroupedProductDataGroupedOptionsInterface Represents `product item id with qty` of a grouped product.
type GroupedProductDataGroupedOptionsInterface struct {
	// Associated product id
	Id *int32 `json:"id,omitempty"`
	// Associated product qty
	Qty *int32 `json:"qty,omitempty"`
	// ExtensionInterface class for @see \\Magento\\GroupedProduct\\Api\\Data\\GroupedOptionsInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupedProductDataGroupedOptionsInterface GroupedProductDataGroupedOptionsInterface

// NewGroupedProductDataGroupedOptionsInterface instantiates a new GroupedProductDataGroupedOptionsInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedProductDataGroupedOptionsInterface() *GroupedProductDataGroupedOptionsInterface {
	this := GroupedProductDataGroupedOptionsInterface{}
	return &this
}

// NewGroupedProductDataGroupedOptionsInterfaceWithDefaults instantiates a new GroupedProductDataGroupedOptionsInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedProductDataGroupedOptionsInterfaceWithDefaults() *GroupedProductDataGroupedOptionsInterface {
	this := GroupedProductDataGroupedOptionsInterface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupedProductDataGroupedOptionsInterface) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedProductDataGroupedOptionsInterface) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupedProductDataGroupedOptionsInterface) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GroupedProductDataGroupedOptionsInterface) SetId(v int32) {
	o.Id = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *GroupedProductDataGroupedOptionsInterface) GetQty() int32 {
	if o == nil || IsNil(o.Qty) {
		var ret int32
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedProductDataGroupedOptionsInterface) GetQtyOk() (*int32, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *GroupedProductDataGroupedOptionsInterface) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given int32 and assigns it to the Qty field.
func (o *GroupedProductDataGroupedOptionsInterface) SetQty(v int32) {
	o.Qty = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *GroupedProductDataGroupedOptionsInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedProductDataGroupedOptionsInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *GroupedProductDataGroupedOptionsInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *GroupedProductDataGroupedOptionsInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o GroupedProductDataGroupedOptionsInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedProductDataGroupedOptionsInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupedProductDataGroupedOptionsInterface) UnmarshalJSON(data []byte) (err error) {
	varGroupedProductDataGroupedOptionsInterface := _GroupedProductDataGroupedOptionsInterface{}

	err = json.Unmarshal(data, &varGroupedProductDataGroupedOptionsInterface)

	if err != nil {
		return err
	}

	*o = GroupedProductDataGroupedOptionsInterface(varGroupedProductDataGroupedOptionsInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *GroupedProductDataGroupedOptionsInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *GroupedProductDataGroupedOptionsInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableGroupedProductDataGroupedOptionsInterface struct {
	value *GroupedProductDataGroupedOptionsInterface
	isSet bool
}

func (v NullableGroupedProductDataGroupedOptionsInterface) Get() *GroupedProductDataGroupedOptionsInterface {
	return v.value
}

func (v *NullableGroupedProductDataGroupedOptionsInterface) Set(val *GroupedProductDataGroupedOptionsInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedProductDataGroupedOptionsInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedProductDataGroupedOptionsInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedProductDataGroupedOptionsInterface(val *GroupedProductDataGroupedOptionsInterface) *NullableGroupedProductDataGroupedOptionsInterface {
	return &NullableGroupedProductDataGroupedOptionsInterface{value: val, isSet: true}
}

func (v NullableGroupedProductDataGroupedOptionsInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedProductDataGroupedOptionsInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}