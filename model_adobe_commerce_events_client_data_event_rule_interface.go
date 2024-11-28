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

// checks if the AdobeCommerceEventsClientDataEventRuleInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdobeCommerceEventsClientDataEventRuleInterface{}

// AdobeCommerceEventsClientDataEventRuleInterface Interface for event rule data from webapi requests
type AdobeCommerceEventsClientDataEventRuleInterface struct {
	// Event rule field name
	Field string `json:"field"`
	// Event rule operator
	Operator string `json:"operator"`
	// Event rule value
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _AdobeCommerceEventsClientDataEventRuleInterface AdobeCommerceEventsClientDataEventRuleInterface

// NewAdobeCommerceEventsClientDataEventRuleInterface instantiates a new AdobeCommerceEventsClientDataEventRuleInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdobeCommerceEventsClientDataEventRuleInterface(field string, operator string, value string) *AdobeCommerceEventsClientDataEventRuleInterface {
	this := AdobeCommerceEventsClientDataEventRuleInterface{}
	this.Field = field
	this.Operator = operator
	this.Value = value
	return &this
}

// NewAdobeCommerceEventsClientDataEventRuleInterfaceWithDefaults instantiates a new AdobeCommerceEventsClientDataEventRuleInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdobeCommerceEventsClientDataEventRuleInterfaceWithDefaults() *AdobeCommerceEventsClientDataEventRuleInterface {
	this := AdobeCommerceEventsClientDataEventRuleInterface{}
	return &this
}

// GetField returns the Field field value
func (o *AdobeCommerceEventsClientDataEventRuleInterface) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *AdobeCommerceEventsClientDataEventRuleInterface) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *AdobeCommerceEventsClientDataEventRuleInterface) SetField(v string) {
	o.Field = v
}

// GetOperator returns the Operator field value
func (o *AdobeCommerceEventsClientDataEventRuleInterface) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *AdobeCommerceEventsClientDataEventRuleInterface) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *AdobeCommerceEventsClientDataEventRuleInterface) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value
func (o *AdobeCommerceEventsClientDataEventRuleInterface) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AdobeCommerceEventsClientDataEventRuleInterface) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AdobeCommerceEventsClientDataEventRuleInterface) SetValue(v string) {
	o.Value = v
}

func (o AdobeCommerceEventsClientDataEventRuleInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdobeCommerceEventsClientDataEventRuleInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdobeCommerceEventsClientDataEventRuleInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
		"operator",
		"value",
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

	varAdobeCommerceEventsClientDataEventRuleInterface := _AdobeCommerceEventsClientDataEventRuleInterface{}

	err = json.Unmarshal(data, &varAdobeCommerceEventsClientDataEventRuleInterface)

	if err != nil {
		return err
	}

	*o = AdobeCommerceEventsClientDataEventRuleInterface(varAdobeCommerceEventsClientDataEventRuleInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AdobeCommerceEventsClientDataEventRuleInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AdobeCommerceEventsClientDataEventRuleInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAdobeCommerceEventsClientDataEventRuleInterface struct {
	value *AdobeCommerceEventsClientDataEventRuleInterface
	isSet bool
}

func (v NullableAdobeCommerceEventsClientDataEventRuleInterface) Get() *AdobeCommerceEventsClientDataEventRuleInterface {
	return v.value
}

func (v *NullableAdobeCommerceEventsClientDataEventRuleInterface) Set(val *AdobeCommerceEventsClientDataEventRuleInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableAdobeCommerceEventsClientDataEventRuleInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableAdobeCommerceEventsClientDataEventRuleInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdobeCommerceEventsClientDataEventRuleInterface(val *AdobeCommerceEventsClientDataEventRuleInterface) *NullableAdobeCommerceEventsClientDataEventRuleInterface {
	return &NullableAdobeCommerceEventsClientDataEventRuleInterface{value: val, isSet: true}
}

func (v NullableAdobeCommerceEventsClientDataEventRuleInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdobeCommerceEventsClientDataEventRuleInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
