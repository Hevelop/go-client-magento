# \BundleProductsOptionsAddAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1BundleproductsOptionsAdd**](BundleProductsOptionsAddAPI.md#PostV1BundleproductsOptionsAdd) | **Post** /V1/bundle-products/options/add | bundle-products/options/add



## PostV1BundleproductsOptionsAdd

> int32 PostV1BundleproductsOptionsAdd(ctx).PostV1BundleproductsOptionsAddRequest(postV1BundleproductsOptionsAddRequest).Execute()

bundle-products/options/add



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
	postV1BundleproductsOptionsAddRequest := *openapiclient.NewPostV1BundleproductsOptionsAddRequest(*openapiclient.NewBundleDataOptionInterface()) // PostV1BundleproductsOptionsAddRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsOptionsAddAPI.PostV1BundleproductsOptionsAdd(context.Background()).PostV1BundleproductsOptionsAddRequest(postV1BundleproductsOptionsAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsOptionsAddAPI.PostV1BundleproductsOptionsAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1BundleproductsOptionsAdd`: int32
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsOptionsAddAPI.PostV1BundleproductsOptionsAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1BundleproductsOptionsAddRequest struct via the builder pattern


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

