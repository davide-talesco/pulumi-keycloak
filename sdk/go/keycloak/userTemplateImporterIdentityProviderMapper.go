// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing an username template importer identity provider mapper within Keycloak.
//
// The username template importer mapper can be used to map externally defined OIDC claims or SAML attributes with a template to the username of the imported Keycloak user:
//
// - Substitutions are enclosed in \${}. For example: '\${ALIAS}.\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
//
// > If you are using Keycloak 10 or higher, you will need to specify the `extraConfig` argument in order to define a `syncMode` for the mapper.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
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
// 		oidc, err := oidc.NewIdentityProvider(ctx, "oidc", &oidc.IdentityProviderArgs{
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
// 		_, err = keycloak.NewUserTemplateImporterIdentityProviderMapper(ctx, "usernameImporter", &keycloak.UserTemplateImporterIdentityProviderMapperArgs{
// 			Realm:                 realm.ID(),
// 			IdentityProviderAlias: oidc.Alias,
// 			Template:              pulumi.String(fmt.Sprintf("%v%v%v", ALIAS, ".", CLAIM.Email)),
// 			ExtraConfig: pulumi.AnyMap{
// 				"syncMode": pulumi.Any("INHERIT"),
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
//  $ pulumi import keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper username_importer my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
// ```
type UserTemplateImporterIdentityProviderMapper struct {
	pulumi.CustomResourceState

	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapOutput `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the realm.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
	Template pulumi.StringPtrOutput `pulumi:"template"`
}

// NewUserTemplateImporterIdentityProviderMapper registers a new resource with the given unique name, arguments, and options.
func NewUserTemplateImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, args *UserTemplateImporterIdentityProviderMapperArgs, opts ...pulumi.ResourceOption) (*UserTemplateImporterIdentityProviderMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	var resource UserTemplateImporterIdentityProviderMapper
	err := ctx.RegisterResource("keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserTemplateImporterIdentityProviderMapper gets an existing UserTemplateImporterIdentityProviderMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserTemplateImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserTemplateImporterIdentityProviderMapperState, opts ...pulumi.ResourceOption) (*UserTemplateImporterIdentityProviderMapper, error) {
	var resource UserTemplateImporterIdentityProviderMapper
	err := ctx.ReadResource("keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserTemplateImporterIdentityProviderMapper resources.
type userTemplateImporterIdentityProviderMapperState struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm *string `pulumi:"realm"`
	// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
	Template *string `pulumi:"template"`
}

type UserTemplateImporterIdentityProviderMapperState struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringPtrInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringPtrInput
	// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
	Template pulumi.StringPtrInput
}

func (UserTemplateImporterIdentityProviderMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userTemplateImporterIdentityProviderMapperState)(nil)).Elem()
}

type userTemplateImporterIdentityProviderMapperArgs struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm string `pulumi:"realm"`
	// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
	Template *string `pulumi:"template"`
}

// The set of arguments for constructing a UserTemplateImporterIdentityProviderMapper resource.
type UserTemplateImporterIdentityProviderMapperArgs struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringInput
	// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
	Template pulumi.StringPtrInput
}

func (UserTemplateImporterIdentityProviderMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userTemplateImporterIdentityProviderMapperArgs)(nil)).Elem()
}

type UserTemplateImporterIdentityProviderMapperInput interface {
	pulumi.Input

	ToUserTemplateImporterIdentityProviderMapperOutput() UserTemplateImporterIdentityProviderMapperOutput
	ToUserTemplateImporterIdentityProviderMapperOutputWithContext(ctx context.Context) UserTemplateImporterIdentityProviderMapperOutput
}

func (*UserTemplateImporterIdentityProviderMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**UserTemplateImporterIdentityProviderMapper)(nil)).Elem()
}

func (i *UserTemplateImporterIdentityProviderMapper) ToUserTemplateImporterIdentityProviderMapperOutput() UserTemplateImporterIdentityProviderMapperOutput {
	return i.ToUserTemplateImporterIdentityProviderMapperOutputWithContext(context.Background())
}

func (i *UserTemplateImporterIdentityProviderMapper) ToUserTemplateImporterIdentityProviderMapperOutputWithContext(ctx context.Context) UserTemplateImporterIdentityProviderMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserTemplateImporterIdentityProviderMapperOutput)
}

// UserTemplateImporterIdentityProviderMapperArrayInput is an input type that accepts UserTemplateImporterIdentityProviderMapperArray and UserTemplateImporterIdentityProviderMapperArrayOutput values.
// You can construct a concrete instance of `UserTemplateImporterIdentityProviderMapperArrayInput` via:
//
//          UserTemplateImporterIdentityProviderMapperArray{ UserTemplateImporterIdentityProviderMapperArgs{...} }
type UserTemplateImporterIdentityProviderMapperArrayInput interface {
	pulumi.Input

	ToUserTemplateImporterIdentityProviderMapperArrayOutput() UserTemplateImporterIdentityProviderMapperArrayOutput
	ToUserTemplateImporterIdentityProviderMapperArrayOutputWithContext(context.Context) UserTemplateImporterIdentityProviderMapperArrayOutput
}

type UserTemplateImporterIdentityProviderMapperArray []UserTemplateImporterIdentityProviderMapperInput

func (UserTemplateImporterIdentityProviderMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserTemplateImporterIdentityProviderMapper)(nil)).Elem()
}

func (i UserTemplateImporterIdentityProviderMapperArray) ToUserTemplateImporterIdentityProviderMapperArrayOutput() UserTemplateImporterIdentityProviderMapperArrayOutput {
	return i.ToUserTemplateImporterIdentityProviderMapperArrayOutputWithContext(context.Background())
}

func (i UserTemplateImporterIdentityProviderMapperArray) ToUserTemplateImporterIdentityProviderMapperArrayOutputWithContext(ctx context.Context) UserTemplateImporterIdentityProviderMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserTemplateImporterIdentityProviderMapperArrayOutput)
}

// UserTemplateImporterIdentityProviderMapperMapInput is an input type that accepts UserTemplateImporterIdentityProviderMapperMap and UserTemplateImporterIdentityProviderMapperMapOutput values.
// You can construct a concrete instance of `UserTemplateImporterIdentityProviderMapperMapInput` via:
//
//          UserTemplateImporterIdentityProviderMapperMap{ "key": UserTemplateImporterIdentityProviderMapperArgs{...} }
type UserTemplateImporterIdentityProviderMapperMapInput interface {
	pulumi.Input

	ToUserTemplateImporterIdentityProviderMapperMapOutput() UserTemplateImporterIdentityProviderMapperMapOutput
	ToUserTemplateImporterIdentityProviderMapperMapOutputWithContext(context.Context) UserTemplateImporterIdentityProviderMapperMapOutput
}

type UserTemplateImporterIdentityProviderMapperMap map[string]UserTemplateImporterIdentityProviderMapperInput

func (UserTemplateImporterIdentityProviderMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserTemplateImporterIdentityProviderMapper)(nil)).Elem()
}

func (i UserTemplateImporterIdentityProviderMapperMap) ToUserTemplateImporterIdentityProviderMapperMapOutput() UserTemplateImporterIdentityProviderMapperMapOutput {
	return i.ToUserTemplateImporterIdentityProviderMapperMapOutputWithContext(context.Background())
}

func (i UserTemplateImporterIdentityProviderMapperMap) ToUserTemplateImporterIdentityProviderMapperMapOutputWithContext(ctx context.Context) UserTemplateImporterIdentityProviderMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserTemplateImporterIdentityProviderMapperMapOutput)
}

type UserTemplateImporterIdentityProviderMapperOutput struct{ *pulumi.OutputState }

func (UserTemplateImporterIdentityProviderMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserTemplateImporterIdentityProviderMapper)(nil)).Elem()
}

func (o UserTemplateImporterIdentityProviderMapperOutput) ToUserTemplateImporterIdentityProviderMapperOutput() UserTemplateImporterIdentityProviderMapperOutput {
	return o
}

func (o UserTemplateImporterIdentityProviderMapperOutput) ToUserTemplateImporterIdentityProviderMapperOutputWithContext(ctx context.Context) UserTemplateImporterIdentityProviderMapperOutput {
	return o
}

type UserTemplateImporterIdentityProviderMapperArrayOutput struct{ *pulumi.OutputState }

func (UserTemplateImporterIdentityProviderMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserTemplateImporterIdentityProviderMapper)(nil)).Elem()
}

func (o UserTemplateImporterIdentityProviderMapperArrayOutput) ToUserTemplateImporterIdentityProviderMapperArrayOutput() UserTemplateImporterIdentityProviderMapperArrayOutput {
	return o
}

func (o UserTemplateImporterIdentityProviderMapperArrayOutput) ToUserTemplateImporterIdentityProviderMapperArrayOutputWithContext(ctx context.Context) UserTemplateImporterIdentityProviderMapperArrayOutput {
	return o
}

func (o UserTemplateImporterIdentityProviderMapperArrayOutput) Index(i pulumi.IntInput) UserTemplateImporterIdentityProviderMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserTemplateImporterIdentityProviderMapper {
		return vs[0].([]*UserTemplateImporterIdentityProviderMapper)[vs[1].(int)]
	}).(UserTemplateImporterIdentityProviderMapperOutput)
}

type UserTemplateImporterIdentityProviderMapperMapOutput struct{ *pulumi.OutputState }

func (UserTemplateImporterIdentityProviderMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserTemplateImporterIdentityProviderMapper)(nil)).Elem()
}

func (o UserTemplateImporterIdentityProviderMapperMapOutput) ToUserTemplateImporterIdentityProviderMapperMapOutput() UserTemplateImporterIdentityProviderMapperMapOutput {
	return o
}

func (o UserTemplateImporterIdentityProviderMapperMapOutput) ToUserTemplateImporterIdentityProviderMapperMapOutputWithContext(ctx context.Context) UserTemplateImporterIdentityProviderMapperMapOutput {
	return o
}

func (o UserTemplateImporterIdentityProviderMapperMapOutput) MapIndex(k pulumi.StringInput) UserTemplateImporterIdentityProviderMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserTemplateImporterIdentityProviderMapper {
		return vs[0].(map[string]*UserTemplateImporterIdentityProviderMapper)[vs[1].(string)]
	}).(UserTemplateImporterIdentityProviderMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserTemplateImporterIdentityProviderMapperInput)(nil)).Elem(), &UserTemplateImporterIdentityProviderMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserTemplateImporterIdentityProviderMapperArrayInput)(nil)).Elem(), UserTemplateImporterIdentityProviderMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserTemplateImporterIdentityProviderMapperMapInput)(nil)).Elem(), UserTemplateImporterIdentityProviderMapperMap{})
	pulumi.RegisterOutputType(UserTemplateImporterIdentityProviderMapperOutput{})
	pulumi.RegisterOutputType(UserTemplateImporterIdentityProviderMapperArrayOutput{})
	pulumi.RegisterOutputType(UserTemplateImporterIdentityProviderMapperMapOutput{})
}
