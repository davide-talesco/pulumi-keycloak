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
    public static class GetAuthenticationExecution
    {
        /// <summary>
        /// This data source can be used to fetch the ID of an authentication execution within Keycloak.
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
        ///         var browserAuthCookie = realm.Id.Apply(id =&gt; Keycloak.GetAuthenticationExecution.InvokeAsync(new Keycloak.GetAuthenticationExecutionArgs
        ///         {
        ///             RealmId = id,
        ///             ParentFlowAlias = "browser",
        ///             ProviderId = "auth-cookie",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthenticationExecutionResult> InvokeAsync(GetAuthenticationExecutionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthenticationExecutionResult>("keycloak:index/getAuthenticationExecution:getAuthenticationExecution", args ?? new GetAuthenticationExecutionArgs(), options.WithVersion());

        /// <summary>
        /// This data source can be used to fetch the ID of an authentication execution within Keycloak.
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
        ///         var browserAuthCookie = realm.Id.Apply(id =&gt; Keycloak.GetAuthenticationExecution.InvokeAsync(new Keycloak.GetAuthenticationExecutionArgs
        ///         {
        ///             RealmId = id,
        ///             ParentFlowAlias = "browser",
        ///             ProviderId = "auth-cookie",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAuthenticationExecutionResult> Invoke(GetAuthenticationExecutionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAuthenticationExecutionResult>("keycloak:index/getAuthenticationExecution:getAuthenticationExecution", args ?? new GetAuthenticationExecutionInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAuthenticationExecutionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alias of the flow this execution is attached to.
        /// </summary>
        [Input("parentFlowAlias", required: true)]
        public string ParentFlowAlias { get; set; } = null!;

        /// <summary>
        /// The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools. This was previously known as the "authenticator".
        /// </summary>
        [Input("providerId", required: true)]
        public string ProviderId { get; set; } = null!;

        /// <summary>
        /// The realm the authentication execution exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetAuthenticationExecutionArgs()
        {
        }
    }

    public sealed class GetAuthenticationExecutionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alias of the flow this execution is attached to.
        /// </summary>
        [Input("parentFlowAlias", required: true)]
        public Input<string> ParentFlowAlias { get; set; } = null!;

        /// <summary>
        /// The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools. This was previously known as the "authenticator".
        /// </summary>
        [Input("providerId", required: true)]
        public Input<string> ProviderId { get; set; } = null!;

        /// <summary>
        /// The realm the authentication execution exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GetAuthenticationExecutionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthenticationExecutionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ParentFlowAlias;
        public readonly string ProviderId;
        public readonly string RealmId;

        [OutputConstructor]
        private GetAuthenticationExecutionResult(
            string id,

            string parentFlowAlias,

            string providerId,

            string realmId)
        {
            Id = id;
            ParentFlowAlias = parentFlowAlias;
            ProviderId = providerId;
            RealmId = realmId;
        }
    }
}
