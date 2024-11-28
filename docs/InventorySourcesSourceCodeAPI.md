# \InventorySourcesSourceCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventorySourcesSourceCode**](InventorySourcesSourceCodeAPI.md#GetV1InventorySourcesSourceCode) | **Get** /V1/inventory/sources/{sourceCode} | inventory/sources/{sourceCode}
[**PutV1InventorySourcesSourceCode**](InventorySourcesSourceCodeAPI.md#PutV1InventorySourcesSourceCode) | **Put** /V1/inventory/sources/{sourceCode} | inventory/sources/{sourceCode}



## GetV1InventorySourcesSourceCode

> InventoryApiDataSourceInterface GetV1InventorySourcesSourceCode(ctx, sourceCode).Execute()

inventory/sources/{sourceCode}



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
	sourceCode := "sourceCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventorySourcesSourceCodeAPI.GetV1InventorySourcesSourceCode(context.Background(), sourceCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesSourceCodeAPI.GetV1InventorySourcesSourceCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventorySourcesSourceCode`: InventoryApiDataSourceInterface
	fmt.Fprintf(os.Stdout, "Response from `InventorySourcesSourceCodeAPI.GetV1InventorySourcesSourceCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventorySourcesSourceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryApiDataSourceInterface**](InventoryApiDataSourceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1InventorySourcesSourceCode

> ErrorResponse PutV1InventorySourcesSourceCode(ctx, sourceCode).PostV1InventorySourcesRequest(postV1InventorySourcesRequest).Execute()

inventory/sources/{sourceCode}



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
	sourceCode := "sourceCode_example" // string | 
	postV1InventorySourcesRequest := *openapiclient.NewPostV1InventorySourcesRequest(*openapiclient.NewInventoryApiDataSourceInterface()) // PostV1InventorySourcesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventorySourcesSourceCodeAPI.PutV1InventorySourcesSourceCode(context.Background(), sourceCode).PostV1InventorySourcesRequest(postV1InventorySourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventorySourcesSourceCodeAPI.PutV1InventorySourcesSourceCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1InventorySourcesSourceCode`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `InventorySourcesSourceCodeAPI.PutV1InventorySourcesSourceCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1InventorySourcesSourceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1InventorySourcesRequest** | [**PostV1InventorySourcesRequest**](PostV1InventorySourcesRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

