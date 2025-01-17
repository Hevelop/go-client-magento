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

// checks if the CatalogDataBasePriceInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDataBasePriceInterface{}

// CatalogDataBasePriceInterface Price interface.
type CatalogDataBasePriceInterface struct {
	// Price.
	Price float32 `json:"price"`
	// Store id.
	StoreId int32 `json:"store_id"`
	// SKU.
	Sku string `json:"sku"`
	// ExtensionInterface class for @see \\Magento\\Catalog\\Api\\Data\\BasePriceInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogDataBasePriceInterface CatalogDataBasePriceInterface

// NewCatalogDataBasePriceInterface instantiates a new CatalogDataBasePriceInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDataBasePriceInterface(price float32, storeId int32, sku string) *CatalogDataBasePriceInterface {
	this := CatalogDataBasePriceInterface{}
	this.Price = price
	this.StoreId = storeId
	this.Sku = sku
	return &this
}

// NewCatalogDataBasePriceInterfaceWithDefaults instantiates a new CatalogDataBasePriceInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDataBasePriceInterfaceWithDefaults() *CatalogDataBasePriceInterface {
	this := CatalogDataBasePriceInterface{}
	return &this
}

// GetPrice returns the Price field value
func (o *CatalogDataBasePriceInterface) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *CatalogDataBasePriceInterface) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *CatalogDataBasePriceInterface) SetPrice(v float32) {
	o.Price = v
}

// GetStoreId returns the StoreId field value
func (o *CatalogDataBasePriceInterface) GetStoreId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value
// and a boolean to check if the value has been set.
func (o *CatalogDataBasePriceInterface) GetStoreIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreId, true
}

// SetStoreId sets field value
func (o *CatalogDataBasePriceInterface) SetStoreId(v int32) {
	o.StoreId = v
}

// GetSku returns the Sku field value
func (o *CatalogDataBasePriceInterface) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *CatalogDataBasePriceInterface) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *CatalogDataBasePriceInterface) SetSku(v string) {
	o.Sku = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *CatalogDataBasePriceInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogDataBasePriceInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *CatalogDataBasePriceInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *CatalogDataBasePriceInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o CatalogDataBasePriceInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDataBasePriceInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["price"] = o.Price
	toSerialize["store_id"] = o.StoreId
	toSerialize["sku"] = o.Sku
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogDataBasePriceInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"price",
		"store_id",
		"sku",
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

	varCatalogDataBasePriceInterface := _CatalogDataBasePriceInterface{}

	err = json.Unmarshal(data, &varCatalogDataBasePriceInterface)

	if err != nil {
		return err
	}

	*o = CatalogDataBasePriceInterface(varCatalogDataBasePriceInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "store_id")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogDataBasePriceInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogDataBasePriceInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogDataBasePriceInterface struct {
	value *CatalogDataBasePriceInterface
	isSet bool
}

func (v NullableCatalogDataBasePriceInterface) Get() *CatalogDataBasePriceInterface {
	return v.value
}

func (v *NullableCatalogDataBasePriceInterface) Set(val *CatalogDataBasePriceInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDataBasePriceInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDataBasePriceInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDataBasePriceInterface(val *CatalogDataBasePriceInterface) *NullableCatalogDataBasePriceInterface {
	return &NullableCatalogDataBasePriceInterface{value: val, isSet: true}
}

func (v NullableCatalogDataBasePriceInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDataBasePriceInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
