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

// checks if the QuoteDataCartItemExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDataCartItemExtensionInterface{}

// QuoteDataCartItemExtensionInterface ExtensionInterface class for @see \\Magento\\Quote\\Api\\Data\\CartItemInterface
type QuoteDataCartItemExtensionInterface struct {
	Discounts            []SalesRuleDataRuleDiscountInterface             `json:"discounts,omitempty"`
	NegotiableQuoteItem  *NegotiableQuoteDataNegotiableQuoteItemInterface `json:"negotiable_quote_item,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuoteDataCartItemExtensionInterface QuoteDataCartItemExtensionInterface

// NewQuoteDataCartItemExtensionInterface instantiates a new QuoteDataCartItemExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDataCartItemExtensionInterface() *QuoteDataCartItemExtensionInterface {
	this := QuoteDataCartItemExtensionInterface{}
	return &this
}

// NewQuoteDataCartItemExtensionInterfaceWithDefaults instantiates a new QuoteDataCartItemExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDataCartItemExtensionInterfaceWithDefaults() *QuoteDataCartItemExtensionInterface {
	this := QuoteDataCartItemExtensionInterface{}
	return &this
}

// GetDiscounts returns the Discounts field value if set, zero value otherwise.
func (o *QuoteDataCartItemExtensionInterface) GetDiscounts() []SalesRuleDataRuleDiscountInterface {
	if o == nil || IsNil(o.Discounts) {
		var ret []SalesRuleDataRuleDiscountInterface
		return ret
	}
	return o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemExtensionInterface) GetDiscountsOk() ([]SalesRuleDataRuleDiscountInterface, bool) {
	if o == nil || IsNil(o.Discounts) {
		return nil, false
	}
	return o.Discounts, true
}

// HasDiscounts returns a boolean if a field has been set.
func (o *QuoteDataCartItemExtensionInterface) HasDiscounts() bool {
	if o != nil && !IsNil(o.Discounts) {
		return true
	}

	return false
}

// SetDiscounts gets a reference to the given []SalesRuleDataRuleDiscountInterface and assigns it to the Discounts field.
func (o *QuoteDataCartItemExtensionInterface) SetDiscounts(v []SalesRuleDataRuleDiscountInterface) {
	o.Discounts = v
}

// GetNegotiableQuoteItem returns the NegotiableQuoteItem field value if set, zero value otherwise.
func (o *QuoteDataCartItemExtensionInterface) GetNegotiableQuoteItem() NegotiableQuoteDataNegotiableQuoteItemInterface {
	if o == nil || IsNil(o.NegotiableQuoteItem) {
		var ret NegotiableQuoteDataNegotiableQuoteItemInterface
		return ret
	}
	return *o.NegotiableQuoteItem
}

// GetNegotiableQuoteItemOk returns a tuple with the NegotiableQuoteItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemExtensionInterface) GetNegotiableQuoteItemOk() (*NegotiableQuoteDataNegotiableQuoteItemInterface, bool) {
	if o == nil || IsNil(o.NegotiableQuoteItem) {
		return nil, false
	}
	return o.NegotiableQuoteItem, true
}

// HasNegotiableQuoteItem returns a boolean if a field has been set.
func (o *QuoteDataCartItemExtensionInterface) HasNegotiableQuoteItem() bool {
	if o != nil && !IsNil(o.NegotiableQuoteItem) {
		return true
	}

	return false
}

// SetNegotiableQuoteItem gets a reference to the given NegotiableQuoteDataNegotiableQuoteItemInterface and assigns it to the NegotiableQuoteItem field.
func (o *QuoteDataCartItemExtensionInterface) SetNegotiableQuoteItem(v NegotiableQuoteDataNegotiableQuoteItemInterface) {
	o.NegotiableQuoteItem = &v
}

func (o QuoteDataCartItemExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDataCartItemExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discounts) {
		toSerialize["discounts"] = o.Discounts
	}
	if !IsNil(o.NegotiableQuoteItem) {
		toSerialize["negotiable_quote_item"] = o.NegotiableQuoteItem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuoteDataCartItemExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varQuoteDataCartItemExtensionInterface := _QuoteDataCartItemExtensionInterface{}

	err = json.Unmarshal(data, &varQuoteDataCartItemExtensionInterface)

	if err != nil {
		return err
	}

	*o = QuoteDataCartItemExtensionInterface(varQuoteDataCartItemExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "discounts")
		delete(additionalProperties, "negotiable_quote_item")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *QuoteDataCartItemExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *QuoteDataCartItemExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableQuoteDataCartItemExtensionInterface struct {
	value *QuoteDataCartItemExtensionInterface
	isSet bool
}

func (v NullableQuoteDataCartItemExtensionInterface) Get() *QuoteDataCartItemExtensionInterface {
	return v.value
}

func (v *NullableQuoteDataCartItemExtensionInterface) Set(val *QuoteDataCartItemExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDataCartItemExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDataCartItemExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDataCartItemExtensionInterface(val *QuoteDataCartItemExtensionInterface) *NullableQuoteDataCartItemExtensionInterface {
	return &NullableQuoteDataCartItemExtensionInterface{value: val, isSet: true}
}

func (v NullableQuoteDataCartItemExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDataCartItemExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
