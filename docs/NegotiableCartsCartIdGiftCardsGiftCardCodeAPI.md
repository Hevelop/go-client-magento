# \NegotiableCartsCartIdGiftCardsGiftCardCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode**](NegotiableCartsCartIdGiftCardsGiftCardCodeAPI.md#DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode) | **Delete** /V1/negotiable-carts/{cartId}/giftCards/{giftCardCode} | negotiable-carts/{cartId}/giftCards/{giftCardCode}



## DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode

> bool DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode(ctx, cartId, giftCardCode).Execute()

negotiable-carts/{cartId}/giftCards/{giftCardCode}



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
	resp, r, err := apiClient.NegotiableCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode(context.Background(), cartId, giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdGiftCardsGiftCardCodeAPI.DeleteV1NegotiablecartsCartIdGiftCardsGiftCardCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 
**giftCardCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1NegotiablecartsCartIdGiftCardsGiftCardCodeRequest struct via the builder pattern


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

