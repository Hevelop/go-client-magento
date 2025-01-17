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

// checks if the TaxDataAppliedTaxRateInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDataAppliedTaxRateInterface{}

// TaxDataAppliedTaxRateInterface Applied tax rate interface.
type TaxDataAppliedTaxRateInterface struct {
	// Code
	Code *string `json:"code,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// Tax Percent
	Percent *float32 `json:"percent,omitempty"`
	// ExtensionInterface class for @see \\Magento\\Tax\\Api\\Data\\AppliedTaxRateInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaxDataAppliedTaxRateInterface TaxDataAppliedTaxRateInterface

// NewTaxDataAppliedTaxRateInterface instantiates a new TaxDataAppliedTaxRateInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDataAppliedTaxRateInterface() *TaxDataAppliedTaxRateInterface {
	this := TaxDataAppliedTaxRateInterface{}
	return &this
}

// NewTaxDataAppliedTaxRateInterfaceWithDefaults instantiates a new TaxDataAppliedTaxRateInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDataAppliedTaxRateInterfaceWithDefaults() *TaxDataAppliedTaxRateInterface {
	this := TaxDataAppliedTaxRateInterface{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TaxDataAppliedTaxRateInterface) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDataAppliedTaxRateInterface) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TaxDataAppliedTaxRateInterface) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TaxDataAppliedTaxRateInterface) SetCode(v string) {
	o.Code = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TaxDataAppliedTaxRateInterface) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDataAppliedTaxRateInterface) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TaxDataAppliedTaxRateInterface) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TaxDataAppliedTaxRateInterface) SetTitle(v string) {
	o.Title = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *TaxDataAppliedTaxRateInterface) GetPercent() float32 {
	if o == nil || IsNil(o.Percent) {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDataAppliedTaxRateInterface) GetPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *TaxDataAppliedTaxRateInterface) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *TaxDataAppliedTaxRateInterface) SetPercent(v float32) {
	o.Percent = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *TaxDataAppliedTaxRateInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDataAppliedTaxRateInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *TaxDataAppliedTaxRateInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *TaxDataAppliedTaxRateInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o TaxDataAppliedTaxRateInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxDataAppliedTaxRateInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaxDataAppliedTaxRateInterface) UnmarshalJSON(data []byte) (err error) {
	varTaxDataAppliedTaxRateInterface := _TaxDataAppliedTaxRateInterface{}

	err = json.Unmarshal(data, &varTaxDataAppliedTaxRateInterface)

	if err != nil {
		return err
	}

	*o = TaxDataAppliedTaxRateInterface(varTaxDataAppliedTaxRateInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "title")
		delete(additionalProperties, "percent")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *TaxDataAppliedTaxRateInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *TaxDataAppliedTaxRateInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableTaxDataAppliedTaxRateInterface struct {
	value *TaxDataAppliedTaxRateInterface
	isSet bool
}

func (v NullableTaxDataAppliedTaxRateInterface) Get() *TaxDataAppliedTaxRateInterface {
	return v.value
}

func (v *NullableTaxDataAppliedTaxRateInterface) Set(val *TaxDataAppliedTaxRateInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDataAppliedTaxRateInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDataAppliedTaxRateInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDataAppliedTaxRateInterface(val *TaxDataAppliedTaxRateInterface) *NullableTaxDataAppliedTaxRateInterface {
	return &NullableTaxDataAppliedTaxRateInterface{value: val, isSet: true}
}

func (v NullableTaxDataAppliedTaxRateInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxDataAppliedTaxRateInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
