# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_timeout: Optional[pulumi.Input[int]] = None,
                 initial_login: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 root_ca_certificate: Optional[pulumi.Input[str]] = None,
                 tls_insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the keycloak package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] client_timeout: Timeout (in seconds) of the Keycloak client
        :param pulumi.Input[bool] initial_login: Whether or not to login to Keycloak instance on provider initialization
        :param pulumi.Input[str] root_ca_certificate: Allows x509 calls using an unknown CA certificate (for development purposes)
        :param pulumi.Input[bool] tls_insecure_skip_verify: Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
               should be avoided.
        :param pulumi.Input[str] url: The base URL of the Keycloak instance, before `/auth`
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

            __props__['base_path'] = base_path
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            if client_timeout is None:
                client_timeout = (_utilities.get_env_int('KEYCLOAK_CLIENT_TIMEOUT') or 5)
            __props__['client_timeout'] = pulumi.Output.from_input(client_timeout).apply(pulumi.runtime.to_json) if client_timeout is not None else None
            __props__['initial_login'] = pulumi.Output.from_input(initial_login).apply(pulumi.runtime.to_json) if initial_login is not None else None
            __props__['password'] = password
            __props__['realm'] = realm
            __props__['root_ca_certificate'] = root_ca_certificate
            __props__['tls_insecure_skip_verify'] = pulumi.Output.from_input(tls_insecure_skip_verify).apply(pulumi.runtime.to_json) if tls_insecure_skip_verify is not None else None
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['username'] = username
        super(Provider, __self__).__init__(
            'keycloak',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

