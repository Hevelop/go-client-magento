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

// checks if the CheckoutAgreementsDataAgreementInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutAgreementsDataAgreementInterface{}

// CheckoutAgreementsDataAgreementInterface Interface AgreementInterface
type CheckoutAgreementsDataAgreementInterface struct {
	// Agreement ID.
	AgreementId int32 `json:"agreement_id"`
	// Agreement name.
	Name string `json:"name"`
	// Agreement content.
	Content string `json:"content"`
	// Agreement content height. Otherwise, null.
	ContentHeight *string `json:"content_height,omitempty"`
	// Agreement checkbox text.
	CheckboxText string `json:"checkbox_text"`
	// Agreement status.
	IsActive bool `json:"is_active"`
	// * true - HTML. * false - plain text.
	IsHtml bool `json:"is_html"`
	// The agreement applied mode.
	Mode int32 `json:"mode"`
	// ExtensionInterface class for @see \\Magento\\CheckoutAgreements\\Api\\Data\\AgreementInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckoutAgreementsDataAgreementInterface CheckoutAgreementsDataAgreementInterface

// NewCheckoutAgreementsDataAgreementInterface instantiates a new CheckoutAgreementsDataAgreementInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutAgreementsDataAgreementInterface(agreementId int32, name string, content string, checkboxText string, isActive bool, isHtml bool, mode int32) *CheckoutAgreementsDataAgreementInterface {
	this := CheckoutAgreementsDataAgreementInterface{}
	this.AgreementId = agreementId
	this.Name = name
	this.Content = content
	this.CheckboxText = checkboxText
	this.IsActive = isActive
	this.IsHtml = isHtml
	this.Mode = mode
	return &this
}

// NewCheckoutAgreementsDataAgreementInterfaceWithDefaults instantiates a new CheckoutAgreementsDataAgreementInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutAgreementsDataAgreementInterfaceWithDefaults() *CheckoutAgreementsDataAgreementInterface {
	this := CheckoutAgreementsDataAgreementInterface{}
	return &this
}

// GetAgreementId returns the AgreementId field value
func (o *CheckoutAgreementsDataAgreementInterface) GetAgreementId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetAgreementIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgreementId, true
}

// SetAgreementId sets field value
func (o *CheckoutAgreementsDataAgreementInterface) SetAgreementId(v int32) {
	o.AgreementId = v
}

// GetName returns the Name field value
func (o *CheckoutAgreementsDataAgreementInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CheckoutAgreementsDataAgreementInterface) SetName(v string) {
	o.Name = v
}

// GetContent returns the Content field value
func (o *CheckoutAgreementsDataAgreementInterface) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CheckoutAgreementsDataAgreementInterface) SetContent(v string) {
	o.Content = v
}

// GetContentHeight returns the ContentHeight field value if set, zero value otherwise.
func (o *CheckoutAgreementsDataAgreementInterface) GetContentHeight() string {
	if o == nil || IsNil(o.ContentHeight) {
		var ret string
		return ret
	}
	return *o.ContentHeight
}

// GetContentHeightOk returns a tuple with the ContentHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetContentHeightOk() (*string, bool) {
	if o == nil || IsNil(o.ContentHeight) {
		return nil, false
	}
	return o.ContentHeight, true
}

// HasContentHeight returns a boolean if a field has been set.
func (o *CheckoutAgreementsDataAgreementInterface) HasContentHeight() bool {
	if o != nil && !IsNil(o.ContentHeight) {
		return true
	}

	return false
}

// SetContentHeight gets a reference to the given string and assigns it to the ContentHeight field.
func (o *CheckoutAgreementsDataAgreementInterface) SetContentHeight(v string) {
	o.ContentHeight = &v
}

// GetCheckboxText returns the CheckboxText field value
func (o *CheckoutAgreementsDataAgreementInterface) GetCheckboxText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckboxText
}

// GetCheckboxTextOk returns a tuple with the CheckboxText field value
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetCheckboxTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckboxText, true
}

// SetCheckboxText sets field value
func (o *CheckoutAgreementsDataAgreementInterface) SetCheckboxText(v string) {
	o.CheckboxText = v
}

// GetIsActive returns the IsActive field value
func (o *CheckoutAgreementsDataAgreementInterface) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *CheckoutAgreementsDataAgreementInterface) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsHtml returns the IsHtml field value
func (o *CheckoutAgreementsDataAgreementInterface) GetIsHtml() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHtml
}

// GetIsHtmlOk returns a tuple with the IsHtml field value
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetIsHtmlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHtml, true
}

// SetIsHtml sets field value
func (o *CheckoutAgreementsDataAgreementInterface) SetIsHtml(v bool) {
	o.IsHtml = v
}

// GetMode returns the Mode field value
func (o *CheckoutAgreementsDataAgreementInterface) GetMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *CheckoutAgreementsDataAgreementInterface) SetMode(v int32) {
	o.Mode = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *CheckoutAgreementsDataAgreementInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutAgreementsDataAgreementInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *CheckoutAgreementsDataAgreementInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *CheckoutAgreementsDataAgreementInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o CheckoutAgreementsDataAgreementInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutAgreementsDataAgreementInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agreement_id"] = o.AgreementId
	toSerialize["name"] = o.Name
	toSerialize["content"] = o.Content
	if !IsNil(o.ContentHeight) {
		toSerialize["content_height"] = o.ContentHeight
	}
	toSerialize["checkbox_text"] = o.CheckboxText
	toSerialize["is_active"] = o.IsActive
	toSerialize["is_html"] = o.IsHtml
	toSerialize["mode"] = o.Mode
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckoutAgreementsDataAgreementInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agreement_id",
		"name",
		"content",
		"checkbox_text",
		"is_active",
		"is_html",
		"mode",
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

	varCheckoutAgreementsDataAgreementInterface := _CheckoutAgreementsDataAgreementInterface{}

	err = json.Unmarshal(data, &varCheckoutAgreementsDataAgreementInterface)

	if err != nil {
		return err
	}

	*o = CheckoutAgreementsDataAgreementInterface(varCheckoutAgreementsDataAgreementInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agreement_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "content")
		delete(additionalProperties, "content_height")
		delete(additionalProperties, "checkbox_text")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "is_html")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CheckoutAgreementsDataAgreementInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CheckoutAgreementsDataAgreementInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCheckoutAgreementsDataAgreementInterface struct {
	value *CheckoutAgreementsDataAgreementInterface
	isSet bool
}

func (v NullableCheckoutAgreementsDataAgreementInterface) Get() *CheckoutAgreementsDataAgreementInterface {
	return v.value
}

func (v *NullableCheckoutAgreementsDataAgreementInterface) Set(val *CheckoutAgreementsDataAgreementInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutAgreementsDataAgreementInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutAgreementsDataAgreementInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutAgreementsDataAgreementInterface(val *CheckoutAgreementsDataAgreementInterface) *NullableCheckoutAgreementsDataAgreementInterface {
	return &NullableCheckoutAgreementsDataAgreementInterface{value: val, isSet: true}
}

func (v NullableCheckoutAgreementsDataAgreementInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutAgreementsDataAgreementInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}