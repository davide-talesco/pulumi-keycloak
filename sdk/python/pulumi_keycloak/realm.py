# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Realm(pulumi.CustomResource):
    access_code_lifespan: pulumi.Output[str]
    access_code_lifespan_login: pulumi.Output[str]
    access_code_lifespan_user_action: pulumi.Output[str]
    access_token_lifespan: pulumi.Output[str]
    access_token_lifespan_for_implicit_flow: pulumi.Output[str]
    account_theme: pulumi.Output[str]
    action_token_generated_by_admin_lifespan: pulumi.Output[str]
    action_token_generated_by_user_lifespan: pulumi.Output[str]
    admin_theme: pulumi.Output[str]
    attributes: pulumi.Output[dict]
    browser_flow: pulumi.Output[str]
    """
    Which flow should be used for BrowserFlow
    """
    client_authentication_flow: pulumi.Output[str]
    """
    Which flow should be used for ClientAuthenticationFlow
    """
    default_signature_algorithm: pulumi.Output[str]
    direct_grant_flow: pulumi.Output[str]
    """
    Which flow should be used for DirectGrantFlow
    """
    display_name: pulumi.Output[str]
    display_name_html: pulumi.Output[str]
    docker_authentication_flow: pulumi.Output[str]
    """
    Which flow should be used for DockerAuthenticationFlow
    """
    duplicate_emails_allowed: pulumi.Output[bool]
    edit_username_allowed: pulumi.Output[bool]
    email_theme: pulumi.Output[str]
    enabled: pulumi.Output[bool]
    internal_id: pulumi.Output[str]
    internationalization: pulumi.Output[dict]
    login_theme: pulumi.Output[str]
    login_with_email_allowed: pulumi.Output[bool]
    offline_session_idle_timeout: pulumi.Output[str]
    offline_session_max_lifespan: pulumi.Output[str]
    password_policy: pulumi.Output[str]
    """
    String that represents the passwordPolicies that are in place. Each policy is separated with " and ". Supported policies
    can be found in the server-info providers page. example: "upperCase(1) and length(8) and forceExpiredPasswordChange(365)
    and notUsername(undefined)"
    """
    realm: pulumi.Output[str]
    refresh_token_max_reuse: pulumi.Output[float]
    registration_allowed: pulumi.Output[bool]
    registration_email_as_username: pulumi.Output[bool]
    registration_flow: pulumi.Output[str]
    """
    Which flow should be used for RegistrationFlow
    """
    remember_me: pulumi.Output[bool]
    reset_credentials_flow: pulumi.Output[str]
    """
    Which flow should be used for ResetCredentialsFlow
    """
    reset_password_allowed: pulumi.Output[bool]
    revoke_refresh_token: pulumi.Output[bool]
    security_defenses: pulumi.Output[dict]
    smtp_server: pulumi.Output[dict]
    ssl_required: pulumi.Output[str]
    """
    SSL Required: Values can be 'none', 'external' or 'all'.
    """
    sso_session_idle_timeout: pulumi.Output[str]
    sso_session_max_lifespan: pulumi.Output[str]
    user_managed_access: pulumi.Output[bool]
    verify_email: pulumi.Output[bool]
    def __init__(__self__, resource_name, opts=None, access_code_lifespan=None, access_code_lifespan_login=None, access_code_lifespan_user_action=None, access_token_lifespan=None, access_token_lifespan_for_implicit_flow=None, account_theme=None, action_token_generated_by_admin_lifespan=None, action_token_generated_by_user_lifespan=None, admin_theme=None, attributes=None, browser_flow=None, client_authentication_flow=None, default_signature_algorithm=None, direct_grant_flow=None, display_name=None, display_name_html=None, docker_authentication_flow=None, duplicate_emails_allowed=None, edit_username_allowed=None, email_theme=None, enabled=None, internationalization=None, login_theme=None, login_with_email_allowed=None, offline_session_idle_timeout=None, offline_session_max_lifespan=None, password_policy=None, realm=None, refresh_token_max_reuse=None, registration_allowed=None, registration_email_as_username=None, registration_flow=None, remember_me=None, reset_credentials_flow=None, reset_password_allowed=None, revoke_refresh_token=None, security_defenses=None, smtp_server=None, ssl_required=None, sso_session_idle_timeout=None, sso_session_max_lifespan=None, user_managed_access=None, verify_email=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Realm resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] browser_flow: Which flow should be used for BrowserFlow
        :param pulumi.Input[str] client_authentication_flow: Which flow should be used for ClientAuthenticationFlow
        :param pulumi.Input[str] direct_grant_flow: Which flow should be used for DirectGrantFlow
        :param pulumi.Input[str] docker_authentication_flow: Which flow should be used for DockerAuthenticationFlow
        :param pulumi.Input[str] password_policy: String that represents the passwordPolicies that are in place. Each policy is separated with " and ". Supported policies
               can be found in the server-info providers page. example: "upperCase(1) and length(8) and forceExpiredPasswordChange(365)
               and notUsername(undefined)"
        :param pulumi.Input[str] registration_flow: Which flow should be used for RegistrationFlow
        :param pulumi.Input[str] reset_credentials_flow: Which flow should be used for ResetCredentialsFlow
        :param pulumi.Input[str] ssl_required: SSL Required: Values can be 'none', 'external' or 'all'.

        The **internationalization** object supports the following:

          * `defaultLocale` (`pulumi.Input[str]`)
          * `supportedLocales` (`pulumi.Input[list]`)

        The **security_defenses** object supports the following:

          * `bruteForceDetection` (`pulumi.Input[dict]`)
            * `failureResetTimeSeconds` (`pulumi.Input[float]`)
            * `maxFailureWaitSeconds` (`pulumi.Input[float]`)
            * `maxLoginFailures` (`pulumi.Input[float]`)
            * `minimumQuickLoginWaitSeconds` (`pulumi.Input[float]`)
            * `permanentLockout` (`pulumi.Input[bool]`)
            * `quickLoginCheckMilliSeconds` (`pulumi.Input[float]`)
            * `waitIncrementSeconds` (`pulumi.Input[float]`)

          * `headers` (`pulumi.Input[dict]`)
            * `contentSecurityPolicy` (`pulumi.Input[str]`)
            * `contentSecurityPolicyReportOnly` (`pulumi.Input[str]`)
            * `strictTransportSecurity` (`pulumi.Input[str]`)
            * `xContentTypeOptions` (`pulumi.Input[str]`)
            * `xFrameOptions` (`pulumi.Input[str]`)
            * `xRobotsTag` (`pulumi.Input[str]`)
            * `xXssProtection` (`pulumi.Input[str]`)

        The **smtp_server** object supports the following:

          * `auth` (`pulumi.Input[dict]`)
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)

          * `envelopeFrom` (`pulumi.Input[str]`)
          * `from` (`pulumi.Input[str]`)
          * `fromDisplayName` (`pulumi.Input[str]`)
          * `host` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[str]`)
          * `replyTo` (`pulumi.Input[str]`)
          * `replyToDisplayName` (`pulumi.Input[str]`)
          * `ssl` (`pulumi.Input[bool]`)
          * `starttls` (`pulumi.Input[bool]`)
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

            __props__['access_code_lifespan'] = access_code_lifespan
            __props__['access_code_lifespan_login'] = access_code_lifespan_login
            __props__['access_code_lifespan_user_action'] = access_code_lifespan_user_action
            __props__['access_token_lifespan'] = access_token_lifespan
            __props__['access_token_lifespan_for_implicit_flow'] = access_token_lifespan_for_implicit_flow
            __props__['account_theme'] = account_theme
            __props__['action_token_generated_by_admin_lifespan'] = action_token_generated_by_admin_lifespan
            __props__['action_token_generated_by_user_lifespan'] = action_token_generated_by_user_lifespan
            __props__['admin_theme'] = admin_theme
            __props__['attributes'] = attributes
            __props__['browser_flow'] = browser_flow
            __props__['client_authentication_flow'] = client_authentication_flow
            __props__['default_signature_algorithm'] = default_signature_algorithm
            __props__['direct_grant_flow'] = direct_grant_flow
            __props__['display_name'] = display_name
            __props__['display_name_html'] = display_name_html
            __props__['docker_authentication_flow'] = docker_authentication_flow
            __props__['duplicate_emails_allowed'] = duplicate_emails_allowed
            __props__['edit_username_allowed'] = edit_username_allowed
            __props__['email_theme'] = email_theme
            __props__['enabled'] = enabled
            __props__['internationalization'] = internationalization
            __props__['login_theme'] = login_theme
            __props__['login_with_email_allowed'] = login_with_email_allowed
            __props__['offline_session_idle_timeout'] = offline_session_idle_timeout
            __props__['offline_session_max_lifespan'] = offline_session_max_lifespan
            __props__['password_policy'] = password_policy
            if realm is None:
                raise TypeError("Missing required property 'realm'")
            __props__['realm'] = realm
            __props__['refresh_token_max_reuse'] = refresh_token_max_reuse
            __props__['registration_allowed'] = registration_allowed
            __props__['registration_email_as_username'] = registration_email_as_username
            __props__['registration_flow'] = registration_flow
            __props__['remember_me'] = remember_me
            __props__['reset_credentials_flow'] = reset_credentials_flow
            __props__['reset_password_allowed'] = reset_password_allowed
            __props__['revoke_refresh_token'] = revoke_refresh_token
            __props__['security_defenses'] = security_defenses
            __props__['smtp_server'] = smtp_server
            __props__['ssl_required'] = ssl_required
            __props__['sso_session_idle_timeout'] = sso_session_idle_timeout
            __props__['sso_session_max_lifespan'] = sso_session_max_lifespan
            __props__['user_managed_access'] = user_managed_access
            __props__['verify_email'] = verify_email
            __props__['internal_id'] = None
        super(Realm, __self__).__init__(
            'keycloak:index/realm:Realm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_code_lifespan=None, access_code_lifespan_login=None, access_code_lifespan_user_action=None, access_token_lifespan=None, access_token_lifespan_for_implicit_flow=None, account_theme=None, action_token_generated_by_admin_lifespan=None, action_token_generated_by_user_lifespan=None, admin_theme=None, attributes=None, browser_flow=None, client_authentication_flow=None, default_signature_algorithm=None, direct_grant_flow=None, display_name=None, display_name_html=None, docker_authentication_flow=None, duplicate_emails_allowed=None, edit_username_allowed=None, email_theme=None, enabled=None, internal_id=None, internationalization=None, login_theme=None, login_with_email_allowed=None, offline_session_idle_timeout=None, offline_session_max_lifespan=None, password_policy=None, realm=None, refresh_token_max_reuse=None, registration_allowed=None, registration_email_as_username=None, registration_flow=None, remember_me=None, reset_credentials_flow=None, reset_password_allowed=None, revoke_refresh_token=None, security_defenses=None, smtp_server=None, ssl_required=None, sso_session_idle_timeout=None, sso_session_max_lifespan=None, user_managed_access=None, verify_email=None):
        """
        Get an existing Realm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] browser_flow: Which flow should be used for BrowserFlow
        :param pulumi.Input[str] client_authentication_flow: Which flow should be used for ClientAuthenticationFlow
        :param pulumi.Input[str] direct_grant_flow: Which flow should be used for DirectGrantFlow
        :param pulumi.Input[str] docker_authentication_flow: Which flow should be used for DockerAuthenticationFlow
        :param pulumi.Input[str] password_policy: String that represents the passwordPolicies that are in place. Each policy is separated with " and ". Supported policies
               can be found in the server-info providers page. example: "upperCase(1) and length(8) and forceExpiredPasswordChange(365)
               and notUsername(undefined)"
        :param pulumi.Input[str] registration_flow: Which flow should be used for RegistrationFlow
        :param pulumi.Input[str] reset_credentials_flow: Which flow should be used for ResetCredentialsFlow
        :param pulumi.Input[str] ssl_required: SSL Required: Values can be 'none', 'external' or 'all'.

        The **internationalization** object supports the following:

          * `defaultLocale` (`pulumi.Input[str]`)
          * `supportedLocales` (`pulumi.Input[list]`)

        The **security_defenses** object supports the following:

          * `bruteForceDetection` (`pulumi.Input[dict]`)
            * `failureResetTimeSeconds` (`pulumi.Input[float]`)
            * `maxFailureWaitSeconds` (`pulumi.Input[float]`)
            * `maxLoginFailures` (`pulumi.Input[float]`)
            * `minimumQuickLoginWaitSeconds` (`pulumi.Input[float]`)
            * `permanentLockout` (`pulumi.Input[bool]`)
            * `quickLoginCheckMilliSeconds` (`pulumi.Input[float]`)
            * `waitIncrementSeconds` (`pulumi.Input[float]`)

          * `headers` (`pulumi.Input[dict]`)
            * `contentSecurityPolicy` (`pulumi.Input[str]`)
            * `contentSecurityPolicyReportOnly` (`pulumi.Input[str]`)
            * `strictTransportSecurity` (`pulumi.Input[str]`)
            * `xContentTypeOptions` (`pulumi.Input[str]`)
            * `xFrameOptions` (`pulumi.Input[str]`)
            * `xRobotsTag` (`pulumi.Input[str]`)
            * `xXssProtection` (`pulumi.Input[str]`)

        The **smtp_server** object supports the following:

          * `auth` (`pulumi.Input[dict]`)
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)

          * `envelopeFrom` (`pulumi.Input[str]`)
          * `from` (`pulumi.Input[str]`)
          * `fromDisplayName` (`pulumi.Input[str]`)
          * `host` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[str]`)
          * `replyTo` (`pulumi.Input[str]`)
          * `replyToDisplayName` (`pulumi.Input[str]`)
          * `ssl` (`pulumi.Input[bool]`)
          * `starttls` (`pulumi.Input[bool]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_code_lifespan"] = access_code_lifespan
        __props__["access_code_lifespan_login"] = access_code_lifespan_login
        __props__["access_code_lifespan_user_action"] = access_code_lifespan_user_action
        __props__["access_token_lifespan"] = access_token_lifespan
        __props__["access_token_lifespan_for_implicit_flow"] = access_token_lifespan_for_implicit_flow
        __props__["account_theme"] = account_theme
        __props__["action_token_generated_by_admin_lifespan"] = action_token_generated_by_admin_lifespan
        __props__["action_token_generated_by_user_lifespan"] = action_token_generated_by_user_lifespan
        __props__["admin_theme"] = admin_theme
        __props__["attributes"] = attributes
        __props__["browser_flow"] = browser_flow
        __props__["client_authentication_flow"] = client_authentication_flow
        __props__["default_signature_algorithm"] = default_signature_algorithm
        __props__["direct_grant_flow"] = direct_grant_flow
        __props__["display_name"] = display_name
        __props__["display_name_html"] = display_name_html
        __props__["docker_authentication_flow"] = docker_authentication_flow
        __props__["duplicate_emails_allowed"] = duplicate_emails_allowed
        __props__["edit_username_allowed"] = edit_username_allowed
        __props__["email_theme"] = email_theme
        __props__["enabled"] = enabled
        __props__["internal_id"] = internal_id
        __props__["internationalization"] = internationalization
        __props__["login_theme"] = login_theme
        __props__["login_with_email_allowed"] = login_with_email_allowed
        __props__["offline_session_idle_timeout"] = offline_session_idle_timeout
        __props__["offline_session_max_lifespan"] = offline_session_max_lifespan
        __props__["password_policy"] = password_policy
        __props__["realm"] = realm
        __props__["refresh_token_max_reuse"] = refresh_token_max_reuse
        __props__["registration_allowed"] = registration_allowed
        __props__["registration_email_as_username"] = registration_email_as_username
        __props__["registration_flow"] = registration_flow
        __props__["remember_me"] = remember_me
        __props__["reset_credentials_flow"] = reset_credentials_flow
        __props__["reset_password_allowed"] = reset_password_allowed
        __props__["revoke_refresh_token"] = revoke_refresh_token
        __props__["security_defenses"] = security_defenses
        __props__["smtp_server"] = smtp_server
        __props__["ssl_required"] = ssl_required
        __props__["sso_session_idle_timeout"] = sso_session_idle_timeout
        __props__["sso_session_max_lifespan"] = sso_session_max_lifespan
        __props__["user_managed_access"] = user_managed_access
        __props__["verify_email"] = verify_email
        return Realm(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
