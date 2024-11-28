# \BundleProductsSkuLinksOptionIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1BundleproductsSkuLinksOptionId**](BundleProductsSkuLinksOptionIdAPI.md#PostV1BundleproductsSkuLinksOptionId) | **Post** /V1/bundle-products/{sku}/links/{optionId} | bundle-products/{sku}/links/{optionId}



## PostV1BundleproductsSkuLinksOptionId

> int32 PostV1BundleproductsSkuLinksOptionId(ctx, sku, optionId).PutV1BundleproductsSkuLinksIdRequest(putV1BundleproductsSkuLinksIdRequest).Execute()

bundle-products/{sku}/links/{optionId}



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
	optionId := int32(56) // int32 | 
	putV1BundleproductsSkuLinksIdRequest := *openapiclient.NewPutV1BundleproductsSkuLinksIdRequest(*openapiclient.NewBundleDataLinkInterface(false, float32(123), int32(123))) // PutV1BundleproductsSkuLinksIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsSkuLinksOptionIdAPI.PostV1BundleproductsSkuLinksOptionId(context.Background(), sku, optionId).PutV1BundleproductsSkuLinksIdRequest(putV1BundleproductsSkuLinksIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsSkuLinksOptionIdAPI.PostV1BundleproductsSkuLinksOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1BundleproductsSkuLinksOptionId`: int32
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsSkuLinksOptionIdAPI.PostV1BundleproductsSkuLinksOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1BundleproductsSkuLinksOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putV1BundleproductsSkuLinksIdRequest** | [**PutV1BundleproductsSkuLinksIdRequest**](PutV1BundleproductsSkuLinksIdRequest.md) |  | 

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

