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

// checks if the SalesDataShipmentCommentInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataShipmentCommentInterface{}

// SalesDataShipmentCommentInterface Shipment comment interface. A shipment is a delivery package that contains products. A shipment document accompanies the shipment. This document lists the products and their quantities in the delivery package. A shipment document can contain comments.
type SalesDataShipmentCommentInterface struct {
	// Is-customer-notified flag value.
	IsCustomerNotified int32 `json:"is_customer_notified"`
	// Parent ID.
	ParentId int32 `json:"parent_id"`
	// ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\ShipmentCommentInterface
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
	// Comment.
	Comment string `json:"comment"`
	// Is-visible-on-storefront flag value.
	IsVisibleOnFront int32 `json:"is_visible_on_front"`
	// Created-at timestamp.
	CreatedAt *string `json:"created_at,omitempty"`
	// Invoice ID.
	EntityId             *int32 `json:"entity_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataShipmentCommentInterface SalesDataShipmentCommentInterface

// NewSalesDataShipmentCommentInterface instantiates a new SalesDataShipmentCommentInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataShipmentCommentInterface(isCustomerNotified int32, parentId int32, comment string, isVisibleOnFront int32) *SalesDataShipmentCommentInterface {
	this := SalesDataShipmentCommentInterface{}
	this.IsCustomerNotified = isCustomerNotified
	this.ParentId = parentId
	this.Comment = comment
	this.IsVisibleOnFront = isVisibleOnFront
	return &this
}

// NewSalesDataShipmentCommentInterfaceWithDefaults instantiates a new SalesDataShipmentCommentInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataShipmentCommentInterfaceWithDefaults() *SalesDataShipmentCommentInterface {
	this := SalesDataShipmentCommentInterface{}
	return &this
}

// GetIsCustomerNotified returns the IsCustomerNotified field value
func (o *SalesDataShipmentCommentInterface) GetIsCustomerNotified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsCustomerNotified
}

// GetIsCustomerNotifiedOk returns a tuple with the IsCustomerNotified field value
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCommentInterface) GetIsCustomerNotifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCustomerNotified, true
}

// SetIsCustomerNotified sets field value
func (o *SalesDataShipmentCommentInterface) SetIsCustomerNotified(v int32) {
	o.IsCustomerNotified = v
}

// GetParentId returns the ParentId field value
func (o *SalesDataShipmentCommentInterface) GetParentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCommentInterface) GetParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *SalesDataShipmentCommentInterface) SetParentId(v int32) {
	o.ParentId = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataShipmentCommentInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCommentInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataShipmentCommentInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *SalesDataShipmentCommentInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

// GetComment returns the Comment field value
func (o *SalesDataShipmentCommentInterface) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCommentInterface) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *SalesDataShipmentCommentInterface) SetComment(v string) {
	o.Comment = v
}

// GetIsVisibleOnFront returns the IsVisibleOnFront field value
func (o *SalesDataShipmentCommentInterface) GetIsVisibleOnFront() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsVisibleOnFront
}

// GetIsVisibleOnFrontOk returns a tuple with the IsVisibleOnFront field value
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCommentInterface) GetIsVisibleOnFrontOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVisibleOnFront, true
}

// SetIsVisibleOnFront sets field value
func (o *SalesDataShipmentCommentInterface) SetIsVisibleOnFront(v int32) {
	o.IsVisibleOnFront = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SalesDataShipmentCommentInterface) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCommentInterface) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SalesDataShipmentCommentInterface) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SalesDataShipmentCommentInterface) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SalesDataShipmentCommentInterface) GetEntityId() int32 {
	if o == nil || IsNil(o.EntityId) {
		var ret int32
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataShipmentCommentInterface) GetEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SalesDataShipmentCommentInterface) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given int32 and assigns it to the EntityId field.
func (o *SalesDataShipmentCommentInterface) SetEntityId(v int32) {
	o.EntityId = &v
}

func (o SalesDataShipmentCommentInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataShipmentCommentInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_customer_notified"] = o.IsCustomerNotified
	toSerialize["parent_id"] = o.ParentId
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}
	toSerialize["comment"] = o.Comment
	toSerialize["is_visible_on_front"] = o.IsVisibleOnFront
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataShipmentCommentInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"is_customer_notified",
		"parent_id",
		"comment",
		"is_visible_on_front",
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

	varSalesDataShipmentCommentInterface := _SalesDataShipmentCommentInterface{}

	err = json.Unmarshal(data, &varSalesDataShipmentCommentInterface)

	if err != nil {
		return err
	}

	*o = SalesDataShipmentCommentInterface(varSalesDataShipmentCommentInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_customer_notified")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "extension_attributes")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "is_visible_on_front")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "entity_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataShipmentCommentInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataShipmentCommentInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataShipmentCommentInterface struct {
	value *SalesDataShipmentCommentInterface
	isSet bool
}

func (v NullableSalesDataShipmentCommentInterface) Get() *SalesDataShipmentCommentInterface {
	return v.value
}

func (v *NullableSalesDataShipmentCommentInterface) Set(val *SalesDataShipmentCommentInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataShipmentCommentInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataShipmentCommentInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataShipmentCommentInterface(val *SalesDataShipmentCommentInterface) *NullableSalesDataShipmentCommentInterface {
	return &NullableSalesDataShipmentCommentInterface{value: val, isSet: true}
}

func (v NullableSalesDataShipmentCommentInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataShipmentCommentInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
