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

// checks if the GiftMessageDataMessageExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftMessageDataMessageExtensionInterface{}

// GiftMessageDataMessageExtensionInterface ExtensionInterface class for @see \\Magento\\GiftMessage\\Api\\Data\\MessageInterface
type GiftMessageDataMessageExtensionInterface struct {
	EntityId                 *string `json:"entity_id,omitempty"`
	EntityType               *string `json:"entity_type,omitempty"`
	WrappingId               *int32  `json:"wrapping_id,omitempty"`
	WrappingAllowGiftReceipt *bool   `json:"wrapping_allow_gift_receipt,omitempty"`
	WrappingAddPrintedCard   *bool   `json:"wrapping_add_printed_card,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _GiftMessageDataMessageExtensionInterface GiftMessageDataMessageExtensionInterface

// NewGiftMessageDataMessageExtensionInterface instantiates a new GiftMessageDataMessageExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftMessageDataMessageExtensionInterface() *GiftMessageDataMessageExtensionInterface {
	this := GiftMessageDataMessageExtensionInterface{}
	return &this
}

// NewGiftMessageDataMessageExtensionInterfaceWithDefaults instantiates a new GiftMessageDataMessageExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftMessageDataMessageExtensionInterfaceWithDefaults() *GiftMessageDataMessageExtensionInterface {
	this := GiftMessageDataMessageExtensionInterface{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *GiftMessageDataMessageExtensionInterface) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftMessageDataMessageExtensionInterface) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *GiftMessageDataMessageExtensionInterface) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *GiftMessageDataMessageExtensionInterface) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *GiftMessageDataMessageExtensionInterface) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftMessageDataMessageExtensionInterface) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *GiftMessageDataMessageExtensionInterface) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *GiftMessageDataMessageExtensionInterface) SetEntityType(v string) {
	o.EntityType = &v
}

// GetWrappingId returns the WrappingId field value if set, zero value otherwise.
func (o *GiftMessageDataMessageExtensionInterface) GetWrappingId() int32 {
	if o == nil || IsNil(o.WrappingId) {
		var ret int32
		return ret
	}
	return *o.WrappingId
}

// GetWrappingIdOk returns a tuple with the WrappingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftMessageDataMessageExtensionInterface) GetWrappingIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WrappingId) {
		return nil, false
	}
	return o.WrappingId, true
}

// HasWrappingId returns a boolean if a field has been set.
func (o *GiftMessageDataMessageExtensionInterface) HasWrappingId() bool {
	if o != nil && !IsNil(o.WrappingId) {
		return true
	}

	return false
}

// SetWrappingId gets a reference to the given int32 and assigns it to the WrappingId field.
func (o *GiftMessageDataMessageExtensionInterface) SetWrappingId(v int32) {
	o.WrappingId = &v
}

// GetWrappingAllowGiftReceipt returns the WrappingAllowGiftReceipt field value if set, zero value otherwise.
func (o *GiftMessageDataMessageExtensionInterface) GetWrappingAllowGiftReceipt() bool {
	if o == nil || IsNil(o.WrappingAllowGiftReceipt) {
		var ret bool
		return ret
	}
	return *o.WrappingAllowGiftReceipt
}

// GetWrappingAllowGiftReceiptOk returns a tuple with the WrappingAllowGiftReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftMessageDataMessageExtensionInterface) GetWrappingAllowGiftReceiptOk() (*bool, bool) {
	if o == nil || IsNil(o.WrappingAllowGiftReceipt) {
		return nil, false
	}
	return o.WrappingAllowGiftReceipt, true
}

// HasWrappingAllowGiftReceipt returns a boolean if a field has been set.
func (o *GiftMessageDataMessageExtensionInterface) HasWrappingAllowGiftReceipt() bool {
	if o != nil && !IsNil(o.WrappingAllowGiftReceipt) {
		return true
	}

	return false
}

// SetWrappingAllowGiftReceipt gets a reference to the given bool and assigns it to the WrappingAllowGiftReceipt field.
func (o *GiftMessageDataMessageExtensionInterface) SetWrappingAllowGiftReceipt(v bool) {
	o.WrappingAllowGiftReceipt = &v
}

// GetWrappingAddPrintedCard returns the WrappingAddPrintedCard field value if set, zero value otherwise.
func (o *GiftMessageDataMessageExtensionInterface) GetWrappingAddPrintedCard() bool {
	if o == nil || IsNil(o.WrappingAddPrintedCard) {
		var ret bool
		return ret
	}
	return *o.WrappingAddPrintedCard
}

// GetWrappingAddPrintedCardOk returns a tuple with the WrappingAddPrintedCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftMessageDataMessageExtensionInterface) GetWrappingAddPrintedCardOk() (*bool, bool) {
	if o == nil || IsNil(o.WrappingAddPrintedCard) {
		return nil, false
	}
	return o.WrappingAddPrintedCard, true
}

// HasWrappingAddPrintedCard returns a boolean if a field has been set.
func (o *GiftMessageDataMessageExtensionInterface) HasWrappingAddPrintedCard() bool {
	if o != nil && !IsNil(o.WrappingAddPrintedCard) {
		return true
	}

	return false
}

// SetWrappingAddPrintedCard gets a reference to the given bool and assigns it to the WrappingAddPrintedCard field.
func (o *GiftMessageDataMessageExtensionInterface) SetWrappingAddPrintedCard(v bool) {
	o.WrappingAddPrintedCard = &v
}

func (o GiftMessageDataMessageExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftMessageDataMessageExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.WrappingId) {
		toSerialize["wrapping_id"] = o.WrappingId
	}
	if !IsNil(o.WrappingAllowGiftReceipt) {
		toSerialize["wrapping_allow_gift_receipt"] = o.WrappingAllowGiftReceipt
	}
	if !IsNil(o.WrappingAddPrintedCard) {
		toSerialize["wrapping_add_printed_card"] = o.WrappingAddPrintedCard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftMessageDataMessageExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varGiftMessageDataMessageExtensionInterface := _GiftMessageDataMessageExtensionInterface{}

	err = json.Unmarshal(data, &varGiftMessageDataMessageExtensionInterface)

	if err != nil {
		return err
	}

	*o = GiftMessageDataMessageExtensionInterface(varGiftMessageDataMessageExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "entity_type")
		delete(additionalProperties, "wrapping_id")
		delete(additionalProperties, "wrapping_allow_gift_receipt")
		delete(additionalProperties, "wrapping_add_printed_card")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *GiftMessageDataMessageExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *GiftMessageDataMessageExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableGiftMessageDataMessageExtensionInterface struct {
	value *GiftMessageDataMessageExtensionInterface
	isSet bool
}

func (v NullableGiftMessageDataMessageExtensionInterface) Get() *GiftMessageDataMessageExtensionInterface {
	return v.value
}

func (v *NullableGiftMessageDataMessageExtensionInterface) Set(val *GiftMessageDataMessageExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftMessageDataMessageExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftMessageDataMessageExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftMessageDataMessageExtensionInterface(val *GiftMessageDataMessageExtensionInterface) *NullableGiftMessageDataMessageExtensionInterface {
	return &NullableGiftMessageDataMessageExtensionInterface{value: val, isSet: true}
}

func (v NullableGiftMessageDataMessageExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftMessageDataMessageExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}