# \PaymentsConfigLocationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PaymentsconfigLocation**](PaymentsConfigLocationAPI.md#GetV1PaymentsconfigLocation) | **Get** /V1/payments-config/{location} | payments-config/{location}



## GetV1PaymentsconfigLocation

> PaymentServicesPaypalPaymentConfigResponseInterface GetV1PaymentsconfigLocation(ctx, location).Execute()

payments-config/{location}



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
	resp, r, err := apiClient.PaymentsConfigLocationAPI.GetV1PaymentsconfigLocation(context.Background(), location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsConfigLocationAPI.GetV1PaymentsconfigLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PaymentsconfigLocation`: PaymentServicesPaypalPaymentConfigResponseInterface
	fmt.Fprintf(os.Stdout, "Response from `PaymentsConfigLocationAPI.GetV1PaymentsconfigLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | sdk location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PaymentsconfigLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentServicesPaypalPaymentConfigResponseInterface**](PaymentServicesPaypalPaymentConfigResponseInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

