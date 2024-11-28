# # TaxDataOrderTaxItemInterface


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxItemId**| **int32** | Tax item ID  | [optional]
**TaxId**| **int32** | Tax ID  | [optional]
**ItemId**| **int32** | Order item ID  | [optional]
**TaxCode**| **string** | Tax code  | [optional]
**TaxPercent**| **float32** | Tax percent  |
**Amount**| **float32** | Tax amount  |
**BaseAmount**| **float32** | Tax amount in base currency  |
**RealAmount**| **float32** | Real tax amount  |
**RealBaseAmount**| **float32** | Real tax amount in base currency  |
**AssociatedItemId**| **int32** | Associated order item ID  | [optional]
**TaxableItemType**| **string** | shipping, product, weee, quote_gw, etc...  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Tax\\Api\\Data\\OrderTaxItemInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

