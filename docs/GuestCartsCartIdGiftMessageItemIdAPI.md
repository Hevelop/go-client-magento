# \GuestCartsCartIdGiftMessageItemIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdGiftmessageItemId**](GuestCartsCartIdGiftMessageItemIdAPI.md#GetV1GuestcartsCartIdGiftmessageItemId) | **Get** /V1/guest-carts/{cartId}/gift-message/{itemId} | guest-carts/{cartId}/gift-message/{itemId}
[**PostV1GuestcartsCartIdGiftmessageItemId**](GuestCartsCartIdGiftMessageItemIdAPI.md#PostV1GuestcartsCartIdGiftmessageItemId) | **Post** /V1/guest-carts/{cartId}/gift-message/{itemId} | guest-carts/{cartId}/gift-message/{itemId}



## GetV1GuestcartsCartIdGiftmessageItemId

> GiftMessageDataMessageInterface GetV1GuestcartsCartIdGiftmessageItemId(ctx, cartId, itemId).Execute()

guest-carts/{cartId}/gift-message/{itemId}



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
	itemId := int32(56) // int32 | The item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdGiftMessageItemIdAPI.GetV1GuestcartsCartIdGiftmessageItemId(context.Background(), cartId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdGiftMessageItemIdAPI.GetV1GuestcartsCartIdGiftmessageItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdGiftmessageItemId`: GiftMessageDataMessageInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdGiftMessageItemIdAPI.GetV1GuestcartsCartIdGiftmessageItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The shopping cart ID. | 
**itemId** | **int32** | The item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdGiftmessageItemIdRequest struct via the builder pattern


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


## PostV1GuestcartsCartIdGiftmessageItemId

> bool PostV1GuestcartsCartIdGiftmessageItemId(ctx, cartId, itemId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()

guest-carts/{cartId}/gift-message/{itemId}



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
	itemId := int32(56) // int32 | The item ID.
	postV1CartsMineGiftmessageRequest := *openapiclient.NewPostV1CartsMineGiftmessageRequest(*openapiclient.NewGiftMessageDataMessageInterface("Sender_example", "Recipient_example", "Message_example")) // PostV1CartsMineGiftmessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdGiftMessageItemIdAPI.PostV1GuestcartsCartIdGiftmessageItemId(context.Background(), cartId, itemId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdGiftMessageItemIdAPI.PostV1GuestcartsCartIdGiftmessageItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdGiftmessageItemId`: bool
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdGiftMessageItemIdAPI.PostV1GuestcartsCartIdGiftmessageItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 
**itemId** | **int32** | The item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdGiftmessageItemIdRequest struct via the builder pattern


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

