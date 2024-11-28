# \CategoriesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1Categories**](CategoriesAPI.md#GetV1Categories) | **Get** /V1/categories | categories
[**PostV1Categories**](CategoriesAPI.md#PostV1Categories) | **Post** /V1/categories | categories



## GetV1Categories

> CatalogDataCategoryTreeInterface GetV1Categories(ctx).RootCategoryId(rootCategoryId).Depth(depth).Execute()

categories



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
	rootCategoryId := int32(56) // int32 |  (optional)
	depth := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.GetV1Categories(context.Background()).RootCategoryId(rootCategoryId).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.GetV1Categories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Categories`: CatalogDataCategoryTreeInterface
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.GetV1Categories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootCategoryId** | **int32** |  | 
 **depth** | **int32** |  | 

### Return type

[**CatalogDataCategoryTreeInterface**](CatalogDataCategoryTreeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1Categories

> CatalogDataCategoryInterface PostV1Categories(ctx).PostV1CategoriesRequest(postV1CategoriesRequest).Execute()

categories



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
	postV1CategoriesRequest := *openapiclient.NewPostV1CategoriesRequest(*openapiclient.NewCatalogDataCategoryInterface()) // PostV1CategoriesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.PostV1Categories(context.Background()).PostV1CategoriesRequest(postV1CategoriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.PostV1Categories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Categories`: CatalogDataCategoryInterface
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.PostV1Categories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CategoriesRequest** | [**PostV1CategoriesRequest**](PostV1CategoriesRequest.md) |  | 

### Return type

[**CatalogDataCategoryInterface**](CatalogDataCategoryInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

