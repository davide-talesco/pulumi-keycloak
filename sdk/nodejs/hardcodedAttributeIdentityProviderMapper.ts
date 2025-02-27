// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class HardcodedAttributeIdentityProviderMapper extends pulumi.CustomResource {
    /**
     * Get an existing HardcodedAttributeIdentityProviderMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HardcodedAttributeIdentityProviderMapperState, opts?: pulumi.CustomResourceOptions): HardcodedAttributeIdentityProviderMapper {
        return new HardcodedAttributeIdentityProviderMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper';

    /**
     * Returns true if the given object is an instance of HardcodedAttributeIdentityProviderMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HardcodedAttributeIdentityProviderMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HardcodedAttributeIdentityProviderMapper.__pulumiType;
    }

    /**
     * OIDC Claim
     */
    public readonly attributeName!: pulumi.Output<string | undefined>;
    /**
     * User Attribute
     */
    public readonly attributeValue!: pulumi.Output<string | undefined>;
    public readonly extraConfig!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * IDP Alias
     */
    public readonly identityProviderAlias!: pulumi.Output<string>;
    /**
     * IDP Mapper Name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Realm Name
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * Is Attribute Related To a User Session
     */
    public readonly userSession!: pulumi.Output<boolean>;

    /**
     * Create a HardcodedAttributeIdentityProviderMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HardcodedAttributeIdentityProviderMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HardcodedAttributeIdentityProviderMapperArgs | HardcodedAttributeIdentityProviderMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HardcodedAttributeIdentityProviderMapperState | undefined;
            resourceInputs["attributeName"] = state ? state.attributeName : undefined;
            resourceInputs["attributeValue"] = state ? state.attributeValue : undefined;
            resourceInputs["extraConfig"] = state ? state.extraConfig : undefined;
            resourceInputs["identityProviderAlias"] = state ? state.identityProviderAlias : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["userSession"] = state ? state.userSession : undefined;
        } else {
            const args = argsOrState as HardcodedAttributeIdentityProviderMapperArgs | undefined;
            if ((!args || args.identityProviderAlias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityProviderAlias'");
            }
            if ((!args || args.realm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realm'");
            }
            if ((!args || args.userSession === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userSession'");
            }
            resourceInputs["attributeName"] = args ? args.attributeName : undefined;
            resourceInputs["attributeValue"] = args ? args.attributeValue : undefined;
            resourceInputs["extraConfig"] = args ? args.extraConfig : undefined;
            resourceInputs["identityProviderAlias"] = args ? args.identityProviderAlias : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["userSession"] = args ? args.userSession : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HardcodedAttributeIdentityProviderMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HardcodedAttributeIdentityProviderMapper resources.
 */
export interface HardcodedAttributeIdentityProviderMapperState {
    /**
     * OIDC Claim
     */
    attributeName?: pulumi.Input<string>;
    /**
     * User Attribute
     */
    attributeValue?: pulumi.Input<string>;
    extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * IDP Alias
     */
    identityProviderAlias?: pulumi.Input<string>;
    /**
     * IDP Mapper Name
     */
    name?: pulumi.Input<string>;
    /**
     * Realm Name
     */
    realm?: pulumi.Input<string>;
    /**
     * Is Attribute Related To a User Session
     */
    userSession?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a HardcodedAttributeIdentityProviderMapper resource.
 */
export interface HardcodedAttributeIdentityProviderMapperArgs {
    /**
     * OIDC Claim
     */
    attributeName?: pulumi.Input<string>;
    /**
     * User Attribute
     */
    attributeValue?: pulumi.Input<string>;
    extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * IDP Alias
     */
    identityProviderAlias: pulumi.Input<string>;
    /**
     * IDP Mapper Name
     */
    name?: pulumi.Input<string>;
    /**
     * Realm Name
     */
    realm: pulumi.Input<string>;
    /**
     * Is Attribute Related To a User Session
     */
    userSession: pulumi.Input<boolean>;
}
