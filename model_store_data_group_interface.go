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

// checks if the StoreDataGroupInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreDataGroupInterface{}

// StoreDataGroupInterface Group interface
type StoreDataGroupInterface struct {
	Id             int32  `json:"id"`
	WebsiteId      int32  `json:"website_id"`
	RootCategoryId int32  `json:"root_category_id"`
	DefaultStoreId int32  `json:"default_store_id"`
	Name           string `json:"name"`
	// Group code.
	Code string `json:"code"`
	// ExtensionInterface class for @see \\Magento\\Store\\Api\\Data\\GroupInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreDataGroupInterface StoreDataGroupInterface

// NewStoreDataGroupInterface instantiates a new StoreDataGroupInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDataGroupInterface(id int32, websiteId int32, rootCategoryId int32, defaultStoreId int32, name string, code string) *StoreDataGroupInterface {
	this := StoreDataGroupInterface{}
	this.Id = id
	this.WebsiteId = websiteId
	this.RootCategoryId = rootCategoryId
	this.DefaultStoreId = defaultStoreId
	this.Name = name
	this.Code = code
	return &this
}

// NewStoreDataGroupInterfaceWithDefaults instantiates a new StoreDataGroupInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDataGroupInterfaceWithDefaults() *StoreDataGroupInterface {
	this := StoreDataGroupInterface{}
	return &this
}

// GetId returns the Id field value
func (o *StoreDataGroupInterface) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StoreDataGroupInterface) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StoreDataGroupInterface) SetId(v int32) {
	o.Id = v
}

// GetWebsiteId returns the WebsiteId field value
func (o *StoreDataGroupInterface) GetWebsiteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WebsiteId
}

// GetWebsiteIdOk returns a tuple with the WebsiteId field value
// and a boolean to check if the value has been set.
func (o *StoreDataGroupInterface) GetWebsiteIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsiteId, true
}

// SetWebsiteId sets field value
func (o *StoreDataGroupInterface) SetWebsiteId(v int32) {
	o.WebsiteId = v
}

// GetRootCategoryId returns the RootCategoryId field value
func (o *StoreDataGroupInterface) GetRootCategoryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RootCategoryId
}

// GetRootCategoryIdOk returns a tuple with the RootCategoryId field value
// and a boolean to check if the value has been set.
func (o *StoreDataGroupInterface) GetRootCategoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootCategoryId, true
}

// SetRootCategoryId sets field value
func (o *StoreDataGroupInterface) SetRootCategoryId(v int32) {
	o.RootCategoryId = v
}

// GetDefaultStoreId returns the DefaultStoreId field value
func (o *StoreDataGroupInterface) GetDefaultStoreId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultStoreId
}

// GetDefaultStoreIdOk returns a tuple with the DefaultStoreId field value
// and a boolean to check if the value has been set.
func (o *StoreDataGroupInterface) GetDefaultStoreIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultStoreId, true
}

// SetDefaultStoreId sets field value
func (o *StoreDataGroupInterface) SetDefaultStoreId(v int32) {
	o.DefaultStoreId = v
}

// GetName returns the Name field value
func (o *StoreDataGroupInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StoreDataGroupInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StoreDataGroupInterface) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *StoreDataGroupInterface) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *StoreDataGroupInterface) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *StoreDataGroupInterface) SetCode(v string) {
	o.Code = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *StoreDataGroupInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDataGroupInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *StoreDataGroupInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *StoreDataGroupInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o StoreDataGroupInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDataGroupInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["website_id"] = o.WebsiteId
	toSerialize["root_category_id"] = o.RootCategoryId
	toSerialize["default_store_id"] = o.DefaultStoreId
	toSerialize["name"] = o.Name
	toSerialize["code"] = o.Code
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoreDataGroupInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"website_id",
		"root_category_id",
		"default_store_id",
		"name",
		"code",
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

	varStoreDataGroupInterface := _StoreDataGroupInterface{}

	err = json.Unmarshal(data, &varStoreDataGroupInterface)

	if err != nil {
		return err
	}

	*o = StoreDataGroupInterface(varStoreDataGroupInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "website_id")
		delete(additionalProperties, "root_category_id")
		delete(additionalProperties, "default_store_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "code")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *StoreDataGroupInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *StoreDataGroupInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableStoreDataGroupInterface struct {
	value *StoreDataGroupInterface
	isSet bool
}

func (v NullableStoreDataGroupInterface) Get() *StoreDataGroupInterface {
	return v.value
}

func (v *NullableStoreDataGroupInterface) Set(val *StoreDataGroupInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDataGroupInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDataGroupInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDataGroupInterface(val *StoreDataGroupInterface) *NullableStoreDataGroupInterface {
	return &NullableStoreDataGroupInterface{value: val, isSet: true}
}

func (v NullableStoreDataGroupInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDataGroupInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
