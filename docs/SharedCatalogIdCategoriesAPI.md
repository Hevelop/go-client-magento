# \SharedCatalogIdCategoriesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1SharedCatalogIdCategories**](SharedCatalogIdCategoriesAPI.md#GetV1SharedCatalogIdCategories) | **Get** /V1/sharedCatalog/{id}/categories | sharedCatalog/{id}/categories



## GetV1SharedCatalogIdCategories

> []int32 GetV1SharedCatalogIdCategories(ctx, id).Execute()

sharedCatalog/{id}/categories



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogIdCategoriesAPI.GetV1SharedCatalogIdCategories(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogIdCategoriesAPI.GetV1SharedCatalogIdCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1SharedCatalogIdCategories`: []int32
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogIdCategoriesAPI.GetV1SharedCatalogIdCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1SharedCatalogIdCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

