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

// checks if the InventoryApiDataStockExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryApiDataStockExtensionInterface{}

// InventoryApiDataStockExtensionInterface TODO: temporal fix of extension classes generation during installation ExtensionInterface class for @see \\Magento\\InventoryApi\\Api\\Data\\StockInterface
type InventoryApiDataStockExtensionInterface struct {
	SalesChannels        []InventorySalesApiDataSalesChannelInterface `json:"sales_channels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryApiDataStockExtensionInterface InventoryApiDataStockExtensionInterface

// NewInventoryApiDataStockExtensionInterface instantiates a new InventoryApiDataStockExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryApiDataStockExtensionInterface() *InventoryApiDataStockExtensionInterface {
	this := InventoryApiDataStockExtensionInterface{}
	return &this
}

// NewInventoryApiDataStockExtensionInterfaceWithDefaults instantiates a new InventoryApiDataStockExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryApiDataStockExtensionInterfaceWithDefaults() *InventoryApiDataStockExtensionInterface {
	this := InventoryApiDataStockExtensionInterface{}
	return &this
}

// GetSalesChannels returns the SalesChannels field value if set, zero value otherwise.
func (o *InventoryApiDataStockExtensionInterface) GetSalesChannels() []InventorySalesApiDataSalesChannelInterface {
	if o == nil || IsNil(o.SalesChannels) {
		var ret []InventorySalesApiDataSalesChannelInterface
		return ret
	}
	return o.SalesChannels
}

// GetSalesChannelsOk returns a tuple with the SalesChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataStockExtensionInterface) GetSalesChannelsOk() ([]InventorySalesApiDataSalesChannelInterface, bool) {
	if o == nil || IsNil(o.SalesChannels) {
		return nil, false
	}
	return o.SalesChannels, true
}

// HasSalesChannels returns a boolean if a field has been set.
func (o *InventoryApiDataStockExtensionInterface) HasSalesChannels() bool {
	if o != nil && !IsNil(o.SalesChannels) {
		return true
	}

	return false
}

// SetSalesChannels gets a reference to the given []InventorySalesApiDataSalesChannelInterface and assigns it to the SalesChannels field.
func (o *InventoryApiDataStockExtensionInterface) SetSalesChannels(v []InventorySalesApiDataSalesChannelInterface) {
	o.SalesChannels = v
}

func (o InventoryApiDataStockExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryApiDataStockExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SalesChannels) {
		toSerialize["sales_channels"] = o.SalesChannels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryApiDataStockExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varInventoryApiDataStockExtensionInterface := _InventoryApiDataStockExtensionInterface{}

	err = json.Unmarshal(data, &varInventoryApiDataStockExtensionInterface)

	if err != nil {
		return err
	}

	*o = InventoryApiDataStockExtensionInterface(varInventoryApiDataStockExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sales_channels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryApiDataStockExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryApiDataStockExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryApiDataStockExtensionInterface struct {
	value *InventoryApiDataStockExtensionInterface
	isSet bool
}

func (v NullableInventoryApiDataStockExtensionInterface) Get() *InventoryApiDataStockExtensionInterface {
	return v.value
}

func (v *NullableInventoryApiDataStockExtensionInterface) Set(val *InventoryApiDataStockExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryApiDataStockExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryApiDataStockExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryApiDataStockExtensionInterface(val *InventoryApiDataStockExtensionInterface) *NullableInventoryApiDataStockExtensionInterface {
	return &NullableInventoryApiDataStockExtensionInterface{value: val, isSet: true}
}

func (v NullableInventoryApiDataStockExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryApiDataStockExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}