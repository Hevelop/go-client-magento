# \CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkTokenAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken**](CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkTokenAPI.md#GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken) | **Get** /V1/customers/{customerId}/password/resetLinkToken/{resetPasswordLinkToken} | customers/{customerId}/password/resetLinkToken/{resetPasswordLinkToken}



## GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken

> bool GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken(ctx, customerId, resetPasswordLinkToken).Execute()

customers/{customerId}/password/resetLinkToken/{resetPasswordLinkToken}



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
	customerId := int32(56) // int32 | If null is given then a customer will be matched by the RP token.
	resetPasswordLinkToken := "resetPasswordLinkToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkTokenAPI.GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken(context.Background(), customerId, resetPasswordLinkToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkTokenAPI.GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkTokenAPI.GetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32** | If null is given then a customer will be matched by the RP token. | 
**resetPasswordLinkToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersCustomerIdPasswordResetLinkTokenResetPasswordLinkTokenRequest struct via the builder pattern


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

