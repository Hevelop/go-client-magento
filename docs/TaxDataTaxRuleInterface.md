# # TaxDataTaxRuleInterface
Tax rule interface.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | Id  | [optional]
**Code**| **string** | Tax rule code  |
**Priority**| **int32** | Priority  |
**Position**| **int32** | Sort order.  |
**CustomerTaxClassIds**| **[]int32** | Customer tax class id  |
**ProductTaxClassIds**| **[]int32** | Product tax class id  |
**TaxRateIds**| **[]int32** | Tax rate ids  |
**CalculateSubtotal**| **bool** | Calculate subtotal.  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Tax\\Api\\Data\\TaxRuleInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

