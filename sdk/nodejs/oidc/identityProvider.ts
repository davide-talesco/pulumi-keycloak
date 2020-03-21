// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class IdentityProvider extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityProviderState, opts?: pulumi.CustomResourceOptions): IdentityProvider {
        return new IdentityProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:oidc/identityProvider:IdentityProvider';

    /**
     * Returns true if the given object is an instance of IdentityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProvider.__pulumiType;
    }

    /**
     * This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity
     * provider. In case that client sends a request with prompt=none and user is not yet authenticated, the error will not
     * be directly returned to client, but the request with prompt=none will be forwarded to this identity provider.
     */
    public readonly acceptsPromptNoneForwardFromClient!: pulumi.Output<boolean | undefined>;
    /**
     * Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
     */
    public readonly addReadTokenRoleOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    public readonly authenticateByDefault!: pulumi.Output<boolean | undefined>;
    /**
     * OIDC authorization URL.
     */
    public readonly authorizationUrl!: pulumi.Output<string>;
    /**
     * Does the external IDP support backchannel logout?
     */
    public readonly backchannelSupported!: pulumi.Output<boolean | undefined>;
    /**
     * Client ID.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client Secret.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to
     * 'openid'.
     */
    public readonly defaultScopes!: pulumi.Output<string | undefined>;
    /**
     * Friendly name for Identity Providers.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable this identity provider.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly extraConfig!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login'
     * means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
     */
    public readonly firstBrokerLoginFlowAlias!: pulumi.Output<string | undefined>;
    /**
     * Hide On Login Page.
     */
    public readonly hideOnLoginPage!: pulumi.Output<boolean | undefined>;
    /**
     * Internal Identity Provider Id
     */
    public /*out*/ readonly internalId!: pulumi.Output<string>;
    /**
     * JSON Web Key Set URL
     */
    public readonly jwksUrl!: pulumi.Output<string | undefined>;
    /**
     * If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
     * want to allow login from the provider, but want to integrate with a provider
     */
    public readonly linkOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Login Hint.
     */
    public readonly loginHint!: pulumi.Output<string | undefined>;
    /**
     * Logout URL
     */
    public readonly logoutUrl!: pulumi.Output<string | undefined>;
    /**
     * Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
     * additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty
     * if you don't want any additional authenticators to be triggered after login with this identity provider. Also note,
     * that authenticator implementations must assume that user is already set in ClientSession as identity provider
     * already set it.
     */
    public readonly postBrokerLoginFlowAlias!: pulumi.Output<string | undefined>;
    /**
     * provider id, is always oidc, unless you have a custom implementation
     */
    public readonly providerId!: pulumi.Output<string | undefined>;
    /**
     * Realm Name
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * Enable/disable if tokens must be stored after authenticating users.
     */
    public readonly storeToken!: pulumi.Output<boolean | undefined>;
    /**
     * Token URL.
     */
    public readonly tokenUrl!: pulumi.Output<string>;
    /**
     * If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
     */
    public readonly trustEmail!: pulumi.Output<boolean | undefined>;
    /**
     * Pass current locale to identity provider
     */
    public readonly uiLocales!: pulumi.Output<boolean | undefined>;
    /**
     * User Info URL
     */
    public readonly userInfoUrl!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable signature validation of external IDP signatures.
     */
    public readonly validateSignature!: pulumi.Output<boolean | undefined>;

    /**
     * Create a IdentityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityProviderArgs | IdentityProviderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IdentityProviderState | undefined;
            inputs["acceptsPromptNoneForwardFromClient"] = state ? state.acceptsPromptNoneForwardFromClient : undefined;
            inputs["addReadTokenRoleOnCreate"] = state ? state.addReadTokenRoleOnCreate : undefined;
            inputs["alias"] = state ? state.alias : undefined;
            inputs["authenticateByDefault"] = state ? state.authenticateByDefault : undefined;
            inputs["authorizationUrl"] = state ? state.authorizationUrl : undefined;
            inputs["backchannelSupported"] = state ? state.backchannelSupported : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["defaultScopes"] = state ? state.defaultScopes : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["extraConfig"] = state ? state.extraConfig : undefined;
            inputs["firstBrokerLoginFlowAlias"] = state ? state.firstBrokerLoginFlowAlias : undefined;
            inputs["hideOnLoginPage"] = state ? state.hideOnLoginPage : undefined;
            inputs["internalId"] = state ? state.internalId : undefined;
            inputs["jwksUrl"] = state ? state.jwksUrl : undefined;
            inputs["linkOnly"] = state ? state.linkOnly : undefined;
            inputs["loginHint"] = state ? state.loginHint : undefined;
            inputs["logoutUrl"] = state ? state.logoutUrl : undefined;
            inputs["postBrokerLoginFlowAlias"] = state ? state.postBrokerLoginFlowAlias : undefined;
            inputs["providerId"] = state ? state.providerId : undefined;
            inputs["realm"] = state ? state.realm : undefined;
            inputs["storeToken"] = state ? state.storeToken : undefined;
            inputs["tokenUrl"] = state ? state.tokenUrl : undefined;
            inputs["trustEmail"] = state ? state.trustEmail : undefined;
            inputs["uiLocales"] = state ? state.uiLocales : undefined;
            inputs["userInfoUrl"] = state ? state.userInfoUrl : undefined;
            inputs["validateSignature"] = state ? state.validateSignature : undefined;
        } else {
            const args = argsOrState as IdentityProviderArgs | undefined;
            if (!args || args.alias === undefined) {
                throw new Error("Missing required property 'alias'");
            }
            if (!args || args.authorizationUrl === undefined) {
                throw new Error("Missing required property 'authorizationUrl'");
            }
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.clientSecret === undefined) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if (!args || args.realm === undefined) {
                throw new Error("Missing required property 'realm'");
            }
            if (!args || args.tokenUrl === undefined) {
                throw new Error("Missing required property 'tokenUrl'");
            }
            inputs["acceptsPromptNoneForwardFromClient"] = args ? args.acceptsPromptNoneForwardFromClient : undefined;
            inputs["addReadTokenRoleOnCreate"] = args ? args.addReadTokenRoleOnCreate : undefined;
            inputs["alias"] = args ? args.alias : undefined;
            inputs["authenticateByDefault"] = args ? args.authenticateByDefault : undefined;
            inputs["authorizationUrl"] = args ? args.authorizationUrl : undefined;
            inputs["backchannelSupported"] = args ? args.backchannelSupported : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["defaultScopes"] = args ? args.defaultScopes : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["extraConfig"] = args ? args.extraConfig : undefined;
            inputs["firstBrokerLoginFlowAlias"] = args ? args.firstBrokerLoginFlowAlias : undefined;
            inputs["hideOnLoginPage"] = args ? args.hideOnLoginPage : undefined;
            inputs["jwksUrl"] = args ? args.jwksUrl : undefined;
            inputs["linkOnly"] = args ? args.linkOnly : undefined;
            inputs["loginHint"] = args ? args.loginHint : undefined;
            inputs["logoutUrl"] = args ? args.logoutUrl : undefined;
            inputs["postBrokerLoginFlowAlias"] = args ? args.postBrokerLoginFlowAlias : undefined;
            inputs["providerId"] = args ? args.providerId : undefined;
            inputs["realm"] = args ? args.realm : undefined;
            inputs["storeToken"] = args ? args.storeToken : undefined;
            inputs["tokenUrl"] = args ? args.tokenUrl : undefined;
            inputs["trustEmail"] = args ? args.trustEmail : undefined;
            inputs["uiLocales"] = args ? args.uiLocales : undefined;
            inputs["userInfoUrl"] = args ? args.userInfoUrl : undefined;
            inputs["validateSignature"] = args ? args.validateSignature : undefined;
            inputs["internalId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IdentityProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityProvider resources.
 */
export interface IdentityProviderState {
    /**
     * This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity
     * provider. In case that client sends a request with prompt=none and user is not yet authenticated, the error will not
     * be directly returned to client, but the request with prompt=none will be forwarded to this identity provider.
     */
    readonly acceptsPromptNoneForwardFromClient?: pulumi.Input<boolean>;
    /**
     * Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
     */
    readonly addReadTokenRoleOnCreate?: pulumi.Input<boolean>;
    /**
     * The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
     */
    readonly alias?: pulumi.Input<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    readonly authenticateByDefault?: pulumi.Input<boolean>;
    /**
     * OIDC authorization URL.
     */
    readonly authorizationUrl?: pulumi.Input<string>;
    /**
     * Does the external IDP support backchannel logout?
     */
    readonly backchannelSupported?: pulumi.Input<boolean>;
    /**
     * Client ID.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Client Secret.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to
     * 'openid'.
     */
    readonly defaultScopes?: pulumi.Input<string>;
    /**
     * Friendly name for Identity Providers.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Enable/disable this identity provider.
     */
    readonly enabled?: pulumi.Input<boolean>;
    readonly extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login'
     * means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
     */
    readonly firstBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * Hide On Login Page.
     */
    readonly hideOnLoginPage?: pulumi.Input<boolean>;
    /**
     * Internal Identity Provider Id
     */
    readonly internalId?: pulumi.Input<string>;
    /**
     * JSON Web Key Set URL
     */
    readonly jwksUrl?: pulumi.Input<string>;
    /**
     * If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
     * want to allow login from the provider, but want to integrate with a provider
     */
    readonly linkOnly?: pulumi.Input<boolean>;
    /**
     * Login Hint.
     */
    readonly loginHint?: pulumi.Input<string>;
    /**
     * Logout URL
     */
    readonly logoutUrl?: pulumi.Input<string>;
    /**
     * Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
     * additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty
     * if you don't want any additional authenticators to be triggered after login with this identity provider. Also note,
     * that authenticator implementations must assume that user is already set in ClientSession as identity provider
     * already set it.
     */
    readonly postBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * provider id, is always oidc, unless you have a custom implementation
     */
    readonly providerId?: pulumi.Input<string>;
    /**
     * Realm Name
     */
    readonly realm?: pulumi.Input<string>;
    /**
     * Enable/disable if tokens must be stored after authenticating users.
     */
    readonly storeToken?: pulumi.Input<boolean>;
    /**
     * Token URL.
     */
    readonly tokenUrl?: pulumi.Input<string>;
    /**
     * If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
     */
    readonly trustEmail?: pulumi.Input<boolean>;
    /**
     * Pass current locale to identity provider
     */
    readonly uiLocales?: pulumi.Input<boolean>;
    /**
     * User Info URL
     */
    readonly userInfoUrl?: pulumi.Input<string>;
    /**
     * Enable/disable signature validation of external IDP signatures.
     */
    readonly validateSignature?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a IdentityProvider resource.
 */
export interface IdentityProviderArgs {
    /**
     * This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity
     * provider. In case that client sends a request with prompt=none and user is not yet authenticated, the error will not
     * be directly returned to client, but the request with prompt=none will be forwarded to this identity provider.
     */
    readonly acceptsPromptNoneForwardFromClient?: pulumi.Input<boolean>;
    /**
     * Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
     */
    readonly addReadTokenRoleOnCreate?: pulumi.Input<boolean>;
    /**
     * The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
     */
    readonly alias: pulumi.Input<string>;
    /**
     * Enable/disable authenticate users by default.
     */
    readonly authenticateByDefault?: pulumi.Input<boolean>;
    /**
     * OIDC authorization URL.
     */
    readonly authorizationUrl: pulumi.Input<string>;
    /**
     * Does the external IDP support backchannel logout?
     */
    readonly backchannelSupported?: pulumi.Input<boolean>;
    /**
     * Client ID.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * Client Secret.
     */
    readonly clientSecret: pulumi.Input<string>;
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to
     * 'openid'.
     */
    readonly defaultScopes?: pulumi.Input<string>;
    /**
     * Friendly name for Identity Providers.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Enable/disable this identity provider.
     */
    readonly enabled?: pulumi.Input<boolean>;
    readonly extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login'
     * means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
     */
    readonly firstBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * Hide On Login Page.
     */
    readonly hideOnLoginPage?: pulumi.Input<boolean>;
    /**
     * JSON Web Key Set URL
     */
    readonly jwksUrl?: pulumi.Input<string>;
    /**
     * If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
     * want to allow login from the provider, but want to integrate with a provider
     */
    readonly linkOnly?: pulumi.Input<boolean>;
    /**
     * Login Hint.
     */
    readonly loginHint?: pulumi.Input<string>;
    /**
     * Logout URL
     */
    readonly logoutUrl?: pulumi.Input<string>;
    /**
     * Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
     * additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty
     * if you don't want any additional authenticators to be triggered after login with this identity provider. Also note,
     * that authenticator implementations must assume that user is already set in ClientSession as identity provider
     * already set it.
     */
    readonly postBrokerLoginFlowAlias?: pulumi.Input<string>;
    /**
     * provider id, is always oidc, unless you have a custom implementation
     */
    readonly providerId?: pulumi.Input<string>;
    /**
     * Realm Name
     */
    readonly realm: pulumi.Input<string>;
    /**
     * Enable/disable if tokens must be stored after authenticating users.
     */
    readonly storeToken?: pulumi.Input<boolean>;
    /**
     * Token URL.
     */
    readonly tokenUrl: pulumi.Input<string>;
    /**
     * If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
     */
    readonly trustEmail?: pulumi.Input<boolean>;
    /**
     * Pass current locale to identity provider
     */
    readonly uiLocales?: pulumi.Input<boolean>;
    /**
     * User Info URL
     */
    readonly userInfoUrl?: pulumi.Input<string>;
    /**
     * Enable/disable signature validation of external IDP signatures.
     */
    readonly validateSignature?: pulumi.Input<boolean>;
}
