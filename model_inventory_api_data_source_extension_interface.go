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

// checks if the InventoryApiDataSourceExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryApiDataSourceExtensionInterface{}

// InventoryApiDataSourceExtensionInterface ExtensionInterface class for @see \\Magento\\InventoryApi\\Api\\Data\\SourceInterface
type InventoryApiDataSourceExtensionInterface struct {
	IsPickupLocationActive *bool   `json:"is_pickup_location_active,omitempty"`
	FrontendName           *string `json:"frontend_name,omitempty"`
	FrontendDescription    *string `json:"frontend_description,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _InventoryApiDataSourceExtensionInterface InventoryApiDataSourceExtensionInterface

// NewInventoryApiDataSourceExtensionInterface instantiates a new InventoryApiDataSourceExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryApiDataSourceExtensionInterface() *InventoryApiDataSourceExtensionInterface {
	this := InventoryApiDataSourceExtensionInterface{}
	return &this
}

// NewInventoryApiDataSourceExtensionInterfaceWithDefaults instantiates a new InventoryApiDataSourceExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryApiDataSourceExtensionInterfaceWithDefaults() *InventoryApiDataSourceExtensionInterface {
	this := InventoryApiDataSourceExtensionInterface{}
	return &this
}

// GetIsPickupLocationActive returns the IsPickupLocationActive field value if set, zero value otherwise.
func (o *InventoryApiDataSourceExtensionInterface) GetIsPickupLocationActive() bool {
	if o == nil || IsNil(o.IsPickupLocationActive) {
		var ret bool
		return ret
	}
	return *o.IsPickupLocationActive
}

// GetIsPickupLocationActiveOk returns a tuple with the IsPickupLocationActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceExtensionInterface) GetIsPickupLocationActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPickupLocationActive) {
		return nil, false
	}
	return o.IsPickupLocationActive, true
}

// HasIsPickupLocationActive returns a boolean if a field has been set.
func (o *InventoryApiDataSourceExtensionInterface) HasIsPickupLocationActive() bool {
	if o != nil && !IsNil(o.IsPickupLocationActive) {
		return true
	}

	return false
}

// SetIsPickupLocationActive gets a reference to the given bool and assigns it to the IsPickupLocationActive field.
func (o *InventoryApiDataSourceExtensionInterface) SetIsPickupLocationActive(v bool) {
	o.IsPickupLocationActive = &v
}

// GetFrontendName returns the FrontendName field value if set, zero value otherwise.
func (o *InventoryApiDataSourceExtensionInterface) GetFrontendName() string {
	if o == nil || IsNil(o.FrontendName) {
		var ret string
		return ret
	}
	return *o.FrontendName
}

// GetFrontendNameOk returns a tuple with the FrontendName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceExtensionInterface) GetFrontendNameOk() (*string, bool) {
	if o == nil || IsNil(o.FrontendName) {
		return nil, false
	}
	return o.FrontendName, true
}

// HasFrontendName returns a boolean if a field has been set.
func (o *InventoryApiDataSourceExtensionInterface) HasFrontendName() bool {
	if o != nil && !IsNil(o.FrontendName) {
		return true
	}

	return false
}

// SetFrontendName gets a reference to the given string and assigns it to the FrontendName field.
func (o *InventoryApiDataSourceExtensionInterface) SetFrontendName(v string) {
	o.FrontendName = &v
}

// GetFrontendDescription returns the FrontendDescription field value if set, zero value otherwise.
func (o *InventoryApiDataSourceExtensionInterface) GetFrontendDescription() string {
	if o == nil || IsNil(o.FrontendDescription) {
		var ret string
		return ret
	}
	return *o.FrontendDescription
}

// GetFrontendDescriptionOk returns a tuple with the FrontendDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryApiDataSourceExtensionInterface) GetFrontendDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.FrontendDescription) {
		return nil, false
	}
	return o.FrontendDescription, true
}

// HasFrontendDescription returns a boolean if a field has been set.
func (o *InventoryApiDataSourceExtensionInterface) HasFrontendDescription() bool {
	if o != nil && !IsNil(o.FrontendDescription) {
		return true
	}

	return false
}

// SetFrontendDescription gets a reference to the given string and assigns it to the FrontendDescription field.
func (o *InventoryApiDataSourceExtensionInterface) SetFrontendDescription(v string) {
	o.FrontendDescription = &v
}

func (o InventoryApiDataSourceExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryApiDataSourceExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsPickupLocationActive) {
		toSerialize["is_pickup_location_active"] = o.IsPickupLocationActive
	}
	if !IsNil(o.FrontendName) {
		toSerialize["frontend_name"] = o.FrontendName
	}
	if !IsNil(o.FrontendDescription) {
		toSerialize["frontend_description"] = o.FrontendDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryApiDataSourceExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varInventoryApiDataSourceExtensionInterface := _InventoryApiDataSourceExtensionInterface{}

	err = json.Unmarshal(data, &varInventoryApiDataSourceExtensionInterface)

	if err != nil {
		return err
	}

	*o = InventoryApiDataSourceExtensionInterface(varInventoryApiDataSourceExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_pickup_location_active")
		delete(additionalProperties, "frontend_name")
		delete(additionalProperties, "frontend_description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *InventoryApiDataSourceExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *InventoryApiDataSourceExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableInventoryApiDataSourceExtensionInterface struct {
	value *InventoryApiDataSourceExtensionInterface
	isSet bool
}

func (v NullableInventoryApiDataSourceExtensionInterface) Get() *InventoryApiDataSourceExtensionInterface {
	return v.value
}

func (v *NullableInventoryApiDataSourceExtensionInterface) Set(val *InventoryApiDataSourceExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryApiDataSourceExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryApiDataSourceExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryApiDataSourceExtensionInterface(val *InventoryApiDataSourceExtensionInterface) *NullableInventoryApiDataSourceExtensionInterface {
	return &NullableInventoryApiDataSourceExtensionInterface{value: val, isSet: true}
}

func (v NullableInventoryApiDataSourceExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryApiDataSourceExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
