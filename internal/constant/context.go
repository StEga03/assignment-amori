package constant

type ContextKey string

// Context keys.
const (
	ContextKeyAPIVersion         ContextKey = "api-version"
	ContextKeyAPISource          ContextKey = "api-source"
	ContextKeyAPIModule          ContextKey = "api-module"
	ContextKeyAPISubModule       ContextKey = "api-sub-module"
	ContextKeyUsecaseKey         ContextKey = "usecase-key"
	ContextKeyLanguage           ContextKey = "context-request-lang"
	ContextKeyUserID             ContextKey = "context-user-id"
	ContextKeyUserScope          ContextKey = "context-user-scope"
	ContextKeyIdentityProviderID ContextKey = "identity-provider-id"
)
