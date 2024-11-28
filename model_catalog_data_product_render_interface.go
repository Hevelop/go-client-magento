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

// checks if the CatalogDataProductRenderInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDataProductRenderInterface{}

// CatalogDataProductRenderInterface Represents Data Object which holds enough information to render product This information is put into part as Add To Cart or Add to Compare Data or Price Data
type CatalogDataProductRenderInterface struct {
	AddToCartButton    CatalogDataProductRenderButtonInterface    `json:"add_to_cart_button"`
	AddToCompareButton CatalogDataProductRenderButtonInterface    `json:"add_to_compare_button"`
	PriceInfo          CatalogDataProductRenderPriceInfoInterface `json:"price_info"`
	// Enough information, that needed to render image on front
	Images []CatalogDataProductRenderImageInterface `json:"images"`
	// Product url
	Url string `json:"url"`
	// Product identifier
	Id int32 `json:"id"`
	// Product name
	Name string `json:"name"`
	// Product type. Such as bundle, grouped, simple, etc...
	Type string `json:"type"`
	// Information about product saleability (In Stock)
	IsSalable string `json:"is_salable"`
	// Information about current store id or requested store id
	StoreId int32 `json:"store_id"`
	// Current or desired currency code to product
	CurrencyCode         string                                     `json:"currency_code"`
	ExtensionAttributes  CatalogDataProductRenderExtensionInterface `json:"extension_attributes"`
	AdditionalProperties map[string]interface{}
}

type _CatalogDataProductRenderInterface CatalogDataProductRenderInterface

// NewCatalogDataProductRenderInterface instantiates a new CatalogDataProductRenderInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDataProductRenderInterface(addToCartButton CatalogDataProductRenderButtonInterface, addToCompareButton CatalogDataProductRenderButtonInterface, priceInfo CatalogDataProductRenderPriceInfoInterface, images []CatalogDataProductRenderImageInterface, url string, id int32, name string, type_ string, isSalable string, storeId int32, currencyCode string, extensionAttributes CatalogDataProductRenderExtensionInterface) *CatalogDataProductRenderInterface {
	this := CatalogDataProductRenderInterface{}
	this.AddToCartButton = addToCartButton
	this.AddToCompareButton = addToCompareButton
	this.PriceInfo = priceInfo
	this.Images = images
	this.Url = url
	this.Id = id
	this.Name = name
	this.Type = type_
	this.IsSalable = isSalable
	this.StoreId = storeId
	this.CurrencyCode = currencyCode
	this.ExtensionAttributes = extensionAttributes
	return &this
}

// NewCatalogDataProductRenderInterfaceWithDefaults instantiates a new CatalogDataProductRenderInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDataProductRenderInterfaceWithDefaults() *CatalogDataProductRenderInterface {
	this := CatalogDataProductRenderInterface{}
	return &this
}

// GetAddToCartButton returns the AddToCartButton field value
func (o *CatalogDataProductRenderInterface) GetAddToCartButton() CatalogDataProductRenderButtonInterface {
	if o == nil {
		var ret CatalogDataProductRenderButtonInterface
		return ret
	}

	return o.AddToCartButton
}

// GetAddToCartButtonOk returns a tuple with the AddToCartButton field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetAddToCartButtonOk() (*CatalogDataProductRenderButtonInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddToCartButton, true
}

// SetAddToCartButton sets field value
func (o *CatalogDataProductRenderInterface) SetAddToCartButton(v CatalogDataProductRenderButtonInterface) {
	o.AddToCartButton = v
}

// GetAddToCompareButton returns the AddToCompareButton field value
func (o *CatalogDataProductRenderInterface) GetAddToCompareButton() CatalogDataProductRenderButtonInterface {
	if o == nil {
		var ret CatalogDataProductRenderButtonInterface
		return ret
	}

	return o.AddToCompareButton
}

// GetAddToCompareButtonOk returns a tuple with the AddToCompareButton field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetAddToCompareButtonOk() (*CatalogDataProductRenderButtonInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddToCompareButton, true
}

// SetAddToCompareButton sets field value
func (o *CatalogDataProductRenderInterface) SetAddToCompareButton(v CatalogDataProductRenderButtonInterface) {
	o.AddToCompareButton = v
}

// GetPriceInfo returns the PriceInfo field value
func (o *CatalogDataProductRenderInterface) GetPriceInfo() CatalogDataProductRenderPriceInfoInterface {
	if o == nil {
		var ret CatalogDataProductRenderPriceInfoInterface
		return ret
	}

	return o.PriceInfo
}

// GetPriceInfoOk returns a tuple with the PriceInfo field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetPriceInfoOk() (*CatalogDataProductRenderPriceInfoInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceInfo, true
}

// SetPriceInfo sets field value
func (o *CatalogDataProductRenderInterface) SetPriceInfo(v CatalogDataProductRenderPriceInfoInterface) {
	o.PriceInfo = v
}

// GetImages returns the Images field value
func (o *CatalogDataProductRenderInterface) GetImages() []CatalogDataProductRenderImageInterface {
	if o == nil {
		var ret []CatalogDataProductRenderImageInterface
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetImagesOk() ([]CatalogDataProductRenderImageInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *CatalogDataProductRenderInterface) SetImages(v []CatalogDataProductRenderImageInterface) {
	o.Images = v
}

// GetUrl returns the Url field value
func (o *CatalogDataProductRenderInterface) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CatalogDataProductRenderInterface) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *CatalogDataProductRenderInterface) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CatalogDataProductRenderInterface) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CatalogDataProductRenderInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CatalogDataProductRenderInterface) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CatalogDataProductRenderInterface) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogDataProductRenderInterface) SetType(v string) {
	o.Type = v
}

// GetIsSalable returns the IsSalable field value
func (o *CatalogDataProductRenderInterface) GetIsSalable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsSalable
}

// GetIsSalableOk returns a tuple with the IsSalable field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetIsSalableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSalable, true
}

// SetIsSalable sets field value
func (o *CatalogDataProductRenderInterface) SetIsSalable(v string) {
	o.IsSalable = v
}

// GetStoreId returns the StoreId field value
func (o *CatalogDataProductRenderInterface) GetStoreId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetStoreIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreId, true
}

// SetStoreId sets field value
func (o *CatalogDataProductRenderInterface) SetStoreId(v int32) {
	o.StoreId = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CatalogDataProductRenderInterface) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CatalogDataProductRenderInterface) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value
func (o *CatalogDataProductRenderInterface) GetExtensionAttributes() CatalogDataProductRenderExtensionInterface {
	if o == nil {
		var ret CatalogDataProductRenderExtensionInterface
		return ret
	}

	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderInterface) GetExtensionAttributesOk() (*CatalogDataProductRenderExtensionInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionAttributes, true
}

// SetExtensionAttributes sets field value
func (o *CatalogDataProductRenderInterface) SetExtensionAttributes(v CatalogDataProductRenderExtensionInterface) {
	o.ExtensionAttributes = v
}

func (o CatalogDataProductRenderInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDataProductRenderInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["add_to_cart_button"] = o.AddToCartButton
	toSerialize["add_to_compare_button"] = o.AddToCompareButton
	toSerialize["price_info"] = o.PriceInfo
	toSerialize["images"] = o.Images
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["is_salable"] = o.IsSalable
	toSerialize["store_id"] = o.StoreId
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["extension_attributes"] = o.ExtensionAttributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogDataProductRenderInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"add_to_cart_button",
		"add_to_compare_button",
		"price_info",
		"images",
		"url",
		"id",
		"name",
		"type",
		"is_salable",
		"store_id",
		"currency_code",
		"extension_attributes",
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

	varCatalogDataProductRenderInterface := _CatalogDataProductRenderInterface{}

	err = json.Unmarshal(data, &varCatalogDataProductRenderInterface)

	if err != nil {
		return err
	}

	*o = CatalogDataProductRenderInterface(varCatalogDataProductRenderInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "add_to_cart_button")
		delete(additionalProperties, "add_to_compare_button")
		delete(additionalProperties, "price_info")
		delete(additionalProperties, "images")
		delete(additionalProperties, "url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "is_salable")
		delete(additionalProperties, "store_id")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogDataProductRenderInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogDataProductRenderInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogDataProductRenderInterface struct {
	value *CatalogDataProductRenderInterface
	isSet bool
}

func (v NullableCatalogDataProductRenderInterface) Get() *CatalogDataProductRenderInterface {
	return v.value
}

func (v *NullableCatalogDataProductRenderInterface) Set(val *CatalogDataProductRenderInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDataProductRenderInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDataProductRenderInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDataProductRenderInterface(val *CatalogDataProductRenderInterface) *NullableCatalogDataProductRenderInterface {
	return &NullableCatalogDataProductRenderInterface{value: val, isSet: true}
}

func (v NullableCatalogDataProductRenderInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDataProductRenderInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}