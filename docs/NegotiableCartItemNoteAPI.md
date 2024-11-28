# \NegotiableCartItemNoteAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1Negotiablecartitemnote**](NegotiableCartItemNoteAPI.md#PostV1Negotiablecartitemnote) | **Post** /V1/negotiable-cart-item-note | negotiable-cart-item-note



## PostV1Negotiablecartitemnote

> int32 PostV1Negotiablecartitemnote(ctx).PostV1NegotiablecartitemnoteRequest(postV1NegotiablecartitemnoteRequest).Execute()

negotiable-cart-item-note



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
	postV1NegotiablecartitemnoteRequest := *openapiclient.NewPostV1NegotiablecartitemnoteRequest(*openapiclient.NewNegotiableQuoteDataItemNoteInterface(int32(123), int32(123), int32(123), "Note_example")) // PostV1NegotiablecartitemnoteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartItemNoteAPI.PostV1Negotiablecartitemnote(context.Background()).PostV1NegotiablecartitemnoteRequest(postV1NegotiablecartitemnoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartItemNoteAPI.PostV1Negotiablecartitemnote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Negotiablecartitemnote`: int32
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartItemNoteAPI.PostV1Negotiablecartitemnote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiablecartitemnoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1NegotiablecartitemnoteRequest** | [**PostV1NegotiablecartitemnoteRequest**](PostV1NegotiablecartitemnoteRequest.md) |  | 

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

