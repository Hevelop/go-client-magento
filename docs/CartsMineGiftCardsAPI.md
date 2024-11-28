# \CartsMineGiftCardsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMineGiftCards**](CartsMineGiftCardsAPI.md#PostV1CartsMineGiftCards) | **Post** /V1/carts/mine/giftCards | carts/mine/giftCards



## PostV1CartsMineGiftCards

> bool PostV1CartsMineGiftCards(ctx).PostV1CartsGuestcartsCartIdGiftCardsRequest(postV1CartsGuestcartsCartIdGiftCardsRequest).Execute()

carts/mine/giftCards



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
	postV1CartsGuestcartsCartIdGiftCardsRequest := *openapiclient.NewPostV1CartsGuestcartsCartIdGiftCardsRequest(*openapiclient.NewGiftCardAccountDataGiftCardAccountInterface(float32(123), float32(123), float32(123), float32(123))) // PostV1CartsGuestcartsCartIdGiftCardsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineGiftCardsAPI.PostV1CartsMineGiftCards(context.Background()).PostV1CartsGuestcartsCartIdGiftCardsRequest(postV1CartsGuestcartsCartIdGiftCardsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineGiftCardsAPI.PostV1CartsMineGiftCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineGiftCards`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineGiftCardsAPI.PostV1CartsMineGiftCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineGiftCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CartsGuestcartsCartIdGiftCardsRequest** | [**PostV1CartsGuestcartsCartIdGiftCardsRequest**](PostV1CartsGuestcartsCartIdGiftCardsRequest.md) |  | 

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

