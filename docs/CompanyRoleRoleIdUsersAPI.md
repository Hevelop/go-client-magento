# \CompanyRoleRoleIdUsersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CompanyRoleRoleIdUsers**](CompanyRoleRoleIdUsersAPI.md#GetV1CompanyRoleRoleIdUsers) | **Get** /V1/company/role/{roleId}/users | company/role/{roleId}/users



## GetV1CompanyRoleRoleIdUsers

> []CustomerDataCustomerInterface GetV1CompanyRoleRoleIdUsers(ctx, roleId).Execute()

company/role/{roleId}/users



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
	roleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyRoleRoleIdUsersAPI.GetV1CompanyRoleRoleIdUsers(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyRoleRoleIdUsersAPI.GetV1CompanyRoleRoleIdUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CompanyRoleRoleIdUsers`: []CustomerDataCustomerInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyRoleRoleIdUsersAPI.GetV1CompanyRoleRoleIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CompanyRoleRoleIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomerDataCustomerInterface**](CustomerDataCustomerInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

