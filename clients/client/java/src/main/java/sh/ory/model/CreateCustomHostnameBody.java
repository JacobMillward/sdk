/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.181
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * CreateCustomHostnameBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-05-24T14:00:22.693216924Z[Etc/UTC]")
public class CreateCustomHostnameBody {
  public static final String SERIALIZED_NAME_COOKIE_DOMAIN = "cookie_domain";
  @SerializedName(SERIALIZED_NAME_COOKIE_DOMAIN)
  private String cookieDomain;

  public static final String SERIALIZED_NAME_CORS_ALLOWED_ORIGINS = "cors_allowed_origins";
  @SerializedName(SERIALIZED_NAME_CORS_ALLOWED_ORIGINS)
  private List<String> corsAllowedOrigins = null;

  public static final String SERIALIZED_NAME_CORS_ENABLED = "cors_enabled";
  @SerializedName(SERIALIZED_NAME_CORS_ENABLED)
  private Boolean corsEnabled;

  public static final String SERIALIZED_NAME_HOSTNAME = "hostname";
  @SerializedName(SERIALIZED_NAME_HOSTNAME)
  private String hostname;

  public CreateCustomHostnameBody() { 
  }

  public CreateCustomHostnameBody cookieDomain(String cookieDomain) {
    
    this.cookieDomain = cookieDomain;
    return this;
  }

   /**
   * The domain where cookies will be set. Has to be a parent domain of the custom hostname to work.
   * @return cookieDomain
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The domain where cookies will be set. Has to be a parent domain of the custom hostname to work.")

  public String getCookieDomain() {
    return cookieDomain;
  }


  public void setCookieDomain(String cookieDomain) {
    this.cookieDomain = cookieDomain;
  }


  public CreateCustomHostnameBody corsAllowedOrigins(List<String> corsAllowedOrigins) {
    
    this.corsAllowedOrigins = corsAllowedOrigins;
    return this;
  }

  public CreateCustomHostnameBody addCorsAllowedOriginsItem(String corsAllowedOriginsItem) {
    if (this.corsAllowedOrigins == null) {
      this.corsAllowedOrigins = new ArrayList<>();
    }
    this.corsAllowedOrigins.add(corsAllowedOriginsItem);
    return this;
  }

   /**
   * CORS Allowed origins for the custom hostname.
   * @return corsAllowedOrigins
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "CORS Allowed origins for the custom hostname.")

  public List<String> getCorsAllowedOrigins() {
    return corsAllowedOrigins;
  }


  public void setCorsAllowedOrigins(List<String> corsAllowedOrigins) {
    this.corsAllowedOrigins = corsAllowedOrigins;
  }


  public CreateCustomHostnameBody corsEnabled(Boolean corsEnabled) {
    
    this.corsEnabled = corsEnabled;
    return this;
  }

   /**
   * CORS Enabled for the custom hostname.
   * @return corsEnabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "CORS Enabled for the custom hostname.")

  public Boolean getCorsEnabled() {
    return corsEnabled;
  }


  public void setCorsEnabled(Boolean corsEnabled) {
    this.corsEnabled = corsEnabled;
  }


  public CreateCustomHostnameBody hostname(String hostname) {
    
    this.hostname = hostname;
    return this;
  }

   /**
   * The custom hostname where the API will be exposed.
   * @return hostname
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The custom hostname where the API will be exposed.")

  public String getHostname() {
    return hostname;
  }


  public void setHostname(String hostname) {
    this.hostname = hostname;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CreateCustomHostnameBody createCustomHostnameBody = (CreateCustomHostnameBody) o;
    return Objects.equals(this.cookieDomain, createCustomHostnameBody.cookieDomain) &&
        Objects.equals(this.corsAllowedOrigins, createCustomHostnameBody.corsAllowedOrigins) &&
        Objects.equals(this.corsEnabled, createCustomHostnameBody.corsEnabled) &&
        Objects.equals(this.hostname, createCustomHostnameBody.hostname);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cookieDomain, corsAllowedOrigins, corsEnabled, hostname);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CreateCustomHostnameBody {\n");
    sb.append("    cookieDomain: ").append(toIndentedString(cookieDomain)).append("\n");
    sb.append("    corsAllowedOrigins: ").append(toIndentedString(corsAllowedOrigins)).append("\n");
    sb.append("    corsEnabled: ").append(toIndentedString(corsEnabled)).append("\n");
    sb.append("    hostname: ").append(toIndentedString(hostname)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

