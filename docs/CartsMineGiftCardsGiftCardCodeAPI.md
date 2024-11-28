# \CartsMineGiftCardsGiftCardCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CartsMineGiftCardsGiftCardCode**](CartsMineGiftCardsGiftCardCodeAPI.md#DeleteV1CartsMineGiftCardsGiftCardCode) | **Delete** /V1/carts/mine/giftCards/{giftCardCode} | carts/mine/giftCards/{giftCardCode}



## DeleteV1CartsMineGiftCardsGiftCardCode

> bool DeleteV1CartsMineGiftCardsGiftCardCode(ctx, giftCardCode).Execute()

carts/mine/giftCards/{giftCardCode}



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
	giftCardCode := "giftCardCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineGiftCardsGiftCardCodeAPI.DeleteV1CartsMineGiftCardsGiftCardCode(context.Background(), giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineGiftCardsGiftCardCodeAPI.DeleteV1CartsMineGiftCardsGiftCardCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CartsMineGiftCardsGiftCardCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineGiftCardsGiftCardCodeAPI.DeleteV1CartsMineGiftCardsGiftCardCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CartsMineGiftCardsGiftCardCodeRequest struct via the builder pattern


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

