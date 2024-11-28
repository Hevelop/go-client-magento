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

// checks if the GiftCardDataGiftCardOptionExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardDataGiftCardOptionExtensionInterface{}

// GiftCardDataGiftCardOptionExtensionInterface ExtensionInterface class for @see \\Magento\\GiftCard\\Api\\Data\\GiftCardOptionInterface
type GiftCardDataGiftCardOptionExtensionInterface struct {
	GiftcardCreatedCodes []string `json:"giftcard_created_codes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftCardDataGiftCardOptionExtensionInterface GiftCardDataGiftCardOptionExtensionInterface

// NewGiftCardDataGiftCardOptionExtensionInterface instantiates a new GiftCardDataGiftCardOptionExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardDataGiftCardOptionExtensionInterface() *GiftCardDataGiftCardOptionExtensionInterface {
	this := GiftCardDataGiftCardOptionExtensionInterface{}
	return &this
}

// NewGiftCardDataGiftCardOptionExtensionInterfaceWithDefaults instantiates a new GiftCardDataGiftCardOptionExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardDataGiftCardOptionExtensionInterfaceWithDefaults() *GiftCardDataGiftCardOptionExtensionInterface {
	this := GiftCardDataGiftCardOptionExtensionInterface{}
	return &this
}

// GetGiftcardCreatedCodes returns the GiftcardCreatedCodes field value if set, zero value otherwise.
func (o *GiftCardDataGiftCardOptionExtensionInterface) GetGiftcardCreatedCodes() []string {
	if o == nil || IsNil(o.GiftcardCreatedCodes) {
		var ret []string
		return ret
	}
	return o.GiftcardCreatedCodes
}

// GetGiftcardCreatedCodesOk returns a tuple with the GiftcardCreatedCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataGiftCardOptionExtensionInterface) GetGiftcardCreatedCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.GiftcardCreatedCodes) {
		return nil, false
	}
	return o.GiftcardCreatedCodes, true
}

// HasGiftcardCreatedCodes returns a boolean if a field has been set.
func (o *GiftCardDataGiftCardOptionExtensionInterface) HasGiftcardCreatedCodes() bool {
	if o != nil && !IsNil(o.GiftcardCreatedCodes) {
		return true
	}

	return false
}

// SetGiftcardCreatedCodes gets a reference to the given []string and assigns it to the GiftcardCreatedCodes field.
func (o *GiftCardDataGiftCardOptionExtensionInterface) SetGiftcardCreatedCodes(v []string) {
	o.GiftcardCreatedCodes = v
}

func (o GiftCardDataGiftCardOptionExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardDataGiftCardOptionExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GiftcardCreatedCodes) {
		toSerialize["giftcard_created_codes"] = o.GiftcardCreatedCodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftCardDataGiftCardOptionExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varGiftCardDataGiftCardOptionExtensionInterface := _GiftCardDataGiftCardOptionExtensionInterface{}

	err = json.Unmarshal(data, &varGiftCardDataGiftCardOptionExtensionInterface)

	if err != nil {
		return err
	}

	*o = GiftCardDataGiftCardOptionExtensionInterface(varGiftCardDataGiftCardOptionExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "giftcard_created_codes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *GiftCardDataGiftCardOptionExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *GiftCardDataGiftCardOptionExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableGiftCardDataGiftCardOptionExtensionInterface struct {
	value *GiftCardDataGiftCardOptionExtensionInterface
	isSet bool
}

func (v NullableGiftCardDataGiftCardOptionExtensionInterface) Get() *GiftCardDataGiftCardOptionExtensionInterface {
	return v.value
}

func (v *NullableGiftCardDataGiftCardOptionExtensionInterface) Set(val *GiftCardDataGiftCardOptionExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardDataGiftCardOptionExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardDataGiftCardOptionExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardDataGiftCardOptionExtensionInterface(val *GiftCardDataGiftCardOptionExtensionInterface) *NullableGiftCardDataGiftCardOptionExtensionInterface {
	return &NullableGiftCardDataGiftCardOptionExtensionInterface{value: val, isSet: true}
}

func (v NullableGiftCardDataGiftCardOptionExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardDataGiftCardOptionExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}