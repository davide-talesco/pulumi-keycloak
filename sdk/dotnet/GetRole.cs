// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public static class GetRole
    {
        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak role for
        /// usage with other resources, such as `keycloak.GroupRoles`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var offlineAccess = realm.Id.Apply(id =&gt; Keycloak.GetRole.InvokeAsync(new Keycloak.GetRoleArgs
        ///         {
        ///             RealmId = id,
        ///             Name = "offline_access",
        ///         }));
        ///         var @group = new Keycloak.Group("group", new Keycloak.GroupArgs
        ///         {
        ///             RealmId = realm.Id,
        ///         });
        ///         var groupRoles = new Keycloak.GroupRoles("groupRoles", new Keycloak.GroupRolesArgs
        ///         {
        ///             RealmId = realm.Id,
        ///             GroupId = @group.Id,
        ///             RoleIds = 
        ///             {
        ///                 offlineAccess.Apply(offlineAccess =&gt; offlineAccess.Id),
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRoleResult> InvokeAsync(GetRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleResult>("keycloak:index/getRole:getRole", args ?? new GetRoleArgs(), options.WithVersion());
    }


    public sealed class GetRoleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// When specified, this role is assumed to be a client role belonging to the client with the provided ID. The `id` attribute of a `keycloak_client` resource should be used here.
        /// </summary>
        [Input("clientId")]
        public string? ClientId { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The realm this role exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetRoleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRoleResult
    {
        public readonly ImmutableDictionary<string, object> Attributes;
        public readonly string? ClientId;
        public readonly ImmutableArray<string> CompositeRoles;
        /// <summary>
        /// (Computed) The description of the role.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string RealmId;

        [OutputConstructor]
        private GetRoleResult(
            ImmutableDictionary<string, object> attributes,

            string? clientId,

            ImmutableArray<string> compositeRoles,

            string description,

            string id,

            string name,

            string realmId)
        {
            Attributes = attributes;
            ClientId = clientId;
            CompositeRoles = compositeRoles;
            Description = description;
            Id = id;
            Name = name;
            RealmId = realmId;
        }
    }
}
