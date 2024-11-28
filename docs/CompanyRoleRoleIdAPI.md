# \CompanyRoleRoleIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CompanyRoleRoleId**](CompanyRoleRoleIdAPI.md#DeleteV1CompanyRoleRoleId) | **Delete** /V1/company/role/{roleId} | company/role/{roleId}
[**GetV1CompanyRoleRoleId**](CompanyRoleRoleIdAPI.md#GetV1CompanyRoleRoleId) | **Get** /V1/company/role/{roleId} | company/role/{roleId}



## DeleteV1CompanyRoleRoleId

> bool DeleteV1CompanyRoleRoleId(ctx, roleId).Execute()

company/role/{roleId}



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
	resp, r, err := apiClient.CompanyRoleRoleIdAPI.DeleteV1CompanyRoleRoleId(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyRoleRoleIdAPI.DeleteV1CompanyRoleRoleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CompanyRoleRoleId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CompanyRoleRoleIdAPI.DeleteV1CompanyRoleRoleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CompanyRoleRoleIdRequest struct via the builder pattern


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


## GetV1CompanyRoleRoleId

> CompanyDataRoleInterface GetV1CompanyRoleRoleId(ctx, roleId).Execute()

company/role/{roleId}



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
	resp, r, err := apiClient.CompanyRoleRoleIdAPI.GetV1CompanyRoleRoleId(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyRoleRoleIdAPI.GetV1CompanyRoleRoleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CompanyRoleRoleId`: CompanyDataRoleInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyRoleRoleIdAPI.GetV1CompanyRoleRoleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CompanyRoleRoleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyDataRoleInterface**](CompanyDataRoleInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

