// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for assigning realm roles to the service account of an openid client.
 * You need to set `serviceAccountsEnabled` to `true` for the openid client that should be assigned the role.
 *
 * If you'd like to attach client roles to a service account, please use the `keycloak.openid.ClientServiceAccountRole`
 * resource.
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
 * const realmRole = new keycloak.Role("realmRole", {realmId: realm.id});
 * const client = new keycloak.openid.Client("client", {
 *     realmId: realm.id,
 *     serviceAccountsEnabled: true,
 * });
 * const clientServiceAccountRole = new keycloak.openid.ClientServiceAccountRealmRole("clientServiceAccountRole", {
 *     realmId: realm.id,
 *     serviceAccountUserId: client.serviceAccountUserId,
 *     role: realmRole.name,
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{roleId}}`. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole client_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/c7230ab7-8e4e-4135-995d-e81b50696ad8
 * ```
 */
export class ClientServiceAccountRealmRole extends pulumi.CustomResource {
    /**
     * Get an existing ClientServiceAccountRealmRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientServiceAccountRealmRoleState, opts?: pulumi.CustomResourceOptions): ClientServiceAccountRealmRole {
        return new ClientServiceAccountRealmRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole';

    /**
     * Returns true if the given object is an instance of ClientServiceAccountRealmRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientServiceAccountRealmRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientServiceAccountRealmRole.__pulumiType;
    }

    /**
     * The realm that the client and role belong to.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The name of the role that is assigned.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
     */
    public readonly serviceAccountUserId!: pulumi.Output<string>;

    /**
     * Create a ClientServiceAccountRealmRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientServiceAccountRealmRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientServiceAccountRealmRoleArgs | ClientServiceAccountRealmRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientServiceAccountRealmRoleState | undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["serviceAccountUserId"] = state ? state.serviceAccountUserId : undefined;
        } else {
            const args = argsOrState as ClientServiceAccountRealmRoleArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.serviceAccountUserId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountUserId'");
            }
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["serviceAccountUserId"] = args ? args.serviceAccountUserId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClientServiceAccountRealmRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientServiceAccountRealmRole resources.
 */
export interface ClientServiceAccountRealmRoleState {
    /**
     * The realm that the client and role belong to.
     */
    readonly realmId?: pulumi.Input<string>;
    /**
     * The name of the role that is assigned.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
     */
    readonly serviceAccountUserId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientServiceAccountRealmRole resource.
 */
export interface ClientServiceAccountRealmRoleArgs {
    /**
     * The realm that the client and role belong to.
     */
    readonly realmId: pulumi.Input<string>;
    /**
     * The name of the role that is assigned.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
     */
    readonly serviceAccountUserId: pulumi.Input<string>;
}
