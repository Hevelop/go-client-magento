# \CmsBlockBlockIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CmsBlockBlockId**](CmsBlockBlockIdAPI.md#DeleteV1CmsBlockBlockId) | **Delete** /V1/cmsBlock/{blockId} | cmsBlock/{blockId}
[**GetV1CmsBlockBlockId**](CmsBlockBlockIdAPI.md#GetV1CmsBlockBlockId) | **Get** /V1/cmsBlock/{blockId} | cmsBlock/{blockId}



## DeleteV1CmsBlockBlockId

> bool DeleteV1CmsBlockBlockId(ctx, blockId).Execute()

cmsBlock/{blockId}



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
	blockId := "blockId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsBlockBlockIdAPI.DeleteV1CmsBlockBlockId(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsBlockBlockIdAPI.DeleteV1CmsBlockBlockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CmsBlockBlockId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CmsBlockBlockIdAPI.DeleteV1CmsBlockBlockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CmsBlockBlockIdRequest struct via the builder pattern


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


## GetV1CmsBlockBlockId

> CmsDataBlockInterface GetV1CmsBlockBlockId(ctx, blockId).Execute()

cmsBlock/{blockId}



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
	blockId := "blockId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsBlockBlockIdAPI.GetV1CmsBlockBlockId(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsBlockBlockIdAPI.GetV1CmsBlockBlockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CmsBlockBlockId`: CmsDataBlockInterface
	fmt.Fprintf(os.Stdout, "Response from `CmsBlockBlockIdAPI.GetV1CmsBlockBlockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CmsBlockBlockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CmsDataBlockInterface**](CmsDataBlockInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

