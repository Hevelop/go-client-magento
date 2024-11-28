# \CompanyRoleIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CompanyRoleId**](CompanyRoleIdAPI.md#PutV1CompanyRoleId) | **Put** /V1/company/role/{id} | company/role/{id}



## PutV1CompanyRoleId

> CompanyDataRoleInterface PutV1CompanyRoleId(ctx, id).PostV1CompanyRoleRequest(postV1CompanyRoleRequest).Execute()

company/role/{id}



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
	id := "id_example" // string | 
	postV1CompanyRoleRequest := *openapiclient.NewPostV1CompanyRoleRequest(*openapiclient.NewCompanyDataRoleInterface([]openapiclient.CompanyDataPermissionInterface{*openapiclient.NewCompanyDataPermissionInterface("ResourceId_example", "Permission_example")})) // PostV1CompanyRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyRoleIdAPI.PutV1CompanyRoleId(context.Background(), id).PostV1CompanyRoleRequest(postV1CompanyRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyRoleIdAPI.PutV1CompanyRoleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CompanyRoleId`: CompanyDataRoleInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyRoleIdAPI.PutV1CompanyRoleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CompanyRoleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CompanyRoleRequest** | [**PostV1CompanyRoleRequest**](PostV1CompanyRoleRequest.md) |  | 

### Return type

[**CompanyDataRoleInterface**](CompanyDataRoleInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

