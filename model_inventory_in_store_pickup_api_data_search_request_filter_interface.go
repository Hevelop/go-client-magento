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

// checks if the InventoryInStorePickupApiDataSearchRequestFilterInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryInStorePickupApiDataSearchRequestFilterInterface{}

// InventoryInStorePickupApiDataSearchRequestFilterInterface Filter for Pickup Location search.
type InventoryInStorePickupApiDataSearchRequestFilterInterface struct {
	// Value.
	Value string `json:"value"`
	// Condition Type.
	ConditionType        string `json:"condition_type"`
	AdditionalProperties map[string]interface{}
}

type _InventoryInStorePickupApiDataSearchRequestFilterInterface InventoryInStorePickupApiDataSearchRequestFilterInterface

// NewInventoryInStorePickupApiDataSearchRequestFilterInterface instantiates a new InventoryInStorePickupApiDataSearchRequestFilterInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryInStorePickupApiDataSearchRequestFilterInterface(value string, conditionType string) *InventoryInStorePickupApiDataSearchRequestFilterInterface {
	this := InventoryInStorePickupApiDataSearchRequestFilterInterface{}
	this.Value = value
	this.ConditionType = conditionType
	return &this
}

// NewInventoryInStorePickupApiDataSearchRequestFilterInterfaceWithDefaults instantiates a new InventoryInStorePickupApiDataSearchRequestFilterInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryInStorePickupApiDataSearchRequestFilterInterfaceWithDefaults() *InventoryInStorePickupApiDataSearchRequestFilterInterface {
	this := InventoryInStorePickupApiDataSearchRequestFilterInterface{}
	return &this
}

// GetValue returns the Value field value
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) SetValue(v string) {
	o.Value = v
}

// GetConditionType returns the ConditionType field value
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) GetConditionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) SetConditionType(v string) {
	o.ConditionType = v
}

func (o InventoryInStorePickupApiDataSearchRequestFilterInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryInStorePickupApiDataSearchRequestFilterInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["condition_type"] = o.ConditionType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"condition_type",
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

	varInventoryInStorePickupApiDataSearchRequestFilterInterface := _InventoryInStorePickupApiDataSearchRequestFilterInterface{}

	err = json.Unmarshal(data, &varInventoryInStorePickupApiDataSearchRequestFilterInterface)

	if err != nil {
		return err
	}

	*o = InventoryInStorePickupApiDataSearchRequestFilterInterface(varInventoryInStorePickupApiDataSearchRequestFilterInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "condition_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryInStorePickupApiDataSearchRequestFilterInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryInStorePickupApiDataSearchRequestFilterInterface struct {
	value *InventoryInStorePickupApiDataSearchRequestFilterInterface
	isSet bool
}

func (v NullableInventoryInStorePickupApiDataSearchRequestFilterInterface) Get() *InventoryInStorePickupApiDataSearchRequestFilterInterface {
	return v.value
}

func (v *NullableInventoryInStorePickupApiDataSearchRequestFilterInterface) Set(val *InventoryInStorePickupApiDataSearchRequestFilterInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryInStorePickupApiDataSearchRequestFilterInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryInStorePickupApiDataSearchRequestFilterInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryInStorePickupApiDataSearchRequestFilterInterface(val *InventoryInStorePickupApiDataSearchRequestFilterInterface) *NullableInventoryInStorePickupApiDataSearchRequestFilterInterface {
	return &NullableInventoryInStorePickupApiDataSearchRequestFilterInterface{value: val, isSet: true}
}

func (v NullableInventoryInStorePickupApiDataSearchRequestFilterInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryInStorePickupApiDataSearchRequestFilterInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}