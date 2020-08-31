# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Realm']


class Realm(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_code_lifespan: Optional[pulumi.Input[str]] = None,
                 access_code_lifespan_login: Optional[pulumi.Input[str]] = None,
                 access_code_lifespan_user_action: Optional[pulumi.Input[str]] = None,
                 access_token_lifespan: Optional[pulumi.Input[str]] = None,
                 access_token_lifespan_for_implicit_flow: Optional[pulumi.Input[str]] = None,
                 account_theme: Optional[pulumi.Input[str]] = None,
                 action_token_generated_by_admin_lifespan: Optional[pulumi.Input[str]] = None,
                 action_token_generated_by_user_lifespan: Optional[pulumi.Input[str]] = None,
                 admin_theme: Optional[pulumi.Input[str]] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 browser_flow: Optional[pulumi.Input[str]] = None,
                 client_authentication_flow: Optional[pulumi.Input[str]] = None,
                 default_signature_algorithm: Optional[pulumi.Input[str]] = None,
                 direct_grant_flow: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 display_name_html: Optional[pulumi.Input[str]] = None,
                 docker_authentication_flow: Optional[pulumi.Input[str]] = None,
                 duplicate_emails_allowed: Optional[pulumi.Input[bool]] = None,
                 edit_username_allowed: Optional[pulumi.Input[bool]] = None,
                 email_theme: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 internationalization: Optional[pulumi.Input[pulumi.InputType['RealmInternationalizationArgs']]] = None,
                 login_theme: Optional[pulumi.Input[str]] = None,
                 login_with_email_allowed: Optional[pulumi.Input[bool]] = None,
                 offline_session_idle_timeout: Optional[pulumi.Input[str]] = None,
                 offline_session_max_lifespan: Optional[pulumi.Input[str]] = None,
                 password_policy: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 refresh_token_max_reuse: Optional[pulumi.Input[float]] = None,
                 registration_allowed: Optional[pulumi.Input[bool]] = None,
                 registration_email_as_username: Optional[pulumi.Input[bool]] = None,
                 registration_flow: Optional[pulumi.Input[str]] = None,
                 remember_me: Optional[pulumi.Input[bool]] = None,
                 reset_credentials_flow: Optional[pulumi.Input[str]] = None,
                 reset_password_allowed: Optional[pulumi.Input[bool]] = None,
                 revoke_refresh_token: Optional[pulumi.Input[bool]] = None,
                 security_defenses: Optional[pulumi.Input[pulumi.InputType['RealmSecurityDefensesArgs']]] = None,
                 smtp_server: Optional[pulumi.Input[pulumi.InputType['RealmSmtpServerArgs']]] = None,
                 ssl_required: Optional[pulumi.Input[str]] = None,
                 sso_session_idle_timeout: Optional[pulumi.Input[str]] = None,
                 sso_session_max_lifespan: Optional[pulumi.Input[str]] = None,
                 user_managed_access: Optional[pulumi.Input[bool]] = None,
                 verify_email: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_code_lifespan: Optional[pulumi.Input[str]] = None,
            access_code_lifespan_login: Optional[pulumi.Input[str]] = None,
            access_code_lifespan_user_action: Optional[pulumi.Input[str]] = None,
            access_token_lifespan: Optional[pulumi.Input[str]] = None,
            access_token_lifespan_for_implicit_flow: Optional[pulumi.Input[str]] = None,
            account_theme: Optional[pulumi.Input[str]] = None,
            action_token_generated_by_admin_lifespan: Optional[pulumi.Input[str]] = None,
            action_token_generated_by_user_lifespan: Optional[pulumi.Input[str]] = None,
            admin_theme: Optional[pulumi.Input[str]] = None,
            attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            browser_flow: Optional[pulumi.Input[str]] = None,
            client_authentication_flow: Optional[pulumi.Input[str]] = None,
            default_signature_algorithm: Optional[pulumi.Input[str]] = None,
            direct_grant_flow: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            display_name_html: Optional[pulumi.Input[str]] = None,
            docker_authentication_flow: Optional[pulumi.Input[str]] = None,
            duplicate_emails_allowed: Optional[pulumi.Input[bool]] = None,
            edit_username_allowed: Optional[pulumi.Input[bool]] = None,
            email_theme: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            internal_id: Optional[pulumi.Input[str]] = None,
            internationalization: Optional[pulumi.Input[pulumi.InputType['RealmInternationalizationArgs']]] = None,
            login_theme: Optional[pulumi.Input[str]] = None,
            login_with_email_allowed: Optional[pulumi.Input[bool]] = None,
            offline_session_idle_timeout: Optional[pulumi.Input[str]] = None,
            offline_session_max_lifespan: Optional[pulumi.Input[str]] = None,
            password_policy: Optional[pulumi.Input[str]] = None,
            realm: Optional[pulumi.Input[str]] = None,
            refresh_token_max_reuse: Optional[pulumi.Input[float]] = None,
            registration_allowed: Optional[pulumi.Input[bool]] = None,
            registration_email_as_username: Optional[pulumi.Input[bool]] = None,
            registration_flow: Optional[pulumi.Input[str]] = None,
            remember_me: Optional[pulumi.Input[bool]] = None,
            reset_credentials_flow: Optional[pulumi.Input[str]] = None,
            reset_password_allowed: Optional[pulumi.Input[bool]] = None,
            revoke_refresh_token: Optional[pulumi.Input[bool]] = None,
            security_defenses: Optional[pulumi.Input[pulumi.InputType['RealmSecurityDefensesArgs']]] = None,
            smtp_server: Optional[pulumi.Input[pulumi.InputType['RealmSmtpServerArgs']]] = None,
            ssl_required: Optional[pulumi.Input[str]] = None,
            sso_session_idle_timeout: Optional[pulumi.Input[str]] = None,
            sso_session_max_lifespan: Optional[pulumi.Input[str]] = None,
            user_managed_access: Optional[pulumi.Input[bool]] = None,
            verify_email: Optional[pulumi.Input[bool]] = None) -> 'Realm':
        """
        Get an existing Realm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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

    @property
    @pulumi.getter(name="accessCodeLifespan")
    def access_code_lifespan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "access_code_lifespan")

    @property
    @pulumi.getter(name="accessCodeLifespanLogin")
    def access_code_lifespan_login(self) -> pulumi.Output[str]:
        return pulumi.get(self, "access_code_lifespan_login")

    @property
    @pulumi.getter(name="accessCodeLifespanUserAction")
    def access_code_lifespan_user_action(self) -> pulumi.Output[str]:
        return pulumi.get(self, "access_code_lifespan_user_action")

    @property
    @pulumi.getter(name="accessTokenLifespan")
    def access_token_lifespan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "access_token_lifespan")

    @property
    @pulumi.getter(name="accessTokenLifespanForImplicitFlow")
    def access_token_lifespan_for_implicit_flow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "access_token_lifespan_for_implicit_flow")

    @property
    @pulumi.getter(name="accountTheme")
    def account_theme(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "account_theme")

    @property
    @pulumi.getter(name="actionTokenGeneratedByAdminLifespan")
    def action_token_generated_by_admin_lifespan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "action_token_generated_by_admin_lifespan")

    @property
    @pulumi.getter(name="actionTokenGeneratedByUserLifespan")
    def action_token_generated_by_user_lifespan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "action_token_generated_by_user_lifespan")

    @property
    @pulumi.getter(name="adminTheme")
    def admin_theme(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "admin_theme")

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="browserFlow")
    def browser_flow(self) -> pulumi.Output[Optional[str]]:
        """
        Which flow should be used for BrowserFlow
        """
        return pulumi.get(self, "browser_flow")

    @property
    @pulumi.getter(name="clientAuthenticationFlow")
    def client_authentication_flow(self) -> pulumi.Output[Optional[str]]:
        """
        Which flow should be used for ClientAuthenticationFlow
        """
        return pulumi.get(self, "client_authentication_flow")

    @property
    @pulumi.getter(name="defaultSignatureAlgorithm")
    def default_signature_algorithm(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "default_signature_algorithm")

    @property
    @pulumi.getter(name="directGrantFlow")
    def direct_grant_flow(self) -> pulumi.Output[Optional[str]]:
        """
        Which flow should be used for DirectGrantFlow
        """
        return pulumi.get(self, "direct_grant_flow")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="displayNameHtml")
    def display_name_html(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "display_name_html")

    @property
    @pulumi.getter(name="dockerAuthenticationFlow")
    def docker_authentication_flow(self) -> pulumi.Output[Optional[str]]:
        """
        Which flow should be used for DockerAuthenticationFlow
        """
        return pulumi.get(self, "docker_authentication_flow")

    @property
    @pulumi.getter(name="duplicateEmailsAllowed")
    def duplicate_emails_allowed(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "duplicate_emails_allowed")

    @property
    @pulumi.getter(name="editUsernameAllowed")
    def edit_username_allowed(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "edit_username_allowed")

    @property
    @pulumi.getter(name="emailTheme")
    def email_theme(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "email_theme")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter
    def internationalization(self) -> pulumi.Output[Optional['outputs.RealmInternationalization']]:
        return pulumi.get(self, "internationalization")

    @property
    @pulumi.getter(name="loginTheme")
    def login_theme(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "login_theme")

    @property
    @pulumi.getter(name="loginWithEmailAllowed")
    def login_with_email_allowed(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "login_with_email_allowed")

    @property
    @pulumi.getter(name="offlineSessionIdleTimeout")
    def offline_session_idle_timeout(self) -> pulumi.Output[str]:
        return pulumi.get(self, "offline_session_idle_timeout")

    @property
    @pulumi.getter(name="offlineSessionMaxLifespan")
    def offline_session_max_lifespan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "offline_session_max_lifespan")

    @property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> pulumi.Output[Optional[str]]:
        """
        String that represents the passwordPolicies that are in place. Each policy is separated with " and ". Supported policies
        can be found in the server-info providers page. example: "upperCase(1) and length(8) and forceExpiredPasswordChange(365)
        and notUsername(undefined)"
        """
        return pulumi.get(self, "password_policy")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter(name="refreshTokenMaxReuse")
    def refresh_token_max_reuse(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "refresh_token_max_reuse")

    @property
    @pulumi.getter(name="registrationAllowed")
    def registration_allowed(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "registration_allowed")

    @property
    @pulumi.getter(name="registrationEmailAsUsername")
    def registration_email_as_username(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "registration_email_as_username")

    @property
    @pulumi.getter(name="registrationFlow")
    def registration_flow(self) -> pulumi.Output[Optional[str]]:
        """
        Which flow should be used for RegistrationFlow
        """
        return pulumi.get(self, "registration_flow")

    @property
    @pulumi.getter(name="rememberMe")
    def remember_me(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "remember_me")

    @property
    @pulumi.getter(name="resetCredentialsFlow")
    def reset_credentials_flow(self) -> pulumi.Output[Optional[str]]:
        """
        Which flow should be used for ResetCredentialsFlow
        """
        return pulumi.get(self, "reset_credentials_flow")

    @property
    @pulumi.getter(name="resetPasswordAllowed")
    def reset_password_allowed(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "reset_password_allowed")

    @property
    @pulumi.getter(name="revokeRefreshToken")
    def revoke_refresh_token(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "revoke_refresh_token")

    @property
    @pulumi.getter(name="securityDefenses")
    def security_defenses(self) -> pulumi.Output[Optional['outputs.RealmSecurityDefenses']]:
        return pulumi.get(self, "security_defenses")

    @property
    @pulumi.getter(name="smtpServer")
    def smtp_server(self) -> pulumi.Output[Optional['outputs.RealmSmtpServer']]:
        return pulumi.get(self, "smtp_server")

    @property
    @pulumi.getter(name="sslRequired")
    def ssl_required(self) -> pulumi.Output[Optional[str]]:
        """
        SSL Required: Values can be 'none', 'external' or 'all'.
        """
        return pulumi.get(self, "ssl_required")

    @property
    @pulumi.getter(name="ssoSessionIdleTimeout")
    def sso_session_idle_timeout(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sso_session_idle_timeout")

    @property
    @pulumi.getter(name="ssoSessionMaxLifespan")
    def sso_session_max_lifespan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sso_session_max_lifespan")

    @property
    @pulumi.getter(name="userManagedAccess")
    def user_managed_access(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "user_managed_access")

    @property
    @pulumi.getter(name="verifyEmail")
    def verify_email(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "verify_email")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

