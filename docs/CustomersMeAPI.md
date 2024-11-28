# \CustomersMeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomersMe**](CustomersMeAPI.md#GetV1CustomersMe) | **Get** /V1/customers/me | customers/me
[**PutV1CustomersMe**](CustomersMeAPI.md#PutV1CustomersMe) | **Put** /V1/customers/me | customers/me



## GetV1CustomersMe

> CustomerDataCustomerInterface GetV1CustomersMe(ctx).Execute()

customers/me



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersMeAPI.GetV1CustomersMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersMeAPI.GetV1CustomersMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersMe`: CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersMeAPI.GetV1CustomersMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersMeRequest struct via the builder pattern


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


## PutV1CustomersMe

> CustomerDataCustomerInterface PutV1CustomersMe(ctx).PutV1CustomersMeRequest(putV1CustomersMeRequest).Execute()

customers/me



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
	putV1CustomersMeRequest := *openapiclient.NewPutV1CustomersMeRequest(*openapiclient.NewCustomerDataCustomerInterface("Email_example", "Firstname_example", "Lastname_example")) // PutV1CustomersMeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersMeAPI.PutV1CustomersMe(context.Background()).PutV1CustomersMeRequest(putV1CustomersMeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersMeAPI.PutV1CustomersMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomersMe`: CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersMeAPI.PutV1CustomersMe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomersMeRequest struct via the builder pattern


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

