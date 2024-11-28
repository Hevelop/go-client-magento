# \StockStatusesProductSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1StockStatusesProductSku**](StockStatusesProductSkuAPI.md#GetV1StockStatusesProductSku) | **Get** /V1/stockStatuses/{productSku} | stockStatuses/{productSku}



## GetV1StockStatusesProductSku

> CatalogInventoryDataStockStatusInterface GetV1StockStatusesProductSku(ctx, productSku).ScopeId(scopeId).Execute()

stockStatuses/{productSku}



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
	resp, r, err := apiClient.StockStatusesProductSkuAPI.GetV1StockStatusesProductSku(context.Background(), productSku).ScopeId(scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockStatusesProductSkuAPI.GetV1StockStatusesProductSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1StockStatusesProductSku`: CatalogInventoryDataStockStatusInterface
	fmt.Fprintf(os.Stdout, "Response from `StockStatusesProductSkuAPI.GetV1StockStatusesProductSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productSku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1StockStatusesProductSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeId** | **int32** |  | 

### Return type

[**CatalogInventoryDataStockStatusInterface**](CatalogInventoryDataStockStatusInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

