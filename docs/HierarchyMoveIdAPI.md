# \HierarchyMoveIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1HierarchyMoveId**](HierarchyMoveIdAPI.md#PutV1HierarchyMoveId) | **Put** /V1/hierarchy/move/{id} | hierarchy/move/{id}



## PutV1HierarchyMoveId

> ErrorResponse PutV1HierarchyMoveId(ctx, id).PutV1HierarchyMoveIdRequest(putV1HierarchyMoveIdRequest).Execute()

hierarchy/move/{id}



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
	putV1HierarchyMoveIdRequest := *openapiclient.NewPutV1HierarchyMoveIdRequest(int32(123)) // PutV1HierarchyMoveIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HierarchyMoveIdAPI.PutV1HierarchyMoveId(context.Background(), id).PutV1HierarchyMoveIdRequest(putV1HierarchyMoveIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HierarchyMoveIdAPI.PutV1HierarchyMoveId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1HierarchyMoveId`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `HierarchyMoveIdAPI.PutV1HierarchyMoveId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1HierarchyMoveIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1HierarchyMoveIdRequest** | [**PutV1HierarchyMoveIdRequest**](PutV1HierarchyMoveIdRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

