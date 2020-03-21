// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # ldap.FullNameMapper
//
// Allows for creating and managing full name mappers for Keycloak users federated
// via LDAP.
//
// The LDAP full name mapper can map a user's full name from an LDAP attribute
// to the first and last name attributes of a Keycloak user.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm that this LDAP mapper will exist in.
// - `ldapUserFederationId` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
// - `name` - (Required) Display name of this mapper when displayed in the console.
// - `ldapFullNameAttribute` - (Required) The name of the LDAP attribute containing the user's full name.
// - `readOnly` - (Optional) When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
// - `writeOnly` - (Optional) When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_ldap_full_name_mapper.html.markdown.
type FullNameMapper struct {
	pulumi.CustomResourceState

	LdapFullNameAttribute pulumi.StringOutput `pulumi:"ldapFullNameAttribute"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	ReadOnly pulumi.BoolPtrOutput `pulumi:"readOnly"`
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	WriteOnly pulumi.BoolPtrOutput `pulumi:"writeOnly"`
}

// NewFullNameMapper registers a new resource with the given unique name, arguments, and options.
func NewFullNameMapper(ctx *pulumi.Context,
	name string, args *FullNameMapperArgs, opts ...pulumi.ResourceOption) (*FullNameMapper, error) {
	if args == nil || args.LdapFullNameAttribute == nil {
		return nil, errors.New("missing required argument 'LdapFullNameAttribute'")
	}
	if args == nil || args.LdapUserFederationId == nil {
		return nil, errors.New("missing required argument 'LdapUserFederationId'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &FullNameMapperArgs{}
	}
	var resource FullNameMapper
	err := ctx.RegisterResource("keycloak:ldap/fullNameMapper:FullNameMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFullNameMapper gets an existing FullNameMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFullNameMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FullNameMapperState, opts ...pulumi.ResourceOption) (*FullNameMapper, error) {
	var resource FullNameMapper
	err := ctx.ReadResource("keycloak:ldap/fullNameMapper:FullNameMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FullNameMapper resources.
type fullNameMapperState struct {
	LdapFullNameAttribute *string `pulumi:"ldapFullNameAttribute"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	ReadOnly *bool `pulumi:"readOnly"`
	// The realm in which the ldap user federation provider exists.
	RealmId *string `pulumi:"realmId"`
	WriteOnly *bool `pulumi:"writeOnly"`
}

type FullNameMapperState struct {
	LdapFullNameAttribute pulumi.StringPtrInput
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	ReadOnly pulumi.BoolPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringPtrInput
	WriteOnly pulumi.BoolPtrInput
}

func (FullNameMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameMapperState)(nil)).Elem()
}

type fullNameMapperArgs struct {
	LdapFullNameAttribute string `pulumi:"ldapFullNameAttribute"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	ReadOnly *bool `pulumi:"readOnly"`
	// The realm in which the ldap user federation provider exists.
	RealmId string `pulumi:"realmId"`
	WriteOnly *bool `pulumi:"writeOnly"`
}

// The set of arguments for constructing a FullNameMapper resource.
type FullNameMapperArgs struct {
	LdapFullNameAttribute pulumi.StringInput
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	ReadOnly pulumi.BoolPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringInput
	WriteOnly pulumi.BoolPtrInput
}

func (FullNameMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameMapperArgs)(nil)).Elem()
}

