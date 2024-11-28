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

// checks if the SalesRuleDataRuleExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesRuleDataRuleExtensionInterface{}

// SalesRuleDataRuleExtensionInterface ExtensionInterface class for @see \\Magento\\SalesRule\\Api\\Data\\RuleInterface
type SalesRuleDataRuleExtensionInterface struct {
	RewardPointsDelta    *int32 `json:"reward_points_delta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesRuleDataRuleExtensionInterface SalesRuleDataRuleExtensionInterface

// NewSalesRuleDataRuleExtensionInterface instantiates a new SalesRuleDataRuleExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesRuleDataRuleExtensionInterface() *SalesRuleDataRuleExtensionInterface {
	this := SalesRuleDataRuleExtensionInterface{}
	return &this
}

// NewSalesRuleDataRuleExtensionInterfaceWithDefaults instantiates a new SalesRuleDataRuleExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesRuleDataRuleExtensionInterfaceWithDefaults() *SalesRuleDataRuleExtensionInterface {
	this := SalesRuleDataRuleExtensionInterface{}
	return &this
}

// GetRewardPointsDelta returns the RewardPointsDelta field value if set, zero value otherwise.
func (o *SalesRuleDataRuleExtensionInterface) GetRewardPointsDelta() int32 {
	if o == nil || IsNil(o.RewardPointsDelta) {
		var ret int32
		return ret
	}
	return *o.RewardPointsDelta
}

// GetRewardPointsDeltaOk returns a tuple with the RewardPointsDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleExtensionInterface) GetRewardPointsDeltaOk() (*int32, bool) {
	if o == nil || IsNil(o.RewardPointsDelta) {
		return nil, false
	}
	return o.RewardPointsDelta, true
}

// HasRewardPointsDelta returns a boolean if a field has been set.
func (o *SalesRuleDataRuleExtensionInterface) HasRewardPointsDelta() bool {
	if o != nil && !IsNil(o.RewardPointsDelta) {
		return true
	}

	return false
}

// SetRewardPointsDelta gets a reference to the given int32 and assigns it to the RewardPointsDelta field.
func (o *SalesRuleDataRuleExtensionInterface) SetRewardPointsDelta(v int32) {
	o.RewardPointsDelta = &v
}

func (o SalesRuleDataRuleExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesRuleDataRuleExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RewardPointsDelta) {
		toSerialize["reward_points_delta"] = o.RewardPointsDelta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesRuleDataRuleExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varSalesRuleDataRuleExtensionInterface := _SalesRuleDataRuleExtensionInterface{}

	err = json.Unmarshal(data, &varSalesRuleDataRuleExtensionInterface)

	if err != nil {
		return err
	}

	*o = SalesRuleDataRuleExtensionInterface(varSalesRuleDataRuleExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reward_points_delta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesRuleDataRuleExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesRuleDataRuleExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesRuleDataRuleExtensionInterface struct {
	value *SalesRuleDataRuleExtensionInterface
	isSet bool
}

func (v NullableSalesRuleDataRuleExtensionInterface) Get() *SalesRuleDataRuleExtensionInterface {
	return v.value
}

func (v *NullableSalesRuleDataRuleExtensionInterface) Set(val *SalesRuleDataRuleExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesRuleDataRuleExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesRuleDataRuleExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesRuleDataRuleExtensionInterface(val *SalesRuleDataRuleExtensionInterface) *NullableSalesRuleDataRuleExtensionInterface {
	return &NullableSalesRuleDataRuleExtensionInterface{value: val, isSet: true}
}

func (v NullableSalesRuleDataRuleExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesRuleDataRuleExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}