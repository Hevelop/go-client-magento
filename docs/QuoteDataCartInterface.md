# # QuoteDataCartInterface
Interface CartInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | Cart/quote ID.  |
**CreatedAt**| **string** | Cart creation date and time. Otherwise, null.  | [optional]
**UpdatedAt**| **string** | Cart last update date and time. Otherwise, null.  | [optional]
**ConvertedAt**| **string** | Cart conversion date and time. Otherwise, null.  | [optional]
**IsActive**| **bool** | Active status flag value. Otherwise, null.  | [optional]
**IsVirtual**| **bool** | Virtual flag value. Otherwise, null.  | [optional]
**Items**| [**[]QuoteDataCartItemInterface**](QuoteDataCartItemInterface.md) | Array of items. Otherwise, null.  | [optional]
**ItemsCount**| **int32** | Number of different items or products in the cart. Otherwise, null.  | [optional]
**ItemsQty**| **float32** | Total quantity of all cart items. Otherwise, null.  | [optional]
**Customer**| [**CustomerDataCustomerInterface**](CustomerDataCustomerInterface.md) |   |
**BillingAddress**| [**QuoteDataAddressInterface**](QuoteDataAddressInterface.md) |   | [optional]
**ReservedOrderId**| **string** | Reserved order ID. Otherwise, null.  | [optional]
**OrigOrderId**| **int32** | Original order ID. Otherwise, null.  | [optional]
**Currency**| [**QuoteDataCurrencyInterface**](QuoteDataCurrencyInterface.md) |   | [optional]
**CustomerIsGuest**| **bool** | For guest customers, false for logged in customers  | [optional]
**CustomerNote**| **string** | Notice text  | [optional]
**CustomerNoteNotify**| **bool** | Customer notification flag  | [optional]
**CustomerTaxClassId**| **int32** | Customer tax class ID.  | [optional]
**StoreId**| **int32** | Store identifier  |
**ExtensionAttributes**| [**QuoteDataCartExtensionInterface**](QuoteDataCartExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

