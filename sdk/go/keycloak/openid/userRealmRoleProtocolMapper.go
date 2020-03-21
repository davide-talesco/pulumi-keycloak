// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # openid.UserRealmRoleProtocolMapper
//
// Allows for creating and managing user realm role protocol mappers within
// Keycloak.
//
// User realm role protocol mappers allow you to define a claim containing the list of the realm roles.
// Protocol mappers can be defined for a single client, or they can
// be defined for a client scope which can be shared between multiple different
// clients.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `claimName` - (Required) The name of the claim to insert into a token.
// - `claimValueType` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
// - `multivalued` - (Optional) Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `true`.
// - `realmRolePrefix` - (Optional) A prefix for each Realm Role.
// - `addToIdToken` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
// - `addToUserinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_openid_user_realm_role_protocol_mapper.html.markdown.
type UserRealmRoleProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	ClaimName pulumi.StringOutput `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrOutput `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrOutput `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Prefix that will be added to each realm role.
	RealmRolePrefix pulumi.StringPtrOutput `pulumi:"realmRolePrefix"`
}

// NewUserRealmRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserRealmRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *UserRealmRoleProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserRealmRoleProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &UserRealmRoleProtocolMapperArgs{}
	}
	var resource UserRealmRoleProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRealmRoleProtocolMapper gets an existing UserRealmRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRealmRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRealmRoleProtocolMapperState, opts ...pulumi.ResourceOption) (*UserRealmRoleProtocolMapper, error) {
	var resource UserRealmRoleProtocolMapper
	err := ctx.ReadResource("keycloak:openid/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRealmRoleProtocolMapper resources.
type userRealmRoleProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	ClaimName *string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued *bool `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
	// Prefix that will be added to each realm role.
	RealmRolePrefix *string `pulumi:"realmRolePrefix"`
}

type UserRealmRoleProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName pulumi.StringPtrInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
	// Prefix that will be added to each realm role.
	RealmRolePrefix pulumi.StringPtrInput
}

func (UserRealmRoleProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRealmRoleProtocolMapperState)(nil)).Elem()
}

type userRealmRoleProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	ClaimName string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued *bool `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
	// Prefix that will be added to each realm role.
	RealmRolePrefix *string `pulumi:"realmRolePrefix"`
}

// The set of arguments for constructing a UserRealmRoleProtocolMapper resource.
type UserRealmRoleProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName pulumi.StringInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
	// Prefix that will be added to each realm role.
	RealmRolePrefix pulumi.StringPtrInput
}

func (UserRealmRoleProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRealmRoleProtocolMapperArgs)(nil)).Elem()
}

