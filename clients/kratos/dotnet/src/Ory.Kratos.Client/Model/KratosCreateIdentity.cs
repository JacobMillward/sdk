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
    /// KratosCreateIdentity
    /// </summary>
    [DataContract(Name = "createIdentity")]
    public partial class KratosCreateIdentity : IEquatable<KratosCreateIdentity>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosCreateIdentity" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosCreateIdentity()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosCreateIdentity" /> class.
        /// </summary>
        /// <param name="schemaId">SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits. (required).</param>
        /// <param name="traits">Traits represent an identity&#39;s traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in &#x60;schema_url&#x60;. (required).</param>
        public KratosCreateIdentity(string schemaId = default(string), Object traits = default(Object))
        {
            // to ensure "schemaId" is required (not null)
            this.SchemaId = schemaId ?? throw new ArgumentNullException("schemaId is a required property for KratosCreateIdentity and cannot be null");
            // to ensure "traits" is required (not null)
            this.Traits = traits ?? throw new ArgumentNullException("traits is a required property for KratosCreateIdentity and cannot be null");
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits.
        /// </summary>
        /// <value>SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits.</value>
        [DataMember(Name = "schema_id", IsRequired = true, EmitDefaultValue = false)]
        public string SchemaId { get; set; }

        /// <summary>
        /// Traits represent an identity&#39;s traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in &#x60;schema_url&#x60;.
        /// </summary>
        /// <value>Traits represent an identity&#39;s traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in &#x60;schema_url&#x60;.</value>
        [DataMember(Name = "traits", IsRequired = true, EmitDefaultValue = false)]
        public Object Traits { get; set; }

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
            sb.Append("class KratosCreateIdentity {\n");
            sb.Append("  SchemaId: ").Append(SchemaId).Append("\n");
            sb.Append("  Traits: ").Append(Traits).Append("\n");
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
            return this.Equals(input as KratosCreateIdentity);
        }

        /// <summary>
        /// Returns true if KratosCreateIdentity instances are equal
        /// </summary>
        /// <param name="input">Instance of KratosCreateIdentity to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(KratosCreateIdentity input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.SchemaId == input.SchemaId ||
                    (this.SchemaId != null &&
                    this.SchemaId.Equals(input.SchemaId))
                ) && 
                (
                    this.Traits == input.Traits ||
                    (this.Traits != null &&
                    this.Traits.Equals(input.Traits))
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
                if (this.SchemaId != null)
                    hashCode = hashCode * 59 + this.SchemaId.GetHashCode();
                if (this.Traits != null)
                    hashCode = hashCode * 59 + this.Traits.GetHashCode();
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
