// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ClientAggregatePolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClientAggregatePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientAggregatePolicyState, opts?: pulumi.CustomResourceOptions): ClientAggregatePolicy {
        return new ClientAggregatePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientAggregatePolicy:ClientAggregatePolicy';

    /**
     * Returns true if the given object is an instance of ClientAggregatePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientAggregatePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientAggregatePolicy.__pulumiType;
    }

    public readonly decisionStrategy!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly logic!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly policies!: pulumi.Output<string[]>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly resourceServerId!: pulumi.Output<string>;

    /**
     * Create a ClientAggregatePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientAggregatePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientAggregatePolicyArgs | ClientAggregatePolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientAggregatePolicyState | undefined;
            inputs["decisionStrategy"] = state ? state.decisionStrategy : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["logic"] = state ? state.logic : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["resourceServerId"] = state ? state.resourceServerId : undefined;
        } else {
            const args = argsOrState as ClientAggregatePolicyArgs | undefined;
            if ((!args || args.decisionStrategy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'decisionStrategy'");
            }
            if ((!args || args.policies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policies'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.resourceServerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            inputs["decisionStrategy"] = args ? args.decisionStrategy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["logic"] = args ? args.logic : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["resourceServerId"] = args ? args.resourceServerId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClientAggregatePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientAggregatePolicy resources.
 */
export interface ClientAggregatePolicyState {
    readonly decisionStrategy?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly logic?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    readonly realmId?: pulumi.Input<string>;
    readonly resourceServerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientAggregatePolicy resource.
 */
export interface ClientAggregatePolicyArgs {
    readonly decisionStrategy: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly logic?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly policies: pulumi.Input<pulumi.Input<string>[]>;
    readonly realmId: pulumi.Input<string>;
    readonly resourceServerId: pulumi.Input<string>;
}
