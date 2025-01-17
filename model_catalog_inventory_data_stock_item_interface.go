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

// checks if the CatalogInventoryDataStockItemInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogInventoryDataStockItemInterface{}

// CatalogInventoryDataStockItemInterface Interface StockItem
type CatalogInventoryDataStockItemInterface struct {
	ItemId    *int32 `json:"item_id,omitempty"`
	ProductId *int32 `json:"product_id,omitempty"`
	// Stock identifier
	StockId *int32  `json:"stock_id,omitempty"`
	Qty     float32 `json:"qty"`
	// Stock Availability
	IsInStock                      bool `json:"is_in_stock"`
	IsQtyDecimal                   bool `json:"is_qty_decimal"`
	ShowDefaultNotificationMessage bool `json:"show_default_notification_message"`
	UseConfigMinQty                bool `json:"use_config_min_qty"`
	// Minimal quantity available for item status in stock
	MinQty              float32 `json:"min_qty"`
	UseConfigMinSaleQty int32   `json:"use_config_min_sale_qty"`
	// Minimum Qty Allowed in Shopping Cart or NULL when there is no limitation
	MinSaleQty          float32 `json:"min_sale_qty"`
	UseConfigMaxSaleQty bool    `json:"use_config_max_sale_qty"`
	// Maximum Qty Allowed in Shopping Cart data wrapper
	MaxSaleQty          float32 `json:"max_sale_qty"`
	UseConfigBackorders bool    `json:"use_config_backorders"`
	// Backorders status
	Backorders              int32 `json:"backorders"`
	UseConfigNotifyStockQty bool  `json:"use_config_notify_stock_qty"`
	// Notify for Quantity Below data wrapper
	NotifyStockQty         float32 `json:"notify_stock_qty"`
	UseConfigQtyIncrements bool    `json:"use_config_qty_increments"`
	// Quantity Increments data wrapper
	QtyIncrements         float32 `json:"qty_increments"`
	UseConfigEnableQtyInc bool    `json:"use_config_enable_qty_inc"`
	// Whether Quantity Increments is enabled
	EnableQtyIncrements  bool `json:"enable_qty_increments"`
	UseConfigManageStock bool `json:"use_config_manage_stock"`
	// Can Manage Stock
	ManageStock            bool   `json:"manage_stock"`
	LowStockDate           string `json:"low_stock_date"`
	IsDecimalDivided       bool   `json:"is_decimal_divided"`
	StockStatusChangedAuto int32  `json:"stock_status_changed_auto"`
	// ExtensionInterface class for @see \\Magento\\CatalogInventory\\Api\\Data\\StockItemInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogInventoryDataStockItemInterface CatalogInventoryDataStockItemInterface

// NewCatalogInventoryDataStockItemInterface instantiates a new CatalogInventoryDataStockItemInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogInventoryDataStockItemInterface(qty float32, isInStock bool, isQtyDecimal bool, showDefaultNotificationMessage bool, useConfigMinQty bool, minQty float32, useConfigMinSaleQty int32, minSaleQty float32, useConfigMaxSaleQty bool, maxSaleQty float32, useConfigBackorders bool, backorders int32, useConfigNotifyStockQty bool, notifyStockQty float32, useConfigQtyIncrements bool, qtyIncrements float32, useConfigEnableQtyInc bool, enableQtyIncrements bool, useConfigManageStock bool, manageStock bool, lowStockDate string, isDecimalDivided bool, stockStatusChangedAuto int32) *CatalogInventoryDataStockItemInterface {
	this := CatalogInventoryDataStockItemInterface{}
	this.Qty = qty
	this.IsInStock = isInStock
	this.IsQtyDecimal = isQtyDecimal
	this.ShowDefaultNotificationMessage = showDefaultNotificationMessage
	this.UseConfigMinQty = useConfigMinQty
	this.MinQty = minQty
	this.UseConfigMinSaleQty = useConfigMinSaleQty
	this.MinSaleQty = minSaleQty
	this.UseConfigMaxSaleQty = useConfigMaxSaleQty
	this.MaxSaleQty = maxSaleQty
	this.UseConfigBackorders = useConfigBackorders
	this.Backorders = backorders
	this.UseConfigNotifyStockQty = useConfigNotifyStockQty
	this.NotifyStockQty = notifyStockQty
	this.UseConfigQtyIncrements = useConfigQtyIncrements
	this.QtyIncrements = qtyIncrements
	this.UseConfigEnableQtyInc = useConfigEnableQtyInc
	this.EnableQtyIncrements = enableQtyIncrements
	this.UseConfigManageStock = useConfigManageStock
	this.ManageStock = manageStock
	this.LowStockDate = lowStockDate
	this.IsDecimalDivided = isDecimalDivided
	this.StockStatusChangedAuto = stockStatusChangedAuto
	return &this
}

// NewCatalogInventoryDataStockItemInterfaceWithDefaults instantiates a new CatalogInventoryDataStockItemInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogInventoryDataStockItemInterfaceWithDefaults() *CatalogInventoryDataStockItemInterface {
	this := CatalogInventoryDataStockItemInterface{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *CatalogInventoryDataStockItemInterface) GetItemId() int32 {
	if o == nil || IsNil(o.ItemId) {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetItemIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *CatalogInventoryDataStockItemInterface) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *CatalogInventoryDataStockItemInterface) SetItemId(v int32) {
	o.ItemId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *CatalogInventoryDataStockItemInterface) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId) {
		var ret int32
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetProductIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *CatalogInventoryDataStockItemInterface) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int32 and assigns it to the ProductId field.
func (o *CatalogInventoryDataStockItemInterface) SetProductId(v int32) {
	o.ProductId = &v
}

// GetStockId returns the StockId field value if set, zero value otherwise.
func (o *CatalogInventoryDataStockItemInterface) GetStockId() int32 {
	if o == nil || IsNil(o.StockId) {
		var ret int32
		return ret
	}
	return *o.StockId
}

// GetStockIdOk returns a tuple with the StockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetStockIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StockId) {
		return nil, false
	}
	return o.StockId, true
}

// HasStockId returns a boolean if a field has been set.
func (o *CatalogInventoryDataStockItemInterface) HasStockId() bool {
	if o != nil && !IsNil(o.StockId) {
		return true
	}

	return false
}

// SetStockId gets a reference to the given int32 and assigns it to the StockId field.
func (o *CatalogInventoryDataStockItemInterface) SetStockId(v int32) {
	o.StockId = &v
}

// GetQty returns the Qty field value
func (o *CatalogInventoryDataStockItemInterface) GetQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetQty(v float32) {
	o.Qty = v
}

// GetIsInStock returns the IsInStock field value
func (o *CatalogInventoryDataStockItemInterface) GetIsInStock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInStock
}

// GetIsInStockOk returns a tuple with the IsInStock field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetIsInStockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInStock, true
}

// SetIsInStock sets field value
func (o *CatalogInventoryDataStockItemInterface) SetIsInStock(v bool) {
	o.IsInStock = v
}

// GetIsQtyDecimal returns the IsQtyDecimal field value
func (o *CatalogInventoryDataStockItemInterface) GetIsQtyDecimal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsQtyDecimal
}

// GetIsQtyDecimalOk returns a tuple with the IsQtyDecimal field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetIsQtyDecimalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsQtyDecimal, true
}

// SetIsQtyDecimal sets field value
func (o *CatalogInventoryDataStockItemInterface) SetIsQtyDecimal(v bool) {
	o.IsQtyDecimal = v
}

// GetShowDefaultNotificationMessage returns the ShowDefaultNotificationMessage field value
func (o *CatalogInventoryDataStockItemInterface) GetShowDefaultNotificationMessage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowDefaultNotificationMessage
}

// GetShowDefaultNotificationMessageOk returns a tuple with the ShowDefaultNotificationMessage field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetShowDefaultNotificationMessageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowDefaultNotificationMessage, true
}

// SetShowDefaultNotificationMessage sets field value
func (o *CatalogInventoryDataStockItemInterface) SetShowDefaultNotificationMessage(v bool) {
	o.ShowDefaultNotificationMessage = v
}

// GetUseConfigMinQty returns the UseConfigMinQty field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigMinQty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseConfigMinQty
}

// GetUseConfigMinQtyOk returns a tuple with the UseConfigMinQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigMinQtyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigMinQty, true
}

// SetUseConfigMinQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigMinQty(v bool) {
	o.UseConfigMinQty = v
}

// GetMinQty returns the MinQty field value
func (o *CatalogInventoryDataStockItemInterface) GetMinQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetMinQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinQty, true
}

// SetMinQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetMinQty(v float32) {
	o.MinQty = v
}

// GetUseConfigMinSaleQty returns the UseConfigMinSaleQty field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigMinSaleQty() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UseConfigMinSaleQty
}

// GetUseConfigMinSaleQtyOk returns a tuple with the UseConfigMinSaleQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigMinSaleQtyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigMinSaleQty, true
}

// SetUseConfigMinSaleQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigMinSaleQty(v int32) {
	o.UseConfigMinSaleQty = v
}

// GetMinSaleQty returns the MinSaleQty field value
func (o *CatalogInventoryDataStockItemInterface) GetMinSaleQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinSaleQty
}

// GetMinSaleQtyOk returns a tuple with the MinSaleQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetMinSaleQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinSaleQty, true
}

// SetMinSaleQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetMinSaleQty(v float32) {
	o.MinSaleQty = v
}

// GetUseConfigMaxSaleQty returns the UseConfigMaxSaleQty field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigMaxSaleQty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseConfigMaxSaleQty
}

// GetUseConfigMaxSaleQtyOk returns a tuple with the UseConfigMaxSaleQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigMaxSaleQtyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigMaxSaleQty, true
}

// SetUseConfigMaxSaleQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigMaxSaleQty(v bool) {
	o.UseConfigMaxSaleQty = v
}

// GetMaxSaleQty returns the MaxSaleQty field value
func (o *CatalogInventoryDataStockItemInterface) GetMaxSaleQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxSaleQty
}

// GetMaxSaleQtyOk returns a tuple with the MaxSaleQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetMaxSaleQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSaleQty, true
}

// SetMaxSaleQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetMaxSaleQty(v float32) {
	o.MaxSaleQty = v
}

// GetUseConfigBackorders returns the UseConfigBackorders field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigBackorders() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseConfigBackorders
}

// GetUseConfigBackordersOk returns a tuple with the UseConfigBackorders field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigBackordersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigBackorders, true
}

// SetUseConfigBackorders sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigBackorders(v bool) {
	o.UseConfigBackorders = v
}

// GetBackorders returns the Backorders field value
func (o *CatalogInventoryDataStockItemInterface) GetBackorders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Backorders
}

// GetBackordersOk returns a tuple with the Backorders field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetBackordersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Backorders, true
}

// SetBackorders sets field value
func (o *CatalogInventoryDataStockItemInterface) SetBackorders(v int32) {
	o.Backorders = v
}

// GetUseConfigNotifyStockQty returns the UseConfigNotifyStockQty field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigNotifyStockQty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseConfigNotifyStockQty
}

// GetUseConfigNotifyStockQtyOk returns a tuple with the UseConfigNotifyStockQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigNotifyStockQtyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigNotifyStockQty, true
}

// SetUseConfigNotifyStockQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigNotifyStockQty(v bool) {
	o.UseConfigNotifyStockQty = v
}

// GetNotifyStockQty returns the NotifyStockQty field value
func (o *CatalogInventoryDataStockItemInterface) GetNotifyStockQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NotifyStockQty
}

// GetNotifyStockQtyOk returns a tuple with the NotifyStockQty field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetNotifyStockQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyStockQty, true
}

// SetNotifyStockQty sets field value
func (o *CatalogInventoryDataStockItemInterface) SetNotifyStockQty(v float32) {
	o.NotifyStockQty = v
}

// GetUseConfigQtyIncrements returns the UseConfigQtyIncrements field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigQtyIncrements() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseConfigQtyIncrements
}

// GetUseConfigQtyIncrementsOk returns a tuple with the UseConfigQtyIncrements field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigQtyIncrementsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigQtyIncrements, true
}

// SetUseConfigQtyIncrements sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigQtyIncrements(v bool) {
	o.UseConfigQtyIncrements = v
}

// GetQtyIncrements returns the QtyIncrements field value
func (o *CatalogInventoryDataStockItemInterface) GetQtyIncrements() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.QtyIncrements
}

// GetQtyIncrementsOk returns a tuple with the QtyIncrements field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetQtyIncrementsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyIncrements, true
}

// SetQtyIncrements sets field value
func (o *CatalogInventoryDataStockItemInterface) SetQtyIncrements(v float32) {
	o.QtyIncrements = v
}

// GetUseConfigEnableQtyInc returns the UseConfigEnableQtyInc field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigEnableQtyInc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseConfigEnableQtyInc
}

// GetUseConfigEnableQtyIncOk returns a tuple with the UseConfigEnableQtyInc field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigEnableQtyIncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigEnableQtyInc, true
}

// SetUseConfigEnableQtyInc sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigEnableQtyInc(v bool) {
	o.UseConfigEnableQtyInc = v
}

// GetEnableQtyIncrements returns the EnableQtyIncrements field value
func (o *CatalogInventoryDataStockItemInterface) GetEnableQtyIncrements() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableQtyIncrements
}

// GetEnableQtyIncrementsOk returns a tuple with the EnableQtyIncrements field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetEnableQtyIncrementsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableQtyIncrements, true
}

// SetEnableQtyIncrements sets field value
func (o *CatalogInventoryDataStockItemInterface) SetEnableQtyIncrements(v bool) {
	o.EnableQtyIncrements = v
}

// GetUseConfigManageStock returns the UseConfigManageStock field value
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigManageStock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseConfigManageStock
}

// GetUseConfigManageStockOk returns a tuple with the UseConfigManageStock field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetUseConfigManageStockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseConfigManageStock, true
}

// SetUseConfigManageStock sets field value
func (o *CatalogInventoryDataStockItemInterface) SetUseConfigManageStock(v bool) {
	o.UseConfigManageStock = v
}

// GetManageStock returns the ManageStock field value
func (o *CatalogInventoryDataStockItemInterface) GetManageStock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ManageStock
}

// GetManageStockOk returns a tuple with the ManageStock field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetManageStockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManageStock, true
}

// SetManageStock sets field value
func (o *CatalogInventoryDataStockItemInterface) SetManageStock(v bool) {
	o.ManageStock = v
}

// GetLowStockDate returns the LowStockDate field value
func (o *CatalogInventoryDataStockItemInterface) GetLowStockDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LowStockDate
}

// GetLowStockDateOk returns a tuple with the LowStockDate field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetLowStockDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowStockDate, true
}

// SetLowStockDate sets field value
func (o *CatalogInventoryDataStockItemInterface) SetLowStockDate(v string) {
	o.LowStockDate = v
}

// GetIsDecimalDivided returns the IsDecimalDivided field value
func (o *CatalogInventoryDataStockItemInterface) GetIsDecimalDivided() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDecimalDivided
}

// GetIsDecimalDividedOk returns a tuple with the IsDecimalDivided field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetIsDecimalDividedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDecimalDivided, true
}

// SetIsDecimalDivided sets field value
func (o *CatalogInventoryDataStockItemInterface) SetIsDecimalDivided(v bool) {
	o.IsDecimalDivided = v
}

// GetStockStatusChangedAuto returns the StockStatusChangedAuto field value
func (o *CatalogInventoryDataStockItemInterface) GetStockStatusChangedAuto() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StockStatusChangedAuto
}

// GetStockStatusChangedAutoOk returns a tuple with the StockStatusChangedAuto field value
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetStockStatusChangedAutoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockStatusChangedAuto, true
}

// SetStockStatusChangedAuto sets field value
func (o *CatalogInventoryDataStockItemInterface) SetStockStatusChangedAuto(v int32) {
	o.StockStatusChangedAuto = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *CatalogInventoryDataStockItemInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogInventoryDataStockItemInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *CatalogInventoryDataStockItemInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *CatalogInventoryDataStockItemInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o CatalogInventoryDataStockItemInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogInventoryDataStockItemInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.StockId) {
		toSerialize["stock_id"] = o.StockId
	}
	toSerialize["qty"] = o.Qty
	toSerialize["is_in_stock"] = o.IsInStock
	toSerialize["is_qty_decimal"] = o.IsQtyDecimal
	toSerialize["show_default_notification_message"] = o.ShowDefaultNotificationMessage
	toSerialize["use_config_min_qty"] = o.UseConfigMinQty
	toSerialize["min_qty"] = o.MinQty
	toSerialize["use_config_min_sale_qty"] = o.UseConfigMinSaleQty
	toSerialize["min_sale_qty"] = o.MinSaleQty
	toSerialize["use_config_max_sale_qty"] = o.UseConfigMaxSaleQty
	toSerialize["max_sale_qty"] = o.MaxSaleQty
	toSerialize["use_config_backorders"] = o.UseConfigBackorders
	toSerialize["backorders"] = o.Backorders
	toSerialize["use_config_notify_stock_qty"] = o.UseConfigNotifyStockQty
	toSerialize["notify_stock_qty"] = o.NotifyStockQty
	toSerialize["use_config_qty_increments"] = o.UseConfigQtyIncrements
	toSerialize["qty_increments"] = o.QtyIncrements
	toSerialize["use_config_enable_qty_inc"] = o.UseConfigEnableQtyInc
	toSerialize["enable_qty_increments"] = o.EnableQtyIncrements
	toSerialize["use_config_manage_stock"] = o.UseConfigManageStock
	toSerialize["manage_stock"] = o.ManageStock
	toSerialize["low_stock_date"] = o.LowStockDate
	toSerialize["is_decimal_divided"] = o.IsDecimalDivided
	toSerialize["stock_status_changed_auto"] = o.StockStatusChangedAuto
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogInventoryDataStockItemInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"qty",
		"is_in_stock",
		"is_qty_decimal",
		"show_default_notification_message",
		"use_config_min_qty",
		"min_qty",
		"use_config_min_sale_qty",
		"min_sale_qty",
		"use_config_max_sale_qty",
		"max_sale_qty",
		"use_config_backorders",
		"backorders",
		"use_config_notify_stock_qty",
		"notify_stock_qty",
		"use_config_qty_increments",
		"qty_increments",
		"use_config_enable_qty_inc",
		"enable_qty_increments",
		"use_config_manage_stock",
		"manage_stock",
		"low_stock_date",
		"is_decimal_divided",
		"stock_status_changed_auto",
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

	varCatalogInventoryDataStockItemInterface := _CatalogInventoryDataStockItemInterface{}

	err = json.Unmarshal(data, &varCatalogInventoryDataStockItemInterface)

	if err != nil {
		return err
	}

	*o = CatalogInventoryDataStockItemInterface(varCatalogInventoryDataStockItemInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "stock_id")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "is_in_stock")
		delete(additionalProperties, "is_qty_decimal")
		delete(additionalProperties, "show_default_notification_message")
		delete(additionalProperties, "use_config_min_qty")
		delete(additionalProperties, "min_qty")
		delete(additionalProperties, "use_config_min_sale_qty")
		delete(additionalProperties, "min_sale_qty")
		delete(additionalProperties, "use_config_max_sale_qty")
		delete(additionalProperties, "max_sale_qty")
		delete(additionalProperties, "use_config_backorders")
		delete(additionalProperties, "backorders")
		delete(additionalProperties, "use_config_notify_stock_qty")
		delete(additionalProperties, "notify_stock_qty")
		delete(additionalProperties, "use_config_qty_increments")
		delete(additionalProperties, "qty_increments")
		delete(additionalProperties, "use_config_enable_qty_inc")
		delete(additionalProperties, "enable_qty_increments")
		delete(additionalProperties, "use_config_manage_stock")
		delete(additionalProperties, "manage_stock")
		delete(additionalProperties, "low_stock_date")
		delete(additionalProperties, "is_decimal_divided")
		delete(additionalProperties, "stock_status_changed_auto")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CatalogInventoryDataStockItemInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CatalogInventoryDataStockItemInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCatalogInventoryDataStockItemInterface struct {
	value *CatalogInventoryDataStockItemInterface
	isSet bool
}

func (v NullableCatalogInventoryDataStockItemInterface) Get() *CatalogInventoryDataStockItemInterface {
	return v.value
}

func (v *NullableCatalogInventoryDataStockItemInterface) Set(val *CatalogInventoryDataStockItemInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogInventoryDataStockItemInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogInventoryDataStockItemInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogInventoryDataStockItemInterface(val *CatalogInventoryDataStockItemInterface) *NullableCatalogInventoryDataStockItemInterface {
	return &NullableCatalogInventoryDataStockItemInterface{value: val, isSet: true}
}

func (v NullableCatalogInventoryDataStockItemInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogInventoryDataStockItemInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
