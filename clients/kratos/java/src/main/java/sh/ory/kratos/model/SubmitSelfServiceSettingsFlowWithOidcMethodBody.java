/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.1-alpha.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.model;

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

/**
 * nolint:deadcode,unused
 */
@ApiModel(description = "nolint:deadcode,unused")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-07-22T18:06:07.711120196Z[Etc/UTC]")
public class SubmitSelfServiceSettingsFlowWithOidcMethodBody {
  public static final String SERIALIZED_NAME_FLOW = "flow";
  @SerializedName(SERIALIZED_NAME_FLOW)
  private String flow;

  public static final String SERIALIZED_NAME_LINK = "link";
  @SerializedName(SERIALIZED_NAME_LINK)
  private String link;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public static final String SERIALIZED_NAME_UNLINK = "unlink";
  @SerializedName(SERIALIZED_NAME_UNLINK)
  private String unlink;


  public SubmitSelfServiceSettingsFlowWithOidcMethodBody flow(String flow) {
    
    this.flow = flow;
    return this;
  }

   /**
   * Flow ID is the flow&#39;s ID.  in: query
   * @return flow
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Flow ID is the flow's ID.  in: query")

  public String getFlow() {
    return flow;
  }


  public void setFlow(String flow) {
    this.flow = flow;
  }


  public SubmitSelfServiceSettingsFlowWithOidcMethodBody link(String link) {
    
    this.link = link;
    return this;
  }

   /**
   * Link this provider  Either this or &#x60;unlink&#x60; must be set.  type: string in: body
   * @return link
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Link this provider  Either this or `unlink` must be set.  type: string in: body")

  public String getLink() {
    return link;
  }


  public void setLink(String link) {
    this.link = link;
  }


  public SubmitSelfServiceSettingsFlowWithOidcMethodBody method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method  Should be set to profile when trying to update a profile.
   * @return method
  **/
  @ApiModelProperty(required = true, value = "Method  Should be set to profile when trying to update a profile.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public SubmitSelfServiceSettingsFlowWithOidcMethodBody unlink(String unlink) {
    
    this.unlink = unlink;
    return this;
  }

   /**
   * Unlink this provider  Either this or &#x60;link&#x60; must be set.  type: string in: body
   * @return unlink
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Unlink this provider  Either this or `link` must be set.  type: string in: body")

  public String getUnlink() {
    return unlink;
  }


  public void setUnlink(String unlink) {
    this.unlink = unlink;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SubmitSelfServiceSettingsFlowWithOidcMethodBody submitSelfServiceSettingsFlowWithOidcMethodBody = (SubmitSelfServiceSettingsFlowWithOidcMethodBody) o;
    return Objects.equals(this.flow, submitSelfServiceSettingsFlowWithOidcMethodBody.flow) &&
        Objects.equals(this.link, submitSelfServiceSettingsFlowWithOidcMethodBody.link) &&
        Objects.equals(this.method, submitSelfServiceSettingsFlowWithOidcMethodBody.method) &&
        Objects.equals(this.unlink, submitSelfServiceSettingsFlowWithOidcMethodBody.unlink);
  }

  @Override
  public int hashCode() {
    return Objects.hash(flow, link, method, unlink);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceSettingsFlowWithOidcMethodBody {\n");
    sb.append("    flow: ").append(toIndentedString(flow)).append("\n");
    sb.append("    link: ").append(toIndentedString(link)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    unlink: ").append(toIndentedString(unlink)).append("\n");
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
