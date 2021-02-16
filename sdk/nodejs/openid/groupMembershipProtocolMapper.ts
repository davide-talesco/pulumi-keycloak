// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing group membership protocol mappers within Keycloak.
 *
 * Group membership protocol mappers allow you to map a user's group memberships to a claim in a token.
 *
 * Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
 * multiple different clients.
 *
 * ## Example Usage
 * ### Client)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const openidClient = new keycloak.openid.Client("openidClient", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     enabled: true,
 *     accessType: "CONFIDENTIAL",
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const groupMembershipMapper = new keycloak.openid.GroupMembershipProtocolMapper("groupMembershipMapper", {
 *     realmId: realm.id,
 *     clientId: openidClient.id,
 *     claimName: "groups",
 * });
 * ```
 * ### Client Scope)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const clientScope = new keycloak.openid.ClientScope("clientScope", {realmId: realm.id});
 * const groupMembershipMapper = new keycloak.openid.GroupMembershipProtocolMapper("groupMembershipMapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 *     claimName: "groups",
 * });
 * ```
 *
 * ## Import
 *
 * Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper group_membership_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 *
 * ```sh
 *  $ pulumi import keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper group_membership_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 */
export class GroupMembershipProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing GroupMembershipProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMembershipProtocolMapperState, opts?: pulumi.CustomResourceOptions): GroupMembershipProtocolMapper {
        return new GroupMembershipProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper';

    /**
     * Returns true if the given object is an instance of GroupMembershipProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMembershipProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMembershipProtocolMapper.__pulumiType;
    }

    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    public readonly addToAccessToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    public readonly addToIdToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     */
    public readonly addToUserinfo!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the claim to insert into a token.
     */
    public readonly claimName!: pulumi.Output<string>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
     */
    public readonly fullPath!: pulumi.Output<boolean | undefined>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a GroupMembershipProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMembershipProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMembershipProtocolMapperArgs | GroupMembershipProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMembershipProtocolMapperState | undefined;
            inputs["addToAccessToken"] = state ? state.addToAccessToken : undefined;
            inputs["addToIdToken"] = state ? state.addToIdToken : undefined;
            inputs["addToUserinfo"] = state ? state.addToUserinfo : undefined;
            inputs["claimName"] = state ? state.claimName : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["fullPath"] = state ? state.fullPath : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as GroupMembershipProtocolMapperArgs | undefined;
            if ((!args || args.claimName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'claimName'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["addToAccessToken"] = args ? args.addToAccessToken : undefined;
            inputs["addToIdToken"] = args ? args.addToIdToken : undefined;
            inputs["addToUserinfo"] = args ? args.addToUserinfo : undefined;
            inputs["claimName"] = args ? args.claimName : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["fullPath"] = args ? args.fullPath : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GroupMembershipProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMembershipProtocolMapper resources.
 */
export interface GroupMembershipProtocolMapperState {
    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    /**
     * The name of the claim to insert into a token.
     */
    readonly claimName?: pulumi.Input<string>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
     */
    readonly fullPath?: pulumi.Input<boolean>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMembershipProtocolMapper resource.
 */
export interface GroupMembershipProtocolMapperArgs {
    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    /**
     * The name of the claim to insert into a token.
     */
    readonly claimName: pulumi.Input<string>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
     */
    readonly fullPath?: pulumi.Input<boolean>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    readonly realmId: pulumi.Input<string>;
}
