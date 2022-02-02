// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can be used to fetch the realm roles of a user within Keycloak.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const masterRealm = keycloak.getRealm({
 *     realm: "master",
 * });
 * const defaultAdminUser = masterRealm.then(masterRealm => keycloak.getUser({
 *     realmId: masterRealm.id,
 *     username: "keycloak",
 * }));
 * const userRealmRoles = Promise.all([masterRealm, defaultAdminUser]).then(([masterRealm, defaultAdminUser]) => keycloak.getUserRealmRoles({
 *     realmId: masterRealm.id,
 *     userId: defaultAdminUser.id,
 * }));
 * export const keycloakUserRoleNames = userRealmRoles.then(userRealmRoles => userRealmRoles.roleNames);
 * ```
 */
export function getUserRealmRoles(args: GetUserRealmRolesArgs, opts?: pulumi.InvokeOptions): Promise<GetUserRealmRolesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("keycloak:index/getUserRealmRoles:getUserRealmRoles", {
        "realmId": args.realmId,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserRealmRoles.
 */
export interface GetUserRealmRolesArgs {
    /**
     * The realm this user belongs to.
     */
    realmId: string;
    /**
     * The ID of the user to query realm roles for.
     */
    userId: string;
}

/**
 * A collection of values returned by getUserRealmRoles.
 */
export interface GetUserRealmRolesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly realmId: string;
    /**
     * (Computed) A list of realm roles that belong to this user.
     */
    readonly roleNames: string[];
    readonly userId: string;
}

export function getUserRealmRolesOutput(args: GetUserRealmRolesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserRealmRolesResult> {
    return pulumi.output(args).apply(a => getUserRealmRoles(a, opts))
}

/**
 * A collection of arguments for invoking getUserRealmRoles.
 */
export interface GetUserRealmRolesOutputArgs {
    /**
     * The realm this user belongs to.
     */
    realmId: pulumi.Input<string>;
    /**
     * The ID of the user to query realm roles for.
     */
    userId: pulumi.Input<string>;
}
