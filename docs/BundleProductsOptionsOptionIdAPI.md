# \BundleProductsOptionsOptionIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1BundleproductsOptionsOptionId**](BundleProductsOptionsOptionIdAPI.md#PutV1BundleproductsOptionsOptionId) | **Put** /V1/bundle-products/options/{optionId} | bundle-products/options/{optionId}



## PutV1BundleproductsOptionsOptionId

> int32 PutV1BundleproductsOptionsOptionId(ctx, optionId).PostV1BundleproductsOptionsAddRequest(postV1BundleproductsOptionsAddRequest).Execute()

bundle-products/options/{optionId}



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
	optionId := "optionId_example" // string | 
	postV1BundleproductsOptionsAddRequest := *openapiclient.NewPostV1BundleproductsOptionsAddRequest(*openapiclient.NewBundleDataOptionInterface()) // PostV1BundleproductsOptionsAddRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsOptionsOptionIdAPI.PutV1BundleproductsOptionsOptionId(context.Background(), optionId).PostV1BundleproductsOptionsAddRequest(postV1BundleproductsOptionsAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsOptionsOptionIdAPI.PutV1BundleproductsOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1BundleproductsOptionsOptionId`: int32
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsOptionsOptionIdAPI.PutV1BundleproductsOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1BundleproductsOptionsOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1BundleproductsOptionsAddRequest** | [**PostV1BundleproductsOptionsAddRequest**](PostV1BundleproductsOptionsAddRequest.md) |  | 

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

