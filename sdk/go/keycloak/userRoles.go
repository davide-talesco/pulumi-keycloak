// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// This resource can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak assigns to the user upon creation. This value can be found in the GUI when editing the user, and is typically a GUID. Examplebash
//
// ```sh
//  $ pulumi import keycloak:index/userRoles:UserRoles user_roles my-realm/b0ae6924-1bd5-4655-9e38-dae7c5e42924
// ```
type UserRoles struct {
	pulumi.CustomResourceState

	// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive pulumi.BoolPtrOutput `pulumi:"exhaustive"`
	// The realm this user exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// A list of role IDs to map to the user
	RoleIds pulumi.StringArrayOutput `pulumi:"roleIds"`
	// The ID of the user this resource should manage roles for.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserRoles registers a new resource with the given unique name, arguments, and options.
func NewUserRoles(ctx *pulumi.Context,
	name string, args *UserRolesArgs, opts ...pulumi.ResourceOption) (*UserRoles, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.RoleIds == nil {
		return nil, errors.New("invalid value for required argument 'RoleIds'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource UserRoles
	err := ctx.RegisterResource("keycloak:index/userRoles:UserRoles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRoles gets an existing UserRoles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRoles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRolesState, opts ...pulumi.ResourceOption) (*UserRoles, error) {
	var resource UserRoles
	err := ctx.ReadResource("keycloak:index/userRoles:UserRoles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRoles resources.
type userRolesState struct {
	// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive *bool `pulumi:"exhaustive"`
	// The realm this user exists in.
	RealmId *string `pulumi:"realmId"`
	// A list of role IDs to map to the user
	RoleIds []string `pulumi:"roleIds"`
	// The ID of the user this resource should manage roles for.
	UserId *string `pulumi:"userId"`
}

type UserRolesState struct {
	// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive pulumi.BoolPtrInput
	// The realm this user exists in.
	RealmId pulumi.StringPtrInput
	// A list of role IDs to map to the user
	RoleIds pulumi.StringArrayInput
	// The ID of the user this resource should manage roles for.
	UserId pulumi.StringPtrInput
}

func (UserRolesState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRolesState)(nil)).Elem()
}

type userRolesArgs struct {
	// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive *bool `pulumi:"exhaustive"`
	// The realm this user exists in.
	RealmId string `pulumi:"realmId"`
	// A list of role IDs to map to the user
	RoleIds []string `pulumi:"roleIds"`
	// The ID of the user this resource should manage roles for.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserRoles resource.
type UserRolesArgs struct {
	// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive pulumi.BoolPtrInput
	// The realm this user exists in.
	RealmId pulumi.StringInput
	// A list of role IDs to map to the user
	RoleIds pulumi.StringArrayInput
	// The ID of the user this resource should manage roles for.
	UserId pulumi.StringInput
}

func (UserRolesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRolesArgs)(nil)).Elem()
}

type UserRolesInput interface {
	pulumi.Input

	ToUserRolesOutput() UserRolesOutput
	ToUserRolesOutputWithContext(ctx context.Context) UserRolesOutput
}

func (*UserRoles) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRoles)(nil))
}

func (i *UserRoles) ToUserRolesOutput() UserRolesOutput {
	return i.ToUserRolesOutputWithContext(context.Background())
}

func (i *UserRoles) ToUserRolesOutputWithContext(ctx context.Context) UserRolesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRolesOutput)
}

func (i *UserRoles) ToUserRolesPtrOutput() UserRolesPtrOutput {
	return i.ToUserRolesPtrOutputWithContext(context.Background())
}

func (i *UserRoles) ToUserRolesPtrOutputWithContext(ctx context.Context) UserRolesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRolesPtrOutput)
}

type UserRolesPtrInput interface {
	pulumi.Input

	ToUserRolesPtrOutput() UserRolesPtrOutput
	ToUserRolesPtrOutputWithContext(ctx context.Context) UserRolesPtrOutput
}

type userRolesPtrType UserRolesArgs

func (*userRolesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRoles)(nil))
}

func (i *userRolesPtrType) ToUserRolesPtrOutput() UserRolesPtrOutput {
	return i.ToUserRolesPtrOutputWithContext(context.Background())
}

func (i *userRolesPtrType) ToUserRolesPtrOutputWithContext(ctx context.Context) UserRolesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRolesPtrOutput)
}

// UserRolesArrayInput is an input type that accepts UserRolesArray and UserRolesArrayOutput values.
// You can construct a concrete instance of `UserRolesArrayInput` via:
//
//          UserRolesArray{ UserRolesArgs{...} }
type UserRolesArrayInput interface {
	pulumi.Input

	ToUserRolesArrayOutput() UserRolesArrayOutput
	ToUserRolesArrayOutputWithContext(context.Context) UserRolesArrayOutput
}

type UserRolesArray []UserRolesInput

func (UserRolesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRoles)(nil)).Elem()
}

func (i UserRolesArray) ToUserRolesArrayOutput() UserRolesArrayOutput {
	return i.ToUserRolesArrayOutputWithContext(context.Background())
}

func (i UserRolesArray) ToUserRolesArrayOutputWithContext(ctx context.Context) UserRolesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRolesArrayOutput)
}

// UserRolesMapInput is an input type that accepts UserRolesMap and UserRolesMapOutput values.
// You can construct a concrete instance of `UserRolesMapInput` via:
//
//          UserRolesMap{ "key": UserRolesArgs{...} }
type UserRolesMapInput interface {
	pulumi.Input

	ToUserRolesMapOutput() UserRolesMapOutput
	ToUserRolesMapOutputWithContext(context.Context) UserRolesMapOutput
}

type UserRolesMap map[string]UserRolesInput

func (UserRolesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRoles)(nil)).Elem()
}

func (i UserRolesMap) ToUserRolesMapOutput() UserRolesMapOutput {
	return i.ToUserRolesMapOutputWithContext(context.Background())
}

func (i UserRolesMap) ToUserRolesMapOutputWithContext(ctx context.Context) UserRolesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRolesMapOutput)
}

type UserRolesOutput struct{ *pulumi.OutputState }

func (UserRolesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRoles)(nil))
}

func (o UserRolesOutput) ToUserRolesOutput() UserRolesOutput {
	return o
}

func (o UserRolesOutput) ToUserRolesOutputWithContext(ctx context.Context) UserRolesOutput {
	return o
}

func (o UserRolesOutput) ToUserRolesPtrOutput() UserRolesPtrOutput {
	return o.ToUserRolesPtrOutputWithContext(context.Background())
}

func (o UserRolesOutput) ToUserRolesPtrOutputWithContext(ctx context.Context) UserRolesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserRoles) *UserRoles {
		return &v
	}).(UserRolesPtrOutput)
}

type UserRolesPtrOutput struct{ *pulumi.OutputState }

func (UserRolesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRoles)(nil))
}

func (o UserRolesPtrOutput) ToUserRolesPtrOutput() UserRolesPtrOutput {
	return o
}

func (o UserRolesPtrOutput) ToUserRolesPtrOutputWithContext(ctx context.Context) UserRolesPtrOutput {
	return o
}

func (o UserRolesPtrOutput) Elem() UserRolesOutput {
	return o.ApplyT(func(v *UserRoles) UserRoles {
		if v != nil {
			return *v
		}
		var ret UserRoles
		return ret
	}).(UserRolesOutput)
}

type UserRolesArrayOutput struct{ *pulumi.OutputState }

func (UserRolesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRoles)(nil))
}

func (o UserRolesArrayOutput) ToUserRolesArrayOutput() UserRolesArrayOutput {
	return o
}

func (o UserRolesArrayOutput) ToUserRolesArrayOutputWithContext(ctx context.Context) UserRolesArrayOutput {
	return o
}

func (o UserRolesArrayOutput) Index(i pulumi.IntInput) UserRolesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserRoles {
		return vs[0].([]UserRoles)[vs[1].(int)]
	}).(UserRolesOutput)
}

type UserRolesMapOutput struct{ *pulumi.OutputState }

func (UserRolesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserRoles)(nil))
}

func (o UserRolesMapOutput) ToUserRolesMapOutput() UserRolesMapOutput {
	return o
}

func (o UserRolesMapOutput) ToUserRolesMapOutputWithContext(ctx context.Context) UserRolesMapOutput {
	return o
}

func (o UserRolesMapOutput) MapIndex(k pulumi.StringInput) UserRolesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserRoles {
		return vs[0].(map[string]UserRoles)[vs[1].(string)]
	}).(UserRolesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserRolesInput)(nil)).Elem(), &UserRoles{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRolesPtrInput)(nil)).Elem(), &UserRoles{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRolesArrayInput)(nil)).Elem(), UserRolesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRolesMapInput)(nil)).Elem(), UserRolesMap{})
	pulumi.RegisterOutputType(UserRolesOutput{})
	pulumi.RegisterOutputType(UserRolesPtrOutput{})
	pulumi.RegisterOutputType(UserRolesArrayOutput{})
	pulumi.RegisterOutputType(UserRolesMapOutput{})
}
