# \BraintreeMinePaymentVaultAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1BraintreeMinePaymentVault**](BraintreeMinePaymentVaultAPI.md#PostV1BraintreeMinePaymentVault) | **Post** /V1/braintree/mine/payment/vault | braintree/mine/payment/vault



## PostV1BraintreeMinePaymentVault

> bool PostV1BraintreeMinePaymentVault(ctx).PostV1BraintreeMinePaymentVaultRequest(postV1BraintreeMinePaymentVaultRequest).Execute()

braintree/mine/payment/vault



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
	postV1BraintreeMinePaymentVaultRequest := *openapiclient.NewPostV1BraintreeMinePaymentVaultRequest(*openapiclient.NewPayPalBraintreeDataPaymentInterface()) // PostV1BraintreeMinePaymentVaultRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BraintreeMinePaymentVaultAPI.PostV1BraintreeMinePaymentVault(context.Background()).PostV1BraintreeMinePaymentVaultRequest(postV1BraintreeMinePaymentVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BraintreeMinePaymentVaultAPI.PostV1BraintreeMinePaymentVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1BraintreeMinePaymentVault`: bool
	fmt.Fprintf(os.Stdout, "Response from `BraintreeMinePaymentVaultAPI.PostV1BraintreeMinePaymentVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1BraintreeMinePaymentVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1BraintreeMinePaymentVaultRequest** | [**PostV1BraintreeMinePaymentVaultRequest**](PostV1BraintreeMinePaymentVaultRequest.md) |  | 

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

