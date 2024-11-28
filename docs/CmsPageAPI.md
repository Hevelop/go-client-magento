# \CmsPageAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CmsPage**](CmsPageAPI.md#PostV1CmsPage) | **Post** /V1/cmsPage | cmsPage



## PostV1CmsPage

> CmsDataPageInterface PostV1CmsPage(ctx).PostV1CmsPageRequest(postV1CmsPageRequest).Execute()

cmsPage



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
	postV1CmsPageRequest := *openapiclient.NewPostV1CmsPageRequest(*openapiclient.NewCmsDataPageInterface("Identifier_example")) // PostV1CmsPageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsPageAPI.PostV1CmsPage(context.Background()).PostV1CmsPageRequest(postV1CmsPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsPageAPI.PostV1CmsPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CmsPage`: CmsDataPageInterface
	fmt.Fprintf(os.Stdout, "Response from `CmsPageAPI.PostV1CmsPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CmsPageRequest struct via the builder pattern


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

