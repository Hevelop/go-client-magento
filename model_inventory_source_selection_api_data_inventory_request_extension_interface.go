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

// checks if the InventorySourceSelectionApiDataInventoryRequestExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventorySourceSelectionApiDataInventoryRequestExtensionInterface{}

// InventorySourceSelectionApiDataInventoryRequestExtensionInterface ExtensionInterface class for @see \\Magento\\InventorySourceSelectionApi\\Api\\Data\\InventoryRequestInterface
type InventorySourceSelectionApiDataInventoryRequestExtensionInterface struct {
	DestinationAddress   *InventorySourceSelectionApiDataAddressInterface `json:"destination_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventorySourceSelectionApiDataInventoryRequestExtensionInterface InventorySourceSelectionApiDataInventoryRequestExtensionInterface

// NewInventorySourceSelectionApiDataInventoryRequestExtensionInterface instantiates a new InventorySourceSelectionApiDataInventoryRequestExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventorySourceSelectionApiDataInventoryRequestExtensionInterface() *InventorySourceSelectionApiDataInventoryRequestExtensionInterface {
	this := InventorySourceSelectionApiDataInventoryRequestExtensionInterface{}
	return &this
}

// NewInventorySourceSelectionApiDataInventoryRequestExtensionInterfaceWithDefaults instantiates a new InventorySourceSelectionApiDataInventoryRequestExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventorySourceSelectionApiDataInventoryRequestExtensionInterfaceWithDefaults() *InventorySourceSelectionApiDataInventoryRequestExtensionInterface {
	this := InventorySourceSelectionApiDataInventoryRequestExtensionInterface{}
	return &this
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) GetDestinationAddress() InventorySourceSelectionApiDataAddressInterface {
	if o == nil || IsNil(o.DestinationAddress) {
		var ret InventorySourceSelectionApiDataAddressInterface
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) GetDestinationAddressOk() (*InventorySourceSelectionApiDataAddressInterface, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) HasDestinationAddress() bool {
	if o != nil && !IsNil(o.DestinationAddress) {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given InventorySourceSelectionApiDataAddressInterface and assigns it to the DestinationAddress field.
func (o *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) SetDestinationAddress(v InventorySourceSelectionApiDataAddressInterface) {
	o.DestinationAddress = &v
}

func (o InventorySourceSelectionApiDataInventoryRequestExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventorySourceSelectionApiDataInventoryRequestExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestinationAddress) {
		toSerialize["destination_address"] = o.DestinationAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varInventorySourceSelectionApiDataInventoryRequestExtensionInterface := _InventorySourceSelectionApiDataInventoryRequestExtensionInterface{}

	err = json.Unmarshal(data, &varInventorySourceSelectionApiDataInventoryRequestExtensionInterface)

	if err != nil {
		return err
	}

	*o = InventorySourceSelectionApiDataInventoryRequestExtensionInterface(varInventorySourceSelectionApiDataInventoryRequestExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "destination_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface struct {
	value *InventorySourceSelectionApiDataInventoryRequestExtensionInterface
	isSet bool
}

func (v NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface) Get() *InventorySourceSelectionApiDataInventoryRequestExtensionInterface {
	return v.value
}

func (v *NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface) Set(val *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface(val *InventorySourceSelectionApiDataInventoryRequestExtensionInterface) *NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface {
	return &NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface{value: val, isSet: true}
}

func (v NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventorySourceSelectionApiDataInventoryRequestExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
