# \PaymentsConfigSmartButtonsLocationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PaymentsconfigSmartbuttonsLocation**](PaymentsConfigSmartButtonsLocationAPI.md#GetV1PaymentsconfigSmartbuttonsLocation) | **Get** /V1/payments-config/smart-buttons/{location} | payments-config/smart-buttons/{location}



## GetV1PaymentsconfigSmartbuttonsLocation

> PaymentServicesPaypalDataPaymentConfigSmartButtonsInterface GetV1PaymentsconfigSmartbuttonsLocation(ctx, location).Execute()

payments-config/smart-buttons/{location}



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
	resp, r, err := apiClient.PaymentsConfigSmartButtonsLocationAPI.GetV1PaymentsconfigSmartbuttonsLocation(context.Background(), location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsConfigSmartButtonsLocationAPI.GetV1PaymentsconfigSmartbuttonsLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PaymentsconfigSmartbuttonsLocation`: PaymentServicesPaypalDataPaymentConfigSmartButtonsInterface
	fmt.Fprintf(os.Stdout, "Response from `PaymentsConfigSmartButtonsLocationAPI.GetV1PaymentsconfigSmartbuttonsLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | sdk location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PaymentsconfigSmartbuttonsLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentServicesPaypalDataPaymentConfigSmartButtonsInterface**](PaymentServicesPaypalDataPaymentConfigSmartButtonsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

