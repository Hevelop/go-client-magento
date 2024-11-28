# \CartsCartIdGiftMessageItemIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsCartIdGiftmessageItemId**](CartsCartIdGiftMessageItemIdAPI.md#GetV1CartsCartIdGiftmessageItemId) | **Get** /V1/carts/{cartId}/gift-message/{itemId} | carts/{cartId}/gift-message/{itemId}
[**PostV1CartsCartIdGiftmessageItemId**](CartsCartIdGiftMessageItemIdAPI.md#PostV1CartsCartIdGiftmessageItemId) | **Post** /V1/carts/{cartId}/gift-message/{itemId} | carts/{cartId}/gift-message/{itemId}



## GetV1CartsCartIdGiftmessageItemId

> GiftMessageDataMessageInterface GetV1CartsCartIdGiftmessageItemId(ctx, cartId, itemId).Execute()

carts/{cartId}/gift-message/{itemId}



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
	itemId := int32(56) // int32 | The item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdGiftMessageItemIdAPI.GetV1CartsCartIdGiftmessageItemId(context.Background(), cartId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdGiftMessageItemIdAPI.GetV1CartsCartIdGiftmessageItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartIdGiftmessageItemId`: GiftMessageDataMessageInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdGiftMessageItemIdAPI.GetV1CartsCartIdGiftmessageItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The shopping cart ID. | 
**itemId** | **int32** | The item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdGiftmessageItemIdRequest struct via the builder pattern


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


## PostV1CartsCartIdGiftmessageItemId

> bool PostV1CartsCartIdGiftmessageItemId(ctx, cartId, itemId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()

carts/{cartId}/gift-message/{itemId}



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
	itemId := int32(56) // int32 | The item ID.
	postV1CartsMineGiftmessageRequest := *openapiclient.NewPostV1CartsMineGiftmessageRequest(*openapiclient.NewGiftMessageDataMessageInterface("Sender_example", "Recipient_example", "Message_example")) // PostV1CartsMineGiftmessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdGiftMessageItemIdAPI.PostV1CartsCartIdGiftmessageItemId(context.Background(), cartId, itemId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdGiftMessageItemIdAPI.PostV1CartsCartIdGiftmessageItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsCartIdGiftmessageItemId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdGiftMessageItemIdAPI.PostV1CartsCartIdGiftmessageItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 
**itemId** | **int32** | The item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsCartIdGiftmessageItemIdRequest struct via the builder pattern


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

