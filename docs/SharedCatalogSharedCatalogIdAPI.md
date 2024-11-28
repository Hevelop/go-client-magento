# \SharedCatalogSharedCatalogIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1SharedCatalogSharedCatalogId**](SharedCatalogSharedCatalogIdAPI.md#DeleteV1SharedCatalogSharedCatalogId) | **Delete** /V1/sharedCatalog/{sharedCatalogId} | sharedCatalog/{sharedCatalogId}
[**GetV1SharedCatalogSharedCatalogId**](SharedCatalogSharedCatalogIdAPI.md#GetV1SharedCatalogSharedCatalogId) | **Get** /V1/sharedCatalog/{sharedCatalogId} | sharedCatalog/{sharedCatalogId}



## DeleteV1SharedCatalogSharedCatalogId

> bool DeleteV1SharedCatalogSharedCatalogId(ctx, sharedCatalogId).Execute()

sharedCatalog/{sharedCatalogId}



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
	sharedCatalogId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogSharedCatalogIdAPI.DeleteV1SharedCatalogSharedCatalogId(context.Background(), sharedCatalogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogSharedCatalogIdAPI.DeleteV1SharedCatalogSharedCatalogId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1SharedCatalogSharedCatalogId`: bool
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogSharedCatalogIdAPI.DeleteV1SharedCatalogSharedCatalogId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedCatalogId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1SharedCatalogSharedCatalogIdRequest struct via the builder pattern


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


## GetV1SharedCatalogSharedCatalogId

> SharedCatalogDataSharedCatalogInterface GetV1SharedCatalogSharedCatalogId(ctx, sharedCatalogId).Execute()

sharedCatalog/{sharedCatalogId}



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
	sharedCatalogId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogSharedCatalogIdAPI.GetV1SharedCatalogSharedCatalogId(context.Background(), sharedCatalogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogSharedCatalogIdAPI.GetV1SharedCatalogSharedCatalogId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1SharedCatalogSharedCatalogId`: SharedCatalogDataSharedCatalogInterface
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogSharedCatalogIdAPI.GetV1SharedCatalogSharedCatalogId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedCatalogId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1SharedCatalogSharedCatalogIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SharedCatalogDataSharedCatalogInterface**](SharedCatalogDataSharedCatalogInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

