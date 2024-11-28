# \InventoryGetProductSalableQuantitySkuStockIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryGetproductsalablequantitySkuStockId**](InventoryGetProductSalableQuantitySkuStockIdAPI.md#GetV1InventoryGetproductsalablequantitySkuStockId) | **Get** /V1/inventory/get-product-salable-quantity/{sku}/{stockId} | inventory/get-product-salable-quantity/{sku}/{stockId}



## GetV1InventoryGetproductsalablequantitySkuStockId

> float32 GetV1InventoryGetproductsalablequantitySkuStockId(ctx, sku, stockId).Execute()

inventory/get-product-salable-quantity/{sku}/{stockId}



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
	resp, r, err := apiClient.InventoryGetProductSalableQuantitySkuStockIdAPI.GetV1InventoryGetproductsalablequantitySkuStockId(context.Background(), sku, stockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGetProductSalableQuantitySkuStockIdAPI.GetV1InventoryGetproductsalablequantitySkuStockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryGetproductsalablequantitySkuStockId`: float32
	fmt.Fprintf(os.Stdout, "Response from `InventoryGetProductSalableQuantitySkuStockIdAPI.GetV1InventoryGetproductsalablequantitySkuStockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**stockId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryGetproductsalablequantitySkuStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

