# \BundleProductsSkuOptionsOptionIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1BundleproductsSkuOptionsOptionId**](BundleProductsSkuOptionsOptionIdAPI.md#DeleteV1BundleproductsSkuOptionsOptionId) | **Delete** /V1/bundle-products/{sku}/options/{optionId} | bundle-products/{sku}/options/{optionId}
[**GetV1BundleproductsSkuOptionsOptionId**](BundleProductsSkuOptionsOptionIdAPI.md#GetV1BundleproductsSkuOptionsOptionId) | **Get** /V1/bundle-products/{sku}/options/{optionId} | bundle-products/{sku}/options/{optionId}



## DeleteV1BundleproductsSkuOptionsOptionId

> bool DeleteV1BundleproductsSkuOptionsOptionId(ctx, sku, optionId).Execute()

bundle-products/{sku}/options/{optionId}



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsSkuOptionsOptionIdAPI.DeleteV1BundleproductsSkuOptionsOptionId(context.Background(), sku, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsSkuOptionsOptionIdAPI.DeleteV1BundleproductsSkuOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1BundleproductsSkuOptionsOptionId`: bool
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsSkuOptionsOptionIdAPI.DeleteV1BundleproductsSkuOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1BundleproductsSkuOptionsOptionIdRequest struct via the builder pattern


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


## GetV1BundleproductsSkuOptionsOptionId

> BundleDataOptionInterface GetV1BundleproductsSkuOptionsOptionId(ctx, sku, optionId).Execute()

bundle-products/{sku}/options/{optionId}



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsSkuOptionsOptionIdAPI.GetV1BundleproductsSkuOptionsOptionId(context.Background(), sku, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsSkuOptionsOptionIdAPI.GetV1BundleproductsSkuOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1BundleproductsSkuOptionsOptionId`: BundleDataOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsSkuOptionsOptionIdAPI.GetV1BundleproductsSkuOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BundleproductsSkuOptionsOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BundleDataOptionInterface**](BundleDataOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

