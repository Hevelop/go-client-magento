# \CartsGuestCartsCartIdGiftCardsGiftCardCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode**](CartsGuestCartsCartIdGiftCardsGiftCardCodeAPI.md#DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode) | **Delete** /V1/carts/guest-carts/{cartId}/giftCards/{giftCardCode} | carts/guest-carts/{cartId}/giftCards/{giftCardCode}



## DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode

> bool DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode(ctx, cartId, giftCardCode).Execute()

carts/guest-carts/{cartId}/giftCards/{giftCardCode}



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
	giftCardCode := "giftCardCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsGuestCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode(context.Background(), cartId, giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsGuestCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsGuestCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 
**giftCardCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CartsGuestcartsCartIdGiftCardsGiftCardCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

