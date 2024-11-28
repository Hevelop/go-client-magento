# \SharedCatalogIdUnassignCategoriesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1SharedCatalogIdUnassignCategories**](SharedCatalogIdUnassignCategoriesAPI.md#PostV1SharedCatalogIdUnassignCategories) | **Post** /V1/sharedCatalog/{id}/unassignCategories | sharedCatalog/{id}/unassignCategories



## PostV1SharedCatalogIdUnassignCategories

> bool PostV1SharedCatalogIdUnassignCategories(ctx, id).PostV1SharedCatalogIdAssignCategoriesRequest(postV1SharedCatalogIdAssignCategoriesRequest).Execute()

sharedCatalog/{id}/unassignCategories



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
	postV1SharedCatalogIdAssignCategoriesRequest := *openapiclient.NewPostV1SharedCatalogIdAssignCategoriesRequest([]openapiclient.CatalogDataCategoryInterface{*openapiclient.NewCatalogDataCategoryInterface()}) // PostV1SharedCatalogIdAssignCategoriesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogIdUnassignCategoriesAPI.PostV1SharedCatalogIdUnassignCategories(context.Background(), id).PostV1SharedCatalogIdAssignCategoriesRequest(postV1SharedCatalogIdAssignCategoriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogIdUnassignCategoriesAPI.PostV1SharedCatalogIdUnassignCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1SharedCatalogIdUnassignCategories`: bool
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogIdUnassignCategoriesAPI.PostV1SharedCatalogIdUnassignCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SharedCatalogIdUnassignCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1SharedCatalogIdAssignCategoriesRequest** | [**PostV1SharedCatalogIdAssignCategoriesRequest**](PostV1SharedCatalogIdAssignCategoriesRequest.md) |  | 

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

