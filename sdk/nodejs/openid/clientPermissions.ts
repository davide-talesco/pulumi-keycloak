// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class ClientPermissions extends pulumi.CustomResource {
    /**
     * Get an existing ClientPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientPermissionsState, opts?: pulumi.CustomResourceOptions): ClientPermissions {
        return new ClientPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientPermissions:ClientPermissions';

    /**
     * Returns true if the given object is an instance of ClientPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientPermissions.__pulumiType;
    }

    /**
     * Resource server id representing the realm management client on which this permission is managed
     */
    public /*out*/ readonly authorizationResourceServerId!: pulumi.Output<string>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly configureScope!: pulumi.Output<outputs.openid.ClientPermissionsConfigureScope | undefined>;
    public /*out*/ readonly enabled!: pulumi.Output<boolean>;
    public readonly manageScope!: pulumi.Output<outputs.openid.ClientPermissionsManageScope | undefined>;
    public readonly mapRolesClientScopeScope!: pulumi.Output<outputs.openid.ClientPermissionsMapRolesClientScopeScope | undefined>;
    public readonly mapRolesCompositeScope!: pulumi.Output<outputs.openid.ClientPermissionsMapRolesCompositeScope | undefined>;
    public readonly mapRolesScope!: pulumi.Output<outputs.openid.ClientPermissionsMapRolesScope | undefined>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly tokenExchangeScope!: pulumi.Output<outputs.openid.ClientPermissionsTokenExchangeScope | undefined>;
    public readonly viewScope!: pulumi.Output<outputs.openid.ClientPermissionsViewScope | undefined>;

    /**
     * Create a ClientPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientPermissionsArgs | ClientPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientPermissionsState | undefined;
            resourceInputs["authorizationResourceServerId"] = state ? state.authorizationResourceServerId : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["configureScope"] = state ? state.configureScope : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["manageScope"] = state ? state.manageScope : undefined;
            resourceInputs["mapRolesClientScopeScope"] = state ? state.mapRolesClientScopeScope : undefined;
            resourceInputs["mapRolesCompositeScope"] = state ? state.mapRolesCompositeScope : undefined;
            resourceInputs["mapRolesScope"] = state ? state.mapRolesScope : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["tokenExchangeScope"] = state ? state.tokenExchangeScope : undefined;
            resourceInputs["viewScope"] = state ? state.viewScope : undefined;
        } else {
            const args = argsOrState as ClientPermissionsArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["configureScope"] = args ? args.configureScope : undefined;
            resourceInputs["manageScope"] = args ? args.manageScope : undefined;
            resourceInputs["mapRolesClientScopeScope"] = args ? args.mapRolesClientScopeScope : undefined;
            resourceInputs["mapRolesCompositeScope"] = args ? args.mapRolesCompositeScope : undefined;
            resourceInputs["mapRolesScope"] = args ? args.mapRolesScope : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["tokenExchangeScope"] = args ? args.tokenExchangeScope : undefined;
            resourceInputs["viewScope"] = args ? args.viewScope : undefined;
            resourceInputs["authorizationResourceServerId"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClientPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientPermissions resources.
 */
export interface ClientPermissionsState {
    /**
     * Resource server id representing the realm management client on which this permission is managed
     */
    authorizationResourceServerId?: pulumi.Input<string>;
    clientId?: pulumi.Input<string>;
    configureScope?: pulumi.Input<inputs.openid.ClientPermissionsConfigureScope>;
    enabled?: pulumi.Input<boolean>;
    manageScope?: pulumi.Input<inputs.openid.ClientPermissionsManageScope>;
    mapRolesClientScopeScope?: pulumi.Input<inputs.openid.ClientPermissionsMapRolesClientScopeScope>;
    mapRolesCompositeScope?: pulumi.Input<inputs.openid.ClientPermissionsMapRolesCompositeScope>;
    mapRolesScope?: pulumi.Input<inputs.openid.ClientPermissionsMapRolesScope>;
    realmId?: pulumi.Input<string>;
    tokenExchangeScope?: pulumi.Input<inputs.openid.ClientPermissionsTokenExchangeScope>;
    viewScope?: pulumi.Input<inputs.openid.ClientPermissionsViewScope>;
}

/**
 * The set of arguments for constructing a ClientPermissions resource.
 */
export interface ClientPermissionsArgs {
    clientId: pulumi.Input<string>;
    configureScope?: pulumi.Input<inputs.openid.ClientPermissionsConfigureScope>;
    manageScope?: pulumi.Input<inputs.openid.ClientPermissionsManageScope>;
    mapRolesClientScopeScope?: pulumi.Input<inputs.openid.ClientPermissionsMapRolesClientScopeScope>;
    mapRolesCompositeScope?: pulumi.Input<inputs.openid.ClientPermissionsMapRolesCompositeScope>;
    mapRolesScope?: pulumi.Input<inputs.openid.ClientPermissionsMapRolesScope>;
    realmId: pulumi.Input<string>;
    tokenExchangeScope?: pulumi.Input<inputs.openid.ClientPermissionsTokenExchangeScope>;
    viewScope?: pulumi.Input<inputs.openid.ClientPermissionsViewScope>;
}
