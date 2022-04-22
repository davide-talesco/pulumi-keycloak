// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing Keycloak client scopes that can be attached to clients that use the OpenID Connect protocol.
//
// Client Scopes can be used to share common protocol and role mappings between multiple clients within a realm. They can also
// be used by clients to conditionally request claims or roles for a user based on the OAuth 2.0 `scope` parameter.
//
// ## Example Usage
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
// 		_, err = openid.NewClientScope(ctx, "openidClientScope", &openid.ClientScopeArgs{
// 			RealmId:             realm.ID(),
// 			Description:         pulumi.String("When requested, this scope will map a user's group memberships to a claim"),
// 			IncludeInTokenScope: pulumi.Bool(true),
// 			GuiOrder:            pulumi.Int(1),
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
// Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `client_scope_id` is the unique ID that Keycloak assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID. Examplebash
//
// ```sh
//  $ pulumi import keycloak:openid/clientScope:ClientScope openid_client_scope my-realm/8e8f7fe1-df9b-40ed-bed3-4597aa0dac52
// ```
type ClientScope struct {
	pulumi.CustomResourceState

	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	ConsentScreenText pulumi.StringPtrOutput `pulumi:"consentScreenText"`
	// The description of this client scope in the GUI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	GuiOrder pulumi.IntPtrOutput `pulumi:"guiOrder"`
	// When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
	IncludeInTokenScope pulumi.BoolPtrOutput `pulumi:"includeInTokenScope"`
	// The display name of this client scope in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this client scope belongs to.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewClientScope registers a new resource with the given unique name, arguments, and options.
func NewClientScope(ctx *pulumi.Context,
	name string, args *ClientScopeArgs, opts ...pulumi.ResourceOption) (*ClientScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource ClientScope
	err := ctx.RegisterResource("keycloak:openid/clientScope:ClientScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientScope gets an existing ClientScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientScopeState, opts ...pulumi.ResourceOption) (*ClientScope, error) {
	var resource ClientScope
	err := ctx.ReadResource("keycloak:openid/clientScope:ClientScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientScope resources.
type clientScopeState struct {
	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	ConsentScreenText *string `pulumi:"consentScreenText"`
	// The description of this client scope in the GUI.
	Description *string `pulumi:"description"`
	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	GuiOrder *int `pulumi:"guiOrder"`
	// When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
	IncludeInTokenScope *bool `pulumi:"includeInTokenScope"`
	// The display name of this client scope in the GUI.
	Name *string `pulumi:"name"`
	// The realm this client scope belongs to.
	RealmId *string `pulumi:"realmId"`
}

type ClientScopeState struct {
	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	ConsentScreenText pulumi.StringPtrInput
	// The description of this client scope in the GUI.
	Description pulumi.StringPtrInput
	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	GuiOrder pulumi.IntPtrInput
	// When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
	IncludeInTokenScope pulumi.BoolPtrInput
	// The display name of this client scope in the GUI.
	Name pulumi.StringPtrInput
	// The realm this client scope belongs to.
	RealmId pulumi.StringPtrInput
}

func (ClientScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientScopeState)(nil)).Elem()
}

type clientScopeArgs struct {
	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	ConsentScreenText *string `pulumi:"consentScreenText"`
	// The description of this client scope in the GUI.
	Description *string `pulumi:"description"`
	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	GuiOrder *int `pulumi:"guiOrder"`
	// When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
	IncludeInTokenScope *bool `pulumi:"includeInTokenScope"`
	// The display name of this client scope in the GUI.
	Name *string `pulumi:"name"`
	// The realm this client scope belongs to.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a ClientScope resource.
type ClientScopeArgs struct {
	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	ConsentScreenText pulumi.StringPtrInput
	// The description of this client scope in the GUI.
	Description pulumi.StringPtrInput
	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	GuiOrder pulumi.IntPtrInput
	// When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
	IncludeInTokenScope pulumi.BoolPtrInput
	// The display name of this client scope in the GUI.
	Name pulumi.StringPtrInput
	// The realm this client scope belongs to.
	RealmId pulumi.StringInput
}

func (ClientScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientScopeArgs)(nil)).Elem()
}

type ClientScopeInput interface {
	pulumi.Input

	ToClientScopeOutput() ClientScopeOutput
	ToClientScopeOutputWithContext(ctx context.Context) ClientScopeOutput
}

func (*ClientScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientScope)(nil)).Elem()
}

func (i *ClientScope) ToClientScopeOutput() ClientScopeOutput {
	return i.ToClientScopeOutputWithContext(context.Background())
}

func (i *ClientScope) ToClientScopeOutputWithContext(ctx context.Context) ClientScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientScopeOutput)
}

// ClientScopeArrayInput is an input type that accepts ClientScopeArray and ClientScopeArrayOutput values.
// You can construct a concrete instance of `ClientScopeArrayInput` via:
//
//          ClientScopeArray{ ClientScopeArgs{...} }
type ClientScopeArrayInput interface {
	pulumi.Input

	ToClientScopeArrayOutput() ClientScopeArrayOutput
	ToClientScopeArrayOutputWithContext(context.Context) ClientScopeArrayOutput
}

type ClientScopeArray []ClientScopeInput

func (ClientScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientScope)(nil)).Elem()
}

func (i ClientScopeArray) ToClientScopeArrayOutput() ClientScopeArrayOutput {
	return i.ToClientScopeArrayOutputWithContext(context.Background())
}

func (i ClientScopeArray) ToClientScopeArrayOutputWithContext(ctx context.Context) ClientScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientScopeArrayOutput)
}

// ClientScopeMapInput is an input type that accepts ClientScopeMap and ClientScopeMapOutput values.
// You can construct a concrete instance of `ClientScopeMapInput` via:
//
//          ClientScopeMap{ "key": ClientScopeArgs{...} }
type ClientScopeMapInput interface {
	pulumi.Input

	ToClientScopeMapOutput() ClientScopeMapOutput
	ToClientScopeMapOutputWithContext(context.Context) ClientScopeMapOutput
}

type ClientScopeMap map[string]ClientScopeInput

func (ClientScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientScope)(nil)).Elem()
}

func (i ClientScopeMap) ToClientScopeMapOutput() ClientScopeMapOutput {
	return i.ToClientScopeMapOutputWithContext(context.Background())
}

func (i ClientScopeMap) ToClientScopeMapOutputWithContext(ctx context.Context) ClientScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientScopeMapOutput)
}

type ClientScopeOutput struct{ *pulumi.OutputState }

func (ClientScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientScope)(nil)).Elem()
}

func (o ClientScopeOutput) ToClientScopeOutput() ClientScopeOutput {
	return o
}

func (o ClientScopeOutput) ToClientScopeOutputWithContext(ctx context.Context) ClientScopeOutput {
	return o
}

type ClientScopeArrayOutput struct{ *pulumi.OutputState }

func (ClientScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientScope)(nil)).Elem()
}

func (o ClientScopeArrayOutput) ToClientScopeArrayOutput() ClientScopeArrayOutput {
	return o
}

func (o ClientScopeArrayOutput) ToClientScopeArrayOutputWithContext(ctx context.Context) ClientScopeArrayOutput {
	return o
}

func (o ClientScopeArrayOutput) Index(i pulumi.IntInput) ClientScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientScope {
		return vs[0].([]*ClientScope)[vs[1].(int)]
	}).(ClientScopeOutput)
}

type ClientScopeMapOutput struct{ *pulumi.OutputState }

func (ClientScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientScope)(nil)).Elem()
}

func (o ClientScopeMapOutput) ToClientScopeMapOutput() ClientScopeMapOutput {
	return o
}

func (o ClientScopeMapOutput) ToClientScopeMapOutputWithContext(ctx context.Context) ClientScopeMapOutput {
	return o
}

func (o ClientScopeMapOutput) MapIndex(k pulumi.StringInput) ClientScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientScope {
		return vs[0].(map[string]*ClientScope)[vs[1].(string)]
	}).(ClientScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientScopeInput)(nil)).Elem(), &ClientScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientScopeArrayInput)(nil)).Elem(), ClientScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientScopeMapInput)(nil)).Elem(), ClientScopeMap{})
	pulumi.RegisterOutputType(ClientScopeOutput{})
	pulumi.RegisterOutputType(ClientScopeArrayOutput{})
	pulumi.RegisterOutputType(ClientScopeMapOutput{})
}
