# # TaxDataOrderTaxDetailsItemInterface
Interface OrderTaxDetailsItemInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type**| **string** | Type (shipping, product, weee, gift wrapping, etc)  | [optional]
**ItemId**| **int32** | Item id if this item is a product  | [optional]
**AssociatedItemId**| **int32** | Associated item id if this item is associated with another item, null otherwise  | [optional]
**AppliedTaxes**| [**[]TaxDataOrderTaxDetailsAppliedTaxInterface**](TaxDataOrderTaxDetailsAppliedTaxInterface.md) | Applied taxes  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Tax\\Api\\Data\\OrderTaxDetailsItemInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

