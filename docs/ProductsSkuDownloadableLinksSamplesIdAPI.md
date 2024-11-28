# \ProductsSkuDownloadableLinksSamplesIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1ProductsSkuDownloadablelinksSamplesId**](ProductsSkuDownloadableLinksSamplesIdAPI.md#PutV1ProductsSkuDownloadablelinksSamplesId) | **Put** /V1/products/{sku}/downloadable-links/samples/{id} | products/{sku}/downloadable-links/samples/{id}



## PutV1ProductsSkuDownloadablelinksSamplesId

> int32 PutV1ProductsSkuDownloadablelinksSamplesId(ctx, sku, id).PostV1ProductsSkuDownloadablelinksSamplesRequest(postV1ProductsSkuDownloadablelinksSamplesRequest).Execute()

products/{sku}/downloadable-links/samples/{id}



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
	postV1ProductsSkuDownloadablelinksSamplesRequest := *openapiclient.NewPostV1ProductsSkuDownloadablelinksSamplesRequest(*openapiclient.NewDownloadableDataSampleInterface("Title_example", int32(123), "SampleType_example")) // PostV1ProductsSkuDownloadablelinksSamplesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuDownloadableLinksSamplesIdAPI.PutV1ProductsSkuDownloadablelinksSamplesId(context.Background(), sku, id).PostV1ProductsSkuDownloadablelinksSamplesRequest(postV1ProductsSkuDownloadablelinksSamplesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuDownloadableLinksSamplesIdAPI.PutV1ProductsSkuDownloadablelinksSamplesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsSkuDownloadablelinksSamplesId`: int32
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuDownloadableLinksSamplesIdAPI.PutV1ProductsSkuDownloadablelinksSamplesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsSkuDownloadablelinksSamplesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postV1ProductsSkuDownloadablelinksSamplesRequest** | [**PostV1ProductsSkuDownloadablelinksSamplesRequest**](PostV1ProductsSkuDownloadablelinksSamplesRequest.md) |  | 

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

