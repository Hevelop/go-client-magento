# \CategoriesCategoryIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CategoriesCategoryId**](CategoriesCategoryIdAPI.md#DeleteV1CategoriesCategoryId) | **Delete** /V1/categories/{categoryId} | categories/{categoryId}
[**GetV1CategoriesCategoryId**](CategoriesCategoryIdAPI.md#GetV1CategoriesCategoryId) | **Get** /V1/categories/{categoryId} | categories/{categoryId}



## DeleteV1CategoriesCategoryId

> bool DeleteV1CategoriesCategoryId(ctx, categoryId).Execute()

categories/{categoryId}



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
	resp, r, err := apiClient.CategoriesCategoryIdAPI.DeleteV1CategoriesCategoryId(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesCategoryIdAPI.DeleteV1CategoriesCategoryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CategoriesCategoryId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CategoriesCategoryIdAPI.DeleteV1CategoriesCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CategoriesCategoryIdRequest struct via the builder pattern


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


## GetV1CategoriesCategoryId

> CatalogDataCategoryInterface GetV1CategoriesCategoryId(ctx, categoryId).StoreId(storeId).Execute()

categories/{categoryId}



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
	storeId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesCategoryIdAPI.GetV1CategoriesCategoryId(context.Background(), categoryId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesCategoryIdAPI.GetV1CategoriesCategoryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CategoriesCategoryId`: CatalogDataCategoryInterface
	fmt.Fprintf(os.Stdout, "Response from `CategoriesCategoryIdAPI.GetV1CategoriesCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CategoriesCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeId** | **int32** |  | 

### Return type

[**CatalogDataCategoryInterface**](CatalogDataCategoryInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

