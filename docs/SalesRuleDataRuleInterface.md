# # SalesRuleDataRuleInterface
Interface RuleInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId**| **int32** | Rule id  | [optional]
**Name**| **string** | Rule name  | [optional]
**StoreLabels**| [**[]SalesRuleDataRuleLabelInterface**](SalesRuleDataRuleLabelInterface.md) | Display label  | [optional]
**Description**| **string** | Description  | [optional]
**WebsiteIds**| **[]int32** | A list of websites the rule applies to  |
**CustomerGroupIds**| **[]int32** | Ids of customer groups that the rule applies to  |
**FromDate**| **string** | The start date when the coupon is active  | [optional]
**ToDate**| **string** | The end date when the coupon is active  | [optional]
**UsesPerCustomer**| **int32** | Number of uses per customer  |
**IsActive**| **bool** | The coupon is active  |
**Condition**| [**SalesRuleDataConditionInterface**](SalesRuleDataConditionInterface.md) |   | [optional]
**ActionCondition**| [**SalesRuleDataConditionInterface**](SalesRuleDataConditionInterface.md) |   | [optional]
**StopRulesProcessing**| **bool** | To stop rule processing  |
**IsAdvanced**| **bool** | Is this field needed  |
**ProductIds**| **[]int32** | Product ids  | [optional]
**SortOrder**| **int32** | Sort order  |
**SimpleAction**| **string** | Simple action of the rule  | [optional]
**DiscountAmount**| **float32** | Discount amount  |
**DiscountQty**| **float32** | Maximum qty discount is applied  | [optional]
**DiscountStep**| **int32** | Discount step  |
**ApplyToShipping**| **bool** | The rule applies to shipping  |
**TimesUsed**| **int32** | How many times the rule has been used  |
**IsRss**| **bool** | Whether the rule is in RSS  |
**CouponType**| **string** | Coupon type  |
**UseAutoGeneration**| **bool** | To auto generate coupon  |
**UsesPerCoupon**| **int32** | Limit of uses per coupon  |
**SimpleFreeShipping**| **string** | To grant free shipping  | [optional]
**ExtensionAttributes**| [**SalesRuleDataRuleExtensionInterface**](SalesRuleDataRuleExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

