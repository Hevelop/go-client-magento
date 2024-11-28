# \CustomersConfirmAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CustomersConfirm**](CustomersConfirmAPI.md#PostV1CustomersConfirm) | **Post** /V1/customers/confirm | customers/confirm



## PostV1CustomersConfirm

> bool PostV1CustomersConfirm(ctx).PostV1CustomersConfirmRequest(postV1CustomersConfirmRequest).Execute()

customers/confirm



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
	postV1CustomersConfirmRequest := *openapiclient.NewPostV1CustomersConfirmRequest("Email_example", int32(123)) // PostV1CustomersConfirmRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersConfirmAPI.PostV1CustomersConfirm(context.Background()).PostV1CustomersConfirmRequest(postV1CustomersConfirmRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersConfirmAPI.PostV1CustomersConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CustomersConfirm`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersConfirmAPI.PostV1CustomersConfirm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CustomersConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CustomersConfirmRequest** | [**PostV1CustomersConfirmRequest**](PostV1CustomersConfirmRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

