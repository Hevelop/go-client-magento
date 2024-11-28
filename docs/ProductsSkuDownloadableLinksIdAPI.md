# \ProductsSkuDownloadableLinksIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1ProductsSkuDownloadablelinksId**](ProductsSkuDownloadableLinksIdAPI.md#PutV1ProductsSkuDownloadablelinksId) | **Put** /V1/products/{sku}/downloadable-links/{id} | products/{sku}/downloadable-links/{id}



## PutV1ProductsSkuDownloadablelinksId

> int32 PutV1ProductsSkuDownloadablelinksId(ctx, sku, id).PostV1ProductsSkuDownloadablelinksRequest(postV1ProductsSkuDownloadablelinksRequest).Execute()

products/{sku}/downloadable-links/{id}



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
	sku := "sku_example" // string | 
	id := "id_example" // string | 
	postV1ProductsSkuDownloadablelinksRequest := *openapiclient.NewPostV1ProductsSkuDownloadablelinksRequest(*openapiclient.NewDownloadableDataLinkInterface(int32(123), int32(123), float32(123), "LinkType_example", "SampleType_example")) // PostV1ProductsSkuDownloadablelinksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuDownloadableLinksIdAPI.PutV1ProductsSkuDownloadablelinksId(context.Background(), sku, id).PostV1ProductsSkuDownloadablelinksRequest(postV1ProductsSkuDownloadablelinksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuDownloadableLinksIdAPI.PutV1ProductsSkuDownloadablelinksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsSkuDownloadablelinksId`: int32
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuDownloadableLinksIdAPI.PutV1ProductsSkuDownloadablelinksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsSkuDownloadablelinksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postV1ProductsSkuDownloadablelinksRequest** | [**PostV1ProductsSkuDownloadablelinksRequest**](PostV1ProductsSkuDownloadablelinksRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

