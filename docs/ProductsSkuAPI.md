# \ProductsSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsSku**](ProductsSkuAPI.md#DeleteV1ProductsSku) | **Delete** /V1/products/{sku} | products/{sku}
[**GetV1ProductsSku**](ProductsSkuAPI.md#GetV1ProductsSku) | **Get** /V1/products/{sku} | products/{sku}
[**PutV1ProductsSku**](ProductsSkuAPI.md#PutV1ProductsSku) | **Put** /V1/products/{sku} | products/{sku}



## DeleteV1ProductsSku

> bool DeleteV1ProductsSku(ctx, sku).Execute()

products/{sku}



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuAPI.DeleteV1ProductsSku(context.Background(), sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuAPI.DeleteV1ProductsSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsSku`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuAPI.DeleteV1ProductsSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsSkuRequest struct via the builder pattern


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


## GetV1ProductsSku

> CatalogDataProductInterface GetV1ProductsSku(ctx, sku).EditMode(editMode).StoreId(storeId).ForceReload(forceReload).Execute()

products/{sku}



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
	editMode := true // bool |  (optional)
	storeId := int32(56) // int32 |  (optional)
	forceReload := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuAPI.GetV1ProductsSku(context.Background(), sku).EditMode(editMode).StoreId(storeId).ForceReload(forceReload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuAPI.GetV1ProductsSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSku`: CatalogDataProductInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuAPI.GetV1ProductsSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editMode** | **bool** |  | 
 **storeId** | **int32** |  | 
 **forceReload** | **bool** |  | 

### Return type

[**CatalogDataProductInterface**](CatalogDataProductInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1ProductsSku

> CatalogDataProductInterface PutV1ProductsSku(ctx, sku).PostV1ProductsRequest(postV1ProductsRequest).Execute()

products/{sku}



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
	postV1ProductsRequest := *openapiclient.NewPostV1ProductsRequest(*openapiclient.NewCatalogDataProductInterface("Sku_example")) // PostV1ProductsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuAPI.PutV1ProductsSku(context.Background(), sku).PostV1ProductsRequest(postV1ProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuAPI.PutV1ProductsSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsSku`: CatalogDataProductInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuAPI.PutV1ProductsSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ProductsRequest** | [**PostV1ProductsRequest**](PostV1ProductsRequest.md) |  | 

### Return type

[**CatalogDataProductInterface**](CatalogDataProductInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

