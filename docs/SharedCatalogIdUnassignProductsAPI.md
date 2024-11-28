# \SharedCatalogIdUnassignProductsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1SharedCatalogIdUnassignProducts**](SharedCatalogIdUnassignProductsAPI.md#PostV1SharedCatalogIdUnassignProducts) | **Post** /V1/sharedCatalog/{id}/unassignProducts | sharedCatalog/{id}/unassignProducts



## PostV1SharedCatalogIdUnassignProducts

> bool PostV1SharedCatalogIdUnassignProducts(ctx, id).PostV1SharedCatalogIdAssignProductsRequest(postV1SharedCatalogIdAssignProductsRequest).Execute()

sharedCatalog/{id}/unassignProducts



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
	postV1SharedCatalogIdAssignProductsRequest := *openapiclient.NewPostV1SharedCatalogIdAssignProductsRequest([]openapiclient.CatalogDataProductInterface{*openapiclient.NewCatalogDataProductInterface("Sku_example")}) // PostV1SharedCatalogIdAssignProductsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogIdUnassignProductsAPI.PostV1SharedCatalogIdUnassignProducts(context.Background(), id).PostV1SharedCatalogIdAssignProductsRequest(postV1SharedCatalogIdAssignProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogIdUnassignProductsAPI.PostV1SharedCatalogIdUnassignProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1SharedCatalogIdUnassignProducts`: bool
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogIdUnassignProductsAPI.PostV1SharedCatalogIdUnassignProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SharedCatalogIdUnassignProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1SharedCatalogIdAssignProductsRequest** | [**PostV1SharedCatalogIdAssignProductsRequest**](PostV1SharedCatalogIdAssignProductsRequest.md) |  | 

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

