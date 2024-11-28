# \InventoryStocksStockIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1InventoryStocksStockId**](InventoryStocksStockIdAPI.md#DeleteV1InventoryStocksStockId) | **Delete** /V1/inventory/stocks/{stockId} | inventory/stocks/{stockId}
[**GetV1InventoryStocksStockId**](InventoryStocksStockIdAPI.md#GetV1InventoryStocksStockId) | **Get** /V1/inventory/stocks/{stockId} | inventory/stocks/{stockId}
[**PutV1InventoryStocksStockId**](InventoryStocksStockIdAPI.md#PutV1InventoryStocksStockId) | **Put** /V1/inventory/stocks/{stockId} | inventory/stocks/{stockId}



## DeleteV1InventoryStocksStockId

> ErrorResponse DeleteV1InventoryStocksStockId(ctx, stockId).Execute()

inventory/stocks/{stockId}



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
	stockId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryStocksStockIdAPI.DeleteV1InventoryStocksStockId(context.Background(), stockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryStocksStockIdAPI.DeleteV1InventoryStocksStockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1InventoryStocksStockId`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryStocksStockIdAPI.DeleteV1InventoryStocksStockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1InventoryStocksStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1InventoryStocksStockId

> InventoryApiDataStockInterface GetV1InventoryStocksStockId(ctx, stockId).Execute()

inventory/stocks/{stockId}



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
	stockId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryStocksStockIdAPI.GetV1InventoryStocksStockId(context.Background(), stockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryStocksStockIdAPI.GetV1InventoryStocksStockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryStocksStockId`: InventoryApiDataStockInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryStocksStockIdAPI.GetV1InventoryStocksStockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryStocksStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryApiDataStockInterface**](InventoryApiDataStockInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1InventoryStocksStockId

> int32 PutV1InventoryStocksStockId(ctx, stockId).PostV1InventoryStocksRequest(postV1InventoryStocksRequest).Execute()

inventory/stocks/{stockId}



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
	stockId := "stockId_example" // string | 
	postV1InventoryStocksRequest := *openapiclient.NewPostV1InventoryStocksRequest(*openapiclient.NewInventoryApiDataStockInterface()) // PostV1InventoryStocksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryStocksStockIdAPI.PutV1InventoryStocksStockId(context.Background(), stockId).PostV1InventoryStocksRequest(postV1InventoryStocksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryStocksStockIdAPI.PutV1InventoryStocksStockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1InventoryStocksStockId`: int32
	fmt.Fprintf(os.Stdout, "Response from `InventoryStocksStockIdAPI.PutV1InventoryStocksStockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1InventoryStocksStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1InventoryStocksRequest** | [**PostV1InventoryStocksRequest**](PostV1InventoryStocksRequest.md) |  | 

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

