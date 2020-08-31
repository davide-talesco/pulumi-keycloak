# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Execution']


class Execution(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticator: Optional[pulumi.Input[str]] = None,
                 parent_flow_alias: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Execution resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            if authenticator is None:
                raise TypeError("Missing required property 'authenticator'")
            __props__['authenticator'] = authenticator
            if parent_flow_alias is None:
                raise TypeError("Missing required property 'parent_flow_alias'")
            __props__['parent_flow_alias'] = parent_flow_alias
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            __props__['requirement'] = requirement
        super(Execution, __self__).__init__(
            'keycloak:authentication/execution:Execution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authenticator: Optional[pulumi.Input[str]] = None,
            parent_flow_alias: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            requirement: Optional[pulumi.Input[str]] = None) -> 'Execution':
        """
        Get an existing Execution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authenticator"] = authenticator
        __props__["parent_flow_alias"] = parent_flow_alias
        __props__["realm_id"] = realm_id
        __props__["requirement"] = requirement
        return Execution(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authenticator(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authenticator")

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> pulumi.Output[str]:
        return pulumi.get(self, "parent_flow_alias")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def requirement(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "requirement")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

