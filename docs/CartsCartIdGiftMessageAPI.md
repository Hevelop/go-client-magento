# \CartsCartIdGiftMessageAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsCartIdGiftmessage**](CartsCartIdGiftMessageAPI.md#GetV1CartsCartIdGiftmessage) | **Get** /V1/carts/{cartId}/gift-message | carts/{cartId}/gift-message
[**PostV1CartsCartIdGiftmessage**](CartsCartIdGiftMessageAPI.md#PostV1CartsCartIdGiftmessage) | **Post** /V1/carts/{cartId}/gift-message | carts/{cartId}/gift-message



## GetV1CartsCartIdGiftmessage

> GiftMessageDataMessageInterface GetV1CartsCartIdGiftmessage(ctx, cartId).Execute()

carts/{cartId}/gift-message



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
	cartId := int32(56) // int32 | The shopping cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdGiftMessageAPI.GetV1CartsCartIdGiftmessage(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdGiftMessageAPI.GetV1CartsCartIdGiftmessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartIdGiftmessage`: GiftMessageDataMessageInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdGiftMessageAPI.GetV1CartsCartIdGiftmessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The shopping cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdGiftmessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftMessageDataMessageInterface**](GiftMessageDataMessageInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1CartsCartIdGiftmessage

> bool PostV1CartsCartIdGiftmessage(ctx, cartId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()

carts/{cartId}/gift-message



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
	cartId := int32(56) // int32 | The cart ID.
	postV1CartsMineGiftmessageRequest := *openapiclient.NewPostV1CartsMineGiftmessageRequest(*openapiclient.NewGiftMessageDataMessageInterface("Sender_example", "Recipient_example", "Message_example")) // PostV1CartsMineGiftmessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdGiftMessageAPI.PostV1CartsCartIdGiftmessage(context.Background(), cartId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdGiftMessageAPI.PostV1CartsCartIdGiftmessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsCartIdGiftmessage`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdGiftMessageAPI.PostV1CartsCartIdGiftmessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsCartIdGiftmessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMineGiftmessageRequest** | [**PostV1CartsMineGiftmessageRequest**](PostV1CartsMineGiftmessageRequest.md) |  | 

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

