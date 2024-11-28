# \CustomersEmailActivateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CustomersEmailActivate**](CustomersEmailActivateAPI.md#PutV1CustomersEmailActivate) | **Put** /V1/customers/{email}/activate | customers/{email}/activate



## PutV1CustomersEmailActivate

> CustomerDataCustomerInterface PutV1CustomersEmailActivate(ctx, email).PutV1CustomersMeActivateRequest(putV1CustomersMeActivateRequest).Execute()

customers/{email}/activate



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
	email := "email_example" // string | 
	putV1CustomersMeActivateRequest := *openapiclient.NewPutV1CustomersMeActivateRequest("ConfirmationKey_example") // PutV1CustomersMeActivateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersEmailActivateAPI.PutV1CustomersEmailActivate(context.Background(), email).PutV1CustomersMeActivateRequest(putV1CustomersMeActivateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersEmailActivateAPI.PutV1CustomersEmailActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomersEmailActivate`: CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersEmailActivateAPI.PutV1CustomersEmailActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomersEmailActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CustomersMeActivateRequest** | [**PutV1CustomersMeActivateRequest**](PutV1CustomersMeActivateRequest.md) |  | 

### Return type

[**CustomerDataCustomerInterface**](CustomerDataCustomerInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

