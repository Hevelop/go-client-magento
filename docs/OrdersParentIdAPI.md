# \OrdersParentIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1OrdersParentId**](OrdersParentIdAPI.md#PutV1OrdersParentId) | **Put** /V1/orders/{parent_id} | orders/{parent_id}



## PutV1OrdersParentId

> SalesDataOrderAddressInterface PutV1OrdersParentId(ctx, parentId).PutV1OrdersParentIdRequest(putV1OrdersParentIdRequest).Execute()

orders/{parent_id}



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
	parentId := "parentId_example" // string | 
	putV1OrdersParentIdRequest := *openapiclient.NewPutV1OrdersParentIdRequest(*openapiclient.NewSalesDataOrderAddressInterface("AddressType_example", "City_example", "CountryId_example", "Firstname_example", "Lastname_example", "Postcode_example", "Telephone_example")) // PutV1OrdersParentIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersParentIdAPI.PutV1OrdersParentId(context.Background(), parentId).PutV1OrdersParentIdRequest(putV1OrdersParentIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersParentIdAPI.PutV1OrdersParentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1OrdersParentId`: SalesDataOrderAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `OrdersParentIdAPI.PutV1OrdersParentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1OrdersParentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1OrdersParentIdRequest** | [**PutV1OrdersParentIdRequest**](PutV1OrdersParentIdRequest.md) |  | 

### Return type

[**SalesDataOrderAddressInterface**](SalesDataOrderAddressInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

