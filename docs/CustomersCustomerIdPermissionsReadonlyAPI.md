# \CustomersCustomerIdPermissionsReadonlyAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomersCustomerIdPermissionsReadonly**](CustomersCustomerIdPermissionsReadonlyAPI.md#GetV1CustomersCustomerIdPermissionsReadonly) | **Get** /V1/customers/{customerId}/permissions/readonly | customers/{customerId}/permissions/readonly



## GetV1CustomersCustomerIdPermissionsReadonly

> bool GetV1CustomersCustomerIdPermissionsReadonly(ctx, customerId).Execute()

customers/{customerId}/permissions/readonly



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
	resp, r, err := apiClient.CustomersCustomerIdPermissionsReadonlyAPI.GetV1CustomersCustomerIdPermissionsReadonly(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersCustomerIdPermissionsReadonlyAPI.GetV1CustomersCustomerIdPermissionsReadonly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersCustomerIdPermissionsReadonly`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersCustomerIdPermissionsReadonlyAPI.GetV1CustomersCustomerIdPermissionsReadonly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersCustomerIdPermissionsReadonlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

