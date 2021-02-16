// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing an attribute importer identity provider mapper within Keycloak.
//
// The attribute importer mapper can be used to map attributes from externally defined users to attributes or properties of the imported Keycloak user:
// - For the OIDC identity provider, this will map a claim on the ID or access token to an attribute for the imported Keycloak user.
// - For the SAML identity provider, this will map a SAML attribute found within the assertion to an attribute for the imported Keycloak user.
// - For social identity providers, this will map a JSON field from the user profile to an attribute for the imported Keycloak user.
//
// > If you are using Keycloak 10 or higher, you will need to specify the `extraConfig` argument in order to define a `syncMode` for the mapper.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/oidc"
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
// 		oidcIdentityProvider, err := oidc.NewIdentityProvider(ctx, "oidcIdentityProvider", &oidc.IdentityProviderArgs{
// 			Realm:            realm.ID(),
// 			Alias:            pulumi.String("oidc"),
// 			AuthorizationUrl: pulumi.String("https://example.com/auth"),
// 			TokenUrl:         pulumi.String("https://example.com/token"),
// 			ClientId:         pulumi.String("example_id"),
// 			ClientSecret:     pulumi.String("example_token"),
// 			DefaultScopes:    pulumi.String("openid random profile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewAttributeImporterIdentityProviderMapper(ctx, "oidcAttributeImporterIdentityProviderMapper", &keycloak.AttributeImporterIdentityProviderMapperArgs{
// 			Realm:                 realm.ID(),
// 			ClaimName:             pulumi.String("my-email-claim"),
// 			IdentityProviderAlias: oidcIdentityProvider.Alias,
// 			UserAttribute:         pulumi.String("email"),
// 			ExtraConfig: pulumi.StringMap{
// 				"syncMode": pulumi.String("INHERIT"),
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
// Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID. Examplebash
//
// ```sh
//  $ pulumi import keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
// ```
type AttributeImporterIdentityProviderMapper struct {
	pulumi.CustomResourceState

	// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
	AttributeFriendlyName pulumi.StringPtrOutput `pulumi:"attributeFriendlyName"`
	// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
	AttributeName pulumi.StringPtrOutput `pulumi:"attributeName"`
	// For OIDC based providers, this is the name of the claim to use.
	ClaimName pulumi.StringPtrOutput `pulumi:"claimName"`
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapOutput `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the realm.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// The user attribute or property name to store the mapped result.
	UserAttribute pulumi.StringOutput `pulumi:"userAttribute"`
}

// NewAttributeImporterIdentityProviderMapper registers a new resource with the given unique name, arguments, and options.
func NewAttributeImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, args *AttributeImporterIdentityProviderMapperArgs, opts ...pulumi.ResourceOption) (*AttributeImporterIdentityProviderMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	if args.UserAttribute == nil {
		return nil, errors.New("invalid value for required argument 'UserAttribute'")
	}
	var resource AttributeImporterIdentityProviderMapper
	err := ctx.RegisterResource("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttributeImporterIdentityProviderMapper gets an existing AttributeImporterIdentityProviderMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttributeImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttributeImporterIdentityProviderMapperState, opts ...pulumi.ResourceOption) (*AttributeImporterIdentityProviderMapper, error) {
	var resource AttributeImporterIdentityProviderMapper
	err := ctx.ReadResource("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttributeImporterIdentityProviderMapper resources.
type attributeImporterIdentityProviderMapperState struct {
	// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
	AttributeName *string `pulumi:"attributeName"`
	// For OIDC based providers, this is the name of the claim to use.
	ClaimName *string `pulumi:"claimName"`
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm *string `pulumi:"realm"`
	// The user attribute or property name to store the mapped result.
	UserAttribute *string `pulumi:"userAttribute"`
}

type AttributeImporterIdentityProviderMapperState struct {
	// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
	AttributeFriendlyName pulumi.StringPtrInput
	// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
	AttributeName pulumi.StringPtrInput
	// For OIDC based providers, this is the name of the claim to use.
	ClaimName pulumi.StringPtrInput
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringPtrInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringPtrInput
	// The user attribute or property name to store the mapped result.
	UserAttribute pulumi.StringPtrInput
}

func (AttributeImporterIdentityProviderMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeImporterIdentityProviderMapperState)(nil)).Elem()
}

type attributeImporterIdentityProviderMapperArgs struct {
	// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
	AttributeName *string `pulumi:"attributeName"`
	// For OIDC based providers, this is the name of the claim to use.
	ClaimName *string `pulumi:"claimName"`
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm string `pulumi:"realm"`
	// The user attribute or property name to store the mapped result.
	UserAttribute string `pulumi:"userAttribute"`
}

// The set of arguments for constructing a AttributeImporterIdentityProviderMapper resource.
type AttributeImporterIdentityProviderMapperArgs struct {
	// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
	AttributeFriendlyName pulumi.StringPtrInput
	// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
	AttributeName pulumi.StringPtrInput
	// For OIDC based providers, this is the name of the claim to use.
	ClaimName pulumi.StringPtrInput
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringInput
	// The user attribute or property name to store the mapped result.
	UserAttribute pulumi.StringInput
}

func (AttributeImporterIdentityProviderMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeImporterIdentityProviderMapperArgs)(nil)).Elem()
}

type AttributeImporterIdentityProviderMapperInput interface {
	pulumi.Input

	ToAttributeImporterIdentityProviderMapperOutput() AttributeImporterIdentityProviderMapperOutput
	ToAttributeImporterIdentityProviderMapperOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperOutput
}

func (*AttributeImporterIdentityProviderMapper) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributeImporterIdentityProviderMapper)(nil))
}

func (i *AttributeImporterIdentityProviderMapper) ToAttributeImporterIdentityProviderMapperOutput() AttributeImporterIdentityProviderMapperOutput {
	return i.ToAttributeImporterIdentityProviderMapperOutputWithContext(context.Background())
}

func (i *AttributeImporterIdentityProviderMapper) ToAttributeImporterIdentityProviderMapperOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperOutput)
}

func (i *AttributeImporterIdentityProviderMapper) ToAttributeImporterIdentityProviderMapperPtrOutput() AttributeImporterIdentityProviderMapperPtrOutput {
	return i.ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(context.Background())
}

func (i *AttributeImporterIdentityProviderMapper) ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperPtrOutput)
}

type AttributeImporterIdentityProviderMapperPtrInput interface {
	pulumi.Input

	ToAttributeImporterIdentityProviderMapperPtrOutput() AttributeImporterIdentityProviderMapperPtrOutput
	ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperPtrOutput
}

type attributeImporterIdentityProviderMapperPtrType AttributeImporterIdentityProviderMapperArgs

func (*attributeImporterIdentityProviderMapperPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeImporterIdentityProviderMapper)(nil))
}

func (i *attributeImporterIdentityProviderMapperPtrType) ToAttributeImporterIdentityProviderMapperPtrOutput() AttributeImporterIdentityProviderMapperPtrOutput {
	return i.ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(context.Background())
}

func (i *attributeImporterIdentityProviderMapperPtrType) ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperPtrOutput)
}

// AttributeImporterIdentityProviderMapperArrayInput is an input type that accepts AttributeImporterIdentityProviderMapperArray and AttributeImporterIdentityProviderMapperArrayOutput values.
// You can construct a concrete instance of `AttributeImporterIdentityProviderMapperArrayInput` via:
//
//          AttributeImporterIdentityProviderMapperArray{ AttributeImporterIdentityProviderMapperArgs{...} }
type AttributeImporterIdentityProviderMapperArrayInput interface {
	pulumi.Input

	ToAttributeImporterIdentityProviderMapperArrayOutput() AttributeImporterIdentityProviderMapperArrayOutput
	ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(context.Context) AttributeImporterIdentityProviderMapperArrayOutput
}

type AttributeImporterIdentityProviderMapperArray []AttributeImporterIdentityProviderMapperInput

func (AttributeImporterIdentityProviderMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AttributeImporterIdentityProviderMapper)(nil))
}

func (i AttributeImporterIdentityProviderMapperArray) ToAttributeImporterIdentityProviderMapperArrayOutput() AttributeImporterIdentityProviderMapperArrayOutput {
	return i.ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(context.Background())
}

func (i AttributeImporterIdentityProviderMapperArray) ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperArrayOutput)
}

// AttributeImporterIdentityProviderMapperMapInput is an input type that accepts AttributeImporterIdentityProviderMapperMap and AttributeImporterIdentityProviderMapperMapOutput values.
// You can construct a concrete instance of `AttributeImporterIdentityProviderMapperMapInput` via:
//
//          AttributeImporterIdentityProviderMapperMap{ "key": AttributeImporterIdentityProviderMapperArgs{...} }
type AttributeImporterIdentityProviderMapperMapInput interface {
	pulumi.Input

	ToAttributeImporterIdentityProviderMapperMapOutput() AttributeImporterIdentityProviderMapperMapOutput
	ToAttributeImporterIdentityProviderMapperMapOutputWithContext(context.Context) AttributeImporterIdentityProviderMapperMapOutput
}

type AttributeImporterIdentityProviderMapperMap map[string]AttributeImporterIdentityProviderMapperInput

func (AttributeImporterIdentityProviderMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AttributeImporterIdentityProviderMapper)(nil))
}

func (i AttributeImporterIdentityProviderMapperMap) ToAttributeImporterIdentityProviderMapperMapOutput() AttributeImporterIdentityProviderMapperMapOutput {
	return i.ToAttributeImporterIdentityProviderMapperMapOutputWithContext(context.Background())
}

func (i AttributeImporterIdentityProviderMapperMap) ToAttributeImporterIdentityProviderMapperMapOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperMapOutput)
}

type AttributeImporterIdentityProviderMapperOutput struct {
	*pulumi.OutputState
}

func (AttributeImporterIdentityProviderMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributeImporterIdentityProviderMapper)(nil))
}

func (o AttributeImporterIdentityProviderMapperOutput) ToAttributeImporterIdentityProviderMapperOutput() AttributeImporterIdentityProviderMapperOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperOutput) ToAttributeImporterIdentityProviderMapperOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperOutput) ToAttributeImporterIdentityProviderMapperPtrOutput() AttributeImporterIdentityProviderMapperPtrOutput {
	return o.ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(context.Background())
}

func (o AttributeImporterIdentityProviderMapperOutput) ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperPtrOutput {
	return o.ApplyT(func(v AttributeImporterIdentityProviderMapper) *AttributeImporterIdentityProviderMapper {
		return &v
	}).(AttributeImporterIdentityProviderMapperPtrOutput)
}

type AttributeImporterIdentityProviderMapperPtrOutput struct {
	*pulumi.OutputState
}

func (AttributeImporterIdentityProviderMapperPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeImporterIdentityProviderMapper)(nil))
}

func (o AttributeImporterIdentityProviderMapperPtrOutput) ToAttributeImporterIdentityProviderMapperPtrOutput() AttributeImporterIdentityProviderMapperPtrOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperPtrOutput) ToAttributeImporterIdentityProviderMapperPtrOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperPtrOutput {
	return o
}

type AttributeImporterIdentityProviderMapperArrayOutput struct{ *pulumi.OutputState }

func (AttributeImporterIdentityProviderMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttributeImporterIdentityProviderMapper)(nil))
}

func (o AttributeImporterIdentityProviderMapperArrayOutput) ToAttributeImporterIdentityProviderMapperArrayOutput() AttributeImporterIdentityProviderMapperArrayOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperArrayOutput) ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperArrayOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperArrayOutput) Index(i pulumi.IntInput) AttributeImporterIdentityProviderMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AttributeImporterIdentityProviderMapper {
		return vs[0].([]AttributeImporterIdentityProviderMapper)[vs[1].(int)]
	}).(AttributeImporterIdentityProviderMapperOutput)
}

type AttributeImporterIdentityProviderMapperMapOutput struct{ *pulumi.OutputState }

func (AttributeImporterIdentityProviderMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AttributeImporterIdentityProviderMapper)(nil))
}

func (o AttributeImporterIdentityProviderMapperMapOutput) ToAttributeImporterIdentityProviderMapperMapOutput() AttributeImporterIdentityProviderMapperMapOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperMapOutput) ToAttributeImporterIdentityProviderMapperMapOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperMapOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperMapOutput) MapIndex(k pulumi.StringInput) AttributeImporterIdentityProviderMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AttributeImporterIdentityProviderMapper {
		return vs[0].(map[string]AttributeImporterIdentityProviderMapper)[vs[1].(string)]
	}).(AttributeImporterIdentityProviderMapperOutput)
}

func init() {
	pulumi.RegisterOutputType(AttributeImporterIdentityProviderMapperOutput{})
	pulumi.RegisterOutputType(AttributeImporterIdentityProviderMapperPtrOutput{})
	pulumi.RegisterOutputType(AttributeImporterIdentityProviderMapperArrayOutput{})
	pulumi.RegisterOutputType(AttributeImporterIdentityProviderMapperMapOutput{})
}
