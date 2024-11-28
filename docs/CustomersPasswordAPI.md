# \CustomersPasswordAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CustomersPassword**](CustomersPasswordAPI.md#PutV1CustomersPassword) | **Put** /V1/customers/password | customers/password



## PutV1CustomersPassword

> bool PutV1CustomersPassword(ctx).PutV1CustomersPasswordRequest(putV1CustomersPasswordRequest).Execute()

customers/password



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
	putV1CustomersPasswordRequest := *openapiclient.NewPutV1CustomersPasswordRequest("Email_example", "Template_example") // PutV1CustomersPasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersPasswordAPI.PutV1CustomersPassword(context.Background()).PutV1CustomersPasswordRequest(putV1CustomersPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersPasswordAPI.PutV1CustomersPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomersPassword`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersPasswordAPI.PutV1CustomersPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomersPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1CustomersPasswordRequest** | [**PutV1CustomersPasswordRequest**](PutV1CustomersPasswordRequest.md) |  | 

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

