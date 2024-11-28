# \StockItemsProductSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1StockItemsProductSku**](StockItemsProductSkuAPI.md#GetV1StockItemsProductSku) | **Get** /V1/stockItems/{productSku} | stockItems/{productSku}



## GetV1StockItemsProductSku

> CatalogInventoryDataStockItemInterface GetV1StockItemsProductSku(ctx, productSku).ScopeId(scopeId).Execute()

stockItems/{productSku}



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-magento"
)

func main() {
	productSku := "productSku_example" // string | 
	scopeId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StockItemsProductSkuAPI.GetV1StockItemsProductSku(context.Background(), productSku).ScopeId(scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockItemsProductSkuAPI.GetV1StockItemsProductSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1StockItemsProductSku`: CatalogInventoryDataStockItemInterface
	fmt.Fprintf(os.Stdout, "Response from `StockItemsProductSkuAPI.GetV1StockItemsProductSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productSku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1StockItemsProductSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeId** | **int32** |  | 

### Return type

[**CatalogInventoryDataStockItemInterface**](CatalogInventoryDataStockItemInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

