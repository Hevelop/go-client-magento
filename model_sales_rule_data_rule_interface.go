/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
	"fmt"
)

// checks if the SalesRuleDataRuleInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesRuleDataRuleInterface{}

// SalesRuleDataRuleInterface Interface RuleInterface
type SalesRuleDataRuleInterface struct {
	// Rule id
	RuleId *int32 `json:"rule_id,omitempty"`
	// Rule name
	Name *string `json:"name,omitempty"`
	// Display label
	StoreLabels []SalesRuleDataRuleLabelInterface `json:"store_labels,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// A list of websites the rule applies to
	WebsiteIds []int32 `json:"website_ids"`
	// Ids of customer groups that the rule applies to
	CustomerGroupIds []int32 `json:"customer_group_ids"`
	// The start date when the coupon is active
	FromDate *string `json:"from_date,omitempty"`
	// The end date when the coupon is active
	ToDate *string `json:"to_date,omitempty"`
	// Number of uses per customer
	UsesPerCustomer int32 `json:"uses_per_customer"`
	// The coupon is active
	IsActive        bool                             `json:"is_active"`
	Condition       *SalesRuleDataConditionInterface `json:"condition,omitempty"`
	ActionCondition *SalesRuleDataConditionInterface `json:"action_condition,omitempty"`
	// To stop rule processing
	StopRulesProcessing bool `json:"stop_rules_processing"`
	// Is this field needed
	IsAdvanced bool `json:"is_advanced"`
	// Product ids
	ProductIds []int32 `json:"product_ids,omitempty"`
	// Sort order
	SortOrder int32 `json:"sort_order"`
	// Simple action of the rule
	SimpleAction *string `json:"simple_action,omitempty"`
	// Discount amount
	DiscountAmount float32 `json:"discount_amount"`
	// Maximum qty discount is applied
	DiscountQty *float32 `json:"discount_qty,omitempty"`
	// Discount step
	DiscountStep int32 `json:"discount_step"`
	// The rule applies to shipping
	ApplyToShipping bool `json:"apply_to_shipping"`
	// How many times the rule has been used
	TimesUsed int32 `json:"times_used"`
	// Whether the rule is in RSS
	IsRss bool `json:"is_rss"`
	// Coupon type
	CouponType string `json:"coupon_type"`
	// To auto generate coupon
	UseAutoGeneration bool `json:"use_auto_generation"`
	// Limit of uses per coupon
	UsesPerCoupon int32 `json:"uses_per_coupon"`
	// To grant free shipping
	SimpleFreeShipping   *string                              `json:"simple_free_shipping,omitempty"`
	ExtensionAttributes  *SalesRuleDataRuleExtensionInterface `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesRuleDataRuleInterface SalesRuleDataRuleInterface

// NewSalesRuleDataRuleInterface instantiates a new SalesRuleDataRuleInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesRuleDataRuleInterface(websiteIds []int32, customerGroupIds []int32, usesPerCustomer int32, isActive bool, stopRulesProcessing bool, isAdvanced bool, sortOrder int32, discountAmount float32, discountStep int32, applyToShipping bool, timesUsed int32, isRss bool, couponType string, useAutoGeneration bool, usesPerCoupon int32) *SalesRuleDataRuleInterface {
	this := SalesRuleDataRuleInterface{}
	this.WebsiteIds = websiteIds
	this.CustomerGroupIds = customerGroupIds
	this.UsesPerCustomer = usesPerCustomer
	this.IsActive = isActive
	this.StopRulesProcessing = stopRulesProcessing
	this.IsAdvanced = isAdvanced
	this.SortOrder = sortOrder
	this.DiscountAmount = discountAmount
	this.DiscountStep = discountStep
	this.ApplyToShipping = applyToShipping
	this.TimesUsed = timesUsed
	this.IsRss = isRss
	this.CouponType = couponType
	this.UseAutoGeneration = useAutoGeneration
	this.UsesPerCoupon = usesPerCoupon
	return &this
}

// NewSalesRuleDataRuleInterfaceWithDefaults instantiates a new SalesRuleDataRuleInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesRuleDataRuleInterfaceWithDefaults() *SalesRuleDataRuleInterface {
	this := SalesRuleDataRuleInterface{}
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetRuleId() int32 {
	if o == nil || IsNil(o.RuleId) {
		var ret int32
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetRuleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given int32 and assigns it to the RuleId field.
func (o *SalesRuleDataRuleInterface) SetRuleId(v int32) {
	o.RuleId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SalesRuleDataRuleInterface) SetName(v string) {
	o.Name = &v
}

// GetStoreLabels returns the StoreLabels field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetStoreLabels() []SalesRuleDataRuleLabelInterface {
	if o == nil || IsNil(o.StoreLabels) {
		var ret []SalesRuleDataRuleLabelInterface
		return ret
	}
	return o.StoreLabels
}

// GetStoreLabelsOk returns a tuple with the StoreLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetStoreLabelsOk() ([]SalesRuleDataRuleLabelInterface, bool) {
	if o == nil || IsNil(o.StoreLabels) {
		return nil, false
	}
	return o.StoreLabels, true
}

// HasStoreLabels returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasStoreLabels() bool {
	if o != nil && !IsNil(o.StoreLabels) {
		return true
	}

	return false
}

// SetStoreLabels gets a reference to the given []SalesRuleDataRuleLabelInterface and assigns it to the StoreLabels field.
func (o *SalesRuleDataRuleInterface) SetStoreLabels(v []SalesRuleDataRuleLabelInterface) {
	o.StoreLabels = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SalesRuleDataRuleInterface) SetDescription(v string) {
	o.Description = &v
}

// GetWebsiteIds returns the WebsiteIds field value
func (o *SalesRuleDataRuleInterface) GetWebsiteIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.WebsiteIds
}

// GetWebsiteIdsOk returns a tuple with the WebsiteIds field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetWebsiteIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteIds, true
}

// SetWebsiteIds sets field value
func (o *SalesRuleDataRuleInterface) SetWebsiteIds(v []int32) {
	o.WebsiteIds = v
}

// GetCustomerGroupIds returns the CustomerGroupIds field value
func (o *SalesRuleDataRuleInterface) GetCustomerGroupIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.CustomerGroupIds
}

// GetCustomerGroupIdsOk returns a tuple with the CustomerGroupIds field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetCustomerGroupIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerGroupIds, true
}

// SetCustomerGroupIds sets field value
func (o *SalesRuleDataRuleInterface) SetCustomerGroupIds(v []int32) {
	o.CustomerGroupIds = v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetFromDate() string {
	if o == nil || IsNil(o.FromDate) {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetFromDateOk() (*string, bool) {
	if o == nil || IsNil(o.FromDate) {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasFromDate() bool {
	if o != nil && !IsNil(o.FromDate) {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *SalesRuleDataRuleInterface) SetFromDate(v string) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetToDate() string {
	if o == nil || IsNil(o.ToDate) {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetToDateOk() (*string, bool) {
	if o == nil || IsNil(o.ToDate) {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasToDate() bool {
	if o != nil && !IsNil(o.ToDate) {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *SalesRuleDataRuleInterface) SetToDate(v string) {
	o.ToDate = &v
}

// GetUsesPerCustomer returns the UsesPerCustomer field value
func (o *SalesRuleDataRuleInterface) GetUsesPerCustomer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsesPerCustomer
}

// GetUsesPerCustomerOk returns a tuple with the UsesPerCustomer field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetUsesPerCustomerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsesPerCustomer, true
}

// SetUsesPerCustomer sets field value
func (o *SalesRuleDataRuleInterface) SetUsesPerCustomer(v int32) {
	o.UsesPerCustomer = v
}

// GetIsActive returns the IsActive field value
func (o *SalesRuleDataRuleInterface) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *SalesRuleDataRuleInterface) SetIsActive(v bool) {
	o.IsActive = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetCondition() SalesRuleDataConditionInterface {
	if o == nil || IsNil(o.Condition) {
		var ret SalesRuleDataConditionInterface
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetConditionOk() (*SalesRuleDataConditionInterface, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given SalesRuleDataConditionInterface and assigns it to the Condition field.
func (o *SalesRuleDataRuleInterface) SetCondition(v SalesRuleDataConditionInterface) {
	o.Condition = &v
}

// GetActionCondition returns the ActionCondition field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetActionCondition() SalesRuleDataConditionInterface {
	if o == nil || IsNil(o.ActionCondition) {
		var ret SalesRuleDataConditionInterface
		return ret
	}
	return *o.ActionCondition
}

// GetActionConditionOk returns a tuple with the ActionCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetActionConditionOk() (*SalesRuleDataConditionInterface, bool) {
	if o == nil || IsNil(o.ActionCondition) {
		return nil, false
	}
	return o.ActionCondition, true
}

// HasActionCondition returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasActionCondition() bool {
	if o != nil && !IsNil(o.ActionCondition) {
		return true
	}

	return false
}

// SetActionCondition gets a reference to the given SalesRuleDataConditionInterface and assigns it to the ActionCondition field.
func (o *SalesRuleDataRuleInterface) SetActionCondition(v SalesRuleDataConditionInterface) {
	o.ActionCondition = &v
}

// GetStopRulesProcessing returns the StopRulesProcessing field value
func (o *SalesRuleDataRuleInterface) GetStopRulesProcessing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StopRulesProcessing
}

// GetStopRulesProcessingOk returns a tuple with the StopRulesProcessing field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetStopRulesProcessingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopRulesProcessing, true
}

// SetStopRulesProcessing sets field value
func (o *SalesRuleDataRuleInterface) SetStopRulesProcessing(v bool) {
	o.StopRulesProcessing = v
}

// GetIsAdvanced returns the IsAdvanced field value
func (o *SalesRuleDataRuleInterface) GetIsAdvanced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAdvanced
}

// GetIsAdvancedOk returns a tuple with the IsAdvanced field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetIsAdvancedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAdvanced, true
}

// SetIsAdvanced sets field value
func (o *SalesRuleDataRuleInterface) SetIsAdvanced(v bool) {
	o.IsAdvanced = v
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetProductIds() []int32 {
	if o == nil || IsNil(o.ProductIds) {
		var ret []int32
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetProductIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []int32 and assigns it to the ProductIds field.
func (o *SalesRuleDataRuleInterface) SetProductIds(v []int32) {
	o.ProductIds = v
}

// GetSortOrder returns the SortOrder field value
func (o *SalesRuleDataRuleInterface) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *SalesRuleDataRuleInterface) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetSimpleAction returns the SimpleAction field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetSimpleAction() string {
	if o == nil || IsNil(o.SimpleAction) {
		var ret string
		return ret
	}
	return *o.SimpleAction
}

// GetSimpleActionOk returns a tuple with the SimpleAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetSimpleActionOk() (*string, bool) {
	if o == nil || IsNil(o.SimpleAction) {
		return nil, false
	}
	return o.SimpleAction, true
}

// HasSimpleAction returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasSimpleAction() bool {
	if o != nil && !IsNil(o.SimpleAction) {
		return true
	}

	return false
}

// SetSimpleAction gets a reference to the given string and assigns it to the SimpleAction field.
func (o *SalesRuleDataRuleInterface) SetSimpleAction(v string) {
	o.SimpleAction = &v
}

// GetDiscountAmount returns the DiscountAmount field value
func (o *SalesRuleDataRuleInterface) GetDiscountAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetDiscountAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountAmount, true
}

// SetDiscountAmount sets field value
func (o *SalesRuleDataRuleInterface) SetDiscountAmount(v float32) {
	o.DiscountAmount = v
}

// GetDiscountQty returns the DiscountQty field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetDiscountQty() float32 {
	if o == nil || IsNil(o.DiscountQty) {
		var ret float32
		return ret
	}
	return *o.DiscountQty
}

// GetDiscountQtyOk returns a tuple with the DiscountQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetDiscountQtyOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountQty) {
		return nil, false
	}
	return o.DiscountQty, true
}

// HasDiscountQty returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasDiscountQty() bool {
	if o != nil && !IsNil(o.DiscountQty) {
		return true
	}

	return false
}

// SetDiscountQty gets a reference to the given float32 and assigns it to the DiscountQty field.
func (o *SalesRuleDataRuleInterface) SetDiscountQty(v float32) {
	o.DiscountQty = &v
}

// GetDiscountStep returns the DiscountStep field value
func (o *SalesRuleDataRuleInterface) GetDiscountStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiscountStep
}

// GetDiscountStepOk returns a tuple with the DiscountStep field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetDiscountStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountStep, true
}

// SetDiscountStep sets field value
func (o *SalesRuleDataRuleInterface) SetDiscountStep(v int32) {
	o.DiscountStep = v
}

// GetApplyToShipping returns the ApplyToShipping field value
func (o *SalesRuleDataRuleInterface) GetApplyToShipping() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ApplyToShipping
}

// GetApplyToShippingOk returns a tuple with the ApplyToShipping field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetApplyToShippingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyToShipping, true
}

// SetApplyToShipping sets field value
func (o *SalesRuleDataRuleInterface) SetApplyToShipping(v bool) {
	o.ApplyToShipping = v
}

// GetTimesUsed returns the TimesUsed field value
func (o *SalesRuleDataRuleInterface) GetTimesUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimesUsed
}

// GetTimesUsedOk returns a tuple with the TimesUsed field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetTimesUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimesUsed, true
}

// SetTimesUsed sets field value
func (o *SalesRuleDataRuleInterface) SetTimesUsed(v int32) {
	o.TimesUsed = v
}

// GetIsRss returns the IsRss field value
func (o *SalesRuleDataRuleInterface) GetIsRss() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRss
}

// GetIsRssOk returns a tuple with the IsRss field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetIsRssOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRss, true
}

// SetIsRss sets field value
func (o *SalesRuleDataRuleInterface) SetIsRss(v bool) {
	o.IsRss = v
}

// GetCouponType returns the CouponType field value
func (o *SalesRuleDataRuleInterface) GetCouponType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponType
}

// GetCouponTypeOk returns a tuple with the CouponType field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetCouponTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponType, true
}

// SetCouponType sets field value
func (o *SalesRuleDataRuleInterface) SetCouponType(v string) {
	o.CouponType = v
}

// GetUseAutoGeneration returns the UseAutoGeneration field value
func (o *SalesRuleDataRuleInterface) GetUseAutoGeneration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseAutoGeneration
}

// GetUseAutoGenerationOk returns a tuple with the UseAutoGeneration field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetUseAutoGenerationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseAutoGeneration, true
}

// SetUseAutoGeneration sets field value
func (o *SalesRuleDataRuleInterface) SetUseAutoGeneration(v bool) {
	o.UseAutoGeneration = v
}

// GetUsesPerCoupon returns the UsesPerCoupon field value
func (o *SalesRuleDataRuleInterface) GetUsesPerCoupon() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsesPerCoupon
}

// GetUsesPerCouponOk returns a tuple with the UsesPerCoupon field value
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetUsesPerCouponOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsesPerCoupon, true
}

// SetUsesPerCoupon sets field value
func (o *SalesRuleDataRuleInterface) SetUsesPerCoupon(v int32) {
	o.UsesPerCoupon = v
}

// GetSimpleFreeShipping returns the SimpleFreeShipping field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetSimpleFreeShipping() string {
	if o == nil || IsNil(o.SimpleFreeShipping) {
		var ret string
		return ret
	}
	return *o.SimpleFreeShipping
}

// GetSimpleFreeShippingOk returns a tuple with the SimpleFreeShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetSimpleFreeShippingOk() (*string, bool) {
	if o == nil || IsNil(o.SimpleFreeShipping) {
		return nil, false
	}
	return o.SimpleFreeShipping, true
}

// HasSimpleFreeShipping returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasSimpleFreeShipping() bool {
	if o != nil && !IsNil(o.SimpleFreeShipping) {
		return true
	}

	return false
}

// SetSimpleFreeShipping gets a reference to the given string and assigns it to the SimpleFreeShipping field.
func (o *SalesRuleDataRuleInterface) SetSimpleFreeShipping(v string) {
	o.SimpleFreeShipping = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesRuleDataRuleInterface) GetExtensionAttributes() SalesRuleDataRuleExtensionInterface {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret SalesRuleDataRuleExtensionInterface
		return ret
	}
	return *o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesRuleDataRuleInterface) GetExtensionAttributesOk() (*SalesRuleDataRuleExtensionInterface, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesRuleDataRuleInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given SalesRuleDataRuleExtensionInterface and assigns it to the ExtensionAttributes field.
func (o *SalesRuleDataRuleInterface) SetExtensionAttributes(v SalesRuleDataRuleExtensionInterface) {
	o.ExtensionAttributes = &v
}

func (o SalesRuleDataRuleInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesRuleDataRuleInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StoreLabels) {
		toSerialize["store_labels"] = o.StoreLabels
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["website_ids"] = o.WebsiteIds
	toSerialize["customer_group_ids"] = o.CustomerGroupIds
	if !IsNil(o.FromDate) {
		toSerialize["from_date"] = o.FromDate
	}
	if !IsNil(o.ToDate) {
		toSerialize["to_date"] = o.ToDate
	}
	toSerialize["uses_per_customer"] = o.UsesPerCustomer
	toSerialize["is_active"] = o.IsActive
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.ActionCondition) {
		toSerialize["action_condition"] = o.ActionCondition
	}
	toSerialize["stop_rules_processing"] = o.StopRulesProcessing
	toSerialize["is_advanced"] = o.IsAdvanced
	if !IsNil(o.ProductIds) {
		toSerialize["product_ids"] = o.ProductIds
	}
	toSerialize["sort_order"] = o.SortOrder
	if !IsNil(o.SimpleAction) {
		toSerialize["simple_action"] = o.SimpleAction
	}
	toSerialize["discount_amount"] = o.DiscountAmount
	if !IsNil(o.DiscountQty) {
		toSerialize["discount_qty"] = o.DiscountQty
	}
	toSerialize["discount_step"] = o.DiscountStep
	toSerialize["apply_to_shipping"] = o.ApplyToShipping
	toSerialize["times_used"] = o.TimesUsed
	toSerialize["is_rss"] = o.IsRss
	toSerialize["coupon_type"] = o.CouponType
	toSerialize["use_auto_generation"] = o.UseAutoGeneration
	toSerialize["uses_per_coupon"] = o.UsesPerCoupon
	if !IsNil(o.SimpleFreeShipping) {
		toSerialize["simple_free_shipping"] = o.SimpleFreeShipping
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesRuleDataRuleInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"website_ids",
		"customer_group_ids",
		"uses_per_customer",
		"is_active",
		"stop_rules_processing",
		"is_advanced",
		"sort_order",
		"discount_amount",
		"discount_step",
		"apply_to_shipping",
		"times_used",
		"is_rss",
		"coupon_type",
		"use_auto_generation",
		"uses_per_coupon",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSalesRuleDataRuleInterface := _SalesRuleDataRuleInterface{}

	err = json.Unmarshal(data, &varSalesRuleDataRuleInterface)

	if err != nil {
		return err
	}

	*o = SalesRuleDataRuleInterface(varSalesRuleDataRuleInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rule_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "store_labels")
		delete(additionalProperties, "description")
		delete(additionalProperties, "website_ids")
		delete(additionalProperties, "customer_group_ids")
		delete(additionalProperties, "from_date")
		delete(additionalProperties, "to_date")
		delete(additionalProperties, "uses_per_customer")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "action_condition")
		delete(additionalProperties, "stop_rules_processing")
		delete(additionalProperties, "is_advanced")
		delete(additionalProperties, "product_ids")
		delete(additionalProperties, "sort_order")
		delete(additionalProperties, "simple_action")
		delete(additionalProperties, "discount_amount")
		delete(additionalProperties, "discount_qty")
		delete(additionalProperties, "discount_step")
		delete(additionalProperties, "apply_to_shipping")
		delete(additionalProperties, "times_used")
		delete(additionalProperties, "is_rss")
		delete(additionalProperties, "coupon_type")
		delete(additionalProperties, "use_auto_generation")
		delete(additionalProperties, "uses_per_coupon")
		delete(additionalProperties, "simple_free_shipping")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesRuleDataRuleInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesRuleDataRuleInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesRuleDataRuleInterface struct {
	value *SalesRuleDataRuleInterface
	isSet bool
}

func (v NullableSalesRuleDataRuleInterface) Get() *SalesRuleDataRuleInterface {
	return v.value
}

func (v *NullableSalesRuleDataRuleInterface) Set(val *SalesRuleDataRuleInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesRuleDataRuleInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesRuleDataRuleInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesRuleDataRuleInterface(val *SalesRuleDataRuleInterface) *NullableSalesRuleDataRuleInterface {
	return &NullableSalesRuleDataRuleInterface{value: val, isSet: true}
}

func (v NullableSalesRuleDataRuleInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesRuleDataRuleInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
