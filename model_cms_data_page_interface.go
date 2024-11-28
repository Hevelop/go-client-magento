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

// checks if the CmsDataPageInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmsDataPageInterface{}

// CmsDataPageInterface CMS page interface.
type CmsDataPageInterface struct {
	// ID
	Id *int32 `json:"id,omitempty"`
	// Identifier
	Identifier string `json:"identifier"`
	// Title
	Title *string `json:"title,omitempty"`
	// Page layout
	PageLayout *string `json:"page_layout,omitempty"`
	// Meta title
	MetaTitle *string `json:"meta_title,omitempty"`
	// Meta keywords
	MetaKeywords *string `json:"meta_keywords,omitempty"`
	// Meta description
	MetaDescription *string `json:"meta_description,omitempty"`
	// Content heading
	ContentHeading *string `json:"content_heading,omitempty"`
	// Content
	Content *string `json:"content,omitempty"`
	// Creation time
	CreationTime *string `json:"creation_time,omitempty"`
	// Update time
	UpdateTime *string `json:"update_time,omitempty"`
	// Sort order
	SortOrder *string `json:"sort_order,omitempty"`
	// Layout update xml
	LayoutUpdateXml *string `json:"layout_update_xml,omitempty"`
	// Custom theme
	CustomTheme *string `json:"custom_theme,omitempty"`
	// Custom root template
	CustomRootTemplate *string `json:"custom_root_template,omitempty"`
	// Custom layout update xml
	CustomLayoutUpdateXml *string `json:"custom_layout_update_xml,omitempty"`
	// Custom theme from
	CustomThemeFrom *string `json:"custom_theme_from,omitempty"`
	// Custom theme to
	CustomThemeTo *string `json:"custom_theme_to,omitempty"`
	// Active
	Active               *bool `json:"active,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CmsDataPageInterface CmsDataPageInterface

// NewCmsDataPageInterface instantiates a new CmsDataPageInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmsDataPageInterface(identifier string) *CmsDataPageInterface {
	this := CmsDataPageInterface{}
	this.Identifier = identifier
	return &this
}

// NewCmsDataPageInterfaceWithDefaults instantiates a new CmsDataPageInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmsDataPageInterfaceWithDefaults() *CmsDataPageInterface {
	this := CmsDataPageInterface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CmsDataPageInterface) SetId(v int32) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value
func (o *CmsDataPageInterface) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *CmsDataPageInterface) SetIdentifier(v string) {
	o.Identifier = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CmsDataPageInterface) SetTitle(v string) {
	o.Title = &v
}

// GetPageLayout returns the PageLayout field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetPageLayout() string {
	if o == nil || IsNil(o.PageLayout) {
		var ret string
		return ret
	}
	return *o.PageLayout
}

// GetPageLayoutOk returns a tuple with the PageLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetPageLayoutOk() (*string, bool) {
	if o == nil || IsNil(o.PageLayout) {
		return nil, false
	}
	return o.PageLayout, true
}

// HasPageLayout returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasPageLayout() bool {
	if o != nil && !IsNil(o.PageLayout) {
		return true
	}

	return false
}

// SetPageLayout gets a reference to the given string and assigns it to the PageLayout field.
func (o *CmsDataPageInterface) SetPageLayout(v string) {
	o.PageLayout = &v
}

// GetMetaTitle returns the MetaTitle field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetMetaTitle() string {
	if o == nil || IsNil(o.MetaTitle) {
		var ret string
		return ret
	}
	return *o.MetaTitle
}

// GetMetaTitleOk returns a tuple with the MetaTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetMetaTitleOk() (*string, bool) {
	if o == nil || IsNil(o.MetaTitle) {
		return nil, false
	}
	return o.MetaTitle, true
}

// HasMetaTitle returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasMetaTitle() bool {
	if o != nil && !IsNil(o.MetaTitle) {
		return true
	}

	return false
}

// SetMetaTitle gets a reference to the given string and assigns it to the MetaTitle field.
func (o *CmsDataPageInterface) SetMetaTitle(v string) {
	o.MetaTitle = &v
}

// GetMetaKeywords returns the MetaKeywords field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetMetaKeywords() string {
	if o == nil || IsNil(o.MetaKeywords) {
		var ret string
		return ret
	}
	return *o.MetaKeywords
}

// GetMetaKeywordsOk returns a tuple with the MetaKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetMetaKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.MetaKeywords) {
		return nil, false
	}
	return o.MetaKeywords, true
}

// HasMetaKeywords returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasMetaKeywords() bool {
	if o != nil && !IsNil(o.MetaKeywords) {
		return true
	}

	return false
}

// SetMetaKeywords gets a reference to the given string and assigns it to the MetaKeywords field.
func (o *CmsDataPageInterface) SetMetaKeywords(v string) {
	o.MetaKeywords = &v
}

// GetMetaDescription returns the MetaDescription field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetMetaDescription() string {
	if o == nil || IsNil(o.MetaDescription) {
		var ret string
		return ret
	}
	return *o.MetaDescription
}

// GetMetaDescriptionOk returns a tuple with the MetaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetMetaDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetaDescription) {
		return nil, false
	}
	return o.MetaDescription, true
}

// HasMetaDescription returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasMetaDescription() bool {
	if o != nil && !IsNil(o.MetaDescription) {
		return true
	}

	return false
}

// SetMetaDescription gets a reference to the given string and assigns it to the MetaDescription field.
func (o *CmsDataPageInterface) SetMetaDescription(v string) {
	o.MetaDescription = &v
}

// GetContentHeading returns the ContentHeading field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetContentHeading() string {
	if o == nil || IsNil(o.ContentHeading) {
		var ret string
		return ret
	}
	return *o.ContentHeading
}

// GetContentHeadingOk returns a tuple with the ContentHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetContentHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.ContentHeading) {
		return nil, false
	}
	return o.ContentHeading, true
}

// HasContentHeading returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasContentHeading() bool {
	if o != nil && !IsNil(o.ContentHeading) {
		return true
	}

	return false
}

// SetContentHeading gets a reference to the given string and assigns it to the ContentHeading field.
func (o *CmsDataPageInterface) SetContentHeading(v string) {
	o.ContentHeading = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *CmsDataPageInterface) SetContent(v string) {
	o.Content = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime) {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetCreationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *CmsDataPageInterface) SetCreationTime(v string) {
	o.CreationTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *CmsDataPageInterface) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetSortOrder() string {
	if o == nil || IsNil(o.SortOrder) {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetSortOrderOk() (*string, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *CmsDataPageInterface) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetLayoutUpdateXml returns the LayoutUpdateXml field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetLayoutUpdateXml() string {
	if o == nil || IsNil(o.LayoutUpdateXml) {
		var ret string
		return ret
	}
	return *o.LayoutUpdateXml
}

// GetLayoutUpdateXmlOk returns a tuple with the LayoutUpdateXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetLayoutUpdateXmlOk() (*string, bool) {
	if o == nil || IsNil(o.LayoutUpdateXml) {
		return nil, false
	}
	return o.LayoutUpdateXml, true
}

// HasLayoutUpdateXml returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasLayoutUpdateXml() bool {
	if o != nil && !IsNil(o.LayoutUpdateXml) {
		return true
	}

	return false
}

// SetLayoutUpdateXml gets a reference to the given string and assigns it to the LayoutUpdateXml field.
func (o *CmsDataPageInterface) SetLayoutUpdateXml(v string) {
	o.LayoutUpdateXml = &v
}

// GetCustomTheme returns the CustomTheme field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetCustomTheme() string {
	if o == nil || IsNil(o.CustomTheme) {
		var ret string
		return ret
	}
	return *o.CustomTheme
}

// GetCustomThemeOk returns a tuple with the CustomTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetCustomThemeOk() (*string, bool) {
	if o == nil || IsNil(o.CustomTheme) {
		return nil, false
	}
	return o.CustomTheme, true
}

// HasCustomTheme returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasCustomTheme() bool {
	if o != nil && !IsNil(o.CustomTheme) {
		return true
	}

	return false
}

// SetCustomTheme gets a reference to the given string and assigns it to the CustomTheme field.
func (o *CmsDataPageInterface) SetCustomTheme(v string) {
	o.CustomTheme = &v
}

// GetCustomRootTemplate returns the CustomRootTemplate field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetCustomRootTemplate() string {
	if o == nil || IsNil(o.CustomRootTemplate) {
		var ret string
		return ret
	}
	return *o.CustomRootTemplate
}

// GetCustomRootTemplateOk returns a tuple with the CustomRootTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetCustomRootTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomRootTemplate) {
		return nil, false
	}
	return o.CustomRootTemplate, true
}

// HasCustomRootTemplate returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasCustomRootTemplate() bool {
	if o != nil && !IsNil(o.CustomRootTemplate) {
		return true
	}

	return false
}

// SetCustomRootTemplate gets a reference to the given string and assigns it to the CustomRootTemplate field.
func (o *CmsDataPageInterface) SetCustomRootTemplate(v string) {
	o.CustomRootTemplate = &v
}

// GetCustomLayoutUpdateXml returns the CustomLayoutUpdateXml field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetCustomLayoutUpdateXml() string {
	if o == nil || IsNil(o.CustomLayoutUpdateXml) {
		var ret string
		return ret
	}
	return *o.CustomLayoutUpdateXml
}

// GetCustomLayoutUpdateXmlOk returns a tuple with the CustomLayoutUpdateXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetCustomLayoutUpdateXmlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomLayoutUpdateXml) {
		return nil, false
	}
	return o.CustomLayoutUpdateXml, true
}

// HasCustomLayoutUpdateXml returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasCustomLayoutUpdateXml() bool {
	if o != nil && !IsNil(o.CustomLayoutUpdateXml) {
		return true
	}

	return false
}

// SetCustomLayoutUpdateXml gets a reference to the given string and assigns it to the CustomLayoutUpdateXml field.
func (o *CmsDataPageInterface) SetCustomLayoutUpdateXml(v string) {
	o.CustomLayoutUpdateXml = &v
}

// GetCustomThemeFrom returns the CustomThemeFrom field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetCustomThemeFrom() string {
	if o == nil || IsNil(o.CustomThemeFrom) {
		var ret string
		return ret
	}
	return *o.CustomThemeFrom
}

// GetCustomThemeFromOk returns a tuple with the CustomThemeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetCustomThemeFromOk() (*string, bool) {
	if o == nil || IsNil(o.CustomThemeFrom) {
		return nil, false
	}
	return o.CustomThemeFrom, true
}

// HasCustomThemeFrom returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasCustomThemeFrom() bool {
	if o != nil && !IsNil(o.CustomThemeFrom) {
		return true
	}

	return false
}

// SetCustomThemeFrom gets a reference to the given string and assigns it to the CustomThemeFrom field.
func (o *CmsDataPageInterface) SetCustomThemeFrom(v string) {
	o.CustomThemeFrom = &v
}

// GetCustomThemeTo returns the CustomThemeTo field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetCustomThemeTo() string {
	if o == nil || IsNil(o.CustomThemeTo) {
		var ret string
		return ret
	}
	return *o.CustomThemeTo
}

// GetCustomThemeToOk returns a tuple with the CustomThemeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetCustomThemeToOk() (*string, bool) {
	if o == nil || IsNil(o.CustomThemeTo) {
		return nil, false
	}
	return o.CustomThemeTo, true
}

// HasCustomThemeTo returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasCustomThemeTo() bool {
	if o != nil && !IsNil(o.CustomThemeTo) {
		return true
	}

	return false
}

// SetCustomThemeTo gets a reference to the given string and assigns it to the CustomThemeTo field.
func (o *CmsDataPageInterface) SetCustomThemeTo(v string) {
	o.CustomThemeTo = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CmsDataPageInterface) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsDataPageInterface) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CmsDataPageInterface) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CmsDataPageInterface) SetActive(v bool) {
	o.Active = &v
}

func (o CmsDataPageInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmsDataPageInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["identifier"] = o.Identifier
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.PageLayout) {
		toSerialize["page_layout"] = o.PageLayout
	}
	if !IsNil(o.MetaTitle) {
		toSerialize["meta_title"] = o.MetaTitle
	}
	if !IsNil(o.MetaKeywords) {
		toSerialize["meta_keywords"] = o.MetaKeywords
	}
	if !IsNil(o.MetaDescription) {
		toSerialize["meta_description"] = o.MetaDescription
	}
	if !IsNil(o.ContentHeading) {
		toSerialize["content_heading"] = o.ContentHeading
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	if !IsNil(o.LayoutUpdateXml) {
		toSerialize["layout_update_xml"] = o.LayoutUpdateXml
	}
	if !IsNil(o.CustomTheme) {
		toSerialize["custom_theme"] = o.CustomTheme
	}
	if !IsNil(o.CustomRootTemplate) {
		toSerialize["custom_root_template"] = o.CustomRootTemplate
	}
	if !IsNil(o.CustomLayoutUpdateXml) {
		toSerialize["custom_layout_update_xml"] = o.CustomLayoutUpdateXml
	}
	if !IsNil(o.CustomThemeFrom) {
		toSerialize["custom_theme_from"] = o.CustomThemeFrom
	}
	if !IsNil(o.CustomThemeTo) {
		toSerialize["custom_theme_to"] = o.CustomThemeTo
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CmsDataPageInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
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

	varCmsDataPageInterface := _CmsDataPageInterface{}

	err = json.Unmarshal(data, &varCmsDataPageInterface)

	if err != nil {
		return err
	}

	*o = CmsDataPageInterface(varCmsDataPageInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "title")
		delete(additionalProperties, "page_layout")
		delete(additionalProperties, "meta_title")
		delete(additionalProperties, "meta_keywords")
		delete(additionalProperties, "meta_description")
		delete(additionalProperties, "content_heading")
		delete(additionalProperties, "content")
		delete(additionalProperties, "creation_time")
		delete(additionalProperties, "update_time")
		delete(additionalProperties, "sort_order")
		delete(additionalProperties, "layout_update_xml")
		delete(additionalProperties, "custom_theme")
		delete(additionalProperties, "custom_root_template")
		delete(additionalProperties, "custom_layout_update_xml")
		delete(additionalProperties, "custom_theme_from")
		delete(additionalProperties, "custom_theme_to")
		delete(additionalProperties, "active")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CmsDataPageInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CmsDataPageInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCmsDataPageInterface struct {
	value *CmsDataPageInterface
	isSet bool
}

func (v NullableCmsDataPageInterface) Get() *CmsDataPageInterface {
	return v.value
}

func (v *NullableCmsDataPageInterface) Set(val *CmsDataPageInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCmsDataPageInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCmsDataPageInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmsDataPageInterface(val *CmsDataPageInterface) *NullableCmsDataPageInterface {
	return &NullableCmsDataPageInterface{value: val, isSet: true}
}

func (v NullableCmsDataPageInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmsDataPageInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
