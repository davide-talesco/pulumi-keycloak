// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ClientJsPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClientJsPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientJsPolicyState, opts?: pulumi.CustomResourceOptions): ClientJsPolicy {
        return new ClientJsPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientJsPolicy:ClientJsPolicy';

    /**
     * Returns true if the given object is an instance of ClientJsPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientJsPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientJsPolicy.__pulumiType;
    }

    public readonly code!: pulumi.Output<string>;
    public readonly decisionStrategy!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly logic!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly resourceServerId!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ClientJsPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientJsPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientJsPolicyArgs | ClientJsPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientJsPolicyState | undefined;
            inputs["code"] = state ? state.code : undefined;
            inputs["decisionStrategy"] = state ? state.decisionStrategy : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["logic"] = state ? state.logic : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["resourceServerId"] = state ? state.resourceServerId : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ClientJsPolicyArgs | undefined;
            if ((!args || args.code === undefined) && !opts.urn) {
                throw new Error("Missing required property 'code'");
            }
            if ((!args || args.decisionStrategy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'decisionStrategy'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.resourceServerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            inputs["code"] = args ? args.code : undefined;
            inputs["decisionStrategy"] = args ? args.decisionStrategy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["logic"] = args ? args.logic : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["resourceServerId"] = args ? args.resourceServerId : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClientJsPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientJsPolicy resources.
 */
export interface ClientJsPolicyState {
    code?: pulumi.Input<string>;
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    logic?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId?: pulumi.Input<string>;
    resourceServerId?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientJsPolicy resource.
 */
export interface ClientJsPolicyArgs {
    code: pulumi.Input<string>;
    decisionStrategy: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    logic?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
    resourceServerId: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
