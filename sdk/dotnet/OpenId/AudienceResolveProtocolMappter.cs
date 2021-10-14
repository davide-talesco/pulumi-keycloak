// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// Allows for creating the "Audience Resolve" OIDC protocol mapper within Keycloak.
    /// 
    /// This protocol mapper is useful to avoid manual management of audiences, instead relying on the presence of client roles
    /// to imply which audiences are appropriate for the token. See the
    /// [Keycloak docs](https://www.keycloak.org/docs/latest/server_admin/#_audience_resolve) for more details.
    /// 
    /// ## Example Usage
    /// ### Client)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             Realm = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var openidClient = new Keycloak.OpenId.Client("openidClient", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = "client",
    ///             Enabled = true,
    ///             AccessType = "CONFIDENTIAL",
    ///             ValidRedirectUris = 
    ///             {
    ///                 "http://localhost:8080/openid-callback",
    ///             },
    ///         });
    ///         var audienceMapper = new Keycloak.OpenId.AudienceResolveProtocolMappter("audienceMapper", new Keycloak.OpenId.AudienceResolveProtocolMappterArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = openidClient.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Client Scope)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             Realm = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var clientScope = new Keycloak.OpenId.ClientScope("clientScope", new Keycloak.OpenId.ClientScopeArgs
    ///         {
    ///             RealmId = realm.Id,
    ///         });
    ///         var audienceMapper = new Keycloak.OpenId.AudienceProtocolMapper("audienceMapper", new Keycloak.OpenId.AudienceProtocolMapperArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientScopeId = clientScope.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter")]
    public partial class AudienceResolveProtocolMappter : Pulumi.CustomResource
    {
        /// <summary>
        /// The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        /// <summary>
        /// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a AudienceResolveProtocolMappter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AudienceResolveProtocolMappter(string name, AudienceResolveProtocolMappterArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter", name, args ?? new AudienceResolveProtocolMappterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AudienceResolveProtocolMappter(string name, Input<string> id, AudienceResolveProtocolMappterState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AudienceResolveProtocolMappter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AudienceResolveProtocolMappter Get(string name, Input<string> id, AudienceResolveProtocolMappterState? state = null, CustomResourceOptions? options = null)
        {
            return new AudienceResolveProtocolMappter(name, id, state, options);
        }
    }

    public sealed class AudienceResolveProtocolMappterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public AudienceResolveProtocolMappterArgs()
        {
        }
    }

    public sealed class AudienceResolveProtocolMappterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public AudienceResolveProtocolMappterState()
        {
        }
    }
}
