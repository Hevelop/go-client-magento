# # SalesDataOrderItemInterface
Order item interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData**| **string** | Additional data.  | [optional]
**AmountRefunded**| **float32** | Amount refunded.  | [optional]
**AppliedRuleIds**| **string** | Applied rule IDs.  | [optional]
**BaseAmountRefunded**| **float32** | Base amount refunded.  | [optional]
**BaseCost**| **float32** | Base cost.  | [optional]
**BaseDiscountAmount**| **float32** | Base discount amount.  | [optional]
**BaseDiscountInvoiced**| **float32** | Base discount invoiced.  | [optional]
**BaseDiscountRefunded**| **float32** | Base discount refunded.  | [optional]
**BaseDiscountTaxCompensationAmount**| **float32** | Base discount tax compensation amount.  | [optional]
**BaseDiscountTaxCompensationInvoiced**| **float32** | Base discount tax compensation invoiced.  | [optional]
**BaseDiscountTaxCompensationRefunded**| **float32** | Base discount tax compensation refunded.  | [optional]
**BaseOriginalPrice**| **float32** | Base original price.  | [optional]
**BasePrice**| **float32** | Base price.  | [optional]
**BasePriceInclTax**| **float32** | Base price including tax.  | [optional]
**BaseRowInvoiced**| **float32** | Base row invoiced.  | [optional]
**BaseRowTotal**| **float32** | Base row total.  | [optional]
**BaseRowTotalInclTax**| **float32** | Base row total including tax.  | [optional]
**BaseTaxAmount**| **float32** | Base tax amount.  | [optional]
**BaseTaxBeforeDiscount**| **float32** | Base tax before discount.  | [optional]
**BaseTaxInvoiced**| **float32** | Base tax invoiced.  | [optional]
**BaseTaxRefunded**| **float32** | Base tax refunded.  | [optional]
**BaseWeeeTaxAppliedAmount**| **float32** | Base WEEE tax applied amount.  | [optional]
**BaseWeeeTaxAppliedRowAmnt**| **float32** | Base WEEE tax applied row amount.  | [optional]
**BaseWeeeTaxDisposition**| **float32** | Base WEEE tax disposition.  | [optional]
**BaseWeeeTaxRowDisposition**| **float32** | Base WEEE tax row disposition.  | [optional]
**CreatedAt**| **string** | Created-at timestamp.  | [optional]
**Description**| **string** | Description.  | [optional]
**DiscountAmount**| **float32** | Discount amount.  | [optional]
**DiscountInvoiced**| **float32** | Discount invoiced.  | [optional]
**DiscountPercent**| **float32** | Discount percent.  | [optional]
**DiscountRefunded**| **float32** | Discount refunded.  | [optional]
**EventId**| **int32** | Event ID.  | [optional]
**ExtOrderItemId**| **string** | External order item ID.  | [optional]
**FreeShipping**| **int32** | Free-shipping flag value.  | [optional]
**GwBasePrice**| **float32** | GW base price.  | [optional]
**GwBasePriceInvoiced**| **float32** | GW base price invoiced.  | [optional]
**GwBasePriceRefunded**| **float32** | GW base price refunded.  | [optional]
**GwBaseTaxAmount**| **float32** | GW base tax amount.  | [optional]
**GwBaseTaxAmountInvoiced**| **float32** | GW base tax amount invoiced.  | [optional]
**GwBaseTaxAmountRefunded**| **float32** | GW base tax amount refunded.  | [optional]
**GwId**| **int32** | GW ID.  | [optional]
**GwPrice**| **float32** | GW price.  | [optional]
**GwPriceInvoiced**| **float32** | GW price invoiced.  | [optional]
**GwPriceRefunded**| **float32** | GW price refunded.  | [optional]
**GwTaxAmount**| **float32** | GW tax amount.  | [optional]
**GwTaxAmountInvoiced**| **float32** | GW tax amount invoiced.  | [optional]
**GwTaxAmountRefunded**| **float32** | GW tax amount refunded.  | [optional]
**DiscountTaxCompensationAmount**| **float32** | Discount tax compensation amount.  | [optional]
**DiscountTaxCompensationCanceled**| **float32** | Discount tax compensation canceled.  | [optional]
**DiscountTaxCompensationInvoiced**| **float32** | Discount tax compensation invoiced.  | [optional]
**DiscountTaxCompensationRefunded**| **float32** | Discount tax compensation refunded.  | [optional]
**IsQtyDecimal**| **int32** | Is-quantity-decimal flag value.  | [optional]
**IsVirtual**| **int32** | Is-virtual flag value.  | [optional]
**ItemId**| **int32** | Item ID.  | [optional]
**LockedDoInvoice**| **int32** | Locked DO invoice flag value.  | [optional]
**LockedDoShip**| **int32** | Locked DO ship flag value.  | [optional]
**Name**| **string** | Name.  | [optional]
**NoDiscount**| **int32** | No-discount flag value.  | [optional]
**OrderId**| **int32** | Order ID.  | [optional]
**OriginalPrice**| **float32** | Original price.  | [optional]
**ParentItemId**| **int32** | Parent item ID.  | [optional]
**Price**| **float32** | Price.  | [optional]
**PriceInclTax**| **float32** | Price including tax.  | [optional]
**ProductId**| **int32** | Product ID.  | [optional]
**ProductType**| **string** | Product type.  | [optional]
**QtyBackordered**| **float32** | Quantity backordered.  | [optional]
**QtyCanceled**| **float32** | Quantity canceled.  | [optional]
**QtyInvoiced**| **float32** | Quantity invoiced.  | [optional]
**QtyOrdered**| **float32** | Quantity ordered.  | [optional]
**QtyRefunded**| **float32** | Quantity refunded.  | [optional]
**QtyReturned**| **float32** | Quantity returned.  | [optional]
**QtyShipped**| **float32** | Quantity shipped.  | [optional]
**QuoteItemId**| **int32** | Quote item ID.  | [optional]
**RowInvoiced**| **float32** | Row invoiced.  | [optional]
**RowTotal**| **float32** | Row total.  | [optional]
**RowTotalInclTax**| **float32** | Row total including tax.  | [optional]
**RowWeight**| **float32** | Row weight.  | [optional]
**Sku**| **string** | SKU.  |
**StoreId**| **int32** | Store ID.  | [optional]
**TaxAmount**| **float32** | Tax amount.  | [optional]
**TaxBeforeDiscount**| **float32** | Tax before discount.  | [optional]
**TaxCanceled**| **float32** | Tax canceled.  | [optional]
**TaxInvoiced**| **float32** | Tax invoiced.  | [optional]
**TaxPercent**| **float32** | Tax percent.  | [optional]
**TaxRefunded**| **float32** | Tax refunded.  | [optional]
**UpdatedAt**| **string** | Updated-at timestamp.  | [optional]
**WeeeTaxApplied**| **string** | WEEE tax applied.  | [optional]
**WeeeTaxAppliedAmount**| **float32** | WEEE tax applied amount.  | [optional]
**WeeeTaxAppliedRowAmount**| **float32** | WEEE tax applied row amount.  | [optional]
**WeeeTaxDisposition**| **float32** | WEEE tax disposition.  | [optional]
**WeeeTaxRowDisposition**| **float32** | WEEE tax row disposition.  | [optional]
**Weight**| **float32** | Weight.  | [optional]
**ParentItem**| [**SalesDataOrderItemInterface**](SalesDataOrderItemInterface.md) |   | [optional]
**ProductOption**| [**CatalogDataProductOptionInterface**](CatalogDataProductOptionInterface.md) |   | [optional]
**ExtensionAttributes**| [**SalesDataOrderItemExtensionInterface**](SalesDataOrderItemExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
