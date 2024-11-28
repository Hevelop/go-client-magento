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

// checks if the SalesDataTransactionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataTransactionInterface{}

// SalesDataTransactionInterface Transaction interface. A transaction is an interaction between a merchant and a customer such as a purchase, a credit, a refund, and so on.
type SalesDataTransactionInterface struct {
	// Transaction ID.
	TransactionId int32 `json:"transaction_id"`
	// The parent ID for the transaction. Otherwise, null.
	ParentId *int32 `json:"parent_id,omitempty"`
	// Order ID.
	OrderId int32 `json:"order_id"`
	// Payment ID.
	PaymentId int32 `json:"payment_id"`
	// Transaction business ID.
	TxnId string `json:"txn_id"`
	// Parent transaction business ID.
	ParentTxnId string `json:"parent_txn_id"`
	// Transaction type.
	TxnType string `json:"txn_type"`
	// Is-closed flag value.
	IsClosed int32 `json:"is_closed"`
	// Array of additional information. Otherwise, null.
	AdditionalInformation []string `json:"additional_information,omitempty"`
	// Created-at timestamp.
	CreatedAt string `json:"created_at"`
	// Array of child transactions.
	ChildTransactions []SalesDataTransactionInterface `json:"child_transactions"`
	// ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\TransactionInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataTransactionInterface SalesDataTransactionInterface

// NewSalesDataTransactionInterface instantiates a new SalesDataTransactionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataTransactionInterface(transactionId int32, orderId int32, paymentId int32, txnId string, parentTxnId string, txnType string, isClosed int32, createdAt string, childTransactions []SalesDataTransactionInterface) *SalesDataTransactionInterface {
	this := SalesDataTransactionInterface{}
	this.TransactionId = transactionId
	this.OrderId = orderId
	this.PaymentId = paymentId
	this.TxnId = txnId
	this.ParentTxnId = parentTxnId
	this.TxnType = txnType
	this.IsClosed = isClosed
	this.CreatedAt = createdAt
	this.ChildTransactions = childTransactions
	return &this
}

// NewSalesDataTransactionInterfaceWithDefaults instantiates a new SalesDataTransactionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataTransactionInterfaceWithDefaults() *SalesDataTransactionInterface {
	this := SalesDataTransactionInterface{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *SalesDataTransactionInterface) GetTransactionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetTransactionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *SalesDataTransactionInterface) SetTransactionId(v int32) {
	o.TransactionId = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *SalesDataTransactionInterface) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *SalesDataTransactionInterface) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *SalesDataTransactionInterface) SetParentId(v int32) {
	o.ParentId = &v
}

// GetOrderId returns the OrderId field value
func (o *SalesDataTransactionInterface) GetOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *SalesDataTransactionInterface) SetOrderId(v int32) {
	o.OrderId = v
}

// GetPaymentId returns the PaymentId field value
func (o *SalesDataTransactionInterface) GetPaymentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetPaymentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *SalesDataTransactionInterface) SetPaymentId(v int32) {
	o.PaymentId = v
}

// GetTxnId returns the TxnId field value
func (o *SalesDataTransactionInterface) GetTxnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetTxnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnId, true
}

// SetTxnId sets field value
func (o *SalesDataTransactionInterface) SetTxnId(v string) {
	o.TxnId = v
}

// GetParentTxnId returns the ParentTxnId field value
func (o *SalesDataTransactionInterface) GetParentTxnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentTxnId
}

// GetParentTxnIdOk returns a tuple with the ParentTxnId field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetParentTxnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentTxnId, true
}

// SetParentTxnId sets field value
func (o *SalesDataTransactionInterface) SetParentTxnId(v string) {
	o.ParentTxnId = v
}

// GetTxnType returns the TxnType field value
func (o *SalesDataTransactionInterface) GetTxnType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnType
}

// GetTxnTypeOk returns a tuple with the TxnType field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetTxnTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnType, true
}

// SetTxnType sets field value
func (o *SalesDataTransactionInterface) SetTxnType(v string) {
	o.TxnType = v
}

// GetIsClosed returns the IsClosed field value
func (o *SalesDataTransactionInterface) GetIsClosed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetIsClosedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClosed, true
}

// SetIsClosed sets field value
func (o *SalesDataTransactionInterface) SetIsClosed(v int32) {
	o.IsClosed = v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *SalesDataTransactionInterface) GetAdditionalInformation() []string {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret []string
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetAdditionalInformationOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *SalesDataTransactionInterface) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given []string and assigns it to the AdditionalInformation field.
func (o *SalesDataTransactionInterface) SetAdditionalInformation(v []string) {
	o.AdditionalInformation = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SalesDataTransactionInterface) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SalesDataTransactionInterface) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetChildTransactions returns the ChildTransactions field value
func (o *SalesDataTransactionInterface) GetChildTransactions() []SalesDataTransactionInterface {
	if o == nil {
		var ret []SalesDataTransactionInterface
		return ret
	}

	return o.ChildTransactions
}

// GetChildTransactionsOk returns a tuple with the ChildTransactions field value
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetChildTransactionsOk() ([]SalesDataTransactionInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChildTransactions, true
}

// SetChildTransactions sets field value
func (o *SalesDataTransactionInterface) SetChildTransactions(v []SalesDataTransactionInterface) {
	o.ChildTransactions = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataTransactionInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataTransactionInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataTransactionInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *SalesDataTransactionInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o SalesDataTransactionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataTransactionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_id"] = o.TransactionId
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	toSerialize["order_id"] = o.OrderId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["txn_id"] = o.TxnId
	toSerialize["parent_txn_id"] = o.ParentTxnId
	toSerialize["txn_type"] = o.TxnType
	toSerialize["is_closed"] = o.IsClosed
	if !IsNil(o.AdditionalInformation) {
		toSerialize["additional_information"] = o.AdditionalInformation
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["child_transactions"] = o.ChildTransactions
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataTransactionInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_id",
		"order_id",
		"payment_id",
		"txn_id",
		"parent_txn_id",
		"txn_type",
		"is_closed",
		"created_at",
		"child_transactions",
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

	varSalesDataTransactionInterface := _SalesDataTransactionInterface{}

	err = json.Unmarshal(data, &varSalesDataTransactionInterface)

	if err != nil {
		return err
	}

	*o = SalesDataTransactionInterface(varSalesDataTransactionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "order_id")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "txn_id")
		delete(additionalProperties, "parent_txn_id")
		delete(additionalProperties, "txn_type")
		delete(additionalProperties, "is_closed")
		delete(additionalProperties, "additional_information")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "child_transactions")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataTransactionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataTransactionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataTransactionInterface struct {
	value *SalesDataTransactionInterface
	isSet bool
}

func (v NullableSalesDataTransactionInterface) Get() *SalesDataTransactionInterface {
	return v.value
}

func (v *NullableSalesDataTransactionInterface) Set(val *SalesDataTransactionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataTransactionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataTransactionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataTransactionInterface(val *SalesDataTransactionInterface) *NullableSalesDataTransactionInterface {
	return &NullableSalesDataTransactionInterface{value: val, isSet: true}
}

func (v NullableSalesDataTransactionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataTransactionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}