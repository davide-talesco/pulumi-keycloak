// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # keycloak..Group data source
 *
 * This data source can be used to fetch properties of a Keycloak group for
 * usage with other resources, such as `keycloak..GroupRoles`.
 *
 * ### Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const offlineAccess = realm.id.apply(id => keycloak.getRole({
 *     name: "offlineAccess",
 *     realmId: id,
 * }, { async: true }));
 * const group = realm.id.apply(id => keycloak.getGroup({
 *     name: "group",
 *     realmId: id,
 * }, { async: true }));
 * const groupRoles = new keycloak.GroupRoles("groupRoles", {
 *     groupId: group.id,
 *     realmId: realm.id,
 *     roles: [offlineAccess.id],
 * });
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this group exists within.
 * - `name` - (Required) The name of the group
 *
 * ### Attributes Reference
 *
 * In addition to the arguments listed above, the following computed attributes are exported:
 *
 * - `id` - The unique ID of the group, which can be used as an argument to
 *   other resources supported by this provider.
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("keycloak:index/getGroup:getGroup", {
        "name": args.name,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    readonly name: string;
    readonly realmId: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    readonly name: string;
    readonly realmId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
