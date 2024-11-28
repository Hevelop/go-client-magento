# \CartsGuestCartsCartIdGiftCardsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsGuestcartsCartIdGiftCards**](CartsGuestCartsCartIdGiftCardsAPI.md#PostV1CartsGuestcartsCartIdGiftCards) | **Post** /V1/carts/guest-carts/{cartId}/giftCards | carts/guest-carts/{cartId}/giftCards



## PostV1CartsGuestcartsCartIdGiftCards

> bool PostV1CartsGuestcartsCartIdGiftCards(ctx, cartId).PostV1CartsGuestcartsCartIdGiftCardsRequest(postV1CartsGuestcartsCartIdGiftCardsRequest).Execute()

carts/guest-carts/{cartId}/giftCards



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
	cartId := "cartId_example" // string | 
	postV1CartsGuestcartsCartIdGiftCardsRequest := *openapiclient.NewPostV1CartsGuestcartsCartIdGiftCardsRequest(*openapiclient.NewGiftCardAccountDataGiftCardAccountInterface(float32(123), float32(123), float32(123), float32(123))) // PostV1CartsGuestcartsCartIdGiftCardsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsGuestCartsCartIdGiftCardsAPI.PostV1CartsGuestcartsCartIdGiftCards(context.Background(), cartId).PostV1CartsGuestcartsCartIdGiftCardsRequest(postV1CartsGuestcartsCartIdGiftCardsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsGuestCartsCartIdGiftCardsAPI.PostV1CartsGuestcartsCartIdGiftCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsGuestcartsCartIdGiftCards`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsGuestCartsCartIdGiftCardsAPI.PostV1CartsGuestcartsCartIdGiftCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsGuestcartsCartIdGiftCardsRequest struct via the builder pattern


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

