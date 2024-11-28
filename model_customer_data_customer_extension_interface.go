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

// checks if the CustomerDataCustomerExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDataCustomerExtensionInterface{}

// CustomerDataCustomerExtensionInterface ExtensionInterface class for @see \\Magento\\Customer\\Api\\Data\\CustomerInterface
type CustomerDataCustomerExtensionInterface struct {
	CompanyAttributes    *CompanyDataCompanyCustomerInterface `json:"company_attributes,omitempty"`
	IsSubscribed         *bool                                `json:"is_subscribed,omitempty"`
	AssistanceAllowed    *int32                               `json:"assistance_allowed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerDataCustomerExtensionInterface CustomerDataCustomerExtensionInterface

// NewCustomerDataCustomerExtensionInterface instantiates a new CustomerDataCustomerExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataCustomerExtensionInterface() *CustomerDataCustomerExtensionInterface {
	this := CustomerDataCustomerExtensionInterface{}
	return &this
}

// NewCustomerDataCustomerExtensionInterfaceWithDefaults instantiates a new CustomerDataCustomerExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataCustomerExtensionInterfaceWithDefaults() *CustomerDataCustomerExtensionInterface {
	this := CustomerDataCustomerExtensionInterface{}
	return &this
}

// GetCompanyAttributes returns the CompanyAttributes field value if set, zero value otherwise.
func (o *CustomerDataCustomerExtensionInterface) GetCompanyAttributes() CompanyDataCompanyCustomerInterface {
	if o == nil || IsNil(o.CompanyAttributes) {
		var ret CompanyDataCompanyCustomerInterface
		return ret
	}
	return *o.CompanyAttributes
}

// GetCompanyAttributesOk returns a tuple with the CompanyAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataCustomerExtensionInterface) GetCompanyAttributesOk() (*CompanyDataCompanyCustomerInterface, bool) {
	if o == nil || IsNil(o.CompanyAttributes) {
		return nil, false
	}
	return o.CompanyAttributes, true
}

// HasCompanyAttributes returns a boolean if a field has been set.
func (o *CustomerDataCustomerExtensionInterface) HasCompanyAttributes() bool {
	if o != nil && !IsNil(o.CompanyAttributes) {
		return true
	}

	return false
}

// SetCompanyAttributes gets a reference to the given CompanyDataCompanyCustomerInterface and assigns it to the CompanyAttributes field.
func (o *CustomerDataCustomerExtensionInterface) SetCompanyAttributes(v CompanyDataCompanyCustomerInterface) {
	o.CompanyAttributes = &v
}

// GetIsSubscribed returns the IsSubscribed field value if set, zero value otherwise.
func (o *CustomerDataCustomerExtensionInterface) GetIsSubscribed() bool {
	if o == nil || IsNil(o.IsSubscribed) {
		var ret bool
		return ret
	}
	return *o.IsSubscribed
}

// GetIsSubscribedOk returns a tuple with the IsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataCustomerExtensionInterface) GetIsSubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscribed) {
		return nil, false
	}
	return o.IsSubscribed, true
}

// HasIsSubscribed returns a boolean if a field has been set.
func (o *CustomerDataCustomerExtensionInterface) HasIsSubscribed() bool {
	if o != nil && !IsNil(o.IsSubscribed) {
		return true
	}

	return false
}

// SetIsSubscribed gets a reference to the given bool and assigns it to the IsSubscribed field.
func (o *CustomerDataCustomerExtensionInterface) SetIsSubscribed(v bool) {
	o.IsSubscribed = &v
}

// GetAssistanceAllowed returns the AssistanceAllowed field value if set, zero value otherwise.
func (o *CustomerDataCustomerExtensionInterface) GetAssistanceAllowed() int32 {
	if o == nil || IsNil(o.AssistanceAllowed) {
		var ret int32
		return ret
	}
	return *o.AssistanceAllowed
}

// GetAssistanceAllowedOk returns a tuple with the AssistanceAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataCustomerExtensionInterface) GetAssistanceAllowedOk() (*int32, bool) {
	if o == nil || IsNil(o.AssistanceAllowed) {
		return nil, false
	}
	return o.AssistanceAllowed, true
}

// HasAssistanceAllowed returns a boolean if a field has been set.
func (o *CustomerDataCustomerExtensionInterface) HasAssistanceAllowed() bool {
	if o != nil && !IsNil(o.AssistanceAllowed) {
		return true
	}

	return false
}

// SetAssistanceAllowed gets a reference to the given int32 and assigns it to the AssistanceAllowed field.
func (o *CustomerDataCustomerExtensionInterface) SetAssistanceAllowed(v int32) {
	o.AssistanceAllowed = &v
}

func (o CustomerDataCustomerExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDataCustomerExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyAttributes) {
		toSerialize["company_attributes"] = o.CompanyAttributes
	}
	if !IsNil(o.IsSubscribed) {
		toSerialize["is_subscribed"] = o.IsSubscribed
	}
	if !IsNil(o.AssistanceAllowed) {
		toSerialize["assistance_allowed"] = o.AssistanceAllowed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerDataCustomerExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varCustomerDataCustomerExtensionInterface := _CustomerDataCustomerExtensionInterface{}

	err = json.Unmarshal(data, &varCustomerDataCustomerExtensionInterface)

	if err != nil {
		return err
	}

	*o = CustomerDataCustomerExtensionInterface(varCustomerDataCustomerExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company_attributes")
		delete(additionalProperties, "is_subscribed")
		delete(additionalProperties, "assistance_allowed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerDataCustomerExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CustomerDataCustomerExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCustomerDataCustomerExtensionInterface struct {
	value *CustomerDataCustomerExtensionInterface
	isSet bool
}

func (v NullableCustomerDataCustomerExtensionInterface) Get() *CustomerDataCustomerExtensionInterface {
	return v.value
}

func (v *NullableCustomerDataCustomerExtensionInterface) Set(val *CustomerDataCustomerExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataCustomerExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataCustomerExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataCustomerExtensionInterface(val *CustomerDataCustomerExtensionInterface) *NullableCustomerDataCustomerExtensionInterface {
	return &NullableCustomerDataCustomerExtensionInterface{value: val, isSet: true}
}

func (v NullableCustomerDataCustomerExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataCustomerExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
