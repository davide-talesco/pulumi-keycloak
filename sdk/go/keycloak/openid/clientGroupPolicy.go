// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientGroupPolicy struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringOutput               `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput            `pulumi:"description"`
	Groups           ClientGroupPolicyGroupArrayOutput `pulumi:"groups"`
	GroupsClaim      pulumi.StringPtrOutput            `pulumi:"groupsClaim"`
	Logic            pulumi.StringPtrOutput            `pulumi:"logic"`
	Name             pulumi.StringOutput               `pulumi:"name"`
	RealmId          pulumi.StringOutput               `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput               `pulumi:"resourceServerId"`
}

// NewClientGroupPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientGroupPolicy(ctx *pulumi.Context,
	name string, args *ClientGroupPolicyArgs, opts ...pulumi.ResourceOption) (*ClientGroupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DecisionStrategy == nil {
		return nil, errors.New("invalid value for required argument 'DecisionStrategy'")
	}
	if args.Groups == nil {
		return nil, errors.New("invalid value for required argument 'Groups'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.ResourceServerId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServerId'")
	}
	var resource ClientGroupPolicy
	err := ctx.RegisterResource("keycloak:openid/clientGroupPolicy:ClientGroupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientGroupPolicy gets an existing ClientGroupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientGroupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientGroupPolicyState, opts ...pulumi.ResourceOption) (*ClientGroupPolicy, error) {
	var resource ClientGroupPolicy
	err := ctx.ReadResource("keycloak:openid/clientGroupPolicy:ClientGroupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientGroupPolicy resources.
type clientGroupPolicyState struct {
	DecisionStrategy *string                  `pulumi:"decisionStrategy"`
	Description      *string                  `pulumi:"description"`
	Groups           []ClientGroupPolicyGroup `pulumi:"groups"`
	GroupsClaim      *string                  `pulumi:"groupsClaim"`
	Logic            *string                  `pulumi:"logic"`
	Name             *string                  `pulumi:"name"`
	RealmId          *string                  `pulumi:"realmId"`
	ResourceServerId *string                  `pulumi:"resourceServerId"`
}

type ClientGroupPolicyState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Groups           ClientGroupPolicyGroupArrayInput
	GroupsClaim      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
}

func (ClientGroupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGroupPolicyState)(nil)).Elem()
}

type clientGroupPolicyArgs struct {
	DecisionStrategy string                   `pulumi:"decisionStrategy"`
	Description      *string                  `pulumi:"description"`
	Groups           []ClientGroupPolicyGroup `pulumi:"groups"`
	GroupsClaim      *string                  `pulumi:"groupsClaim"`
	Logic            *string                  `pulumi:"logic"`
	Name             *string                  `pulumi:"name"`
	RealmId          string                   `pulumi:"realmId"`
	ResourceServerId string                   `pulumi:"resourceServerId"`
}

// The set of arguments for constructing a ClientGroupPolicy resource.
type ClientGroupPolicyArgs struct {
	DecisionStrategy pulumi.StringInput
	Description      pulumi.StringPtrInput
	Groups           ClientGroupPolicyGroupArrayInput
	GroupsClaim      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
}

func (ClientGroupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGroupPolicyArgs)(nil)).Elem()
}

type ClientGroupPolicyInput interface {
	pulumi.Input

	ToClientGroupPolicyOutput() ClientGroupPolicyOutput
	ToClientGroupPolicyOutputWithContext(ctx context.Context) ClientGroupPolicyOutput
}

func (*ClientGroupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGroupPolicy)(nil)).Elem()
}

func (i *ClientGroupPolicy) ToClientGroupPolicyOutput() ClientGroupPolicyOutput {
	return i.ToClientGroupPolicyOutputWithContext(context.Background())
}

func (i *ClientGroupPolicy) ToClientGroupPolicyOutputWithContext(ctx context.Context) ClientGroupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupPolicyOutput)
}

// ClientGroupPolicyArrayInput is an input type that accepts ClientGroupPolicyArray and ClientGroupPolicyArrayOutput values.
// You can construct a concrete instance of `ClientGroupPolicyArrayInput` via:
//
//          ClientGroupPolicyArray{ ClientGroupPolicyArgs{...} }
type ClientGroupPolicyArrayInput interface {
	pulumi.Input

	ToClientGroupPolicyArrayOutput() ClientGroupPolicyArrayOutput
	ToClientGroupPolicyArrayOutputWithContext(context.Context) ClientGroupPolicyArrayOutput
}

type ClientGroupPolicyArray []ClientGroupPolicyInput

func (ClientGroupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientGroupPolicy)(nil)).Elem()
}

func (i ClientGroupPolicyArray) ToClientGroupPolicyArrayOutput() ClientGroupPolicyArrayOutput {
	return i.ToClientGroupPolicyArrayOutputWithContext(context.Background())
}

func (i ClientGroupPolicyArray) ToClientGroupPolicyArrayOutputWithContext(ctx context.Context) ClientGroupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupPolicyArrayOutput)
}

// ClientGroupPolicyMapInput is an input type that accepts ClientGroupPolicyMap and ClientGroupPolicyMapOutput values.
// You can construct a concrete instance of `ClientGroupPolicyMapInput` via:
//
//          ClientGroupPolicyMap{ "key": ClientGroupPolicyArgs{...} }
type ClientGroupPolicyMapInput interface {
	pulumi.Input

	ToClientGroupPolicyMapOutput() ClientGroupPolicyMapOutput
	ToClientGroupPolicyMapOutputWithContext(context.Context) ClientGroupPolicyMapOutput
}

type ClientGroupPolicyMap map[string]ClientGroupPolicyInput

func (ClientGroupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientGroupPolicy)(nil)).Elem()
}

func (i ClientGroupPolicyMap) ToClientGroupPolicyMapOutput() ClientGroupPolicyMapOutput {
	return i.ToClientGroupPolicyMapOutputWithContext(context.Background())
}

func (i ClientGroupPolicyMap) ToClientGroupPolicyMapOutputWithContext(ctx context.Context) ClientGroupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupPolicyMapOutput)
}

type ClientGroupPolicyOutput struct{ *pulumi.OutputState }

func (ClientGroupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGroupPolicy)(nil)).Elem()
}

func (o ClientGroupPolicyOutput) ToClientGroupPolicyOutput() ClientGroupPolicyOutput {
	return o
}

func (o ClientGroupPolicyOutput) ToClientGroupPolicyOutputWithContext(ctx context.Context) ClientGroupPolicyOutput {
	return o
}

type ClientGroupPolicyArrayOutput struct{ *pulumi.OutputState }

func (ClientGroupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientGroupPolicy)(nil)).Elem()
}

func (o ClientGroupPolicyArrayOutput) ToClientGroupPolicyArrayOutput() ClientGroupPolicyArrayOutput {
	return o
}

func (o ClientGroupPolicyArrayOutput) ToClientGroupPolicyArrayOutputWithContext(ctx context.Context) ClientGroupPolicyArrayOutput {
	return o
}

func (o ClientGroupPolicyArrayOutput) Index(i pulumi.IntInput) ClientGroupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientGroupPolicy {
		return vs[0].([]*ClientGroupPolicy)[vs[1].(int)]
	}).(ClientGroupPolicyOutput)
}

type ClientGroupPolicyMapOutput struct{ *pulumi.OutputState }

func (ClientGroupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientGroupPolicy)(nil)).Elem()
}

func (o ClientGroupPolicyMapOutput) ToClientGroupPolicyMapOutput() ClientGroupPolicyMapOutput {
	return o
}

func (o ClientGroupPolicyMapOutput) ToClientGroupPolicyMapOutputWithContext(ctx context.Context) ClientGroupPolicyMapOutput {
	return o
}

func (o ClientGroupPolicyMapOutput) MapIndex(k pulumi.StringInput) ClientGroupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientGroupPolicy {
		return vs[0].(map[string]*ClientGroupPolicy)[vs[1].(string)]
	}).(ClientGroupPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientGroupPolicyInput)(nil)).Elem(), &ClientGroupPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientGroupPolicyArrayInput)(nil)).Elem(), ClientGroupPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientGroupPolicyMapInput)(nil)).Elem(), ClientGroupPolicyMap{})
	pulumi.RegisterOutputType(ClientGroupPolicyOutput{})
	pulumi.RegisterOutputType(ClientGroupPolicyArrayOutput{})
	pulumi.RegisterOutputType(ClientGroupPolicyMapOutput{})
}
