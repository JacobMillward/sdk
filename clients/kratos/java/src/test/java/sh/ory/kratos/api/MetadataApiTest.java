/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.6-alpha.5
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.api;

import sh.ory.kratos.ApiException;
import sh.ory.kratos.model.GenericError;
import sh.ory.kratos.model.InlineResponse200;
import sh.ory.kratos.model.InlineResponse2001;
import sh.ory.kratos.model.InlineResponse503;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for MetadataApi
 */
@Ignore
public class MetadataApiTest {

    private final MetadataApi api = new MetadataApi();

    
    /**
     * Return Running Software Version.
     *
     * This endpoint returns the version of Ory Kratos.  If the service supports TLS Edge Termination, this endpoint does not require the &#x60;X-Forwarded-Proto&#x60; header to be set.  Be aware that if you are running multiple nodes of this service, the version will never refer to the cluster state, only to a single instance.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getVersionTest() throws ApiException {
        InlineResponse2001 response = api.getVersion();

        // TODO: test validations
    }
    
    /**
     * Check HTTP Server Status
     *
     * This endpoint returns a HTTP 200 status code when Ory Kratos is accepting incoming HTTP requests. This status does currently not include checks whether the database connection is working.  If the service supports TLS Edge Termination, this endpoint does not require the &#x60;X-Forwarded-Proto&#x60; header to be set.  Be aware that if you are running multiple nodes of this service, the health status will never refer to the cluster state, only to a single instance.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void isAliveTest() throws ApiException {
        InlineResponse200 response = api.isAlive();

        // TODO: test validations
    }
    
    /**
     * Check HTTP Server and Database Status
     *
     * This endpoint returns a HTTP 200 status code when Ory Kratos is up running and the environment dependencies (e.g. the database) are responsive as well.  If the service supports TLS Edge Termination, this endpoint does not require the &#x60;X-Forwarded-Proto&#x60; header to be set.  Be aware that if you are running multiple nodes of Ory Kratos, the health status will never refer to the cluster state, only to a single instance.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void isReadyTest() throws ApiException {
        InlineResponse200 response = api.isReady();

        // TODO: test validations
    }
    
    /**
     * Get snapshot metrics from the service. If you&#39;re using k8s, you can then add annotations to your deployment like so:
     *
     * &#x60;&#x60;&#x60; metadata: annotations: prometheus.io/port: \&quot;4434\&quot; prometheus.io/path: \&quot;/metrics/prometheus\&quot; &#x60;&#x60;&#x60;
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void prometheusTest() throws ApiException {
        api.prometheus();

        // TODO: test validations
    }
    
}
