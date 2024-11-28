# \PaymentsSdkLocationMethodCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PaymentssdkLocationMethodCode**](PaymentsSdkLocationMethodCodeAPI.md#GetV1PaymentssdkLocationMethodCode) | **Get** /V1/payments-sdk/{location}/{methodCode} | payments-sdk/{location}/{methodCode}



## GetV1PaymentssdkLocationMethodCode

> PaymentServicesPaypalDataPaymentSdkParamsInterface GetV1PaymentssdkLocationMethodCode(ctx, location, methodCode).Execute()

payments-sdk/{location}/{methodCode}



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
	methodCode := "methodCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsSdkLocationMethodCodeAPI.GetV1PaymentssdkLocationMethodCode(context.Background(), location, methodCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsSdkLocationMethodCodeAPI.GetV1PaymentssdkLocationMethodCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PaymentssdkLocationMethodCode`: PaymentServicesPaypalDataPaymentSdkParamsInterface
	fmt.Fprintf(os.Stdout, "Response from `PaymentsSdkLocationMethodCodeAPI.GetV1PaymentssdkLocationMethodCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** |  | 
**methodCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PaymentssdkLocationMethodCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentServicesPaypalDataPaymentSdkParamsInterface**](PaymentServicesPaypalDataPaymentSdkParamsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

