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

// checks if the CompanyDataHierarchyInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyDataHierarchyInterface{}

// CompanyDataHierarchyInterface Company hierarchy DTO interface for WebAPI.
type CompanyDataHierarchyInterface struct {
	// Structure ID.
	StructureId *int32 `json:"structure_id,omitempty"`
	// Entity ID.
	EntityId *int32 `json:"entity_id,omitempty"`
	// Entity type.
	EntityType *string `json:"entity_type,omitempty"`
	// Structure parent ID.
	StructureParentId *int32 `json:"structure_parent_id,omitempty"`
	// ExtensionInterface class for @see \\Magento\\Company\\Api\\Data\\HierarchyInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompanyDataHierarchyInterface CompanyDataHierarchyInterface

// NewCompanyDataHierarchyInterface instantiates a new CompanyDataHierarchyInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyDataHierarchyInterface() *CompanyDataHierarchyInterface {
	this := CompanyDataHierarchyInterface{}
	return &this
}

// NewCompanyDataHierarchyInterfaceWithDefaults instantiates a new CompanyDataHierarchyInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyDataHierarchyInterfaceWithDefaults() *CompanyDataHierarchyInterface {
	this := CompanyDataHierarchyInterface{}
	return &this
}

// GetStructureId returns the StructureId field value if set, zero value otherwise.
func (o *CompanyDataHierarchyInterface) GetStructureId() int32 {
	if o == nil || IsNil(o.StructureId) {
		var ret int32
		return ret
	}
	return *o.StructureId
}

// GetStructureIdOk returns a tuple with the StructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataHierarchyInterface) GetStructureIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StructureId) {
		return nil, false
	}
	return o.StructureId, true
}

// HasStructureId returns a boolean if a field has been set.
func (o *CompanyDataHierarchyInterface) HasStructureId() bool {
	if o != nil && !IsNil(o.StructureId) {
		return true
	}

	return false
}

// SetStructureId gets a reference to the given int32 and assigns it to the StructureId field.
func (o *CompanyDataHierarchyInterface) SetStructureId(v int32) {
	o.StructureId = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *CompanyDataHierarchyInterface) GetEntityId() int32 {
	if o == nil || IsNil(o.EntityId) {
		var ret int32
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataHierarchyInterface) GetEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *CompanyDataHierarchyInterface) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given int32 and assigns it to the EntityId field.
func (o *CompanyDataHierarchyInterface) SetEntityId(v int32) {
	o.EntityId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CompanyDataHierarchyInterface) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataHierarchyInterface) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CompanyDataHierarchyInterface) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *CompanyDataHierarchyInterface) SetEntityType(v string) {
	o.EntityType = &v
}

// GetStructureParentId returns the StructureParentId field value if set, zero value otherwise.
func (o *CompanyDataHierarchyInterface) GetStructureParentId() int32 {
	if o == nil || IsNil(o.StructureParentId) {
		var ret int32
		return ret
	}
	return *o.StructureParentId
}

// GetStructureParentIdOk returns a tuple with the StructureParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataHierarchyInterface) GetStructureParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StructureParentId) {
		return nil, false
	}
	return o.StructureParentId, true
}

// HasStructureParentId returns a boolean if a field has been set.
func (o *CompanyDataHierarchyInterface) HasStructureParentId() bool {
	if o != nil && !IsNil(o.StructureParentId) {
		return true
	}

	return false
}

// SetStructureParentId gets a reference to the given int32 and assigns it to the StructureParentId field.
func (o *CompanyDataHierarchyInterface) SetStructureParentId(v int32) {
	o.StructureParentId = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *CompanyDataHierarchyInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyDataHierarchyInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *CompanyDataHierarchyInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *CompanyDataHierarchyInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o CompanyDataHierarchyInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyDataHierarchyInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StructureId) {
		toSerialize["structure_id"] = o.StructureId
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.StructureParentId) {
		toSerialize["structure_parent_id"] = o.StructureParentId
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompanyDataHierarchyInterface) UnmarshalJSON(data []byte) (err error) {
	varCompanyDataHierarchyInterface := _CompanyDataHierarchyInterface{}

	err = json.Unmarshal(data, &varCompanyDataHierarchyInterface)

	if err != nil {
		return err
	}

	*o = CompanyDataHierarchyInterface(varCompanyDataHierarchyInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "structure_id")
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "entity_type")
		delete(additionalProperties, "structure_parent_id")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CompanyDataHierarchyInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CompanyDataHierarchyInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCompanyDataHierarchyInterface struct {
	value *CompanyDataHierarchyInterface
	isSet bool
}

func (v NullableCompanyDataHierarchyInterface) Get() *CompanyDataHierarchyInterface {
	return v.value
}

func (v *NullableCompanyDataHierarchyInterface) Set(val *CompanyDataHierarchyInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyDataHierarchyInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyDataHierarchyInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyDataHierarchyInterface(val *CompanyDataHierarchyInterface) *NullableCompanyDataHierarchyInterface {
	return &NullableCompanyDataHierarchyInterface{value: val, isSet: true}
}

func (v NullableCompanyDataHierarchyInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyDataHierarchyInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
