# \CustomersValidateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CustomersValidate**](CustomersValidateAPI.md#PutV1CustomersValidate) | **Put** /V1/customers/validate | customers/validate



## PutV1CustomersValidate

> CustomerDataValidationResultsInterface PutV1CustomersValidate(ctx).PutV1CustomersValidateRequest(putV1CustomersValidateRequest).Execute()

customers/validate



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
	putV1CustomersValidateRequest := *openapiclient.NewPutV1CustomersValidateRequest(*openapiclient.NewCustomerDataCustomerInterface("Email_example", "Firstname_example", "Lastname_example")) // PutV1CustomersValidateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersValidateAPI.PutV1CustomersValidate(context.Background()).PutV1CustomersValidateRequest(putV1CustomersValidateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersValidateAPI.PutV1CustomersValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomersValidate`: CustomerDataValidationResultsInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersValidateAPI.PutV1CustomersValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomersValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1CustomersValidateRequest** | [**PutV1CustomersValidateRequest**](PutV1CustomersValidateRequest.md) |  | 

### Return type

[**CustomerDataValidationResultsInterface**](CustomerDataValidationResultsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

