# \GiftWrappingsWrappingIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1GiftwrappingsWrappingId**](GiftWrappingsWrappingIdAPI.md#PutV1GiftwrappingsWrappingId) | **Put** /V1/gift-wrappings/{wrappingId} | gift-wrappings/{wrappingId}



## PutV1GiftwrappingsWrappingId

> GiftWrappingDataWrappingInterface PutV1GiftwrappingsWrappingId(ctx, wrappingId).PostV1GiftwrappingsRequest(postV1GiftwrappingsRequest).Execute()

gift-wrappings/{wrappingId}



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
	wrappingId := "wrappingId_example" // string | 
	postV1GiftwrappingsRequest := *openapiclient.NewPostV1GiftwrappingsRequest(*openapiclient.NewGiftWrappingDataWrappingInterface("Design_example", int32(123), float32(123))) // PostV1GiftwrappingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftWrappingsWrappingIdAPI.PutV1GiftwrappingsWrappingId(context.Background(), wrappingId).PostV1GiftwrappingsRequest(postV1GiftwrappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftWrappingsWrappingIdAPI.PutV1GiftwrappingsWrappingId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GiftwrappingsWrappingId`: GiftWrappingDataWrappingInterface
	fmt.Fprintf(os.Stdout, "Response from `GiftWrappingsWrappingIdAPI.PutV1GiftwrappingsWrappingId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wrappingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GiftwrappingsWrappingIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1GiftwrappingsRequest** | [**PostV1GiftwrappingsRequest**](PostV1GiftwrappingsRequest.md) |  | 

### Return type

[**GiftWrappingDataWrappingInterface**](GiftWrappingDataWrappingInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

