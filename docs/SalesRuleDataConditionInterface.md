# # SalesRuleDataConditionInterface
Interface ConditionInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType**| **string** | Condition type  |
**Conditions**| [**[]SalesRuleDataConditionInterface**](SalesRuleDataConditionInterface.md) | List of conditions  | [optional]
**AggregatorType**| **string** | The aggregator type  | [optional]
**Operator**| **string** | The operator of the condition  |
**AttributeName**| **string** | The attribute name of the condition  | [optional]
**Value**| **string** | The value of the condition  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\SalesRule\\Api\\Data\\ConditionInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

