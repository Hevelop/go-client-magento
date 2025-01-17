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

// checks if the AsynchronousOperationsDataOperationInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsynchronousOperationsDataOperationInterface{}

// AsynchronousOperationsDataOperationInterface Class OperationInterface
type AsynchronousOperationsDataOperationInterface struct {
	ExtensionAttributes *AsynchronousOperationsDataOperationExtensionInterface `json:"extension_attributes,omitempty"`
	// Id
	Id int32 `json:"id"`
	// Bulk uuid
	BulkUuid string `json:"bulk_uuid"`
	// Queue Topic
	TopicName string `json:"topic_name"`
	// Data
	SerializedData string `json:"serialized_data"`
	// Serialized Data
	ResultSerializedData string `json:"result_serialized_data"`
	// Operation status
	Status int32 `json:"status"`
	// Result message
	ResultMessage string `json:"result_message"`
	// Error code
	ErrorCode            int32 `json:"error_code"`
	AdditionalProperties map[string]interface{}
}

type _AsynchronousOperationsDataOperationInterface AsynchronousOperationsDataOperationInterface

// NewAsynchronousOperationsDataOperationInterface instantiates a new AsynchronousOperationsDataOperationInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsynchronousOperationsDataOperationInterface(id int32, bulkUuid string, topicName string, serializedData string, resultSerializedData string, status int32, resultMessage string, errorCode int32) *AsynchronousOperationsDataOperationInterface {
	this := AsynchronousOperationsDataOperationInterface{}
	this.Id = id
	this.BulkUuid = bulkUuid
	this.TopicName = topicName
	this.SerializedData = serializedData
	this.ResultSerializedData = resultSerializedData
	this.Status = status
	this.ResultMessage = resultMessage
	this.ErrorCode = errorCode
	return &this
}

// NewAsynchronousOperationsDataOperationInterfaceWithDefaults instantiates a new AsynchronousOperationsDataOperationInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsynchronousOperationsDataOperationInterfaceWithDefaults() *AsynchronousOperationsDataOperationInterface {
	this := AsynchronousOperationsDataOperationInterface{}
	return &this
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *AsynchronousOperationsDataOperationInterface) GetExtensionAttributes() AsynchronousOperationsDataOperationExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret AsynchronousOperationsDataOperationExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetExtensionAttributesOk() (*AsynchronousOperationsDataOperationExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *AsynchronousOperationsDataOperationInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given AsynchronousOperationsDataOperationExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *AsynchronousOperationsDataOperationInterface) SetExtensionAttributes(v AsynchronousOperationsDataOperationExtensionInterface) {
	o.ExtensionAttributes = &v
}

// GetId returns the Id field value
func (o *AsynchronousOperationsDataOperationInterface) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetId(v int32) {
	o.Id = v
}

// GetBulkUuid returns the BulkUuid field value
func (o *AsynchronousOperationsDataOperationInterface) GetBulkUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BulkUuid
}

// GetBulkUuidOk returns a tuple with the BulkUuid field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetBulkUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BulkUuid, true
}

// SetBulkUuid sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetBulkUuid(v string) {
	o.BulkUuid = v
}

// GetTopicName returns the TopicName field value
func (o *AsynchronousOperationsDataOperationInterface) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetTopicName(v string) {
	o.TopicName = v
}

// GetSerializedData returns the SerializedData field value
func (o *AsynchronousOperationsDataOperationInterface) GetSerializedData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerializedData
}

// GetSerializedDataOk returns a tuple with the SerializedData field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetSerializedDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerializedData, true
}

// SetSerializedData sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetSerializedData(v string) {
	o.SerializedData = v
}

// GetResultSerializedData returns the ResultSerializedData field value
func (o *AsynchronousOperationsDataOperationInterface) GetResultSerializedData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultSerializedData
}

// GetResultSerializedDataOk returns a tuple with the ResultSerializedData field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetResultSerializedDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultSerializedData, true
}

// SetResultSerializedData sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetResultSerializedData(v string) {
	o.ResultSerializedData = v
}

// GetStatus returns the Status field value
func (o *AsynchronousOperationsDataOperationInterface) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetStatus(v int32) {
	o.Status = v
}

// GetResultMessage returns the ResultMessage field value
func (o *AsynchronousOperationsDataOperationInterface) GetResultMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetResultMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultMessage, true
}

// SetResultMessage sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetResultMessage(v string) {
	o.ResultMessage = v
}

// GetErrorCode returns the ErrorCode field value
func (o *AsynchronousOperationsDataOperationInterface) GetErrorCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataOperationInterface) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *AsynchronousOperationsDataOperationInterface) SetErrorCode(v int32) {
	o.ErrorCode = v
}

func (o AsynchronousOperationsDataOperationInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsynchronousOperationsDataOperationInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}
	toSerialize["id"] = o.Id
	toSerialize["bulk_uuid"] = o.BulkUuid
	toSerialize["topic_name"] = o.TopicName
	toSerialize["serialized_data"] = o.SerializedData
	toSerialize["result_serialized_data"] = o.ResultSerializedData
	toSerialize["status"] = o.Status
	toSerialize["result_message"] = o.ResultMessage
	toSerialize["error_code"] = o.ErrorCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AsynchronousOperationsDataOperationInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"bulk_uuid",
		"topic_name",
		"serialized_data",
		"result_serialized_data",
		"status",
		"result_message",
		"error_code",
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

	varAsynchronousOperationsDataOperationInterface := _AsynchronousOperationsDataOperationInterface{}

	err = json.Unmarshal(data, &varAsynchronousOperationsDataOperationInterface)

	if err != nil {
		return err
	}

	*o = AsynchronousOperationsDataOperationInterface(varAsynchronousOperationsDataOperationInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extension_attributes")
		delete(additionalProperties, "id")
		delete(additionalProperties, "bulk_uuid")
		delete(additionalProperties, "topic_name")
		delete(additionalProperties, "serialized_data")
		delete(additionalProperties, "result_serialized_data")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result_message")
		delete(additionalProperties, "error_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AsynchronousOperationsDataOperationInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AsynchronousOperationsDataOperationInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAsynchronousOperationsDataOperationInterface struct {
	value *AsynchronousOperationsDataOperationInterface
	isSet bool
}

func (v NullableAsynchronousOperationsDataOperationInterface) Get() *AsynchronousOperationsDataOperationInterface {
	return v.value
}

func (v *NullableAsynchronousOperationsDataOperationInterface) Set(val *AsynchronousOperationsDataOperationInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableAsynchronousOperationsDataOperationInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableAsynchronousOperationsDataOperationInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsynchronousOperationsDataOperationInterface(val *AsynchronousOperationsDataOperationInterface) *NullableAsynchronousOperationsDataOperationInterface {
	return &NullableAsynchronousOperationsDataOperationInterface{value: val, isSet: true}
}

func (v NullableAsynchronousOperationsDataOperationInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsynchronousOperationsDataOperationInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
