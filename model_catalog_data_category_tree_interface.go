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

// checks if the CatalogDataCategoryTreeInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDataCategoryTreeInterface{}

// CatalogDataCategoryTreeInterface
type CatalogDataCategoryTreeInterface struct {
	Id *int32 `json:"id,omitempty"`
	// Parent category ID
	ParentId int32 `json:"parent_id"`
	// Category name
	Name string `json:"name"`
	// Whether category is active
	IsActive bool `json:"is_active"`
	// Category position
	Position int32 `json:"position"`
	// Category level
	Level int32 `json:"level"`
	// Product count
	ProductCount         int32                              `json:"product_count"`
	ChildrenData         []CatalogDataCategoryTreeInterface `json:"children_data"`
	AdditionalProperties map[string]interface{}
}

type _CatalogDataCategoryTreeInterface CatalogDataCategoryTreeInterface

// NewCatalogDataCategoryTreeInterface instantiates a new CatalogDataCategoryTreeInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDataCategoryTreeInterface(parentId int32, name string, isActive bool, position int32, level int32, productCount int32, childrenData []CatalogDataCategoryTreeInterface) *CatalogDataCategoryTreeInterface {
	this := CatalogDataCategoryTreeInterface{}
	this.ParentId = parentId
	this.Name = name
	this.IsActive = isActive
	this.Position = position
	this.Level = level
	this.ProductCount = productCount
	this.ChildrenData = childrenData
	return &this
}

// NewCatalogDataCategoryTreeInterfaceWithDefaults instantiates a new CatalogDataCategoryTreeInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDataCategoryTreeInterfaceWithDefaults() *CatalogDataCategoryTreeInterface {
	this := CatalogDataCategoryTreeInterface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogDataCategoryTreeInterface) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogDataCategoryTreeInterface) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CatalogDataCategoryTreeInterface) SetId(v int32) {
	o.Id = &v
}

// GetParentId returns the ParentId field value
func (o *CatalogDataCategoryTreeInterface) GetParentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *CatalogDataCategoryTreeInterface) SetParentId(v int32) {
	o.ParentId = v
}

// GetName returns the Name field value
func (o *CatalogDataCategoryTreeInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CatalogDataCategoryTreeInterface) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value
func (o *CatalogDataCategoryTreeInterface) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *CatalogDataCategoryTreeInterface) SetIsActive(v bool) {
	o.IsActive = v
}

// GetPosition returns the Position field value
func (o *CatalogDataCategoryTreeInterface) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *CatalogDataCategoryTreeInterface) SetPosition(v int32) {
	o.Position = v
}

// GetLevel returns the Level field value
func (o *CatalogDataCategoryTreeInterface) GetLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CatalogDataCategoryTreeInterface) SetLevel(v int32) {
	o.Level = v
}

// GetProductCount returns the ProductCount field value
func (o *CatalogDataCategoryTreeInterface) GetProductCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductCount
}

// GetProductCountOk returns a tuple with the ProductCount field value
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetProductCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCount, true
}

// SetProductCount sets field value
func (o *CatalogDataCategoryTreeInterface) SetProductCount(v int32) {
	o.ProductCount = v
}

// GetChildrenData returns the ChildrenData field value
func (o *CatalogDataCategoryTreeInterface) GetChildrenData() []CatalogDataCategoryTreeInterface {
	if o == nil {
		var ret []CatalogDataCategoryTreeInterface
		return ret
	}

	return o.ChildrenData
}

// GetChildrenDataOk returns a tuple with the ChildrenData field value
// and a boolean to check if the value has been set.
func (o *CatalogDataCategoryTreeInterface) GetChildrenDataOk() ([]CatalogDataCategoryTreeInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChildrenData, true
}

// SetChildrenData sets field value
func (o *CatalogDataCategoryTreeInterface) SetChildrenData(v []CatalogDataCategoryTreeInterface) {
	o.ChildrenData = v
}

func (o CatalogDataCategoryTreeInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDataCategoryTreeInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["parent_id"] = o.ParentId
	toSerialize["name"] = o.Name
	toSerialize["is_active"] = o.IsActive
	toSerialize["position"] = o.Position
	toSerialize["level"] = o.Level
	toSerialize["product_count"] = o.ProductCount
	toSerialize["children_data"] = o.ChildrenData

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogDataCategoryTreeInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parent_id",
		"name",
		"is_active",
		"position",
		"level",
		"product_count",
		"children_data",
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

	varCatalogDataCategoryTreeInterface := _CatalogDataCategoryTreeInterface{}

	err = json.Unmarshal(data, &varCatalogDataCategoryTreeInterface)

	if err != nil {
		return err
	}

	*o = CatalogDataCategoryTreeInterface(varCatalogDataCategoryTreeInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "position")
		delete(additionalProperties, "level")
		delete(additionalProperties, "product_count")
		delete(additionalProperties, "children_data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogDataCategoryTreeInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogDataCategoryTreeInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogDataCategoryTreeInterface struct {
	value *CatalogDataCategoryTreeInterface
	isSet bool
}

func (v NullableCatalogDataCategoryTreeInterface) Get() *CatalogDataCategoryTreeInterface {
	return v.value
}

func (v *NullableCatalogDataCategoryTreeInterface) Set(val *CatalogDataCategoryTreeInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDataCategoryTreeInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDataCategoryTreeInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDataCategoryTreeInterface(val *CatalogDataCategoryTreeInterface) *NullableCatalogDataCategoryTreeInterface {
	return &NullableCatalogDataCategoryTreeInterface{value: val, isSet: true}
}

func (v NullableCatalogDataCategoryTreeInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDataCategoryTreeInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
