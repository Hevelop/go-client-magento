/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
)

// checks if the QuoteDataTotalsExtensionInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDataTotalsExtensionInterface{}

// QuoteDataTotalsExtensionInterface ExtensionInterface class for @see \\Magento\\Quote\\Api\\Data\\TotalsInterface
type QuoteDataTotalsExtensionInterface struct {
	CouponLabel               *string                                            `json:"coupon_label,omitempty"`
	NegotiableQuoteTotals     *NegotiableQuoteDataNegotiableQuoteTotalsInterface `json:"negotiable_quote_totals,omitempty"`
	BaseCustomerBalanceAmount *float32                                           `json:"base_customer_balance_amount,omitempty"`
	CustomerBalanceAmount     *float32                                           `json:"customer_balance_amount,omitempty"`
	CouponCodes               []string                                           `json:"coupon_codes,omitempty"`
	CouponsLabels             []string                                           `json:"coupons_labels,omitempty"`
	RewardPointsBalance       *float32                                           `json:"reward_points_balance,omitempty"`
	RewardCurrencyAmount      *float32                                           `json:"reward_currency_amount,omitempty"`
	BaseRewardCurrencyAmount  *float32                                           `json:"base_reward_currency_amount,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _QuoteDataTotalsExtensionInterface QuoteDataTotalsExtensionInterface

// NewQuoteDataTotalsExtensionInterface instantiates a new QuoteDataTotalsExtensionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDataTotalsExtensionInterface() *QuoteDataTotalsExtensionInterface {
	this := QuoteDataTotalsExtensionInterface{}
	return &this
}

// NewQuoteDataTotalsExtensionInterfaceWithDefaults instantiates a new QuoteDataTotalsExtensionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDataTotalsExtensionInterfaceWithDefaults() *QuoteDataTotalsExtensionInterface {
	this := QuoteDataTotalsExtensionInterface{}
	return &this
}

// GetCouponLabel returns the CouponLabel field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetCouponLabel() string {
	if o == nil || IsNil(o.CouponLabel) {
		var ret string
		return ret
	}
	return *o.CouponLabel
}

// GetCouponLabelOk returns a tuple with the CouponLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetCouponLabelOk() (*string, bool) {
	if o == nil || IsNil(o.CouponLabel) {
		return nil, false
	}
	return o.CouponLabel, true
}

// HasCouponLabel returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasCouponLabel() bool {
	if o != nil && !IsNil(o.CouponLabel) {
		return true
	}

	return false
}

// SetCouponLabel gets a reference to the given string and assigns it to the CouponLabel field.
func (o *QuoteDataTotalsExtensionInterface) SetCouponLabel(v string) {
	o.CouponLabel = &v
}

// GetNegotiableQuoteTotals returns the NegotiableQuoteTotals field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetNegotiableQuoteTotals() NegotiableQuoteDataNegotiableQuoteTotalsInterface {
	if o == nil || IsNil(o.NegotiableQuoteTotals) {
		var ret NegotiableQuoteDataNegotiableQuoteTotalsInterface
		return ret
	}
	return *o.NegotiableQuoteTotals
}

// GetNegotiableQuoteTotalsOk returns a tuple with the NegotiableQuoteTotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetNegotiableQuoteTotalsOk() (*NegotiableQuoteDataNegotiableQuoteTotalsInterface, bool) {
	if o == nil || IsNil(o.NegotiableQuoteTotals) {
		return nil, false
	}
	return o.NegotiableQuoteTotals, true
}

// HasNegotiableQuoteTotals returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasNegotiableQuoteTotals() bool {
	if o != nil && !IsNil(o.NegotiableQuoteTotals) {
		return true
	}

	return false
}

// SetNegotiableQuoteTotals gets a reference to the given NegotiableQuoteDataNegotiableQuoteTotalsInterface and assigns it to the NegotiableQuoteTotals field.
func (o *QuoteDataTotalsExtensionInterface) SetNegotiableQuoteTotals(v NegotiableQuoteDataNegotiableQuoteTotalsInterface) {
	o.NegotiableQuoteTotals = &v
}

// GetBaseCustomerBalanceAmount returns the BaseCustomerBalanceAmount field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetBaseCustomerBalanceAmount() float32 {
	if o == nil || IsNil(o.BaseCustomerBalanceAmount) {
		var ret float32
		return ret
	}
	return *o.BaseCustomerBalanceAmount
}

// GetBaseCustomerBalanceAmountOk returns a tuple with the BaseCustomerBalanceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetBaseCustomerBalanceAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseCustomerBalanceAmount) {
		return nil, false
	}
	return o.BaseCustomerBalanceAmount, true
}

// HasBaseCustomerBalanceAmount returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasBaseCustomerBalanceAmount() bool {
	if o != nil && !IsNil(o.BaseCustomerBalanceAmount) {
		return true
	}

	return false
}

// SetBaseCustomerBalanceAmount gets a reference to the given float32 and assigns it to the BaseCustomerBalanceAmount field.
func (o *QuoteDataTotalsExtensionInterface) SetBaseCustomerBalanceAmount(v float32) {
	o.BaseCustomerBalanceAmount = &v
}

// GetCustomerBalanceAmount returns the CustomerBalanceAmount field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetCustomerBalanceAmount() float32 {
	if o == nil || IsNil(o.CustomerBalanceAmount) {
		var ret float32
		return ret
	}
	return *o.CustomerBalanceAmount
}

// GetCustomerBalanceAmountOk returns a tuple with the CustomerBalanceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetCustomerBalanceAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CustomerBalanceAmount) {
		return nil, false
	}
	return o.CustomerBalanceAmount, true
}

// HasCustomerBalanceAmount returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasCustomerBalanceAmount() bool {
	if o != nil && !IsNil(o.CustomerBalanceAmount) {
		return true
	}

	return false
}

// SetCustomerBalanceAmount gets a reference to the given float32 and assigns it to the CustomerBalanceAmount field.
func (o *QuoteDataTotalsExtensionInterface) SetCustomerBalanceAmount(v float32) {
	o.CustomerBalanceAmount = &v
}

// GetCouponCodes returns the CouponCodes field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetCouponCodes() []string {
	if o == nil || IsNil(o.CouponCodes) {
		var ret []string
		return ret
	}
	return o.CouponCodes
}

// GetCouponCodesOk returns a tuple with the CouponCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetCouponCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CouponCodes) {
		return nil, false
	}
	return o.CouponCodes, true
}

// HasCouponCodes returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasCouponCodes() bool {
	if o != nil && !IsNil(o.CouponCodes) {
		return true
	}

	return false
}

// SetCouponCodes gets a reference to the given []string and assigns it to the CouponCodes field.
func (o *QuoteDataTotalsExtensionInterface) SetCouponCodes(v []string) {
	o.CouponCodes = v
}

// GetCouponsLabels returns the CouponsLabels field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetCouponsLabels() []string {
	if o == nil || IsNil(o.CouponsLabels) {
		var ret []string
		return ret
	}
	return o.CouponsLabels
}

// GetCouponsLabelsOk returns a tuple with the CouponsLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetCouponsLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.CouponsLabels) {
		return nil, false
	}
	return o.CouponsLabels, true
}

// HasCouponsLabels returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasCouponsLabels() bool {
	if o != nil && !IsNil(o.CouponsLabels) {
		return true
	}

	return false
}

// SetCouponsLabels gets a reference to the given []string and assigns it to the CouponsLabels field.
func (o *QuoteDataTotalsExtensionInterface) SetCouponsLabels(v []string) {
	o.CouponsLabels = v
}

// GetRewardPointsBalance returns the RewardPointsBalance field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetRewardPointsBalance() float32 {
	if o == nil || IsNil(o.RewardPointsBalance) {
		var ret float32
		return ret
	}
	return *o.RewardPointsBalance
}

// GetRewardPointsBalanceOk returns a tuple with the RewardPointsBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetRewardPointsBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.RewardPointsBalance) {
		return nil, false
	}
	return o.RewardPointsBalance, true
}

// HasRewardPointsBalance returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasRewardPointsBalance() bool {
	if o != nil && !IsNil(o.RewardPointsBalance) {
		return true
	}

	return false
}

// SetRewardPointsBalance gets a reference to the given float32 and assigns it to the RewardPointsBalance field.
func (o *QuoteDataTotalsExtensionInterface) SetRewardPointsBalance(v float32) {
	o.RewardPointsBalance = &v
}

// GetRewardCurrencyAmount returns the RewardCurrencyAmount field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetRewardCurrencyAmount() float32 {
	if o == nil || IsNil(o.RewardCurrencyAmount) {
		var ret float32
		return ret
	}
	return *o.RewardCurrencyAmount
}

// GetRewardCurrencyAmountOk returns a tuple with the RewardCurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetRewardCurrencyAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RewardCurrencyAmount) {
		return nil, false
	}
	return o.RewardCurrencyAmount, true
}

// HasRewardCurrencyAmount returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasRewardCurrencyAmount() bool {
	if o != nil && !IsNil(o.RewardCurrencyAmount) {
		return true
	}

	return false
}

// SetRewardCurrencyAmount gets a reference to the given float32 and assigns it to the RewardCurrencyAmount field.
func (o *QuoteDataTotalsExtensionInterface) SetRewardCurrencyAmount(v float32) {
	o.RewardCurrencyAmount = &v
}

// GetBaseRewardCurrencyAmount returns the BaseRewardCurrencyAmount field value if set, zero value otherwise.
func (o *QuoteDataTotalsExtensionInterface) GetBaseRewardCurrencyAmount() float32 {
	if o == nil || IsNil(o.BaseRewardCurrencyAmount) {
		var ret float32
		return ret
	}
	return *o.BaseRewardCurrencyAmount
}

// GetBaseRewardCurrencyAmountOk returns a tuple with the BaseRewardCurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDataTotalsExtensionInterface) GetBaseRewardCurrencyAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseRewardCurrencyAmount) {
		return nil, false
	}
	return o.BaseRewardCurrencyAmount, true
}

// HasBaseRewardCurrencyAmount returns a boolean if a field has been set.
func (o *QuoteDataTotalsExtensionInterface) HasBaseRewardCurrencyAmount() bool {
	if o != nil && !IsNil(o.BaseRewardCurrencyAmount) {
		return true
	}

	return false
}

// SetBaseRewardCurrencyAmount gets a reference to the given float32 and assigns it to the BaseRewardCurrencyAmount field.
func (o *QuoteDataTotalsExtensionInterface) SetBaseRewardCurrencyAmount(v float32) {
	o.BaseRewardCurrencyAmount = &v
}

func (o QuoteDataTotalsExtensionInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDataTotalsExtensionInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CouponLabel) {
		toSerialize["coupon_label"] = o.CouponLabel
	}
	if !IsNil(o.NegotiableQuoteTotals) {
		toSerialize["negotiable_quote_totals"] = o.NegotiableQuoteTotals
	}
	if !IsNil(o.BaseCustomerBalanceAmount) {
		toSerialize["base_customer_balance_amount"] = o.BaseCustomerBalanceAmount
	}
	if !IsNil(o.CustomerBalanceAmount) {
		toSerialize["customer_balance_amount"] = o.CustomerBalanceAmount
	}
	if !IsNil(o.CouponCodes) {
		toSerialize["coupon_codes"] = o.CouponCodes
	}
	if !IsNil(o.CouponsLabels) {
		toSerialize["coupons_labels"] = o.CouponsLabels
	}
	if !IsNil(o.RewardPointsBalance) {
		toSerialize["reward_points_balance"] = o.RewardPointsBalance
	}
	if !IsNil(o.RewardCurrencyAmount) {
		toSerialize["reward_currency_amount"] = o.RewardCurrencyAmount
	}
	if !IsNil(o.BaseRewardCurrencyAmount) {
		toSerialize["base_reward_currency_amount"] = o.BaseRewardCurrencyAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuoteDataTotalsExtensionInterface) UnmarshalJSON(data []byte) (err error) {
	varQuoteDataTotalsExtensionInterface := _QuoteDataTotalsExtensionInterface{}

	err = json.Unmarshal(data, &varQuoteDataTotalsExtensionInterface)

	if err != nil {
		return err
	}

	*o = QuoteDataTotalsExtensionInterface(varQuoteDataTotalsExtensionInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coupon_label")
		delete(additionalProperties, "negotiable_quote_totals")
		delete(additionalProperties, "base_customer_balance_amount")
		delete(additionalProperties, "customer_balance_amount")
		delete(additionalProperties, "coupon_codes")
		delete(additionalProperties, "coupons_labels")
		delete(additionalProperties, "reward_points_balance")
		delete(additionalProperties, "reward_currency_amount")
		delete(additionalProperties, "base_reward_currency_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *QuoteDataTotalsExtensionInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *QuoteDataTotalsExtensionInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableQuoteDataTotalsExtensionInterface struct {
	value *QuoteDataTotalsExtensionInterface
	isSet bool
}

func (v NullableQuoteDataTotalsExtensionInterface) Get() *QuoteDataTotalsExtensionInterface {
	return v.value
}

func (v *NullableQuoteDataTotalsExtensionInterface) Set(val *QuoteDataTotalsExtensionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDataTotalsExtensionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDataTotalsExtensionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDataTotalsExtensionInterface(val *QuoteDataTotalsExtensionInterface) *NullableQuoteDataTotalsExtensionInterface {
	return &NullableQuoteDataTotalsExtensionInterface{value: val, isSet: true}
}

func (v NullableQuoteDataTotalsExtensionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDataTotalsExtensionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
