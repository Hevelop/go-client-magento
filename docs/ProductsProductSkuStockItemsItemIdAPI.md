# \ProductsProductSkuStockItemsItemIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1ProductsProductSkuStockItemsItemId**](ProductsProductSkuStockItemsItemIdAPI.md#PutV1ProductsProductSkuStockItemsItemId) | **Put** /V1/products/{productSku}/stockItems/{itemId} | products/{productSku}/stockItems/{itemId}



## PutV1ProductsProductSkuStockItemsItemId

> int32 PutV1ProductsProductSkuStockItemsItemId(ctx, productSku, itemId).PutV1ProductsProductSkuStockItemsItemIdRequest(putV1ProductsProductSkuStockItemsItemIdRequest).Execute()

products/{productSku}/stockItems/{itemId}



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
	itemId := "itemId_example" // string | 
	putV1ProductsProductSkuStockItemsItemIdRequest := *openapiclient.NewPutV1ProductsProductSkuStockItemsItemIdRequest(*openapiclient.NewCatalogInventoryDataStockItemInterface(float32(123), false, false, false, false, float32(123), int32(123), float32(123), false, float32(123), false, int32(123), false, float32(123), false, float32(123), false, false, false, false, "LowStockDate_example", false, int32(123))) // PutV1ProductsProductSkuStockItemsItemIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsProductSkuStockItemsItemIdAPI.PutV1ProductsProductSkuStockItemsItemId(context.Background(), productSku, itemId).PutV1ProductsProductSkuStockItemsItemIdRequest(putV1ProductsProductSkuStockItemsItemIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsProductSkuStockItemsItemIdAPI.PutV1ProductsProductSkuStockItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsProductSkuStockItemsItemId`: int32
	fmt.Fprintf(os.Stdout, "Response from `ProductsProductSkuStockItemsItemIdAPI.PutV1ProductsProductSkuStockItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productSku** | **string** |  | 
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsProductSkuStockItemsItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putV1ProductsProductSkuStockItemsItemIdRequest** | [**PutV1ProductsProductSkuStockItemsItemIdRequest**](PutV1ProductsProductSkuStockItemsItemIdRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

