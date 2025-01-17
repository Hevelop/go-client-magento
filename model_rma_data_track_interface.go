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

// checks if the RmaDataTrackInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaDataTrackInterface{}

// RmaDataTrackInterface Interface TrackInterface
type RmaDataTrackInterface struct {
	// Entity id
	EntityId int32 `json:"entity_id"`
	// Rma entity id
	RmaEntityId int32 `json:"rma_entity_id"`
	// Track number
	TrackNumber string `json:"track_number"`
	// Carrier title
	CarrierTitle string `json:"carrier_title"`
	// Carrier code
	CarrierCode string `json:"carrier_code"`
	// ExtensionInterface class for @see \\Magento\\Rma\\Api\\Data\\TrackInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RmaDataTrackInterface RmaDataTrackInterface

// NewRmaDataTrackInterface instantiates a new RmaDataTrackInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaDataTrackInterface(entityId int32, rmaEntityId int32, trackNumber string, carrierTitle string, carrierCode string) *RmaDataTrackInterface {
	this := RmaDataTrackInterface{}
	this.EntityId = entityId
	this.RmaEntityId = rmaEntityId
	this.TrackNumber = trackNumber
	this.CarrierTitle = carrierTitle
	this.CarrierCode = carrierCode
	return &this
}

// NewRmaDataTrackInterfaceWithDefaults instantiates a new RmaDataTrackInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaDataTrackInterfaceWithDefaults() *RmaDataTrackInterface {
	this := RmaDataTrackInterface{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *RmaDataTrackInterface) GetEntityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *RmaDataTrackInterface) GetEntityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *RmaDataTrackInterface) SetEntityId(v int32) {
	o.EntityId = v
}

// GetRmaEntityId returns the RmaEntityId field value
func (o *RmaDataTrackInterface) GetRmaEntityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RmaEntityId
}

// GetRmaEntityIdOk returns a tuple with the RmaEntityId field value
// and a boolean to check if the value has been set.
func (o *RmaDataTrackInterface) GetRmaEntityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RmaEntityId, true
}

// SetRmaEntityId sets field value
func (o *RmaDataTrackInterface) SetRmaEntityId(v int32) {
	o.RmaEntityId = v
}

// GetTrackNumber returns the TrackNumber field value
func (o *RmaDataTrackInterface) GetTrackNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackNumber
}

// GetTrackNumberOk returns a tuple with the TrackNumber field value
// and a boolean to check if the value has been set.
func (o *RmaDataTrackInterface) GetTrackNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackNumber, true
}

// SetTrackNumber sets field value
func (o *RmaDataTrackInterface) SetTrackNumber(v string) {
	o.TrackNumber = v
}

// GetCarrierTitle returns the CarrierTitle field value
func (o *RmaDataTrackInterface) GetCarrierTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierTitle
}

// GetCarrierTitleOk returns a tuple with the CarrierTitle field value
// and a boolean to check if the value has been set.
func (o *RmaDataTrackInterface) GetCarrierTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierTitle, true
}

// SetCarrierTitle sets field value
func (o *RmaDataTrackInterface) SetCarrierTitle(v string) {
	o.CarrierTitle = v
}

// GetCarrierCode returns the CarrierCode field value
func (o *RmaDataTrackInterface) GetCarrierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value
// and a boolean to check if the value has been set.
func (o *RmaDataTrackInterface) GetCarrierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierCode, true
}

// SetCarrierCode sets field value
func (o *RmaDataTrackInterface) SetCarrierCode(v string) {
	o.CarrierCode = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *RmaDataTrackInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaDataTrackInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *RmaDataTrackInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *RmaDataTrackInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o RmaDataTrackInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaDataTrackInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["rma_entity_id"] = o.RmaEntityId
	toSerialize["track_number"] = o.TrackNumber
	toSerialize["carrier_title"] = o.CarrierTitle
	toSerialize["carrier_code"] = o.CarrierCode
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RmaDataTrackInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
		"rma_entity_id",
		"track_number",
		"carrier_title",
		"carrier_code",
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

	varRmaDataTrackInterface := _RmaDataTrackInterface{}

	err = json.Unmarshal(data, &varRmaDataTrackInterface)

	if err != nil {
		return err
	}

	*o = RmaDataTrackInterface(varRmaDataTrackInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "rma_entity_id")
		delete(additionalProperties, "track_number")
		delete(additionalProperties, "carrier_title")
		delete(additionalProperties, "carrier_code")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *RmaDataTrackInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *RmaDataTrackInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableRmaDataTrackInterface struct {
	value *RmaDataTrackInterface
	isSet bool
}

func (v NullableRmaDataTrackInterface) Get() *RmaDataTrackInterface {
	return v.value
}

func (v *NullableRmaDataTrackInterface) Set(val *RmaDataTrackInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaDataTrackInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaDataTrackInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaDataTrackInterface(val *RmaDataTrackInterface) *NullableRmaDataTrackInterface {
	return &NullableRmaDataTrackInterface{value: val, isSet: true}
}

func (v NullableRmaDataTrackInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaDataTrackInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
