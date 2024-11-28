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

// checks if the AnalyticsDataLinkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsDataLinkInterface{}

// AnalyticsDataLinkInterface Represents link with collected data and initialized vector for decryption.
type AnalyticsDataLinkInterface struct {
	Url                  string `json:"url"`
	InitializationVector string `json:"initialization_vector"`
	AdditionalProperties map[string]interface{}
}

type _AnalyticsDataLinkInterface AnalyticsDataLinkInterface

// NewAnalyticsDataLinkInterface instantiates a new AnalyticsDataLinkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsDataLinkInterface(url string, initializationVector string) *AnalyticsDataLinkInterface {
	this := AnalyticsDataLinkInterface{}
	this.Url = url
	this.InitializationVector = initializationVector
	return &this
}

// NewAnalyticsDataLinkInterfaceWithDefaults instantiates a new AnalyticsDataLinkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsDataLinkInterfaceWithDefaults() *AnalyticsDataLinkInterface {
	this := AnalyticsDataLinkInterface{}
	return &this
}

// GetUrl returns the Url field value
func (o *AnalyticsDataLinkInterface) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDataLinkInterface) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AnalyticsDataLinkInterface) SetUrl(v string) {
	o.Url = v
}

// GetInitializationVector returns the InitializationVector field value
func (o *AnalyticsDataLinkInterface) GetInitializationVector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitializationVector
}

// GetInitializationVectorOk returns a tuple with the InitializationVector field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDataLinkInterface) GetInitializationVectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitializationVector, true
}

// SetInitializationVector sets field value
func (o *AnalyticsDataLinkInterface) SetInitializationVector(v string) {
	o.InitializationVector = v
}

func (o AnalyticsDataLinkInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsDataLinkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["initialization_vector"] = o.InitializationVector

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnalyticsDataLinkInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"initialization_vector",
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

	varAnalyticsDataLinkInterface := _AnalyticsDataLinkInterface{}

	err = json.Unmarshal(data, &varAnalyticsDataLinkInterface)

	if err != nil {
		return err
	}

	*o = AnalyticsDataLinkInterface(varAnalyticsDataLinkInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "initialization_vector")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AnalyticsDataLinkInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AnalyticsDataLinkInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAnalyticsDataLinkInterface struct {
	value *AnalyticsDataLinkInterface
	isSet bool
}

func (v NullableAnalyticsDataLinkInterface) Get() *AnalyticsDataLinkInterface {
	return v.value
}

func (v *NullableAnalyticsDataLinkInterface) Set(val *AnalyticsDataLinkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsDataLinkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsDataLinkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsDataLinkInterface(val *AnalyticsDataLinkInterface) *NullableAnalyticsDataLinkInterface {
	return &NullableAnalyticsDataLinkInterface{value: val, isSet: true}
}

func (v NullableAnalyticsDataLinkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsDataLinkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
