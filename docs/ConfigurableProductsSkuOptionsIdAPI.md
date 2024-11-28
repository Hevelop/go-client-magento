# \ConfigurableProductsSkuOptionsIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ConfigurableproductsSkuOptionsId**](ConfigurableProductsSkuOptionsIdAPI.md#DeleteV1ConfigurableproductsSkuOptionsId) | **Delete** /V1/configurable-products/{sku}/options/{id} | configurable-products/{sku}/options/{id}
[**GetV1ConfigurableproductsSkuOptionsId**](ConfigurableProductsSkuOptionsIdAPI.md#GetV1ConfigurableproductsSkuOptionsId) | **Get** /V1/configurable-products/{sku}/options/{id} | configurable-products/{sku}/options/{id}
[**PutV1ConfigurableproductsSkuOptionsId**](ConfigurableProductsSkuOptionsIdAPI.md#PutV1ConfigurableproductsSkuOptionsId) | **Put** /V1/configurable-products/{sku}/options/{id} | configurable-products/{sku}/options/{id}



## DeleteV1ConfigurableproductsSkuOptionsId

> bool DeleteV1ConfigurableproductsSkuOptionsId(ctx, sku, id).Execute()

configurable-products/{sku}/options/{id}



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurableProductsSkuOptionsIdAPI.DeleteV1ConfigurableproductsSkuOptionsId(context.Background(), sku, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuOptionsIdAPI.DeleteV1ConfigurableproductsSkuOptionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ConfigurableproductsSkuOptionsId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuOptionsIdAPI.DeleteV1ConfigurableproductsSkuOptionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ConfigurableproductsSkuOptionsIdRequest struct via the builder pattern


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


## GetV1ConfigurableproductsSkuOptionsId

> ConfigurableProductDataOptionInterface GetV1ConfigurableproductsSkuOptionsId(ctx, sku, id).Execute()

configurable-products/{sku}/options/{id}



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurableProductsSkuOptionsIdAPI.GetV1ConfigurableproductsSkuOptionsId(context.Background(), sku, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuOptionsIdAPI.GetV1ConfigurableproductsSkuOptionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ConfigurableproductsSkuOptionsId`: ConfigurableProductDataOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuOptionsIdAPI.GetV1ConfigurableproductsSkuOptionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ConfigurableproductsSkuOptionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfigurableProductDataOptionInterface**](ConfigurableProductDataOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1ConfigurableproductsSkuOptionsId

> int32 PutV1ConfigurableproductsSkuOptionsId(ctx, sku, id).PostV1ConfigurableproductsSkuOptionsRequest(postV1ConfigurableproductsSkuOptionsRequest).Execute()

configurable-products/{sku}/options/{id}



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
	postV1ConfigurableproductsSkuOptionsRequest := *openapiclient.NewPostV1ConfigurableproductsSkuOptionsRequest(*openapiclient.NewConfigurableProductDataOptionInterface()) // PostV1ConfigurableproductsSkuOptionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurableProductsSkuOptionsIdAPI.PutV1ConfigurableproductsSkuOptionsId(context.Background(), sku, id).PostV1ConfigurableproductsSkuOptionsRequest(postV1ConfigurableproductsSkuOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuOptionsIdAPI.PutV1ConfigurableproductsSkuOptionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ConfigurableproductsSkuOptionsId`: int32
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuOptionsIdAPI.PutV1ConfigurableproductsSkuOptionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ConfigurableproductsSkuOptionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postV1ConfigurableproductsSkuOptionsRequest** | [**PostV1ConfigurableproductsSkuOptionsRequest**](PostV1ConfigurableproductsSkuOptionsRequest.md) |  | 

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

