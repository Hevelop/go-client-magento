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

// checks if the DownloadableDataSampleInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadableDataSampleInterface{}

// DownloadableDataSampleInterface
type DownloadableDataSampleInterface struct {
	// Sample(or link) id
	Id *int32 `json:"id,omitempty"`
	// Title
	Title string `json:"title"`
	// Order index for sample
	SortOrder  int32  `json:"sort_order"`
	SampleType string `json:"sample_type"`
	// relative file path
	SampleFile        *string                               `json:"sample_file,omitempty"`
	SampleFileContent *DownloadableDataFileContentInterface `json:"sample_file_content,omitempty"`
	// file URL
	SampleUrl *string `json:"sample_url,omitempty"`
	// ExtensionInterface class for @see \\Magento\\Downloadable\\Api\\Data\\SampleInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DownloadableDataSampleInterface DownloadableDataSampleInterface

// NewDownloadableDataSampleInterface instantiates a new DownloadableDataSampleInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadableDataSampleInterface(title string, sortOrder int32, sampleType string) *DownloadableDataSampleInterface {
	this := DownloadableDataSampleInterface{}
	this.Title = title
	this.SortOrder = sortOrder
	this.SampleType = sampleType
	return &this
}

// NewDownloadableDataSampleInterfaceWithDefaults instantiates a new DownloadableDataSampleInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadableDataSampleInterfaceWithDefaults() *DownloadableDataSampleInterface {
	this := DownloadableDataSampleInterface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DownloadableDataSampleInterface) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DownloadableDataSampleInterface) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DownloadableDataSampleInterface) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value
func (o *DownloadableDataSampleInterface) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DownloadableDataSampleInterface) SetTitle(v string) {
	o.Title = v
}

// GetSortOrder returns the SortOrder field value
func (o *DownloadableDataSampleInterface) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *DownloadableDataSampleInterface) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetSampleType returns the SampleType field value
func (o *DownloadableDataSampleInterface) GetSampleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SampleType
}

// GetSampleTypeOk returns a tuple with the SampleType field value
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetSampleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleType, true
}

// SetSampleType sets field value
func (o *DownloadableDataSampleInterface) SetSampleType(v string) {
	o.SampleType = v
}

// GetSampleFile returns the SampleFile field value if set, zero value otherwise.
func (o *DownloadableDataSampleInterface) GetSampleFile() string {
	if o == nil || IsNil(o.SampleFile) {
		var ret string
		return ret
	}
	return *o.SampleFile
}

// GetSampleFileOk returns a tuple with the SampleFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetSampleFileOk() (*string, bool) {
	if o == nil || IsNil(o.SampleFile) {
		return nil, false
	}
	return o.SampleFile, true
}

// HasSampleFile returns a boolean if a field has been set.
func (o *DownloadableDataSampleInterface) HasSampleFile() bool {
	if o != nil && !IsNil(o.SampleFile) {
		return true
	}

	return false
}

// SetSampleFile gets a reference to the given string and assigns it to the SampleFile field.
func (o *DownloadableDataSampleInterface) SetSampleFile(v string) {
	o.SampleFile = &v
}

// GetSampleFileContent returns the SampleFileContent field value if set, zero value otherwise.
func (o *DownloadableDataSampleInterface) GetSampleFileContent() DownloadableDataFileContentInterface {
	if o == nil || IsNil(o.SampleFileContent) {
		var ret DownloadableDataFileContentInterface
		return ret
	}
	return *o.SampleFileContent
}

// GetSampleFileContentOk returns a tuple with the SampleFileContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetSampleFileContentOk() (*DownloadableDataFileContentInterface, bool) {
	if o == nil || IsNil(o.SampleFileContent) {
		return nil, false
	}
	return o.SampleFileContent, true
}

// HasSampleFileContent returns a boolean if a field has been set.
func (o *DownloadableDataSampleInterface) HasSampleFileContent() bool {
	if o != nil && !IsNil(o.SampleFileContent) {
		return true
	}

	return false
}

// SetSampleFileContent gets a reference to the given DownloadableDataFileContentInterface and assigns it to the SampleFileContent field.
func (o *DownloadableDataSampleInterface) SetSampleFileContent(v DownloadableDataFileContentInterface) {
	o.SampleFileContent = &v
}

// GetSampleUrl returns the SampleUrl field value if set, zero value otherwise.
func (o *DownloadableDataSampleInterface) GetSampleUrl() string {
	if o == nil || IsNil(o.SampleUrl) {
		var ret string
		return ret
	}
	return *o.SampleUrl
}

// GetSampleUrlOk returns a tuple with the SampleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetSampleUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SampleUrl) {
		return nil, false
	}
	return o.SampleUrl, true
}

// HasSampleUrl returns a boolean if a field has been set.
func (o *DownloadableDataSampleInterface) HasSampleUrl() bool {
	if o != nil && !IsNil(o.SampleUrl) {
		return true
	}

	return false
}

// SetSampleUrl gets a reference to the given string and assigns it to the SampleUrl field.
func (o *DownloadableDataSampleInterface) SetSampleUrl(v string) {
	o.SampleUrl = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *DownloadableDataSampleInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadableDataSampleInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *DownloadableDataSampleInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *DownloadableDataSampleInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o DownloadableDataSampleInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadableDataSampleInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["title"] = o.Title
	toSerialize["sort_order"] = o.SortOrder
	toSerialize["sample_type"] = o.SampleType
	if !IsNil(o.SampleFile) {
		toSerialize["sample_file"] = o.SampleFile
	}
	if !IsNil(o.SampleFileContent) {
		toSerialize["sample_file_content"] = o.SampleFileContent
	}
	if !IsNil(o.SampleUrl) {
		toSerialize["sample_url"] = o.SampleUrl
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DownloadableDataSampleInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"sort_order",
		"sample_type",
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

	varDownloadableDataSampleInterface := _DownloadableDataSampleInterface{}

	err = json.Unmarshal(data, &varDownloadableDataSampleInterface)

	if err != nil {
		return err
	}

	*o = DownloadableDataSampleInterface(varDownloadableDataSampleInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "sort_order")
		delete(additionalProperties, "sample_type")
		delete(additionalProperties, "sample_file")
		delete(additionalProperties, "sample_file_content")
		delete(additionalProperties, "sample_url")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DownloadableDataSampleInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *DownloadableDataSampleInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableDownloadableDataSampleInterface struct {
	value *DownloadableDataSampleInterface
	isSet bool
}

func (v NullableDownloadableDataSampleInterface) Get() *DownloadableDataSampleInterface {
	return v.value
}

func (v *NullableDownloadableDataSampleInterface) Set(val *DownloadableDataSampleInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadableDataSampleInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadableDataSampleInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadableDataSampleInterface(val *DownloadableDataSampleInterface) *NullableDownloadableDataSampleInterface {
	return &NullableDownloadableDataSampleInterface{value: val, isSet: true}
}

func (v NullableDownloadableDataSampleInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadableDataSampleInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
