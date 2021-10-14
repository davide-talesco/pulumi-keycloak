// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    /// <summary>
    /// Allows for creating and managing `rsa-generated` Realm keystores within Keycloak.
    /// 
    /// A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
    /// 
    /// ## Import
    /// 
    /// Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:index/realmKeystoreRsaGenerated:RealmKeystoreRsaGenerated keystore_rsa_generated my-realm/my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/realmKeystoreRsaGenerated:RealmKeystoreRsaGenerated")]
    public partial class RealmKeystoreRsaGenerated : Pulumi.CustomResource
    {
        /// <summary>
        /// When `false`, key in not used for signing. Defaults to `true`.
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// Intended algorithm for the key. Defaults to `RS256`
        /// </summary>
        [Output("algorithm")]
        public Output<string?> Algorithm { get; private set; } = null!;

        /// <summary>
        /// When `false`, key is not accessible in this realm. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Size for the generated keys
        /// </summary>
        [Output("keySize")]
        public Output<int?> KeySize { get; private set; } = null!;

        /// <summary>
        /// Display name of provider when linked in admin console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Priority for the provider. Defaults to `0`
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The realm this keystore exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a RealmKeystoreRsaGenerated resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RealmKeystoreRsaGenerated(string name, RealmKeystoreRsaGeneratedArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/realmKeystoreRsaGenerated:RealmKeystoreRsaGenerated", name, args ?? new RealmKeystoreRsaGeneratedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RealmKeystoreRsaGenerated(string name, Input<string> id, RealmKeystoreRsaGeneratedState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/realmKeystoreRsaGenerated:RealmKeystoreRsaGenerated", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RealmKeystoreRsaGenerated resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RealmKeystoreRsaGenerated Get(string name, Input<string> id, RealmKeystoreRsaGeneratedState? state = null, CustomResourceOptions? options = null)
        {
            return new RealmKeystoreRsaGenerated(name, id, state, options);
        }
    }

    public sealed class RealmKeystoreRsaGeneratedArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `false`, key in not used for signing. Defaults to `true`.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Intended algorithm for the key. Defaults to `RS256`
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// When `false`, key is not accessible in this realm. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Size for the generated keys
        /// </summary>
        [Input("keySize")]
        public Input<int>? KeySize { get; set; }

        /// <summary>
        /// Display name of provider when linked in admin console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority for the provider. Defaults to `0`
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The realm this keystore exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public RealmKeystoreRsaGeneratedArgs()
        {
        }
    }

    public sealed class RealmKeystoreRsaGeneratedState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `false`, key in not used for signing. Defaults to `true`.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Intended algorithm for the key. Defaults to `RS256`
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// When `false`, key is not accessible in this realm. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Size for the generated keys
        /// </summary>
        [Input("keySize")]
        public Input<int>? KeySize { get; set; }

        /// <summary>
        /// Display name of provider when linked in admin console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority for the provider. Defaults to `0`
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The realm this keystore exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public RealmKeystoreRsaGeneratedState()
        {
        }
    }
}
