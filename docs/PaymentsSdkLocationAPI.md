# \PaymentsSdkLocationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PaymentssdkLocation**](PaymentsSdkLocationAPI.md#GetV1PaymentssdkLocation) | **Get** /V1/payments-sdk/{location} | payments-sdk/{location}



## GetV1PaymentssdkLocation

> []PaymentServicesPaypalDataPaymentSdkParamsInterface GetV1PaymentssdkLocation(ctx, location).Execute()

payments-sdk/{location}



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
	location := "location_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsSdkLocationAPI.GetV1PaymentssdkLocation(context.Background(), location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsSdkLocationAPI.GetV1PaymentssdkLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PaymentssdkLocation`: []PaymentServicesPaypalDataPaymentSdkParamsInterface
	fmt.Fprintf(os.Stdout, "Response from `PaymentsSdkLocationAPI.GetV1PaymentssdkLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PaymentssdkLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PaymentServicesPaypalDataPaymentSdkParamsInterface**](PaymentServicesPaypalDataPaymentSdkParamsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

