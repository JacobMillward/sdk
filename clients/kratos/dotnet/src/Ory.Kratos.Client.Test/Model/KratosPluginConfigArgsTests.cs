/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.0.0-alpha.38
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
    ///  Class for testing KratosPluginConfigArgs
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class KratosPluginConfigArgsTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for KratosPluginConfigArgs
        //private KratosPluginConfigArgs instance;

        public KratosPluginConfigArgsTests()
        {
            // TODO uncomment below to create an instance of KratosPluginConfigArgs
            //instance = new KratosPluginConfigArgs();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of KratosPluginConfigArgs
        /// </summary>
        [Fact]
        public void KratosPluginConfigArgsInstanceTest()
        {
            // TODO uncomment below to test "IsType" KratosPluginConfigArgs
            //Assert.IsType<KratosPluginConfigArgs>(instance);
        }


        /// <summary>
        /// Test the property 'Description'
        /// </summary>
        [Fact]
        public void DescriptionTest()
        {
            // TODO unit test for the property 'Description'
        }
        /// <summary>
        /// Test the property 'Name'
        /// </summary>
        [Fact]
        public void NameTest()
        {
            // TODO unit test for the property 'Name'
        }
        /// <summary>
        /// Test the property 'Settable'
        /// </summary>
        [Fact]
        public void SettableTest()
        {
            // TODO unit test for the property 'Settable'
        }
        /// <summary>
        /// Test the property 'Value'
        /// </summary>
        [Fact]
        public void ValueTest()
        {
            // TODO unit test for the property 'Value'
        }

    }

}