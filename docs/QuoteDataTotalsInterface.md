# # QuoteDataTotalsInterface
Interface TotalsInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrandTotal**| **float32** | Grand total in quote currency  | [optional]
**BaseGrandTotal**| **float32** | Grand total in base currency  | [optional]
**Subtotal**| **float32** | Subtotal in quote currency  | [optional]
**BaseSubtotal**| **float32** | Subtotal in base currency  | [optional]
**DiscountAmount**| **float32** | Discount amount in quote currency  | [optional]
**BaseDiscountAmount**| **float32** | Discount amount in base currency  | [optional]
**SubtotalWithDiscount**| **float32** | Subtotal in quote currency with applied discount  | [optional]
**BaseSubtotalWithDiscount**| **float32** | Subtotal in base currency with applied discount  | [optional]
**ShippingAmount**| **float32** | Shipping amount in quote currency  | [optional]
**BaseShippingAmount**| **float32** | Shipping amount in base currency  | [optional]
**ShippingDiscountAmount**| **float32** | Shipping discount amount in quote currency  | [optional]
**BaseShippingDiscountAmount**| **float32** | Shipping discount amount in base currency  | [optional]
**TaxAmount**| **float32** | Tax amount in quote currency  | [optional]
**BaseTaxAmount**| **float32** | Tax amount in base currency  | [optional]
**WeeeTaxAppliedAmount**| **float32** | Item weee tax applied amount in quote currency.  |
**ShippingTaxAmount**| **float32** | Shipping tax amount in quote currency  | [optional]
**BaseShippingTaxAmount**| **float32** | Shipping tax amount in base currency  | [optional]
**SubtotalInclTax**| **float32** | Subtotal including tax in quote currency  | [optional]
**BaseSubtotalInclTax**| **float32** | Subtotal including tax in base currency  | [optional]
**ShippingInclTax**| **float32** | Shipping including tax in quote currency  | [optional]
**BaseShippingInclTax**| **float32** | Shipping including tax in base currency  | [optional]
**BaseCurrencyCode**| **string** | Base currency code  | [optional]
**QuoteCurrencyCode**| **string** | Quote currency code  | [optional]
**CouponCode**| **string** | Applied coupon code  | [optional]
**ItemsQty**| **int32** | Items qty  | [optional]
**Items**| [**[]QuoteDataTotalsItemInterface**](QuoteDataTotalsItemInterface.md) | Totals by items  | [optional]
**TotalSegments**| [**[]QuoteDataTotalSegmentInterface**](QuoteDataTotalSegmentInterface.md) | Dynamically calculated totals  |
**ExtensionAttributes**| [**QuoteDataTotalsExtensionInterface**](QuoteDataTotalsExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

