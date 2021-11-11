// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch policy and permission information for an OpenID client that has authorization enabled.
//
// ## Example Usage
//
// In this example, we'll create a new OpenID client with authorization enabled. This will cause Keycloak to create a default
// permission for this client called "Default Permission". We'll use the `openid.getClientAuthorizationPolicy` data
// source to fetch information about this permission, so we can use it to create a new resource-based authorization permission.
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
// 		clientWithAuthz, err := openid.NewClient(ctx, "clientWithAuthz", &openid.ClientArgs{
// 			ClientId:               pulumi.String("client-with-authz"),
// 			RealmId:                realm.ID(),
// 			AccessType:             pulumi.String("CONFIDENTIAL"),
// 			ServiceAccountsEnabled: pulumi.Bool(true),
// 			Authorization: &openid.ClientAuthorizationArgs{
// 				PolicyEnforcementMode: pulumi.String("ENFORCING"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		resource, err := openid.NewClientAuthorizationResource(ctx, "resource", &openid.ClientAuthorizationResourceArgs{
// 			ResourceServerId: clientWithAuthz.ResourceServerId,
// 			RealmId:          realm.ID(),
// 			Uris: pulumi.StringArray{
// 				pulumi.String("/endpoint/*"),
// 			},
// 			Attributes: pulumi.AnyMap{
// 				"foo": pulumi.Any("bar"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewClientAuthorizationPermission(ctx, "permission", &openid.ClientAuthorizationPermissionArgs{
// 			ResourceServerId: clientWithAuthz.ResourceServerId,
// 			RealmId:          realm.ID(),
// 			Policies: pulumi.StringArray{
// 				defaultPermission.ApplyT(func(defaultPermission openid.GetClientAuthorizationPolicyResult) (string, error) {
// 					return defaultPermission.Id, nil
// 				}).(pulumi.StringOutput),
// 			},
// 			Resources: pulumi.StringArray{
// 				resource.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetClientAuthorizationPolicy(ctx *pulumi.Context, args *GetClientAuthorizationPolicyArgs, opts ...pulumi.InvokeOption) (*GetClientAuthorizationPolicyResult, error) {
	var rv GetClientAuthorizationPolicyResult
	err := ctx.Invoke("keycloak:openid/getClientAuthorizationPolicy:getClientAuthorizationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClientAuthorizationPolicy.
type GetClientAuthorizationPolicyArgs struct {
	// The name of the authorization policy.
	Name string `pulumi:"name"`
	// The realm this authorization policy exists within.
	RealmId string `pulumi:"realmId"`
	// The ID of the resource server this authorization policy is attached to.
	ResourceServerId string `pulumi:"resourceServerId"`
}

// A collection of values returned by getClientAuthorizationPolicy.
type GetClientAuthorizationPolicyResult struct {
	// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
	DecisionStrategy string `pulumi:"decisionStrategy"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
	Logic string `pulumi:"logic"`
	Name  string `pulumi:"name"`
	// (Computed) The ID of the owning resource. Applies to resources.
	Owner string `pulumi:"owner"`
	// (Computed) The IDs of the policies that must be applied to scopes/resources for this policy/permission. Applies to policies and permissions.
	Policies         []string `pulumi:"policies"`
	RealmId          string   `pulumi:"realmId"`
	ResourceServerId string   `pulumi:"resourceServerId"`
	// (Computed) The IDs of the resources that this permission applies to. Applies to resource-based permissions.
	Resources []string `pulumi:"resources"`
	// (Computed) The IDs of the scopes that this permission applies to. Applies to scope-based permissions.
	Scopes []string `pulumi:"scopes"`
	// (Computed) The type of this policy / permission. For permissions, this could be `resource` or `scope`. For policies, this could be any type of authorization policy, such as `js`.
	Type string `pulumi:"type"`
}

func GetClientAuthorizationPolicyOutput(ctx *pulumi.Context, args GetClientAuthorizationPolicyOutputArgs, opts ...pulumi.InvokeOption) GetClientAuthorizationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClientAuthorizationPolicyResult, error) {
			args := v.(GetClientAuthorizationPolicyArgs)
			r, err := GetClientAuthorizationPolicy(ctx, &args, opts...)
			return *r, err
		}).(GetClientAuthorizationPolicyResultOutput)
}

// A collection of arguments for invoking getClientAuthorizationPolicy.
type GetClientAuthorizationPolicyOutputArgs struct {
	// The name of the authorization policy.
	Name pulumi.StringInput `pulumi:"name"`
	// The realm this authorization policy exists within.
	RealmId pulumi.StringInput `pulumi:"realmId"`
	// The ID of the resource server this authorization policy is attached to.
	ResourceServerId pulumi.StringInput `pulumi:"resourceServerId"`
}

func (GetClientAuthorizationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthorizationPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getClientAuthorizationPolicy.
type GetClientAuthorizationPolicyResultOutput struct{ *pulumi.OutputState }

func (GetClientAuthorizationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthorizationPolicyResult)(nil)).Elem()
}

func (o GetClientAuthorizationPolicyResultOutput) ToGetClientAuthorizationPolicyResultOutput() GetClientAuthorizationPolicyResultOutput {
	return o
}

func (o GetClientAuthorizationPolicyResultOutput) ToGetClientAuthorizationPolicyResultOutputWithContext(ctx context.Context) GetClientAuthorizationPolicyResultOutput {
	return o
}

// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
func (o GetClientAuthorizationPolicyResultOutput) DecisionStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.DecisionStrategy }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClientAuthorizationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
func (o GetClientAuthorizationPolicyResultOutput) Logic() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.Logic }).(pulumi.StringOutput)
}

func (o GetClientAuthorizationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// (Computed) The ID of the owning resource. Applies to resources.
func (o GetClientAuthorizationPolicyResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.Owner }).(pulumi.StringOutput)
}

// (Computed) The IDs of the policies that must be applied to scopes/resources for this policy/permission. Applies to policies and permissions.
func (o GetClientAuthorizationPolicyResultOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

func (o GetClientAuthorizationPolicyResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func (o GetClientAuthorizationPolicyResultOutput) ResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.ResourceServerId }).(pulumi.StringOutput)
}

// (Computed) The IDs of the resources that this permission applies to. Applies to resource-based permissions.
func (o GetClientAuthorizationPolicyResultOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

// (Computed) The IDs of the scopes that this permission applies to. Applies to scope-based permissions.
func (o GetClientAuthorizationPolicyResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// (Computed) The type of this policy / permission. For permissions, this could be `resource` or `scope`. For policies, this could be any type of authorization policy, such as `js`.
func (o GetClientAuthorizationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorizationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientAuthorizationPolicyResultOutput{})
}
