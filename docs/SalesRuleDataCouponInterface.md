# # SalesRuleDataCouponInterface
Interface CouponInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponId**| **int32** | Coupon id  | [optional]
**RuleId**| **int32** | The id of the rule associated with the coupon  |
**Code**| **string** | Coupon code  | [optional]
**UsageLimit**| **int32** | Usage limit  | [optional]
**UsagePerCustomer**| **int32** | Usage limit per customer  | [optional]
**TimesUsed**| **int32** | The number of times the coupon has been used  |
**ExpirationDate**| **string** | Expiration date  | [optional]
**IsPrimary**| **bool** | The coupon is primary coupon for the rule that it&#39;s associated with  |
**CreatedAt**| **string** | When the coupon is created  | [optional]
**Type**| **int32** | Of coupon  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\SalesRule\\Api\\Data\\CouponInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

