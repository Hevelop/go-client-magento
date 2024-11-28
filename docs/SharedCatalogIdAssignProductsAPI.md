# \SharedCatalogIdAssignProductsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1SharedCatalogIdAssignProducts**](SharedCatalogIdAssignProductsAPI.md#PostV1SharedCatalogIdAssignProducts) | **Post** /V1/sharedCatalog/{id}/assignProducts | sharedCatalog/{id}/assignProducts



## PostV1SharedCatalogIdAssignProducts

> bool PostV1SharedCatalogIdAssignProducts(ctx, id).PostV1SharedCatalogIdAssignProductsRequest(postV1SharedCatalogIdAssignProductsRequest).Execute()

sharedCatalog/{id}/assignProducts



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
	resp, r, err := apiClient.SharedCatalogIdAssignProductsAPI.PostV1SharedCatalogIdAssignProducts(context.Background(), id).PostV1SharedCatalogIdAssignProductsRequest(postV1SharedCatalogIdAssignProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogIdAssignProductsAPI.PostV1SharedCatalogIdAssignProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1SharedCatalogIdAssignProducts`: bool
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogIdAssignProductsAPI.PostV1SharedCatalogIdAssignProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SharedCatalogIdAssignProductsRequest struct via the builder pattern


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

