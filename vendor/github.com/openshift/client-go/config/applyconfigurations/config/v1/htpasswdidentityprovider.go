// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// HTPasswdIdentityProviderApplyConfiguration represents a declarative configuration of the HTPasswdIdentityProvider type for use
// with apply.
type HTPasswdIdentityProviderApplyConfiguration struct {
	FileData *SecretNameReferenceApplyConfiguration `json:"fileData,omitempty"`
}

// HTPasswdIdentityProviderApplyConfiguration constructs a declarative configuration of the HTPasswdIdentityProvider type for use with
// apply.
func HTPasswdIdentityProvider() *HTPasswdIdentityProviderApplyConfiguration {
	return &HTPasswdIdentityProviderApplyConfiguration{}
}

// WithFileData sets the FileData field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FileData field is set to the value of the last call.
func (b *HTPasswdIdentityProviderApplyConfiguration) WithFileData(value *SecretNameReferenceApplyConfiguration) *HTPasswdIdentityProviderApplyConfiguration {
	b.FileData = value
	return b
}
