// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing SAML Identity Providers within Keycloak.
//
// SAML (Security Assertion Markup Language) identity providers allows users to authenticate through a third-party system using the SAML protocol.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/saml"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = saml.NewIdentityProvider(ctx, "realmSamlIdentityProvider", &saml.IdentityProviderArgs{
// 			Realm:                   realm.ID(),
// 			Alias:                   pulumi.String("my-saml-idp"),
// 			SingleSignOnServiceUrl:  pulumi.String("https://domain.com/adfs/ls/"),
// 			SingleLogoutServiceUrl:  pulumi.String("https://domain.com/adfs/ls/?wa=wsignout1.0"),
// 			BackchannelSupported:    pulumi.Bool(true),
// 			PostBindingResponse:     pulumi.Bool(true),
// 			PostBindingLogout:       pulumi.Bool(true),
// 			PostBindingAuthnRequest: pulumi.Bool(true),
// 			StoreToken:              pulumi.Bool(false),
// 			TrustEmail:              pulumi.Bool(true),
// 			ForceAuthn:              pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Identity providers can be imported using the format `{{realm_id}}/{{idp_alias}}`, where `idp_alias` is the identity provider alias. Examplebash
//
// ```sh
//  $ pulumi import keycloak:saml/identityProvider:IdentityProvider realm_saml_identity_provider my-realm/my-saml-idp
// ```
type IdentityProvider struct {
	pulumi.CustomResourceState

	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrOutput `pulumi:"addReadTokenRoleOnCreate"`
	// The unique name of identity provider.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Authenticate users by default. Defaults to `false`.
	AuthenticateByDefault pulumi.BoolPtrOutput `pulumi:"authenticateByDefault"`
	// Does the external IDP support back-channel logout ?.
	BackchannelSupported pulumi.BoolPtrOutput `pulumi:"backchannelSupported"`
	// The display name for the realm that is shown when logging in to the admin console.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"firstBrokerLoginFlowAlias"`
	// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
	ForceAuthn pulumi.BoolPtrOutput `pulumi:"forceAuthn"`
	// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
	HideOnLoginPage pulumi.BoolPtrOutput `pulumi:"hideOnLoginPage"`
	// Internal Identity Provider Id
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrOutput `pulumi:"linkOnly"`
	// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
	NameIdPolicyFormat pulumi.StringPtrOutput `pulumi:"nameIdPolicyFormat"`
	// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingAuthnRequest pulumi.BoolPtrOutput `pulumi:"postBindingAuthnRequest"`
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingLogout pulumi.BoolPtrOutput `pulumi:"postBindingLogout"`
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
	PostBindingResponse pulumi.BoolPtrOutput `pulumi:"postBindingResponse"`
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
	PostBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"postBrokerLoginFlowAlias"`
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Signing Algorithm. Defaults to empty.
	SignatureAlgorithm pulumi.StringPtrOutput `pulumi:"signatureAlgorithm"`
	// Signing Certificate.
	SigningCertificate pulumi.StringPtrOutput `pulumi:"signingCertificate"`
	// The Url that must be used to send logout requests.
	SingleLogoutServiceUrl pulumi.StringPtrOutput `pulumi:"singleLogoutServiceUrl"`
	// The Url that must be used to send authentication requests (SAML AuthnRequest).
	SingleSignOnServiceUrl pulumi.StringOutput `pulumi:"singleSignOnServiceUrl"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrOutput `pulumi:"storeToken"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrOutput `pulumi:"trustEmail"`
	// Enable/disable signature validation of SAML responses.
	ValidateSignature pulumi.BoolPtrOutput `pulumi:"validateSignature"`
	// Indicates whether this service provider expects an encrypted Assertion.
	WantAssertionsEncrypted pulumi.BoolPtrOutput `pulumi:"wantAssertionsEncrypted"`
	// Indicates whether this service provider expects a signed Assertion.
	WantAssertionsSigned pulumi.BoolPtrOutput `pulumi:"wantAssertionsSigned"`
	// Sign Key Transformer. Defaults to empty.
	XmlSignKeyInfoKeyNameTransformer pulumi.StringPtrOutput `pulumi:"xmlSignKeyInfoKeyNameTransformer"`
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	if args.SingleSignOnServiceUrl == nil {
		return nil, errors.New("invalid value for required argument 'SingleSignOnServiceUrl'")
	}
	var resource IdentityProvider
	err := ctx.RegisterResource("keycloak:saml/identityProvider:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("keycloak:saml/identityProvider:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// The unique name of identity provider.
	Alias *string `pulumi:"alias"`
	// Authenticate users by default. Defaults to `false`.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// Does the external IDP support back-channel logout ?.
	BackchannelSupported *bool `pulumi:"backchannelSupported"`
	// The display name for the realm that is shown when logging in to the admin console.
	DisplayName *string `pulumi:"displayName"`
	// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
	ForceAuthn *bool `pulumi:"forceAuthn"`
	// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// Internal Identity Provider Id
	InternalId *string `pulumi:"internalId"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly *bool `pulumi:"linkOnly"`
	// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
	NameIdPolicyFormat *string `pulumi:"nameIdPolicyFormat"`
	// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingAuthnRequest *bool `pulumi:"postBindingAuthnRequest"`
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingLogout *bool `pulumi:"postBindingLogout"`
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
	PostBindingResponse *bool `pulumi:"postBindingResponse"`
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// The name of the realm. This is unique across Keycloak.
	Realm *string `pulumi:"realm"`
	// Signing Algorithm. Defaults to empty.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// Signing Certificate.
	SigningCertificate *string `pulumi:"signingCertificate"`
	// The Url that must be used to send logout requests.
	SingleLogoutServiceUrl *string `pulumi:"singleLogoutServiceUrl"`
	// The Url that must be used to send authentication requests (SAML AuthnRequest).
	SingleSignOnServiceUrl *string `pulumi:"singleSignOnServiceUrl"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken *bool `pulumi:"storeToken"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Enable/disable signature validation of SAML responses.
	ValidateSignature *bool `pulumi:"validateSignature"`
	// Indicates whether this service provider expects an encrypted Assertion.
	WantAssertionsEncrypted *bool `pulumi:"wantAssertionsEncrypted"`
	// Indicates whether this service provider expects a signed Assertion.
	WantAssertionsSigned *bool `pulumi:"wantAssertionsSigned"`
	// Sign Key Transformer. Defaults to empty.
	XmlSignKeyInfoKeyNameTransformer *string `pulumi:"xmlSignKeyInfoKeyNameTransformer"`
}

type IdentityProviderState struct {
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// The unique name of identity provider.
	Alias pulumi.StringPtrInput
	// Authenticate users by default. Defaults to `false`.
	AuthenticateByDefault pulumi.BoolPtrInput
	// Does the external IDP support back-channel logout ?.
	BackchannelSupported pulumi.BoolPtrInput
	// The display name for the realm that is shown when logging in to the admin console.
	DisplayName pulumi.StringPtrInput
	// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
	ForceAuthn pulumi.BoolPtrInput
	// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
	HideOnLoginPage pulumi.BoolPtrInput
	// Internal Identity Provider Id
	InternalId pulumi.StringPtrInput
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrInput
	// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
	NameIdPolicyFormat pulumi.StringPtrInput
	// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingAuthnRequest pulumi.BoolPtrInput
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingLogout pulumi.BoolPtrInput
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
	PostBindingResponse pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringPtrInput
	// Signing Algorithm. Defaults to empty.
	SignatureAlgorithm pulumi.StringPtrInput
	// Signing Certificate.
	SigningCertificate pulumi.StringPtrInput
	// The Url that must be used to send logout requests.
	SingleLogoutServiceUrl pulumi.StringPtrInput
	// The Url that must be used to send authentication requests (SAML AuthnRequest).
	SingleSignOnServiceUrl pulumi.StringPtrInput
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrInput
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrInput
	// Enable/disable signature validation of SAML responses.
	ValidateSignature pulumi.BoolPtrInput
	// Indicates whether this service provider expects an encrypted Assertion.
	WantAssertionsEncrypted pulumi.BoolPtrInput
	// Indicates whether this service provider expects a signed Assertion.
	WantAssertionsSigned pulumi.BoolPtrInput
	// Sign Key Transformer. Defaults to empty.
	XmlSignKeyInfoKeyNameTransformer pulumi.StringPtrInput
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// The unique name of identity provider.
	Alias string `pulumi:"alias"`
	// Authenticate users by default. Defaults to `false`.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// Does the external IDP support back-channel logout ?.
	BackchannelSupported *bool `pulumi:"backchannelSupported"`
	// The display name for the realm that is shown when logging in to the admin console.
	DisplayName *string `pulumi:"displayName"`
	// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
	ForceAuthn *bool `pulumi:"forceAuthn"`
	// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly *bool `pulumi:"linkOnly"`
	// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
	NameIdPolicyFormat *string `pulumi:"nameIdPolicyFormat"`
	// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingAuthnRequest *bool `pulumi:"postBindingAuthnRequest"`
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingLogout *bool `pulumi:"postBindingLogout"`
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
	PostBindingResponse *bool `pulumi:"postBindingResponse"`
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// The name of the realm. This is unique across Keycloak.
	Realm string `pulumi:"realm"`
	// Signing Algorithm. Defaults to empty.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// Signing Certificate.
	SigningCertificate *string `pulumi:"signingCertificate"`
	// The Url that must be used to send logout requests.
	SingleLogoutServiceUrl *string `pulumi:"singleLogoutServiceUrl"`
	// The Url that must be used to send authentication requests (SAML AuthnRequest).
	SingleSignOnServiceUrl string `pulumi:"singleSignOnServiceUrl"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken *bool `pulumi:"storeToken"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Enable/disable signature validation of SAML responses.
	ValidateSignature *bool `pulumi:"validateSignature"`
	// Indicates whether this service provider expects an encrypted Assertion.
	WantAssertionsEncrypted *bool `pulumi:"wantAssertionsEncrypted"`
	// Indicates whether this service provider expects a signed Assertion.
	WantAssertionsSigned *bool `pulumi:"wantAssertionsSigned"`
	// Sign Key Transformer. Defaults to empty.
	XmlSignKeyInfoKeyNameTransformer *string `pulumi:"xmlSignKeyInfoKeyNameTransformer"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// The unique name of identity provider.
	Alias pulumi.StringInput
	// Authenticate users by default. Defaults to `false`.
	AuthenticateByDefault pulumi.BoolPtrInput
	// Does the external IDP support back-channel logout ?.
	BackchannelSupported pulumi.BoolPtrInput
	// The display name for the realm that is shown when logging in to the admin console.
	DisplayName pulumi.StringPtrInput
	// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
	ForceAuthn pulumi.BoolPtrInput
	// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
	HideOnLoginPage pulumi.BoolPtrInput
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrInput
	// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
	NameIdPolicyFormat pulumi.StringPtrInput
	// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingAuthnRequest pulumi.BoolPtrInput
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
	PostBindingLogout pulumi.BoolPtrInput
	// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
	PostBindingResponse pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringInput
	// Signing Algorithm. Defaults to empty.
	SignatureAlgorithm pulumi.StringPtrInput
	// Signing Certificate.
	SigningCertificate pulumi.StringPtrInput
	// The Url that must be used to send logout requests.
	SingleLogoutServiceUrl pulumi.StringPtrInput
	// The Url that must be used to send authentication requests (SAML AuthnRequest).
	SingleSignOnServiceUrl pulumi.StringInput
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrInput
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrInput
	// Enable/disable signature validation of SAML responses.
	ValidateSignature pulumi.BoolPtrInput
	// Indicates whether this service provider expects an encrypted Assertion.
	WantAssertionsEncrypted pulumi.BoolPtrInput
	// Indicates whether this service provider expects a signed Assertion.
	WantAssertionsSigned pulumi.BoolPtrInput
	// Sign Key Transformer. Defaults to empty.
	XmlSignKeyInfoKeyNameTransformer pulumi.StringPtrInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}

type IdentityProviderInput interface {
	pulumi.Input

	ToIdentityProviderOutput() IdentityProviderOutput
	ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput
}

func (*IdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProvider)(nil))
}

func (i *IdentityProvider) ToIdentityProviderOutput() IdentityProviderOutput {
	return i.ToIdentityProviderOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderOutput)
}

func (i *IdentityProvider) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return i.ToIdentityProviderPtrOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderPtrOutput)
}

type IdentityProviderPtrInput interface {
	pulumi.Input

	ToIdentityProviderPtrOutput() IdentityProviderPtrOutput
	ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput
}

type identityProviderPtrType IdentityProviderArgs

func (*identityProviderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil))
}

func (i *identityProviderPtrType) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return i.ToIdentityProviderPtrOutputWithContext(context.Background())
}

func (i *identityProviderPtrType) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderPtrOutput)
}

// IdentityProviderArrayInput is an input type that accepts IdentityProviderArray and IdentityProviderArrayOutput values.
// You can construct a concrete instance of `IdentityProviderArrayInput` via:
//
//          IdentityProviderArray{ IdentityProviderArgs{...} }
type IdentityProviderArrayInput interface {
	pulumi.Input

	ToIdentityProviderArrayOutput() IdentityProviderArrayOutput
	ToIdentityProviderArrayOutputWithContext(context.Context) IdentityProviderArrayOutput
}

type IdentityProviderArray []IdentityProviderInput

func (IdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IdentityProvider)(nil))
}

func (i IdentityProviderArray) ToIdentityProviderArrayOutput() IdentityProviderArrayOutput {
	return i.ToIdentityProviderArrayOutputWithContext(context.Background())
}

func (i IdentityProviderArray) ToIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderArrayOutput)
}

// IdentityProviderMapInput is an input type that accepts IdentityProviderMap and IdentityProviderMapOutput values.
// You can construct a concrete instance of `IdentityProviderMapInput` via:
//
//          IdentityProviderMap{ "key": IdentityProviderArgs{...} }
type IdentityProviderMapInput interface {
	pulumi.Input

	ToIdentityProviderMapOutput() IdentityProviderMapOutput
	ToIdentityProviderMapOutputWithContext(context.Context) IdentityProviderMapOutput
}

type IdentityProviderMap map[string]IdentityProviderInput

func (IdentityProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IdentityProvider)(nil))
}

func (i IdentityProviderMap) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return i.ToIdentityProviderMapOutputWithContext(context.Background())
}

func (i IdentityProviderMap) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderMapOutput)
}

type IdentityProviderOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProvider)(nil))
}

func (o IdentityProviderOutput) ToIdentityProviderOutput() IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return o.ToIdentityProviderPtrOutputWithContext(context.Background())
}

func (o IdentityProviderOutput) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return o.ApplyT(func(v IdentityProvider) *IdentityProvider {
		return &v
	}).(IdentityProviderPtrOutput)
}

type IdentityProviderPtrOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil))
}

func (o IdentityProviderPtrOutput) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return o
}

func (o IdentityProviderPtrOutput) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return o
}

type IdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityProvider)(nil))
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutput() IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) Index(i pulumi.IntInput) IdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentityProvider {
		return vs[0].([]IdentityProvider)[vs[1].(int)]
	}).(IdentityProviderOutput)
}

type IdentityProviderMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityProvider)(nil))
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityProvider {
		return vs[0].(map[string]IdentityProvider)[vs[1].(string)]
	}).(IdentityProviderOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityProviderOutput{})
	pulumi.RegisterOutputType(IdentityProviderPtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderMapOutput{})
}
