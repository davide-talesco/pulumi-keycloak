# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['RoleMapper']


class RoleMapper(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 ldap_roles_dn: Optional[pulumi.Input[str]] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 memberof_ldap_attribute: Optional[pulumi.Input[str]] = None,
                 membership_attribute_type: Optional[pulumi.Input[str]] = None,
                 membership_ldap_attribute: Optional[pulumi.Input[str]] = None,
                 membership_user_ldap_attribute: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role_name_ldap_attribute: Optional[pulumi.Input[str]] = None,
                 role_object_classes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 roles_ldap_filter: Optional[pulumi.Input[str]] = None,
                 use_realm_roles_mapping: Optional[pulumi.Input[bool]] = None,
                 user_roles_retrieve_strategy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a RoleMapper resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['client_id'] = client_id
            if ldap_roles_dn is None:
                raise TypeError("Missing required property 'ldap_roles_dn'")
            __props__['ldap_roles_dn'] = ldap_roles_dn
            if ldap_user_federation_id is None:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__['ldap_user_federation_id'] = ldap_user_federation_id
            __props__['memberof_ldap_attribute'] = memberof_ldap_attribute
            __props__['membership_attribute_type'] = membership_attribute_type
            if membership_ldap_attribute is None:
                raise TypeError("Missing required property 'membership_ldap_attribute'")
            __props__['membership_ldap_attribute'] = membership_ldap_attribute
            if membership_user_ldap_attribute is None:
                raise TypeError("Missing required property 'membership_user_ldap_attribute'")
            __props__['membership_user_ldap_attribute'] = membership_user_ldap_attribute
            __props__['mode'] = mode
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if role_name_ldap_attribute is None:
                raise TypeError("Missing required property 'role_name_ldap_attribute'")
            __props__['role_name_ldap_attribute'] = role_name_ldap_attribute
            if role_object_classes is None:
                raise TypeError("Missing required property 'role_object_classes'")
            __props__['role_object_classes'] = role_object_classes
            __props__['roles_ldap_filter'] = roles_ldap_filter
            __props__['use_realm_roles_mapping'] = use_realm_roles_mapping
            __props__['user_roles_retrieve_strategy'] = user_roles_retrieve_strategy
        super(RoleMapper, __self__).__init__(
            'keycloak:ldap/roleMapper:RoleMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            ldap_roles_dn: Optional[pulumi.Input[str]] = None,
            ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
            memberof_ldap_attribute: Optional[pulumi.Input[str]] = None,
            membership_attribute_type: Optional[pulumi.Input[str]] = None,
            membership_ldap_attribute: Optional[pulumi.Input[str]] = None,
            membership_user_ldap_attribute: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            role_name_ldap_attribute: Optional[pulumi.Input[str]] = None,
            role_object_classes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            roles_ldap_filter: Optional[pulumi.Input[str]] = None,
            use_realm_roles_mapping: Optional[pulumi.Input[bool]] = None,
            user_roles_retrieve_strategy: Optional[pulumi.Input[str]] = None) -> 'RoleMapper':
        """
        Get an existing RoleMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["ldap_roles_dn"] = ldap_roles_dn
        __props__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__["memberof_ldap_attribute"] = memberof_ldap_attribute
        __props__["membership_attribute_type"] = membership_attribute_type
        __props__["membership_ldap_attribute"] = membership_ldap_attribute
        __props__["membership_user_ldap_attribute"] = membership_user_ldap_attribute
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        __props__["role_name_ldap_attribute"] = role_name_ldap_attribute
        __props__["role_object_classes"] = role_object_classes
        __props__["roles_ldap_filter"] = roles_ldap_filter
        __props__["use_realm_roles_mapping"] = use_realm_roles_mapping
        __props__["user_roles_retrieve_strategy"] = user_roles_retrieve_strategy
        return RoleMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="ldapRolesDn")
    def ldap_roles_dn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ldap_roles_dn")

    @property
    @pulumi.getter(name="ldapUserFederationId")
    def ldap_user_federation_id(self) -> pulumi.Output[str]:
        """
        The ldap user federation provider to attach this mapper to.
        """
        return pulumi.get(self, "ldap_user_federation_id")

    @property
    @pulumi.getter(name="memberofLdapAttribute")
    def memberof_ldap_attribute(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "memberof_ldap_attribute")

    @property
    @pulumi.getter(name="membershipAttributeType")
    def membership_attribute_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "membership_attribute_type")

    @property
    @pulumi.getter(name="membershipLdapAttribute")
    def membership_ldap_attribute(self) -> pulumi.Output[str]:
        return pulumi.get(self, "membership_ldap_attribute")

    @property
    @pulumi.getter(name="membershipUserLdapAttribute")
    def membership_user_ldap_attribute(self) -> pulumi.Output[str]:
        return pulumi.get(self, "membership_user_ldap_attribute")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Display name of the mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm in which the ldap user federation provider exists.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="roleNameLdapAttribute")
    def role_name_ldap_attribute(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role_name_ldap_attribute")

    @property
    @pulumi.getter(name="roleObjectClasses")
    def role_object_classes(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "role_object_classes")

    @property
    @pulumi.getter(name="rolesLdapFilter")
    def roles_ldap_filter(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "roles_ldap_filter")

    @property
    @pulumi.getter(name="useRealmRolesMapping")
    def use_realm_roles_mapping(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "use_realm_roles_mapping")

    @property
    @pulumi.getter(name="userRolesRetrieveStrategy")
    def user_roles_retrieve_strategy(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_roles_retrieve_strategy")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

