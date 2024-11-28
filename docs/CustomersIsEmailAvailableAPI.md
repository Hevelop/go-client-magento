# \CustomersIsEmailAvailableAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CustomersIsEmailAvailable**](CustomersIsEmailAvailableAPI.md#PostV1CustomersIsEmailAvailable) | **Post** /V1/customers/isEmailAvailable | customers/isEmailAvailable



## PostV1CustomersIsEmailAvailable

> bool PostV1CustomersIsEmailAvailable(ctx).PostV1CustomersIsEmailAvailableRequest(postV1CustomersIsEmailAvailableRequest).Execute()

customers/isEmailAvailable



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
	postV1CustomersIsEmailAvailableRequest := *openapiclient.NewPostV1CustomersIsEmailAvailableRequest("CustomerEmail_example") // PostV1CustomersIsEmailAvailableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersIsEmailAvailableAPI.PostV1CustomersIsEmailAvailable(context.Background()).PostV1CustomersIsEmailAvailableRequest(postV1CustomersIsEmailAvailableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersIsEmailAvailableAPI.PostV1CustomersIsEmailAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CustomersIsEmailAvailable`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersIsEmailAvailableAPI.PostV1CustomersIsEmailAvailable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CustomersIsEmailAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CustomersIsEmailAvailableRequest** | [**PostV1CustomersIsEmailAvailableRequest**](PostV1CustomersIsEmailAvailableRequest.md) |  | 

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

