// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public partial class HardcodedAttributeIdentityProviderMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// OIDC Claim
        /// </summary>
        [Output("attributeName")]
        public Output<string?> AttributeName { get; private set; } = null!;

        /// <summary>
        /// User Attribute
        /// </summary>
        [Output("attributeValue")]
        public Output<string?> AttributeValue { get; private set; } = null!;

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Output("identityProviderAlias")]
        public Output<string> IdentityProviderAlias { get; private set; } = null!;

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Realm Name
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// Is Attribute Related To a User Session
        /// </summary>
        [Output("userSession")]
        public Output<bool> UserSession { get; private set; } = null!;


        /// <summary>
        /// Create a HardcodedAttributeIdentityProviderMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HardcodedAttributeIdentityProviderMapper(string name, HardcodedAttributeIdentityProviderMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private HardcodedAttributeIdentityProviderMapper(string name, Input<string> id, HardcodedAttributeIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HardcodedAttributeIdentityProviderMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HardcodedAttributeIdentityProviderMapper Get(string name, Input<string> id, HardcodedAttributeIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new HardcodedAttributeIdentityProviderMapper(name, id, state, options);
        }
    }

    public sealed class HardcodedAttributeIdentityProviderMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// OIDC Claim
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// User Attribute
        /// </summary>
        [Input("attributeValue")]
        public Input<string>? AttributeValue { get; set; }

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Input("identityProviderAlias", required: true)]
        public Input<string> IdentityProviderAlias { get; set; } = null!;

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// Is Attribute Related To a User Session
        /// </summary>
        [Input("userSession", required: true)]
        public Input<bool> UserSession { get; set; } = null!;

        public HardcodedAttributeIdentityProviderMapperArgs()
        {
        }
    }

    public sealed class HardcodedAttributeIdentityProviderMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// OIDC Claim
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// User Attribute
        /// </summary>
        [Input("attributeValue")]
        public Input<string>? AttributeValue { get; set; }

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Input("identityProviderAlias")]
        public Input<string>? IdentityProviderAlias { get; set; }

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Is Attribute Related To a User Session
        /// </summary>
        [Input("userSession")]
        public Input<bool>? UserSession { get; set; }

        public HardcodedAttributeIdentityProviderMapperState()
        {
        }
    }
}
