# \CategoriesIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CategoriesId**](CategoriesIdAPI.md#PutV1CategoriesId) | **Put** /V1/categories/{id} | categories/{id}



## PutV1CategoriesId

> CatalogDataCategoryInterface PutV1CategoriesId(ctx, id).PostV1CategoriesRequest(postV1CategoriesRequest).Execute()

categories/{id}



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
	id := "id_example" // string | 
	postV1CategoriesRequest := *openapiclient.NewPostV1CategoriesRequest(*openapiclient.NewCatalogDataCategoryInterface()) // PostV1CategoriesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesIdAPI.PutV1CategoriesId(context.Background(), id).PostV1CategoriesRequest(postV1CategoriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesIdAPI.PutV1CategoriesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CategoriesId`: CatalogDataCategoryInterface
	fmt.Fprintf(os.Stdout, "Response from `CategoriesIdAPI.PutV1CategoriesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CategoriesIdRequest struct via the builder pattern


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

