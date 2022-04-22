// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oidc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing OIDC Identity Providers within Keycloak.
//
// OIDC (OpenID Connect) identity providers allows users to authenticate through a third party system using the OIDC standard.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/oidc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		_, err = oidc.NewIdentityProvider(ctx, "realmIdentityProvider", &oidc.IdentityProviderArgs{
// 			Realm:            realm.ID(),
// 			Alias:            pulumi.String("my-idp"),
// 			AuthorizationUrl: pulumi.String("https://authorizationurl.com"),
// 			ClientId:         pulumi.String("clientID"),
// 			ClientSecret:     pulumi.String("clientSecret"),
// 			TokenUrl:         pulumi.String("https://tokenurl.com"),
// 			ExtraConfig: pulumi.AnyMap{
// 				"clientAuthMethod": pulumi.Any("client_secret_post"),
// 			},
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
//  $ pulumi import keycloak:oidc/identityProvider:IdentityProvider realm_identity_provider my-realm/my-idp
// ```
type IdentityProvider struct {
	pulumi.CustomResourceState

	// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient pulumi.BoolPtrOutput `pulumi:"acceptsPromptNoneForwardFromClient"`
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrOutput `pulumi:"addReadTokenRoleOnCreate"`
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrOutput `pulumi:"authenticateByDefault"`
	// The Authorization Url.
	AuthorizationUrl pulumi.StringOutput `pulumi:"authorizationUrl"`
	// Does the external IDP support backchannel logout? Defaults to `true`.
	BackchannelSupported pulumi.BoolPtrOutput `pulumi:"backchannelSupported"`
	// The client or client identifier registered within the identity provider.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
	DefaultScopes pulumi.StringPtrOutput `pulumi:"defaultScopes"`
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo pulumi.BoolPtrOutput `pulumi:"disableUserInfo"`
	// Display name for the identity provider in the GUI.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     pulumi.BoolPtrOutput `pulumi:"enabled"`
	ExtraConfig pulumi.MapOutput     `pulumi:"extraConfig"`
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"firstBrokerLoginFlowAlias"`
	// A number defining the order of this identity provider in the GUI.
	GuiOrder pulumi.StringPtrOutput `pulumi:"guiOrder"`
	// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
	HideOnLoginPage pulumi.BoolPtrOutput `pulumi:"hideOnLoginPage"`
	// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// JSON Web Key Set URL.
	JwksUrl pulumi.StringPtrOutput `pulumi:"jwksUrl"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrOutput `pulumi:"linkOnly"`
	// Pass login hint to identity provider.
	LoginHint pulumi.StringPtrOutput `pulumi:"loginHint"`
	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	LogoutUrl pulumi.StringPtrOutput `pulumi:"logoutUrl"`
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"postBrokerLoginFlowAlias"`
	// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId pulumi.StringPtrOutput `pulumi:"providerId"`
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrOutput `pulumi:"storeToken"`
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode pulumi.StringPtrOutput `pulumi:"syncMode"`
	// The Token URL.
	TokenUrl pulumi.StringOutput `pulumi:"tokenUrl"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrOutput `pulumi:"trustEmail"`
	// Pass current locale to identity provider. Defaults to `false`.
	UiLocales pulumi.BoolPtrOutput `pulumi:"uiLocales"`
	// User Info URL.
	UserInfoUrl pulumi.StringPtrOutput `pulumi:"userInfoUrl"`
	// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
	ValidateSignature pulumi.BoolPtrOutput `pulumi:"validateSignature"`
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
	if args.AuthorizationUrl == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationUrl'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	if args.TokenUrl == nil {
		return nil, errors.New("invalid value for required argument 'TokenUrl'")
	}
	var resource IdentityProvider
	err := ctx.RegisterResource("keycloak:oidc/identityProvider:IdentityProvider", name, args, &resource, opts...)
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
	err := ctx.ReadResource("keycloak:oidc/identityProvider:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
	// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient *bool `pulumi:"acceptsPromptNoneForwardFromClient"`
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias *string `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// The Authorization Url.
	AuthorizationUrl *string `pulumi:"authorizationUrl"`
	// Does the external IDP support backchannel logout? Defaults to `true`.
	BackchannelSupported *bool `pulumi:"backchannelSupported"`
	// The client or client identifier registered within the identity provider.
	ClientId *string `pulumi:"clientId"`
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret *string `pulumi:"clientSecret"`
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
	DefaultScopes *string `pulumi:"defaultScopes"`
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo *bool `pulumi:"disableUserInfo"`
	// Display name for the identity provider in the GUI.
	DisplayName *string `pulumi:"displayName"`
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     *bool                  `pulumi:"enabled"`
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// A number defining the order of this identity provider in the GUI.
	GuiOrder *string `pulumi:"guiOrder"`
	// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
	InternalId *string `pulumi:"internalId"`
	// JSON Web Key Set URL.
	JwksUrl *string `pulumi:"jwksUrl"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly *bool `pulumi:"linkOnly"`
	// Pass login hint to identity provider.
	LoginHint *string `pulumi:"loginHint"`
	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId *string `pulumi:"providerId"`
	// The name of the realm. This is unique across Keycloak.
	Realm *string `pulumi:"realm"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken *bool `pulumi:"storeToken"`
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode *string `pulumi:"syncMode"`
	// The Token URL.
	TokenUrl *string `pulumi:"tokenUrl"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Pass current locale to identity provider. Defaults to `false`.
	UiLocales *bool `pulumi:"uiLocales"`
	// User Info URL.
	UserInfoUrl *string `pulumi:"userInfoUrl"`
	// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
	ValidateSignature *bool `pulumi:"validateSignature"`
}

type IdentityProviderState struct {
	// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient pulumi.BoolPtrInput
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias pulumi.StringPtrInput
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrInput
	// The Authorization Url.
	AuthorizationUrl pulumi.StringPtrInput
	// Does the external IDP support backchannel logout? Defaults to `true`.
	BackchannelSupported pulumi.BoolPtrInput
	// The client or client identifier registered within the identity provider.
	ClientId pulumi.StringPtrInput
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret pulumi.StringPtrInput
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
	DefaultScopes pulumi.StringPtrInput
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo pulumi.BoolPtrInput
	// Display name for the identity provider in the GUI.
	DisplayName pulumi.StringPtrInput
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     pulumi.BoolPtrInput
	ExtraConfig pulumi.MapInput
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// A number defining the order of this identity provider in the GUI.
	GuiOrder pulumi.StringPtrInput
	// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
	HideOnLoginPage pulumi.BoolPtrInput
	// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
	InternalId pulumi.StringPtrInput
	// JSON Web Key Set URL.
	JwksUrl pulumi.StringPtrInput
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrInput
	// Pass login hint to identity provider.
	LoginHint pulumi.StringPtrInput
	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	LogoutUrl pulumi.StringPtrInput
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId pulumi.StringPtrInput
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringPtrInput
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrInput
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode pulumi.StringPtrInput
	// The Token URL.
	TokenUrl pulumi.StringPtrInput
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrInput
	// Pass current locale to identity provider. Defaults to `false`.
	UiLocales pulumi.BoolPtrInput
	// User Info URL.
	UserInfoUrl pulumi.StringPtrInput
	// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
	ValidateSignature pulumi.BoolPtrInput
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient *bool `pulumi:"acceptsPromptNoneForwardFromClient"`
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias string `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// The Authorization Url.
	AuthorizationUrl string `pulumi:"authorizationUrl"`
	// Does the external IDP support backchannel logout? Defaults to `true`.
	BackchannelSupported *bool `pulumi:"backchannelSupported"`
	// The client or client identifier registered within the identity provider.
	ClientId string `pulumi:"clientId"`
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret string `pulumi:"clientSecret"`
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
	DefaultScopes *string `pulumi:"defaultScopes"`
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo *bool `pulumi:"disableUserInfo"`
	// Display name for the identity provider in the GUI.
	DisplayName *string `pulumi:"displayName"`
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     *bool                  `pulumi:"enabled"`
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// A number defining the order of this identity provider in the GUI.
	GuiOrder *string `pulumi:"guiOrder"`
	// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// JSON Web Key Set URL.
	JwksUrl *string `pulumi:"jwksUrl"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly *bool `pulumi:"linkOnly"`
	// Pass login hint to identity provider.
	LoginHint *string `pulumi:"loginHint"`
	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId *string `pulumi:"providerId"`
	// The name of the realm. This is unique across Keycloak.
	Realm string `pulumi:"realm"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken *bool `pulumi:"storeToken"`
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode *string `pulumi:"syncMode"`
	// The Token URL.
	TokenUrl string `pulumi:"tokenUrl"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Pass current locale to identity provider. Defaults to `false`.
	UiLocales *bool `pulumi:"uiLocales"`
	// User Info URL.
	UserInfoUrl *string `pulumi:"userInfoUrl"`
	// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
	ValidateSignature *bool `pulumi:"validateSignature"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient pulumi.BoolPtrInput
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias pulumi.StringInput
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrInput
	// The Authorization Url.
	AuthorizationUrl pulumi.StringInput
	// Does the external IDP support backchannel logout? Defaults to `true`.
	BackchannelSupported pulumi.BoolPtrInput
	// The client or client identifier registered within the identity provider.
	ClientId pulumi.StringInput
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret pulumi.StringInput
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
	DefaultScopes pulumi.StringPtrInput
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo pulumi.BoolPtrInput
	// Display name for the identity provider in the GUI.
	DisplayName pulumi.StringPtrInput
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     pulumi.BoolPtrInput
	ExtraConfig pulumi.MapInput
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// A number defining the order of this identity provider in the GUI.
	GuiOrder pulumi.StringPtrInput
	// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
	HideOnLoginPage pulumi.BoolPtrInput
	// JSON Web Key Set URL.
	JwksUrl pulumi.StringPtrInput
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrInput
	// Pass login hint to identity provider.
	LoginHint pulumi.StringPtrInput
	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	LogoutUrl pulumi.StringPtrInput
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId pulumi.StringPtrInput
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringInput
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrInput
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode pulumi.StringPtrInput
	// The Token URL.
	TokenUrl pulumi.StringInput
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrInput
	// Pass current locale to identity provider. Defaults to `false`.
	UiLocales pulumi.BoolPtrInput
	// User Info URL.
	UserInfoUrl pulumi.StringPtrInput
	// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
	ValidateSignature pulumi.BoolPtrInput
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
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (i *IdentityProvider) ToIdentityProviderOutput() IdentityProviderOutput {
	return i.ToIdentityProviderOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderOutput)
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
	return reflect.TypeOf((*[]*IdentityProvider)(nil)).Elem()
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
	return reflect.TypeOf((*map[string]*IdentityProvider)(nil)).Elem()
}

func (i IdentityProviderMap) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return i.ToIdentityProviderMapOutputWithContext(context.Background())
}

func (i IdentityProviderMap) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderMapOutput)
}

type IdentityProviderOutput struct{ *pulumi.OutputState }

func (IdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderOutput) ToIdentityProviderOutput() IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return o
}

type IdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutput() IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) Index(i pulumi.IntInput) IdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityProvider {
		return vs[0].([]*IdentityProvider)[vs[1].(int)]
	}).(IdentityProviderOutput)
}

type IdentityProviderMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityProvider {
		return vs[0].(map[string]*IdentityProvider)[vs[1].(string)]
	}).(IdentityProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderInput)(nil)).Elem(), &IdentityProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderArrayInput)(nil)).Elem(), IdentityProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderMapInput)(nil)).Elem(), IdentityProviderMap{})
	pulumi.RegisterOutputType(IdentityProviderOutput{})
	pulumi.RegisterOutputType(IdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderMapOutput{})
}
