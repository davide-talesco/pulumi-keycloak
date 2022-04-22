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
// 		_, err = oidc.NewGoogleIdentityProvider(ctx, "google", &oidc.GoogleIdentityProviderArgs{
// 			Realm:        realm.ID(),
// 			ClientId:     pulumi.Any(_var.Google_identity_provider_client_id),
// 			ClientSecret: pulumi.Any(_var.Google_identity_provider_client_secret),
// 			TrustEmail:   pulumi.Bool(true),
// 			HostedDomain: pulumi.String("example.com"),
// 			SyncMode:     pulumi.String("IMPORT"),
// 			ExtraConfig: pulumi.AnyMap{
// 				"myCustomConfigKey": pulumi.Any("myValue"),
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
// This resource does not yet support importing.
type GoogleIdentityProvider struct {
	pulumi.CustomResourceState

	// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient pulumi.BoolPtrOutput `pulumi:"acceptsPromptNoneForwardFromClient"`
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrOutput `pulumi:"addReadTokenRoleOnCreate"`
	// (Computed) The alias for the Google identity provider.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrOutput `pulumi:"authenticateByDefault"`
	// The client or client identifier registered within the identity provider.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
	DefaultScopes pulumi.StringPtrOutput `pulumi:"defaultScopes"`
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo pulumi.BoolPtrOutput `pulumi:"disableUserInfo"`
	// (Computed) Display name for the Google identity provider in the GUI.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     pulumi.BoolPtrOutput `pulumi:"enabled"`
	ExtraConfig pulumi.MapOutput     `pulumi:"extraConfig"`
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"firstBrokerLoginFlowAlias"`
	// A number defining the order of this identity provider in the GUI.
	GuiOrder pulumi.StringPtrOutput `pulumi:"guiOrder"`
	// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
	HideOnLoginPage pulumi.BoolPtrOutput `pulumi:"hideOnLoginPage"`
	// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
	HostedDomain pulumi.StringPtrOutput `pulumi:"hostedDomain"`
	// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrOutput `pulumi:"linkOnly"`
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"postBrokerLoginFlowAlias"`
	// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId pulumi.StringPtrOutput `pulumi:"providerId"`
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
	RequestRefreshToken pulumi.BoolPtrOutput `pulumi:"requestRefreshToken"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrOutput `pulumi:"storeToken"`
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode pulumi.StringPtrOutput `pulumi:"syncMode"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrOutput `pulumi:"trustEmail"`
	// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
	UseUserIpParam pulumi.BoolPtrOutput `pulumi:"useUserIpParam"`
}

// NewGoogleIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewGoogleIdentityProvider(ctx *pulumi.Context,
	name string, args *GoogleIdentityProviderArgs, opts ...pulumi.ResourceOption) (*GoogleIdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource GoogleIdentityProvider
	err := ctx.RegisterResource("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGoogleIdentityProvider gets an existing GoogleIdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGoogleIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GoogleIdentityProviderState, opts ...pulumi.ResourceOption) (*GoogleIdentityProvider, error) {
	var resource GoogleIdentityProvider
	err := ctx.ReadResource("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GoogleIdentityProvider resources.
type googleIdentityProviderState struct {
	// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient *bool `pulumi:"acceptsPromptNoneForwardFromClient"`
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// (Computed) The alias for the Google identity provider.
	Alias *string `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// The client or client identifier registered within the identity provider.
	ClientId *string `pulumi:"clientId"`
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret *string `pulumi:"clientSecret"`
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
	DefaultScopes *string `pulumi:"defaultScopes"`
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo *bool `pulumi:"disableUserInfo"`
	// (Computed) Display name for the Google identity provider in the GUI.
	DisplayName *string `pulumi:"displayName"`
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     *bool                  `pulumi:"enabled"`
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// A number defining the order of this identity provider in the GUI.
	GuiOrder *string `pulumi:"guiOrder"`
	// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
	HostedDomain *string `pulumi:"hostedDomain"`
	// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
	InternalId *string `pulumi:"internalId"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly *bool `pulumi:"linkOnly"`
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId *string `pulumi:"providerId"`
	// The name of the realm. This is unique across Keycloak.
	Realm *string `pulumi:"realm"`
	// Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
	RequestRefreshToken *bool `pulumi:"requestRefreshToken"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken *bool `pulumi:"storeToken"`
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode *string `pulumi:"syncMode"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
	UseUserIpParam *bool `pulumi:"useUserIpParam"`
}

type GoogleIdentityProviderState struct {
	// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient pulumi.BoolPtrInput
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// (Computed) The alias for the Google identity provider.
	Alias pulumi.StringPtrInput
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrInput
	// The client or client identifier registered within the identity provider.
	ClientId pulumi.StringPtrInput
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret pulumi.StringPtrInput
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
	DefaultScopes pulumi.StringPtrInput
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo pulumi.BoolPtrInput
	// (Computed) Display name for the Google identity provider in the GUI.
	DisplayName pulumi.StringPtrInput
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     pulumi.BoolPtrInput
	ExtraConfig pulumi.MapInput
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// A number defining the order of this identity provider in the GUI.
	GuiOrder pulumi.StringPtrInput
	// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
	HideOnLoginPage pulumi.BoolPtrInput
	// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
	HostedDomain pulumi.StringPtrInput
	// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
	InternalId pulumi.StringPtrInput
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrInput
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId pulumi.StringPtrInput
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringPtrInput
	// Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
	RequestRefreshToken pulumi.BoolPtrInput
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrInput
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode pulumi.StringPtrInput
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrInput
	// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
	UseUserIpParam pulumi.BoolPtrInput
}

func (GoogleIdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*googleIdentityProviderState)(nil)).Elem()
}

type googleIdentityProviderArgs struct {
	// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient *bool `pulumi:"acceptsPromptNoneForwardFromClient"`
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// The client or client identifier registered within the identity provider.
	ClientId string `pulumi:"clientId"`
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret string `pulumi:"clientSecret"`
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
	DefaultScopes *string `pulumi:"defaultScopes"`
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo *bool `pulumi:"disableUserInfo"`
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     *bool                  `pulumi:"enabled"`
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// A number defining the order of this identity provider in the GUI.
	GuiOrder *string `pulumi:"guiOrder"`
	// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
	HostedDomain *string `pulumi:"hostedDomain"`
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly *bool `pulumi:"linkOnly"`
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId *string `pulumi:"providerId"`
	// The name of the realm. This is unique across Keycloak.
	Realm string `pulumi:"realm"`
	// Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
	RequestRefreshToken *bool `pulumi:"requestRefreshToken"`
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken *bool `pulumi:"storeToken"`
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode *string `pulumi:"syncMode"`
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
	UseUserIpParam *bool `pulumi:"useUserIpParam"`
}

// The set of arguments for constructing a GoogleIdentityProvider resource.
type GoogleIdentityProviderArgs struct {
	// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
	AcceptsPromptNoneForwardFromClient pulumi.BoolPtrInput
	// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrInput
	// The client or client identifier registered within the identity provider.
	ClientId pulumi.StringInput
	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	ClientSecret pulumi.StringInput
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
	DefaultScopes pulumi.StringPtrInput
	// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
	DisableUserInfo pulumi.BoolPtrInput
	// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
	Enabled     pulumi.BoolPtrInput
	ExtraConfig pulumi.MapInput
	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// A number defining the order of this identity provider in the GUI.
	GuiOrder pulumi.StringPtrInput
	// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
	HideOnLoginPage pulumi.BoolPtrInput
	// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
	HostedDomain pulumi.StringPtrInput
	// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
	LinkOnly pulumi.BoolPtrInput
	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
	ProviderId pulumi.StringPtrInput
	// The name of the realm. This is unique across Keycloak.
	Realm pulumi.StringInput
	// Sets the "accessType" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
	RequestRefreshToken pulumi.BoolPtrInput
	// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
	StoreToken pulumi.BoolPtrInput
	// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
	SyncMode pulumi.StringPtrInput
	// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
	TrustEmail pulumi.BoolPtrInput
	// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
	UseUserIpParam pulumi.BoolPtrInput
}

func (GoogleIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*googleIdentityProviderArgs)(nil)).Elem()
}

type GoogleIdentityProviderInput interface {
	pulumi.Input

	ToGoogleIdentityProviderOutput() GoogleIdentityProviderOutput
	ToGoogleIdentityProviderOutputWithContext(ctx context.Context) GoogleIdentityProviderOutput
}

func (*GoogleIdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleIdentityProvider)(nil)).Elem()
}

func (i *GoogleIdentityProvider) ToGoogleIdentityProviderOutput() GoogleIdentityProviderOutput {
	return i.ToGoogleIdentityProviderOutputWithContext(context.Background())
}

func (i *GoogleIdentityProvider) ToGoogleIdentityProviderOutputWithContext(ctx context.Context) GoogleIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleIdentityProviderOutput)
}

// GoogleIdentityProviderArrayInput is an input type that accepts GoogleIdentityProviderArray and GoogleIdentityProviderArrayOutput values.
// You can construct a concrete instance of `GoogleIdentityProviderArrayInput` via:
//
//          GoogleIdentityProviderArray{ GoogleIdentityProviderArgs{...} }
type GoogleIdentityProviderArrayInput interface {
	pulumi.Input

	ToGoogleIdentityProviderArrayOutput() GoogleIdentityProviderArrayOutput
	ToGoogleIdentityProviderArrayOutputWithContext(context.Context) GoogleIdentityProviderArrayOutput
}

type GoogleIdentityProviderArray []GoogleIdentityProviderInput

func (GoogleIdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GoogleIdentityProvider)(nil)).Elem()
}

func (i GoogleIdentityProviderArray) ToGoogleIdentityProviderArrayOutput() GoogleIdentityProviderArrayOutput {
	return i.ToGoogleIdentityProviderArrayOutputWithContext(context.Background())
}

func (i GoogleIdentityProviderArray) ToGoogleIdentityProviderArrayOutputWithContext(ctx context.Context) GoogleIdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleIdentityProviderArrayOutput)
}

// GoogleIdentityProviderMapInput is an input type that accepts GoogleIdentityProviderMap and GoogleIdentityProviderMapOutput values.
// You can construct a concrete instance of `GoogleIdentityProviderMapInput` via:
//
//          GoogleIdentityProviderMap{ "key": GoogleIdentityProviderArgs{...} }
type GoogleIdentityProviderMapInput interface {
	pulumi.Input

	ToGoogleIdentityProviderMapOutput() GoogleIdentityProviderMapOutput
	ToGoogleIdentityProviderMapOutputWithContext(context.Context) GoogleIdentityProviderMapOutput
}

type GoogleIdentityProviderMap map[string]GoogleIdentityProviderInput

func (GoogleIdentityProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GoogleIdentityProvider)(nil)).Elem()
}

func (i GoogleIdentityProviderMap) ToGoogleIdentityProviderMapOutput() GoogleIdentityProviderMapOutput {
	return i.ToGoogleIdentityProviderMapOutputWithContext(context.Background())
}

func (i GoogleIdentityProviderMap) ToGoogleIdentityProviderMapOutputWithContext(ctx context.Context) GoogleIdentityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleIdentityProviderMapOutput)
}

type GoogleIdentityProviderOutput struct{ *pulumi.OutputState }

func (GoogleIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleIdentityProvider)(nil)).Elem()
}

func (o GoogleIdentityProviderOutput) ToGoogleIdentityProviderOutput() GoogleIdentityProviderOutput {
	return o
}

func (o GoogleIdentityProviderOutput) ToGoogleIdentityProviderOutputWithContext(ctx context.Context) GoogleIdentityProviderOutput {
	return o
}

type GoogleIdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (GoogleIdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GoogleIdentityProvider)(nil)).Elem()
}

func (o GoogleIdentityProviderArrayOutput) ToGoogleIdentityProviderArrayOutput() GoogleIdentityProviderArrayOutput {
	return o
}

func (o GoogleIdentityProviderArrayOutput) ToGoogleIdentityProviderArrayOutputWithContext(ctx context.Context) GoogleIdentityProviderArrayOutput {
	return o
}

func (o GoogleIdentityProviderArrayOutput) Index(i pulumi.IntInput) GoogleIdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GoogleIdentityProvider {
		return vs[0].([]*GoogleIdentityProvider)[vs[1].(int)]
	}).(GoogleIdentityProviderOutput)
}

type GoogleIdentityProviderMapOutput struct{ *pulumi.OutputState }

func (GoogleIdentityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GoogleIdentityProvider)(nil)).Elem()
}

func (o GoogleIdentityProviderMapOutput) ToGoogleIdentityProviderMapOutput() GoogleIdentityProviderMapOutput {
	return o
}

func (o GoogleIdentityProviderMapOutput) ToGoogleIdentityProviderMapOutputWithContext(ctx context.Context) GoogleIdentityProviderMapOutput {
	return o
}

func (o GoogleIdentityProviderMapOutput) MapIndex(k pulumi.StringInput) GoogleIdentityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GoogleIdentityProvider {
		return vs[0].(map[string]*GoogleIdentityProvider)[vs[1].(string)]
	}).(GoogleIdentityProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleIdentityProviderInput)(nil)).Elem(), &GoogleIdentityProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleIdentityProviderArrayInput)(nil)).Elem(), GoogleIdentityProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleIdentityProviderMapInput)(nil)).Elem(), GoogleIdentityProviderMap{})
	pulumi.RegisterOutputType(GoogleIdentityProviderOutput{})
	pulumi.RegisterOutputType(GoogleIdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(GoogleIdentityProviderMapOutput{})
}
