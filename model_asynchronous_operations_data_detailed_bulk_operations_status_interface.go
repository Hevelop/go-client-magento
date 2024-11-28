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

// checks if the AsynchronousOperationsDataDetailedBulkOperationsStatusInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsynchronousOperationsDataDetailedBulkOperationsStatusInterface{}

// AsynchronousOperationsDataDetailedBulkOperationsStatusInterface Interface BulkStatusInterface Bulk summary data with list of operations items full data.
type AsynchronousOperationsDataDetailedBulkOperationsStatusInterface struct {
	// Operations list.
	OperationsList []AsynchronousOperationsDataOperationInterface `json:"operations_list"`
	// ExtensionInterface class for @see \\Magento\\AsynchronousOperations\\Api\\Data\\BulkSummaryInterface
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
	// User type
	UserType int32 `json:"user_type"`
	// Bulk uuid
	BulkId string `json:"bulk_id"`
	// Bulk description
	Description string `json:"description"`
	// Bulk scheduled time
	StartTime string `json:"start_time"`
	// User id
	UserId int32 `json:"user_id"`
	// Total number of operations scheduled in scope of this bulk
	OperationCount       int32 `json:"operation_count"`
	AdditionalProperties map[string]interface{}
}

type _AsynchronousOperationsDataDetailedBulkOperationsStatusInterface AsynchronousOperationsDataDetailedBulkOperationsStatusInterface

// NewAsynchronousOperationsDataDetailedBulkOperationsStatusInterface instantiates a new AsynchronousOperationsDataDetailedBulkOperationsStatusInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsynchronousOperationsDataDetailedBulkOperationsStatusInterface(operationsList []AsynchronousOperationsDataOperationInterface, userType int32, bulkId string, description string, startTime string, userId int32, operationCount int32) *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface {
	this := AsynchronousOperationsDataDetailedBulkOperationsStatusInterface{}
	this.OperationsList = operationsList
	this.UserType = userType
	this.BulkId = bulkId
	this.Description = description
	this.StartTime = startTime
	this.UserId = userId
	this.OperationCount = operationCount
	return &this
}

// NewAsynchronousOperationsDataDetailedBulkOperationsStatusInterfaceWithDefaults instantiates a new AsynchronousOperationsDataDetailedBulkOperationsStatusInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsynchronousOperationsDataDetailedBulkOperationsStatusInterfaceWithDefaults() *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface {
	this := AsynchronousOperationsDataDetailedBulkOperationsStatusInterface{}
	return &this
}

// GetOperationsList returns the OperationsList field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetOperationsList() []AsynchronousOperationsDataOperationInterface {
	if o == nil {
		var ret []AsynchronousOperationsDataOperationInterface
		return ret
	}

	return o.OperationsList
}

// GetOperationsListOk returns a tuple with the OperationsList field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetOperationsListOk() ([]AsynchronousOperationsDataOperationInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationsList, true
}

// SetOperationsList sets field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetOperationsList(v []AsynchronousOperationsDataOperationInterface) {
	o.OperationsList = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

// GetUserType returns the UserType field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetUserType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetUserTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserType, true
}

// SetUserType sets field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetUserType(v int32) {
	o.UserType = v
}

// GetBulkId returns the BulkId field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetBulkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetBulkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BulkId, true
}

// SetBulkId sets field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetBulkId(v string) {
	o.BulkId = v
}

// GetDescription returns the Description field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetDescription(v string) {
	o.Description = v
}

// GetStartTime returns the StartTime field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetStartTime(v string) {
	o.StartTime = v
}

// GetUserId returns the UserId field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetUserId(v int32) {
	o.UserId = v
}

// GetOperationCount returns the OperationCount field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetOperationCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OperationCount
}

// GetOperationCountOk returns a tuple with the OperationCount field value
// and a boolean to check if the value has been set.
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetOperationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationCount, true
}

// SetOperationCount sets field value
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetOperationCount(v int32) {
	o.OperationCount = v
}

func (o AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operations_list"] = o.OperationsList
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}
	toSerialize["user_type"] = o.UserType
	toSerialize["bulk_id"] = o.BulkId
	toSerialize["description"] = o.Description
	toSerialize["start_time"] = o.StartTime
	toSerialize["user_id"] = o.UserId
	toSerialize["operation_count"] = o.OperationCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operations_list",
		"user_type",
		"bulk_id",
		"description",
		"start_time",
		"user_id",
		"operation_count",
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

	varAsynchronousOperationsDataDetailedBulkOperationsStatusInterface := _AsynchronousOperationsDataDetailedBulkOperationsStatusInterface{}

	err = json.Unmarshal(data, &varAsynchronousOperationsDataDetailedBulkOperationsStatusInterface)

	if err != nil {
		return err
	}

	*o = AsynchronousOperationsDataDetailedBulkOperationsStatusInterface(varAsynchronousOperationsDataDetailedBulkOperationsStatusInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operations_list")
		delete(additionalProperties, "extension_attributes")
		delete(additionalProperties, "user_type")
		delete(additionalProperties, "bulk_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "operation_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface struct {
	value *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface
	isSet bool
}

func (v NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface) Get() *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface {
	return v.value
}

func (v *NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface) Set(val *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface(val *AsynchronousOperationsDataDetailedBulkOperationsStatusInterface) *NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface {
	return &NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface{value: val, isSet: true}
}

func (v NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsynchronousOperationsDataDetailedBulkOperationsStatusInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
