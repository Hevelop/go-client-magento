# \ProductsDownloadableLinksIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsDownloadablelinksId**](ProductsDownloadableLinksIdAPI.md#DeleteV1ProductsDownloadablelinksId) | **Delete** /V1/products/downloadable-links/{id} | products/downloadable-links/{id}



## DeleteV1ProductsDownloadablelinksId

> bool DeleteV1ProductsDownloadablelinksId(ctx, id).Execute()

products/downloadable-links/{id}



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
	resp, r, err := apiClient.ProductsDownloadableLinksIdAPI.DeleteV1ProductsDownloadablelinksId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsDownloadableLinksIdAPI.DeleteV1ProductsDownloadablelinksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsDownloadablelinksId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsDownloadableLinksIdAPI.DeleteV1ProductsDownloadablelinksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsDownloadablelinksIdRequest struct via the builder pattern


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

