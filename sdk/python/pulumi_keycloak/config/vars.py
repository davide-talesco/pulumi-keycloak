# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('keycloak')

client_id = __config__.get('clientId') or utilities.get_env('KEYCLOAK_CLIENT_ID')

client_secret = __config__.get('clientSecret') or utilities.get_env('KEYCLOAK_CLIENT_SECRET')

client_timeout = __config__.get('clientTimeout') or (utilities.get_env_int('KEYCLOAK_CLIENT_TIMEOUT') or 5)
"""
Timeout (in seconds) of the Keycloak client
"""

initial_login = __config__.get('initialLogin')
"""
Whether or not to login to Keycloak instance on provider initialization
"""

password = __config__.get('password') or utilities.get_env('KEYCLOAK_PASSWORD')

realm = __config__.get('realm') or (utilities.get_env('KEYCLOAK_REALM') or 'master')

root_ca_certificate = __config__.get('rootCaCertificate')
"""
Allows x509 calls using an unknown CA certificate (for development purposes)
"""

tls_insecure_skip_verify = __config__.get('tlsInsecureSkipVerify')
"""
Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
should be avoided.
"""

url = __config__.get('url') or utilities.get_env('KEYCLOAK_URL')
"""
The base URL of the Keycloak instance, before `/auth`
"""

username = __config__.get('username') or utilities.get_env('KEYCLOAK_USER')

