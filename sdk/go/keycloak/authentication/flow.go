// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing an authentication flow within Keycloak.
//
// [Authentication flows](https://www.keycloak.org/docs/11.0/server_admin/index.html#_authentication-flows) describe a sequence
// of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
// is a container for these actions, which are otherwise known as executions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/authentication"
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
// 		flow, err := authentication.NewFlow(ctx, "flow", &authentication.FlowArgs{
// 			RealmId: realm.ID(),
// 			Alias:   pulumi.String("my-flow-alias"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authentication.NewExecution(ctx, "execution", &authentication.ExecutionArgs{
// 			RealmId:         realm.ID(),
// 			ParentFlowAlias: flow.Alias,
// 			Authenticator:   pulumi.String("identity-provider-redirector"),
// 			Requirement:     pulumi.String("REQUIRED"),
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
// Authentication flows can be imported using the format `{{realmId}}/{{authenticationFlowId}}`. The authentication flow ID is typically a GUID which is autogenerated when the flow is created via Keycloak. Unfortunately, it is not trivial to retrieve the authentication flow ID from the UI. The best way to do this is to visit the "Authentication" page in Keycloak, and use the network tab of your browser to view the response of the API call to `/auth/admin/realms/${realm}/authentication/flows`, which will be a list of authentication flows. Examplebash
//
// ```sh
//  $ pulumi import keycloak:authentication/flow:Flow flow my-realm/e9a5641e-778c-4daf-89c0-f4ef617987d1
// ```
type Flow struct {
	pulumi.CustomResourceState

	// The alias for this authentication flow.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// A description for the authentication flow.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId pulumi.StringPtrOutput `pulumi:"providerId"`
	// The realm that the authentication flow exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource Flow
	err := ctx.RegisterResource("keycloak:authentication/flow:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("keycloak:authentication/flow:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
	// The alias for this authentication flow.
	Alias *string `pulumi:"alias"`
	// A description for the authentication flow.
	Description *string `pulumi:"description"`
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId *string `pulumi:"providerId"`
	// The realm that the authentication flow exists in.
	RealmId *string `pulumi:"realmId"`
}

type FlowState struct {
	// The alias for this authentication flow.
	Alias pulumi.StringPtrInput
	// A description for the authentication flow.
	Description pulumi.StringPtrInput
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId pulumi.StringPtrInput
	// The realm that the authentication flow exists in.
	RealmId pulumi.StringPtrInput
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	// The alias for this authentication flow.
	Alias string `pulumi:"alias"`
	// A description for the authentication flow.
	Description *string `pulumi:"description"`
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId *string `pulumi:"providerId"`
	// The realm that the authentication flow exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	// The alias for this authentication flow.
	Alias pulumi.StringInput
	// A description for the authentication flow.
	Description pulumi.StringPtrInput
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId pulumi.StringPtrInput
	// The realm that the authentication flow exists in.
	RealmId pulumi.StringInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}

type FlowInput interface {
	pulumi.Input

	ToFlowOutput() FlowOutput
	ToFlowOutputWithContext(ctx context.Context) FlowOutput
}

func (*Flow) ElementType() reflect.Type {
	return reflect.TypeOf((*Flow)(nil))
}

func (i *Flow) ToFlowOutput() FlowOutput {
	return i.ToFlowOutputWithContext(context.Background())
}

func (i *Flow) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowOutput)
}

func (i *Flow) ToFlowPtrOutput() FlowPtrOutput {
	return i.ToFlowPtrOutputWithContext(context.Background())
}

func (i *Flow) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowPtrOutput)
}

type FlowPtrInput interface {
	pulumi.Input

	ToFlowPtrOutput() FlowPtrOutput
	ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput
}

type flowPtrType FlowArgs

func (*flowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil))
}

func (i *flowPtrType) ToFlowPtrOutput() FlowPtrOutput {
	return i.ToFlowPtrOutputWithContext(context.Background())
}

func (i *flowPtrType) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowPtrOutput)
}

// FlowArrayInput is an input type that accepts FlowArray and FlowArrayOutput values.
// You can construct a concrete instance of `FlowArrayInput` via:
//
//          FlowArray{ FlowArgs{...} }
type FlowArrayInput interface {
	pulumi.Input

	ToFlowArrayOutput() FlowArrayOutput
	ToFlowArrayOutputWithContext(context.Context) FlowArrayOutput
}

type FlowArray []FlowInput

func (FlowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flow)(nil)).Elem()
}

func (i FlowArray) ToFlowArrayOutput() FlowArrayOutput {
	return i.ToFlowArrayOutputWithContext(context.Background())
}

func (i FlowArray) ToFlowArrayOutputWithContext(ctx context.Context) FlowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowArrayOutput)
}

// FlowMapInput is an input type that accepts FlowMap and FlowMapOutput values.
// You can construct a concrete instance of `FlowMapInput` via:
//
//          FlowMap{ "key": FlowArgs{...} }
type FlowMapInput interface {
	pulumi.Input

	ToFlowMapOutput() FlowMapOutput
	ToFlowMapOutputWithContext(context.Context) FlowMapOutput
}

type FlowMap map[string]FlowInput

func (FlowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flow)(nil)).Elem()
}

func (i FlowMap) ToFlowMapOutput() FlowMapOutput {
	return i.ToFlowMapOutputWithContext(context.Background())
}

func (i FlowMap) ToFlowMapOutputWithContext(ctx context.Context) FlowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowMapOutput)
}

type FlowOutput struct{ *pulumi.OutputState }

func (FlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Flow)(nil))
}

func (o FlowOutput) ToFlowOutput() FlowOutput {
	return o
}

func (o FlowOutput) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return o
}

func (o FlowOutput) ToFlowPtrOutput() FlowPtrOutput {
	return o.ToFlowPtrOutputWithContext(context.Background())
}

func (o FlowOutput) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Flow) *Flow {
		return &v
	}).(FlowPtrOutput)
}

type FlowPtrOutput struct{ *pulumi.OutputState }

func (FlowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil))
}

func (o FlowPtrOutput) ToFlowPtrOutput() FlowPtrOutput {
	return o
}

func (o FlowPtrOutput) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return o
}

func (o FlowPtrOutput) Elem() FlowOutput {
	return o.ApplyT(func(v *Flow) Flow {
		if v != nil {
			return *v
		}
		var ret Flow
		return ret
	}).(FlowOutput)
}

type FlowArrayOutput struct{ *pulumi.OutputState }

func (FlowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Flow)(nil))
}

func (o FlowArrayOutput) ToFlowArrayOutput() FlowArrayOutput {
	return o
}

func (o FlowArrayOutput) ToFlowArrayOutputWithContext(ctx context.Context) FlowArrayOutput {
	return o
}

func (o FlowArrayOutput) Index(i pulumi.IntInput) FlowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Flow {
		return vs[0].([]Flow)[vs[1].(int)]
	}).(FlowOutput)
}

type FlowMapOutput struct{ *pulumi.OutputState }

func (FlowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Flow)(nil))
}

func (o FlowMapOutput) ToFlowMapOutput() FlowMapOutput {
	return o
}

func (o FlowMapOutput) ToFlowMapOutputWithContext(ctx context.Context) FlowMapOutput {
	return o
}

func (o FlowMapOutput) MapIndex(k pulumi.StringInput) FlowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Flow {
		return vs[0].(map[string]Flow)[vs[1].(string)]
	}).(FlowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowInput)(nil)).Elem(), &Flow{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowPtrInput)(nil)).Elem(), &Flow{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowArrayInput)(nil)).Elem(), FlowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowMapInput)(nil)).Elem(), FlowMap{})
	pulumi.RegisterOutputType(FlowOutput{})
	pulumi.RegisterOutputType(FlowPtrOutput{})
	pulumi.RegisterOutputType(FlowArrayOutput{})
	pulumi.RegisterOutputType(FlowMapOutput{})
}
