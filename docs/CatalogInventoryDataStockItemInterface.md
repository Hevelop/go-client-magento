# # CatalogInventoryDataStockItemInterface
Interface StockItem

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId**| **int32** |   | [optional]
**ProductId**| **int32** |   | [optional]
**StockId**| **int32** | Stock identifier  | [optional]
**Qty**| **float32** |   |
**IsInStock**| **bool** | Stock Availability  |
**IsQtyDecimal**| **bool** |   |
**ShowDefaultNotificationMessage**| **bool** |   |
**UseConfigMinQty**| **bool** |   |
**MinQty**| **float32** | Minimal quantity available for item status in stock  |
**UseConfigMinSaleQty**| **int32** |   |
**MinSaleQty**| **float32** | Minimum Qty Allowed in Shopping Cart or NULL when there is no limitation  |
**UseConfigMaxSaleQty**| **bool** |   |
**MaxSaleQty**| **float32** | Maximum Qty Allowed in Shopping Cart data wrapper  |
**UseConfigBackorders**| **bool** |   |
**Backorders**| **int32** | Backorders status  |
**UseConfigNotifyStockQty**| **bool** |   |
**NotifyStockQty**| **float32** | Notify for Quantity Below data wrapper  |
**UseConfigQtyIncrements**| **bool** |   |
**QtyIncrements**| **float32** | Quantity Increments data wrapper  |
**UseConfigEnableQtyInc**| **bool** |   |
**EnableQtyIncrements**| **bool** | Whether Quantity Increments is enabled  |
**UseConfigManageStock**| **bool** |   |
**ManageStock**| **bool** | Can Manage Stock  |
**LowStockDate**| **string** |   |
**IsDecimalDivided**| **bool** |   |
**StockStatusChangedAuto**| **int32** |   |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\CatalogInventory\\Api\\Data\\StockItemInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

