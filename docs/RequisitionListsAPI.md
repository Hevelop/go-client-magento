# \RequisitionListsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1RequisitionLists**](RequisitionListsAPI.md#PostV1RequisitionLists) | **Post** /V1/requisition_lists | requisition_lists



## PostV1RequisitionLists

> RequisitionListDataRequisitionListInterface PostV1RequisitionLists(ctx).PostV1RequisitionListsRequest(postV1RequisitionListsRequest).Execute()

requisition_lists



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
	postV1RequisitionListsRequest := *openapiclient.NewPostV1RequisitionListsRequest(*openapiclient.NewRequisitionListDataRequisitionListInterface(int32(123), int32(123), "Name_example", "UpdatedAt_example", "Description_example", []openapiclient.RequisitionListDataRequisitionListItemInterface{*openapiclient.NewRequisitionListDataRequisitionListItemInterface(int32(123), "Sku_example", int32(123), float32(123), []string{"Options_example"}, int32(123), "AddedAt_example")})) // PostV1RequisitionListsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequisitionListsAPI.PostV1RequisitionLists(context.Background()).PostV1RequisitionListsRequest(postV1RequisitionListsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequisitionListsAPI.PostV1RequisitionLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1RequisitionLists`: RequisitionListDataRequisitionListInterface
	fmt.Fprintf(os.Stdout, "Response from `RequisitionListsAPI.PostV1RequisitionLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1RequisitionListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1RequisitionListsRequest** | [**PostV1RequisitionListsRequest**](PostV1RequisitionListsRequest.md) |  | 

### Return type

[**RequisitionListDataRequisitionListInterface**](RequisitionListDataRequisitionListInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

