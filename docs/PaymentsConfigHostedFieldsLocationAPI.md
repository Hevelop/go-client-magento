# \PaymentsConfigHostedFieldsLocationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PaymentsconfigHostedfieldsLocation**](PaymentsConfigHostedFieldsLocationAPI.md#GetV1PaymentsconfigHostedfieldsLocation) | **Get** /V1/payments-config/hosted-fields/{location} | payments-config/hosted-fields/{location}



## GetV1PaymentsconfigHostedfieldsLocation

> PaymentServicesPaypalDataPaymentConfigHostedFieldsInterface GetV1PaymentsconfigHostedfieldsLocation(ctx, location).Execute()

payments-config/hosted-fields/{location}



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
	resp, r, err := apiClient.PaymentsConfigHostedFieldsLocationAPI.GetV1PaymentsconfigHostedfieldsLocation(context.Background(), location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsConfigHostedFieldsLocationAPI.GetV1PaymentsconfigHostedfieldsLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PaymentsconfigHostedfieldsLocation`: PaymentServicesPaypalDataPaymentConfigHostedFieldsInterface
	fmt.Fprintf(os.Stdout, "Response from `PaymentsConfigHostedFieldsLocationAPI.GetV1PaymentsconfigHostedfieldsLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | sdk location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PaymentsconfigHostedfieldsLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentServicesPaypalDataPaymentConfigHostedFieldsInterface**](PaymentServicesPaypalDataPaymentConfigHostedFieldsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

