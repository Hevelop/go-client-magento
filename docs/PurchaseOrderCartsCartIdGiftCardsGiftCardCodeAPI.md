# \PurchaseOrderCartsCartIdGiftCardsGiftCardCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode**](PurchaseOrderCartsCartIdGiftCardsGiftCardCodeAPI.md#DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode) | **Delete** /V1/purchase-order-carts/{cartId}/giftCards/{giftCardCode} | purchase-order-carts/{cartId}/giftCards/{giftCardCode}



## DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode

> bool DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode(ctx, cartId, giftCardCode).Execute()

purchase-order-carts/{cartId}/giftCards/{giftCardCode}



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
	giftCardCode := "giftCardCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode(context.Background(), cartId, giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 
**giftCardCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1PurchaseordercartsCartIdGiftCardsGiftCardCodeRequest struct via the builder pattern


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

