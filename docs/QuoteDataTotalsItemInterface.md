# # QuoteDataTotalsItemInterface
Interface TotalsItemInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId**| **int32** | Item id  |
**Price**| **float32** | Item price in quote currency.  |
**BasePrice**| **float32** | Item price in base currency.  |
**Qty**| **float32** | Item quantity.  |
**RowTotal**| **float32** | Row total in quote currency.  |
**BaseRowTotal**| **float32** | Row total in base currency.  |
**RowTotalWithDiscount**| **float32** | Row total with discount in quote currency. Otherwise, null.  | [optional]
**TaxAmount**| **float32** | Tax amount in quote currency. Otherwise, null.  | [optional]
**BaseTaxAmount**| **float32** | Tax amount in base currency. Otherwise, null.  | [optional]
**TaxPercent**| **float32** | Tax percent. Otherwise, null.  | [optional]
**DiscountAmount**| **float32** | Discount amount in quote currency. Otherwise, null.  | [optional]
**BaseDiscountAmount**| **float32** | Discount amount in base currency. Otherwise, null.  | [optional]
**DiscountPercent**| **float32** | Discount percent. Otherwise, null.  | [optional]
**PriceInclTax**| **float32** | Price including tax in quote currency. Otherwise, null.  | [optional]
**BasePriceInclTax**| **float32** | Price including tax in base currency. Otherwise, null.  | [optional]
**RowTotalInclTax**| **float32** | Row total including tax in quote currency. Otherwise, null.  | [optional]
**BaseRowTotalInclTax**| **float32** | Row total including tax in base currency. Otherwise, null.  | [optional]
**Options**| **string** | Item price in quote currency.  |
**WeeeTaxAppliedAmount**| **float32** | Item weee tax applied amount in quote currency.  |
**WeeeTaxApplied**| **string** | Item weee tax applied in quote currency.  |
**ExtensionAttributes**| [**QuoteDataTotalsItemExtensionInterface**](QuoteDataTotalsItemExtensionInterface.md) |   | [optional]
**Name**| **string** | Product name. Otherwise, null.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

