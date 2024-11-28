# \GiftWrappingsIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1GiftwrappingsId**](GiftWrappingsIdAPI.md#DeleteV1GiftwrappingsId) | **Delete** /V1/gift-wrappings/{id} | gift-wrappings/{id}
[**GetV1GiftwrappingsId**](GiftWrappingsIdAPI.md#GetV1GiftwrappingsId) | **Get** /V1/gift-wrappings/{id} | gift-wrappings/{id}



## DeleteV1GiftwrappingsId

> bool DeleteV1GiftwrappingsId(ctx, id).Execute()

gift-wrappings/{id}



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
	resp, r, err := apiClient.GiftWrappingsIdAPI.DeleteV1GiftwrappingsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftWrappingsIdAPI.DeleteV1GiftwrappingsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1GiftwrappingsId`: bool
	fmt.Fprintf(os.Stdout, "Response from `GiftWrappingsIdAPI.DeleteV1GiftwrappingsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1GiftwrappingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1GiftwrappingsId

> GiftWrappingDataWrappingInterface GetV1GiftwrappingsId(ctx, id).StoreId(storeId).Execute()

gift-wrappings/{id}



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
	storeId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftWrappingsIdAPI.GetV1GiftwrappingsId(context.Background(), id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftWrappingsIdAPI.GetV1GiftwrappingsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GiftwrappingsId`: GiftWrappingDataWrappingInterface
	fmt.Fprintf(os.Stdout, "Response from `GiftWrappingsIdAPI.GetV1GiftwrappingsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GiftwrappingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeId** | **int32** |  | 

### Return type

[**GiftWrappingDataWrappingInterface**](GiftWrappingDataWrappingInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

