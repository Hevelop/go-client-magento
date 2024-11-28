# # QuoteDataShippingMethodInterface
Interface ShippingMethodInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierCode**| **string** | Shipping carrier code.  |
**MethodCode**| **string** | Shipping method code.  |
**CarrierTitle**| **string** | Shipping carrier title. Otherwise, null.  | [optional]
**MethodTitle**| **string** | Shipping method title. Otherwise, null.  | [optional]
**Amount**| **float32** | Shipping amount in store currency.  |
**BaseAmount**| **float32** | Shipping amount in base currency.  |
**Available**| **bool** | The value of the availability flag for the current shipping method.  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Quote\\Api\\Data\\ShippingMethodInterface  | [optional]
**ErrorMessage**| **string** | Shipping Error message.  |
**PriceExclTax**| **float32** | Shipping price excl tax.  |
**PriceInclTax**| **float32** | Shipping price incl tax.  |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

