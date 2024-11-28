# # SalesDataOrderInterface
Order interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustmentNegative**| **float32** | Negative adjustment value.  | [optional]
**AdjustmentPositive**| **float32** | Positive adjustment value.  | [optional]
**AppliedRuleIds**| **string** | Applied rule IDs.  | [optional]
**BaseAdjustmentNegative**| **float32** | Base negative adjustment value.  | [optional]
**BaseAdjustmentPositive**| **float32** | Base positive adjustment value.  | [optional]
**BaseCurrencyCode**| **string** | Base currency code.  | [optional]
**BaseDiscountAmount**| **float32** | Base discount amount.  | [optional]
**BaseDiscountCanceled**| **float32** | Base discount canceled.  | [optional]
**BaseDiscountInvoiced**| **float32** | Base discount invoiced.  | [optional]
**BaseDiscountRefunded**| **float32** | Base discount refunded.  | [optional]
**BaseGrandTotal**| **float32** | Base grand total.  |
**BaseDiscountTaxCompensationAmount**| **float32** | Base discount tax compensation amount.  | [optional]
**BaseDiscountTaxCompensationInvoiced**| **float32** | Base discount tax compensation invoiced.  | [optional]
**BaseDiscountTaxCompensationRefunded**| **float32** | Base discount tax compensation refunded.  | [optional]
**BaseShippingAmount**| **float32** | Base shipping amount.  | [optional]
**BaseShippingCanceled**| **float32** | Base shipping canceled.  | [optional]
**BaseShippingDiscountAmount**| **float32** | Base shipping discount amount.  | [optional]
**BaseShippingDiscountTaxCompensationAmnt**| **float32** | Base shipping discount tax compensation amount.  | [optional]
**BaseShippingInclTax**| **float32** | Base shipping including tax.  | [optional]
**BaseShippingInvoiced**| **float32** | Base shipping invoiced.  | [optional]
**BaseShippingRefunded**| **float32** | Base shipping refunded.  | [optional]
**BaseShippingTaxAmount**| **float32** | Base shipping tax amount.  | [optional]
**BaseShippingTaxRefunded**| **float32** | Base shipping tax refunded.  | [optional]
**BaseSubtotal**| **float32** | Base subtotal.  | [optional]
**BaseSubtotalCanceled**| **float32** | Base subtotal canceled.  | [optional]
**BaseSubtotalInclTax**| **float32** | Base subtotal including tax.  | [optional]
**BaseSubtotalInvoiced**| **float32** | Base subtotal invoiced.  | [optional]
**BaseSubtotalRefunded**| **float32** | Base subtotal refunded.  | [optional]
**BaseTaxAmount**| **float32** | Base tax amount.  | [optional]
**BaseTaxCanceled**| **float32** | Base tax canceled.  | [optional]
**BaseTaxInvoiced**| **float32** | Base tax invoiced.  | [optional]
**BaseTaxRefunded**| **float32** | Base tax refunded.  | [optional]
**BaseTotalCanceled**| **float32** | Base total canceled.  | [optional]
**BaseTotalDue**| **float32** | Base total due.  | [optional]
**BaseTotalInvoiced**| **float32** | Base total invoiced.  | [optional]
**BaseTotalInvoicedCost**| **float32** | Base total invoiced cost.  | [optional]
**BaseTotalOfflineRefunded**| **float32** | Base total offline refunded.  | [optional]
**BaseTotalOnlineRefunded**| **float32** | Base total online refunded.  | [optional]
**BaseTotalPaid**| **float32** | Base total paid.  | [optional]
**BaseTotalQtyOrdered**| **float32** | Base total quantity ordered.  | [optional]
**BaseTotalRefunded**| **float32** | Base total refunded.  | [optional]
**BaseToGlobalRate**| **float32** | Base-to-global rate.  | [optional]
**BaseToOrderRate**| **float32** | Base-to-order rate.  | [optional]
**BillingAddressId**| **int32** | Billing address ID.  | [optional]
**CanShipPartially**| **int32** | Can-ship-partially flag value.  | [optional]
**CanShipPartiallyItem**| **int32** | Can-ship-partially-item flag value.  | [optional]
**CouponCode**| **string** | Coupon code.  | [optional]
**CreatedAt**| **string** | Created-at timestamp.  | [optional]
**CustomerDob**| **string** | In keeping with current security and privacy best practices, be sure you are aware of any potential legal and security risks associated with the storage of customers’ full date of birth (month, day, year) along with other personal identifiers (e.g., full name) before collecting or processing such data.  | [optional]
**CustomerEmail**| **string** | Customer email address.  |
**CustomerFirstname**| **string** | Customer first name.  | [optional]
**CustomerGender**| **int32** | Customer gender.  | [optional]
**CustomerGroupId**| **int32** | Customer group ID.  | [optional]
**CustomerId**| **int32** | Customer ID.  | [optional]
**CustomerIsGuest**| **int32** | Customer-is-guest flag value.  | [optional]
**CustomerLastname**| **string** | Customer last name.  | [optional]
**CustomerMiddlename**| **string** | Customer middle name.  | [optional]
**CustomerNote**| **string** | Customer note.  | [optional]
**CustomerNoteNotify**| **int32** | Customer-note-notify flag value.  | [optional]
**CustomerPrefix**| **string** | Customer prefix.  | [optional]
**CustomerSuffix**| **string** | Customer suffix.  | [optional]
**CustomerTaxvat**| **string** | Customer value-added tax (VAT).  | [optional]
**DiscountAmount**| **float32** | Discount amount.  | [optional]
**DiscountCanceled**| **float32** | Discount canceled.  | [optional]
**DiscountDescription**| **string** | Discount description.  | [optional]
**DiscountInvoiced**| **float32** | Discount invoiced.  | [optional]
**DiscountRefunded**| **float32** | Discount refunded amount.  | [optional]
**EditIncrement**| **int32** | Edit increment value.  | [optional]
**EmailSent**| **int32** | Email-sent flag value.  | [optional]
**EntityId**| **int32** | Order ID.  | [optional]
**ExtCustomerId**| **string** | External customer ID.  | [optional]
**ExtOrderId**| **string** | External order ID.  | [optional]
**ForcedShipmentWithInvoice**| **int32** | Forced-shipment-with-invoice flag value.  | [optional]
**GlobalCurrencyCode**| **string** | Global currency code.  | [optional]
**GrandTotal**| **float32** | Grand total.  |
**DiscountTaxCompensationAmount**| **float32** | Discount tax compensation amount.  | [optional]
**DiscountTaxCompensationInvoiced**| **float32** | Discount tax compensation invoiced amount.  | [optional]
**DiscountTaxCompensationRefunded**| **float32** | Discount tax compensation refunded amount.  | [optional]
**HoldBeforeState**| **string** | Hold before state.  | [optional]
**HoldBeforeStatus**| **string** | Hold before status.  | [optional]
**IncrementId**| **string** | Increment ID.  | [optional]
**IsVirtual**| **int32** | Is-virtual flag value.  | [optional]
**OrderCurrencyCode**| **string** | Order currency code.  | [optional]
**OriginalIncrementId**| **string** | Original increment ID.  | [optional]
**PaymentAuthorizationAmount**| **float32** | Payment authorization amount.  | [optional]
**PaymentAuthExpiration**| **int32** | Payment authorization expiration date.  | [optional]
**ProtectCode**| **string** | Protect code.  | [optional]
**QuoteAddressId**| **int32** | Quote address ID.  | [optional]
**QuoteId**| **int32** | Quote ID.  | [optional]
**RelationChildId**| **string** | Relation child ID.  | [optional]
**RelationChildRealId**| **string** | Relation child real ID.  | [optional]
**RelationParentId**| **string** | Relation parent ID.  | [optional]
**RelationParentRealId**| **string** | Relation parent real ID.  | [optional]
**RemoteIp**| **string** | Remote IP address.  | [optional]
**ShippingAmount**| **float32** | Shipping amount.  | [optional]
**ShippingCanceled**| **float32** | Shipping canceled amount.  | [optional]
**ShippingDescription**| **string** | Shipping description.  | [optional]
**ShippingDiscountAmount**| **float32** | Shipping discount amount.  | [optional]
**ShippingDiscountTaxCompensationAmount**| **float32** | Shipping discount tax compensation amount.  | [optional]
**ShippingInclTax**| **float32** | Shipping including tax amount.  | [optional]
**ShippingInvoiced**| **float32** | Shipping invoiced amount.  | [optional]
**ShippingRefunded**| **float32** | Shipping refunded amount.  | [optional]
**ShippingTaxAmount**| **float32** | Shipping tax amount.  | [optional]
**ShippingTaxRefunded**| **float32** | Shipping tax refunded amount.  | [optional]
**State**| **string** | State.  | [optional]
**Status**| **string** | Status.  | [optional]
**StoreCurrencyCode**| **string** | Store currency code.  | [optional]
**StoreId**| **int32** | Store ID.  | [optional]
**StoreName**| **string** | Store name.  | [optional]
**StoreToBaseRate**| **float32** | Store-to-base rate.  | [optional]
**StoreToOrderRate**| **float32** | Store-to-order rate.  | [optional]
**Subtotal**| **float32** | Subtotal.  | [optional]
**SubtotalCanceled**| **float32** | Subtotal canceled amount.  | [optional]
**SubtotalInclTax**| **float32** | Subtotal including tax amount.  | [optional]
**SubtotalInvoiced**| **float32** | Subtotal invoiced amount.  | [optional]
**SubtotalRefunded**| **float32** | Subtotal refunded amount.  | [optional]
**TaxAmount**| **float32** | Tax amount.  | [optional]
**TaxCanceled**| **float32** | Tax canceled amount.  | [optional]
**TaxInvoiced**| **float32** | Tax invoiced amount.  | [optional]
**TaxRefunded**| **float32** | Tax refunded amount.  | [optional]
**TotalCanceled**| **float32** | Total canceled.  | [optional]
**TotalDue**| **float32** | Total due.  | [optional]
**TotalInvoiced**| **float32** | Total invoiced amount.  | [optional]
**TotalItemCount**| **int32** | Total item count.  | [optional]
**TotalOfflineRefunded**| **float32** | Total offline refunded amount.  | [optional]
**TotalOnlineRefunded**| **float32** | Total online refunded amount.  | [optional]
**TotalPaid**| **float32** | Total paid.  | [optional]
**TotalQtyOrdered**| **float32** | Total quantity ordered.  | [optional]
**TotalRefunded**| **float32** | Total amount refunded.  | [optional]
**UpdatedAt**| **string** | Updated-at timestamp.  | [optional]
**Weight**| **float32** | Weight.  | [optional]
**XForwardedFor**| **string** | X-Forwarded-For field value.  | [optional]
**Items**| [**[]SalesDataOrderItemInterface**](SalesDataOrderItemInterface.md) | Array of items.  |
**BillingAddress**| [**SalesDataOrderAddressInterface**](SalesDataOrderAddressInterface.md) |   | [optional]
**Payment**| [**SalesDataOrderPaymentInterface**](SalesDataOrderPaymentInterface.md) |   | [optional]
**StatusHistories**| [**[]SalesDataOrderStatusHistoryInterface**](SalesDataOrderStatusHistoryInterface.md) | Array of status histories.  | [optional]
**ExtensionAttributes**| [**SalesDataOrderExtensionInterface**](SalesDataOrderExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
