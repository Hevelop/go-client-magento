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

// checks if the NegotiableQuoteDataCommentInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegotiableQuoteDataCommentInterface{}

// NegotiableQuoteDataCommentInterface Interface CommentInterface
type NegotiableQuoteDataCommentInterface struct {
	// Comment ID.
	EntityId int32 `json:"entity_id"`
	// Negotiable quote ID, that this comment belongs to.
	ParentId int32 `json:"parent_id"`
	// The comment creator type.
	CreatorType int32 `json:"creator_type"`
	// Is quote was declined by seller.
	IsDecline int32 `json:"is_decline"`
	// Is quote draft flag.
	IsDraft int32 `json:"is_draft"`
	// Comment creator ID.
	CreatorId int32 `json:"creator_id"`
	// Comment.
	Comment string `json:"comment"`
	// Comment created at.
	CreatedAt string `json:"created_at"`
	// ExtensionInterface class for @see \\Magento\\NegotiableQuote\\Api\\Data\\CommentInterface
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
	// Existing attachments.
	Attachments          []NegotiableQuoteDataCommentAttachmentInterface `json:"attachments"`
	AdditionalProperties map[string]interface{}
}

type _NegotiableQuoteDataCommentInterface NegotiableQuoteDataCommentInterface

// NewNegotiableQuoteDataCommentInterface instantiates a new NegotiableQuoteDataCommentInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegotiableQuoteDataCommentInterface(entityId int32, parentId int32, creatorType int32, isDecline int32, isDraft int32, creatorId int32, comment string, createdAt string, attachments []NegotiableQuoteDataCommentAttachmentInterface) *NegotiableQuoteDataCommentInterface {
	this := NegotiableQuoteDataCommentInterface{}
	this.EntityId = entityId
	this.ParentId = parentId
	this.CreatorType = creatorType
	this.IsDecline = isDecline
	this.IsDraft = isDraft
	this.CreatorId = creatorId
	this.Comment = comment
	this.CreatedAt = createdAt
	this.Attachments = attachments
	return &this
}

// NewNegotiableQuoteDataCommentInterfaceWithDefaults instantiates a new NegotiableQuoteDataCommentInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegotiableQuoteDataCommentInterfaceWithDefaults() *NegotiableQuoteDataCommentInterface {
	this := NegotiableQuoteDataCommentInterface{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *NegotiableQuoteDataCommentInterface) GetEntityId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetEntityIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *NegotiableQuoteDataCommentInterface) SetEntityId(v int32) {
	o.EntityId = v
}

// GetParentId returns the ParentId field value
func (o *NegotiableQuoteDataCommentInterface) GetParentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *NegotiableQuoteDataCommentInterface) SetParentId(v int32) {
	o.ParentId = v
}

// GetCreatorType returns the CreatorType field value
func (o *NegotiableQuoteDataCommentInterface) GetCreatorType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorType
}

// GetCreatorTypeOk returns a tuple with the CreatorType field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetCreatorTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorType, true
}

// SetCreatorType sets field value
func (o *NegotiableQuoteDataCommentInterface) SetCreatorType(v int32) {
	o.CreatorType = v
}

// GetIsDecline returns the IsDecline field value
func (o *NegotiableQuoteDataCommentInterface) GetIsDecline() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsDecline
}

// GetIsDeclineOk returns a tuple with the IsDecline field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetIsDeclineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDecline, true
}

// SetIsDecline sets field value
func (o *NegotiableQuoteDataCommentInterface) SetIsDecline(v int32) {
	o.IsDecline = v
}

// GetIsDraft returns the IsDraft field value
func (o *NegotiableQuoteDataCommentInterface) GetIsDraft() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsDraft
}

// GetIsDraftOk returns a tuple with the IsDraft field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetIsDraftOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDraft, true
}

// SetIsDraft sets field value
func (o *NegotiableQuoteDataCommentInterface) SetIsDraft(v int32) {
	o.IsDraft = v
}

// GetCreatorId returns the CreatorId field value
func (o *NegotiableQuoteDataCommentInterface) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *NegotiableQuoteDataCommentInterface) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetComment returns the Comment field value
func (o *NegotiableQuoteDataCommentInterface) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *NegotiableQuoteDataCommentInterface) SetComment(v string) {
	o.Comment = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NegotiableQuoteDataCommentInterface) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NegotiableQuoteDataCommentInterface) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *NegotiableQuoteDataCommentInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *NegotiableQuoteDataCommentInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *NegotiableQuoteDataCommentInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

// GetAttachments returns the Attachments field value
func (o *NegotiableQuoteDataCommentInterface) GetAttachments() []NegotiableQuoteDataCommentAttachmentInterface {
	if o == nil {
		var ret []NegotiableQuoteDataCommentAttachmentInterface
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataCommentInterface) GetAttachmentsOk() ([]NegotiableQuoteDataCommentAttachmentInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *NegotiableQuoteDataCommentInterface) SetAttachments(v []NegotiableQuoteDataCommentAttachmentInterface) {
	o.Attachments = v
}

func (o NegotiableQuoteDataCommentInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegotiableQuoteDataCommentInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["parent_id"] = o.ParentId
	toSerialize["creator_type"] = o.CreatorType
	toSerialize["is_decline"] = o.IsDecline
	toSerialize["is_draft"] = o.IsDraft
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["comment"] = o.Comment
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}
	toSerialize["attachments"] = o.Attachments

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NegotiableQuoteDataCommentInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_id",
		"parent_id",
		"creator_type",
		"is_decline",
		"is_draft",
		"creator_id",
		"comment",
		"created_at",
		"attachments",
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

	varNegotiableQuoteDataCommentInterface := _NegotiableQuoteDataCommentInterface{}

	err = json.Unmarshal(data, &varNegotiableQuoteDataCommentInterface)

	if err != nil {
		return err
	}

	*o = NegotiableQuoteDataCommentInterface(varNegotiableQuoteDataCommentInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "creator_type")
		delete(additionalProperties, "is_decline")
		delete(additionalProperties, "is_draft")
		delete(additionalProperties, "creator_id")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "extension_attributes")
		delete(additionalProperties, "attachments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *NegotiableQuoteDataCommentInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *NegotiableQuoteDataCommentInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableNegotiableQuoteDataCommentInterface struct {
	value *NegotiableQuoteDataCommentInterface
	isSet bool
}

func (v NullableNegotiableQuoteDataCommentInterface) Get() *NegotiableQuoteDataCommentInterface {
	return v.value
}

func (v *NullableNegotiableQuoteDataCommentInterface) Set(val *NegotiableQuoteDataCommentInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNegotiableQuoteDataCommentInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNegotiableQuoteDataCommentInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegotiableQuoteDataCommentInterface(val *NegotiableQuoteDataCommentInterface) *NullableNegotiableQuoteDataCommentInterface {
	return &NullableNegotiableQuoteDataCommentInterface{value: val, isSet: true}
}

func (v NullableNegotiableQuoteDataCommentInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegotiableQuoteDataCommentInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
