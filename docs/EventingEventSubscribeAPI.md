# \EventingEventSubscribeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1EventingEventSubscribe**](EventingEventSubscribeAPI.md#PostV1EventingEventSubscribe) | **Post** /V1/eventing/eventSubscribe | eventing/eventSubscribe



## PostV1EventingEventSubscribe

> ErrorResponse PostV1EventingEventSubscribe(ctx).PostV1EventingEventSubscribeRequest(postV1EventingEventSubscribeRequest).Execute()

eventing/eventSubscribe



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
	postV1EventingEventSubscribeRequest := *openapiclient.NewPostV1EventingEventSubscribeRequest(*openapiclient.NewAdobeCommerceEventsClientDataEventDataInterface("Name_example", "Parent_example", []openapiclient.AdobeCommerceEventsClientDataEventFieldInterface{*openapiclient.NewAdobeCommerceEventsClientDataEventFieldInterface("Name_example", "Converter_example")}, []openapiclient.AdobeCommerceEventsClientDataEventRuleInterface{*openapiclient.NewAdobeCommerceEventsClientDataEventRuleInterface("Field_example", "Operator_example", "Value_example")}, "Destination_example", false)) // PostV1EventingEventSubscribeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventingEventSubscribeAPI.PostV1EventingEventSubscribe(context.Background()).PostV1EventingEventSubscribeRequest(postV1EventingEventSubscribeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventingEventSubscribeAPI.PostV1EventingEventSubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1EventingEventSubscribe`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `EventingEventSubscribeAPI.PostV1EventingEventSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1EventingEventSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1EventingEventSubscribeRequest** | [**PostV1EventingEventSubscribeRequest**](PostV1EventingEventSubscribeRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

