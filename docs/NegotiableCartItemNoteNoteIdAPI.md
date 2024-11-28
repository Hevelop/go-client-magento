# \NegotiableCartItemNoteNoteIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1NegotiablecartitemnoteNoteId**](NegotiableCartItemNoteNoteIdAPI.md#DeleteV1NegotiablecartitemnoteNoteId) | **Delete** /V1/negotiable-cart-item-note/{noteId} | negotiable-cart-item-note/{noteId}
[**GetV1NegotiablecartitemnoteNoteId**](NegotiableCartItemNoteNoteIdAPI.md#GetV1NegotiablecartitemnoteNoteId) | **Get** /V1/negotiable-cart-item-note/{noteId} | negotiable-cart-item-note/{noteId}



## DeleteV1NegotiablecartitemnoteNoteId

> ErrorResponse DeleteV1NegotiablecartitemnoteNoteId(ctx, noteId).Execute()

negotiable-cart-item-note/{noteId}



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
	noteId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartItemNoteNoteIdAPI.DeleteV1NegotiablecartitemnoteNoteId(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartItemNoteNoteIdAPI.DeleteV1NegotiablecartitemnoteNoteId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1NegotiablecartitemnoteNoteId`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartItemNoteNoteIdAPI.DeleteV1NegotiablecartitemnoteNoteId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1NegotiablecartitemnoteNoteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1NegotiablecartitemnoteNoteId

> NegotiableQuoteDataItemNoteInterface GetV1NegotiablecartitemnoteNoteId(ctx, noteId).Execute()

negotiable-cart-item-note/{noteId}



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
	noteId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartItemNoteNoteIdAPI.GetV1NegotiablecartitemnoteNoteId(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartItemNoteNoteIdAPI.GetV1NegotiablecartitemnoteNoteId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1NegotiablecartitemnoteNoteId`: NegotiableQuoteDataItemNoteInterface
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartItemNoteNoteIdAPI.GetV1NegotiablecartitemnoteNoteId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1NegotiablecartitemnoteNoteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NegotiableQuoteDataItemNoteInterface**](NegotiableQuoteDataItemNoteInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

