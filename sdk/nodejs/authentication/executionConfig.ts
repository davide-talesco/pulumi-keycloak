// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for managing an authentication execution's configuration. If a particular authentication execution supports additional
 * configuration (such as with the `identity-provider-redirector` execution), this can be managed with this resource.
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
 * const flow = new keycloak.authentication.Flow("flow", {
 *     realmId: realm.id,
 *     alias: "my-flow-alias",
 * });
 * const execution = new keycloak.authentication.Execution("execution", {
 *     realmId: realm.id,
 *     parentFlowAlias: flow.alias,
 *     authenticator: "identity-provider-redirector",
 * });
 * const config = new keycloak.authentication.ExecutionConfig("config", {
 *     realmId: realm.id,
 *     executionId: execution.id,
 *     alias: "my-config-alias",
 *     config: {
 *         defaultProvider: "my-config-default-idp",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Configurations can be imported using the format `{{realm}}/{{authenticationExecutionId}}/{{authenticationExecutionConfigId}}`. If the `authenticationExecutionId` is incorrect, the import will still be successful. A subsequent apply will change the `authenticationExecutionId` to the correct one, which causes the configuration to be replaced. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:authentication/executionConfig:ExecutionConfig config my-realm/be081463-ddbf-4b42-9eff-9c97886f24ff/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
 * ```
 */
export class ExecutionConfig extends pulumi.CustomResource {
    /**
     * Get an existing ExecutionConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExecutionConfigState, opts?: pulumi.CustomResourceOptions): ExecutionConfig {
        return new ExecutionConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:authentication/executionConfig:ExecutionConfig';

    /**
     * Returns true if the given object is an instance of ExecutionConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExecutionConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExecutionConfig.__pulumiType;
    }

    /**
     * The name of the configuration.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     */
    public readonly config!: pulumi.Output<{[key: string]: string}>;
    /**
     * The authentication execution this configuration is attached to.
     */
    public readonly executionId!: pulumi.Output<string>;
    /**
     * The realm the authentication execution exists in.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a ExecutionConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExecutionConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExecutionConfigArgs | ExecutionConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExecutionConfigState | undefined;
            inputs["alias"] = state ? state.alias : undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["executionId"] = state ? state.executionId : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as ExecutionConfigArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.executionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'executionId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["alias"] = args ? args.alias : undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["executionId"] = args ? args.executionId : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ExecutionConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExecutionConfig resources.
 */
export interface ExecutionConfigState {
    /**
     * The name of the configuration.
     */
    readonly alias?: pulumi.Input<string>;
    /**
     * The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     */
    readonly config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication execution this configuration is attached to.
     */
    readonly executionId?: pulumi.Input<string>;
    /**
     * The realm the authentication execution exists in.
     */
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExecutionConfig resource.
 */
export interface ExecutionConfigArgs {
    /**
     * The name of the configuration.
     */
    readonly alias: pulumi.Input<string>;
    /**
     * The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
     */
    readonly config: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The authentication execution this configuration is attached to.
     */
    readonly executionId: pulumi.Input<string>;
    /**
     * The realm the authentication execution exists in.
     */
    readonly realmId: pulumi.Input<string>;
}
