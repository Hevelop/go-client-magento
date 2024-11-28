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

// checks if the SalesDataCreditmemoExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataCreditmemoExtensionInterface{}

// SalesDataCreditmemoExtensionInterface ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\CreditmemoInterface
type SalesDataCreditmemoExtensionInterface struct {
	BaseCustomerBalanceAmount *float32 `json:"base_customer_balance_amount,omitempty"`
	CustomerBalanceAmount     *float32 `json:"customer_balance_amount,omitempty"`
	BaseGiftCardsAmount       *float32 `json:"base_gift_cards_amount,omitempty"`
	GiftCardsAmount           *float32 `json:"gift_cards_amount,omitempty"`
	GwBasePrice               *string  `json:"gw_base_price,omitempty"`
	GwPrice                   *string  `json:"gw_price,omitempty"`
	GwItemsBasePrice          *string  `json:"gw_items_base_price,omitempty"`
	GwItemsPrice              *string  `json:"gw_items_price,omitempty"`
	GwCardBasePrice           *string  `json:"gw_card_base_price,omitempty"`
	GwCardPrice               *string  `json:"gw_card_price,omitempty"`
	GwBaseTaxAmount           *string  `json:"gw_base_tax_amount,omitempty"`
	GwTaxAmount               *string  `json:"gw_tax_amount,omitempty"`
	GwItemsBaseTaxAmount      *string  `json:"gw_items_base_tax_amount,omitempty"`
	GwItemsTaxAmount          *string  `json:"gw_items_tax_amount,omitempty"`
	GwCardBaseTaxAmount       *string  `json:"gw_card_base_tax_amount,omitempty"`
	GwCardTaxAmount           *string  `json:"gw_card_tax_amount,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _SalesDataCreditmemoExtensionInterface SalesDataCreditmemoExtensionInterface

// NewSalesDataCreditmemoExtensionInterface instantiates a new SalesDataCreditmemoExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataCreditmemoExtensionInterface() *SalesDataCreditmemoExtensionInterface {
	this := SalesDataCreditmemoExtensionInterface{}
	return &this
}

// NewSalesDataCreditmemoExtensionInterfaceWithDefaults instantiates a new SalesDataCreditmemoExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataCreditmemoExtensionInterfaceWithDefaults() *SalesDataCreditmemoExtensionInterface {
	this := SalesDataCreditmemoExtensionInterface{}
	return &this
}

// GetBaseCustomerBalanceAmount returns the BaseCustomerBalanceAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetBaseCustomerBalanceAmount() float32 {
	if o == nil || IsNil(o.BaseCustomerBalanceAmount) {
		var ret float32
		return ret
	}
	return *o.BaseCustomerBalanceAmount
}

// GetBaseCustomerBalanceAmountOk returns a tuple with the BaseCustomerBalanceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetBaseCustomerBalanceAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseCustomerBalanceAmount) {
		return nil, false
	}
	return o.BaseCustomerBalanceAmount, true
}

// HasBaseCustomerBalanceAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasBaseCustomerBalanceAmount() bool {
	if o != nil && !IsNil(o.BaseCustomerBalanceAmount) {
		return true
	}

	return false
}

// SetBaseCustomerBalanceAmount gets a reference to the given float32 and assigns it to the BaseCustomerBalanceAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetBaseCustomerBalanceAmount(v float32) {
	o.BaseCustomerBalanceAmount = &v
}

// GetCustomerBalanceAmount returns the CustomerBalanceAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetCustomerBalanceAmount() float32 {
	if o == nil || IsNil(o.CustomerBalanceAmount) {
		var ret float32
		return ret
	}
	return *o.CustomerBalanceAmount
}

// GetCustomerBalanceAmountOk returns a tuple with the CustomerBalanceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetCustomerBalanceAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CustomerBalanceAmount) {
		return nil, false
	}
	return o.CustomerBalanceAmount, true
}

// HasCustomerBalanceAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasCustomerBalanceAmount() bool {
	if o != nil && !IsNil(o.CustomerBalanceAmount) {
		return true
	}

	return false
}

// SetCustomerBalanceAmount gets a reference to the given float32 and assigns it to the CustomerBalanceAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetCustomerBalanceAmount(v float32) {
	o.CustomerBalanceAmount = &v
}

// GetBaseGiftCardsAmount returns the BaseGiftCardsAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetBaseGiftCardsAmount() float32 {
	if o == nil || IsNil(o.BaseGiftCardsAmount) {
		var ret float32
		return ret
	}
	return *o.BaseGiftCardsAmount
}

// GetBaseGiftCardsAmountOk returns a tuple with the BaseGiftCardsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetBaseGiftCardsAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseGiftCardsAmount) {
		return nil, false
	}
	return o.BaseGiftCardsAmount, true
}

// HasBaseGiftCardsAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasBaseGiftCardsAmount() bool {
	if o != nil && !IsNil(o.BaseGiftCardsAmount) {
		return true
	}

	return false
}

// SetBaseGiftCardsAmount gets a reference to the given float32 and assigns it to the BaseGiftCardsAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetBaseGiftCardsAmount(v float32) {
	o.BaseGiftCardsAmount = &v
}

// GetGiftCardsAmount returns the GiftCardsAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGiftCardsAmount() float32 {
	if o == nil || IsNil(o.GiftCardsAmount) {
		var ret float32
		return ret
	}
	return *o.GiftCardsAmount
}

// GetGiftCardsAmountOk returns a tuple with the GiftCardsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGiftCardsAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.GiftCardsAmount) {
		return nil, false
	}
	return o.GiftCardsAmount, true
}

// HasGiftCardsAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGiftCardsAmount() bool {
	if o != nil && !IsNil(o.GiftCardsAmount) {
		return true
	}

	return false
}

// SetGiftCardsAmount gets a reference to the given float32 and assigns it to the GiftCardsAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetGiftCardsAmount(v float32) {
	o.GiftCardsAmount = &v
}

// GetGwBasePrice returns the GwBasePrice field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwBasePrice() string {
	if o == nil || IsNil(o.GwBasePrice) {
		var ret string
		return ret
	}
	return *o.GwBasePrice
}

// GetGwBasePriceOk returns a tuple with the GwBasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwBasePriceOk() (*string, bool) {
	if o == nil || IsNil(o.GwBasePrice) {
		return nil, false
	}
	return o.GwBasePrice, true
}

// HasGwBasePrice returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwBasePrice() bool {
	if o != nil && !IsNil(o.GwBasePrice) {
		return true
	}

	return false
}

// SetGwBasePrice gets a reference to the given string and assigns it to the GwBasePrice field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwBasePrice(v string) {
	o.GwBasePrice = &v
}

// GetGwPrice returns the GwPrice field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwPrice() string {
	if o == nil || IsNil(o.GwPrice) {
		var ret string
		return ret
	}
	return *o.GwPrice
}

// GetGwPriceOk returns a tuple with the GwPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GwPrice) {
		return nil, false
	}
	return o.GwPrice, true
}

// HasGwPrice returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwPrice() bool {
	if o != nil && !IsNil(o.GwPrice) {
		return true
	}

	return false
}

// SetGwPrice gets a reference to the given string and assigns it to the GwPrice field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwPrice(v string) {
	o.GwPrice = &v
}

// GetGwItemsBasePrice returns the GwItemsBasePrice field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsBasePrice() string {
	if o == nil || IsNil(o.GwItemsBasePrice) {
		var ret string
		return ret
	}
	return *o.GwItemsBasePrice
}

// GetGwItemsBasePriceOk returns a tuple with the GwItemsBasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsBasePriceOk() (*string, bool) {
	if o == nil || IsNil(o.GwItemsBasePrice) {
		return nil, false
	}
	return o.GwItemsBasePrice, true
}

// HasGwItemsBasePrice returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwItemsBasePrice() bool {
	if o != nil && !IsNil(o.GwItemsBasePrice) {
		return true
	}

	return false
}

// SetGwItemsBasePrice gets a reference to the given string and assigns it to the GwItemsBasePrice field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwItemsBasePrice(v string) {
	o.GwItemsBasePrice = &v
}

// GetGwItemsPrice returns the GwItemsPrice field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsPrice() string {
	if o == nil || IsNil(o.GwItemsPrice) {
		var ret string
		return ret
	}
	return *o.GwItemsPrice
}

// GetGwItemsPriceOk returns a tuple with the GwItemsPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GwItemsPrice) {
		return nil, false
	}
	return o.GwItemsPrice, true
}

// HasGwItemsPrice returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwItemsPrice() bool {
	if o != nil && !IsNil(o.GwItemsPrice) {
		return true
	}

	return false
}

// SetGwItemsPrice gets a reference to the given string and assigns it to the GwItemsPrice field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwItemsPrice(v string) {
	o.GwItemsPrice = &v
}

// GetGwCardBasePrice returns the GwCardBasePrice field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardBasePrice() string {
	if o == nil || IsNil(o.GwCardBasePrice) {
		var ret string
		return ret
	}
	return *o.GwCardBasePrice
}

// GetGwCardBasePriceOk returns a tuple with the GwCardBasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardBasePriceOk() (*string, bool) {
	if o == nil || IsNil(o.GwCardBasePrice) {
		return nil, false
	}
	return o.GwCardBasePrice, true
}

// HasGwCardBasePrice returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwCardBasePrice() bool {
	if o != nil && !IsNil(o.GwCardBasePrice) {
		return true
	}

	return false
}

// SetGwCardBasePrice gets a reference to the given string and assigns it to the GwCardBasePrice field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwCardBasePrice(v string) {
	o.GwCardBasePrice = &v
}

// GetGwCardPrice returns the GwCardPrice field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardPrice() string {
	if o == nil || IsNil(o.GwCardPrice) {
		var ret string
		return ret
	}
	return *o.GwCardPrice
}

// GetGwCardPriceOk returns a tuple with the GwCardPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GwCardPrice) {
		return nil, false
	}
	return o.GwCardPrice, true
}

// HasGwCardPrice returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwCardPrice() bool {
	if o != nil && !IsNil(o.GwCardPrice) {
		return true
	}

	return false
}

// SetGwCardPrice gets a reference to the given string and assigns it to the GwCardPrice field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwCardPrice(v string) {
	o.GwCardPrice = &v
}

// GetGwBaseTaxAmount returns the GwBaseTaxAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwBaseTaxAmount() string {
	if o == nil || IsNil(o.GwBaseTaxAmount) {
		var ret string
		return ret
	}
	return *o.GwBaseTaxAmount
}

// GetGwBaseTaxAmountOk returns a tuple with the GwBaseTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwBaseTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.GwBaseTaxAmount) {
		return nil, false
	}
	return o.GwBaseTaxAmount, true
}

// HasGwBaseTaxAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwBaseTaxAmount() bool {
	if o != nil && !IsNil(o.GwBaseTaxAmount) {
		return true
	}

	return false
}

// SetGwBaseTaxAmount gets a reference to the given string and assigns it to the GwBaseTaxAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwBaseTaxAmount(v string) {
	o.GwBaseTaxAmount = &v
}

// GetGwTaxAmount returns the GwTaxAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwTaxAmount() string {
	if o == nil || IsNil(o.GwTaxAmount) {
		var ret string
		return ret
	}
	return *o.GwTaxAmount
}

// GetGwTaxAmountOk returns a tuple with the GwTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.GwTaxAmount) {
		return nil, false
	}
	return o.GwTaxAmount, true
}

// HasGwTaxAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwTaxAmount() bool {
	if o != nil && !IsNil(o.GwTaxAmount) {
		return true
	}

	return false
}

// SetGwTaxAmount gets a reference to the given string and assigns it to the GwTaxAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwTaxAmount(v string) {
	o.GwTaxAmount = &v
}

// GetGwItemsBaseTaxAmount returns the GwItemsBaseTaxAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsBaseTaxAmount() string {
	if o == nil || IsNil(o.GwItemsBaseTaxAmount) {
		var ret string
		return ret
	}
	return *o.GwItemsBaseTaxAmount
}

// GetGwItemsBaseTaxAmountOk returns a tuple with the GwItemsBaseTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsBaseTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.GwItemsBaseTaxAmount) {
		return nil, false
	}
	return o.GwItemsBaseTaxAmount, true
}

// HasGwItemsBaseTaxAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwItemsBaseTaxAmount() bool {
	if o != nil && !IsNil(o.GwItemsBaseTaxAmount) {
		return true
	}

	return false
}

// SetGwItemsBaseTaxAmount gets a reference to the given string and assigns it to the GwItemsBaseTaxAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwItemsBaseTaxAmount(v string) {
	o.GwItemsBaseTaxAmount = &v
}

// GetGwItemsTaxAmount returns the GwItemsTaxAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsTaxAmount() string {
	if o == nil || IsNil(o.GwItemsTaxAmount) {
		var ret string
		return ret
	}
	return *o.GwItemsTaxAmount
}

// GetGwItemsTaxAmountOk returns a tuple with the GwItemsTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwItemsTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.GwItemsTaxAmount) {
		return nil, false
	}
	return o.GwItemsTaxAmount, true
}

// HasGwItemsTaxAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwItemsTaxAmount() bool {
	if o != nil && !IsNil(o.GwItemsTaxAmount) {
		return true
	}

	return false
}

// SetGwItemsTaxAmount gets a reference to the given string and assigns it to the GwItemsTaxAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwItemsTaxAmount(v string) {
	o.GwItemsTaxAmount = &v
}

// GetGwCardBaseTaxAmount returns the GwCardBaseTaxAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardBaseTaxAmount() string {
	if o == nil || IsNil(o.GwCardBaseTaxAmount) {
		var ret string
		return ret
	}
	return *o.GwCardBaseTaxAmount
}

// GetGwCardBaseTaxAmountOk returns a tuple with the GwCardBaseTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardBaseTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.GwCardBaseTaxAmount) {
		return nil, false
	}
	return o.GwCardBaseTaxAmount, true
}

// HasGwCardBaseTaxAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwCardBaseTaxAmount() bool {
	if o != nil && !IsNil(o.GwCardBaseTaxAmount) {
		return true
	}

	return false
}

// SetGwCardBaseTaxAmount gets a reference to the given string and assigns it to the GwCardBaseTaxAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwCardBaseTaxAmount(v string) {
	o.GwCardBaseTaxAmount = &v
}

// GetGwCardTaxAmount returns the GwCardTaxAmount field value if set, zero value otherwise.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardTaxAmount() string {
	if o == nil || IsNil(o.GwCardTaxAmount) {
		var ret string
		return ret
	}
	return *o.GwCardTaxAmount
}

// GetGwCardTaxAmountOk returns a tuple with the GwCardTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataCreditmemoExtensionInterface) GetGwCardTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.GwCardTaxAmount) {
		return nil, false
	}
	return o.GwCardTaxAmount, true
}

// HasGwCardTaxAmount returns a boolean if a field has been set.
func (o *SalesDataCreditmemoExtensionInterface) HasGwCardTaxAmount() bool {
	if o != nil && !IsNil(o.GwCardTaxAmount) {
		return true
	}

	return false
}

// SetGwCardTaxAmount gets a reference to the given string and assigns it to the GwCardTaxAmount field.
func (o *SalesDataCreditmemoExtensionInterface) SetGwCardTaxAmount(v string) {
	o.GwCardTaxAmount = &v
}

func (o SalesDataCreditmemoExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataCreditmemoExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseCustomerBalanceAmount) {
		toSerialize["base_customer_balance_amount"] = o.BaseCustomerBalanceAmount
	}
	if !IsNil(o.CustomerBalanceAmount) {
		toSerialize["customer_balance_amount"] = o.CustomerBalanceAmount
	}
	if !IsNil(o.BaseGiftCardsAmount) {
		toSerialize["base_gift_cards_amount"] = o.BaseGiftCardsAmount
	}
	if !IsNil(o.GiftCardsAmount) {
		toSerialize["gift_cards_amount"] = o.GiftCardsAmount
	}
	if !IsNil(o.GwBasePrice) {
		toSerialize["gw_base_price"] = o.GwBasePrice
	}
	if !IsNil(o.GwPrice) {
		toSerialize["gw_price"] = o.GwPrice
	}
	if !IsNil(o.GwItemsBasePrice) {
		toSerialize["gw_items_base_price"] = o.GwItemsBasePrice
	}
	if !IsNil(o.GwItemsPrice) {
		toSerialize["gw_items_price"] = o.GwItemsPrice
	}
	if !IsNil(o.GwCardBasePrice) {
		toSerialize["gw_card_base_price"] = o.GwCardBasePrice
	}
	if !IsNil(o.GwCardPrice) {
		toSerialize["gw_card_price"] = o.GwCardPrice
	}
	if !IsNil(o.GwBaseTaxAmount) {
		toSerialize["gw_base_tax_amount"] = o.GwBaseTaxAmount
	}
	if !IsNil(o.GwTaxAmount) {
		toSerialize["gw_tax_amount"] = o.GwTaxAmount
	}
	if !IsNil(o.GwItemsBaseTaxAmount) {
		toSerialize["gw_items_base_tax_amount"] = o.GwItemsBaseTaxAmount
	}
	if !IsNil(o.GwItemsTaxAmount) {
		toSerialize["gw_items_tax_amount"] = o.GwItemsTaxAmount
	}
	if !IsNil(o.GwCardBaseTaxAmount) {
		toSerialize["gw_card_base_tax_amount"] = o.GwCardBaseTaxAmount
	}
	if !IsNil(o.GwCardTaxAmount) {
		toSerialize["gw_card_tax_amount"] = o.GwCardTaxAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataCreditmemoExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varSalesDataCreditmemoExtensionInterface := _SalesDataCreditmemoExtensionInterface{}

	err = json.Unmarshal(data, &varSalesDataCreditmemoExtensionInterface)

	if err != nil {
		return err
	}

	*o = SalesDataCreditmemoExtensionInterface(varSalesDataCreditmemoExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base_customer_balance_amount")
		delete(additionalProperties, "customer_balance_amount")
		delete(additionalProperties, "base_gift_cards_amount")
		delete(additionalProperties, "gift_cards_amount")
		delete(additionalProperties, "gw_base_price")
		delete(additionalProperties, "gw_price")
		delete(additionalProperties, "gw_items_base_price")
		delete(additionalProperties, "gw_items_price")
		delete(additionalProperties, "gw_card_base_price")
		delete(additionalProperties, "gw_card_price")
		delete(additionalProperties, "gw_base_tax_amount")
		delete(additionalProperties, "gw_tax_amount")
		delete(additionalProperties, "gw_items_base_tax_amount")
		delete(additionalProperties, "gw_items_tax_amount")
		delete(additionalProperties, "gw_card_base_tax_amount")
		delete(additionalProperties, "gw_card_tax_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataCreditmemoExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataCreditmemoExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataCreditmemoExtensionInterface struct {
	value *SalesDataCreditmemoExtensionInterface
	isSet bool
}

func (v NullableSalesDataCreditmemoExtensionInterface) Get() *SalesDataCreditmemoExtensionInterface {
	return v.value
}

func (v *NullableSalesDataCreditmemoExtensionInterface) Set(val *SalesDataCreditmemoExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataCreditmemoExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataCreditmemoExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataCreditmemoExtensionInterface(val *SalesDataCreditmemoExtensionInterface) *NullableSalesDataCreditmemoExtensionInterface {
	return &NullableSalesDataCreditmemoExtensionInterface{value: val, isSet: true}
}

func (v NullableSalesDataCreditmemoExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataCreditmemoExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
