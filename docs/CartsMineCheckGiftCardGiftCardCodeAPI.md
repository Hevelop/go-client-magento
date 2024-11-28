# \CartsMineCheckGiftCardGiftCardCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMineCheckGiftCardGiftCardCode**](CartsMineCheckGiftCardGiftCardCodeAPI.md#GetV1CartsMineCheckGiftCardGiftCardCode) | **Get** /V1/carts/mine/checkGiftCard/{giftCardCode} | carts/mine/checkGiftCard/{giftCardCode}



## GetV1CartsMineCheckGiftCardGiftCardCode

> float32 GetV1CartsMineCheckGiftCardGiftCardCode(ctx, giftCardCode).Execute()

carts/mine/checkGiftCard/{giftCardCode}



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
	resp, r, err := apiClient.CartsMineCheckGiftCardGiftCardCodeAPI.GetV1CartsMineCheckGiftCardGiftCardCode(context.Background(), giftCardCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCheckGiftCardGiftCardCodeAPI.GetV1CartsMineCheckGiftCardGiftCardCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMineCheckGiftCardGiftCardCode`: float32
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCheckGiftCardGiftCardCodeAPI.GetV1CartsMineCheckGiftCardGiftCardCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineCheckGiftCardGiftCardCodeRequest struct via the builder pattern


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

