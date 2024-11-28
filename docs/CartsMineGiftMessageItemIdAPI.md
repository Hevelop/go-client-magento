# \CartsMineGiftMessageItemIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMineGiftmessageItemId**](CartsMineGiftMessageItemIdAPI.md#GetV1CartsMineGiftmessageItemId) | **Get** /V1/carts/mine/gift-message/{itemId} | carts/mine/gift-message/{itemId}
[**PostV1CartsMineGiftmessageItemId**](CartsMineGiftMessageItemIdAPI.md#PostV1CartsMineGiftmessageItemId) | **Post** /V1/carts/mine/gift-message/{itemId} | carts/mine/gift-message/{itemId}



## GetV1CartsMineGiftmessageItemId

> GiftMessageDataMessageInterface GetV1CartsMineGiftmessageItemId(ctx, itemId).Execute()

carts/mine/gift-message/{itemId}



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
	itemId := int32(56) // int32 | The item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineGiftMessageItemIdAPI.GetV1CartsMineGiftmessageItemId(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineGiftMessageItemIdAPI.GetV1CartsMineGiftmessageItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMineGiftmessageItemId`: GiftMessageDataMessageInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineGiftMessageItemIdAPI.GetV1CartsMineGiftmessageItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **int32** | The item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineGiftmessageItemIdRequest struct via the builder pattern


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


## PostV1CartsMineGiftmessageItemId

> bool PostV1CartsMineGiftmessageItemId(ctx, itemId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()

carts/mine/gift-message/{itemId}



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
	itemId := int32(56) // int32 | The item ID.
	postV1CartsMineGiftmessageRequest := *openapiclient.NewPostV1CartsMineGiftmessageRequest(*openapiclient.NewGiftMessageDataMessageInterface("Sender_example", "Recipient_example", "Message_example")) // PostV1CartsMineGiftmessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineGiftMessageItemIdAPI.PostV1CartsMineGiftmessageItemId(context.Background(), itemId).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineGiftMessageItemIdAPI.PostV1CartsMineGiftmessageItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineGiftmessageItemId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineGiftMessageItemIdAPI.PostV1CartsMineGiftmessageItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **int32** | The item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineGiftmessageItemIdRequest struct via the builder pattern


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

