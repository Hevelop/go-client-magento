# \CmsPagePageIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CmsPagePageId**](CmsPagePageIdAPI.md#DeleteV1CmsPagePageId) | **Delete** /V1/cmsPage/{pageId} | cmsPage/{pageId}
[**GetV1CmsPagePageId**](CmsPagePageIdAPI.md#GetV1CmsPagePageId) | **Get** /V1/cmsPage/{pageId} | cmsPage/{pageId}



## DeleteV1CmsPagePageId

> bool DeleteV1CmsPagePageId(ctx, pageId).Execute()

cmsPage/{pageId}



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
	pageId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsPagePageIdAPI.DeleteV1CmsPagePageId(context.Background(), pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsPagePageIdAPI.DeleteV1CmsPagePageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CmsPagePageId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CmsPagePageIdAPI.DeleteV1CmsPagePageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CmsPagePageIdRequest struct via the builder pattern


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


## GetV1CmsPagePageId

> CmsDataPageInterface GetV1CmsPagePageId(ctx, pageId).Execute()

cmsPage/{pageId}



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
	pageId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsPagePageIdAPI.GetV1CmsPagePageId(context.Background(), pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsPagePageIdAPI.GetV1CmsPagePageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CmsPagePageId`: CmsDataPageInterface
	fmt.Fprintf(os.Stdout, "Response from `CmsPagePageIdAPI.GetV1CmsPagePageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CmsPagePageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CmsDataPageInterface**](CmsDataPageInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

