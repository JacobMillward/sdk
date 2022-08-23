/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.15
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using Ory.Client.Api;
using Ory.Client.Model;
using Ory.Client.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace Ory.Client.Test.Model
{
    /// <summary>
    ///  Class for testing ClientOAuth2Client
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class ClientOAuth2ClientTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for ClientOAuth2Client
        //private ClientOAuth2Client instance;

        public ClientOAuth2ClientTests()
        {
            // TODO uncomment below to create an instance of ClientOAuth2Client
            //instance = new ClientOAuth2Client();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of ClientOAuth2Client
        /// </summary>
        [Fact]
        public void ClientOAuth2ClientInstanceTest()
        {
            // TODO uncomment below to test "IsType" ClientOAuth2Client
            //Assert.IsType<ClientOAuth2Client>(instance);
        }


        /// <summary>
        /// Test the property 'AllowedCorsOrigins'
        /// </summary>
        [Fact]
        public void AllowedCorsOriginsTest()
        {
            // TODO unit test for the property 'AllowedCorsOrigins'
        }
        /// <summary>
        /// Test the property 'Audience'
        /// </summary>
        [Fact]
        public void AudienceTest()
        {
            // TODO unit test for the property 'Audience'
        }
        /// <summary>
        /// Test the property 'AuthorizationCodeGrantAccessTokenLifespan'
        /// </summary>
        [Fact]
        public void AuthorizationCodeGrantAccessTokenLifespanTest()
        {
            // TODO unit test for the property 'AuthorizationCodeGrantAccessTokenLifespan'
        }
        /// <summary>
        /// Test the property 'AuthorizationCodeGrantIdTokenLifespan'
        /// </summary>
        [Fact]
        public void AuthorizationCodeGrantIdTokenLifespanTest()
        {
            // TODO unit test for the property 'AuthorizationCodeGrantIdTokenLifespan'
        }
        /// <summary>
        /// Test the property 'AuthorizationCodeGrantRefreshTokenLifespan'
        /// </summary>
        [Fact]
        public void AuthorizationCodeGrantRefreshTokenLifespanTest()
        {
            // TODO unit test for the property 'AuthorizationCodeGrantRefreshTokenLifespan'
        }
        /// <summary>
        /// Test the property 'BackchannelLogoutSessionRequired'
        /// </summary>
        [Fact]
        public void BackchannelLogoutSessionRequiredTest()
        {
            // TODO unit test for the property 'BackchannelLogoutSessionRequired'
        }
        /// <summary>
        /// Test the property 'BackchannelLogoutUri'
        /// </summary>
        [Fact]
        public void BackchannelLogoutUriTest()
        {
            // TODO unit test for the property 'BackchannelLogoutUri'
        }
        /// <summary>
        /// Test the property 'ClientCredentialsGrantAccessTokenLifespan'
        /// </summary>
        [Fact]
        public void ClientCredentialsGrantAccessTokenLifespanTest()
        {
            // TODO unit test for the property 'ClientCredentialsGrantAccessTokenLifespan'
        }
        /// <summary>
        /// Test the property 'ClientId'
        /// </summary>
        [Fact]
        public void ClientIdTest()
        {
            // TODO unit test for the property 'ClientId'
        }
        /// <summary>
        /// Test the property 'ClientName'
        /// </summary>
        [Fact]
        public void ClientNameTest()
        {
            // TODO unit test for the property 'ClientName'
        }
        /// <summary>
        /// Test the property 'ClientSecret'
        /// </summary>
        [Fact]
        public void ClientSecretTest()
        {
            // TODO unit test for the property 'ClientSecret'
        }
        /// <summary>
        /// Test the property 'ClientSecretExpiresAt'
        /// </summary>
        [Fact]
        public void ClientSecretExpiresAtTest()
        {
            // TODO unit test for the property 'ClientSecretExpiresAt'
        }
        /// <summary>
        /// Test the property 'ClientUri'
        /// </summary>
        [Fact]
        public void ClientUriTest()
        {
            // TODO unit test for the property 'ClientUri'
        }
        /// <summary>
        /// Test the property 'Contacts'
        /// </summary>
        [Fact]
        public void ContactsTest()
        {
            // TODO unit test for the property 'Contacts'
        }
        /// <summary>
        /// Test the property 'CreatedAt'
        /// </summary>
        [Fact]
        public void CreatedAtTest()
        {
            // TODO unit test for the property 'CreatedAt'
        }
        /// <summary>
        /// Test the property 'FrontchannelLogoutSessionRequired'
        /// </summary>
        [Fact]
        public void FrontchannelLogoutSessionRequiredTest()
        {
            // TODO unit test for the property 'FrontchannelLogoutSessionRequired'
        }
        /// <summary>
        /// Test the property 'FrontchannelLogoutUri'
        /// </summary>
        [Fact]
        public void FrontchannelLogoutUriTest()
        {
            // TODO unit test for the property 'FrontchannelLogoutUri'
        }
        /// <summary>
        /// Test the property 'GrantTypes'
        /// </summary>
        [Fact]
        public void GrantTypesTest()
        {
            // TODO unit test for the property 'GrantTypes'
        }
        /// <summary>
        /// Test the property 'ImplicitGrantAccessTokenLifespan'
        /// </summary>
        [Fact]
        public void ImplicitGrantAccessTokenLifespanTest()
        {
            // TODO unit test for the property 'ImplicitGrantAccessTokenLifespan'
        }
        /// <summary>
        /// Test the property 'ImplicitGrantIdTokenLifespan'
        /// </summary>
        [Fact]
        public void ImplicitGrantIdTokenLifespanTest()
        {
            // TODO unit test for the property 'ImplicitGrantIdTokenLifespan'
        }
        /// <summary>
        /// Test the property 'Jwks'
        /// </summary>
        [Fact]
        public void JwksTest()
        {
            // TODO unit test for the property 'Jwks'
        }
        /// <summary>
        /// Test the property 'JwksUri'
        /// </summary>
        [Fact]
        public void JwksUriTest()
        {
            // TODO unit test for the property 'JwksUri'
        }
        /// <summary>
        /// Test the property 'JwtBearerGrantAccessTokenLifespan'
        /// </summary>
        [Fact]
        public void JwtBearerGrantAccessTokenLifespanTest()
        {
            // TODO unit test for the property 'JwtBearerGrantAccessTokenLifespan'
        }
        /// <summary>
        /// Test the property 'LogoUri'
        /// </summary>
        [Fact]
        public void LogoUriTest()
        {
            // TODO unit test for the property 'LogoUri'
        }
        /// <summary>
        /// Test the property 'Metadata'
        /// </summary>
        [Fact]
        public void MetadataTest()
        {
            // TODO unit test for the property 'Metadata'
        }
        /// <summary>
        /// Test the property 'Owner'
        /// </summary>
        [Fact]
        public void OwnerTest()
        {
            // TODO unit test for the property 'Owner'
        }
        /// <summary>
        /// Test the property 'PasswordGrantAccessTokenLifespan'
        /// </summary>
        [Fact]
        public void PasswordGrantAccessTokenLifespanTest()
        {
            // TODO unit test for the property 'PasswordGrantAccessTokenLifespan'
        }
        /// <summary>
        /// Test the property 'PasswordGrantRefreshTokenLifespan'
        /// </summary>
        [Fact]
        public void PasswordGrantRefreshTokenLifespanTest()
        {
            // TODO unit test for the property 'PasswordGrantRefreshTokenLifespan'
        }
        /// <summary>
        /// Test the property 'PolicyUri'
        /// </summary>
        [Fact]
        public void PolicyUriTest()
        {
            // TODO unit test for the property 'PolicyUri'
        }
        /// <summary>
        /// Test the property 'PostLogoutRedirectUris'
        /// </summary>
        [Fact]
        public void PostLogoutRedirectUrisTest()
        {
            // TODO unit test for the property 'PostLogoutRedirectUris'
        }
        /// <summary>
        /// Test the property 'RedirectUris'
        /// </summary>
        [Fact]
        public void RedirectUrisTest()
        {
            // TODO unit test for the property 'RedirectUris'
        }
        /// <summary>
        /// Test the property 'RefreshTokenGrantAccessTokenLifespan'
        /// </summary>
        [Fact]
        public void RefreshTokenGrantAccessTokenLifespanTest()
        {
            // TODO unit test for the property 'RefreshTokenGrantAccessTokenLifespan'
        }
        /// <summary>
        /// Test the property 'RefreshTokenGrantIdTokenLifespan'
        /// </summary>
        [Fact]
        public void RefreshTokenGrantIdTokenLifespanTest()
        {
            // TODO unit test for the property 'RefreshTokenGrantIdTokenLifespan'
        }
        /// <summary>
        /// Test the property 'RefreshTokenGrantRefreshTokenLifespan'
        /// </summary>
        [Fact]
        public void RefreshTokenGrantRefreshTokenLifespanTest()
        {
            // TODO unit test for the property 'RefreshTokenGrantRefreshTokenLifespan'
        }
        /// <summary>
        /// Test the property 'RegistrationAccessToken'
        /// </summary>
        [Fact]
        public void RegistrationAccessTokenTest()
        {
            // TODO unit test for the property 'RegistrationAccessToken'
        }
        /// <summary>
        /// Test the property 'RegistrationClientUri'
        /// </summary>
        [Fact]
        public void RegistrationClientUriTest()
        {
            // TODO unit test for the property 'RegistrationClientUri'
        }
        /// <summary>
        /// Test the property 'RequestObjectSigningAlg'
        /// </summary>
        [Fact]
        public void RequestObjectSigningAlgTest()
        {
            // TODO unit test for the property 'RequestObjectSigningAlg'
        }
        /// <summary>
        /// Test the property 'RequestUris'
        /// </summary>
        [Fact]
        public void RequestUrisTest()
        {
            // TODO unit test for the property 'RequestUris'
        }
        /// <summary>
        /// Test the property 'ResponseTypes'
        /// </summary>
        [Fact]
        public void ResponseTypesTest()
        {
            // TODO unit test for the property 'ResponseTypes'
        }
        /// <summary>
        /// Test the property 'Scope'
        /// </summary>
        [Fact]
        public void ScopeTest()
        {
            // TODO unit test for the property 'Scope'
        }
        /// <summary>
        /// Test the property 'SectorIdentifierUri'
        /// </summary>
        [Fact]
        public void SectorIdentifierUriTest()
        {
            // TODO unit test for the property 'SectorIdentifierUri'
        }
        /// <summary>
        /// Test the property 'SubjectType'
        /// </summary>
        [Fact]
        public void SubjectTypeTest()
        {
            // TODO unit test for the property 'SubjectType'
        }
        /// <summary>
        /// Test the property 'TokenEndpointAuthMethod'
        /// </summary>
        [Fact]
        public void TokenEndpointAuthMethodTest()
        {
            // TODO unit test for the property 'TokenEndpointAuthMethod'
        }
        /// <summary>
        /// Test the property 'TokenEndpointAuthSigningAlg'
        /// </summary>
        [Fact]
        public void TokenEndpointAuthSigningAlgTest()
        {
            // TODO unit test for the property 'TokenEndpointAuthSigningAlg'
        }
        /// <summary>
        /// Test the property 'TosUri'
        /// </summary>
        [Fact]
        public void TosUriTest()
        {
            // TODO unit test for the property 'TosUri'
        }
        /// <summary>
        /// Test the property 'UpdatedAt'
        /// </summary>
        [Fact]
        public void UpdatedAtTest()
        {
            // TODO unit test for the property 'UpdatedAt'
        }
        /// <summary>
        /// Test the property 'UserinfoSignedResponseAlg'
        /// </summary>
        [Fact]
        public void UserinfoSignedResponseAlgTest()
        {
            // TODO unit test for the property 'UserinfoSignedResponseAlg'
        }

    }

}