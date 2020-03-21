# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ClientJsPolicy(pulumi.CustomResource):
    code: pulumi.Output[str]
    decision_strategy: pulumi.Output[str]
    description: pulumi.Output[str]
    logic: pulumi.Output[str]
    name: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    resource_server_id: pulumi.Output[str]
    type: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, code=None, decision_strategy=None, description=None, logic=None, name=None, realm_id=None, resource_server_id=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a ClientJsPolicy resource with the given unique name, props, and options.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if code is None:
                raise TypeError("Missing required property 'code'")
            __props__['code'] = code
            if decision_strategy is None:
                raise TypeError("Missing required property 'decision_strategy'")
            __props__['decision_strategy'] = decision_strategy
            __props__['description'] = description
            __props__['logic'] = logic
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if resource_server_id is None:
                raise TypeError("Missing required property 'resource_server_id'")
            __props__['resource_server_id'] = resource_server_id
            __props__['type'] = type
        super(ClientJsPolicy, __self__).__init__(
            'keycloak:openid/clientJsPolicy:ClientJsPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, code=None, decision_strategy=None, description=None, logic=None, name=None, realm_id=None, resource_server_id=None, type=None):
        """
        Get an existing ClientJsPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["code"] = code
        __props__["decision_strategy"] = decision_strategy
        __props__["description"] = description
        __props__["logic"] = logic
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        __props__["resource_server_id"] = resource_server_id
        __props__["type"] = type
        return ClientJsPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

