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

// checks if the CatalogDataProductLinkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDataProductLinkInterface{}

// CatalogDataProductLinkInterface
type CatalogDataProductLinkInterface struct {
	// SKU
	Sku string `json:"sku"`
	// Link type
	LinkType string `json:"link_type"`
	// Linked product sku
	LinkedProductSku string `json:"linked_product_sku"`
	// Linked product type (simple, virtual, etc)
	LinkedProductType string `json:"linked_product_type"`
	// Linked item position
	Position             int32                                     `json:"position"`
	ExtensionAttributes  *CatalogDataProductLinkExtensionInterface `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogDataProductLinkInterface CatalogDataProductLinkInterface

// NewCatalogDataProductLinkInterface instantiates a new CatalogDataProductLinkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDataProductLinkInterface(sku string, linkType string, linkedProductSku string, linkedProductType string, position int32) *CatalogDataProductLinkInterface {
	this := CatalogDataProductLinkInterface{}
	this.Sku = sku
	this.LinkType = linkType
	this.LinkedProductSku = linkedProductSku
	this.LinkedProductType = linkedProductType
	this.Position = position
	return &this
}

// NewCatalogDataProductLinkInterfaceWithDefaults instantiates a new CatalogDataProductLinkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDataProductLinkInterfaceWithDefaults() *CatalogDataProductLinkInterface {
	this := CatalogDataProductLinkInterface{}
	return &this
}

// GetSku returns the Sku field value
func (o *CatalogDataProductLinkInterface) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductLinkInterface) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *CatalogDataProductLinkInterface) SetSku(v string) {
	o.Sku = v
}

// GetLinkType returns the LinkType field value
func (o *CatalogDataProductLinkInterface) GetLinkType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductLinkInterface) GetLinkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkType, true
}

// SetLinkType sets field value
func (o *CatalogDataProductLinkInterface) SetLinkType(v string) {
	o.LinkType = v
}

// GetLinkedProductSku returns the LinkedProductSku field value
func (o *CatalogDataProductLinkInterface) GetLinkedProductSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkedProductSku
}

// GetLinkedProductSkuOk returns a tuple with the LinkedProductSku field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductLinkInterface) GetLinkedProductSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkedProductSku, true
}

// SetLinkedProductSku sets field value
func (o *CatalogDataProductLinkInterface) SetLinkedProductSku(v string) {
	o.LinkedProductSku = v
}

// GetLinkedProductType returns the LinkedProductType field value
func (o *CatalogDataProductLinkInterface) GetLinkedProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkedProductType
}

// GetLinkedProductTypeOk returns a tuple with the LinkedProductType field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductLinkInterface) GetLinkedProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkedProductType, true
}

// SetLinkedProductType sets field value
func (o *CatalogDataProductLinkInterface) SetLinkedProductType(v string) {
	o.LinkedProductType = v
}

// GetPosition returns the Position field value
func (o *CatalogDataProductLinkInterface) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *CatalogDataProductLinkInterface) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *CatalogDataProductLinkInterface) SetPosition(v int32) {
	o.Position = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *CatalogDataProductLinkInterface) GetExtensionAttributes() CatalogDataProductLinkExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret CatalogDataProductLinkExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogDataProductLinkInterface) GetExtensionAttributesOk() (*CatalogDataProductLinkExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *CatalogDataProductLinkInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given CatalogDataProductLinkExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *CatalogDataProductLinkInterface) SetExtensionAttributes(v CatalogDataProductLinkExtensionInterface) {
	o.ExtensionAttributes = &v
}

func (o CatalogDataProductLinkInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDataProductLinkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	toSerialize["link_type"] = o.LinkType
	toSerialize["linked_product_sku"] = o.LinkedProductSku
	toSerialize["linked_product_type"] = o.LinkedProductType
	toSerialize["position"] = o.Position
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogDataProductLinkInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku",
		"link_type",
		"linked_product_sku",
		"linked_product_type",
		"position",
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

	varCatalogDataProductLinkInterface := _CatalogDataProductLinkInterface{}

	err = json.Unmarshal(data, &varCatalogDataProductLinkInterface)

	if err != nil {
		return err
	}

	*o = CatalogDataProductLinkInterface(varCatalogDataProductLinkInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sku")
		delete(additionalProperties, "link_type")
		delete(additionalProperties, "linked_product_sku")
		delete(additionalProperties, "linked_product_type")
		delete(additionalProperties, "position")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogDataProductLinkInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogDataProductLinkInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogDataProductLinkInterface struct {
	value *CatalogDataProductLinkInterface
	isSet bool
}

func (v NullableCatalogDataProductLinkInterface) Get() *CatalogDataProductLinkInterface {
	return v.value
}

func (v *NullableCatalogDataProductLinkInterface) Set(val *CatalogDataProductLinkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDataProductLinkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDataProductLinkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDataProductLinkInterface(val *CatalogDataProductLinkInterface) *NullableCatalogDataProductLinkInterface {
	return &NullableCatalogDataProductLinkInterface{value: val, isSet: true}
}

func (v NullableCatalogDataProductLinkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDataProductLinkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
