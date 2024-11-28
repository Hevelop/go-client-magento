# \InventoryStockResolverTypeCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryStockresolverTypeCode**](InventoryStockResolverTypeCodeAPI.md#GetV1InventoryStockresolverTypeCode) | **Get** /V1/inventory/stock-resolver/{type}/{code} | inventory/stock-resolver/{type}/{code}



## GetV1InventoryStockresolverTypeCode

> InventoryApiDataStockInterface GetV1InventoryStockresolverTypeCode(ctx, type_, code).Execute()

inventory/stock-resolver/{type}/{code}



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
	type_ := "type__example" // string | 
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryStockResolverTypeCodeAPI.GetV1InventoryStockresolverTypeCode(context.Background(), type_, code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockResolverTypeCodeAPI.GetV1InventoryStockresolverTypeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryStockresolverTypeCode`: InventoryApiDataStockInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryStockResolverTypeCodeAPI.GetV1InventoryStockresolverTypeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryStockresolverTypeCodeRequest struct via the builder pattern


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

