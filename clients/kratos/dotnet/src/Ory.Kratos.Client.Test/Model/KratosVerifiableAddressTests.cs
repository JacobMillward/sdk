/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.15
 * Contact: hi@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using Ory.Kratos.Client.Api;
using Ory.Kratos.Client.Model;
using Ory.Kratos.Client.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace Ory.Kratos.Client.Test.Model
{
    /// <summary>
    ///  Class for testing KratosVerifiableAddress
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class KratosVerifiableAddressTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for KratosVerifiableAddress
        //private KratosVerifiableAddress instance;

        public KratosVerifiableAddressTests()
        {
            // TODO uncomment below to create an instance of KratosVerifiableAddress
            //instance = new KratosVerifiableAddress();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of KratosVerifiableAddress
        /// </summary>
        [Fact]
        public void KratosVerifiableAddressInstanceTest()
        {
            // TODO uncomment below to test "IsType" KratosVerifiableAddress
            //Assert.IsType<KratosVerifiableAddress>(instance);
        }


        /// <summary>
        /// Test the property 'Id'
        /// </summary>
        [Fact]
        public void IdTest()
        {
            // TODO unit test for the property 'Id'
        }
        /// <summary>
        /// Test the property 'Status'
        /// </summary>
        [Fact]
        public void StatusTest()
        {
            // TODO unit test for the property 'Status'
        }
        /// <summary>
        /// Test the property 'Value'
        /// </summary>
        [Fact]
        public void ValueTest()
        {
            // TODO unit test for the property 'Value'
        }
        /// <summary>
        /// Test the property 'Verified'
        /// </summary>
        [Fact]
        public void VerifiedTest()
        {
            // TODO unit test for the property 'Verified'
        }
        /// <summary>
        /// Test the property 'VerifiedAt'
        /// </summary>
        [Fact]
        public void VerifiedAtTest()
        {
            // TODO unit test for the property 'VerifiedAt'
        }
        /// <summary>
        /// Test the property 'Via'
        /// </summary>
        [Fact]
        public void ViaTest()
        {
            // TODO unit test for the property 'Via'
        }

    }

}