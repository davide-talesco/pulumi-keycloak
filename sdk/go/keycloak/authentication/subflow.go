// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Subflow struct {
	pulumi.CustomResourceState

	Alias pulumi.StringOutput `pulumi:"alias"`
	// Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
	Authenticator   pulumi.StringPtrOutput `pulumi:"authenticator"`
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	ParentFlowAlias pulumi.StringOutput    `pulumi:"parentFlowAlias"`
	ProviderId      pulumi.StringPtrOutput `pulumi:"providerId"`
	RealmId         pulumi.StringOutput    `pulumi:"realmId"`
	Requirement     pulumi.StringPtrOutput `pulumi:"requirement"`
}

// NewSubflow registers a new resource with the given unique name, arguments, and options.
func NewSubflow(ctx *pulumi.Context,
	name string, args *SubflowArgs, opts ...pulumi.ResourceOption) (*Subflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.ParentFlowAlias == nil {
		return nil, errors.New("invalid value for required argument 'ParentFlowAlias'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource Subflow
	err := ctx.RegisterResource("keycloak:authentication/subflow:Subflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubflow gets an existing Subflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubflowState, opts ...pulumi.ResourceOption) (*Subflow, error) {
	var resource Subflow
	err := ctx.ReadResource("keycloak:authentication/subflow:Subflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subflow resources.
type subflowState struct {
	Alias *string `pulumi:"alias"`
	// Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
	Authenticator   *string `pulumi:"authenticator"`
	Description     *string `pulumi:"description"`
	ParentFlowAlias *string `pulumi:"parentFlowAlias"`
	ProviderId      *string `pulumi:"providerId"`
	RealmId         *string `pulumi:"realmId"`
	Requirement     *string `pulumi:"requirement"`
}

type SubflowState struct {
	Alias pulumi.StringPtrInput
	// Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
	Authenticator   pulumi.StringPtrInput
	Description     pulumi.StringPtrInput
	ParentFlowAlias pulumi.StringPtrInput
	ProviderId      pulumi.StringPtrInput
	RealmId         pulumi.StringPtrInput
	Requirement     pulumi.StringPtrInput
}

func (SubflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*subflowState)(nil)).Elem()
}

type subflowArgs struct {
	Alias string `pulumi:"alias"`
	// Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
	Authenticator   *string `pulumi:"authenticator"`
	Description     *string `pulumi:"description"`
	ParentFlowAlias string  `pulumi:"parentFlowAlias"`
	ProviderId      *string `pulumi:"providerId"`
	RealmId         string  `pulumi:"realmId"`
	Requirement     *string `pulumi:"requirement"`
}

// The set of arguments for constructing a Subflow resource.
type SubflowArgs struct {
	Alias pulumi.StringInput
	// Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
	Authenticator   pulumi.StringPtrInput
	Description     pulumi.StringPtrInput
	ParentFlowAlias pulumi.StringInput
	ProviderId      pulumi.StringPtrInput
	RealmId         pulumi.StringInput
	Requirement     pulumi.StringPtrInput
}

func (SubflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subflowArgs)(nil)).Elem()
}

type SubflowInput interface {
	pulumi.Input

	ToSubflowOutput() SubflowOutput
	ToSubflowOutputWithContext(ctx context.Context) SubflowOutput
}

func (*Subflow) ElementType() reflect.Type {
	return reflect.TypeOf((*Subflow)(nil))
}

func (i *Subflow) ToSubflowOutput() SubflowOutput {
	return i.ToSubflowOutputWithContext(context.Background())
}

func (i *Subflow) ToSubflowOutputWithContext(ctx context.Context) SubflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubflowOutput)
}

func (i *Subflow) ToSubflowPtrOutput() SubflowPtrOutput {
	return i.ToSubflowPtrOutputWithContext(context.Background())
}

func (i *Subflow) ToSubflowPtrOutputWithContext(ctx context.Context) SubflowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubflowPtrOutput)
}

type SubflowPtrInput interface {
	pulumi.Input

	ToSubflowPtrOutput() SubflowPtrOutput
	ToSubflowPtrOutputWithContext(ctx context.Context) SubflowPtrOutput
}

type subflowPtrType SubflowArgs

func (*subflowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subflow)(nil))
}

func (i *subflowPtrType) ToSubflowPtrOutput() SubflowPtrOutput {
	return i.ToSubflowPtrOutputWithContext(context.Background())
}

func (i *subflowPtrType) ToSubflowPtrOutputWithContext(ctx context.Context) SubflowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubflowPtrOutput)
}

// SubflowArrayInput is an input type that accepts SubflowArray and SubflowArrayOutput values.
// You can construct a concrete instance of `SubflowArrayInput` via:
//
//          SubflowArray{ SubflowArgs{...} }
type SubflowArrayInput interface {
	pulumi.Input

	ToSubflowArrayOutput() SubflowArrayOutput
	ToSubflowArrayOutputWithContext(context.Context) SubflowArrayOutput
}

type SubflowArray []SubflowInput

func (SubflowArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Subflow)(nil))
}

func (i SubflowArray) ToSubflowArrayOutput() SubflowArrayOutput {
	return i.ToSubflowArrayOutputWithContext(context.Background())
}

func (i SubflowArray) ToSubflowArrayOutputWithContext(ctx context.Context) SubflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubflowArrayOutput)
}

// SubflowMapInput is an input type that accepts SubflowMap and SubflowMapOutput values.
// You can construct a concrete instance of `SubflowMapInput` via:
//
//          SubflowMap{ "key": SubflowArgs{...} }
type SubflowMapInput interface {
	pulumi.Input

	ToSubflowMapOutput() SubflowMapOutput
	ToSubflowMapOutputWithContext(context.Context) SubflowMapOutput
}

type SubflowMap map[string]SubflowInput

func (SubflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Subflow)(nil))
}

func (i SubflowMap) ToSubflowMapOutput() SubflowMapOutput {
	return i.ToSubflowMapOutputWithContext(context.Background())
}

func (i SubflowMap) ToSubflowMapOutputWithContext(ctx context.Context) SubflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubflowMapOutput)
}

type SubflowOutput struct {
	*pulumi.OutputState
}

func (SubflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subflow)(nil))
}

func (o SubflowOutput) ToSubflowOutput() SubflowOutput {
	return o
}

func (o SubflowOutput) ToSubflowOutputWithContext(ctx context.Context) SubflowOutput {
	return o
}

func (o SubflowOutput) ToSubflowPtrOutput() SubflowPtrOutput {
	return o.ToSubflowPtrOutputWithContext(context.Background())
}

func (o SubflowOutput) ToSubflowPtrOutputWithContext(ctx context.Context) SubflowPtrOutput {
	return o.ApplyT(func(v Subflow) *Subflow {
		return &v
	}).(SubflowPtrOutput)
}

type SubflowPtrOutput struct {
	*pulumi.OutputState
}

func (SubflowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subflow)(nil))
}

func (o SubflowPtrOutput) ToSubflowPtrOutput() SubflowPtrOutput {
	return o
}

func (o SubflowPtrOutput) ToSubflowPtrOutputWithContext(ctx context.Context) SubflowPtrOutput {
	return o
}

type SubflowArrayOutput struct{ *pulumi.OutputState }

func (SubflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subflow)(nil))
}

func (o SubflowArrayOutput) ToSubflowArrayOutput() SubflowArrayOutput {
	return o
}

func (o SubflowArrayOutput) ToSubflowArrayOutputWithContext(ctx context.Context) SubflowArrayOutput {
	return o
}

func (o SubflowArrayOutput) Index(i pulumi.IntInput) SubflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subflow {
		return vs[0].([]Subflow)[vs[1].(int)]
	}).(SubflowOutput)
}

type SubflowMapOutput struct{ *pulumi.OutputState }

func (SubflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Subflow)(nil))
}

func (o SubflowMapOutput) ToSubflowMapOutput() SubflowMapOutput {
	return o
}

func (o SubflowMapOutput) ToSubflowMapOutputWithContext(ctx context.Context) SubflowMapOutput {
	return o
}

func (o SubflowMapOutput) MapIndex(k pulumi.StringInput) SubflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Subflow {
		return vs[0].(map[string]Subflow)[vs[1].(string)]
	}).(SubflowOutput)
}

func init() {
	pulumi.RegisterOutputType(SubflowOutput{})
	pulumi.RegisterOutputType(SubflowPtrOutput{})
	pulumi.RegisterOutputType(SubflowArrayOutput{})
	pulumi.RegisterOutputType(SubflowMapOutput{})
}
