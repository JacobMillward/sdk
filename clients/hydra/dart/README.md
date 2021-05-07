# ory_hydra_client
Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.

This Dart package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: v1.9.1
- Build package: org.openapitools.codegen.languages.DartClientCodegen

## Requirements

Dart 2.0 or later

## Installation & Usage

### Github
If this Dart package is published to Github, add the following dependency to your pubspec.yaml
```
dependencies:
  ory_hydra_client:
    git: https://github.com/ory/sdk.git
```

### Local
To use the package in your local drive, add the following dependency to your pubspec.yaml
```
dependencies:
  ory_hydra_client:
    path: /path/to/ory_hydra_client
```

## Tests

TODO

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```dart
import 'package:ory_hydra_client/api.dart';


final api_instance = AdminApi();
final consentChallenge = consentChallenge_example; // String | 
final body = AcceptConsentRequest(); // AcceptConsentRequest | 

try {
    final result = api_instance.acceptConsentRequest(consentChallenge, body);
    print(result);
} catch (e) {
    print('Exception when calling AdminApi->acceptConsentRequest: $e\n');
}

```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**acceptConsentRequest**](doc//AdminApi.md#acceptconsentrequest) | **PUT** /oauth2/auth/requests/consent/accept | Accept a Consent Request
*AdminApi* | [**acceptLoginRequest**](doc//AdminApi.md#acceptloginrequest) | **PUT** /oauth2/auth/requests/login/accept | Accept a Login Request
*AdminApi* | [**acceptLogoutRequest**](doc//AdminApi.md#acceptlogoutrequest) | **PUT** /oauth2/auth/requests/logout/accept | Accept a Logout Request
*AdminApi* | [**createJsonWebKeySet**](doc//AdminApi.md#createjsonwebkeyset) | **POST** /keys/{set} | Generate a New JSON Web Key
*AdminApi* | [**createOAuth2Client**](doc//AdminApi.md#createoauth2client) | **POST** /clients | Create an OAuth 2.0 Client
*AdminApi* | [**deleteJsonWebKey**](doc//AdminApi.md#deletejsonwebkey) | **DELETE** /keys/{set}/{kid} | Delete a JSON Web Key
*AdminApi* | [**deleteJsonWebKeySet**](doc//AdminApi.md#deletejsonwebkeyset) | **DELETE** /keys/{set} | Delete a JSON Web Key Set
*AdminApi* | [**deleteOAuth2Client**](doc//AdminApi.md#deleteoauth2client) | **DELETE** /clients/{id} | Deletes an OAuth 2.0 Client
*AdminApi* | [**deleteOAuth2Token**](doc//AdminApi.md#deleteoauth2token) | **DELETE** /oauth2/tokens | Delete OAuth2 Access Tokens from a Client
*AdminApi* | [**flushInactiveOAuth2Tokens**](doc//AdminApi.md#flushinactiveoauth2tokens) | **POST** /oauth2/flush | Flush Expired OAuth2 Access Tokens
*AdminApi* | [**getConsentRequest**](doc//AdminApi.md#getconsentrequest) | **GET** /oauth2/auth/requests/consent | Get Consent Request Information
*AdminApi* | [**getJsonWebKey**](doc//AdminApi.md#getjsonwebkey) | **GET** /keys/{set}/{kid} | Fetch a JSON Web Key
*AdminApi* | [**getJsonWebKeySet**](doc//AdminApi.md#getjsonwebkeyset) | **GET** /keys/{set} | Retrieve a JSON Web Key Set
*AdminApi* | [**getLoginRequest**](doc//AdminApi.md#getloginrequest) | **GET** /oauth2/auth/requests/login | Get a Login Request
*AdminApi* | [**getLogoutRequest**](doc//AdminApi.md#getlogoutrequest) | **GET** /oauth2/auth/requests/logout | Get a Logout Request
*AdminApi* | [**getOAuth2Client**](doc//AdminApi.md#getoauth2client) | **GET** /clients/{id} | Get an OAuth 2.0 Client.
*AdminApi* | [**getVersion**](doc//AdminApi.md#getversion) | **GET** /version | Get Service Version
*AdminApi* | [**introspectOAuth2Token**](doc//AdminApi.md#introspectoauth2token) | **POST** /oauth2/introspect | Introspect OAuth2 Tokens
*AdminApi* | [**isInstanceAlive**](doc//AdminApi.md#isinstancealive) | **GET** /health/alive | Check Alive Status
*AdminApi* | [**listOAuth2Clients**](doc//AdminApi.md#listoauth2clients) | **GET** /clients | List OAuth 2.0 Clients
*AdminApi* | [**listSubjectConsentSessions**](doc//AdminApi.md#listsubjectconsentsessions) | **GET** /oauth2/auth/sessions/consent | Lists All Consent Sessions of a Subject
*AdminApi* | [**prometheus**](doc//AdminApi.md#prometheus) | **GET** /metrics/prometheus | Get Snapshot Metrics from the Hydra Service.
*AdminApi* | [**rejectConsentRequest**](doc//AdminApi.md#rejectconsentrequest) | **PUT** /oauth2/auth/requests/consent/reject | Reject a Consent Request
*AdminApi* | [**rejectLoginRequest**](doc//AdminApi.md#rejectloginrequest) | **PUT** /oauth2/auth/requests/login/reject | Reject a Login Request
*AdminApi* | [**rejectLogoutRequest**](doc//AdminApi.md#rejectlogoutrequest) | **PUT** /oauth2/auth/requests/logout/reject | Reject a Logout Request
*AdminApi* | [**revokeAuthenticationSession**](doc//AdminApi.md#revokeauthenticationsession) | **DELETE** /oauth2/auth/sessions/login | Invalidates All Login Sessions of a Certain User Invalidates a Subject's Authentication Session
*AdminApi* | [**revokeConsentSessions**](doc//AdminApi.md#revokeconsentsessions) | **DELETE** /oauth2/auth/sessions/consent | Revokes Consent Sessions of a Subject for a Specific OAuth 2.0 Client
*AdminApi* | [**updateJsonWebKey**](doc//AdminApi.md#updatejsonwebkey) | **PUT** /keys/{set}/{kid} | Update a JSON Web Key
*AdminApi* | [**updateJsonWebKeySet**](doc//AdminApi.md#updatejsonwebkeyset) | **PUT** /keys/{set} | Update a JSON Web Key Set
*AdminApi* | [**updateOAuth2Client**](doc//AdminApi.md#updateoauth2client) | **PUT** /clients/{id} | Update an OAuth 2.0 Client
*PublicApi* | [**disconnectUser**](doc//PublicApi.md#disconnectuser) | **GET** /oauth2/sessions/logout | OpenID Connect Front-Backchannel Enabled Logout
*PublicApi* | [**discoverOpenIDConfiguration**](doc//PublicApi.md#discoveropenidconfiguration) | **GET** /.well-known/openid-configuration | OpenID Connect Discovery
*PublicApi* | [**isInstanceReady**](doc//PublicApi.md#isinstanceready) | **GET** /health/ready | Check Readiness Status
*PublicApi* | [**oauth2Token**](doc//PublicApi.md#oauth2token) | **POST** /oauth2/token | The OAuth 2.0 Token Endpoint
*PublicApi* | [**oauthAuth**](doc//PublicApi.md#oauthauth) | **GET** /oauth2/auth | The OAuth 2.0 Authorize Endpoint
*PublicApi* | [**revokeOAuth2Token**](doc//PublicApi.md#revokeoauth2token) | **POST** /oauth2/revoke | Revoke OAuth2 Tokens
*PublicApi* | [**userinfo**](doc//PublicApi.md#userinfo) | **GET** /userinfo | OpenID Connect Userinfo
*PublicApi* | [**wellKnown**](doc//PublicApi.md#wellknown) | **GET** /.well-known/jwks.json | JSON Web Keys Discovery


## Documentation For Models

 - [AcceptConsentRequest](doc//AcceptConsentRequest.md)
 - [AcceptLoginRequest](doc//AcceptLoginRequest.md)
 - [CompletedRequest](doc//CompletedRequest.md)
 - [ConsentRequest](doc//ConsentRequest.md)
 - [ConsentRequestSession](doc//ConsentRequestSession.md)
 - [ContainerWaitOKBodyError](doc//ContainerWaitOKBodyError.md)
 - [FlushInactiveOAuth2TokensRequest](doc//FlushInactiveOAuth2TokensRequest.md)
 - [GenericError](doc//GenericError.md)
 - [HealthNotReadyStatus](doc//HealthNotReadyStatus.md)
 - [HealthStatus](doc//HealthStatus.md)
 - [JSONWebKey](doc//JSONWebKey.md)
 - [JSONWebKeySet](doc//JSONWebKeySet.md)
 - [JsonWebKeySetGeneratorRequest](doc//JsonWebKeySetGeneratorRequest.md)
 - [LoginRequest](doc//LoginRequest.md)
 - [LogoutRequest](doc//LogoutRequest.md)
 - [OAuth2Client](doc//OAuth2Client.md)
 - [OAuth2TokenIntrospection](doc//OAuth2TokenIntrospection.md)
 - [Oauth2TokenResponse](doc//Oauth2TokenResponse.md)
 - [OpenIDConnectContext](doc//OpenIDConnectContext.md)
 - [PluginConfig](doc//PluginConfig.md)
 - [PluginConfigArgs](doc//PluginConfigArgs.md)
 - [PluginConfigInterface](doc//PluginConfigInterface.md)
 - [PluginConfigLinux](doc//PluginConfigLinux.md)
 - [PluginConfigNetwork](doc//PluginConfigNetwork.md)
 - [PluginConfigRootfs](doc//PluginConfigRootfs.md)
 - [PluginConfigUser](doc//PluginConfigUser.md)
 - [PluginDevice](doc//PluginDevice.md)
 - [PluginEnv](doc//PluginEnv.md)
 - [PluginInterfaceType](doc//PluginInterfaceType.md)
 - [PluginMount](doc//PluginMount.md)
 - [PluginSettings](doc//PluginSettings.md)
 - [PreviousConsentSession](doc//PreviousConsentSession.md)
 - [RejectRequest](doc//RejectRequest.md)
 - [UserinfoResponse](doc//UserinfoResponse.md)
 - [Version](doc//Version.md)
 - [VolumeUsageData](doc//VolumeUsageData.md)
 - [WellKnown](doc//WellKnown.md)


## Documentation For Authorization


## basic

- **Type**: HTTP Basic authentication

## oauth2

- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://hydra.demo.ory.sh/oauth2/auth
- **Scopes**: 
 - **offline**: A scope required when requesting refresh tokens (alias for `offline_access`)
 - **offline_access**: A scope required when requesting refresh tokens
 - **openid**: Request an OpenID Connect ID Token


## Author



