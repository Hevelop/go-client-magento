# \CustomersCustomerIdCartsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CustomersCustomerIdCarts**](CustomersCustomerIdCartsAPI.md#PostV1CustomersCustomerIdCarts) | **Post** /V1/customers/{customerId}/carts | customers/{customerId}/carts



## PostV1CustomersCustomerIdCarts

> int32 PostV1CustomersCustomerIdCarts(ctx, customerId).Execute()

customers/{customerId}/carts



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
	customerId := int32(56) // int32 | The customer ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersCustomerIdCartsAPI.PostV1CustomersCustomerIdCarts(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersCustomerIdCartsAPI.PostV1CustomersCustomerIdCarts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CustomersCustomerIdCarts`: int32
	fmt.Fprintf(os.Stdout, "Response from `CustomersCustomerIdCartsAPI.PostV1CustomersCustomerIdCarts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32** | The customer ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CustomersCustomerIdCartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

