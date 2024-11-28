# # CatalogDataProductRenderInterface
Represents Data Object which holds enough information to render product This information is put into part as Add To Cart or Add to Compare Data or Price Data

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToCartButton**| [**CatalogDataProductRenderButtonInterface**](CatalogDataProductRenderButtonInterface.md) |   |
**AddToCompareButton**| [**CatalogDataProductRenderButtonInterface**](CatalogDataProductRenderButtonInterface.md) |   |
**PriceInfo**| [**CatalogDataProductRenderPriceInfoInterface**](CatalogDataProductRenderPriceInfoInterface.md) |   |
**Images**| [**[]CatalogDataProductRenderImageInterface**](CatalogDataProductRenderImageInterface.md) | Enough information, that needed to render image on front  |
**Url**| **string** | Product url  |
**Id**| **int32** | Product identifier  |
**Name**| **string** | Product name  |
**Type**| **string** | Product type. Such as bundle, grouped, simple, etc...  |
**IsSalable**| **string** | Information about product saleability (In Stock)  |
**StoreId**| **int32** | Information about current store id or requested store id  |
**CurrencyCode**| **string** | Current or desired currency code to product  |
**ExtensionAttributes**| [**CatalogDataProductRenderExtensionInterface**](CatalogDataProductRenderExtensionInterface.md) |   |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

