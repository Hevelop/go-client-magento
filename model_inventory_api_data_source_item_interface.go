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

// checks if the InventoryApiDataSourceItemInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryApiDataSourceItemInterface{}

// InventoryApiDataSourceItemInterface Represents amount of product on physical storage Entity id getter is missed because entity identifies by compound identifier (sku and source_code) Used fully qualified namespaces in annotations for proper work of WebApi request parser
type InventoryApiDataSourceItemInterface struct {
	// Source item sku
	Sku *string `json:"sku,omitempty"`
	// Source code
	SourceCode *string `json:"source_code,omitempty"`
	// Source item quantity
	Quantity *float32 `json:"quantity,omitempty"`
	// Source item status (One of self::STATUS_*)
	Status *int32 `json:"status,omitempty"`
	// ExtensionInterface class for @see \\Magento\\InventoryApi\\Api\\Data\\SourceItemInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryApiDataSourceItemInterface InventoryApiDataSourceItemInterface

// NewInventoryApiDataSourceItemInterface instantiates a new InventoryApiDataSourceItemInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryApiDataSourceItemInterface() *InventoryApiDataSourceItemInterface {
	this := InventoryApiDataSourceItemInterface{}
	return &this
}

// NewInventoryApiDataSourceItemInterfaceWithDefaults instantiates a new InventoryApiDataSourceItemInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryApiDataSourceItemInterfaceWithDefaults() *InventoryApiDataSourceItemInterface {
	this := InventoryApiDataSourceItemInterface{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InventoryApiDataSourceItemInterface) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceItemInterface) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InventoryApiDataSourceItemInterface) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *InventoryApiDataSourceItemInterface) SetSku(v string) {
	o.Sku = &v
}

// GetSourceCode returns the SourceCode field value if set, zero value otherwise.
func (o *InventoryApiDataSourceItemInterface) GetSourceCode() string {
	if o == nil || IsNil(o.SourceCode) {
		var ret string
		return ret
	}
	return *o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceItemInterface) GetSourceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCode) {
		return nil, false
	}
	return o.SourceCode, true
}

// HasSourceCode returns a boolean if a field has been set.
func (o *InventoryApiDataSourceItemInterface) HasSourceCode() bool {
	if o != nil && !IsNil(o.SourceCode) {
		return true
	}

	return false
}

// SetSourceCode gets a reference to the given string and assigns it to the SourceCode field.
func (o *InventoryApiDataSourceItemInterface) SetSourceCode(v string) {
	o.SourceCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *InventoryApiDataSourceItemInterface) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceItemInterface) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *InventoryApiDataSourceItemInterface) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *InventoryApiDataSourceItemInterface) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InventoryApiDataSourceItemInterface) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceItemInterface) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InventoryApiDataSourceItemInterface) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *InventoryApiDataSourceItemInterface) SetStatus(v int32) {
	o.Status = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventoryApiDataSourceItemInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceItemInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventoryApiDataSourceItemInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *InventoryApiDataSourceItemInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o InventoryApiDataSourceItemInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryApiDataSourceItemInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.SourceCode) {
		toSerialize["source_code"] = o.SourceCode
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryApiDataSourceItemInterface) UnmarshalJSON(data []byte) (err error) {
	varInventoryApiDataSourceItemInterface := _InventoryApiDataSourceItemInterface{}

	err = json.Unmarshal(data, &varInventoryApiDataSourceItemInterface)

	if err != nil {
		return err
	}

	*o = InventoryApiDataSourceItemInterface(varInventoryApiDataSourceItemInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sku")
		delete(additionalProperties, "source_code")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "status")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryApiDataSourceItemInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryApiDataSourceItemInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryApiDataSourceItemInterface struct {
	value *InventoryApiDataSourceItemInterface
	isSet bool
}

func (v NullableInventoryApiDataSourceItemInterface) Get() *InventoryApiDataSourceItemInterface {
	return v.value
}

func (v *NullableInventoryApiDataSourceItemInterface) Set(val *InventoryApiDataSourceItemInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryApiDataSourceItemInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryApiDataSourceItemInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryApiDataSourceItemInterface(val *InventoryApiDataSourceItemInterface) *NullableInventoryApiDataSourceItemInterface {
	return &NullableInventoryApiDataSourceItemInterface{value: val, isSet: true}
}

func (v NullableInventoryApiDataSourceItemInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryApiDataSourceItemInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
