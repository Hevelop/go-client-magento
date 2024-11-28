# \CartsMineGiftMessageAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMineGiftmessage**](CartsMineGiftMessageAPI.md#GetV1CartsMineGiftmessage) | **Get** /V1/carts/mine/gift-message | carts/mine/gift-message
[**PostV1CartsMineGiftmessage**](CartsMineGiftMessageAPI.md#PostV1CartsMineGiftmessage) | **Post** /V1/carts/mine/gift-message | carts/mine/gift-message



## GetV1CartsMineGiftmessage

> GiftMessageDataMessageInterface GetV1CartsMineGiftmessage(ctx).Execute()

carts/mine/gift-message



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineGiftMessageAPI.GetV1CartsMineGiftmessage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineGiftMessageAPI.GetV1CartsMineGiftmessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMineGiftmessage`: GiftMessageDataMessageInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineGiftMessageAPI.GetV1CartsMineGiftmessage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineGiftmessageRequest struct via the builder pattern


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


## PostV1CartsMineGiftmessage

> bool PostV1CartsMineGiftmessage(ctx).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()

carts/mine/gift-message



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
	postV1CartsMineGiftmessageRequest := *openapiclient.NewPostV1CartsMineGiftmessageRequest(*openapiclient.NewGiftMessageDataMessageInterface("Sender_example", "Recipient_example", "Message_example")) // PostV1CartsMineGiftmessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineGiftMessageAPI.PostV1CartsMineGiftmessage(context.Background()).PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineGiftMessageAPI.PostV1CartsMineGiftmessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineGiftmessage`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineGiftMessageAPI.PostV1CartsMineGiftmessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineGiftmessageRequest struct via the builder pattern


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

