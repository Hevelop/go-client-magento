# \PaymentsConfigApplePayLocationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PaymentsconfigApplepayLocation**](PaymentsConfigApplePayLocationAPI.md#GetV1PaymentsconfigApplepayLocation) | **Get** /V1/payments-config/apple-pay/{location} | payments-config/apple-pay/{location}



## GetV1PaymentsconfigApplepayLocation

> PaymentServicesPaypalDataPaymentConfigApplePayInterface GetV1PaymentsconfigApplepayLocation(ctx, location).Execute()

payments-config/apple-pay/{location}



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
	location := "location_example" // string | sdk location.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsConfigApplePayLocationAPI.GetV1PaymentsconfigApplepayLocation(context.Background(), location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsConfigApplePayLocationAPI.GetV1PaymentsconfigApplepayLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PaymentsconfigApplepayLocation`: PaymentServicesPaypalDataPaymentConfigApplePayInterface
	fmt.Fprintf(os.Stdout, "Response from `PaymentsConfigApplePayLocationAPI.GetV1PaymentsconfigApplepayLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | sdk location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PaymentsconfigApplepayLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentServicesPaypalDataPaymentConfigApplePayInterface**](PaymentServicesPaypalDataPaymentConfigApplePayInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

