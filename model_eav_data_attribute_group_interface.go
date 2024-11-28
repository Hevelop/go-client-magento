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

// checks if the EavDataAttributeGroupInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EavDataAttributeGroupInterface{}

// EavDataAttributeGroupInterface Interface AttributeGroupInterface
type EavDataAttributeGroupInterface struct {
	// Id
	AttributeGroupId *string `json:"attribute_group_id,omitempty"`
	// Name
	AttributeGroupName *string `json:"attribute_group_name,omitempty"`
	// Attribute set id
	AttributeSetId       *int32                                   `json:"attribute_set_id,omitempty"`
	ExtensionAttributes  *EavDataAttributeGroupExtensionInterface `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EavDataAttributeGroupInterface EavDataAttributeGroupInterface

// NewEavDataAttributeGroupInterface instantiates a new EavDataAttributeGroupInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEavDataAttributeGroupInterface() *EavDataAttributeGroupInterface {
	this := EavDataAttributeGroupInterface{}
	return &this
}

// NewEavDataAttributeGroupInterfaceWithDefaults instantiates a new EavDataAttributeGroupInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEavDataAttributeGroupInterfaceWithDefaults() *EavDataAttributeGroupInterface {
	this := EavDataAttributeGroupInterface{}
	return &this
}

// GetAttributeGroupId returns the AttributeGroupId field value if set, zero value otherwise.
func (o *EavDataAttributeGroupInterface) GetAttributeGroupId() string {
	if o == nil || IsNil(o.AttributeGroupId) {
		var ret string
		return ret
	}
	return *o.AttributeGroupId
}

// GetAttributeGroupIdOk returns a tuple with the AttributeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EavDataAttributeGroupInterface) GetAttributeGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeGroupId) {
		return nil, false
	}
	return o.AttributeGroupId, true
}

// HasAttributeGroupId returns a boolean if a field has been set.
func (o *EavDataAttributeGroupInterface) HasAttributeGroupId() bool {
	if o != nil && !IsNil(o.AttributeGroupId) {
		return true
	}

	return false
}

// SetAttributeGroupId gets a reference to the given string and assigns it to the AttributeGroupId field.
func (o *EavDataAttributeGroupInterface) SetAttributeGroupId(v string) {
	o.AttributeGroupId = &v
}

// GetAttributeGroupName returns the AttributeGroupName field value if set, zero value otherwise.
func (o *EavDataAttributeGroupInterface) GetAttributeGroupName() string {
	if o == nil || IsNil(o.AttributeGroupName) {
		var ret string
		return ret
	}
	return *o.AttributeGroupName
}

// GetAttributeGroupNameOk returns a tuple with the AttributeGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EavDataAttributeGroupInterface) GetAttributeGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeGroupName) {
		return nil, false
	}
	return o.AttributeGroupName, true
}

// HasAttributeGroupName returns a boolean if a field has been set.
func (o *EavDataAttributeGroupInterface) HasAttributeGroupName() bool {
	if o != nil && !IsNil(o.AttributeGroupName) {
		return true
	}

	return false
}

// SetAttributeGroupName gets a reference to the given string and assigns it to the AttributeGroupName field.
func (o *EavDataAttributeGroupInterface) SetAttributeGroupName(v string) {
	o.AttributeGroupName = &v
}

// GetAttributeSetId returns the AttributeSetId field value if set, zero value otherwise.
func (o *EavDataAttributeGroupInterface) GetAttributeSetId() int32 {
	if o == nil || IsNil(o.AttributeSetId) {
		var ret int32
		return ret
	}
	return *o.AttributeSetId
}

// GetAttributeSetIdOk returns a tuple with the AttributeSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EavDataAttributeGroupInterface) GetAttributeSetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AttributeSetId) {
		return nil, false
	}
	return o.AttributeSetId, true
}

// HasAttributeSetId returns a boolean if a field has been set.
func (o *EavDataAttributeGroupInterface) HasAttributeSetId() bool {
	if o != nil && !IsNil(o.AttributeSetId) {
		return true
	}

	return false
}

// SetAttributeSetId gets a reference to the given int32 and assigns it to the AttributeSetId field.
func (o *EavDataAttributeGroupInterface) SetAttributeSetId(v int32) {
	o.AttributeSetId = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *EavDataAttributeGroupInterface) GetExtensionAttributes() EavDataAttributeGroupExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret EavDataAttributeGroupExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EavDataAttributeGroupInterface) GetExtensionAttributesOk() (*EavDataAttributeGroupExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *EavDataAttributeGroupInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given EavDataAttributeGroupExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *EavDataAttributeGroupInterface) SetExtensionAttributes(v EavDataAttributeGroupExtensionInterface) {
	o.ExtensionAttributes = &v
}

func (o EavDataAttributeGroupInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EavDataAttributeGroupInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeGroupId) {
		toSerialize["attribute_group_id"] = o.AttributeGroupId
	}
	if !IsNil(o.AttributeGroupName) {
		toSerialize["attribute_group_name"] = o.AttributeGroupName
	}
	if !IsNil(o.AttributeSetId) {
		toSerialize["attribute_set_id"] = o.AttributeSetId
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EavDataAttributeGroupInterface) UnmarshalJSON(data []byte) (err error) {
	varEavDataAttributeGroupInterface := _EavDataAttributeGroupInterface{}

	err = json.Unmarshal(data, &varEavDataAttributeGroupInterface)

	if err != nil {
		return err
	}

	*o = EavDataAttributeGroupInterface(varEavDataAttributeGroupInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute_group_id")
		delete(additionalProperties, "attribute_group_name")
		delete(additionalProperties, "attribute_set_id")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *EavDataAttributeGroupInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *EavDataAttributeGroupInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableEavDataAttributeGroupInterface struct {
	value *EavDataAttributeGroupInterface
	isSet bool
}

func (v NullableEavDataAttributeGroupInterface) Get() *EavDataAttributeGroupInterface {
	return v.value
}

func (v *NullableEavDataAttributeGroupInterface) Set(val *EavDataAttributeGroupInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableEavDataAttributeGroupInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableEavDataAttributeGroupInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEavDataAttributeGroupInterface(val *EavDataAttributeGroupInterface) *NullableEavDataAttributeGroupInterface {
	return &NullableEavDataAttributeGroupInterface{value: val, isSet: true}
}

func (v NullableEavDataAttributeGroupInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEavDataAttributeGroupInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}