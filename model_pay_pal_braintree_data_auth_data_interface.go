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

// checks if the PayPalBraintreeDataAuthDataInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayPalBraintreeDataAuthDataInterface{}

// PayPalBraintreeDataAuthDataInterface Interface AuthDataInterface
type PayPalBraintreeDataAuthDataInterface struct {
	// Client token
	ClientToken *string `json:"client_token,omitempty"`
	// Display name
	DisplayName string `json:"display_name"`
	// To success page
	ActionSuccess string `json:"action_success"`
	LoggedIn      bool   `json:"logged_in"`
	// Current store code
	StoreCode            string `json:"store_code"`
	AdditionalProperties map[string]interface{}
}

type _PayPalBraintreeDataAuthDataInterface PayPalBraintreeDataAuthDataInterface

// NewPayPalBraintreeDataAuthDataInterface instantiates a new PayPalBraintreeDataAuthDataInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayPalBraintreeDataAuthDataInterface(displayName string, actionSuccess string, loggedIn bool, storeCode string) *PayPalBraintreeDataAuthDataInterface {
	this := PayPalBraintreeDataAuthDataInterface{}
	this.DisplayName = displayName
	this.ActionSuccess = actionSuccess
	this.LoggedIn = loggedIn
	this.StoreCode = storeCode
	return &this
}

// NewPayPalBraintreeDataAuthDataInterfaceWithDefaults instantiates a new PayPalBraintreeDataAuthDataInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayPalBraintreeDataAuthDataInterfaceWithDefaults() *PayPalBraintreeDataAuthDataInterface {
	this := PayPalBraintreeDataAuthDataInterface{}
	return &this
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *PayPalBraintreeDataAuthDataInterface) GetClientToken() string {
	if o == nil || IsNil(o.ClientToken) {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPalBraintreeDataAuthDataInterface) GetClientTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ClientToken) {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *PayPalBraintreeDataAuthDataInterface) HasClientToken() bool {
	if o != nil && !IsNil(o.ClientToken) {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *PayPalBraintreeDataAuthDataInterface) SetClientToken(v string) {
	o.ClientToken = &v
}

// GetDisplayName returns the DisplayName field value
func (o *PayPalBraintreeDataAuthDataInterface) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PayPalBraintreeDataAuthDataInterface) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PayPalBraintreeDataAuthDataInterface) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetActionSuccess returns the ActionSuccess field value
func (o *PayPalBraintreeDataAuthDataInterface) GetActionSuccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionSuccess
}

// GetActionSuccessOk returns a tuple with the ActionSuccess field value
// and a boolean to check if the value has been set.
func (o *PayPalBraintreeDataAuthDataInterface) GetActionSuccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionSuccess, true
}

// SetActionSuccess sets field value
func (o *PayPalBraintreeDataAuthDataInterface) SetActionSuccess(v string) {
	o.ActionSuccess = v
}

// GetLoggedIn returns the LoggedIn field value
func (o *PayPalBraintreeDataAuthDataInterface) GetLoggedIn() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LoggedIn
}

// GetLoggedInOk returns a tuple with the LoggedIn field value
// and a boolean to check if the value has been set.
func (o *PayPalBraintreeDataAuthDataInterface) GetLoggedInOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoggedIn, true
}

// SetLoggedIn sets field value
func (o *PayPalBraintreeDataAuthDataInterface) SetLoggedIn(v bool) {
	o.LoggedIn = v
}

// GetStoreCode returns the StoreCode field value
func (o *PayPalBraintreeDataAuthDataInterface) GetStoreCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoreCode
}

// GetStoreCodeOk returns a tuple with the StoreCode field value
// and a boolean to check if the value has been set.
func (o *PayPalBraintreeDataAuthDataInterface) GetStoreCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreCode, true
}

// SetStoreCode sets field value
func (o *PayPalBraintreeDataAuthDataInterface) SetStoreCode(v string) {
	o.StoreCode = v
}

func (o PayPalBraintreeDataAuthDataInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayPalBraintreeDataAuthDataInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientToken) {
		toSerialize["client_token"] = o.ClientToken
	}
	toSerialize["display_name"] = o.DisplayName
	toSerialize["action_success"] = o.ActionSuccess
	toSerialize["logged_in"] = o.LoggedIn
	toSerialize["store_code"] = o.StoreCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PayPalBraintreeDataAuthDataInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display_name",
		"action_success",
		"logged_in",
		"store_code",
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

	varPayPalBraintreeDataAuthDataInterface := _PayPalBraintreeDataAuthDataInterface{}

	err = json.Unmarshal(data, &varPayPalBraintreeDataAuthDataInterface)

	if err != nil {
		return err
	}

	*o = PayPalBraintreeDataAuthDataInterface(varPayPalBraintreeDataAuthDataInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_token")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "action_success")
		delete(additionalProperties, "logged_in")
		delete(additionalProperties, "store_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PayPalBraintreeDataAuthDataInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PayPalBraintreeDataAuthDataInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePayPalBraintreeDataAuthDataInterface struct {
	value *PayPalBraintreeDataAuthDataInterface
	isSet bool
}

func (v NullablePayPalBraintreeDataAuthDataInterface) Get() *PayPalBraintreeDataAuthDataInterface {
	return v.value
}

func (v *NullablePayPalBraintreeDataAuthDataInterface) Set(val *PayPalBraintreeDataAuthDataInterface) {
	v.value = val
	v.isSet = true
}

func (v NullablePayPalBraintreeDataAuthDataInterface) IsSet() bool {
	return v.isSet
}

func (v *NullablePayPalBraintreeDataAuthDataInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayPalBraintreeDataAuthDataInterface(val *PayPalBraintreeDataAuthDataInterface) *NullablePayPalBraintreeDataAuthDataInterface {
	return &NullablePayPalBraintreeDataAuthDataInterface{value: val, isSet: true}
}

func (v NullablePayPalBraintreeDataAuthDataInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayPalBraintreeDataAuthDataInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
