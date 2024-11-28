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

// checks if the InventoryApiDataStockSourceLinkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryApiDataStockSourceLinkInterface{}

// InventoryApiDataStockSourceLinkInterface Represents relation between Stock and Source entities. Used fully qualified namespaces in annotations for proper work of WebApi request parser
type InventoryApiDataStockSourceLinkInterface struct {
	// Stock id
	StockId *int32 `json:"stock_id,omitempty"`
	// Source code of the link
	SourceCode *string `json:"source_code,omitempty"`
	// Priority of the link
	Priority             *int32                                             `json:"priority,omitempty"`
	ExtensionAttributes  *InventoryApiDataStockSourceLinkExtensionInterface `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryApiDataStockSourceLinkInterface InventoryApiDataStockSourceLinkInterface

// NewInventoryApiDataStockSourceLinkInterface instantiates a new InventoryApiDataStockSourceLinkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryApiDataStockSourceLinkInterface() *InventoryApiDataStockSourceLinkInterface {
	this := InventoryApiDataStockSourceLinkInterface{}
	return &this
}

// NewInventoryApiDataStockSourceLinkInterfaceWithDefaults instantiates a new InventoryApiDataStockSourceLinkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryApiDataStockSourceLinkInterfaceWithDefaults() *InventoryApiDataStockSourceLinkInterface {
	this := InventoryApiDataStockSourceLinkInterface{}
	return &this
}

// GetStockId returns the StockId field value if set, zero value otherwise.
func (o *InventoryApiDataStockSourceLinkInterface) GetStockId() int32 {
	if o == nil || IsNil(o.StockId) {
		var ret int32
		return ret
	}
	return *o.StockId
}

// GetStockIdOk returns a tuple with the StockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataStockSourceLinkInterface) GetStockIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StockId) {
		return nil, false
	}
	return o.StockId, true
}

// HasStockId returns a boolean if a field has been set.
func (o *InventoryApiDataStockSourceLinkInterface) HasStockId() bool {
	if o != nil && !IsNil(o.StockId) {
		return true
	}

	return false
}

// SetStockId gets a reference to the given int32 and assigns it to the StockId field.
func (o *InventoryApiDataStockSourceLinkInterface) SetStockId(v int32) {
	o.StockId = &v
}

// GetSourceCode returns the SourceCode field value if set, zero value otherwise.
func (o *InventoryApiDataStockSourceLinkInterface) GetSourceCode() string {
	if o == nil || IsNil(o.SourceCode) {
		var ret string
		return ret
	}
	return *o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataStockSourceLinkInterface) GetSourceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCode) {
		return nil, false
	}
	return o.SourceCode, true
}

// HasSourceCode returns a boolean if a field has been set.
func (o *InventoryApiDataStockSourceLinkInterface) HasSourceCode() bool {
	if o != nil && !IsNil(o.SourceCode) {
		return true
	}

	return false
}

// SetSourceCode gets a reference to the given string and assigns it to the SourceCode field.
func (o *InventoryApiDataStockSourceLinkInterface) SetSourceCode(v string) {
	o.SourceCode = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *InventoryApiDataStockSourceLinkInterface) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataStockSourceLinkInterface) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InventoryApiDataStockSourceLinkInterface) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *InventoryApiDataStockSourceLinkInterface) SetPriority(v int32) {
	o.Priority = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventoryApiDataStockSourceLinkInterface) GetExtensionAttributes() InventoryApiDataStockSourceLinkExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret InventoryApiDataStockSourceLinkExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataStockSourceLinkInterface) GetExtensionAttributesOk() (*InventoryApiDataStockSourceLinkExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventoryApiDataStockSourceLinkInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given InventoryApiDataStockSourceLinkExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *InventoryApiDataStockSourceLinkInterface) SetExtensionAttributes(v InventoryApiDataStockSourceLinkExtensionInterface) {
	o.ExtensionAttributes = &v
}

func (o InventoryApiDataStockSourceLinkInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryApiDataStockSourceLinkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StockId) {
		toSerialize["stock_id"] = o.StockId
	}
	if !IsNil(o.SourceCode) {
		toSerialize["source_code"] = o.SourceCode
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryApiDataStockSourceLinkInterface) UnmarshalJSON(data []byte) (err error) {
	varInventoryApiDataStockSourceLinkInterface := _InventoryApiDataStockSourceLinkInterface{}

	err = json.Unmarshal(data, &varInventoryApiDataStockSourceLinkInterface)

	if err != nil {
		return err
	}

	*o = InventoryApiDataStockSourceLinkInterface(varInventoryApiDataStockSourceLinkInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "stock_id")
		delete(additionalProperties, "source_code")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryApiDataStockSourceLinkInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryApiDataStockSourceLinkInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryApiDataStockSourceLinkInterface struct {
	value *InventoryApiDataStockSourceLinkInterface
	isSet bool
}

func (v NullableInventoryApiDataStockSourceLinkInterface) Get() *InventoryApiDataStockSourceLinkInterface {
	return v.value
}

func (v *NullableInventoryApiDataStockSourceLinkInterface) Set(val *InventoryApiDataStockSourceLinkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryApiDataStockSourceLinkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryApiDataStockSourceLinkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryApiDataStockSourceLinkInterface(val *InventoryApiDataStockSourceLinkInterface) *NullableInventoryApiDataStockSourceLinkInterface {
	return &NullableInventoryApiDataStockSourceLinkInterface{value: val, isSet: true}
}

func (v NullableInventoryApiDataStockSourceLinkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryApiDataStockSourceLinkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}