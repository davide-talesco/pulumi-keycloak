// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttributeToRoleIdentityMapper struct {
	pulumi.CustomResourceState

	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrOutput `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName pulumi.StringPtrOutput `pulumi:"attributeName"`
	// Attribute Value
	AttributeValue pulumi.StringPtrOutput `pulumi:"attributeValue"`
	// OIDC Claim Name
	ClaimName pulumi.StringPtrOutput `pulumi:"claimName"`
	// OIDC Claim Value
	ClaimValue  pulumi.StringPtrOutput `pulumi:"claimValue"`
	ExtraConfig pulumi.MapOutput       `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Role Name
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAttributeToRoleIdentityMapper registers a new resource with the given unique name, arguments, and options.
func NewAttributeToRoleIdentityMapper(ctx *pulumi.Context,
	name string, args *AttributeToRoleIdentityMapperArgs, opts ...pulumi.ResourceOption) (*AttributeToRoleIdentityMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource AttributeToRoleIdentityMapper
	err := ctx.RegisterResource("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttributeToRoleIdentityMapper gets an existing AttributeToRoleIdentityMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttributeToRoleIdentityMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttributeToRoleIdentityMapperState, opts ...pulumi.ResourceOption) (*AttributeToRoleIdentityMapper, error) {
	var resource AttributeToRoleIdentityMapper
	err := ctx.ReadResource("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttributeToRoleIdentityMapper resources.
type attributeToRoleIdentityMapperState struct {
	// Attribute Friendly Name
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName *string `pulumi:"attributeName"`
	// Attribute Value
	AttributeValue *string `pulumi:"attributeValue"`
	// OIDC Claim Name
	ClaimName *string `pulumi:"claimName"`
	// OIDC Claim Value
	ClaimValue  *string                `pulumi:"claimValue"`
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// Role Name
	Role *string `pulumi:"role"`
}

type AttributeToRoleIdentityMapperState struct {
	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name
	AttributeName pulumi.StringPtrInput
	// Attribute Value
	AttributeValue pulumi.StringPtrInput
	// OIDC Claim Name
	ClaimName pulumi.StringPtrInput
	// OIDC Claim Value
	ClaimValue  pulumi.StringPtrInput
	ExtraConfig pulumi.MapInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringPtrInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// Role Name
	Role pulumi.StringPtrInput
}

func (AttributeToRoleIdentityMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeToRoleIdentityMapperState)(nil)).Elem()
}

type attributeToRoleIdentityMapperArgs struct {
	// Attribute Friendly Name
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName *string `pulumi:"attributeName"`
	// Attribute Value
	AttributeValue *string `pulumi:"attributeValue"`
	// OIDC Claim Name
	ClaimName *string `pulumi:"claimName"`
	// OIDC Claim Value
	ClaimValue  *string                `pulumi:"claimValue"`
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// Role Name
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AttributeToRoleIdentityMapper resource.
type AttributeToRoleIdentityMapperArgs struct {
	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name
	AttributeName pulumi.StringPtrInput
	// Attribute Value
	AttributeValue pulumi.StringPtrInput
	// OIDC Claim Name
	ClaimName pulumi.StringPtrInput
	// OIDC Claim Value
	ClaimValue  pulumi.StringPtrInput
	ExtraConfig pulumi.MapInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// Role Name
	Role pulumi.StringInput
}

func (AttributeToRoleIdentityMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeToRoleIdentityMapperArgs)(nil)).Elem()
}

type AttributeToRoleIdentityMapperInput interface {
	pulumi.Input

	ToAttributeToRoleIdentityMapperOutput() AttributeToRoleIdentityMapperOutput
	ToAttributeToRoleIdentityMapperOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperOutput
}

func (*AttributeToRoleIdentityMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (i *AttributeToRoleIdentityMapper) ToAttributeToRoleIdentityMapperOutput() AttributeToRoleIdentityMapperOutput {
	return i.ToAttributeToRoleIdentityMapperOutputWithContext(context.Background())
}

func (i *AttributeToRoleIdentityMapper) ToAttributeToRoleIdentityMapperOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeToRoleIdentityMapperOutput)
}

// AttributeToRoleIdentityMapperArrayInput is an input type that accepts AttributeToRoleIdentityMapperArray and AttributeToRoleIdentityMapperArrayOutput values.
// You can construct a concrete instance of `AttributeToRoleIdentityMapperArrayInput` via:
//
//          AttributeToRoleIdentityMapperArray{ AttributeToRoleIdentityMapperArgs{...} }
type AttributeToRoleIdentityMapperArrayInput interface {
	pulumi.Input

	ToAttributeToRoleIdentityMapperArrayOutput() AttributeToRoleIdentityMapperArrayOutput
	ToAttributeToRoleIdentityMapperArrayOutputWithContext(context.Context) AttributeToRoleIdentityMapperArrayOutput
}

type AttributeToRoleIdentityMapperArray []AttributeToRoleIdentityMapperInput

func (AttributeToRoleIdentityMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (i AttributeToRoleIdentityMapperArray) ToAttributeToRoleIdentityMapperArrayOutput() AttributeToRoleIdentityMapperArrayOutput {
	return i.ToAttributeToRoleIdentityMapperArrayOutputWithContext(context.Background())
}

func (i AttributeToRoleIdentityMapperArray) ToAttributeToRoleIdentityMapperArrayOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeToRoleIdentityMapperArrayOutput)
}

// AttributeToRoleIdentityMapperMapInput is an input type that accepts AttributeToRoleIdentityMapperMap and AttributeToRoleIdentityMapperMapOutput values.
// You can construct a concrete instance of `AttributeToRoleIdentityMapperMapInput` via:
//
//          AttributeToRoleIdentityMapperMap{ "key": AttributeToRoleIdentityMapperArgs{...} }
type AttributeToRoleIdentityMapperMapInput interface {
	pulumi.Input

	ToAttributeToRoleIdentityMapperMapOutput() AttributeToRoleIdentityMapperMapOutput
	ToAttributeToRoleIdentityMapperMapOutputWithContext(context.Context) AttributeToRoleIdentityMapperMapOutput
}

type AttributeToRoleIdentityMapperMap map[string]AttributeToRoleIdentityMapperInput

func (AttributeToRoleIdentityMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (i AttributeToRoleIdentityMapperMap) ToAttributeToRoleIdentityMapperMapOutput() AttributeToRoleIdentityMapperMapOutput {
	return i.ToAttributeToRoleIdentityMapperMapOutputWithContext(context.Background())
}

func (i AttributeToRoleIdentityMapperMap) ToAttributeToRoleIdentityMapperMapOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeToRoleIdentityMapperMapOutput)
}

type AttributeToRoleIdentityMapperOutput struct{ *pulumi.OutputState }

func (AttributeToRoleIdentityMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (o AttributeToRoleIdentityMapperOutput) ToAttributeToRoleIdentityMapperOutput() AttributeToRoleIdentityMapperOutput {
	return o
}

func (o AttributeToRoleIdentityMapperOutput) ToAttributeToRoleIdentityMapperOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperOutput {
	return o
}

type AttributeToRoleIdentityMapperArrayOutput struct{ *pulumi.OutputState }

func (AttributeToRoleIdentityMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (o AttributeToRoleIdentityMapperArrayOutput) ToAttributeToRoleIdentityMapperArrayOutput() AttributeToRoleIdentityMapperArrayOutput {
	return o
}

func (o AttributeToRoleIdentityMapperArrayOutput) ToAttributeToRoleIdentityMapperArrayOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperArrayOutput {
	return o
}

func (o AttributeToRoleIdentityMapperArrayOutput) Index(i pulumi.IntInput) AttributeToRoleIdentityMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AttributeToRoleIdentityMapper {
		return vs[0].([]*AttributeToRoleIdentityMapper)[vs[1].(int)]
	}).(AttributeToRoleIdentityMapperOutput)
}

type AttributeToRoleIdentityMapperMapOutput struct{ *pulumi.OutputState }

func (AttributeToRoleIdentityMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (o AttributeToRoleIdentityMapperMapOutput) ToAttributeToRoleIdentityMapperMapOutput() AttributeToRoleIdentityMapperMapOutput {
	return o
}

func (o AttributeToRoleIdentityMapperMapOutput) ToAttributeToRoleIdentityMapperMapOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperMapOutput {
	return o
}

func (o AttributeToRoleIdentityMapperMapOutput) MapIndex(k pulumi.StringInput) AttributeToRoleIdentityMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AttributeToRoleIdentityMapper {
		return vs[0].(map[string]*AttributeToRoleIdentityMapper)[vs[1].(string)]
	}).(AttributeToRoleIdentityMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeToRoleIdentityMapperInput)(nil)).Elem(), &AttributeToRoleIdentityMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeToRoleIdentityMapperArrayInput)(nil)).Elem(), AttributeToRoleIdentityMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeToRoleIdentityMapperMapInput)(nil)).Elem(), AttributeToRoleIdentityMapperMap{})
	pulumi.RegisterOutputType(AttributeToRoleIdentityMapperOutput{})
	pulumi.RegisterOutputType(AttributeToRoleIdentityMapperArrayOutput{})
	pulumi.RegisterOutputType(AttributeToRoleIdentityMapperMapOutput{})
}
