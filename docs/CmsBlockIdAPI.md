# \CmsBlockIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CmsBlockId**](CmsBlockIdAPI.md#PutV1CmsBlockId) | **Put** /V1/cmsBlock/{id} | cmsBlock/{id}



## PutV1CmsBlockId

> CmsDataBlockInterface PutV1CmsBlockId(ctx, id).PostV1CmsBlockRequest(postV1CmsBlockRequest).Execute()

cmsBlock/{id}



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
	postV1CmsBlockRequest := *openapiclient.NewPostV1CmsBlockRequest(*openapiclient.NewCmsDataBlockInterface("Identifier_example")) // PostV1CmsBlockRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CmsBlockIdAPI.PutV1CmsBlockId(context.Background(), id).PostV1CmsBlockRequest(postV1CmsBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CmsBlockIdAPI.PutV1CmsBlockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CmsBlockId`: CmsDataBlockInterface
	fmt.Fprintf(os.Stdout, "Response from `CmsBlockIdAPI.PutV1CmsBlockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CmsBlockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CmsBlockRequest** | [**PostV1CmsBlockRequest**](PostV1CmsBlockRequest.md) |  | 

### Return type

[**CmsDataBlockInterface**](CmsDataBlockInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

