# \GuestCartsCartIdGiftMessageAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdGiftmessage**](GuestCartsCartIdGiftMessageAPI.md#GetV1GuestcartsCartIdGiftmessage) | **Get** /V1/guest-carts/{cartId}/gift-message | guest-carts/{cartId}/gift-message
[**PostV1GuestcartsCartIdGiftmessage**](GuestCartsCartIdGiftMessageAPI.md#PostV1GuestcartsCartIdGiftmessage) | **Post** /V1/guest-carts/{cartId}/gift-message | guest-carts/{cartId}/gift-message



## GetV1GuestcartsCartIdGiftmessage

> GiftMessageDataMessageInterface GetV1GuestcartsCartIdGiftmessage(ctx, cartId).Execute()

guest-carts/{cartId}/gift-message



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
	cartId := "cartId_example" // string | The shopping cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdGiftMessageAPI.GetV1GuestcartsCartIdGiftmessage(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdGiftMessageAPI.GetV1GuestcartsCartIdGiftmessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdGiftmessage`: GiftMessageDataMessageInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdGiftMessageAPI.GetV1GuestcartsCartIdGiftmessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The shopping cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdGiftmessageRequest struct via the builder pattern


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


## PostV1GuestcartsCartIdGiftmessage

> bool PostV1GuestcartsCartIdGiftmessage(ctx, cartId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()

guest-carts/{cartId}/gift-message



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
	cartId := "cartId_example" // string | The cart ID.
	postV1CartsMineGiftmessageRequest := *openapiclient.NewPostV1CartsMineGiftmessageRequest(*openapiclient.NewGiftMessageDataMessageInterface("Sender_example", "Recipient_example", "Message_example")) // PostV1CartsMineGiftmessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdGiftMessageAPI.PostV1GuestcartsCartIdGiftmessage(context.Background(), cartId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdGiftMessageAPI.PostV1GuestcartsCartIdGiftmessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdGiftmessage`: bool
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdGiftMessageAPI.PostV1GuestcartsCartIdGiftmessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdGiftmessageRequest struct via the builder pattern


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

