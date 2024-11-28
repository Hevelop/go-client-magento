# \CustomersAddressesAddressIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomersAddressesAddressId**](CustomersAddressesAddressIdAPI.md#GetV1CustomersAddressesAddressId) | **Get** /V1/customers/addresses/{addressId} | customers/addresses/{addressId}



## GetV1CustomersAddressesAddressId

> CustomerDataAddressInterface GetV1CustomersAddressesAddressId(ctx, addressId).Execute()

customers/addresses/{addressId}



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
	addressId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAddressesAddressIdAPI.GetV1CustomersAddressesAddressId(context.Background(), addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAddressesAddressIdAPI.GetV1CustomersAddressesAddressId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersAddressesAddressId`: CustomerDataAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersAddressesAddressIdAPI.GetV1CustomersAddressesAddressId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersAddressesAddressIdRequest struct via the builder pattern


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

