# \BundleProductsSkuLinksIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1BundleproductsSkuLinksId**](BundleProductsSkuLinksIdAPI.md#PutV1BundleproductsSkuLinksId) | **Put** /V1/bundle-products/{sku}/links/{id} | bundle-products/{sku}/links/{id}



## PutV1BundleproductsSkuLinksId

> bool PutV1BundleproductsSkuLinksId(ctx, sku, id).PutV1BundleproductsSkuLinksIdRequest(putV1BundleproductsSkuLinksIdRequest).Execute()

bundle-products/{sku}/links/{id}



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
	putV1BundleproductsSkuLinksIdRequest := *openapiclient.NewPutV1BundleproductsSkuLinksIdRequest(*openapiclient.NewBundleDataLinkInterface(false, float32(123), int32(123))) // PutV1BundleproductsSkuLinksIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsSkuLinksIdAPI.PutV1BundleproductsSkuLinksId(context.Background(), sku, id).PutV1BundleproductsSkuLinksIdRequest(putV1BundleproductsSkuLinksIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsSkuLinksIdAPI.PutV1BundleproductsSkuLinksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1BundleproductsSkuLinksId`: bool
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsSkuLinksIdAPI.PutV1BundleproductsSkuLinksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1BundleproductsSkuLinksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putV1BundleproductsSkuLinksIdRequest** | [**PutV1BundleproductsSkuLinksIdRequest**](PutV1BundleproductsSkuLinksIdRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

