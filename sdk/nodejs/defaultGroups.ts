// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for managing a realm's default groups.
 *
 * > You should not use `keycloak.DefaultGroups` with a group whose members are managed by `keycloak.GroupMemberships`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const group = new keycloak.Group("group", {realmId: realm.id});
 * const _default = new keycloak.DefaultGroups("default", {
 *     realmId: realm.id,
 *     groupIds: [group.id],
 * });
 * ```
 *
 * ## Import
 *
 * Default groups can be imported using the format `{{realm_id}}` where `realm_id` is the realm the group exists in. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:index/defaultGroups:DefaultGroups default my-realm
 * ```
 */
export class DefaultGroups extends pulumi.CustomResource {
    /**
     * Get an existing DefaultGroups resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultGroupsState, opts?: pulumi.CustomResourceOptions): DefaultGroups {
        return new DefaultGroups(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/defaultGroups:DefaultGroups';

    /**
     * Returns true if the given object is an instance of DefaultGroups.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultGroups {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultGroups.__pulumiType;
    }

    /**
     * A set of group ids that should be default groups on the realm referenced by `realmId`.
     */
    public readonly groupIds!: pulumi.Output<string[]>;
    /**
     * The realm this group exists in.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a DefaultGroups resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultGroupsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultGroupsArgs | DefaultGroupsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultGroupsState | undefined;
            inputs["groupIds"] = state ? state.groupIds : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as DefaultGroupsArgs | undefined;
            if ((!args || args.groupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupIds'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["groupIds"] = args ? args.groupIds : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DefaultGroups.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultGroups resources.
 */
export interface DefaultGroupsState {
    /**
     * A set of group ids that should be default groups on the realm referenced by `realmId`.
     */
    readonly groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The realm this group exists in.
     */
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefaultGroups resource.
 */
export interface DefaultGroupsArgs {
    /**
     * A set of group ids that should be default groups on the realm referenced by `realmId`.
     */
    readonly groupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The realm this group exists in.
     */
    readonly realmId: pulumi.Input<string>;
}
