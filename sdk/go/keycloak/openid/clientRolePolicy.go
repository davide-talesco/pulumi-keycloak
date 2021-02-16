// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientRolePolicy struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringPtrOutput          `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput          `pulumi:"description"`
	Logic            pulumi.StringPtrOutput          `pulumi:"logic"`
	Name             pulumi.StringOutput             `pulumi:"name"`
	RealmId          pulumi.StringOutput             `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput             `pulumi:"resourceServerId"`
	Roles            ClientRolePolicyRoleArrayOutput `pulumi:"roles"`
	Type             pulumi.StringOutput             `pulumi:"type"`
}

// NewClientRolePolicy registers a new resource with the given unique name, arguments, and options.
func NewClientRolePolicy(ctx *pulumi.Context,
	name string, args *ClientRolePolicyArgs, opts ...pulumi.ResourceOption) (*ClientRolePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.ResourceServerId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServerId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ClientRolePolicy
	err := ctx.RegisterResource("keycloak:openid/clientRolePolicy:ClientRolePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientRolePolicy gets an existing ClientRolePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientRolePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientRolePolicyState, opts ...pulumi.ResourceOption) (*ClientRolePolicy, error) {
	var resource ClientRolePolicy
	err := ctx.ReadResource("keycloak:openid/clientRolePolicy:ClientRolePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientRolePolicy resources.
type clientRolePolicyState struct {
	DecisionStrategy *string                `pulumi:"decisionStrategy"`
	Description      *string                `pulumi:"description"`
	Logic            *string                `pulumi:"logic"`
	Name             *string                `pulumi:"name"`
	RealmId          *string                `pulumi:"realmId"`
	ResourceServerId *string                `pulumi:"resourceServerId"`
	Roles            []ClientRolePolicyRole `pulumi:"roles"`
	Type             *string                `pulumi:"type"`
}

type ClientRolePolicyState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Roles            ClientRolePolicyRoleArrayInput
	Type             pulumi.StringPtrInput
}

func (ClientRolePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientRolePolicyState)(nil)).Elem()
}

type clientRolePolicyArgs struct {
	DecisionStrategy *string                `pulumi:"decisionStrategy"`
	Description      *string                `pulumi:"description"`
	Logic            *string                `pulumi:"logic"`
	Name             *string                `pulumi:"name"`
	RealmId          string                 `pulumi:"realmId"`
	ResourceServerId string                 `pulumi:"resourceServerId"`
	Roles            []ClientRolePolicyRole `pulumi:"roles"`
	Type             string                 `pulumi:"type"`
}

// The set of arguments for constructing a ClientRolePolicy resource.
type ClientRolePolicyArgs struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Roles            ClientRolePolicyRoleArrayInput
	Type             pulumi.StringInput
}

func (ClientRolePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientRolePolicyArgs)(nil)).Elem()
}

type ClientRolePolicyInput interface {
	pulumi.Input

	ToClientRolePolicyOutput() ClientRolePolicyOutput
	ToClientRolePolicyOutputWithContext(ctx context.Context) ClientRolePolicyOutput
}

func (*ClientRolePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRolePolicy)(nil))
}

func (i *ClientRolePolicy) ToClientRolePolicyOutput() ClientRolePolicyOutput {
	return i.ToClientRolePolicyOutputWithContext(context.Background())
}

func (i *ClientRolePolicy) ToClientRolePolicyOutputWithContext(ctx context.Context) ClientRolePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRolePolicyOutput)
}

func (i *ClientRolePolicy) ToClientRolePolicyPtrOutput() ClientRolePolicyPtrOutput {
	return i.ToClientRolePolicyPtrOutputWithContext(context.Background())
}

func (i *ClientRolePolicy) ToClientRolePolicyPtrOutputWithContext(ctx context.Context) ClientRolePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRolePolicyPtrOutput)
}

type ClientRolePolicyPtrInput interface {
	pulumi.Input

	ToClientRolePolicyPtrOutput() ClientRolePolicyPtrOutput
	ToClientRolePolicyPtrOutputWithContext(ctx context.Context) ClientRolePolicyPtrOutput
}

type clientRolePolicyPtrType ClientRolePolicyArgs

func (*clientRolePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRolePolicy)(nil))
}

func (i *clientRolePolicyPtrType) ToClientRolePolicyPtrOutput() ClientRolePolicyPtrOutput {
	return i.ToClientRolePolicyPtrOutputWithContext(context.Background())
}

func (i *clientRolePolicyPtrType) ToClientRolePolicyPtrOutputWithContext(ctx context.Context) ClientRolePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRolePolicyPtrOutput)
}

// ClientRolePolicyArrayInput is an input type that accepts ClientRolePolicyArray and ClientRolePolicyArrayOutput values.
// You can construct a concrete instance of `ClientRolePolicyArrayInput` via:
//
//          ClientRolePolicyArray{ ClientRolePolicyArgs{...} }
type ClientRolePolicyArrayInput interface {
	pulumi.Input

	ToClientRolePolicyArrayOutput() ClientRolePolicyArrayOutput
	ToClientRolePolicyArrayOutputWithContext(context.Context) ClientRolePolicyArrayOutput
}

type ClientRolePolicyArray []ClientRolePolicyInput

func (ClientRolePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClientRolePolicy)(nil))
}

func (i ClientRolePolicyArray) ToClientRolePolicyArrayOutput() ClientRolePolicyArrayOutput {
	return i.ToClientRolePolicyArrayOutputWithContext(context.Background())
}

func (i ClientRolePolicyArray) ToClientRolePolicyArrayOutputWithContext(ctx context.Context) ClientRolePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRolePolicyArrayOutput)
}

// ClientRolePolicyMapInput is an input type that accepts ClientRolePolicyMap and ClientRolePolicyMapOutput values.
// You can construct a concrete instance of `ClientRolePolicyMapInput` via:
//
//          ClientRolePolicyMap{ "key": ClientRolePolicyArgs{...} }
type ClientRolePolicyMapInput interface {
	pulumi.Input

	ToClientRolePolicyMapOutput() ClientRolePolicyMapOutput
	ToClientRolePolicyMapOutputWithContext(context.Context) ClientRolePolicyMapOutput
}

type ClientRolePolicyMap map[string]ClientRolePolicyInput

func (ClientRolePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClientRolePolicy)(nil))
}

func (i ClientRolePolicyMap) ToClientRolePolicyMapOutput() ClientRolePolicyMapOutput {
	return i.ToClientRolePolicyMapOutputWithContext(context.Background())
}

func (i ClientRolePolicyMap) ToClientRolePolicyMapOutputWithContext(ctx context.Context) ClientRolePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRolePolicyMapOutput)
}

type ClientRolePolicyOutput struct {
	*pulumi.OutputState
}

func (ClientRolePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRolePolicy)(nil))
}

func (o ClientRolePolicyOutput) ToClientRolePolicyOutput() ClientRolePolicyOutput {
	return o
}

func (o ClientRolePolicyOutput) ToClientRolePolicyOutputWithContext(ctx context.Context) ClientRolePolicyOutput {
	return o
}

func (o ClientRolePolicyOutput) ToClientRolePolicyPtrOutput() ClientRolePolicyPtrOutput {
	return o.ToClientRolePolicyPtrOutputWithContext(context.Background())
}

func (o ClientRolePolicyOutput) ToClientRolePolicyPtrOutputWithContext(ctx context.Context) ClientRolePolicyPtrOutput {
	return o.ApplyT(func(v ClientRolePolicy) *ClientRolePolicy {
		return &v
	}).(ClientRolePolicyPtrOutput)
}

type ClientRolePolicyPtrOutput struct {
	*pulumi.OutputState
}

func (ClientRolePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRolePolicy)(nil))
}

func (o ClientRolePolicyPtrOutput) ToClientRolePolicyPtrOutput() ClientRolePolicyPtrOutput {
	return o
}

func (o ClientRolePolicyPtrOutput) ToClientRolePolicyPtrOutputWithContext(ctx context.Context) ClientRolePolicyPtrOutput {
	return o
}

type ClientRolePolicyArrayOutput struct{ *pulumi.OutputState }

func (ClientRolePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientRolePolicy)(nil))
}

func (o ClientRolePolicyArrayOutput) ToClientRolePolicyArrayOutput() ClientRolePolicyArrayOutput {
	return o
}

func (o ClientRolePolicyArrayOutput) ToClientRolePolicyArrayOutputWithContext(ctx context.Context) ClientRolePolicyArrayOutput {
	return o
}

func (o ClientRolePolicyArrayOutput) Index(i pulumi.IntInput) ClientRolePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientRolePolicy {
		return vs[0].([]ClientRolePolicy)[vs[1].(int)]
	}).(ClientRolePolicyOutput)
}

type ClientRolePolicyMapOutput struct{ *pulumi.OutputState }

func (ClientRolePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClientRolePolicy)(nil))
}

func (o ClientRolePolicyMapOutput) ToClientRolePolicyMapOutput() ClientRolePolicyMapOutput {
	return o
}

func (o ClientRolePolicyMapOutput) ToClientRolePolicyMapOutputWithContext(ctx context.Context) ClientRolePolicyMapOutput {
	return o
}

func (o ClientRolePolicyMapOutput) MapIndex(k pulumi.StringInput) ClientRolePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClientRolePolicy {
		return vs[0].(map[string]ClientRolePolicy)[vs[1].(string)]
	}).(ClientRolePolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientRolePolicyOutput{})
	pulumi.RegisterOutputType(ClientRolePolicyPtrOutput{})
	pulumi.RegisterOutputType(ClientRolePolicyArrayOutput{})
	pulumi.RegisterOutputType(ClientRolePolicyMapOutput{})
}
