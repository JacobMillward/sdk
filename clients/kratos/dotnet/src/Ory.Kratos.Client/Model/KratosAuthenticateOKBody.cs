/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.3-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Kratos.Client.Client.OpenAPIDateConverter;

namespace Ory.Kratos.Client.Model
{
    /// <summary>
    /// AuthenticateOKBody authenticate o k body
    /// </summary>
    [DataContract(Name = "AuthenticateOKBody")]
    public partial class KratosAuthenticateOKBody : IEquatable<KratosAuthenticateOKBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosAuthenticateOKBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosAuthenticateOKBody()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosAuthenticateOKBody" /> class.
        /// </summary>
        /// <param name="identityToken">An opaque token used to authenticate a user after a successful login (required).</param>
        /// <param name="status">The status of the authentication (required).</param>
        public KratosAuthenticateOKBody(string identityToken = default(string), string status = default(string))
        {
            // to ensure "identityToken" is required (not null)
            this.IdentityToken = identityToken ?? throw new ArgumentNullException("identityToken is a required property for KratosAuthenticateOKBody and cannot be null");
            // to ensure "status" is required (not null)
            this.Status = status ?? throw new ArgumentNullException("status is a required property for KratosAuthenticateOKBody and cannot be null");
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// An opaque token used to authenticate a user after a successful login
        /// </summary>
        /// <value>An opaque token used to authenticate a user after a successful login</value>
        [DataMember(Name = "IdentityToken", IsRequired = true, EmitDefaultValue = false)]
        public string IdentityToken { get; set; }

        /// <summary>
        /// The status of the authentication
        /// </summary>
        /// <value>The status of the authentication</value>
        [DataMember(Name = "Status", IsRequired = true, EmitDefaultValue = false)]
        public string Status { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class KratosAuthenticateOKBody {\n");
            sb.Append("  IdentityToken: ").Append(IdentityToken).Append("\n");
            sb.Append("  Status: ").Append(Status).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as KratosAuthenticateOKBody);
        }

        /// <summary>
        /// Returns true if KratosAuthenticateOKBody instances are equal
        /// </summary>
        /// <param name="input">Instance of KratosAuthenticateOKBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(KratosAuthenticateOKBody input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.IdentityToken == input.IdentityToken ||
                    (this.IdentityToken != null &&
                    this.IdentityToken.Equals(input.IdentityToken))
                ) && 
                (
                    this.Status == input.Status ||
                    (this.Status != null &&
                    this.Status.Equals(input.Status))
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.IdentityToken != null)
                    hashCode = hashCode * 59 + this.IdentityToken.GetHashCode();
                if (this.Status != null)
                    hashCode = hashCode * 59 + this.Status.GetHashCode();
                if (this.AdditionalProperties != null)
                    hashCode = hashCode * 59 + this.AdditionalProperties.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
