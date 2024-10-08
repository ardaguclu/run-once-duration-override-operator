// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// DeveloperConsoleCatalogCustomizationApplyConfiguration represents a declarative configuration of the DeveloperConsoleCatalogCustomization type for use
// with apply.
type DeveloperConsoleCatalogCustomizationApplyConfiguration struct {
	Categories []DeveloperConsoleCatalogCategoryApplyConfiguration `json:"categories,omitempty"`
	Types      *DeveloperConsoleCatalogTypesApplyConfiguration     `json:"types,omitempty"`
}

// DeveloperConsoleCatalogCustomizationApplyConfiguration constructs a declarative configuration of the DeveloperConsoleCatalogCustomization type for use with
// apply.
func DeveloperConsoleCatalogCustomization() *DeveloperConsoleCatalogCustomizationApplyConfiguration {
	return &DeveloperConsoleCatalogCustomizationApplyConfiguration{}
}

// WithCategories adds the given value to the Categories field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Categories field.
func (b *DeveloperConsoleCatalogCustomizationApplyConfiguration) WithCategories(values ...*DeveloperConsoleCatalogCategoryApplyConfiguration) *DeveloperConsoleCatalogCustomizationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCategories")
		}
		b.Categories = append(b.Categories, *values[i])
	}
	return b
}

// WithTypes sets the Types field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Types field is set to the value of the last call.
func (b *DeveloperConsoleCatalogCustomizationApplyConfiguration) WithTypes(value *DeveloperConsoleCatalogTypesApplyConfiguration) *DeveloperConsoleCatalogCustomizationApplyConfiguration {
	b.Types = value
	return b
}
