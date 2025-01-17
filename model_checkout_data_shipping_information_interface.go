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

// checks if the CheckoutDataShippingInformationInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutDataShippingInformationInterface{}

// CheckoutDataShippingInformationInterface Interface ShippingInformationInterface
type CheckoutDataShippingInformationInterface struct {
	ShippingAddress QuoteDataAddressInterface  `json:"shipping_address"`
	BillingAddress  *QuoteDataAddressInterface `json:"billing_address,omitempty"`
	// Shipping method code
	ShippingMethodCode string `json:"shipping_method_code"`
	// Carrier code
	ShippingCarrierCode string `json:"shipping_carrier_code"`
	// ExtensionInterface class for @see \\Magento\\Checkout\\Api\\Data\\ShippingInformationInterface
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
	// Custom attributes values.
	CustomAttributes     []FrameworkAttributeInterface `json:"custom_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckoutDataShippingInformationInterface CheckoutDataShippingInformationInterface

// NewCheckoutDataShippingInformationInterface instantiates a new CheckoutDataShippingInformationInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutDataShippingInformationInterface(shippingAddress QuoteDataAddressInterface, shippingMethodCode string, shippingCarrierCode string) *CheckoutDataShippingInformationInterface {
	this := CheckoutDataShippingInformationInterface{}
	this.ShippingAddress = shippingAddress
	this.ShippingMethodCode = shippingMethodCode
	this.ShippingCarrierCode = shippingCarrierCode
	return &this
}

// NewCheckoutDataShippingInformationInterfaceWithDefaults instantiates a new CheckoutDataShippingInformationInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutDataShippingInformationInterfaceWithDefaults() *CheckoutDataShippingInformationInterface {
	this := CheckoutDataShippingInformationInterface{}
	return &this
}

// GetShippingAddress returns the ShippingAddress field value
func (o *CheckoutDataShippingInformationInterface) GetShippingAddress() QuoteDataAddressInterface {
	if o == nil {
		var ret QuoteDataAddressInterface
		return ret
	}

	return o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value
// and a boolean to check if the value has been set.
func (o *CheckoutDataShippingInformationInterface) GetShippingAddressOk() (*QuoteDataAddressInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingAddress, true
}

// SetShippingAddress sets field value
func (o *CheckoutDataShippingInformationInterface) SetShippingAddress(v QuoteDataAddressInterface) {
	o.ShippingAddress = v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *CheckoutDataShippingInformationInterface) GetBillingAddress() QuoteDataAddressInterface {
	if o == nil || IsNil(o.BillingAddress) {
		var ret QuoteDataAddressInterface
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDataShippingInformationInterface) GetBillingAddressOk() (*QuoteDataAddressInterface, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *CheckoutDataShippingInformationInterface) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given QuoteDataAddressInterface and assigns it to the BillingAddress field.
func (o *CheckoutDataShippingInformationInterface) SetBillingAddress(v QuoteDataAddressInterface) {
	o.BillingAddress = &v
}

// GetShippingMethodCode returns the ShippingMethodCode field value
func (o *CheckoutDataShippingInformationInterface) GetShippingMethodCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingMethodCode
}

// GetShippingMethodCodeOk returns a tuple with the ShippingMethodCode field value
// and a boolean to check if the value has been set.
func (o *CheckoutDataShippingInformationInterface) GetShippingMethodCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingMethodCode, true
}

// SetShippingMethodCode sets field value
func (o *CheckoutDataShippingInformationInterface) SetShippingMethodCode(v string) {
	o.ShippingMethodCode = v
}

// GetShippingCarrierCode returns the ShippingCarrierCode field value
func (o *CheckoutDataShippingInformationInterface) GetShippingCarrierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingCarrierCode
}

// GetShippingCarrierCodeOk returns a tuple with the ShippingCarrierCode field value
// and a boolean to check if the value has been set.
func (o *CheckoutDataShippingInformationInterface) GetShippingCarrierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingCarrierCode, true
}

// SetShippingCarrierCode sets field value
func (o *CheckoutDataShippingInformationInterface) SetShippingCarrierCode(v string) {
	o.ShippingCarrierCode = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *CheckoutDataShippingInformationInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDataShippingInformationInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *CheckoutDataShippingInformationInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *CheckoutDataShippingInformationInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise.
func (o *CheckoutDataShippingInformationInterface) GetCustomAttributes() []FrameworkAttributeInterface {
	if o == nil || IsNil(o.CustomAttributes) {
		var ret []FrameworkAttributeInterface
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDataShippingInformationInterface) GetCustomAttributesOk() ([]FrameworkAttributeInterface, bool) {
	if o == nil || IsNil(o.CustomAttributes) {
		return nil, false
	}
	return o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *CheckoutDataShippingInformationInterface) HasCustomAttributes() bool {
	if o != nil && !IsNil(o.CustomAttributes) {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []FrameworkAttributeInterface and assigns it to the CustomAttributes field.
func (o *CheckoutDataShippingInformationInterface) SetCustomAttributes(v []FrameworkAttributeInterface) {
	o.CustomAttributes = v
}

func (o CheckoutDataShippingInformationInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutDataShippingInformationInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipping_address"] = o.ShippingAddress
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	toSerialize["shipping_method_code"] = o.ShippingMethodCode
	toSerialize["shipping_carrier_code"] = o.ShippingCarrierCode
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}
	if !IsNil(o.CustomAttributes) {
		toSerialize["custom_attributes"] = o.CustomAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckoutDataShippingInformationInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shipping_address",
		"shipping_method_code",
		"shipping_carrier_code",
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

	varCheckoutDataShippingInformationInterface := _CheckoutDataShippingInformationInterface{}

	err = json.Unmarshal(data, &varCheckoutDataShippingInformationInterface)

	if err != nil {
		return err
	}

	*o = CheckoutDataShippingInformationInterface(varCheckoutDataShippingInformationInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "shipping_address")
		delete(additionalProperties, "billing_address")
		delete(additionalProperties, "shipping_method_code")
		delete(additionalProperties, "shipping_carrier_code")
		delete(additionalProperties, "extension_attributes")
		delete(additionalProperties, "custom_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CheckoutDataShippingInformationInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CheckoutDataShippingInformationInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCheckoutDataShippingInformationInterface struct {
	value *CheckoutDataShippingInformationInterface
	isSet bool
}

func (v NullableCheckoutDataShippingInformationInterface) Get() *CheckoutDataShippingInformationInterface {
	return v.value
}

func (v *NullableCheckoutDataShippingInformationInterface) Set(val *CheckoutDataShippingInformationInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutDataShippingInformationInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutDataShippingInformationInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutDataShippingInformationInterface(val *CheckoutDataShippingInformationInterface) *NullableCheckoutDataShippingInformationInterface {
	return &NullableCheckoutDataShippingInformationInterface{value: val, isSet: true}
}

func (v NullableCheckoutDataShippingInformationInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutDataShippingInformationInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
