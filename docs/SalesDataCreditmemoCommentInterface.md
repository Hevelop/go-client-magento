# # SalesDataCreditmemoCommentInterface
Credit memo comment interface. After a customer places and pays for an order and an invoice has been issued, the merchant can create a credit memo to refund all or part of the amount paid for any returned or undelivered items. The memo restores funds to the customer account so that the customer can make future purchases. A credit memo usually includes comments that detail why the credit memo amount was credited to the customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment**| **string** | Comment.  |
**CreatedAt**| **string** | Created-at timestamp.  | [optional]
**EntityId**| **int32** | Credit memo ID.  | [optional]
**IsCustomerNotified**| **int32** | Is-customer-notified flag value.  |
**IsVisibleOnFront**| **int32** | Is-visible-on-storefront flag value.  |
**ParentId**| **int32** | Parent ID.  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\CreditmemoCommentInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

