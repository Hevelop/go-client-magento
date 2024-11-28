# \CartsQuoteIdGiftCardsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsQuoteIdGiftCards**](CartsQuoteIdGiftCardsAPI.md#GetV1CartsQuoteIdGiftCards) | **Get** /V1/carts/{quoteId}/giftCards | carts/{quoteId}/giftCards



## GetV1CartsQuoteIdGiftCards

> GiftCardAccountDataGiftCardAccountInterface GetV1CartsQuoteIdGiftCards(ctx, quoteId).Execute()

carts/{quoteId}/giftCards



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
	quoteId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsQuoteIdGiftCardsAPI.GetV1CartsQuoteIdGiftCards(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsQuoteIdGiftCardsAPI.GetV1CartsQuoteIdGiftCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsQuoteIdGiftCards`: GiftCardAccountDataGiftCardAccountInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsQuoteIdGiftCardsAPI.GetV1CartsQuoteIdGiftCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsQuoteIdGiftCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftCardAccountDataGiftCardAccountInterface**](GiftCardAccountDataGiftCardAccountInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

