# # CatalogDataProductRenderFormattedPriceInfoInterface
Formatted Price interface. Aggregate formatted html with price representations. E.g.: &lt;span class&#x3D;\&quot;price\&quot;&gt;$9.00&lt;/span&gt; Consider currency, rounding and html

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinalPrice**| **string** | Html with final price  |
**MaxPrice**| **string** | Max price of a product  |
**MinimalPrice**| **string** | The minimal price of the product or variation  |
**MaxRegularPrice**| **string** | Max regular price  |
**MinimalRegularPrice**| **string** | Minimal regular price  |
**SpecialPrice**| **string** | Special price  |
**RegularPrice**| **string** | Price - is price of product without discounts and special price with taxes and fixed product tax  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Catalog\\Api\\Data\\ProductRender\\FormattedPriceInfoInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

