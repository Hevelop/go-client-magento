# # SalesDataOrderStatusHistoryInterface
Order status history interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment**| **string** | Comment.  |
**CreatedAt**| **string** | Created-at timestamp.  | [optional]
**EntityId**| **int32** | Order status history ID.  | [optional]
**EntityName**| **string** | Entity name.  | [optional]
**IsCustomerNotified**| **int32** | Is-customer-notified flag value.  |
**IsVisibleOnFront**| **int32** | Is-visible-on-storefront flag value.  |
**ParentId**| **int32** | Parent ID.  |
**Status**| **string** | Status.  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\OrderStatusHistoryInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

