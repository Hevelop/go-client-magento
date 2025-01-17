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

// checks if the CatalogDataProductTierPriceInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDataProductTierPriceInterface{}

// CatalogDataProductTierPriceInterface
type CatalogDataProductTierPriceInterface struct {
	// Customer group id
	CustomerGroupId int32 `json:"customer_group_id"`
	// Tier qty
	Qty float32 `json:"qty"`
	// Price value
	Value                float32                                        `json:"value"`
	ExtensionAttributes  *CatalogDataProductTierPriceExtensionInterface `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogDataProductTierPriceInterface CatalogDataProductTierPriceInterface

// NewCatalogDataProductTierPriceInterface instantiates a new CatalogDataProductTierPriceInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDataProductTierPriceInterface(customerGroupId int32, qty float32, value float32) *CatalogDataProductTierPriceInterface {
	this := CatalogDataProductTierPriceInterface{}
	this.CustomerGroupId = customerGroupId
	this.Qty = qty
	this.Value = value
	return &this
}

// NewCatalogDataProductTierPriceInterfaceWithDefaults instantiates a new CatalogDataProductTierPriceInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDataProductTierPriceInterfaceWithDefaults() *CatalogDataProductTierPriceInterface {
	this := CatalogDataProductTierPriceInterface{}
	return &this
}

// GetCustomerGroupId returns the CustomerGroupId field value
func (o *CatalogDataProductTierPriceInterface) GetCustomerGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomerGroupId
}

// GetCustomerGroupIdOk returns a tuple with the CustomerGroupId field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductTierPriceInterface) GetCustomerGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerGroupId, true
}

// SetCustomerGroupId sets field value
func (o *CatalogDataProductTierPriceInterface) SetCustomerGroupId(v int32) {
	o.CustomerGroupId = v
}

// GetQty returns the Qty field value
func (o *CatalogDataProductTierPriceInterface) GetQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductTierPriceInterface) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *CatalogDataProductTierPriceInterface) SetQty(v float32) {
	o.Qty = v
}

// GetValue returns the Value field value
func (o *CatalogDataProductTierPriceInterface) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductTierPriceInterface) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CatalogDataProductTierPriceInterface) SetValue(v float32) {
	o.Value = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *CatalogDataProductTierPriceInterface) GetExtensionAttributes() CatalogDataProductTierPriceExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret CatalogDataProductTierPriceExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogDataProductTierPriceInterface) GetExtensionAttributesOk() (*CatalogDataProductTierPriceExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *CatalogDataProductTierPriceInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given CatalogDataProductTierPriceExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *CatalogDataProductTierPriceInterface) SetExtensionAttributes(v CatalogDataProductTierPriceExtensionInterface) {
	o.ExtensionAttributes = &v
}

func (o CatalogDataProductTierPriceInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDataProductTierPriceInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer_group_id"] = o.CustomerGroupId
	toSerialize["qty"] = o.Qty
	toSerialize["value"] = o.Value
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogDataProductTierPriceInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer_group_id",
		"qty",
		"value",
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

	varCatalogDataProductTierPriceInterface := _CatalogDataProductTierPriceInterface{}

	err = json.Unmarshal(data, &varCatalogDataProductTierPriceInterface)

	if err != nil {
		return err
	}

	*o = CatalogDataProductTierPriceInterface(varCatalogDataProductTierPriceInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_group_id")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "value")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogDataProductTierPriceInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogDataProductTierPriceInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogDataProductTierPriceInterface struct {
	value *CatalogDataProductTierPriceInterface
	isSet bool
}

func (v NullableCatalogDataProductTierPriceInterface) Get() *CatalogDataProductTierPriceInterface {
	return v.value
}

func (v *NullableCatalogDataProductTierPriceInterface) Set(val *CatalogDataProductTierPriceInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDataProductTierPriceInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDataProductTierPriceInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDataProductTierPriceInterface(val *CatalogDataProductTierPriceInterface) *NullableCatalogDataProductTierPriceInterface {
	return &NullableCatalogDataProductTierPriceInterface{value: val, isSet: true}
}

func (v NullableCatalogDataProductTierPriceInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDataProductTierPriceInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
