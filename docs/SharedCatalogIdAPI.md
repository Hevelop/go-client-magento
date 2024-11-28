# \SharedCatalogIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1SharedCatalogId**](SharedCatalogIdAPI.md#PutV1SharedCatalogId) | **Put** /V1/sharedCatalog/{id} | sharedCatalog/{id}



## PutV1SharedCatalogId

> int32 PutV1SharedCatalogId(ctx, id).PostV1SharedCatalogRequest(postV1SharedCatalogRequest).Execute()

sharedCatalog/{id}



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
	postV1SharedCatalogRequest := *openapiclient.NewPostV1SharedCatalogRequest(*openapiclient.NewSharedCatalogDataSharedCatalogInterface("Name_example", "Description_example", int32(123), int32(123), "CreatedAt_example", int32(123), int32(123), int32(123))) // PostV1SharedCatalogRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCatalogIdAPI.PutV1SharedCatalogId(context.Background(), id).PostV1SharedCatalogRequest(postV1SharedCatalogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCatalogIdAPI.PutV1SharedCatalogId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1SharedCatalogId`: int32
	fmt.Fprintf(os.Stdout, "Response from `SharedCatalogIdAPI.PutV1SharedCatalogId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1SharedCatalogIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1SharedCatalogRequest** | [**PostV1SharedCatalogRequest**](PostV1SharedCatalogRequest.md) |  | 

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

