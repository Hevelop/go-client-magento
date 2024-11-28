# # InventoryApiDataSourceItemInterface
Represents amount of product on physical storage Entity id getter is missed because entity identifies by compound identifier (sku and source_code) Used fully qualified namespaces in annotations for proper work of WebApi request parser

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku**| **string** | Source item sku  | [optional]
**SourceCode**| **string** | Source code  | [optional]
**Quantity**| **float32** | Source item quantity  | [optional]
**Status**| **int32** | Source item status (One of self::STATUS_*)  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\InventoryApi\\Api\\Data\\SourceItemInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

