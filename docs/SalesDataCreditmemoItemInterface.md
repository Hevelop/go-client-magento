# # SalesDataCreditmemoItemInterface
Credit memo item interface. After a customer places and pays for an order and an invoice has been issued, the merchant can create a credit memo to refund all or part of the amount paid for any returned or undelivered items. The memo restores funds to the customer account so that the customer can make future purchases. A credit memo item is an invoiced item for which a merchant creates a credit memo.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData**| **string** | Additional data.  | [optional]
**BaseCost**| **float32** | The base cost for a credit memo item.  |
**BaseDiscountAmount**| **float32** | The base discount amount for a credit memo item.  | [optional]
**BaseDiscountTaxCompensationAmount**| **float32** | The base discount tax compensation amount for a credit memo item.  | [optional]
**BasePrice**| **float32** | The base price for a credit memo item.  |
**BasePriceInclTax**| **float32** | Base price including tax.  | [optional]
**BaseRowTotal**| **float32** | Base row total.  | [optional]
**BaseRowTotalInclTax**| **float32** | Base row total including tax.  | [optional]
**BaseTaxAmount**| **float32** | Base tax amount.  | [optional]
**BaseWeeeTaxAppliedAmount**| **float32** | Base WEEE tax applied amount.  | [optional]
**BaseWeeeTaxAppliedRowAmnt**| **float32** | Base WEEE tax applied row amount.  | [optional]
**BaseWeeeTaxDisposition**| **float32** | Base WEEE tax disposition.  | [optional]
**BaseWeeeTaxRowDisposition**| **float32** | Base WEEE tax row disposition.  | [optional]
**Description**| **string** | Description.  | [optional]
**DiscountAmount**| **float32** | Discount amount.  | [optional]
**EntityId**| **int32** | Credit memo item ID.  |
**DiscountTaxCompensationAmount**| **float32** | Discount tax compensation amount.  | [optional]
**Name**| **string** | Name.  | [optional]
**OrderItemId**| **int32** | Order item ID.  |
**ParentId**| **int32** | Parent ID.  | [optional]
**Price**| **float32** | Price.  | [optional]
**PriceInclTax**| **float32** | Price including tax.  | [optional]
**ProductId**| **int32** | Product ID.  | [optional]
**Qty**| **float32** | Quantity.  |
**RowTotal**| **float32** | Row total.  | [optional]
**RowTotalInclTax**| **float32** | Row total including tax.  | [optional]
**Sku**| **string** | SKU.  | [optional]
**TaxAmount**| **float32** | Tax amount.  | [optional]
**WeeeTaxApplied**| **string** | WEEE tax applied.  | [optional]
**WeeeTaxAppliedAmount**| **float32** | WEEE tax applied amount.  | [optional]
**WeeeTaxAppliedRowAmount**| **float32** | WEEE tax applied row amount.  | [optional]
**WeeeTaxDisposition**| **float32** | WEEE tax disposition.  | [optional]
**WeeeTaxRowDisposition**| **float32** | WEEE tax row disposition.  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\CreditmemoItemInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

