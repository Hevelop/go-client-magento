# \CompanyRoleAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CompanyRole**](CompanyRoleAPI.md#GetV1CompanyRole) | **Get** /V1/company/role/ | company/role/
[**PostV1CompanyRole**](CompanyRoleAPI.md#PostV1CompanyRole) | **Post** /V1/company/role/ | company/role/



## GetV1CompanyRole

> CompanyDataRoleSearchResultsInterface GetV1CompanyRole(ctx).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()

company/role/



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
	searchCriteriaFilterGroups0Filters0Field := "searchCriteriaFilterGroups0Filters0Field_example" // string | Field (optional)
	searchCriteriaFilterGroups0Filters0Value := "searchCriteriaFilterGroups0Filters0Value_example" // string | Value (optional)
	searchCriteriaFilterGroups0Filters0ConditionType := "searchCriteriaFilterGroups0Filters0ConditionType_example" // string | Condition type (optional)
	searchCriteriaSortOrders0Field := "searchCriteriaSortOrders0Field_example" // string | Sorting field. (optional)
	searchCriteriaSortOrders0Direction := "searchCriteriaSortOrders0Direction_example" // string | Sorting direction. (optional)
	searchCriteriaPageSize := int32(56) // int32 | Page size. (optional)
	searchCriteriaCurrentPage := int32(56) // int32 | Current page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyRoleAPI.GetV1CompanyRole(context.Background()).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyRoleAPI.GetV1CompanyRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CompanyRole`: CompanyDataRoleSearchResultsInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyRoleAPI.GetV1CompanyRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CompanyRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchCriteriaFilterGroups0Filters0Field** | **string** | Field | 
 **searchCriteriaFilterGroups0Filters0Value** | **string** | Value | 
 **searchCriteriaFilterGroups0Filters0ConditionType** | **string** | Condition type | 
 **searchCriteriaSortOrders0Field** | **string** | Sorting field. | 
 **searchCriteriaSortOrders0Direction** | **string** | Sorting direction. | 
 **searchCriteriaPageSize** | **int32** | Page size. | 
 **searchCriteriaCurrentPage** | **int32** | Current page. | 

### Return type

[**CompanyDataRoleSearchResultsInterface**](CompanyDataRoleSearchResultsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1CompanyRole

> CompanyDataRoleInterface PostV1CompanyRole(ctx).PostV1CompanyRoleRequest(postV1CompanyRoleRequest).Execute()

company/role/



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
	postV1CompanyRoleRequest := *openapiclient.NewPostV1CompanyRoleRequest(*openapiclient.NewCompanyDataRoleInterface([]openapiclient.CompanyDataPermissionInterface{*openapiclient.NewCompanyDataPermissionInterface("ResourceId_example", "Permission_example")})) // PostV1CompanyRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyRoleAPI.PostV1CompanyRole(context.Background()).PostV1CompanyRoleRequest(postV1CompanyRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyRoleAPI.PostV1CompanyRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CompanyRole`: CompanyDataRoleInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyRoleAPI.PostV1CompanyRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CompanyRoleRequest struct via the builder pattern


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

