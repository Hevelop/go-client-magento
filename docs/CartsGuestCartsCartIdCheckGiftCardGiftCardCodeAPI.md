# \CartsGuestCartsCartIdCheckGiftCardGiftCardCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode**](CartsGuestCartsCartIdCheckGiftCardGiftCardCodeAPI.md#GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode) | **Get** /V1/carts/guest-carts/{cartId}/checkGiftCard/{giftCardCode} | carts/guest-carts/{cartId}/checkGiftCard/{giftCardCode}



## GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode

> float32 GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode(ctx, cartId, giftCardCode).Execute()

carts/guest-carts/{cartId}/checkGiftCard/{giftCardCode}



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
	resp, r, err := apiClient.CartsGuestCartsCartIdCheckGiftCardGiftCardCodeAPI.GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode(context.Background(), cartId, giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsGuestCartsCartIdCheckGiftCardGiftCardCodeAPI.GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode`: float32
	fmt.Fprintf(os.Stdout, "Response from `CartsGuestCartsCartIdCheckGiftCardGiftCardCodeAPI.GetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 
**giftCardCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsGuestcartsCartIdCheckGiftCardGiftCardCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

