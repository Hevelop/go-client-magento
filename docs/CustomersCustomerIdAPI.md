# \CustomersCustomerIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CustomersCustomerId**](CustomersCustomerIdAPI.md#DeleteV1CustomersCustomerId) | **Delete** /V1/customers/{customerId} | customers/{customerId}
[**GetV1CustomersCustomerId**](CustomersCustomerIdAPI.md#GetV1CustomersCustomerId) | **Get** /V1/customers/{customerId} | customers/{customerId}
[**PutV1CustomersCustomerId**](CustomersCustomerIdAPI.md#PutV1CustomersCustomerId) | **Put** /V1/customers/{customerId} | customers/{customerId}



## DeleteV1CustomersCustomerId

> bool DeleteV1CustomersCustomerId(ctx, customerId).Execute()

customers/{customerId}



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
	resp, r, err := apiClient.CustomersCustomerIdAPI.DeleteV1CustomersCustomerId(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersCustomerIdAPI.DeleteV1CustomersCustomerId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CustomersCustomerId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersCustomerIdAPI.DeleteV1CustomersCustomerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CustomersCustomerIdRequest struct via the builder pattern


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


## GetV1CustomersCustomerId

> CustomerDataCustomerInterface GetV1CustomersCustomerId(ctx, customerId).Execute()

customers/{customerId}



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
	resp, r, err := apiClient.CustomersCustomerIdAPI.GetV1CustomersCustomerId(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersCustomerIdAPI.GetV1CustomersCustomerId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersCustomerId`: CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersCustomerIdAPI.GetV1CustomersCustomerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersCustomerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDataCustomerInterface**](CustomerDataCustomerInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CustomersCustomerId

> CustomerDataCustomerInterface PutV1CustomersCustomerId(ctx, customerId).PutV1CustomersMeRequest(putV1CustomersMeRequest).Execute()

customers/{customerId}



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
	customerId := "customerId_example" // string | 
	putV1CustomersMeRequest := *openapiclient.NewPutV1CustomersMeRequest(*openapiclient.NewCustomerDataCustomerInterface("Email_example", "Firstname_example", "Lastname_example")) // PutV1CustomersMeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersCustomerIdAPI.PutV1CustomersCustomerId(context.Background(), customerId).PutV1CustomersMeRequest(putV1CustomersMeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersCustomerIdAPI.PutV1CustomersCustomerId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomersCustomerId`: CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersCustomerIdAPI.PutV1CustomersCustomerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomersCustomerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CustomersMeRequest** | [**PutV1CustomersMeRequest**](PutV1CustomersMeRequest.md) |  | 

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

