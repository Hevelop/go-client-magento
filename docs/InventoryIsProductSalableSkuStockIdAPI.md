# \InventoryIsProductSalableSkuStockIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryIsproductsalableSkuStockId**](InventoryIsProductSalableSkuStockIdAPI.md#GetV1InventoryIsproductsalableSkuStockId) | **Get** /V1/inventory/is-product-salable/{sku}/{stockId} | inventory/is-product-salable/{sku}/{stockId}



## GetV1InventoryIsproductsalableSkuStockId

> bool GetV1InventoryIsproductsalableSkuStockId(ctx, sku, stockId).Execute()

inventory/is-product-salable/{sku}/{stockId}



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
	sku := "sku_example" // string | 
	stockId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryIsProductSalableSkuStockIdAPI.GetV1InventoryIsproductsalableSkuStockId(context.Background(), sku, stockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryIsProductSalableSkuStockIdAPI.GetV1InventoryIsproductsalableSkuStockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryIsproductsalableSkuStockId`: bool
	fmt.Fprintf(os.Stdout, "Response from `InventoryIsProductSalableSkuStockIdAPI.GetV1InventoryIsproductsalableSkuStockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**stockId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryIsproductsalableSkuStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

