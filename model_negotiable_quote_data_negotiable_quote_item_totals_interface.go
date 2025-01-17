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

// checks if the NegotiableQuoteDataNegotiableQuoteItemTotalsInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegotiableQuoteDataNegotiableQuoteItemTotalsInterface{}

// NegotiableQuoteDataNegotiableQuoteItemTotalsInterface Extension attribute for quote item totals model.
type NegotiableQuoteDataNegotiableQuoteItemTotalsInterface struct {
	// Cost for quote item.
	Cost float32 `json:"cost"`
	// Catalog price for quote item.
	CatalogPrice float32 `json:"catalog_price"`
	// Catalog price for quote item in base currency.
	BaseCatalogPrice float32 `json:"base_catalog_price"`
	// Catalog price with included tax for quote item.
	CatalogPriceInclTax float32 `json:"catalog_price_incl_tax"`
	// Catalog price with included tax for quote item in base currency.
	BaseCatalogPriceInclTax float32 `json:"base_catalog_price_incl_tax"`
	// Cart price for quote item.
	CartPrice float32 `json:"cart_price"`
	// Cart price for quote item in base currency.
	BaseCartPrice float32 `json:"base_cart_price"`
	// Tax from catalog price for quote item.
	CartTax float32 `json:"cart_tax"`
	// Tax from catalog price for quote item in base currency.
	BaseCartTax float32 `json:"base_cart_tax"`
	// Cart price with included tax for quote item.
	CartPriceInclTax float32 `json:"cart_price_incl_tax"`
	// Cart price with included tax for quote item in base currency.
	BaseCartPriceInclTax float32 `json:"base_cart_price_incl_tax"`
	// ExtensionInterface class for @see \\Magento\\NegotiableQuote\\Api\\Data\\NegotiableQuoteItemTotalsInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NegotiableQuoteDataNegotiableQuoteItemTotalsInterface NegotiableQuoteDataNegotiableQuoteItemTotalsInterface

// NewNegotiableQuoteDataNegotiableQuoteItemTotalsInterface instantiates a new NegotiableQuoteDataNegotiableQuoteItemTotalsInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegotiableQuoteDataNegotiableQuoteItemTotalsInterface(cost float32, catalogPrice float32, baseCatalogPrice float32, catalogPriceInclTax float32, baseCatalogPriceInclTax float32, cartPrice float32, baseCartPrice float32, cartTax float32, baseCartTax float32, cartPriceInclTax float32, baseCartPriceInclTax float32) *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface {
	this := NegotiableQuoteDataNegotiableQuoteItemTotalsInterface{}
	this.Cost = cost
	this.CatalogPrice = catalogPrice
	this.BaseCatalogPrice = baseCatalogPrice
	this.CatalogPriceInclTax = catalogPriceInclTax
	this.BaseCatalogPriceInclTax = baseCatalogPriceInclTax
	this.CartPrice = cartPrice
	this.BaseCartPrice = baseCartPrice
	this.CartTax = cartTax
	this.BaseCartTax = baseCartTax
	this.CartPriceInclTax = cartPriceInclTax
	this.BaseCartPriceInclTax = baseCartPriceInclTax
	return &this
}

// NewNegotiableQuoteDataNegotiableQuoteItemTotalsInterfaceWithDefaults instantiates a new NegotiableQuoteDataNegotiableQuoteItemTotalsInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegotiableQuoteDataNegotiableQuoteItemTotalsInterfaceWithDefaults() *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface {
	this := NegotiableQuoteDataNegotiableQuoteItemTotalsInterface{}
	return &this
}

// GetCost returns the Cost field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCost() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetCost(v float32) {
	o.Cost = v
}

// GetCatalogPrice returns the CatalogPrice field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCatalogPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CatalogPrice
}

// GetCatalogPriceOk returns a tuple with the CatalogPrice field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCatalogPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CatalogPrice, true
}

// SetCatalogPrice sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetCatalogPrice(v float32) {
	o.CatalogPrice = v
}

// GetBaseCatalogPrice returns the BaseCatalogPrice field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCatalogPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaseCatalogPrice
}

// GetBaseCatalogPriceOk returns a tuple with the BaseCatalogPrice field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCatalogPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCatalogPrice, true
}

// SetBaseCatalogPrice sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetBaseCatalogPrice(v float32) {
	o.BaseCatalogPrice = v
}

// GetCatalogPriceInclTax returns the CatalogPriceInclTax field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCatalogPriceInclTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CatalogPriceInclTax
}

// GetCatalogPriceInclTaxOk returns a tuple with the CatalogPriceInclTax field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCatalogPriceInclTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CatalogPriceInclTax, true
}

// SetCatalogPriceInclTax sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetCatalogPriceInclTax(v float32) {
	o.CatalogPriceInclTax = v
}

// GetBaseCatalogPriceInclTax returns the BaseCatalogPriceInclTax field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCatalogPriceInclTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaseCatalogPriceInclTax
}

// GetBaseCatalogPriceInclTaxOk returns a tuple with the BaseCatalogPriceInclTax field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCatalogPriceInclTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCatalogPriceInclTax, true
}

// SetBaseCatalogPriceInclTax sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetBaseCatalogPriceInclTax(v float32) {
	o.BaseCatalogPriceInclTax = v
}

// GetCartPrice returns the CartPrice field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCartPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CartPrice
}

// GetCartPriceOk returns a tuple with the CartPrice field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCartPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CartPrice, true
}

// SetCartPrice sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetCartPrice(v float32) {
	o.CartPrice = v
}

// GetBaseCartPrice returns the BaseCartPrice field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCartPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaseCartPrice
}

// GetBaseCartPriceOk returns a tuple with the BaseCartPrice field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCartPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCartPrice, true
}

// SetBaseCartPrice sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetBaseCartPrice(v float32) {
	o.BaseCartPrice = v
}

// GetCartTax returns the CartTax field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCartTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CartTax
}

// GetCartTaxOk returns a tuple with the CartTax field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCartTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CartTax, true
}

// SetCartTax sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetCartTax(v float32) {
	o.CartTax = v
}

// GetBaseCartTax returns the BaseCartTax field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCartTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaseCartTax
}

// GetBaseCartTaxOk returns a tuple with the BaseCartTax field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCartTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCartTax, true
}

// SetBaseCartTax sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetBaseCartTax(v float32) {
	o.BaseCartTax = v
}

// GetCartPriceInclTax returns the CartPriceInclTax field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCartPriceInclTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CartPriceInclTax
}

// GetCartPriceInclTaxOk returns a tuple with the CartPriceInclTax field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetCartPriceInclTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CartPriceInclTax, true
}

// SetCartPriceInclTax sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetCartPriceInclTax(v float32) {
	o.CartPriceInclTax = v
}

// GetBaseCartPriceInclTax returns the BaseCartPriceInclTax field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCartPriceInclTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaseCartPriceInclTax
}

// GetBaseCartPriceInclTaxOk returns a tuple with the BaseCartPriceInclTax field value
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetBaseCartPriceInclTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCartPriceInclTax, true
}

// SetBaseCartPriceInclTax sets field value
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetBaseCartPriceInclTax(v float32) {
	o.BaseCartPriceInclTax = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cost"] = o.Cost
	toSerialize["catalog_price"] = o.CatalogPrice
	toSerialize["base_catalog_price"] = o.BaseCatalogPrice
	toSerialize["catalog_price_incl_tax"] = o.CatalogPriceInclTax
	toSerialize["base_catalog_price_incl_tax"] = o.BaseCatalogPriceInclTax
	toSerialize["cart_price"] = o.CartPrice
	toSerialize["base_cart_price"] = o.BaseCartPrice
	toSerialize["cart_tax"] = o.CartTax
	toSerialize["base_cart_tax"] = o.BaseCartTax
	toSerialize["cart_price_incl_tax"] = o.CartPriceInclTax
	toSerialize["base_cart_price_incl_tax"] = o.BaseCartPriceInclTax
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cost",
		"catalog_price",
		"base_catalog_price",
		"catalog_price_incl_tax",
		"base_catalog_price_incl_tax",
		"cart_price",
		"base_cart_price",
		"cart_tax",
		"base_cart_tax",
		"cart_price_incl_tax",
		"base_cart_price_incl_tax",
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

	varNegotiableQuoteDataNegotiableQuoteItemTotalsInterface := _NegotiableQuoteDataNegotiableQuoteItemTotalsInterface{}

	err = json.Unmarshal(data, &varNegotiableQuoteDataNegotiableQuoteItemTotalsInterface)

	if err != nil {
		return err
	}

	*o = NegotiableQuoteDataNegotiableQuoteItemTotalsInterface(varNegotiableQuoteDataNegotiableQuoteItemTotalsInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cost")
		delete(additionalProperties, "catalog_price")
		delete(additionalProperties, "base_catalog_price")
		delete(additionalProperties, "catalog_price_incl_tax")
		delete(additionalProperties, "base_catalog_price_incl_tax")
		delete(additionalProperties, "cart_price")
		delete(additionalProperties, "base_cart_price")
		delete(additionalProperties, "cart_tax")
		delete(additionalProperties, "base_cart_tax")
		delete(additionalProperties, "cart_price_incl_tax")
		delete(additionalProperties, "base_cart_price_incl_tax")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface struct {
	value *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface
	isSet bool
}

func (v NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface) Get() *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface {
	return v.value
}

func (v *NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface) Set(val *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface(val *NegotiableQuoteDataNegotiableQuoteItemTotalsInterface) *NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface {
	return &NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface{value: val, isSet: true}
}

func (v NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegotiableQuoteDataNegotiableQuoteItemTotalsInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
