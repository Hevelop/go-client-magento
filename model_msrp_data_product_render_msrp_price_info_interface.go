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

// checks if the MsrpDataProductRenderMsrpPriceInfoInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MsrpDataProductRenderMsrpPriceInfoInterface{}

// MsrpDataProductRenderMsrpPriceInfoInterface Price interface.
type MsrpDataProductRenderMsrpPriceInfoInterface struct {
	MsrpPrice             string `json:"msrp_price"`
	IsApplicable          string `json:"is_applicable"`
	IsShownPriceOnGesture string `json:"is_shown_price_on_gesture"`
	MsrpMessage           string `json:"msrp_message"`
	ExplanationMessage    string `json:"explanation_message"`
	// ExtensionInterface class for @see \\Magento\\Msrp\\Api\\Data\\ProductRender\\MsrpPriceInfoInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MsrpDataProductRenderMsrpPriceInfoInterface MsrpDataProductRenderMsrpPriceInfoInterface

// NewMsrpDataProductRenderMsrpPriceInfoInterface instantiates a new MsrpDataProductRenderMsrpPriceInfoInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsrpDataProductRenderMsrpPriceInfoInterface(msrpPrice string, isApplicable string, isShownPriceOnGesture string, msrpMessage string, explanationMessage string) *MsrpDataProductRenderMsrpPriceInfoInterface {
	this := MsrpDataProductRenderMsrpPriceInfoInterface{}
	this.MsrpPrice = msrpPrice
	this.IsApplicable = isApplicable
	this.IsShownPriceOnGesture = isShownPriceOnGesture
	this.MsrpMessage = msrpMessage
	this.ExplanationMessage = explanationMessage
	return &this
}

// NewMsrpDataProductRenderMsrpPriceInfoInterfaceWithDefaults instantiates a new MsrpDataProductRenderMsrpPriceInfoInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsrpDataProductRenderMsrpPriceInfoInterfaceWithDefaults() *MsrpDataProductRenderMsrpPriceInfoInterface {
	this := MsrpDataProductRenderMsrpPriceInfoInterface{}
	return &this
}

// GetMsrpPrice returns the MsrpPrice field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetMsrpPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsrpPrice
}

// GetMsrpPriceOk returns a tuple with the MsrpPrice field value
// and a boolean to check if the value has been set.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetMsrpPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsrpPrice, true
}

// SetMsrpPrice sets field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) SetMsrpPrice(v string) {
	o.MsrpPrice = v
}

// GetIsApplicable returns the IsApplicable field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetIsApplicable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsApplicable
}

// GetIsApplicableOk returns a tuple with the IsApplicable field value
// and a boolean to check if the value has been set.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetIsApplicableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsApplicable, true
}

// SetIsApplicable sets field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) SetIsApplicable(v string) {
	o.IsApplicable = v
}

// GetIsShownPriceOnGesture returns the IsShownPriceOnGesture field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetIsShownPriceOnGesture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsShownPriceOnGesture
}

// GetIsShownPriceOnGestureOk returns a tuple with the IsShownPriceOnGesture field value
// and a boolean to check if the value has been set.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetIsShownPriceOnGestureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsShownPriceOnGesture, true
}

// SetIsShownPriceOnGesture sets field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) SetIsShownPriceOnGesture(v string) {
	o.IsShownPriceOnGesture = v
}

// GetMsrpMessage returns the MsrpMessage field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetMsrpMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsrpMessage
}

// GetMsrpMessageOk returns a tuple with the MsrpMessage field value
// and a boolean to check if the value has been set.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetMsrpMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsrpMessage, true
}

// SetMsrpMessage sets field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) SetMsrpMessage(v string) {
	o.MsrpMessage = v
}

// GetExplanationMessage returns the ExplanationMessage field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetExplanationMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExplanationMessage
}

// GetExplanationMessageOk returns a tuple with the ExplanationMessage field value
// and a boolean to check if the value has been set.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetExplanationMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplanationMessage, true
}

// SetExplanationMessage sets field value
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) SetExplanationMessage(v string) {
	o.ExplanationMessage = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o MsrpDataProductRenderMsrpPriceInfoInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MsrpDataProductRenderMsrpPriceInfoInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msrp_price"] = o.MsrpPrice
	toSerialize["is_applicable"] = o.IsApplicable
	toSerialize["is_shown_price_on_gesture"] = o.IsShownPriceOnGesture
	toSerialize["msrp_message"] = o.MsrpMessage
	toSerialize["explanation_message"] = o.ExplanationMessage
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MsrpDataProductRenderMsrpPriceInfoInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"msrp_price",
		"is_applicable",
		"is_shown_price_on_gesture",
		"msrp_message",
		"explanation_message",
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

	varMsrpDataProductRenderMsrpPriceInfoInterface := _MsrpDataProductRenderMsrpPriceInfoInterface{}

	err = json.Unmarshal(data, &varMsrpDataProductRenderMsrpPriceInfoInterface)

	if err != nil {
		return err
	}

	*o = MsrpDataProductRenderMsrpPriceInfoInterface(varMsrpDataProductRenderMsrpPriceInfoInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msrp_price")
		delete(additionalProperties, "is_applicable")
		delete(additionalProperties, "is_shown_price_on_gesture")
		delete(additionalProperties, "msrp_message")
		delete(additionalProperties, "explanation_message")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *MsrpDataProductRenderMsrpPriceInfoInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableMsrpDataProductRenderMsrpPriceInfoInterface struct {
	value *MsrpDataProductRenderMsrpPriceInfoInterface
	isSet bool
}

func (v NullableMsrpDataProductRenderMsrpPriceInfoInterface) Get() *MsrpDataProductRenderMsrpPriceInfoInterface {
	return v.value
}

func (v *NullableMsrpDataProductRenderMsrpPriceInfoInterface) Set(val *MsrpDataProductRenderMsrpPriceInfoInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableMsrpDataProductRenderMsrpPriceInfoInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableMsrpDataProductRenderMsrpPriceInfoInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsrpDataProductRenderMsrpPriceInfoInterface(val *MsrpDataProductRenderMsrpPriceInfoInterface) *NullableMsrpDataProductRenderMsrpPriceInfoInterface {
	return &NullableMsrpDataProductRenderMsrpPriceInfoInterface{value: val, isSet: true}
}

func (v NullableMsrpDataProductRenderMsrpPriceInfoInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsrpDataProductRenderMsrpPriceInfoInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}