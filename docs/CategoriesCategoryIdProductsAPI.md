# \CategoriesCategoryIdProductsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CategoriesCategoryIdProducts**](CategoriesCategoryIdProductsAPI.md#GetV1CategoriesCategoryIdProducts) | **Get** /V1/categories/{categoryId}/products | categories/{categoryId}/products
[**PostV1CategoriesCategoryIdProducts**](CategoriesCategoryIdProductsAPI.md#PostV1CategoriesCategoryIdProducts) | **Post** /V1/categories/{categoryId}/products | categories/{categoryId}/products
[**PutV1CategoriesCategoryIdProducts**](CategoriesCategoryIdProductsAPI.md#PutV1CategoriesCategoryIdProducts) | **Put** /V1/categories/{categoryId}/products | categories/{categoryId}/products



## GetV1CategoriesCategoryIdProducts

> []CatalogDataCategoryProductLinkInterface GetV1CategoriesCategoryIdProducts(ctx, categoryId).Execute()

categories/{categoryId}/products



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
	categoryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesCategoryIdProductsAPI.GetV1CategoriesCategoryIdProducts(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesCategoryIdProductsAPI.GetV1CategoriesCategoryIdProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CategoriesCategoryIdProducts`: []CatalogDataCategoryProductLinkInterface
	fmt.Fprintf(os.Stdout, "Response from `CategoriesCategoryIdProductsAPI.GetV1CategoriesCategoryIdProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CategoriesCategoryIdProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogDataCategoryProductLinkInterface**](CatalogDataCategoryProductLinkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1CategoriesCategoryIdProducts

> bool PostV1CategoriesCategoryIdProducts(ctx, categoryId).PutV1CategoriesCategoryIdProductsRequest(putV1CategoriesCategoryIdProductsRequest).Execute()

categories/{categoryId}/products



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
	categoryId := "categoryId_example" // string | 
	putV1CategoriesCategoryIdProductsRequest := *openapiclient.NewPutV1CategoriesCategoryIdProductsRequest(*openapiclient.NewCatalogDataCategoryProductLinkInterface("CategoryId_example")) // PutV1CategoriesCategoryIdProductsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesCategoryIdProductsAPI.PostV1CategoriesCategoryIdProducts(context.Background(), categoryId).PutV1CategoriesCategoryIdProductsRequest(putV1CategoriesCategoryIdProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesCategoryIdProductsAPI.PostV1CategoriesCategoryIdProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CategoriesCategoryIdProducts`: bool
	fmt.Fprintf(os.Stdout, "Response from `CategoriesCategoryIdProductsAPI.PostV1CategoriesCategoryIdProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CategoriesCategoryIdProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CategoriesCategoryIdProductsRequest** | [**PutV1CategoriesCategoryIdProductsRequest**](PutV1CategoriesCategoryIdProductsRequest.md) |  | 

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


## PutV1CategoriesCategoryIdProducts

> bool PutV1CategoriesCategoryIdProducts(ctx, categoryId).PutV1CategoriesCategoryIdProductsRequest(putV1CategoriesCategoryIdProductsRequest).Execute()

categories/{categoryId}/products



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
	categoryId := "categoryId_example" // string | 
	putV1CategoriesCategoryIdProductsRequest := *openapiclient.NewPutV1CategoriesCategoryIdProductsRequest(*openapiclient.NewCatalogDataCategoryProductLinkInterface("CategoryId_example")) // PutV1CategoriesCategoryIdProductsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesCategoryIdProductsAPI.PutV1CategoriesCategoryIdProducts(context.Background(), categoryId).PutV1CategoriesCategoryIdProductsRequest(putV1CategoriesCategoryIdProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesCategoryIdProductsAPI.PutV1CategoriesCategoryIdProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CategoriesCategoryIdProducts`: bool
	fmt.Fprintf(os.Stdout, "Response from `CategoriesCategoryIdProductsAPI.PutV1CategoriesCategoryIdProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CategoriesCategoryIdProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CategoriesCategoryIdProductsRequest** | [**PutV1CategoriesCategoryIdProductsRequest**](PutV1CategoriesCategoryIdProductsRequest.md) |  | 

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

