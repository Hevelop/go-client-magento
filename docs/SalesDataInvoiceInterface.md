# # SalesDataInvoiceInterface
Invoice interface. An invoice is a record of the receipt of payment for an order.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCurrencyCode**| **string** | Base currency code.  | [optional]
**BaseDiscountAmount**| **float32** | Base discount amount.  | [optional]
**BaseGrandTotal**| **float32** | Base grand total.  | [optional]
**BaseDiscountTaxCompensationAmount**| **float32** | Base discount tax compensation amount.  | [optional]
**BaseShippingAmount**| **float32** | Base shipping amount.  | [optional]
**BaseShippingDiscountTaxCompensationAmnt**| **float32** | Base shipping discount tax compensation amount.  | [optional]
**BaseShippingInclTax**| **float32** | Base shipping including tax.  | [optional]
**BaseShippingTaxAmount**| **float32** | Base shipping tax amount.  | [optional]
**BaseSubtotal**| **float32** | Base subtotal.  | [optional]
**BaseSubtotalInclTax**| **float32** | Base subtotal including tax.  | [optional]
**BaseTaxAmount**| **float32** | Base tax amount.  | [optional]
**BaseTotalRefunded**| **float32** | Base total refunded.  | [optional]
**BaseToGlobalRate**| **float32** | Base-to-global rate.  | [optional]
**BaseToOrderRate**| **float32** | Base-to-order rate.  | [optional]
**BillingAddressId**| **int32** | Billing address ID.  | [optional]
**CanVoidFlag**| **int32** | Can void flag value.  | [optional]
**CreatedAt**| **string** | Created-at timestamp.  | [optional]
**DiscountAmount**| **float32** | Discount amount.  | [optional]
**DiscountDescription**| **string** | Discount description.  | [optional]
**EmailSent**| **int32** | Email-sent flag value.  | [optional]
**EntityId**| **int32** | Invoice ID.  | [optional]
**GlobalCurrencyCode**| **string** | Global currency code.  | [optional]
**GrandTotal**| **float32** | Grand total.  | [optional]
**DiscountTaxCompensationAmount**| **float32** | Discount tax compensation amount.  | [optional]
**IncrementId**| **string** | Increment ID.  | [optional]
**IsUsedForRefund**| **int32** | Is-used-for-refund flag value.  | [optional]
**OrderCurrencyCode**| **string** | Order currency code.  | [optional]
**OrderId**| **int32** | Order ID.  |
**ShippingAddressId**| **int32** | Shipping address ID.  | [optional]
**ShippingAmount**| **float32** | Shipping amount.  | [optional]
**ShippingDiscountTaxCompensationAmount**| **float32** | Shipping discount tax compensation amount.  | [optional]
**ShippingInclTax**| **float32** | Shipping including tax.  | [optional]
**ShippingTaxAmount**| **float32** | Shipping tax amount.  | [optional]
**State**| **int32** | State.  | [optional]
**StoreCurrencyCode**| **string** | Store currency code.  | [optional]
**StoreId**| **int32** | Store ID.  | [optional]
**StoreToBaseRate**| **float32** | Store-to-base rate.  | [optional]
**StoreToOrderRate**| **float32** | Store-to-order rate.  | [optional]
**Subtotal**| **float32** | Subtotal.  | [optional]
**SubtotalInclTax**| **float32** | Subtotal including tax.  | [optional]
**TaxAmount**| **float32** | Tax amount.  | [optional]
**TotalQty**| **float32** | Total quantity.  |
**TransactionId**| **string** | Transaction ID.  | [optional]
**UpdatedAt**| **string** | Updated-at timestamp.  | [optional]
**Items**| [**[]SalesDataInvoiceItemInterface**](SalesDataInvoiceItemInterface.md) | Array of invoice items.  |
**Comments**| [**[]SalesDataInvoiceCommentInterface**](SalesDataInvoiceCommentInterface.md) | Array of any invoice comments. Otherwise, null.  | [optional]
**ExtensionAttributes**| [**SalesDataInvoiceExtensionInterface**](SalesDataInvoiceExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

