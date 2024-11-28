# \CmsBlockAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CmsBlock**](CmsBlockAPI.md#PostV1CmsBlock) | **Post** /V1/cmsBlock | cmsBlock



## PostV1CmsBlock

> CmsDataBlockInterface PostV1CmsBlock(ctx).PostV1CmsBlockRequest(postV1CmsBlockRequest).Execute()

cmsBlock



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
	postV1CmsBlockRequest := *openapiclient.NewPostV1CmsBlockRequest(*openapiclient.NewCmsDataBlockInterface("Identifier_example")) // PostV1CmsBlockRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsBlockAPI.PostV1CmsBlock(context.Background()).PostV1CmsBlockRequest(postV1CmsBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsBlockAPI.PostV1CmsBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CmsBlock`: CmsDataBlockInterface
	fmt.Fprintf(os.Stdout, "Response from `CmsBlockAPI.PostV1CmsBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CmsBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CmsBlockRequest** | [**PostV1CmsBlockRequest**](PostV1CmsBlockRequest.md) |  | 

### Return type

[**CmsDataBlockInterface**](CmsDataBlockInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

