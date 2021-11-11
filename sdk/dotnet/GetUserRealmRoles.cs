// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Keycloak
{
    public static class GetUserRealmRoles
    {
        /// <summary>
        /// This data source can be used to fetch the realm roles of a user within Keycloak.
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
        ///         var masterRealm = Output.Create(Keycloak.GetRealm.InvokeAsync(new Keycloak.GetRealmArgs
        ///         {
        ///             Realm = "master",
        ///         }));
        ///         var defaultAdminUser = masterRealm.Apply(masterRealm =&gt; Output.Create(Keycloak.GetUser.InvokeAsync(new Keycloak.GetUserArgs
        ///         {
        ///             RealmId = masterRealm.Id,
        ///             Username = "keycloak",
        ///         })));
        ///         var userRealmRoles = Output.Tuple(masterRealm, defaultAdminUser).Apply(values =&gt;
        ///         {
        ///             var masterRealm = values.Item1;
        ///             var defaultAdminUser = values.Item2;
        ///             return Output.Create(Keycloak.GetUserRealmRoles.InvokeAsync(new Keycloak.GetUserRealmRolesArgs
        ///             {
        ///                 RealmId = masterRealm.Id,
        ///                 UserId = defaultAdminUser.Id,
        ///             }));
        ///         });
        ///         this.KeycloakUserRoleNames = userRealmRoles.Apply(userRealmRoles =&gt; userRealmRoles.RoleNames);
        ///     }
        /// 
        ///     [Output("keycloakUserRoleNames")]
        ///     public Output&lt;string&gt; KeycloakUserRoleNames { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserRealmRolesResult> InvokeAsync(GetUserRealmRolesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserRealmRolesResult>("keycloak:index/getUserRealmRoles:getUserRealmRoles", args ?? new GetUserRealmRolesArgs(), options.WithVersion());

        /// <summary>
        /// This data source can be used to fetch the realm roles of a user within Keycloak.
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
        ///         var masterRealm = Output.Create(Keycloak.GetRealm.InvokeAsync(new Keycloak.GetRealmArgs
        ///         {
        ///             Realm = "master",
        ///         }));
        ///         var defaultAdminUser = masterRealm.Apply(masterRealm =&gt; Output.Create(Keycloak.GetUser.InvokeAsync(new Keycloak.GetUserArgs
        ///         {
        ///             RealmId = masterRealm.Id,
        ///             Username = "keycloak",
        ///         })));
        ///         var userRealmRoles = Output.Tuple(masterRealm, defaultAdminUser).Apply(values =&gt;
        ///         {
        ///             var masterRealm = values.Item1;
        ///             var defaultAdminUser = values.Item2;
        ///             return Output.Create(Keycloak.GetUserRealmRoles.InvokeAsync(new Keycloak.GetUserRealmRolesArgs
        ///             {
        ///                 RealmId = masterRealm.Id,
        ///                 UserId = defaultAdminUser.Id,
        ///             }));
        ///         });
        ///         this.KeycloakUserRoleNames = userRealmRoles.Apply(userRealmRoles =&gt; userRealmRoles.RoleNames);
        ///     }
        /// 
        ///     [Output("keycloakUserRoleNames")]
        ///     public Output&lt;string&gt; KeycloakUserRoleNames { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserRealmRolesResult> Invoke(GetUserRealmRolesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUserRealmRolesResult>("keycloak:index/getUserRealmRoles:getUserRealmRoles", args ?? new GetUserRealmRolesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetUserRealmRolesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The realm this user belongs to.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        /// <summary>
        /// The ID of the user to query realm roles for.
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetUserRealmRolesArgs()
        {
        }
    }

    public sealed class GetUserRealmRolesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The realm this user belongs to.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The ID of the user to query realm roles for.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetUserRealmRolesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserRealmRolesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RealmId;
        /// <summary>
        /// (Computed) A list of realm roles that belong to this user.
        /// </summary>
        public readonly ImmutableArray<string> RoleNames;
        public readonly string UserId;

        [OutputConstructor]
        private GetUserRealmRolesResult(
            string id,

            string realmId,

            ImmutableArray<string> roleNames,

            string userId)
        {
            Id = id;
            RealmId = realmId;
            RoleNames = roleNames;
            UserId = userId;
        }
    }
}
