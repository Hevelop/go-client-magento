# \ProductsSkuOptionsOptionIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsSkuOptionsOptionId**](ProductsSkuOptionsOptionIdAPI.md#DeleteV1ProductsSkuOptionsOptionId) | **Delete** /V1/products/{sku}/options/{optionId} | products/{sku}/options/{optionId}
[**GetV1ProductsSkuOptionsOptionId**](ProductsSkuOptionsOptionIdAPI.md#GetV1ProductsSkuOptionsOptionId) | **Get** /V1/products/{sku}/options/{optionId} | products/{sku}/options/{optionId}



## DeleteV1ProductsSkuOptionsOptionId

> bool DeleteV1ProductsSkuOptionsOptionId(ctx, sku, optionId).Execute()

products/{sku}/options/{optionId}



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
	resp, r, err := apiClient.ProductsSkuOptionsOptionIdAPI.DeleteV1ProductsSkuOptionsOptionId(context.Background(), sku, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuOptionsOptionIdAPI.DeleteV1ProductsSkuOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsSkuOptionsOptionId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuOptionsOptionIdAPI.DeleteV1ProductsSkuOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsSkuOptionsOptionIdRequest struct via the builder pattern


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


## GetV1ProductsSkuOptionsOptionId

> CatalogDataProductCustomOptionInterface GetV1ProductsSkuOptionsOptionId(ctx, sku, optionId).Execute()

products/{sku}/options/{optionId}



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
	resp, r, err := apiClient.ProductsSkuOptionsOptionIdAPI.GetV1ProductsSkuOptionsOptionId(context.Background(), sku, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuOptionsOptionIdAPI.GetV1ProductsSkuOptionsOptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSkuOptionsOptionId`: CatalogDataProductCustomOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuOptionsOptionIdAPI.GetV1ProductsSkuOptionsOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuOptionsOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CatalogDataProductCustomOptionInterface**](CatalogDataProductCustomOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

