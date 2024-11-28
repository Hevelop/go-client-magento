# \AdobestockAssetIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1AdobestockAssetId**](AdobestockAssetIdAPI.md#DeleteV1AdobestockAssetId) | **Delete** /V1/adobestock/asset/{id} | adobestock/asset/{id}
[**GetV1AdobestockAssetId**](AdobestockAssetIdAPI.md#GetV1AdobestockAssetId) | **Get** /V1/adobestock/asset/{id} | adobestock/asset/{id}



## DeleteV1AdobestockAssetId

> ErrorResponse DeleteV1AdobestockAssetId(ctx, id).Execute()

adobestock/asset/{id}



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdobestockAssetIdAPI.DeleteV1AdobestockAssetId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdobestockAssetIdAPI.DeleteV1AdobestockAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1AdobestockAssetId`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `AdobestockAssetIdAPI.DeleteV1AdobestockAssetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1AdobestockAssetIdRequest struct via the builder pattern


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


## GetV1AdobestockAssetId

> AdobeStockAssetApiDataAssetInterface GetV1AdobestockAssetId(ctx, id).Execute()

adobestock/asset/{id}



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdobestockAssetIdAPI.GetV1AdobestockAssetId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdobestockAssetIdAPI.GetV1AdobestockAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AdobestockAssetId`: AdobeStockAssetApiDataAssetInterface
	fmt.Fprintf(os.Stdout, "Response from `AdobestockAssetIdAPI.GetV1AdobestockAssetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AdobestockAssetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdobeStockAssetApiDataAssetInterface**](AdobeStockAssetApiDataAssetInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

