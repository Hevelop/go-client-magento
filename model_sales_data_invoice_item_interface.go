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

// checks if the SalesDataInvoiceItemInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataInvoiceItemInterface{}

// SalesDataInvoiceItemInterface Invoice item interface. An invoice is a record of the receipt of payment for an order. An invoice item is a purchased item in an invoice.
type SalesDataInvoiceItemInterface struct {
	// Additional data.
	AdditionalData *string `json:"additional_data,omitempty"`
	// Base cost.
	BaseCost *float32 `json:"base_cost,omitempty"`
	// Base discount amount.
	BaseDiscountAmount *float32 `json:"base_discount_amount,omitempty"`
	// Base discount tax compensation amount.
	BaseDiscountTaxCompensationAmount *float32 `json:"base_discount_tax_compensation_amount,omitempty"`
	// Base price.
	BasePrice *float32 `json:"base_price,omitempty"`
	// Base price including tax.
	BasePriceInclTax *float32 `json:"base_price_incl_tax,omitempty"`
	// Base row total.
	BaseRowTotal *float32 `json:"base_row_total,omitempty"`
	// Base row total including tax.
	BaseRowTotalInclTax *float32 `json:"base_row_total_incl_tax,omitempty"`
	// Base tax amount.
	BaseTaxAmount *float32 `json:"base_tax_amount,omitempty"`
	// Description.
	Description *string `json:"description,omitempty"`
	// Discount amount.
	DiscountAmount *float32 `json:"discount_amount,omitempty"`
	// Invoice item ID.
	EntityId *int32 `json:"entity_id,omitempty"`
	// Discount tax compensation amount.
	DiscountTaxCompensationAmount *float32 `json:"discount_tax_compensation_amount,omitempty"`
	// Name.
	Name *string `json:"name,omitempty"`
	// Parent ID.
	ParentId *int32 `json:"parent_id,omitempty"`
	// Price.
	Price *float32 `json:"price,omitempty"`
	// Price including tax.
	PriceInclTax *float32 `json:"price_incl_tax,omitempty"`
	// Product ID.
	ProductId *int32 `json:"product_id,omitempty"`
	// Row total.
	RowTotal *float32 `json:"row_total,omitempty"`
	// Row total including tax.
	RowTotalInclTax *float32 `json:"row_total_incl_tax,omitempty"`
	// SKU.
	Sku string `json:"sku"`
	// Tax amount.
	TaxAmount *float32 `json:"tax_amount,omitempty"`
	// ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\InvoiceItemInterface
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
	// Order item ID.
	OrderItemId int32 `json:"order_item_id"`
	// Quantity.
	Qty                  float32 `json:"qty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataInvoiceItemInterface SalesDataInvoiceItemInterface

// NewSalesDataInvoiceItemInterface instantiates a new SalesDataInvoiceItemInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataInvoiceItemInterface(sku string, orderItemId int32, qty float32) *SalesDataInvoiceItemInterface {
	this := SalesDataInvoiceItemInterface{}
	this.Sku = sku
	this.OrderItemId = orderItemId
	this.Qty = qty
	return &this
}

// NewSalesDataInvoiceItemInterfaceWithDefaults instantiates a new SalesDataInvoiceItemInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataInvoiceItemInterfaceWithDefaults() *SalesDataInvoiceItemInterface {
	this := SalesDataInvoiceItemInterface{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetAdditionalData() string {
	if o == nil || IsNil(o.AdditionalData) {
		var ret string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetAdditionalDataOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given string and assigns it to the AdditionalData field.
func (o *SalesDataInvoiceItemInterface) SetAdditionalData(v string) {
	o.AdditionalData = &v
}

// GetBaseCost returns the BaseCost field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBaseCost() float32 {
	if o == nil || IsNil(o.BaseCost) {
		var ret float32
		return ret
	}
	return *o.BaseCost
}

// GetBaseCostOk returns a tuple with the BaseCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBaseCostOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseCost) {
		return nil, false
	}
	return o.BaseCost, true
}

// HasBaseCost returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBaseCost() bool {
	if o != nil && !IsNil(o.BaseCost) {
		return true
	}

	return false
}

// SetBaseCost gets a reference to the given float32 and assigns it to the BaseCost field.
func (o *SalesDataInvoiceItemInterface) SetBaseCost(v float32) {
	o.BaseCost = &v
}

// GetBaseDiscountAmount returns the BaseDiscountAmount field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBaseDiscountAmount() float32 {
	if o == nil || IsNil(o.BaseDiscountAmount) {
		var ret float32
		return ret
	}
	return *o.BaseDiscountAmount
}

// GetBaseDiscountAmountOk returns a tuple with the BaseDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBaseDiscountAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseDiscountAmount) {
		return nil, false
	}
	return o.BaseDiscountAmount, true
}

// HasBaseDiscountAmount returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBaseDiscountAmount() bool {
	if o != nil && !IsNil(o.BaseDiscountAmount) {
		return true
	}

	return false
}

// SetBaseDiscountAmount gets a reference to the given float32 and assigns it to the BaseDiscountAmount field.
func (o *SalesDataInvoiceItemInterface) SetBaseDiscountAmount(v float32) {
	o.BaseDiscountAmount = &v
}

// GetBaseDiscountTaxCompensationAmount returns the BaseDiscountTaxCompensationAmount field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBaseDiscountTaxCompensationAmount() float32 {
	if o == nil || IsNil(o.BaseDiscountTaxCompensationAmount) {
		var ret float32
		return ret
	}
	return *o.BaseDiscountTaxCompensationAmount
}

// GetBaseDiscountTaxCompensationAmountOk returns a tuple with the BaseDiscountTaxCompensationAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBaseDiscountTaxCompensationAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseDiscountTaxCompensationAmount) {
		return nil, false
	}
	return o.BaseDiscountTaxCompensationAmount, true
}

// HasBaseDiscountTaxCompensationAmount returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBaseDiscountTaxCompensationAmount() bool {
	if o != nil && !IsNil(o.BaseDiscountTaxCompensationAmount) {
		return true
	}

	return false
}

// SetBaseDiscountTaxCompensationAmount gets a reference to the given float32 and assigns it to the BaseDiscountTaxCompensationAmount field.
func (o *SalesDataInvoiceItemInterface) SetBaseDiscountTaxCompensationAmount(v float32) {
	o.BaseDiscountTaxCompensationAmount = &v
}

// GetBasePrice returns the BasePrice field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBasePrice() float32 {
	if o == nil || IsNil(o.BasePrice) {
		var ret float32
		return ret
	}
	return *o.BasePrice
}

// GetBasePriceOk returns a tuple with the BasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBasePriceOk() (*float32, bool) {
	if o == nil || IsNil(o.BasePrice) {
		return nil, false
	}
	return o.BasePrice, true
}

// HasBasePrice returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBasePrice() bool {
	if o != nil && !IsNil(o.BasePrice) {
		return true
	}

	return false
}

// SetBasePrice gets a reference to the given float32 and assigns it to the BasePrice field.
func (o *SalesDataInvoiceItemInterface) SetBasePrice(v float32) {
	o.BasePrice = &v
}

// GetBasePriceInclTax returns the BasePriceInclTax field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBasePriceInclTax() float32 {
	if o == nil || IsNil(o.BasePriceInclTax) {
		var ret float32
		return ret
	}
	return *o.BasePriceInclTax
}

// GetBasePriceInclTaxOk returns a tuple with the BasePriceInclTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBasePriceInclTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.BasePriceInclTax) {
		return nil, false
	}
	return o.BasePriceInclTax, true
}

// HasBasePriceInclTax returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBasePriceInclTax() bool {
	if o != nil && !IsNil(o.BasePriceInclTax) {
		return true
	}

	return false
}

// SetBasePriceInclTax gets a reference to the given float32 and assigns it to the BasePriceInclTax field.
func (o *SalesDataInvoiceItemInterface) SetBasePriceInclTax(v float32) {
	o.BasePriceInclTax = &v
}

// GetBaseRowTotal returns the BaseRowTotal field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBaseRowTotal() float32 {
	if o == nil || IsNil(o.BaseRowTotal) {
		var ret float32
		return ret
	}
	return *o.BaseRowTotal
}

// GetBaseRowTotalOk returns a tuple with the BaseRowTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBaseRowTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseRowTotal) {
		return nil, false
	}
	return o.BaseRowTotal, true
}

// HasBaseRowTotal returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBaseRowTotal() bool {
	if o != nil && !IsNil(o.BaseRowTotal) {
		return true
	}

	return false
}

// SetBaseRowTotal gets a reference to the given float32 and assigns it to the BaseRowTotal field.
func (o *SalesDataInvoiceItemInterface) SetBaseRowTotal(v float32) {
	o.BaseRowTotal = &v
}

// GetBaseRowTotalInclTax returns the BaseRowTotalInclTax field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBaseRowTotalInclTax() float32 {
	if o == nil || IsNil(o.BaseRowTotalInclTax) {
		var ret float32
		return ret
	}
	return *o.BaseRowTotalInclTax
}

// GetBaseRowTotalInclTaxOk returns a tuple with the BaseRowTotalInclTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBaseRowTotalInclTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseRowTotalInclTax) {
		return nil, false
	}
	return o.BaseRowTotalInclTax, true
}

// HasBaseRowTotalInclTax returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBaseRowTotalInclTax() bool {
	if o != nil && !IsNil(o.BaseRowTotalInclTax) {
		return true
	}

	return false
}

// SetBaseRowTotalInclTax gets a reference to the given float32 and assigns it to the BaseRowTotalInclTax field.
func (o *SalesDataInvoiceItemInterface) SetBaseRowTotalInclTax(v float32) {
	o.BaseRowTotalInclTax = &v
}

// GetBaseTaxAmount returns the BaseTaxAmount field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetBaseTaxAmount() float32 {
	if o == nil || IsNil(o.BaseTaxAmount) {
		var ret float32
		return ret
	}
	return *o.BaseTaxAmount
}

// GetBaseTaxAmountOk returns a tuple with the BaseTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetBaseTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseTaxAmount) {
		return nil, false
	}
	return o.BaseTaxAmount, true
}

// HasBaseTaxAmount returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasBaseTaxAmount() bool {
	if o != nil && !IsNil(o.BaseTaxAmount) {
		return true
	}

	return false
}

// SetBaseTaxAmount gets a reference to the given float32 and assigns it to the BaseTaxAmount field.
func (o *SalesDataInvoiceItemInterface) SetBaseTaxAmount(v float32) {
	o.BaseTaxAmount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SalesDataInvoiceItemInterface) SetDescription(v string) {
	o.Description = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetDiscountAmount() float32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret float32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetDiscountAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given float32 and assigns it to the DiscountAmount field.
func (o *SalesDataInvoiceItemInterface) SetDiscountAmount(v float32) {
	o.DiscountAmount = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetEntityId() int32 {
	if o == nil || IsNil(o.EntityId) {
		var ret int32
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given int32 and assigns it to the EntityId field.
func (o *SalesDataInvoiceItemInterface) SetEntityId(v int32) {
	o.EntityId = &v
}

// GetDiscountTaxCompensationAmount returns the DiscountTaxCompensationAmount field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetDiscountTaxCompensationAmount() float32 {
	if o == nil || IsNil(o.DiscountTaxCompensationAmount) {
		var ret float32
		return ret
	}
	return *o.DiscountTaxCompensationAmount
}

// GetDiscountTaxCompensationAmountOk returns a tuple with the DiscountTaxCompensationAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetDiscountTaxCompensationAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountTaxCompensationAmount) {
		return nil, false
	}
	return o.DiscountTaxCompensationAmount, true
}

// HasDiscountTaxCompensationAmount returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasDiscountTaxCompensationAmount() bool {
	if o != nil && !IsNil(o.DiscountTaxCompensationAmount) {
		return true
	}

	return false
}

// SetDiscountTaxCompensationAmount gets a reference to the given float32 and assigns it to the DiscountTaxCompensationAmount field.
func (o *SalesDataInvoiceItemInterface) SetDiscountTaxCompensationAmount(v float32) {
	o.DiscountTaxCompensationAmount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SalesDataInvoiceItemInterface) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *SalesDataInvoiceItemInterface) SetParentId(v int32) {
	o.ParentId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *SalesDataInvoiceItemInterface) SetPrice(v float32) {
	o.Price = &v
}

// GetPriceInclTax returns the PriceInclTax field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetPriceInclTax() float32 {
	if o == nil || IsNil(o.PriceInclTax) {
		var ret float32
		return ret
	}
	return *o.PriceInclTax
}

// GetPriceInclTaxOk returns a tuple with the PriceInclTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetPriceInclTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceInclTax) {
		return nil, false
	}
	return o.PriceInclTax, true
}

// HasPriceInclTax returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasPriceInclTax() bool {
	if o != nil && !IsNil(o.PriceInclTax) {
		return true
	}

	return false
}

// SetPriceInclTax gets a reference to the given float32 and assigns it to the PriceInclTax field.
func (o *SalesDataInvoiceItemInterface) SetPriceInclTax(v float32) {
	o.PriceInclTax = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId) {
		var ret int32
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetProductIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int32 and assigns it to the ProductId field.
func (o *SalesDataInvoiceItemInterface) SetProductId(v int32) {
	o.ProductId = &v
}

// GetRowTotal returns the RowTotal field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetRowTotal() float32 {
	if o == nil || IsNil(o.RowTotal) {
		var ret float32
		return ret
	}
	return *o.RowTotal
}

// GetRowTotalOk returns a tuple with the RowTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetRowTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.RowTotal) {
		return nil, false
	}
	return o.RowTotal, true
}

// HasRowTotal returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasRowTotal() bool {
	if o != nil && !IsNil(o.RowTotal) {
		return true
	}

	return false
}

// SetRowTotal gets a reference to the given float32 and assigns it to the RowTotal field.
func (o *SalesDataInvoiceItemInterface) SetRowTotal(v float32) {
	o.RowTotal = &v
}

// GetRowTotalInclTax returns the RowTotalInclTax field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetRowTotalInclTax() float32 {
	if o == nil || IsNil(o.RowTotalInclTax) {
		var ret float32
		return ret
	}
	return *o.RowTotalInclTax
}

// GetRowTotalInclTaxOk returns a tuple with the RowTotalInclTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetRowTotalInclTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.RowTotalInclTax) {
		return nil, false
	}
	return o.RowTotalInclTax, true
}

// HasRowTotalInclTax returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasRowTotalInclTax() bool {
	if o != nil && !IsNil(o.RowTotalInclTax) {
		return true
	}

	return false
}

// SetRowTotalInclTax gets a reference to the given float32 and assigns it to the RowTotalInclTax field.
func (o *SalesDataInvoiceItemInterface) SetRowTotalInclTax(v float32) {
	o.RowTotalInclTax = &v
}

// GetSku returns the Sku field value
func (o *SalesDataInvoiceItemInterface) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *SalesDataInvoiceItemInterface) SetSku(v string) {
	o.Sku = v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetTaxAmount() float32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float32 and assigns it to the TaxAmount field.
func (o *SalesDataInvoiceItemInterface) SetTaxAmount(v float32) {
	o.TaxAmount = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataInvoiceItemInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataInvoiceItemInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *SalesDataInvoiceItemInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

// GetOrderItemId returns the OrderItemId field value
func (o *SalesDataInvoiceItemInterface) GetOrderItemId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetOrderItemIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemId, true
}

// SetOrderItemId sets field value
func (o *SalesDataInvoiceItemInterface) SetOrderItemId(v int32) {
	o.OrderItemId = v
}

// GetQty returns the Qty field value
func (o *SalesDataInvoiceItemInterface) GetQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *SalesDataInvoiceItemInterface) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *SalesDataInvoiceItemInterface) SetQty(v float32) {
	o.Qty = v
}

func (o SalesDataInvoiceItemInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataInvoiceItemInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalData) {
		toSerialize["additional_data"] = o.AdditionalData
	}
	if !IsNil(o.BaseCost) {
		toSerialize["base_cost"] = o.BaseCost
	}
	if !IsNil(o.BaseDiscountAmount) {
		toSerialize["base_discount_amount"] = o.BaseDiscountAmount
	}
	if !IsNil(o.BaseDiscountTaxCompensationAmount) {
		toSerialize["base_discount_tax_compensation_amount"] = o.BaseDiscountTaxCompensationAmount
	}
	if !IsNil(o.BasePrice) {
		toSerialize["base_price"] = o.BasePrice
	}
	if !IsNil(o.BasePriceInclTax) {
		toSerialize["base_price_incl_tax"] = o.BasePriceInclTax
	}
	if !IsNil(o.BaseRowTotal) {
		toSerialize["base_row_total"] = o.BaseRowTotal
	}
	if !IsNil(o.BaseRowTotalInclTax) {
		toSerialize["base_row_total_incl_tax"] = o.BaseRowTotalInclTax
	}
	if !IsNil(o.BaseTaxAmount) {
		toSerialize["base_tax_amount"] = o.BaseTaxAmount
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discount_amount"] = o.DiscountAmount
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.DiscountTaxCompensationAmount) {
		toSerialize["discount_tax_compensation_amount"] = o.DiscountTaxCompensationAmount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceInclTax) {
		toSerialize["price_incl_tax"] = o.PriceInclTax
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.RowTotal) {
		toSerialize["row_total"] = o.RowTotal
	}
	if !IsNil(o.RowTotalInclTax) {
		toSerialize["row_total_incl_tax"] = o.RowTotalInclTax
	}
	toSerialize["sku"] = o.Sku
	if !IsNil(o.TaxAmount) {
		toSerialize["tax_amount"] = o.TaxAmount
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}
	toSerialize["order_item_id"] = o.OrderItemId
	toSerialize["qty"] = o.Qty

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataInvoiceItemInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku",
		"order_item_id",
		"qty",
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

	varSalesDataInvoiceItemInterface := _SalesDataInvoiceItemInterface{}

	err = json.Unmarshal(data, &varSalesDataInvoiceItemInterface)

	if err != nil {
		return err
	}

	*o = SalesDataInvoiceItemInterface(varSalesDataInvoiceItemInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_data")
		delete(additionalProperties, "base_cost")
		delete(additionalProperties, "base_discount_amount")
		delete(additionalProperties, "base_discount_tax_compensation_amount")
		delete(additionalProperties, "base_price")
		delete(additionalProperties, "base_price_incl_tax")
		delete(additionalProperties, "base_row_total")
		delete(additionalProperties, "base_row_total_incl_tax")
		delete(additionalProperties, "base_tax_amount")
		delete(additionalProperties, "description")
		delete(additionalProperties, "discount_amount")
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "discount_tax_compensation_amount")
		delete(additionalProperties, "name")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "price_incl_tax")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "row_total")
		delete(additionalProperties, "row_total_incl_tax")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "tax_amount")
		delete(additionalProperties, "extension_attributes")
		delete(additionalProperties, "order_item_id")
		delete(additionalProperties, "qty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataInvoiceItemInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataInvoiceItemInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataInvoiceItemInterface struct {
	value *SalesDataInvoiceItemInterface
	isSet bool
}

func (v NullableSalesDataInvoiceItemInterface) Get() *SalesDataInvoiceItemInterface {
	return v.value
}

func (v *NullableSalesDataInvoiceItemInterface) Set(val *SalesDataInvoiceItemInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataInvoiceItemInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataInvoiceItemInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataInvoiceItemInterface(val *SalesDataInvoiceItemInterface) *NullableSalesDataInvoiceItemInterface {
	return &NullableSalesDataInvoiceItemInterface{value: val, isSet: true}
}

func (v NullableSalesDataInvoiceItemInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataInvoiceItemInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}