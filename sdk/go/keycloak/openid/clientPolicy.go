// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientPolicy struct {
	pulumi.CustomResourceState

	Clients          pulumi.StringArrayOutput `pulumi:"clients"`
	DecisionStrategy pulumi.StringPtrOutput   `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Logic            pulumi.StringPtrOutput   `pulumi:"logic"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	RealmId          pulumi.StringOutput      `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput      `pulumi:"resourceServerId"`
}

// NewClientPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientPolicy(ctx *pulumi.Context,
	name string, args *ClientPolicyArgs, opts ...pulumi.ResourceOption) (*ClientPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Clients == nil {
		return nil, errors.New("invalid value for required argument 'Clients'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.ResourceServerId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServerId'")
	}
	var resource ClientPolicy
	err := ctx.RegisterResource("keycloak:openid/clientPolicy:ClientPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientPolicy gets an existing ClientPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientPolicyState, opts ...pulumi.ResourceOption) (*ClientPolicy, error) {
	var resource ClientPolicy
	err := ctx.ReadResource("keycloak:openid/clientPolicy:ClientPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientPolicy resources.
type clientPolicyState struct {
	Clients          []string `pulumi:"clients"`
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Logic            *string  `pulumi:"logic"`
	Name             *string  `pulumi:"name"`
	RealmId          *string  `pulumi:"realmId"`
	ResourceServerId *string  `pulumi:"resourceServerId"`
}

type ClientPolicyState struct {
	Clients          pulumi.StringArrayInput
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
}

func (ClientPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientPolicyState)(nil)).Elem()
}

type clientPolicyArgs struct {
	Clients          []string `pulumi:"clients"`
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Logic            *string  `pulumi:"logic"`
	Name             *string  `pulumi:"name"`
	RealmId          string   `pulumi:"realmId"`
	ResourceServerId string   `pulumi:"resourceServerId"`
}

// The set of arguments for constructing a ClientPolicy resource.
type ClientPolicyArgs struct {
	Clients          pulumi.StringArrayInput
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
}

func (ClientPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientPolicyArgs)(nil)).Elem()
}

type ClientPolicyInput interface {
	pulumi.Input

	ToClientPolicyOutput() ClientPolicyOutput
	ToClientPolicyOutputWithContext(ctx context.Context) ClientPolicyOutput
}

func (*ClientPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientPolicy)(nil)).Elem()
}

func (i *ClientPolicy) ToClientPolicyOutput() ClientPolicyOutput {
	return i.ToClientPolicyOutputWithContext(context.Background())
}

func (i *ClientPolicy) ToClientPolicyOutputWithContext(ctx context.Context) ClientPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPolicyOutput)
}

// ClientPolicyArrayInput is an input type that accepts ClientPolicyArray and ClientPolicyArrayOutput values.
// You can construct a concrete instance of `ClientPolicyArrayInput` via:
//
//          ClientPolicyArray{ ClientPolicyArgs{...} }
type ClientPolicyArrayInput interface {
	pulumi.Input

	ToClientPolicyArrayOutput() ClientPolicyArrayOutput
	ToClientPolicyArrayOutputWithContext(context.Context) ClientPolicyArrayOutput
}

type ClientPolicyArray []ClientPolicyInput

func (ClientPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientPolicy)(nil)).Elem()
}

func (i ClientPolicyArray) ToClientPolicyArrayOutput() ClientPolicyArrayOutput {
	return i.ToClientPolicyArrayOutputWithContext(context.Background())
}

func (i ClientPolicyArray) ToClientPolicyArrayOutputWithContext(ctx context.Context) ClientPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPolicyArrayOutput)
}

// ClientPolicyMapInput is an input type that accepts ClientPolicyMap and ClientPolicyMapOutput values.
// You can construct a concrete instance of `ClientPolicyMapInput` via:
//
//          ClientPolicyMap{ "key": ClientPolicyArgs{...} }
type ClientPolicyMapInput interface {
	pulumi.Input

	ToClientPolicyMapOutput() ClientPolicyMapOutput
	ToClientPolicyMapOutputWithContext(context.Context) ClientPolicyMapOutput
}

type ClientPolicyMap map[string]ClientPolicyInput

func (ClientPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientPolicy)(nil)).Elem()
}

func (i ClientPolicyMap) ToClientPolicyMapOutput() ClientPolicyMapOutput {
	return i.ToClientPolicyMapOutputWithContext(context.Background())
}

func (i ClientPolicyMap) ToClientPolicyMapOutputWithContext(ctx context.Context) ClientPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPolicyMapOutput)
}

type ClientPolicyOutput struct{ *pulumi.OutputState }

func (ClientPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientPolicy)(nil)).Elem()
}

func (o ClientPolicyOutput) ToClientPolicyOutput() ClientPolicyOutput {
	return o
}

func (o ClientPolicyOutput) ToClientPolicyOutputWithContext(ctx context.Context) ClientPolicyOutput {
	return o
}

type ClientPolicyArrayOutput struct{ *pulumi.OutputState }

func (ClientPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientPolicy)(nil)).Elem()
}

func (o ClientPolicyArrayOutput) ToClientPolicyArrayOutput() ClientPolicyArrayOutput {
	return o
}

func (o ClientPolicyArrayOutput) ToClientPolicyArrayOutputWithContext(ctx context.Context) ClientPolicyArrayOutput {
	return o
}

func (o ClientPolicyArrayOutput) Index(i pulumi.IntInput) ClientPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientPolicy {
		return vs[0].([]*ClientPolicy)[vs[1].(int)]
	}).(ClientPolicyOutput)
}

type ClientPolicyMapOutput struct{ *pulumi.OutputState }

func (ClientPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientPolicy)(nil)).Elem()
}

func (o ClientPolicyMapOutput) ToClientPolicyMapOutput() ClientPolicyMapOutput {
	return o
}

func (o ClientPolicyMapOutput) ToClientPolicyMapOutputWithContext(ctx context.Context) ClientPolicyMapOutput {
	return o
}

func (o ClientPolicyMapOutput) MapIndex(k pulumi.StringInput) ClientPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientPolicy {
		return vs[0].(map[string]*ClientPolicy)[vs[1].(string)]
	}).(ClientPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientPolicyInput)(nil)).Elem(), &ClientPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientPolicyArrayInput)(nil)).Elem(), ClientPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientPolicyMapInput)(nil)).Elem(), ClientPolicyMap{})
	pulumi.RegisterOutputType(ClientPolicyOutput{})
	pulumi.RegisterOutputType(ClientPolicyArrayOutput{})
	pulumi.RegisterOutputType(ClientPolicyMapOutput{})
}
