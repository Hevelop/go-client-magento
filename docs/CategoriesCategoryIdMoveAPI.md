# \CategoriesCategoryIdMoveAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CategoriesCategoryIdMove**](CategoriesCategoryIdMoveAPI.md#PutV1CategoriesCategoryIdMove) | **Put** /V1/categories/{categoryId}/move | categories/{categoryId}/move



## PutV1CategoriesCategoryIdMove

> bool PutV1CategoriesCategoryIdMove(ctx, categoryId).PutV1CategoriesCategoryIdMoveRequest(putV1CategoriesCategoryIdMoveRequest).Execute()

categories/{categoryId}/move



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
	putV1CategoriesCategoryIdMoveRequest := *openapiclient.NewPutV1CategoriesCategoryIdMoveRequest(int32(123)) // PutV1CategoriesCategoryIdMoveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesCategoryIdMoveAPI.PutV1CategoriesCategoryIdMove(context.Background(), categoryId).PutV1CategoriesCategoryIdMoveRequest(putV1CategoriesCategoryIdMoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesCategoryIdMoveAPI.PutV1CategoriesCategoryIdMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CategoriesCategoryIdMove`: bool
	fmt.Fprintf(os.Stdout, "Response from `CategoriesCategoryIdMoveAPI.PutV1CategoriesCategoryIdMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CategoriesCategoryIdMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CategoriesCategoryIdMoveRequest** | [**PutV1CategoriesCategoryIdMoveRequest**](PutV1CategoriesCategoryIdMoveRequest.md) |  | 

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

