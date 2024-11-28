# # CheckoutDataShippingInformationInterface
Interface ShippingInformationInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingAddress**| [**QuoteDataAddressInterface**](QuoteDataAddressInterface.md) |   |
**BillingAddress**| [**QuoteDataAddressInterface**](QuoteDataAddressInterface.md) |   | [optional]
**ShippingMethodCode**| **string** | Shipping method code  |
**ShippingCarrierCode**| **string** | Carrier code  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Checkout\\Api\\Data\\ShippingInformationInterface  | [optional]
**CustomAttributes**| [**[]FrameworkAttributeInterface**](FrameworkAttributeInterface.md) | Custom attributes values.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

