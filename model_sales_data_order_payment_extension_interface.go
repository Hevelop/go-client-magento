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

// checks if the SalesDataOrderPaymentExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataOrderPaymentExtensionInterface{}

// SalesDataOrderPaymentExtensionInterface ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\OrderPaymentInterface
type SalesDataOrderPaymentExtensionInterface struct {
	NotificationMessage  *string                         `json:"notification_message,omitempty"`
	VaultPaymentToken    *VaultDataPaymentTokenInterface `json:"vault_payment_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataOrderPaymentExtensionInterface SalesDataOrderPaymentExtensionInterface

// NewSalesDataOrderPaymentExtensionInterface instantiates a new SalesDataOrderPaymentExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataOrderPaymentExtensionInterface() *SalesDataOrderPaymentExtensionInterface {
	this := SalesDataOrderPaymentExtensionInterface{}
	return &this
}

// NewSalesDataOrderPaymentExtensionInterfaceWithDefaults instantiates a new SalesDataOrderPaymentExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataOrderPaymentExtensionInterfaceWithDefaults() *SalesDataOrderPaymentExtensionInterface {
	this := SalesDataOrderPaymentExtensionInterface{}
	return &this
}

// GetNotificationMessage returns the NotificationMessage field value if set, zero value otherwise.
func (o *SalesDataOrderPaymentExtensionInterface) GetNotificationMessage() string {
	if o == nil || IsNil(o.NotificationMessage) {
		var ret string
		return ret
	}
	return *o.NotificationMessage
}

// GetNotificationMessageOk returns a tuple with the NotificationMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderPaymentExtensionInterface) GetNotificationMessageOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationMessage) {
		return nil, false
	}
	return o.NotificationMessage, true
}

// HasNotificationMessage returns a boolean if a field has been set.
func (o *SalesDataOrderPaymentExtensionInterface) HasNotificationMessage() bool {
	if o != nil && !IsNil(o.NotificationMessage) {
		return true
	}

	return false
}

// SetNotificationMessage gets a reference to the given string and assigns it to the NotificationMessage field.
func (o *SalesDataOrderPaymentExtensionInterface) SetNotificationMessage(v string) {
	o.NotificationMessage = &v
}

// GetVaultPaymentToken returns the VaultPaymentToken field value if set, zero value otherwise.
func (o *SalesDataOrderPaymentExtensionInterface) GetVaultPaymentToken() VaultDataPaymentTokenInterface {
	if o == nil || IsNil(o.VaultPaymentToken) {
		var ret VaultDataPaymentTokenInterface
		return ret
	}
	return *o.VaultPaymentToken
}

// GetVaultPaymentTokenOk returns a tuple with the VaultPaymentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderPaymentExtensionInterface) GetVaultPaymentTokenOk() (*VaultDataPaymentTokenInterface, bool) {
	if o == nil || IsNil(o.VaultPaymentToken) {
		return nil, false
	}
	return o.VaultPaymentToken, true
}

// HasVaultPaymentToken returns a boolean if a field has been set.
func (o *SalesDataOrderPaymentExtensionInterface) HasVaultPaymentToken() bool {
	if o != nil && !IsNil(o.VaultPaymentToken) {
		return true
	}

	return false
}

// SetVaultPaymentToken gets a reference to the given VaultDataPaymentTokenInterface and assigns it to the VaultPaymentToken field.
func (o *SalesDataOrderPaymentExtensionInterface) SetVaultPaymentToken(v VaultDataPaymentTokenInterface) {
	o.VaultPaymentToken = &v
}

func (o SalesDataOrderPaymentExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataOrderPaymentExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationMessage) {
		toSerialize["notification_message"] = o.NotificationMessage
	}
	if !IsNil(o.VaultPaymentToken) {
		toSerialize["vault_payment_token"] = o.VaultPaymentToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataOrderPaymentExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varSalesDataOrderPaymentExtensionInterface := _SalesDataOrderPaymentExtensionInterface{}

	err = json.Unmarshal(data, &varSalesDataOrderPaymentExtensionInterface)

	if err != nil {
		return err
	}

	*o = SalesDataOrderPaymentExtensionInterface(varSalesDataOrderPaymentExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "notification_message")
		delete(additionalProperties, "vault_payment_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataOrderPaymentExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataOrderPaymentExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataOrderPaymentExtensionInterface struct {
	value *SalesDataOrderPaymentExtensionInterface
	isSet bool
}

func (v NullableSalesDataOrderPaymentExtensionInterface) Get() *SalesDataOrderPaymentExtensionInterface {
	return v.value
}

func (v *NullableSalesDataOrderPaymentExtensionInterface) Set(val *SalesDataOrderPaymentExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataOrderPaymentExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataOrderPaymentExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataOrderPaymentExtensionInterface(val *SalesDataOrderPaymentExtensionInterface) *NullableSalesDataOrderPaymentExtensionInterface {
	return &NullableSalesDataOrderPaymentExtensionInterface{value: val, isSet: true}
}

func (v NullableSalesDataOrderPaymentExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataOrderPaymentExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
