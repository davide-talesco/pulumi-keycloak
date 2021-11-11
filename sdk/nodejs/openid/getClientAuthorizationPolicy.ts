// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can be used to fetch policy and permission information for an OpenID client that has authorization enabled.
 *
 * ## Example Usage
 *
 * In this example, we'll create a new OpenID client with authorization enabled. This will cause Keycloak to create a default
 * permission for this client called "Default Permission". We'll use the `keycloak.openid.getClientAuthorizationPolicy` data
 * source to fetch information about this permission, so we can use it to create a new resource-based authorization permission.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const clientWithAuthz = new keycloak.openid.Client("clientWithAuthz", {
 *     clientId: "client-with-authz",
 *     realmId: realm.id,
 *     accessType: "CONFIDENTIAL",
 *     serviceAccountsEnabled: true,
 *     authorization: {
 *         policyEnforcementMode: "ENFORCING",
 *     },
 * });
 * const defaultPermission = pulumi.all([realm.id, clientWithAuthz.resourceServerId]).apply(([id, resourceServerId]) => keycloak.openid.getClientAuthorizationPolicy({
 *     realmId: id,
 *     resourceServerId: resourceServerId,
 *     name: "Default Permission",
 * }));
 * const resource = new keycloak.openid.ClientAuthorizationResource("resource", {
 *     resourceServerId: clientWithAuthz.resourceServerId,
 *     realmId: realm.id,
 *     uris: ["/endpoint/*"],
 *     attributes: {
 *         foo: "bar",
 *     },
 * });
 * const permission = new keycloak.openid.ClientAuthorizationPermission("permission", {
 *     resourceServerId: clientWithAuthz.resourceServerId,
 *     realmId: realm.id,
 *     policies: [defaultPermission.apply(defaultPermission => defaultPermission.id)],
 *     resources: [resource.id],
 * });
 * ```
 */
export function getClientAuthorizationPolicy(args: GetClientAuthorizationPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetClientAuthorizationPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("keycloak:openid/getClientAuthorizationPolicy:getClientAuthorizationPolicy", {
        "name": args.name,
        "realmId": args.realmId,
        "resourceServerId": args.resourceServerId,
    }, opts);
}

/**
 * A collection of arguments for invoking getClientAuthorizationPolicy.
 */
export interface GetClientAuthorizationPolicyArgs {
    /**
     * The name of the authorization policy.
     */
    name: string;
    /**
     * The realm this authorization policy exists within.
     */
    realmId: string;
    /**
     * The ID of the resource server this authorization policy is attached to.
     */
    resourceServerId: string;
}

/**
 * A collection of values returned by getClientAuthorizationPolicy.
 */
export interface GetClientAuthorizationPolicyResult {
    /**
     * (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
     */
    readonly decisionStrategy: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
     */
    readonly logic: string;
    readonly name: string;
    /**
     * (Computed) The ID of the owning resource. Applies to resources.
     */
    readonly owner: string;
    /**
     * (Computed) The IDs of the policies that must be applied to scopes/resources for this policy/permission. Applies to policies and permissions.
     */
    readonly policies: string[];
    readonly realmId: string;
    readonly resourceServerId: string;
    /**
     * (Computed) The IDs of the resources that this permission applies to. Applies to resource-based permissions.
     */
    readonly resources: string[];
    /**
     * (Computed) The IDs of the scopes that this permission applies to. Applies to scope-based permissions.
     */
    readonly scopes: string[];
    /**
     * (Computed) The type of this policy / permission. For permissions, this could be `resource` or `scope`. For policies, this could be any type of authorization policy, such as `js`.
     */
    readonly type: string;
}

export function getClientAuthorizationPolicyOutput(args: GetClientAuthorizationPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClientAuthorizationPolicyResult> {
    return pulumi.output(args).apply(a => getClientAuthorizationPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getClientAuthorizationPolicy.
 */
export interface GetClientAuthorizationPolicyOutputArgs {
    /**
     * The name of the authorization policy.
     */
    name: pulumi.Input<string>;
    /**
     * The realm this authorization policy exists within.
     */
    realmId: pulumi.Input<string>;
    /**
     * The ID of the resource server this authorization policy is attached to.
     */
    resourceServerId: pulumi.Input<string>;
}
