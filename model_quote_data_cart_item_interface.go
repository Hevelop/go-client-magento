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

// checks if the QuoteDataCartItemInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDataCartItemInterface{}

// QuoteDataCartItemInterface Interface CartItemInterface
type QuoteDataCartItemInterface struct {
	// Item ID. Otherwise, null.
	ItemId *int32 `json:"item_id,omitempty"`
	// Product SKU. Otherwise, null.
	Sku *string `json:"sku,omitempty"`
	// Product quantity.
	Qty float32 `json:"qty"`
	// Product name. Otherwise, null.
	Name *string `json:"name,omitempty"`
	// Product price. Otherwise, null.
	Price *float32 `json:"price,omitempty"`
	// Product type. Otherwise, null.
	ProductType *string `json:"product_type,omitempty"`
	// Quote id.
	QuoteId              string                               `json:"quote_id"`
	ProductOption        *QuoteDataProductOptionInterface     `json:"product_option,omitempty"`
	ExtensionAttributes  *QuoteDataCartItemExtensionInterface `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuoteDataCartItemInterface QuoteDataCartItemInterface

// NewQuoteDataCartItemInterface instantiates a new QuoteDataCartItemInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDataCartItemInterface(qty float32, quoteId string) *QuoteDataCartItemInterface {
	this := QuoteDataCartItemInterface{}
	this.Qty = qty
	this.QuoteId = quoteId
	return &this
}

// NewQuoteDataCartItemInterfaceWithDefaults instantiates a new QuoteDataCartItemInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDataCartItemInterfaceWithDefaults() *QuoteDataCartItemInterface {
	this := QuoteDataCartItemInterface{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *QuoteDataCartItemInterface) GetItemId() int32 {
	if o == nil || IsNil(o.ItemId) {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetItemIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *QuoteDataCartItemInterface) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *QuoteDataCartItemInterface) SetItemId(v int32) {
	o.ItemId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *QuoteDataCartItemInterface) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *QuoteDataCartItemInterface) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *QuoteDataCartItemInterface) SetSku(v string) {
	o.Sku = &v
}

// GetQty returns the Qty field value
func (o *QuoteDataCartItemInterface) GetQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *QuoteDataCartItemInterface) SetQty(v float32) {
	o.Qty = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QuoteDataCartItemInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QuoteDataCartItemInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QuoteDataCartItemInterface) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QuoteDataCartItemInterface) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QuoteDataCartItemInterface) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *QuoteDataCartItemInterface) SetPrice(v float32) {
	o.Price = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *QuoteDataCartItemInterface) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *QuoteDataCartItemInterface) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *QuoteDataCartItemInterface) SetProductType(v string) {
	o.ProductType = &v
}

// GetQuoteId returns the QuoteId field value
func (o *QuoteDataCartItemInterface) GetQuoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetQuoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteId, true
}

// SetQuoteId sets field value
func (o *QuoteDataCartItemInterface) SetQuoteId(v string) {
	o.QuoteId = v
}

// GetProductOption returns the ProductOption field value if set, zero value otherwise.
func (o *QuoteDataCartItemInterface) GetProductOption() QuoteDataProductOptionInterface {
	if o == nil || IsNil(o.ProductOption) {
		var ret QuoteDataProductOptionInterface
		return ret
	}
	return *o.ProductOption
}

// GetProductOptionOk returns a tuple with the ProductOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetProductOptionOk() (*QuoteDataProductOptionInterface, bool) {
	if o == nil || IsNil(o.ProductOption) {
		return nil, false
	}
	return o.ProductOption, true
}

// HasProductOption returns a boolean if a field has been set.
func (o *QuoteDataCartItemInterface) HasProductOption() bool {
	if o != nil && !IsNil(o.ProductOption) {
		return true
	}

	return false
}

// SetProductOption gets a reference to the given QuoteDataProductOptionInterface and assigns it to the ProductOption field.
func (o *QuoteDataCartItemInterface) SetProductOption(v QuoteDataProductOptionInterface) {
	o.ProductOption = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *QuoteDataCartItemInterface) GetExtensionAttributes() QuoteDataCartItemExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret QuoteDataCartItemExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataCartItemInterface) GetExtensionAttributesOk() (*QuoteDataCartItemExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *QuoteDataCartItemInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given QuoteDataCartItemExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *QuoteDataCartItemInterface) SetExtensionAttributes(v QuoteDataCartItemExtensionInterface) {
	o.ExtensionAttributes = &v
}

func (o QuoteDataCartItemInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDataCartItemInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	toSerialize["qty"] = o.Qty
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ProductType) {
		toSerialize["product_type"] = o.ProductType
	}
	toSerialize["quote_id"] = o.QuoteId
	if !IsNil(o.ProductOption) {
		toSerialize["product_option"] = o.ProductOption
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuoteDataCartItemInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"qty",
		"quote_id",
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

	varQuoteDataCartItemInterface := _QuoteDataCartItemInterface{}

	err = json.Unmarshal(data, &varQuoteDataCartItemInterface)

	if err != nil {
		return err
	}

	*o = QuoteDataCartItemInterface(varQuoteDataCartItemInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "name")
		delete(additionalProperties, "price")
		delete(additionalProperties, "product_type")
		delete(additionalProperties, "quote_id")
		delete(additionalProperties, "product_option")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *QuoteDataCartItemInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *QuoteDataCartItemInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableQuoteDataCartItemInterface struct {
	value *QuoteDataCartItemInterface
	isSet bool
}

func (v NullableQuoteDataCartItemInterface) Get() *QuoteDataCartItemInterface {
	return v.value
}

func (v *NullableQuoteDataCartItemInterface) Set(val *QuoteDataCartItemInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDataCartItemInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDataCartItemInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDataCartItemInterface(val *QuoteDataCartItemInterface) *NullableQuoteDataCartItemInterface {
	return &NullableQuoteDataCartItemInterface{value: val, isSet: true}
}

func (v NullableQuoteDataCartItemInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDataCartItemInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
