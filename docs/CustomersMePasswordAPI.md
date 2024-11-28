# \CustomersMePasswordAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CustomersMePassword**](CustomersMePasswordAPI.md#PutV1CustomersMePassword) | **Put** /V1/customers/me/password | customers/me/password



## PutV1CustomersMePassword

> bool PutV1CustomersMePassword(ctx).PutV1CustomersMePasswordRequest(putV1CustomersMePasswordRequest).Execute()

customers/me/password



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
	putV1CustomersMePasswordRequest := *openapiclient.NewPutV1CustomersMePasswordRequest("CurrentPassword_example", "NewPassword_example") // PutV1CustomersMePasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersMePasswordAPI.PutV1CustomersMePassword(context.Background()).PutV1CustomersMePasswordRequest(putV1CustomersMePasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersMePasswordAPI.PutV1CustomersMePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomersMePassword`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersMePasswordAPI.PutV1CustomersMePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomersMePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1CustomersMePasswordRequest** | [**PutV1CustomersMePasswordRequest**](PutV1CustomersMePasswordRequest.md) |  | 

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

