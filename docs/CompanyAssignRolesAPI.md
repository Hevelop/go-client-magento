# \CompanyAssignRolesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CompanyAssignRoles**](CompanyAssignRolesAPI.md#PutV1CompanyAssignRoles) | **Put** /V1/company/assignRoles | company/assignRoles



## PutV1CompanyAssignRoles

> bool PutV1CompanyAssignRoles(ctx).PutV1CompanyAssignRolesRequest(putV1CompanyAssignRolesRequest).Execute()

company/assignRoles



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
	putV1CompanyAssignRolesRequest := *openapiclient.NewPutV1CompanyAssignRolesRequest(int32(123), []openapiclient.CompanyDataRoleInterface{*openapiclient.NewCompanyDataRoleInterface([]openapiclient.CompanyDataPermissionInterface{*openapiclient.NewCompanyDataPermissionInterface("ResourceId_example", "Permission_example")})}) // PutV1CompanyAssignRolesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAssignRolesAPI.PutV1CompanyAssignRoles(context.Background()).PutV1CompanyAssignRolesRequest(putV1CompanyAssignRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAssignRolesAPI.PutV1CompanyAssignRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CompanyAssignRoles`: bool
	fmt.Fprintf(os.Stdout, "Response from `CompanyAssignRolesAPI.PutV1CompanyAssignRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CompanyAssignRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1CompanyAssignRolesRequest** | [**PutV1CompanyAssignRolesRequest**](PutV1CompanyAssignRolesRequest.md) |  | 

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

