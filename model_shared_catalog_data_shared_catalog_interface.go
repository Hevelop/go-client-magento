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

// checks if the SharedCatalogDataSharedCatalogInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedCatalogDataSharedCatalogInterface{}

// SharedCatalogDataSharedCatalogInterface SharedCatalogInterface interface.
type SharedCatalogDataSharedCatalogInterface struct {
	// ID.
	Id *int32 `json:"id,omitempty"`
	// Shared Catalog name.
	Name string `json:"name"`
	// Shared Catalog description.
	Description string `json:"description"`
	// Customer Group Id.
	CustomerGroupId int32 `json:"customer_group_id"`
	// Shared Catalog type.
	Type int32 `json:"type"`
	// Created time for Shared Catalog.
	CreatedAt string `json:"created_at"`
	// Admin id for Shared Catalog.
	CreatedBy int32 `json:"created_by"`
	// Store group id for Shared Catalog.
	StoreId int32 `json:"store_id"`
	// Tax class id.
	TaxClassId           int32 `json:"tax_class_id"`
	AdditionalProperties map[string]interface{}
}

type _SharedCatalogDataSharedCatalogInterface SharedCatalogDataSharedCatalogInterface

// NewSharedCatalogDataSharedCatalogInterface instantiates a new SharedCatalogDataSharedCatalogInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedCatalogDataSharedCatalogInterface(name string, description string, customerGroupId int32, type_ int32, createdAt string, createdBy int32, storeId int32, taxClassId int32) *SharedCatalogDataSharedCatalogInterface {
	this := SharedCatalogDataSharedCatalogInterface{}
	this.Name = name
	this.Description = description
	this.CustomerGroupId = customerGroupId
	this.Type = type_
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.StoreId = storeId
	this.TaxClassId = taxClassId
	return &this
}

// NewSharedCatalogDataSharedCatalogInterfaceWithDefaults instantiates a new SharedCatalogDataSharedCatalogInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedCatalogDataSharedCatalogInterfaceWithDefaults() *SharedCatalogDataSharedCatalogInterface {
	this := SharedCatalogDataSharedCatalogInterface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SharedCatalogDataSharedCatalogInterface) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SharedCatalogDataSharedCatalogInterface) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SharedCatalogDataSharedCatalogInterface) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SharedCatalogDataSharedCatalogInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *SharedCatalogDataSharedCatalogInterface) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetDescription(v string) {
	o.Description = v
}

// GetCustomerGroupId returns the CustomerGroupId field value
func (o *SharedCatalogDataSharedCatalogInterface) GetCustomerGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomerGroupId
}

// GetCustomerGroupIdOk returns a tuple with the CustomerGroupId field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetCustomerGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerGroupId, true
}

// SetCustomerGroupId sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetCustomerGroupId(v int32) {
	o.CustomerGroupId = v
}

// GetType returns the Type field value
func (o *SharedCatalogDataSharedCatalogInterface) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetType(v int32) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SharedCatalogDataSharedCatalogInterface) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SharedCatalogDataSharedCatalogInterface) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetCreatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetStoreId returns the StoreId field value
func (o *SharedCatalogDataSharedCatalogInterface) GetStoreId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetStoreIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreId, true
}

// SetStoreId sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetStoreId(v int32) {
	o.StoreId = v
}

// GetTaxClassId returns the TaxClassId field value
func (o *SharedCatalogDataSharedCatalogInterface) GetTaxClassId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TaxClassId
}

// GetTaxClassIdOk returns a tuple with the TaxClassId field value
// and a boolean to check if the value has been set.
func (o *SharedCatalogDataSharedCatalogInterface) GetTaxClassIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxClassId, true
}

// SetTaxClassId sets field value
func (o *SharedCatalogDataSharedCatalogInterface) SetTaxClassId(v int32) {
	o.TaxClassId = v
}

func (o SharedCatalogDataSharedCatalogInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedCatalogDataSharedCatalogInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["customer_group_id"] = o.CustomerGroupId
	toSerialize["type"] = o.Type
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["store_id"] = o.StoreId
	toSerialize["tax_class_id"] = o.TaxClassId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SharedCatalogDataSharedCatalogInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"customer_group_id",
		"type",
		"created_at",
		"created_by",
		"store_id",
		"tax_class_id",
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

	varSharedCatalogDataSharedCatalogInterface := _SharedCatalogDataSharedCatalogInterface{}

	err = json.Unmarshal(data, &varSharedCatalogDataSharedCatalogInterface)

	if err != nil {
		return err
	}

	*o = SharedCatalogDataSharedCatalogInterface(varSharedCatalogDataSharedCatalogInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "customer_group_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "store_id")
		delete(additionalProperties, "tax_class_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SharedCatalogDataSharedCatalogInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SharedCatalogDataSharedCatalogInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSharedCatalogDataSharedCatalogInterface struct {
	value *SharedCatalogDataSharedCatalogInterface
	isSet bool
}

func (v NullableSharedCatalogDataSharedCatalogInterface) Get() *SharedCatalogDataSharedCatalogInterface {
	return v.value
}

func (v *NullableSharedCatalogDataSharedCatalogInterface) Set(val *SharedCatalogDataSharedCatalogInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedCatalogDataSharedCatalogInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedCatalogDataSharedCatalogInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedCatalogDataSharedCatalogInterface(val *SharedCatalogDataSharedCatalogInterface) *NullableSharedCatalogDataSharedCatalogInterface {
	return &NullableSharedCatalogDataSharedCatalogInterface{value: val, isSet: true}
}

func (v NullableSharedCatalogDataSharedCatalogInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedCatalogDataSharedCatalogInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}