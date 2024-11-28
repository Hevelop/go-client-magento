# \PurchaseOrderCartsCartIdGiftCardsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1PurchaseordercartsCartIdGiftCards**](PurchaseOrderCartsCartIdGiftCardsAPI.md#PostV1PurchaseordercartsCartIdGiftCards) | **Post** /V1/purchase-order-carts/{cartId}/giftCards | purchase-order-carts/{cartId}/giftCards



## PostV1PurchaseordercartsCartIdGiftCards

> bool PostV1PurchaseordercartsCartIdGiftCards(ctx, cartId).PostV1CartsGuestcartsCartIdGiftCardsRequest(postV1CartsGuestcartsCartIdGiftCardsRequest).Execute()

purchase-order-carts/{cartId}/giftCards



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
	cartId := int32(56) // int32 | 
	postV1CartsGuestcartsCartIdGiftCardsRequest := *openapiclient.NewPostV1CartsGuestcartsCartIdGiftCardsRequest(*openapiclient.NewGiftCardAccountDataGiftCardAccountInterface(float32(123), float32(123), float32(123), float32(123))) // PostV1CartsGuestcartsCartIdGiftCardsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderCartsCartIdGiftCardsAPI.PostV1PurchaseordercartsCartIdGiftCards(context.Background(), cartId).PostV1CartsGuestcartsCartIdGiftCardsRequest(postV1CartsGuestcartsCartIdGiftCardsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderCartsCartIdGiftCardsAPI.PostV1PurchaseordercartsCartIdGiftCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1PurchaseordercartsCartIdGiftCards`: bool
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderCartsCartIdGiftCardsAPI.PostV1PurchaseordercartsCartIdGiftCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1PurchaseordercartsCartIdGiftCardsRequest struct via the builder pattern


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

