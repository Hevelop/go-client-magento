# \CmsPageIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CmsPageId**](CmsPageIdAPI.md#PutV1CmsPageId) | **Put** /V1/cmsPage/{id} | cmsPage/{id}



## PutV1CmsPageId

> CmsDataPageInterface PutV1CmsPageId(ctx, id).PostV1CmsPageRequest(postV1CmsPageRequest).Execute()

cmsPage/{id}



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
	id := "id_example" // string | 
	postV1CmsPageRequest := *openapiclient.NewPostV1CmsPageRequest(*openapiclient.NewCmsDataPageInterface("Identifier_example")) // PostV1CmsPageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsPageIdAPI.PutV1CmsPageId(context.Background(), id).PostV1CmsPageRequest(postV1CmsPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsPageIdAPI.PutV1CmsPageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CmsPageId`: CmsDataPageInterface
	fmt.Fprintf(os.Stdout, "Response from `CmsPageIdAPI.PutV1CmsPageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CmsPageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CmsPageRequest** | [**PostV1CmsPageRequest**](PostV1CmsPageRequest.md) |  | 

### Return type

[**CmsDataPageInterface**](CmsDataPageInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

