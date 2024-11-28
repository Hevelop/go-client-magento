# \CustomersResetPasswordAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CustomersResetPassword**](CustomersResetPasswordAPI.md#PostV1CustomersResetPassword) | **Post** /V1/customers/resetPassword | customers/resetPassword



## PostV1CustomersResetPassword

> bool PostV1CustomersResetPassword(ctx).PostV1CustomersResetPasswordRequest(postV1CustomersResetPasswordRequest).Execute()

customers/resetPassword



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
	postV1CustomersResetPasswordRequest := *openapiclient.NewPostV1CustomersResetPasswordRequest("Email_example", "ResetToken_example", "NewPassword_example") // PostV1CustomersResetPasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersResetPasswordAPI.PostV1CustomersResetPassword(context.Background()).PostV1CustomersResetPasswordRequest(postV1CustomersResetPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersResetPasswordAPI.PostV1CustomersResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CustomersResetPassword`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersResetPasswordAPI.PostV1CustomersResetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CustomersResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CustomersResetPasswordRequest** | [**PostV1CustomersResetPasswordRequest**](PostV1CustomersResetPasswordRequest.md) |  | 

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

