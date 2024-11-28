# \CustomersMeActivateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CustomersMeActivate**](CustomersMeActivateAPI.md#PutV1CustomersMeActivate) | **Put** /V1/customers/me/activate | customers/me/activate



## PutV1CustomersMeActivate

> CustomerDataCustomerInterface PutV1CustomersMeActivate(ctx).PutV1CustomersMeActivateRequest(putV1CustomersMeActivateRequest).Execute()

customers/me/activate



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
	putV1CustomersMeActivateRequest := *openapiclient.NewPutV1CustomersMeActivateRequest("ConfirmationKey_example") // PutV1CustomersMeActivateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersMeActivateAPI.PutV1CustomersMeActivate(context.Background()).PutV1CustomersMeActivateRequest(putV1CustomersMeActivateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersMeActivateAPI.PutV1CustomersMeActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomersMeActivate`: CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersMeActivateAPI.PutV1CustomersMeActivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomersMeActivateRequest struct via the builder pattern


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

