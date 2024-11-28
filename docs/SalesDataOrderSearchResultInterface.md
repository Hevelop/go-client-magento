# # SalesDataOrderSearchResultInterface
Order search result interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items**| [**[]SalesDataOrderInterface**](SalesDataOrderInterface.md) | Array of collection items.  |
**SearchCriteria**| [**FrameworkSearchCriteriaInterface**](FrameworkSearchCriteriaInterface.md) |   |
**TotalCount**| **int32** | Total count.  |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

