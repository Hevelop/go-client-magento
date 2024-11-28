# \CustomersCustomerIdBillingAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomersCustomerIdBillingAddress**](CustomersCustomerIdBillingAddressAPI.md#GetV1CustomersCustomerIdBillingAddress) | **Get** /V1/customers/{customerId}/billingAddress | customers/{customerId}/billingAddress



## GetV1CustomersCustomerIdBillingAddress

> CustomerDataAddressInterface GetV1CustomersCustomerIdBillingAddress(ctx, customerId).Execute()

customers/{customerId}/billingAddress



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
	customerId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersCustomerIdBillingAddressAPI.GetV1CustomersCustomerIdBillingAddress(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersCustomerIdBillingAddressAPI.GetV1CustomersCustomerIdBillingAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersCustomerIdBillingAddress`: CustomerDataAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersCustomerIdBillingAddressAPI.GetV1CustomersCustomerIdBillingAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersCustomerIdBillingAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDataAddressInterface**](CustomerDataAddressInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

