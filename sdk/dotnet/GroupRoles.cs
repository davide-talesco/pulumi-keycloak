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
    /// ## Example Usage
    /// ### Exhaustive Roles)
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
    ///         var realmRole = new Keycloak.Role("realmRole", new Keycloak.RoleArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             Description = "My Realm Role",
    ///         });
    ///         var client = new Keycloak.OpenId.Client("client", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = "client",
    ///             Enabled = true,
    ///             AccessType = "BEARER-ONLY",
    ///         });
    ///         var clientRole = new Keycloak.Role("clientRole", new Keycloak.RoleArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = keycloak_client.Client.Id,
    ///             Description = "My Client Role",
    ///         });
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
    ///                 realmRole.Id,
    ///                 clientRole.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Non Exhaustive Roles)
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
    ///         var realmRole = new Keycloak.Role("realmRole", new Keycloak.RoleArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             Description = "My Realm Role",
    ///         });
    ///         var client = new Keycloak.OpenId.Client("client", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = "client",
    ///             Enabled = true,
    ///             AccessType = "BEARER-ONLY",
    ///         });
    ///         var clientRole = new Keycloak.Role("clientRole", new Keycloak.RoleArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = keycloak_client.Client.Id,
    ///             Description = "My Client Role",
    ///         });
    ///         var @group = new Keycloak.Group("group", new Keycloak.GroupArgs
    ///         {
    ///             RealmId = realm.Id,
    ///         });
    ///         var groupRoleAssociation1 = new Keycloak.GroupRoles("groupRoleAssociation1", new Keycloak.GroupRolesArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             GroupId = @group.Id,
    ///             Exhaustive = false,
    ///             RoleIds = 
    ///             {
    ///                 realmRole.Id,
    ///             },
    ///         });
    ///         var groupRoleAssociation2 = new Keycloak.GroupRoles("groupRoleAssociation2", new Keycloak.GroupRolesArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             GroupId = @group.Id,
    ///             Exhaustive = false,
    ///             RoleIds = 
    ///             {
    ///                 clientRole.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the format `{{realm_id}}/{{group_id}}`, where `group_id` is the unique ID that Keycloak assigns to the group upon creation. This value can be found in the URI when editing this group in the GUI, and is typically a GUID. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:index/groupRoles:GroupRoles group_roles my-realm/18cc6b87-2ce7-4e59-bdc8-b9d49ec98a94
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/groupRoles:GroupRoles")]
    public partial class GroupRoles : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the group will be removed. Defaults to `true`.
        /// </summary>
        [Output("exhaustive")]
        public Output<bool?> Exhaustive { get; private set; } = null!;

        /// <summary>
        /// The ID of the group this resource should manage roles for.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// A list of role IDs to map to the group.
        /// </summary>
        [Output("roleIds")]
        public Output<ImmutableArray<string>> RoleIds { get; private set; } = null!;


        /// <summary>
        /// Create a GroupRoles resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupRoles(string name, GroupRolesArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/groupRoles:GroupRoles", name, args ?? new GroupRolesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupRoles(string name, Input<string> id, GroupRolesState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/groupRoles:GroupRoles", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupRoles resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupRoles Get(string name, Input<string> id, GroupRolesState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupRoles(name, id, state, options);
        }
    }

    public sealed class GroupRolesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the group will be removed. Defaults to `true`.
        /// </summary>
        [Input("exhaustive")]
        public Input<bool>? Exhaustive { get; set; }

        /// <summary>
        /// The ID of the group this resource should manage roles for.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("roleIds", required: true)]
        private InputList<string>? _roleIds;

        /// <summary>
        /// A list of role IDs to map to the group.
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        public GroupRolesArgs()
        {
        }
    }

    public sealed class GroupRolesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the group will be removed. Defaults to `true`.
        /// </summary>
        [Input("exhaustive")]
        public Input<bool>? Exhaustive { get; set; }

        /// <summary>
        /// The ID of the group this resource should manage roles for.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("roleIds")]
        private InputList<string>? _roleIds;

        /// <summary>
        /// A list of role IDs to map to the group.
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        public GroupRolesState()
        {
        }
    }
}
