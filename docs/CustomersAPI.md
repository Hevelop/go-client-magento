# \CustomersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1Customers**](CustomersAPI.md#PostV1Customers) | **Post** /V1/customers | customers



## PostV1Customers

> CustomerDataCustomerInterface PostV1Customers(ctx).PostV1CustomersRequest(postV1CustomersRequest).Execute()

customers



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
	postV1CustomersRequest := *openapiclient.NewPostV1CustomersRequest(*openapiclient.NewCustomerDataCustomerInterface("Email_example", "Firstname_example", "Lastname_example")) // PostV1CustomersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.PostV1Customers(context.Background()).PostV1CustomersRequest(postV1CustomersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.PostV1Customers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Customers`: CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.PostV1Customers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CustomersRequest** | [**PostV1CustomersRequest**](PostV1CustomersRequest.md) |  | 

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

