# \TaxClassesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TaxClasses**](TaxClassesAPI.md#PostV1TaxClasses) | **Post** /V1/taxClasses | taxClasses



## PostV1TaxClasses

> string PostV1TaxClasses(ctx).PostV1TaxClassesRequest(postV1TaxClassesRequest).Execute()

taxClasses



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
	postV1TaxClassesRequest := *openapiclient.NewPostV1TaxClassesRequest(*openapiclient.NewTaxDataTaxClassInterface("ClassName_example", "ClassType_example")) // PostV1TaxClassesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxClassesAPI.PostV1TaxClasses(context.Background()).PostV1TaxClassesRequest(postV1TaxClassesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxClassesAPI.PostV1TaxClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TaxClasses`: string
	fmt.Fprintf(os.Stdout, "Response from `TaxClassesAPI.PostV1TaxClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TaxClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TaxClassesRequest** | [**PostV1TaxClassesRequest**](PostV1TaxClassesRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

