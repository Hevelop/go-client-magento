# \AddressesAddressIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1AddressesAddressId**](AddressesAddressIdAPI.md#DeleteV1AddressesAddressId) | **Delete** /V1/addresses/{addressId} | addresses/{addressId}



## DeleteV1AddressesAddressId

> bool DeleteV1AddressesAddressId(ctx, addressId).Execute()

addresses/{addressId}



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
	addressId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAddressIdAPI.DeleteV1AddressesAddressId(context.Background(), addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAddressIdAPI.DeleteV1AddressesAddressId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1AddressesAddressId`: bool
	fmt.Fprintf(os.Stdout, "Response from `AddressesAddressIdAPI.DeleteV1AddressesAddressId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1AddressesAddressIdRequest struct via the builder pattern


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

