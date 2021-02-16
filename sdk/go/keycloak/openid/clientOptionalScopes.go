// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server.
type ClientOptionalScopes struct {
	pulumi.CustomResourceState

	// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// An array of client scope names to attach to this client as optional scopes.
	OptionalScopes pulumi.StringArrayOutput `pulumi:"optionalScopes"`
	// The realm this client and scopes exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewClientOptionalScopes registers a new resource with the given unique name, arguments, and options.
func NewClientOptionalScopes(ctx *pulumi.Context,
	name string, args *ClientOptionalScopesArgs, opts ...pulumi.ResourceOption) (*ClientOptionalScopes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.OptionalScopes == nil {
		return nil, errors.New("invalid value for required argument 'OptionalScopes'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource ClientOptionalScopes
	err := ctx.RegisterResource("keycloak:openid/clientOptionalScopes:ClientOptionalScopes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientOptionalScopes gets an existing ClientOptionalScopes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientOptionalScopes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientOptionalScopesState, opts ...pulumi.ResourceOption) (*ClientOptionalScopes, error) {
	var resource ClientOptionalScopes
	err := ctx.ReadResource("keycloak:openid/clientOptionalScopes:ClientOptionalScopes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientOptionalScopes resources.
type clientOptionalScopesState struct {
	// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId *string `pulumi:"clientId"`
	// An array of client scope names to attach to this client as optional scopes.
	OptionalScopes []string `pulumi:"optionalScopes"`
	// The realm this client and scopes exists in.
	RealmId *string `pulumi:"realmId"`
}

type ClientOptionalScopesState struct {
	// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringPtrInput
	// An array of client scope names to attach to this client as optional scopes.
	OptionalScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringPtrInput
}

func (ClientOptionalScopesState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientOptionalScopesState)(nil)).Elem()
}

type clientOptionalScopesArgs struct {
	// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId string `pulumi:"clientId"`
	// An array of client scope names to attach to this client as optional scopes.
	OptionalScopes []string `pulumi:"optionalScopes"`
	// The realm this client and scopes exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a ClientOptionalScopes resource.
type ClientOptionalScopesArgs struct {
	// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringInput
	// An array of client scope names to attach to this client as optional scopes.
	OptionalScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringInput
}

func (ClientOptionalScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientOptionalScopesArgs)(nil)).Elem()
}

type ClientOptionalScopesInput interface {
	pulumi.Input

	ToClientOptionalScopesOutput() ClientOptionalScopesOutput
	ToClientOptionalScopesOutputWithContext(ctx context.Context) ClientOptionalScopesOutput
}

func (*ClientOptionalScopes) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientOptionalScopes)(nil))
}

func (i *ClientOptionalScopes) ToClientOptionalScopesOutput() ClientOptionalScopesOutput {
	return i.ToClientOptionalScopesOutputWithContext(context.Background())
}

func (i *ClientOptionalScopes) ToClientOptionalScopesOutputWithContext(ctx context.Context) ClientOptionalScopesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOptionalScopesOutput)
}

func (i *ClientOptionalScopes) ToClientOptionalScopesPtrOutput() ClientOptionalScopesPtrOutput {
	return i.ToClientOptionalScopesPtrOutputWithContext(context.Background())
}

func (i *ClientOptionalScopes) ToClientOptionalScopesPtrOutputWithContext(ctx context.Context) ClientOptionalScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOptionalScopesPtrOutput)
}

type ClientOptionalScopesPtrInput interface {
	pulumi.Input

	ToClientOptionalScopesPtrOutput() ClientOptionalScopesPtrOutput
	ToClientOptionalScopesPtrOutputWithContext(ctx context.Context) ClientOptionalScopesPtrOutput
}

type clientOptionalScopesPtrType ClientOptionalScopesArgs

func (*clientOptionalScopesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientOptionalScopes)(nil))
}

func (i *clientOptionalScopesPtrType) ToClientOptionalScopesPtrOutput() ClientOptionalScopesPtrOutput {
	return i.ToClientOptionalScopesPtrOutputWithContext(context.Background())
}

func (i *clientOptionalScopesPtrType) ToClientOptionalScopesPtrOutputWithContext(ctx context.Context) ClientOptionalScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOptionalScopesPtrOutput)
}

// ClientOptionalScopesArrayInput is an input type that accepts ClientOptionalScopesArray and ClientOptionalScopesArrayOutput values.
// You can construct a concrete instance of `ClientOptionalScopesArrayInput` via:
//
//          ClientOptionalScopesArray{ ClientOptionalScopesArgs{...} }
type ClientOptionalScopesArrayInput interface {
	pulumi.Input

	ToClientOptionalScopesArrayOutput() ClientOptionalScopesArrayOutput
	ToClientOptionalScopesArrayOutputWithContext(context.Context) ClientOptionalScopesArrayOutput
}

type ClientOptionalScopesArray []ClientOptionalScopesInput

func (ClientOptionalScopesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClientOptionalScopes)(nil))
}

func (i ClientOptionalScopesArray) ToClientOptionalScopesArrayOutput() ClientOptionalScopesArrayOutput {
	return i.ToClientOptionalScopesArrayOutputWithContext(context.Background())
}

func (i ClientOptionalScopesArray) ToClientOptionalScopesArrayOutputWithContext(ctx context.Context) ClientOptionalScopesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOptionalScopesArrayOutput)
}

// ClientOptionalScopesMapInput is an input type that accepts ClientOptionalScopesMap and ClientOptionalScopesMapOutput values.
// You can construct a concrete instance of `ClientOptionalScopesMapInput` via:
//
//          ClientOptionalScopesMap{ "key": ClientOptionalScopesArgs{...} }
type ClientOptionalScopesMapInput interface {
	pulumi.Input

	ToClientOptionalScopesMapOutput() ClientOptionalScopesMapOutput
	ToClientOptionalScopesMapOutputWithContext(context.Context) ClientOptionalScopesMapOutput
}

type ClientOptionalScopesMap map[string]ClientOptionalScopesInput

func (ClientOptionalScopesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClientOptionalScopes)(nil))
}

func (i ClientOptionalScopesMap) ToClientOptionalScopesMapOutput() ClientOptionalScopesMapOutput {
	return i.ToClientOptionalScopesMapOutputWithContext(context.Background())
}

func (i ClientOptionalScopesMap) ToClientOptionalScopesMapOutputWithContext(ctx context.Context) ClientOptionalScopesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOptionalScopesMapOutput)
}

type ClientOptionalScopesOutput struct {
	*pulumi.OutputState
}

func (ClientOptionalScopesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientOptionalScopes)(nil))
}

func (o ClientOptionalScopesOutput) ToClientOptionalScopesOutput() ClientOptionalScopesOutput {
	return o
}

func (o ClientOptionalScopesOutput) ToClientOptionalScopesOutputWithContext(ctx context.Context) ClientOptionalScopesOutput {
	return o
}

func (o ClientOptionalScopesOutput) ToClientOptionalScopesPtrOutput() ClientOptionalScopesPtrOutput {
	return o.ToClientOptionalScopesPtrOutputWithContext(context.Background())
}

func (o ClientOptionalScopesOutput) ToClientOptionalScopesPtrOutputWithContext(ctx context.Context) ClientOptionalScopesPtrOutput {
	return o.ApplyT(func(v ClientOptionalScopes) *ClientOptionalScopes {
		return &v
	}).(ClientOptionalScopesPtrOutput)
}

type ClientOptionalScopesPtrOutput struct {
	*pulumi.OutputState
}

func (ClientOptionalScopesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientOptionalScopes)(nil))
}

func (o ClientOptionalScopesPtrOutput) ToClientOptionalScopesPtrOutput() ClientOptionalScopesPtrOutput {
	return o
}

func (o ClientOptionalScopesPtrOutput) ToClientOptionalScopesPtrOutputWithContext(ctx context.Context) ClientOptionalScopesPtrOutput {
	return o
}

type ClientOptionalScopesArrayOutput struct{ *pulumi.OutputState }

func (ClientOptionalScopesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientOptionalScopes)(nil))
}

func (o ClientOptionalScopesArrayOutput) ToClientOptionalScopesArrayOutput() ClientOptionalScopesArrayOutput {
	return o
}

func (o ClientOptionalScopesArrayOutput) ToClientOptionalScopesArrayOutputWithContext(ctx context.Context) ClientOptionalScopesArrayOutput {
	return o
}

func (o ClientOptionalScopesArrayOutput) Index(i pulumi.IntInput) ClientOptionalScopesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientOptionalScopes {
		return vs[0].([]ClientOptionalScopes)[vs[1].(int)]
	}).(ClientOptionalScopesOutput)
}

type ClientOptionalScopesMapOutput struct{ *pulumi.OutputState }

func (ClientOptionalScopesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClientOptionalScopes)(nil))
}

func (o ClientOptionalScopesMapOutput) ToClientOptionalScopesMapOutput() ClientOptionalScopesMapOutput {
	return o
}

func (o ClientOptionalScopesMapOutput) ToClientOptionalScopesMapOutputWithContext(ctx context.Context) ClientOptionalScopesMapOutput {
	return o
}

func (o ClientOptionalScopesMapOutput) MapIndex(k pulumi.StringInput) ClientOptionalScopesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClientOptionalScopes {
		return vs[0].(map[string]ClientOptionalScopes)[vs[1].(string)]
	}).(ClientOptionalScopesOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientOptionalScopesOutput{})
	pulumi.RegisterOutputType(ClientOptionalScopesPtrOutput{})
	pulumi.RegisterOutputType(ClientOptionalScopesArrayOutput{})
	pulumi.RegisterOutputType(ClientOptionalScopesMapOutput{})
}
