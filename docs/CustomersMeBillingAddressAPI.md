# \CustomersMeBillingAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomersMeBillingAddress**](CustomersMeBillingAddressAPI.md#GetV1CustomersMeBillingAddress) | **Get** /V1/customers/me/billingAddress | customers/me/billingAddress



## GetV1CustomersMeBillingAddress

> CustomerDataAddressInterface GetV1CustomersMeBillingAddress(ctx).Execute()

customers/me/billingAddress



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
	resp, r, err := apiClient.CustomersMeBillingAddressAPI.GetV1CustomersMeBillingAddress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersMeBillingAddressAPI.GetV1CustomersMeBillingAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomersMeBillingAddress`: CustomerDataAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomersMeBillingAddressAPI.GetV1CustomersMeBillingAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomersMeBillingAddressRequest struct via the builder pattern


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

