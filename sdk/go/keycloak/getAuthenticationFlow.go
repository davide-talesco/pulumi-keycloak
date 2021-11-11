// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch the ID of an authentication flow within Keycloak.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
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
// 		return nil
// 	})
// }
// ```
func GetAuthenticationFlow(ctx *pulumi.Context, args *GetAuthenticationFlowArgs, opts ...pulumi.InvokeOption) (*GetAuthenticationFlowResult, error) {
	var rv GetAuthenticationFlowResult
	err := ctx.Invoke("keycloak:index/getAuthenticationFlow:getAuthenticationFlow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthenticationFlow.
type GetAuthenticationFlowArgs struct {
	// The alias of the flow.
	Alias string `pulumi:"alias"`
	// The realm the authentication flow exists in.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getAuthenticationFlow.
type GetAuthenticationFlowResult struct {
	Alias string `pulumi:"alias"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	RealmId string `pulumi:"realmId"`
}

func GetAuthenticationFlowOutput(ctx *pulumi.Context, args GetAuthenticationFlowOutputArgs, opts ...pulumi.InvokeOption) GetAuthenticationFlowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAuthenticationFlowResult, error) {
			args := v.(GetAuthenticationFlowArgs)
			r, err := GetAuthenticationFlow(ctx, &args, opts...)
			return *r, err
		}).(GetAuthenticationFlowResultOutput)
}

// A collection of arguments for invoking getAuthenticationFlow.
type GetAuthenticationFlowOutputArgs struct {
	// The alias of the flow.
	Alias pulumi.StringInput `pulumi:"alias"`
	// The realm the authentication flow exists in.
	RealmId pulumi.StringInput `pulumi:"realmId"`
}

func (GetAuthenticationFlowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthenticationFlowArgs)(nil)).Elem()
}

// A collection of values returned by getAuthenticationFlow.
type GetAuthenticationFlowResultOutput struct{ *pulumi.OutputState }

func (GetAuthenticationFlowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthenticationFlowResult)(nil)).Elem()
}

func (o GetAuthenticationFlowResultOutput) ToGetAuthenticationFlowResultOutput() GetAuthenticationFlowResultOutput {
	return o
}

func (o GetAuthenticationFlowResultOutput) ToGetAuthenticationFlowResultOutputWithContext(ctx context.Context) GetAuthenticationFlowResultOutput {
	return o
}

func (o GetAuthenticationFlowResultOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthenticationFlowResult) string { return v.Alias }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAuthenticationFlowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthenticationFlowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAuthenticationFlowResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthenticationFlowResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAuthenticationFlowResultOutput{})
}
