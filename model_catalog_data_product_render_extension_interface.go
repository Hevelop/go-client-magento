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

// checks if the CatalogDataProductRenderExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDataProductRenderExtensionInterface{}

// CatalogDataProductRenderExtensionInterface ExtensionInterface class for @see \\Magento\\Catalog\\Api\\Data\\ProductRenderInterface
type CatalogDataProductRenderExtensionInterface struct {
	WishlistButton       *CatalogDataProductRenderButtonInterface `json:"wishlist_button,omitempty"`
	ReviewHtml           *string                                  `json:"review_html,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogDataProductRenderExtensionInterface CatalogDataProductRenderExtensionInterface

// NewCatalogDataProductRenderExtensionInterface instantiates a new CatalogDataProductRenderExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDataProductRenderExtensionInterface() *CatalogDataProductRenderExtensionInterface {
	this := CatalogDataProductRenderExtensionInterface{}
	return &this
}

// NewCatalogDataProductRenderExtensionInterfaceWithDefaults instantiates a new CatalogDataProductRenderExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDataProductRenderExtensionInterfaceWithDefaults() *CatalogDataProductRenderExtensionInterface {
	this := CatalogDataProductRenderExtensionInterface{}
	return &this
}

// GetWishlistButton returns the WishlistButton field value if set, zero value otherwise.
func (o *CatalogDataProductRenderExtensionInterface) GetWishlistButton() CatalogDataProductRenderButtonInterface {
	if o == nil || IsNil(o.WishlistButton) {
		var ret CatalogDataProductRenderButtonInterface
		return ret
	}
	return *o.WishlistButton
}

// GetWishlistButtonOk returns a tuple with the WishlistButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderExtensionInterface) GetWishlistButtonOk() (*CatalogDataProductRenderButtonInterface, bool) {
	if o == nil || IsNil(o.WishlistButton) {
		return nil, false
	}
	return o.WishlistButton, true
}

// HasWishlistButton returns a boolean if a field has been set.
func (o *CatalogDataProductRenderExtensionInterface) HasWishlistButton() bool {
	if o != nil && !IsNil(o.WishlistButton) {
		return true
	}

	return false
}

// SetWishlistButton gets a reference to the given CatalogDataProductRenderButtonInterface and assigns it to the WishlistButton field.
func (o *CatalogDataProductRenderExtensionInterface) SetWishlistButton(v CatalogDataProductRenderButtonInterface) {
	o.WishlistButton = &v
}

// GetReviewHtml returns the ReviewHtml field value if set, zero value otherwise.
func (o *CatalogDataProductRenderExtensionInterface) GetReviewHtml() string {
	if o == nil || IsNil(o.ReviewHtml) {
		var ret string
		return ret
	}
	return *o.ReviewHtml
}

// GetReviewHtmlOk returns a tuple with the ReviewHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogDataProductRenderExtensionInterface) GetReviewHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewHtml) {
		return nil, false
	}
	return o.ReviewHtml, true
}

// HasReviewHtml returns a boolean if a field has been set.
func (o *CatalogDataProductRenderExtensionInterface) HasReviewHtml() bool {
	if o != nil && !IsNil(o.ReviewHtml) {
		return true
	}

	return false
}

// SetReviewHtml gets a reference to the given string and assigns it to the ReviewHtml field.
func (o *CatalogDataProductRenderExtensionInterface) SetReviewHtml(v string) {
	o.ReviewHtml = &v
}

func (o CatalogDataProductRenderExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDataProductRenderExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WishlistButton) {
		toSerialize["wishlist_button"] = o.WishlistButton
	}
	if !IsNil(o.ReviewHtml) {
		toSerialize["review_html"] = o.ReviewHtml
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogDataProductRenderExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varCatalogDataProductRenderExtensionInterface := _CatalogDataProductRenderExtensionInterface{}

	err = json.Unmarshal(data, &varCatalogDataProductRenderExtensionInterface)

	if err != nil {
		return err
	}

	*o = CatalogDataProductRenderExtensionInterface(varCatalogDataProductRenderExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "wishlist_button")
		delete(additionalProperties, "review_html")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogDataProductRenderExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogDataProductRenderExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogDataProductRenderExtensionInterface struct {
	value *CatalogDataProductRenderExtensionInterface
	isSet bool
}

func (v NullableCatalogDataProductRenderExtensionInterface) Get() *CatalogDataProductRenderExtensionInterface {
	return v.value
}

func (v *NullableCatalogDataProductRenderExtensionInterface) Set(val *CatalogDataProductRenderExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDataProductRenderExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDataProductRenderExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDataProductRenderExtensionInterface(val *CatalogDataProductRenderExtensionInterface) *NullableCatalogDataProductRenderExtensionInterface {
	return &NullableCatalogDataProductRenderExtensionInterface{value: val, isSet: true}
}

func (v NullableCatalogDataProductRenderExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDataProductRenderExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
