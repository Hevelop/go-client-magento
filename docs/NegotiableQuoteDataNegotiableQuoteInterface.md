# # NegotiableQuoteDataNegotiableQuoteInterface
Interface NegotiableQuoteInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId**| **int32** | Negotiable quote ID.  |
**IsRegularQuote**| **bool** | Is regular quote.  |
**Status**| **string** | Negotiable quote status.  |
**NegotiatedPriceType**| **int32** | Negotiated price type.  |
**NegotiatedPriceValue**| **float32** | Negotiated price value.  |
**ShippingPrice**| **float32** | Proposed shipping price.  |
**QuoteName**| **string** | Negotiable quote name.  |
**ExpirationPeriod**| **string** | Expiration period.  |
**EmailNotificationStatus**| **int32** | Email notification status.  |
**HasUnconfirmedChanges**| **bool** | Has unconfirmed changes.  |
**IsShippingTaxChanged**| **bool** | Shipping tax changes.  |
**IsCustomerPriceChanged**| **bool** | Customer price changes.  |
**Notifications**| **int32** | Quote notifications.  |
**AppliedRuleIds**| **string** | Quote rules.  |
**IsAddressDraft**| **bool** | Is address draft.  |
**DeletedSku**| **string** | Deleted products sku.  |
**CreatorId**| **int32** | Quote creator id.  |
**CreatorType**| **int32** | Quote creator type.  |
**OriginalTotalPrice**| **float32** | Quote original total price.  | [optional]
**BaseOriginalTotalPrice**| **float32** | Quote original total price in base currency.  | [optional]
**NegotiatedTotalPrice**| **float32** | Quote negotiated total price.  | [optional]
**BaseNegotiatedTotalPrice**| **float32** | Quote negotiated total price in base currency.  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\NegotiableQuote\\Api\\Data\\NegotiableQuoteInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

