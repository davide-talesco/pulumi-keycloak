// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class ClientGroupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClientGroupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientGroupPolicyState, opts?: pulumi.CustomResourceOptions): ClientGroupPolicy {
        return new ClientGroupPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientGroupPolicy:ClientGroupPolicy';

    /**
     * Returns true if the given object is an instance of ClientGroupPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientGroupPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientGroupPolicy.__pulumiType;
    }

    public readonly decisionStrategy!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly groups!: pulumi.Output<outputs.openid.ClientGroupPolicyGroup[]>;
    public readonly groupsClaim!: pulumi.Output<string | undefined>;
    public readonly logic!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly resourceServerId!: pulumi.Output<string>;

    /**
     * Create a ClientGroupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientGroupPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientGroupPolicyArgs | ClientGroupPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientGroupPolicyState | undefined;
            resourceInputs["decisionStrategy"] = state ? state.decisionStrategy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["groupsClaim"] = state ? state.groupsClaim : undefined;
            resourceInputs["logic"] = state ? state.logic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["resourceServerId"] = state ? state.resourceServerId : undefined;
        } else {
            const args = argsOrState as ClientGroupPolicyArgs | undefined;
            if ((!args || args.decisionStrategy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'decisionStrategy'");
            }
            if ((!args || args.groups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groups'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.resourceServerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            resourceInputs["decisionStrategy"] = args ? args.decisionStrategy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["groupsClaim"] = args ? args.groupsClaim : undefined;
            resourceInputs["logic"] = args ? args.logic : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["resourceServerId"] = args ? args.resourceServerId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClientGroupPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientGroupPolicy resources.
 */
export interface ClientGroupPolicyState {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<inputs.openid.ClientGroupPolicyGroup>[]>;
    groupsClaim?: pulumi.Input<string>;
    logic?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId?: pulumi.Input<string>;
    resourceServerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientGroupPolicy resource.
 */
export interface ClientGroupPolicyArgs {
    decisionStrategy: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    groups: pulumi.Input<pulumi.Input<inputs.openid.ClientGroupPolicyGroup>[]>;
    groupsClaim?: pulumi.Input<string>;
    logic?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
    resourceServerId: pulumi.Input<string>;
}
