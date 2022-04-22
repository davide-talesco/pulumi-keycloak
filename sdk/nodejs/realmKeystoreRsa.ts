// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing `rsa` Realm keystores within Keycloak.
 *
 * A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
 *
 * ## Import
 *
 * Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:index/realmKeystoreRsa:RealmKeystoreRsa keystore_rsa my-realm/my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
 * ```
 */
export class RealmKeystoreRsa extends pulumi.CustomResource {
    /**
     * Get an existing RealmKeystoreRsa resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RealmKeystoreRsaState, opts?: pulumi.CustomResourceOptions): RealmKeystoreRsa {
        return new RealmKeystoreRsa(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/realmKeystoreRsa:RealmKeystoreRsa';

    /**
     * Returns true if the given object is an instance of RealmKeystoreRsa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RealmKeystoreRsa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RealmKeystoreRsa.__pulumiType;
    }

    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     */
    public readonly algorithm!: pulumi.Output<string | undefined>;
    /**
     * X509 Certificate encoded in PEM format.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Display name of provider when linked in admin console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Priority for the provider. Defaults to `0`
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Private RSA Key encoded in PEM format.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * The realm this keystore exists in.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a RealmKeystoreRsa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealmKeystoreRsaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RealmKeystoreRsaArgs | RealmKeystoreRsaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RealmKeystoreRsaState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as RealmKeystoreRsaArgs | undefined;
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RealmKeystoreRsa.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RealmKeystoreRsa resources.
 */
export interface RealmKeystoreRsaState {
    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     */
    algorithm?: pulumi.Input<string>;
    /**
     * X509 Certificate encoded in PEM format.
     */
    certificate?: pulumi.Input<string>;
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Display name of provider when linked in admin console.
     */
    name?: pulumi.Input<string>;
    /**
     * Priority for the provider. Defaults to `0`
     */
    priority?: pulumi.Input<number>;
    /**
     * Private RSA Key encoded in PEM format.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The realm this keystore exists in.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RealmKeystoreRsa resource.
 */
export interface RealmKeystoreRsaArgs {
    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     */
    algorithm?: pulumi.Input<string>;
    /**
     * X509 Certificate encoded in PEM format.
     */
    certificate: pulumi.Input<string>;
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Display name of provider when linked in admin console.
     */
    name?: pulumi.Input<string>;
    /**
     * Priority for the provider. Defaults to `0`
     */
    priority?: pulumi.Input<number>;
    /**
     * Private RSA Key encoded in PEM format.
     */
    privateKey: pulumi.Input<string>;
    /**
     * The realm this keystore exists in.
     */
    realmId: pulumi.Input<string>;
}
