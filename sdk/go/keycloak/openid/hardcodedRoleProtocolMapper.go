// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing hardcoded role protocol mappers within Keycloak.
//
// Hardcoded role protocol mappers allow you to specify a single role to always map to an access token for a client.
//
// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
// multiple different clients.
//
// ## Example Usage
// ### Client)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
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
// 		role, err := keycloak.NewRole(ctx, "role", &keycloak.RoleArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		openidClient, err := openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
// 			RealmId:    realm.ID(),
// 			ClientId:   pulumi.String("client"),
// 			Enabled:    pulumi.Bool(true),
// 			AccessType: pulumi.String("CONFIDENTIAL"),
// 			ValidRedirectUris: pulumi.StringArray{
// 				pulumi.String("http://localhost:8080/openid-callback"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewHardcodedRoleProtocolMapper(ctx, "hardcodedRoleMapper", &openid.HardcodedRoleProtocolMapperArgs{
// 			RealmId:  realm.ID(),
// 			ClientId: openidClient.ID(),
// 			RoleId:   role.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Client Scope)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
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
// 		role, err := keycloak.NewRole(ctx, "role", &keycloak.RoleArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewHardcodedRoleProtocolMapper(ctx, "hardcodedRoleMapper", &openid.HardcodedRoleProtocolMapperArgs{
// 			RealmId:       realm.ID(),
// 			ClientScopeId: clientScope.ID(),
// 			RoleId:        role.ID(),
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
// Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
//
// ```sh
//  $ pulumi import keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper hardcoded_role_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
//
// ```sh
//  $ pulumi import keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper hardcoded_role_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
type HardcodedRoleProtocolMapper struct {
	pulumi.CustomResourceState

	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The ID of the role to map to an access token.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewHardcodedRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *HardcodedRoleProtocolMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedRoleProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	var resource HardcodedRoleProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedRoleProtocolMapper gets an existing HardcodedRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedRoleProtocolMapperState, opts ...pulumi.ResourceOption) (*HardcodedRoleProtocolMapper, error) {
	var resource HardcodedRoleProtocolMapper
	err := ctx.ReadResource("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedRoleProtocolMapper resources.
type hardcodedRoleProtocolMapperState struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
	// The ID of the role to map to an access token.
	RoleId *string `pulumi:"roleId"`
}

type HardcodedRoleProtocolMapperState struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
	// The ID of the role to map to an access token.
	RoleId pulumi.StringPtrInput
}

func (HardcodedRoleProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleProtocolMapperState)(nil)).Elem()
}

type hardcodedRoleProtocolMapperArgs struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
	// The ID of the role to map to an access token.
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a HardcodedRoleProtocolMapper resource.
type HardcodedRoleProtocolMapperArgs struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
	// The ID of the role to map to an access token.
	RoleId pulumi.StringInput
}

func (HardcodedRoleProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleProtocolMapperArgs)(nil)).Elem()
}

type HardcodedRoleProtocolMapperInput interface {
	pulumi.Input

	ToHardcodedRoleProtocolMapperOutput() HardcodedRoleProtocolMapperOutput
	ToHardcodedRoleProtocolMapperOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperOutput
}

func (*HardcodedRoleProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((*HardcodedRoleProtocolMapper)(nil))
}

func (i *HardcodedRoleProtocolMapper) ToHardcodedRoleProtocolMapperOutput() HardcodedRoleProtocolMapperOutput {
	return i.ToHardcodedRoleProtocolMapperOutputWithContext(context.Background())
}

func (i *HardcodedRoleProtocolMapper) ToHardcodedRoleProtocolMapperOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleProtocolMapperOutput)
}

func (i *HardcodedRoleProtocolMapper) ToHardcodedRoleProtocolMapperPtrOutput() HardcodedRoleProtocolMapperPtrOutput {
	return i.ToHardcodedRoleProtocolMapperPtrOutputWithContext(context.Background())
}

func (i *HardcodedRoleProtocolMapper) ToHardcodedRoleProtocolMapperPtrOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleProtocolMapperPtrOutput)
}

type HardcodedRoleProtocolMapperPtrInput interface {
	pulumi.Input

	ToHardcodedRoleProtocolMapperPtrOutput() HardcodedRoleProtocolMapperPtrOutput
	ToHardcodedRoleProtocolMapperPtrOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperPtrOutput
}

type hardcodedRoleProtocolMapperPtrType HardcodedRoleProtocolMapperArgs

func (*hardcodedRoleProtocolMapperPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedRoleProtocolMapper)(nil))
}

func (i *hardcodedRoleProtocolMapperPtrType) ToHardcodedRoleProtocolMapperPtrOutput() HardcodedRoleProtocolMapperPtrOutput {
	return i.ToHardcodedRoleProtocolMapperPtrOutputWithContext(context.Background())
}

func (i *hardcodedRoleProtocolMapperPtrType) ToHardcodedRoleProtocolMapperPtrOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleProtocolMapperPtrOutput)
}

// HardcodedRoleProtocolMapperArrayInput is an input type that accepts HardcodedRoleProtocolMapperArray and HardcodedRoleProtocolMapperArrayOutput values.
// You can construct a concrete instance of `HardcodedRoleProtocolMapperArrayInput` via:
//
//          HardcodedRoleProtocolMapperArray{ HardcodedRoleProtocolMapperArgs{...} }
type HardcodedRoleProtocolMapperArrayInput interface {
	pulumi.Input

	ToHardcodedRoleProtocolMapperArrayOutput() HardcodedRoleProtocolMapperArrayOutput
	ToHardcodedRoleProtocolMapperArrayOutputWithContext(context.Context) HardcodedRoleProtocolMapperArrayOutput
}

type HardcodedRoleProtocolMapperArray []HardcodedRoleProtocolMapperInput

func (HardcodedRoleProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedRoleProtocolMapper)(nil)).Elem()
}

func (i HardcodedRoleProtocolMapperArray) ToHardcodedRoleProtocolMapperArrayOutput() HardcodedRoleProtocolMapperArrayOutput {
	return i.ToHardcodedRoleProtocolMapperArrayOutputWithContext(context.Background())
}

func (i HardcodedRoleProtocolMapperArray) ToHardcodedRoleProtocolMapperArrayOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleProtocolMapperArrayOutput)
}

// HardcodedRoleProtocolMapperMapInput is an input type that accepts HardcodedRoleProtocolMapperMap and HardcodedRoleProtocolMapperMapOutput values.
// You can construct a concrete instance of `HardcodedRoleProtocolMapperMapInput` via:
//
//          HardcodedRoleProtocolMapperMap{ "key": HardcodedRoleProtocolMapperArgs{...} }
type HardcodedRoleProtocolMapperMapInput interface {
	pulumi.Input

	ToHardcodedRoleProtocolMapperMapOutput() HardcodedRoleProtocolMapperMapOutput
	ToHardcodedRoleProtocolMapperMapOutputWithContext(context.Context) HardcodedRoleProtocolMapperMapOutput
}

type HardcodedRoleProtocolMapperMap map[string]HardcodedRoleProtocolMapperInput

func (HardcodedRoleProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedRoleProtocolMapper)(nil)).Elem()
}

func (i HardcodedRoleProtocolMapperMap) ToHardcodedRoleProtocolMapperMapOutput() HardcodedRoleProtocolMapperMapOutput {
	return i.ToHardcodedRoleProtocolMapperMapOutputWithContext(context.Background())
}

func (i HardcodedRoleProtocolMapperMap) ToHardcodedRoleProtocolMapperMapOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleProtocolMapperMapOutput)
}

type HardcodedRoleProtocolMapperOutput struct{ *pulumi.OutputState }

func (HardcodedRoleProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardcodedRoleProtocolMapper)(nil))
}

func (o HardcodedRoleProtocolMapperOutput) ToHardcodedRoleProtocolMapperOutput() HardcodedRoleProtocolMapperOutput {
	return o
}

func (o HardcodedRoleProtocolMapperOutput) ToHardcodedRoleProtocolMapperOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperOutput {
	return o
}

func (o HardcodedRoleProtocolMapperOutput) ToHardcodedRoleProtocolMapperPtrOutput() HardcodedRoleProtocolMapperPtrOutput {
	return o.ToHardcodedRoleProtocolMapperPtrOutputWithContext(context.Background())
}

func (o HardcodedRoleProtocolMapperOutput) ToHardcodedRoleProtocolMapperPtrOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HardcodedRoleProtocolMapper) *HardcodedRoleProtocolMapper {
		return &v
	}).(HardcodedRoleProtocolMapperPtrOutput)
}

type HardcodedRoleProtocolMapperPtrOutput struct{ *pulumi.OutputState }

func (HardcodedRoleProtocolMapperPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedRoleProtocolMapper)(nil))
}

func (o HardcodedRoleProtocolMapperPtrOutput) ToHardcodedRoleProtocolMapperPtrOutput() HardcodedRoleProtocolMapperPtrOutput {
	return o
}

func (o HardcodedRoleProtocolMapperPtrOutput) ToHardcodedRoleProtocolMapperPtrOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperPtrOutput {
	return o
}

func (o HardcodedRoleProtocolMapperPtrOutput) Elem() HardcodedRoleProtocolMapperOutput {
	return o.ApplyT(func(v *HardcodedRoleProtocolMapper) HardcodedRoleProtocolMapper {
		if v != nil {
			return *v
		}
		var ret HardcodedRoleProtocolMapper
		return ret
	}).(HardcodedRoleProtocolMapperOutput)
}

type HardcodedRoleProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (HardcodedRoleProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HardcodedRoleProtocolMapper)(nil))
}

func (o HardcodedRoleProtocolMapperArrayOutput) ToHardcodedRoleProtocolMapperArrayOutput() HardcodedRoleProtocolMapperArrayOutput {
	return o
}

func (o HardcodedRoleProtocolMapperArrayOutput) ToHardcodedRoleProtocolMapperArrayOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperArrayOutput {
	return o
}

func (o HardcodedRoleProtocolMapperArrayOutput) Index(i pulumi.IntInput) HardcodedRoleProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HardcodedRoleProtocolMapper {
		return vs[0].([]HardcodedRoleProtocolMapper)[vs[1].(int)]
	}).(HardcodedRoleProtocolMapperOutput)
}

type HardcodedRoleProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (HardcodedRoleProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HardcodedRoleProtocolMapper)(nil))
}

func (o HardcodedRoleProtocolMapperMapOutput) ToHardcodedRoleProtocolMapperMapOutput() HardcodedRoleProtocolMapperMapOutput {
	return o
}

func (o HardcodedRoleProtocolMapperMapOutput) ToHardcodedRoleProtocolMapperMapOutputWithContext(ctx context.Context) HardcodedRoleProtocolMapperMapOutput {
	return o
}

func (o HardcodedRoleProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) HardcodedRoleProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HardcodedRoleProtocolMapper {
		return vs[0].(map[string]HardcodedRoleProtocolMapper)[vs[1].(string)]
	}).(HardcodedRoleProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleProtocolMapperInput)(nil)).Elem(), &HardcodedRoleProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleProtocolMapperPtrInput)(nil)).Elem(), &HardcodedRoleProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleProtocolMapperArrayInput)(nil)).Elem(), HardcodedRoleProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleProtocolMapperMapInput)(nil)).Elem(), HardcodedRoleProtocolMapperMap{})
	pulumi.RegisterOutputType(HardcodedRoleProtocolMapperOutput{})
	pulumi.RegisterOutputType(HardcodedRoleProtocolMapperPtrOutput{})
	pulumi.RegisterOutputType(HardcodedRoleProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(HardcodedRoleProtocolMapperMapOutput{})
}
