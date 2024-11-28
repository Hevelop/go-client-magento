# # SalesDataCreditmemoInterface
Credit memo interface. After a customer places and pays for an order and an invoice has been issued, the merchant can create a credit memo to refund all or part of the amount paid for any returned or undelivered items. The memo restores funds to the customer account so that the customer can make future purchases.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adjustment**| **float32** | Credit memo adjustment.  | [optional]
**AdjustmentNegative**| **float32** | Credit memo negative adjustment.  | [optional]
**AdjustmentPositive**| **float32** | Credit memo positive adjustment.  | [optional]
**BaseAdjustment**| **float32** | Credit memo base adjustment.  | [optional]
**BaseAdjustmentNegative**| **float32** | Credit memo negative base adjustment.  | [optional]
**BaseAdjustmentPositive**| **float32** | Credit memo positive base adjustment.  | [optional]
**BaseCurrencyCode**| **string** | Credit memo base currency code.  | [optional]
**BaseDiscountAmount**| **float32** | Credit memo base discount amount.  | [optional]
**BaseGrandTotal**| **float32** | Credit memo base grand total.  | [optional]
**BaseDiscountTaxCompensationAmount**| **float32** | Credit memo base discount tax compensation amount.  | [optional]
**BaseShippingAmount**| **float32** | Credit memo base shipping amount.  | [optional]
**BaseShippingDiscountTaxCompensationAmnt**| **float32** | Credit memo base shipping discount tax compensation amount.  | [optional]
**BaseShippingInclTax**| **float32** | Credit memo base shipping including tax.  | [optional]
**BaseShippingTaxAmount**| **float32** | Credit memo base shipping tax amount.  | [optional]
**BaseSubtotal**| **float32** | Credit memo base subtotal.  | [optional]
**BaseSubtotalInclTax**| **float32** | Credit memo base subtotal including tax.  | [optional]
**BaseTaxAmount**| **float32** | Credit memo base tax amount.  | [optional]
**BaseToGlobalRate**| **float32** | Credit memo base-to-global rate.  | [optional]
**BaseToOrderRate**| **float32** | Credit memo base-to-order rate.  | [optional]
**BillingAddressId**| **int32** | Credit memo billing address ID.  | [optional]
**CreatedAt**| **string** | Credit memo created-at timestamp.  | [optional]
**CreditmemoStatus**| **int32** | Credit memo status.  | [optional]
**DiscountAmount**| **float32** | Credit memo discount amount.  | [optional]
**DiscountDescription**| **string** | Credit memo discount description.  | [optional]
**EmailSent**| **int32** | Credit memo email sent flag value.  | [optional]
**EntityId**| **int32** | Credit memo ID.  | [optional]
**GlobalCurrencyCode**| **string** | Credit memo global currency code.  | [optional]
**GrandTotal**| **float32** | Credit memo grand total.  | [optional]
**DiscountTaxCompensationAmount**| **float32** | Credit memo discount tax compensation amount.  | [optional]
**IncrementId**| **string** | Credit memo increment ID.  | [optional]
**InvoiceId**| **int32** | Credit memo invoice ID.  | [optional]
**OrderCurrencyCode**| **string** | Credit memo order currency code.  | [optional]
**OrderId**| **int32** | Credit memo order ID.  |
**ShippingAddressId**| **int32** | Credit memo shipping address ID.  | [optional]
**ShippingAmount**| **float32** | Credit memo shipping amount.  | [optional]
**ShippingDiscountTaxCompensationAmount**| **float32** | Credit memo shipping discount tax compensation amount.  | [optional]
**ShippingInclTax**| **float32** | Credit memo shipping including tax.  | [optional]
**ShippingTaxAmount**| **float32** | Credit memo shipping tax amount.  | [optional]
**State**| **int32** | Credit memo state.  | [optional]
**StoreCurrencyCode**| **string** | Credit memo store currency code.  | [optional]
**StoreId**| **int32** | Credit memo store ID.  | [optional]
**StoreToBaseRate**| **float32** | Credit memo store-to-base rate.  | [optional]
**StoreToOrderRate**| **float32** | Credit memo store-to-order rate.  | [optional]
**Subtotal**| **float32** | Credit memo subtotal.  | [optional]
**SubtotalInclTax**| **float32** | Credit memo subtotal including tax.  | [optional]
**TaxAmount**| **float32** | Credit memo tax amount.  | [optional]
**TransactionId**| **string** | Credit memo transaction ID.  | [optional]
**UpdatedAt**| **string** | Credit memo updated-at timestamp.  | [optional]
**Items**| [**[]SalesDataCreditmemoItemInterface**](SalesDataCreditmemoItemInterface.md) | Array of credit memo items.  |
**Comments**| [**[]SalesDataCreditmemoCommentInterface**](SalesDataCreditmemoCommentInterface.md) | Array of any credit memo comments. Otherwise, null.  | [optional]
**ExtensionAttributes**| [**SalesDataCreditmemoExtensionInterface**](SalesDataCreditmemoExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

