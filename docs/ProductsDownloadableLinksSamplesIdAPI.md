# \ProductsDownloadableLinksSamplesIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsDownloadablelinksSamplesId**](ProductsDownloadableLinksSamplesIdAPI.md#DeleteV1ProductsDownloadablelinksSamplesId) | **Delete** /V1/products/downloadable-links/samples/{id} | products/downloadable-links/samples/{id}



## DeleteV1ProductsDownloadablelinksSamplesId

> bool DeleteV1ProductsDownloadablelinksSamplesId(ctx, id).Execute()

products/downloadable-links/samples/{id}



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
	resp, r, err := apiClient.ProductsDownloadableLinksSamplesIdAPI.DeleteV1ProductsDownloadablelinksSamplesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsDownloadableLinksSamplesIdAPI.DeleteV1ProductsDownloadablelinksSamplesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsDownloadablelinksSamplesId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsDownloadableLinksSamplesIdAPI.DeleteV1ProductsDownloadablelinksSamplesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsDownloadablelinksSamplesIdRequest struct via the builder pattern


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

