
## DaVinci Connection Definitions

Below is a list of all available DaVinci Connections available for use in `davinci_connection` resource. 
If the `value` type of a property is not defined it must be inferred.


  
### 1Kosmos connector

**Connector Display Name**: 1Kosmos connector

**Connector ID** - schema `connectorId`: connector1Kosmos

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### AWS Lambda

**Connector Display Name**: AWS Lambda

**Connector ID** - schema `connectorId`: connectorAWSLambda

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Access Key Id | `accessKeyId` | `` | Access key to connect to AWS Environment |
| AWS Region | `region` | `string` | AWS Region where the Lambda function is created |
| AWS Secret Key | `secretAccessKey` | `` | Secret Key to access the AWS | 




### AWS Login

**Connector Display Name**: AWS Login

**Connector ID** - schema `connectorId`: awsIdpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### AWS Secrets Manager

**Connector Display Name**: AWS Secrets Manager

**Connector ID** - schema `connectorId`: connectorAmazonAwsSecretsManager

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| AWS Access Key | `accessKeyId` | `` | The AWS Access Key |
| AWS Region | `region` | `string` | The AWS Region |
| AWS Access Secret | `secretAccessKey` | `` | The AWS Access Secret | 




### AbuseIPDB

**Connector Display Name**: AbuseIPDB

**Connector ID** - schema `connectorId`: connectorAbuseipdb

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | API Key gathered from AbuseIPDB tenant | 




### Acuant

**Connector Display Name**: Acuant

**Connector ID** - schema `connectorId`: connectorAcuant

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Amazon DynamoDB

**Connector Display Name**: Amazon DynamoDB

**Connector ID** - schema `connectorId`: connectorAmazonDynamoDB

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| AWS Access Key | `awsAccessKey` | `string` | Your AWS Access Key |
| AWS Access Secret | `awsAccessSecret` | `string` | Access Secret corresponding with Access Key found in Your Security Credentials |
| AWS Region | `awsRegion` | `string` | The AWS Region you are using the connector for. | 




### Amazon Simple Email Service

**Connector Display Name**: Amazon Simple Email Service

**Connector ID** - schema `connectorId`: amazonSimpleEmailConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| AWS Access Key | `awsAccessKey` | `string` |  |
| AWS Access Secret | `awsAccessSecret` | `string` |  |
| AWS Region | `awsRegion` | `string` |  |
| From | `from` | `` |  | 




### Amazon Simple Notification Service

**Connector Display Name**: Amazon Simple Notification Service

**Connector ID** - schema `connectorId`: amazonSimpleNotificationConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Annotation

**Connector Display Name**: Annotation

**Connector ID** - schema `connectorId`: annotationConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Apple Login

**Connector Display Name**: Apple Login

**Connector ID** - schema `connectorId`: appleConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Argyle

**Connector Display Name**: Argyle

**Connector ID** - schema `connectorId`: argyleConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Server URL | `apiUrl` | `string` |  |
| Client ID | `clientId` | `string` |  |
| Client Secret | `clientSecret` | `string` |  |
| Argyle Loader Javascript Web URL | `javascriptWebUrl` | `` | Argyle loader javascript web URL |
| Plugin Key | `pluginKey` | `` |  | 




### AuthID

**Connector Display Name**: AuthID

**Connector ID** - schema `connectorId`: connectorAuthid

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### AuthenticID

**Connector Display Name**: AuthenticID

**Connector ID** - schema `connectorId`: authenticIdConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Account Access Key | `accountAccessKey` | `string` | Your Account Access Key provided by AuthenticID  |
| Android SDK Licence Key | `androidSDKLicenseKey` | `string` | License key is whitelisted for specific package name |
| REST API URL | `apiUrl` | `string` | AuthenticID REST API URL for sandbox/production environments |
| Base URL | `baseUrl` | `string` | AuthenticID API URL for sandbox/production environments |
| Client Certificate | `clientCertificate` | `` | Your Client Certificate provided by AuthenticID |
| Client Key | `clientKey` | `` | Your Client Key provided by AuthenticID |
| iOS SDK Licence Key | `iOSSDKLicenseKey` | `string` | License key is whitelisted for specific bundle id |
| Certificate Passphrase | `passphrase` | `` | Your Certificate Passphrase provided by AuthenticID |
| Secret Token | `secretToken` | `string` | Your Secret Token provided by AuthenticID | 




### Azure AD User Management

**Connector Display Name**: Azure AD User Management

**Connector ID** - schema `connectorId`: azureUserManagementConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API URL | `baseUrl` | `string` | The Microsoft API URL to target. For a custom value, select Use Custom API URL and enter a value in the Custom API URL field. |
| Custom API URL | `customApiUrl` | `` | The URL for the Microsoft Graph API, such as "https://graph.microsoft.com/v1.0". |
| Custom Parameters | `customAuth` | `array` |  | 




### BambooHR

**Connector Display Name**: BambooHR

**Connector ID** - schema `connectorId`: bambooConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` |  |
| Base URL | `baseUrl` | `string` |  BambooHR Base URL |
| Company Sub Domain | `companySubDomain` | `string` |  Your BambooHR subdomain |
| Flow ID | `flowId` | `string` | Select ID of the flow to execute when BambooHR sends a webhook |
| Singular Key Webhook URL | `skWebhookUri` | `string` | Use this url as the Webhook URL in the Third Party Integration's configuration |
| Webhook Token | `webhookToken` | `string` | Create a webhook token and configure it in the bambooHR webhook url. | 




### Berbix

**Connector Display Name**: Berbix

**Connector ID** - schema `connectorId`: connectorBerbix

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Domain Name | `domainName` | `` | Provide Berbix domain name |
| Path | `path` | `` | Provide path of the API |
| User Name | `username` | `string` | Provide your Berbix user name | 




### Beyond Identity

**Connector Display Name**: Beyond Identity

**Connector ID** - schema `connectorId`: connectorBeyondIdentity

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### BioCatch

**Connector Display Name**: BioCatch

**Connector ID** - schema `connectorId`: biocatchConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Server URL | `apiUrl` | `string` |  |
| Customer ID | `customerId` | `string` |  |
| Javascript CDN URL | `javascriptCdnUrl` | `string` |  |
| SDK Token | `sdkToken` | `string` |  |
| Truth-mapping API Key | `truthApiKey` | `string` | Fraudulent/Genuine Session Reporting API Key |
| Truth-mapping API URL | `truthApiUrl` | `string` | Fraudulent/Genuine Session Reporting | 




### Bitbucket Login

**Connector Display Name**: Bitbucket Login

**Connector ID** - schema `connectorId`: bitbucketIdpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### Challenge

**Connector Display Name**: Challenge

**Connector ID** - schema `connectorId`: challengeConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Circle Access

**Connector Display Name**: Circle Access

**Connector ID** - schema `connectorId`: connectorCircleAccess

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| App Key | `appKey` | `` | App Key |
| Custom Parameters | `customAuth` | `array` |  |
| Login Url | `loginUrl` | `` | The URL of your Circle Access login |
| Read Key | `readKey` | `` | Read Key |
| Application Return To URL | `returnToUrl` | `string` | When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application. |
| Write Key | `writeKey` | `` | Write key | 




### Clearbit

**Connector Display Name**: Clearbit

**Connector ID** - schema `connectorId`: connectorClearbit

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | Clearbit API Key |
| Risk API Version | `riskApiVersion` | `` | Clearbit - Risk API Version |
| Person API Version | `version` | `string` | Clearbit - Person API Version | 




### Cloudflare

**Connector Display Name**: Cloudflare

**Connector ID** - schema `connectorId`: connectorCloudflare

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Account ID | `accountId` | `` | Cloudflare Account ID |
| API Token | `apiToken` | `` | Cloudflare API Token | 




### Code Snippet

**Connector Display Name**: Code Snippet

**Connector ID** - schema `connectorId`: codeSnippetConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Code Snippet | `code` | `string` | Follow example for code. |
| Input Schema | `inputSchema` | `string` | Follow example for JSON schema. |
| Output Schema | `outputSchema` | `string` | Follow example for JSON schema. | 




### Comply Advantage

**Connector Display Name**: Comply Advantage

**Connector ID** - schema `connectorId`: complyAdvatangeConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | API Key is the API key that you can retrieve from Comply Advantage Admin Portal |
| Base URL | `baseUrl` | `string` | Comply Advantage API URL for sandbox/production environments | 




### Cookie

**Connector Display Name**: Cookie

**Connector ID** - schema `connectorId`: cookieConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| HMAC Signing Key | `hmacSigningKey` | `string` | Base64 encoded 256 bit key | 




### Credova

**Connector Display Name**: Credova

**Connector ID** - schema `connectorId`: credovaConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Base URL | `baseUrl` | `string` | Base URL for Credova API |
| Credova Password | `password` | `string` | Password for the Credova Developer Portal |
| Credova Username | `username` | `string` | Username for the Credova Developer Portal | 




### CrowdStrike

**Connector Display Name**: CrowdStrike

**Connector ID** - schema `connectorId`: crowdStrikeConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| CrowdStrike Base URL | `baseURL` | `` | The base URL of the CrowdStrike environment. |
| Client ID | `clientId` | `string` | The Client ID of the application in CrowdStrike. |
| Client Secret | `clientSecret` | `string` | The Client Secret provided by CrowdStrike. | 




### Daon IDV

**Connector Display Name**: Daon IDV

**Connector ID** - schema `connectorId`: connectorDaonidv

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### Daon IdentityX

**Connector Display Name**: Daon IdentityX

**Connector ID** - schema `connectorId`: daonConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Base URL | `apiUrl` | `string` | The protocol, host and base path to the IdX API. E.g. https://api.identityx-cloud.com/tenant1/IdentityXServices/rest/v1 |
| Admin Password | `password` | `string` | The password of the user to authenticate API calls. |
| Admin Username | `username` | `string` | The userId to authenticate API calls. | 




### Data Zoo

**Connector Display Name**: Data Zoo

**Connector ID** - schema `connectorId`: dataZooConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Data Zoo Password | `password` | `string` |  |
| Data Zoo Username | `username` | `string` |  | 




### Datadog API

**Connector Display Name**: Datadog API

**Connector ID** - schema `connectorId`: connector-oai-datadogapi

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authentication API Key | `authApiKey` | `` | The API key for an account that has access to the Datadog API. |
| Authentication Application Key | `authApplicationKey` | `` | The Application key for an account that has access to the Datadog API. |
| API URL | `basePath` | `` | The base URL for contacting the Datadog API, such as "https://api.us3.datadoghq.com". | 




### DeBounce

**Connector Display Name**: DeBounce

**Connector ID** - schema `connectorId`: connectorDeBounce

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | A DeBounce API Key is physically a token/code of 13 random alphanumeric characters. If you need to create an API key, please log in to your DeBounce account and then navigate to the API section. | 




### Device Policy

**Connector Display Name**: Device Policy

**Connector ID** - schema `connectorId`: devicePolicyConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### DigiLocker

**Connector Display Name**: DigiLocker

**Connector ID** - schema `connectorId`: digilockerConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### Digidentity

**Connector Display Name**: Digidentity

**Connector ID** - schema `connectorId`: digidentityConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### Duo

**Connector Display Name**: Duo

**Connector ID** - schema `connectorId`: duoConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Entrust

**Connector Display Name**: Entrust

**Connector ID** - schema `connectorId`: entrustConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Application ID | `applicationId` | `string` | The application ID for the Identity as a Service application. |
| Service Domain | `serviceDomain` | `` | The domain of the Entrust service. Format is '<customer>.<region>.trustedauth.com'. For example, 'mycompany.us.trustedauth.com'. | 




### Equifax

**Connector Display Name**: Equifax

**Connector ID** - schema `connectorId`: equifaxConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Base URL | `baseUrl` | `string` | Base URL for Equifax API |
| Client ID | `clientId` | `string` | When you Create a New App, Equifax will assign a Client ID per environment for the API Product |
| Client Secret | `clientSecret` | `string` | When you Create a New App, Equifax will assign a Client Secret per environment for the API Product |
| SOAP API Environment | `equifaxSoapApiEnvironment` | `string` | SOAP API WSDL Environment. |
| Member Number | `memberNumber` | `` | Unique Identifier of Customer. Please contact Equifax Sales Representative during client onboarding for this value. |
| Password for SOAP API | `password` | `string` | Password provided by Equifax for SOAP API |
| Username for SOAP API | `username` | `string` | Username provided by Equifax for SOAP API | 




### Error Message

**Connector Display Name**: Error Message

**Connector ID** - schema `connectorId`: errorConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Facebook Login

**Connector Display Name**: Facebook Login

**Connector ID** - schema `connectorId`: facebookIdpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### Fingerprint JS

**Connector Display Name**: Fingerprint JS

**Connector ID** - schema `connectorId`: fingerprintjsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Fingerprint Subscription API Token | `apiToken` | `` |  |
| Javascript CDN URL | `javascriptCdnUrl` | `string` |  |
| Fingerprint Subscription Browser Token | `token` | `string` |  | 




### Finicity

**Connector Display Name**: Finicity

**Connector ID** - schema `connectorId`: finicityConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Finicity App Key | `appKey` | `` | Finicity App Key from Finicity Developer Portal |
| Base URL | `baseUrl` | `string` | Base URL for Finicity API |
| Partner ID | `partnerId` | `` | The partner id you can obtain from your Finicity developer dashboard |
| Partner Secret | `partnerSecret` | `` | Partner Secret from Finicity Developer Portal | 




### Flow Analytics

**Connector Display Name**: Flow Analytics

**Connector ID** - schema `connectorId`: analyticsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Flow Conductor

**Connector Display Name**: Flow Conductor

**Connector ID** - schema `connectorId`: flowConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Enforce Signed Token | `enforcedSignedToken` | `boolean` |  |
| Input Schema | `inputSchema` | `string` | Follow example for JSON schema. |
| Public Key | `pemPublicKey` | `string` | pem public key | 




### Freshdesk

**Connector Display Name**: Freshdesk

**Connector ID** - schema `connectorId`: connectorFreshdesk

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Freshdesk API Key | `apiKey` | `string` | Make sure that the "APIkey:X" is Base64-encoded before pasting into the text field. |
| Freshdesk Base URL (or Domain) | `baseURL` | `` | The <tenant>.freshdesk.com URL or custom domain |
| Freshdesk API Version | `version` | `string` | The current Freshdesk API Version | 




### Freshservice

**Connector Display Name**: Freshservice

**Connector ID** - schema `connectorId`: connectorFreshservice

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | Your Freshservice API key. |
| Domain | `domain` | `string` | Your Freshservice domain. Example: https://domain.freshservice.com/ | 




### Function

**Connector Display Name**: Function

**Connector ID** - schema `connectorId`: functionsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### GBG

**Connector Display Name**: GBG

**Connector ID** - schema `connectorId`: gbgConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| GBG Password | `password` | `string` |  |
| Request URL | `requestUrl` | `string` |  |
| Soap Action URL | `soapAction` | `string` | SOAP Action is a header required for the soap request |
| GBG Username | `username` | `string` |  | 




### GitHub Login

**Connector Display Name**: GitHub Login

**Connector ID** - schema `connectorId`: githubIdpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### Google Analytics (Universal Analytics)

**Connector Display Name**: Google Analytics (Universal Analytics)

**Connector ID** - schema `connectorId`: connectorGoogleanalyticsUA

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Tracking ID | `trackingID` | `` | The tracking ID / web property ID. The format is UA-XXXX-Y. All collected data is associated by this ID. |
| Version | `version` | `string` | The Protocol version. The current value is '1'. This will only change when there are changes made that are not backwards compatible. | 




### Google Chrome Device Trust

**Connector Display Name**: Google Chrome Device Trust

**Connector ID** - schema `connectorId`: connectorGoogleChromeEnterprise

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Google Login

**Connector Display Name**: Google Login

**Connector ID** - schema `connectorId`: googleConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### Google Workspace Admin

**Connector Display Name**: Google Workspace Admin

**Connector ID** - schema `connectorId`: googleWorkSpaceAdminConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Service Account Email Address | `iss` | `string` | The email address associated with the Google Workspace service, such as "google-workspace-admin@xenon-set-123456.iam.gserviceaccount.com" |
| Private Key | `privateKey` | `` | The private key associated with the public key that you added to the Google Workspace service. |
| Admin Email Address | `sub` | `string` | The administrator's email address. | 




### HTTP

**Connector Display Name**: HTTP

**Connector ID** - schema `connectorId`: httpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Select an OpenID token management connection for signed HTTP responses. | `connectionId` | `string` |  |
| reCAPTCHA v2 Secret Key | `recaptchaSecretKey` | `string` | The Secret Key from reCAPTCHA Admin dashboard. |
| reCAPTCHA v2 Site Key | `recaptchaSiteKey` | `string` | The Site Key from reCAPTCHA Admin dashboard. | 




### HYPR

**Connector Display Name**: HYPR

**Connector ID** - schema `connectorId`: hyprConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Have I Been Pwned

**Connector Display Name**: Have I Been Pwned

**Connector ID** - schema `connectorId`: haveIBeenPwnedConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Have I Been Pwned API Key | `apiKey` | `string` |  |
| API Server URL | `apiUrl` | `string` |  |
|  | `userAgent` | `` |  | 




### Hellō Connector

**Connector Display Name**: Hellō Connector

**Connector ID** - schema `connectorId`: connectorHello

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Hubspot

**Connector Display Name**: Hubspot

**Connector ID** - schema `connectorId`: connectorHubspot

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `bearerToken` | `` | Your unique API key | 




### ID DataWeb

**Connector Display Name**: ID DataWeb

**Connector ID** - schema `connectorId`: idDatawebConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### ID R&D

**Connector Display Name**: ID R&D

**Connector ID** - schema `connectorId`: idranddConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` |  |
| API Server URL | `apiUrl` | `string` |  | 




### ID.me

**Connector Display Name**: ID.me

**Connector ID** - schema `connectorId`: idMeConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### ID.me - Identity Verification

**Connector Display Name**: ID.me - Identity Verification

**Connector ID** - schema `connectorId`: connectorIdMeIdentity

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### IDEMIA

**Connector Display Name**: IDEMIA

**Connector ID** - schema `connectorId`: idemiaConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apikey` | `` |  |
| IDEMIA API base URL | `baseUrl` | `string` | Base Url for IDEMIA API. Can be found in the dashboard documents. | 




### IDI Data

**Connector Display Name**: IDI Data

**Connector ID** - schema `connectorId`: skPeopleIntelligenceConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authorization URL | `authUrl` | `string` |  |
| Client ID | `clientId` | `string` |  |
| Client Secret | `clientSecret` | `string` |  |
| DPPA | `dppa` | `string` |  |
| GLBA | `glba` | `string` |  |
| Search URL | `searchUrl` | `string` |  | 




### IDmission

**Connector Display Name**: IDmission

**Connector ID** - schema `connectorId`: idmissionConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authentication Description | `authDescription` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| Sign On ID | `loginId` | `string` |  |
| Merchant ID | `merchantId` | `string` |  |
| Password | `password` | `string` |  |
| Product ID | `productId` | `string` |  |
| Product Name | `productName` | `string` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  |
| IDmission Server URL | `url` | `string` |  | 




### Image

**Connector Display Name**: Image

**Connector ID** - schema `connectorId`: imageConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Infinipoint

**Connector Display Name**: Infinipoint

**Connector ID** - schema `connectorId`: connectorInfinipoint

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Jamf

**Connector Display Name**: Jamf

**Connector ID** - schema `connectorId`: connectorJamf

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| JAMF Password | `jamfPassword` | `` | Enter Password for token |
| JAMF Username | `jamfUsername` | `` | Enter Username for token |
| Server Name | `serverName` | `` | Enter Server Name for Base URL | 




### Jira

**Connector Display Name**: Jira

**Connector ID** - schema `connectorId`: jiraConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Jira API token | `apiKey` | `string` | You may need to create a token from Jira with your credentials, if you haven't created one. |
| Base Url | `apiUrl` | `string` | Base URL of the Jira instance. |
| Email Address | `email` | `string` | Email used for your Jira account. | 




### Jira Service Desk

**Connector Display Name**: Jira Service Desk

**Connector ID** - schema `connectorId`: connectorJiraServiceDesk

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Bearer Authorization Token for JIRA Service Desk | `JIRAServiceDeskAuth` | `` | Bearer Authorization Token for JIRA Service Desk |
| Raw JSON for creating new JIRA service desk request | `JIRAServiceDeskCreateData` | `` | Raw JSON body to create new JIRA service desk request. Example: {   "requestParticipants": ["qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d69abfa3980ce712caae"   ],   "serviceDeskId": "10",   "requestTypeId": "25",   "requestFieldValues": {     "summary": "Request JSD help via REST",     "description": "I need a new *mouse* for my Mac"   } } |
| JIRA Service Desk URL | `JIRAServiceDeskURL` | `` | URL for JIRA Service Desk. Example: your-domain.atlassian.net |
| Raw JSON for updating JIRA service desk | `JIRAServiceDeskUpdateData` | `` | Raw JSON body to update JIRA service desk request. Example: {"id": "1","additionalComment": {"body": "I have fixed the problem."}} |
| Method | `method` | `` | The HTTP Method. | 




### Jumio

**Connector Display Name**: Jumio

**Connector ID** - schema `connectorId`: jumioConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` |  |
| Authentication Description | `authDescription` | `string` |  |
| Base URL for Authentication | `authUrl` | `string` |  |
| Time Transaction URL Valid (seconds) | `authorizationTokenLifetime` | `number` | default: 1800 (30 minutes). maximum: 5184000 (60 days) |
| HEX Main Color | `baseColor` | `string` | Must be passed with bgColor. |
| HEX Background Color. | `bgColor` | `string` | Must be passed with baseColor. |
| Callback URL | `callbackUrl` | `` |  |
| API Secret | `clientSecret` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Do not show in iFrame | `doNotShowInIframe` | `boolean` | If this is true, user will be redirected to the verification url and then redirected back when complete |
| Document Verification Url | `docVerificationUrl` | `string` |  |
| Custom Header Logo URL | `headerImageUrl` | `string` | Logo must be: landscape (16:9 or 4:3), min. height of 192 pixels, size 8-64 KB. |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| Locale | `locale` | `string` | Renders content in the specified language. |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  | 




### KBA

**Connector Display Name**: KBA

**Connector ID** - schema `connectorId`: kbaConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authentication Description | `authDescription` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Fields List | `formFieldsList` | `array` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  | 




### Kaizen Secure Voiz

**Connector Display Name**: Kaizen Secure Voiz

**Connector ID** - schema `connectorId`: kaizenVoizConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Server URL | `apiUrl` | `string` | example: http://<server_root>/ksvvoiceservice/rest/service |
| Application Name | `applicationName` | `string` |  |
| Authentication Description | `authDescription` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  | 




### Keyri QR Login

**Connector Display Name**: Keyri QR Login

**Connector ID** - schema `connectorId`: connectorKeyri

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### LDAP

**Connector Display Name**: LDAP

**Connector ID** - schema `connectorId`: pingOneLDAPConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | The Client ID of your PingOne Worker application. |
| Client Secret | `clientSecret` | `string` | The Client Secret of your PingOne Worker application. |
| Environment ID | `envId` | `string` | Your PingOne environment ID. |
| Gateway ID | `gatewayId` | `` | Your PingOne LDAP gateway ID. |
| Region | `region` | `string` | The region in which your PingOne environment exists. | 




### LexisNexis

**Connector Display Name**: LexisNexis

**Connector ID** - schema `connectorId`: lexisNexisConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| ACAS Endpoint | `acasEndpoint` | `` | ACAS Endpoint |
| Account ID | `accountId` | `` | Account ID provided by LexisNexis. |
| API Key | `apiKey` | `string` | API Key provided by LexisNexis. |
| API Key | `apiKey2` | `` | API Key provided by LexisNexis. |
| API Base URL | `apiUrl` | `string` | The Base URL for Phone Finder, ID Verification and ThreatMetrix Capability for LexisNexis. |
| API Base URL | `apiUrl2` | `` | The Base URL for OTP Verification, KBA and Document Verification for LexisNexis API. |
| API Base URL | `apiUrl3` | `` | The Base URL for Emailage using LexisNexis. |
| API Base URL | `apiUrl4` | `` | The Base URL for Emailage using LexisNexis. Remember to add the / in the end. |
| Client ID | `clientId` | `string` | Account SID for Emailage provided by LexisNexis. |
| Client Secret | `clientSecret` | `string` | OAuth Secret for Emailage provided by LexisNexis. |
| Javascript CDN URL | `javascriptCdnUrl` | `string` | This script is used for ThreatMetrix Profiling. |
| Organization ID | `orgId` | `` | Organization ID provided by LexisNexis. |
| Organization ID | `orgId2` | `` | Organization ID provided by LexisNexis. |
| Password | `password` | `string` | Account Password provided by LexisNexis for OTP, KBA or Document Verification |
| Password | `trueIdPassword` | `` | Account Password provided by LexisNexis for True ID |
| Username | `trueIdUsername` | `` | Account Username provided by LexisNexis for True ID |
| Username | `username` | `string` | Account Username provided by LexisNexis for OTP, KBA or Document Verification | 




### LinkedIn Login

**Connector Display Name**: LinkedIn Login

**Connector ID** - schema `connectorId`: linkedInConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### Location Policy

**Connector Display Name**: Location Policy

**Connector ID** - schema `connectorId`: locationPolicyConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### MFA Container

**Connector Display Name**: MFA Container

**Connector ID** - schema `connectorId`: mfaContainerConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Mailchimp

**Connector Display Name**: Mailchimp

**Connector ID** - schema `connectorId`: connectorMailchimp

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Transactional API Key | `transactionalApiKey` | `` | The Transactional API Key is used to send data to the transactional API. |
| Transactional API Version | `transactionalApiVersion` | `` | Mailchimp - Transactional API Version | 




### Mailgun

**Connector Display Name**: Mailgun

**Connector ID** - schema `connectorId`: connectorMailgun

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | Mailgun API Key |
| API Version | `apiVersion` | `` | Mailgun API Version |
| Domain | `mailgunDomain` | `` | Name of the desired domain (e.g. mail.mycompany.com) | 




### Melissa Global Address

**Connector Display Name**: Melissa Global Address

**Connector ID** - schema `connectorId`: melissaConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| License Key | `apiKey` | `string` | License Key is the API key that you can retrieve from Melissa Admin Portal | 




### Microsoft Intune

**Connector Display Name**: Microsoft Intune

**Connector ID** - schema `connectorId`: connectorMicrosoftIntune

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | Client ID |
| Client Secret | `clientSecret` | `string` | Client Secret |
| Domain Name | `domainName` | `` | Domain Name |
| Grant Type | `grantType` | `string` | Grant Type |
| Scope | `scope` | `string` | Scope |
| Tenant | `tenant` | `` | Tenant | 




### Microsoft Login

**Connector Display Name**: Microsoft Login

**Connector ID** - schema `connectorId`: microsoftIdpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### Microsoft Teams

**Connector Display Name**: Microsoft Teams

**Connector ID** - schema `connectorId`: microsoftTeamsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### NuData Security

**Connector Display Name**: NuData Security

**Connector ID** - schema `connectorId`: nudataConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Nuance

**Connector Display Name**: Nuance

**Connector ID** - schema `connectorId`: nuanceConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authentication Description | `authDescription` | `string` |  |
| Config Set Name | `configSetName` | `` | The Config Set Name for accessing Nuance API. |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| Passphrase One | `passphrase1` | `` | Passphrase that the user will need to speak for voice sample. |
| Passphrase Two | `passphrase2` | `` | Passphrase that the user will need to speak for voice sample. |
| Passphrase Three | `passphrase3` | `` | Passphrase that the user will need to speak for voice sample. |
| Passphrase Four | `passphrase4` | `` | Passphrase that the user will need to speak for voice sample. |
| Passphrase Five | `passphrase5` | `` | Passphrase that the user will need to speak for voice sample. |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  | 




### OIDC & OAuth IdP

**Connector Display Name**: OIDC & OAuth IdP

**Connector ID** - schema `connectorId`: genericConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### OPSWAT MetaAccess

**Connector Display Name**: OPSWAT MetaAccess

**Connector ID** - schema `connectorId`: connectorOpswat

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth Client Key | `clientID` | `` | Oauth client key for authenticating API calls with MetaAccess. |
| Oauth Client Secret | `clientSecret` | `string` | Oauth client secret for authenticating API calls with MetaAccess. |
| Cross-Domain API Port | `crossDomainApiPort` | `` | MetaAccess Cross-Domain API integration port. |
| MetaAccess Domain | `maDomain` | `` | MetaAccess domain for your environment. | 




### OneTrust

**Connector Display Name**: OneTrust

**Connector ID** - schema `connectorId`: oneTrustConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | Your OneTrust application client ID. |
| Client Secret | `clientSecret` | `string` | Your OneTrust application client secret. | 




### Onfido

**Connector Display Name**: Onfido

**Connector ID** - schema `connectorId`: onfidoConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Android Application Package Name | `androidPackageName` | `string` | Your Android Application's Package Name |
| API Key | `apiKey` | `string` |  |
| Authentication Description | `authDescription` | `string` |  |
| Base URL | `baseUrl` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Customize Steps | `customizeSteps` | `boolean` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| iOS Application Bundle ID | `iOSBundleId` | `string` | Your iOS Application's Bundle ID |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| CSS URL | `javascriptCSSUrl` | `string` |  |
| Javascript CDN URL | `javascriptCdnUrl` | `string` |  |
| Language | `language` | `string` |  |
|  | `referenceStepsList` | `array` |  |
| Referrer URL | `referrerUrl` | `string` |  |
| Retrieve Reports | `retrieveReports` | `boolean` |  |
| Close on Overlay Click | `shouldCloseOnOverlayClick` | `boolean` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| ID Verification Steps | `stepsList` | `boolean` | The Proof of Address document capture is currently a BETA feature, and it cannot be used in conjunction with the document and face steps as part of a single SDK flow. |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  |
| Customize Language | `useLanguage` | `boolean` |  |
| Modal | `useModal` | `boolean` |  |
| OnFido Description | `viewDescriptions` | `string` |  |
| OnFido Title | `viewTitle` | `string` |  | 




### PaloAlto Prisma Connector

**Connector Display Name**: PaloAlto Prisma Connector

**Connector ID** - schema `connectorId`: connectorPaloAltoPrisma

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Prisma Base URL | `baseURL` | `` | Prisma Base URL |
| Prisma - Secret Key | `prismaPassword` | `` | Secret Key |
| Prisma - Access Key | `prismaUsername` | `` | Access Key | 




### PingAccess Administration

**Connector Display Name**: PingAccess Administration

**Connector ID** - schema `connectorId`: connector-oai-pingaccessadministrativeapi

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authenticating Password | `authPassword` | `` | The password for an account that has access to the PingAccess administrative API. |
| Authenticating Username | `authUsername` | `` | The username for an account that has access to the PingAccess administrative API. |
| API URL | `basePath` | `` | The base URL for the PingAccess Administrative API, such as "https://localhost:9000/pa-admin-api/v3". |
| Use SSL Verification | `sslVerification` | `` | When enabled, DaVinci verifies the PingAccess SSL certificate and uses encrypted communication. | 




### PingFederate

**Connector Display Name**: PingFederate

**Connector ID** - schema `connectorId`: pingFederateConnectorV2

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### PingFederate Administration

**Connector Display Name**: PingFederate Administration

**Connector ID** - schema `connectorId`: connector-oai-pfadminapi

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authenticating Password | `authPassword` | `` | The password for an account that has access to the PingFederate administrative API. |
| Authenticating Username | `authUsername` | `` | The username for an account that has access to the PingFederate administrative API. |
| API URL | `basePath` | `` | The base URL for the PingFederate administrative API, such as "https://8.8.4.4:9999/pf-admin-api/v1". |
| Use SSL Verification | `sslVerification` | `` | When enabled, DaVinci verifies the PingFederate SSL certificate and uses encrypted communication. | 




### PingID

**Connector Display Name**: PingID

**Connector ID** - schema `connectorId`: pingIdConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### PingOne

**Connector Display Name**: PingOne

**Connector ID** - schema `connectorId`: pingOneSSOConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | The Client ID of your PingOne Worker application. |
| Client Secret | `clientSecret` | `string` | The Client Secret of your PingOne Worker application. |
| Environment ID | `envId` | `string` | Your PingOne environment ID. |
| Region | `region` | `string` | The region in which your PingOne environment exists. | 




### PingOne Authentication

**Connector Display Name**: PingOne Authentication

**Connector ID** - schema `connectorId`: pingOneAuthenticationConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### PingOne Authorize

**Connector Display Name**: PingOne Authorize

**Connector ID** - schema `connectorId`: pingOneAuthorizeConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | The Client ID of your PingOne Worker application |
| Client Secret | `clientSecret` | `string` | The Client Secret of your PingOne Worker application |
| Endpoint URL | `endpointURL` | `` | The url of the decision endpoint to submit a decision request to | 




### PingOne Forms

**Connector Display Name**: PingOne Forms

**Connector ID** - schema `connectorId`: pingOneFormsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### PingOne MFA

**Connector Display Name**: PingOne MFA

**Connector ID** - schema `connectorId`: pingOneMfaConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | The Client ID of your PingOne Worker application. |
| Client Secret | `clientSecret` | `string` | The Client Secret of your PingOne Worker application. |
| Environment ID | `envId` | `string` | Your PingOne Environment ID. |
| Policy ID | `policyId` | `` | The ID of your PingOne MFA device authentication policy. |
| Region | `region` | `string` | The region in which your PingOne environment exists. | 




### PingOne Notifications

**Connector Display Name**: PingOne Notifications

**Connector ID** - schema `connectorId`: notificationsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | The Client ID of your PingOne Worker application. |
| Client Secret | `clientSecret` | `string` | The Client Secret of your PingOne Worker application. |
| Environment ID | `envId` | `string` | Your PingOne Environment ID. |
| Notification Policy ID | `notificationPolicyId` | `` | A unique identifier for the policy. |
| Region | `region` | `string` | The region in which your PingOne environment exists. | 




### PingOne RADIUS Gateway

**Connector Display Name**: PingOne RADIUS Gateway

**Connector ID** - schema `connectorId`: pingOneIntegrationsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### PingOne Risk

**Connector Display Name**: PingOne Risk

**Connector ID** - schema `connectorId`: pingOneRiskConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | The id for your Application found in Ping's Dashboard |
| Client Secret | `clientSecret` | `string` | Client Secret from your App in Ping's Dashboard |
| Environment ID | `envId` | `string` | Your Environment ID provided by Ping. |
| Region | `region` | `string` | The region your PingOne environment is in. | 




### PingOne Verify

**Connector Display Name**: PingOne Verify

**Connector ID** - schema `connectorId`: pingOneVerifyConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | The Client ID of your PingOne Worker application |
| Client Secret | `clientSecret` | `string` | The Client Secret of your PingOne Worker application |
| Environment ID | `envId` | `string` | Your PingOne Environment ID. |
| Region | `region` | `string` | The region your PingOne environment is in. | 




### Prove

**Connector Display Name**: Prove

**Connector ID** - schema `connectorId`: payfoneConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| App Client ID | `appClientId` | `string` |  |
| Prove Base URL | `baseUrl` | `string` |  |
| Client ID | `clientId` | `string` |  |
| Password | `password` | `string` |  |
| Simulator Mode? | `simulatorMode` | `boolean` |  |
| Simulator Phone Number | `simulatorPhoneNumber` | `string` |  |
| Callback Base URL | `skCallbackBaseUrl` | `string` | Use this url as the callback base URL |
| Username | `username` | `string` |  | 




### Prove International

**Connector Display Name**: Prove International

**Connector ID** - schema `connectorId`: proveConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Prove Base URL | `baseUrl` | `string` |  |
| Prove Client ID | `clientId` | `string` |  |
| Prove Grant Type | `grantType` | `string` |  |
| Prove Password | `password` | `string` |  |
| Prove Username | `username` | `string` |  | 




### RSA

**Connector Display Name**: RSA

**Connector ID** - schema `connectorId`: rsaConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Access ID | `accessId` | `` | RSA Access ID from Administration API key file |
| Access Key | `accessKey` | `` | RSA Access Key from Administration API key file |
| Base URL | `baseUrl` | `string` | Base URL for RSA API that is provided in Administration API key file | 




### SAML

**Connector Display Name**: SAML

**Connector ID** - schema `connectorId`: samlConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### SAML IdP

**Connector Display Name**: SAML IdP

**Connector ID** - schema `connectorId`: samlIdpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| SAML Parameters | `saml` | `array` |  | 




### SMTP Client

**Connector Display Name**: SMTP Client

**Connector ID** - schema `connectorId`: smtpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| SMTP Server/Host | `hostname` | `string` | Example: smtp-relay.gmail.com |
| Client Name | `name` | `string` | Optional hostname of the client, used for identifying to the server, defaults to hostname of the machine |
| Password | `password` | `string` |  |
| SMTP Port | `port` | `number` | Example: 25 |
| Secure Flag? | `secureFlag` | `boolean` |  |
| Username | `username` | `string` |  | 




### SailPoint IdentityNow

**Connector Display Name**: SailPoint IdentityNow

**Connector ID** - schema `connectorId`: connectorIdentityNow

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Client ID | `clientId` | `string` | Client Id for your client found in IdentityNow's Dashboard |
| Client Secret | `clientSecret` | `string` | Client Secret from your client in IdentityNow's Dashboard |
| IdentityNow Tenant | `tenant` | `` | The org name is displayed within the Org Details section of the dashboard | 




### Salesforce

**Connector Display Name**: Salesforce

**Connector ID** - schema `connectorId`: salesforceConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Username | `adminUsername` | `` | The username of your Salesforce administrator account. |
| Consumer Key | `consumerKey` | `` | The consumer key shown on your Salesforce connected app. |
| Domain Name | `domainName` | `` | Your Salesforce domain name, such as "mycompany-dev-ed". |
| Private Key | `privateKey` | `` | The private key that corresponds to the X.509 certificate you added to your Salesforce connected app. | 




### Salesforce Marketing Cloud (BETA)

**Connector Display Name**: Salesforce Marketing Cloud (BETA)

**Connector ID** - schema `connectorId`: connectorSalesforceMarketingCloud

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Salesforce Marketing Cloud URL | `SalesforceMarketingCloudURL` | `` | URL for Salesforce Marketing Cloud. Example: https://YOUR_SUBDOMAIN.rest.marketingcloudapis.com |
| Account ID | `accountId` | `` | Account identifier, or MID, of the target business unit. Use to switch between business units. If you don’t specify account_id, the returned access token is in the context of the business unit that created the integration. |
| Client ID | `clientId` | `string` | Client ID issued when you create the API integration in Installed Packages. |
| Client Secret | `clientSecret` | `string` | Client secret issued when you create the API integration in Installed Packages. |
| Scope | `scope` | `string` | Space-separated list of data-access permissions for your application. | 




### Saviynt Connector Flows

**Connector Display Name**: Saviynt Connector Flows

**Connector ID** - schema `connectorId`: connectorSaviyntFlow

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Saviynt Domain Name | `domainName` | `` | Provide your Saviynt domain name |
| Saviynt Path Name | `path` | `` | Provide your Saviynt path name |
| Saviynt Password | `saviyntPassword` | `` | Provide your Saviynt password |
| Saviynt User Name | `saviyntUserName` | `` | Provide your Saviynt user name | 




### Screen

**Connector Display Name**: Screen

**Connector ID** - schema `connectorId`: screenConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### SecurID

**Connector Display Name**: SecurID

**Connector ID** - schema `connectorId`: securIdConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| SecurID Authentication API REST URL | `apiUrl` | `string` | The URL of your SecurID authentication API, such as "https://company.auth.securid.com" |
| Client Key | `clientKey` | `` | Your SecurID authentication client key, such as "vowc450ahs6nry66vok0pvaizwnfr43ewsqcm7tz". | 




### Securonix

**Connector Display Name**: Securonix

**Connector ID** - schema `connectorId`: connectorSecuronix

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Domain Name | `domainName` | `` | Domain Name |
| Token | `token` | `string` | Token for authentication | 




### Segment

**Connector Display Name**: Segment

**Connector ID** - schema `connectorId`: connectorSegment

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| HTTP Tracking API Version | `version` | `string` | Segment - HTTP Tracking API Version |
| Write Key | `writeKey` | `` | The Write Key is used to send data to a specific workplace | 




### SentiLink

**Connector Display Name**: SentiLink

**Connector ID** - schema `connectorId`: sentilinkConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Account ID | `account` | `` | Account ID of SentiLink |
| API URL | `apiUrl` | `string` |  |
| Javascript CDN URL | `javascriptCdnUrl` | `string` |  |
| Token ID | `token` | `string` | Token ID for SentiLink account. | 




### Shopify Connector

**Connector Display Name**: Shopify Connector

**Connector ID** - schema `connectorId`: connectorShopify

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Admin API Access Token | `accessToken` | `string` | Your store's unique Admin API Access Token that goes into the X-Shopify-Access-Token property. Required scopes when generating Admin API Access Token: 'read_customers' and 'write_customers'. Note any Custom Shopify API calls you intend to use with this connector via Make Custom API Call capability, will have to be added as well. |
| API Version Name | `apiVersion` | `` | The Shopify version name ( ex. 2022-04 ) |
| Store Name | `yourStoreName` | `` | The name of your store as Shopify identifies you ( first text that comes after HTTPS:// ) | 




### Signicat

**Connector Display Name**: Signicat

**Connector ID** - schema `connectorId`: connectorSignicat

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### Slack Login

**Connector Display Name**: Slack Login

**Connector ID** - schema `connectorId`: slackConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### Smarty Address Validator

**Connector Display Name**: Smarty Address Validator

**Connector ID** - schema `connectorId`: connectorSmarty

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Auth ID | `authId` | `` | Smarty Authentication ID (Found on 'API Keys' tab in Smarty tenant) |
| Auth Token | `authToken` | `` | Smarty Authentication Token (Found on 'API Keys' tab in Smarty tenant) |
| License | `license` | `` | Smarty License Value (Found on 'Subscriptions' tab in Smarty tenant) | 




### Socure

**Connector Display Name**: Socure

**Connector ID** - schema `connectorId`: socureConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| ID+ Key | `apiKey` | `string` | ID+ Key is the API key that you can retrieve from Socure Admin Portal |
| API URL | `baseUrl` | `string` | The Socure API URL to target. For a custom value, select Use Custom API URL and enter a value in the Custom API URL field. |
| SDK Key | `sdkKey` | `` | SDK Key that you can retrieve from Socure Admin Portal | 




### Splunk

**Connector Display Name**: Splunk

**Connector ID** - schema `connectorId`: splunkConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Base URL | `apiUrl` | `string` | The Base API URL for Splunk. |
| Port | `port` | `number` | API Server Port. |
| Token | `token` | `string` | Splunk Token to make API requests. | 




### Spotify

**Connector Display Name**: Spotify

**Connector ID** - schema `connectorId`: connectorSpotify

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Oauth2 Parameters | `oauth2` | `array` |  | 




### String Manipulation

**Connector Display Name**: String Manipulation

**Connector ID** - schema `connectorId`: stringsConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Symantec VIP

**Connector Display Name**: Symantec VIP

**Connector ID** - schema `connectorId`: symc

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authentication Description | `authDescription` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| PFX File (Base64 encoded) | `pfxBase64` | `string` |  |
| PFX Password | `pfxPassword` | `string` |  |
| Enable Push Sign On | `pushLoginEnabled` | `boolean` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  | 




### TMT Analysis

**Connector Display Name**: TMT Analysis

**Connector ID** - schema `connectorId`: tmtConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | API Key for TMT Analysis. |
| API Secret | `apiSecret` | `` | API Secret for TMT Analysis. |
| Base URL | `apiUrl` | `string` | The Base API URL for TMT Analysis. | 




### Tableau

**Connector Display Name**: Tableau

**Connector ID** - schema `connectorId`: connectorTableau

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Add Flow Permissions Request Body in XML Format. | `addFlowPermissionsRequestBody` | `` | Add Flow Permissions Request Body in XML Format. Example: <tsRequest><task><flowRun><flow id="flow-id"/><flowRunSpec><flowParameterSpecs><flowParameterSpec parameterId="parameter-id" overrideValue= "overrideValue"/><flowParameterSpecs><flowRunSpec></flowRun></task></tsRequest> |
| Add User to Site Request Body in XML Format. | `addUsertoSiteRequestBody` | `` | Add User to Site Request Body in XML Format. Example: <tsRequest><user name="user-name" siteRole="site-role" authSetting="auth-setting" /></tsRequest> |
| api-version | `apiVersion` | `` | The version of the API to use, such as 3.16. |
| auth-ID | `authId` | `` | The Tableau-Auth sent along with every request |
| XML file format to be used for creating schedule | `createScheduleBody` | `` | This should contain the entire XML. Eg: <tsRequest><schedule name="schedule-name"priority="schedule-priority"type="schedule-type"frequency="schedule-frequency"executionOrder="schedule-execution-order"><frequencyDetails start="start-time" end="end-time"><intervals><interval interval-expression /></intervals></frequencyDetails></schedule></tsRequest> |
| datasource-id | `datasourceId` | `` | The ID of the flow. |
| flow-id | `flowId` | `string` | The flow-id value for the flow you want to add permissions to. |
| group-id | `groupId` | `` | The ID of the group. |
| job-id | `jobId` | `` | The ID of the job. |
| schedule-id | `scheduleId` | `` | The ID of the schedule that you are associating with the data source. |
| server-url | `serverUrl` | `` | The tableau server URL Example: https://www.tableau.com:8030 |
| site-id | `siteId` | `` | The ID of the site that contains the view. |
| task-id | `taskId` | `` | The ID of the extract refresh task. |
| XML file format to be used for updating schedule | `updateScheduleRequestBody` | `` | This should contain the entire XML. Eg: <tsRequest><schedule name="hourly-schedule-1" priority="50" type="Extract" frequency="Hourly" executionOrder="Parallel"><frequencyDetails start="18:30:00" end="23:00:00"><intervals><interval hours="2" /></intervals></frequencyDetails></schedule></tsRequest> |
| Update User Request Body in XML Format. | `updateUserRequestBody` | `` | Update User Request Body in XML Format. <tsRequest><user fullName="new-full-name" email="new-email" password="new-password" siteRole="new-site-role" authSetting="new-auth-setting" /></tsRequest> |
| user-id | `userId` | `string` | The ID of the user to get/give information for. |
| workbook-id | `workbookId` | `` | The ID of the workbook to add to the schedule. | 




### TeleSign

**Connector Display Name**: TeleSign

**Connector ID** - schema `connectorId`: telesignConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Authentication Description | `authDescription` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| Password | `password` | `string` |  |
| Provider Name | `providerName` | `string` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  |
| Username | `username` | `string` |  | 




### TeleSign IVR

**Connector Display Name**: TeleSign IVR

**Connector ID** - schema `connectorId`: telesignIvrConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` |  |
| Customer ID | `customerId` | `string` |  |
| Flow ID | `flowId` | `string` |  |
| TeleSign IVR Phone Number | `telesignIvrPhoneNumber` | `` |  | 




### Teleport

**Connector Display Name**: Teleport

**Connector ID** - schema `connectorId`: nodeConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Token Management

**Connector Display Name**: Token Management

**Connector ID** - schema `connectorId`: skOpenIdConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### TransUnion TLOxp

**Connector Display Name**: TransUnion TLOxp

**Connector ID** - schema `connectorId`: tutloxpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API URL | `apiUrl` | `string` | The URL for your TransUnion API. Unnecessary to change unless you're testing against a demo tenant. |
| DPPA Purpose Code | `dppaCode` | `` | The DPPA code that determines the level of data access in the API. |
| GLB Purpose Code | `glbCode` | `` | The GLB code that determines the level of data access in the API. |
| Password | `password` | `string` | The password for your API User |
| Username | `username` | `string` | The username for your API user. | 




### TransUnion TruValidate

**Connector Display Name**: TransUnion TruValidate

**Connector ID** - schema `connectorId`: transunionConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Base URL | `apiUrl` | `string` | The Base API URL for TransUnion. |
| Password | `docVerificationPassword` | `` | Password for Document Verification, provided by TransUnion |
| Public Key | `docVerificationPublicKey` | `` | Public Key for Document Verification, provided by TransUnion |
| Secret | `docVerificationSecret` | `` | Secret for Document Verification, provided by TransUnion |
| Site ID | `docVerificationSiteId` | `` | Site ID for Document Verification, provided by TransUnion |
| Username | `docVerificationUsername` | `` | Username for Document Verification, provided by TransUnion |
| Password | `idVerificationPassword` | `` | Password for ID Verification, provided by TransUnion |
| Public Key | `idVerificationPublicKey` | `` | Public Key for ID Verification, provided by TransUnion |
| Secret | `idVerificationSecret` | `` | Secret for ID Verification, provided by TransUnion |
| Site ID | `idVerificationSiteId` | `` | Site ID for ID Verification, provided by TransUnion |
| Username | `idVerificationUsername` | `` | Username for ID Verification, provided by TransUnion |
| Password | `kbaPassword` | `` | Password for KBA, provided by TransUnion |
| Public Key | `kbaPublicKey` | `` | Public Key for KBA, provided by TransUnion |
| Secret | `kbaSecret` | `` | Secret for KBA, provided by TransUnion |
| Site ID | `kbaSiteId` | `` | Site ID for KBA, provided by TransUnion |
| Username | `kbaUsername` | `` | Username for KBA, provided by TransUnion |
| Password | `otpPassword` | `` | Password for otp Verification, provided by TransUnion |
| Public Key | `otpPublicKey` | `` | Public Key for otp Verification, provided by TransUnion |
| Secret | `otpSecret` | `` | Secret for otp Verification, provided by TransUnion |
| Site ID | `otpSiteId` | `` | Site ID for otp Verification, provided by TransUnion |
| Username | `otpUsername` | `` | Username for otp Verification, provided by TransUnion | 




### Twilio

**Connector Display Name**: Twilio

**Connector ID** - schema `connectorId`: twilioConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Account Sid | `accountSid` | `` |  |
| Authentication Description | `authDescription` | `string` |  |
| Text Message Template (Authentication) | `authMessageTemplate` | `string` |  |
| Auth Token | `authToken` | `` |  |
| Connector Name | `connectorName` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| Text Message Template (Registration) | `registerMessageTemplate` | `string` |  |
| Sender Phone Number | `senderPhoneNumber` | `` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  | 




### Twitter Login

**Connector Display Name**: Twitter Login

**Connector ID** - schema `connectorId`: twitterIdpConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




### UnifyID

**Connector Display Name**: UnifyID

**Connector ID** - schema `connectorId`: unifyIdConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Account ID | `accountId` | `` |  |
| API Key | `apiKey` | `string` |  |
| Connector Name | `connectorName` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| SDK Token | `sdkToken` | `string` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Tooltip | `toolTip` | `string` |  | 




### User Policy

**Connector Display Name**: User Policy

**Connector ID** - schema `connectorId`: userPolicyConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Expires in the specified number of days | `passwordExpiryInDays` | `number` | Choose 0 for never expire |
| Notify user before password expires | `passwordExpiryNotification` | `boolean` |  |
| Maximum Password Length | `passwordLengthMax` | `number` |  |
| Minimum Password Length | `passwordLengthMin` | `number` |  |
| Number of failed login attempts before account is locked | `passwordLockoutAttempts` | `number` |  |
| Number of unique user passwords associated with a user | `passwordPreviousXPasswords` | `number` | Choose 0 if any previous passwords are allowed. This is not recommended. |
| Require Lowercase Characters | `passwordRequireLowercase` | `boolean` | Should the password contain lowercase characters? |
| Require Numbers | `passwordRequireNumbers` | `boolean` | Should the password contain numbers? |
| Require Special Characters | `passwordRequireSpecial` | `boolean` | Should the password contain special character? |
| Require Uppercase Characters | `passwordRequireUppercase` | `boolean` | Should the password contain uppercase characters? |
| Spaces Accepted | `passwordSpacesOk` | `boolean` | Are spaces allowed in the password? |
| Passwords Feature Enabled? | `passwordsEnabled` | `boolean` |  |
| Temporary password expires in the specified number of days | `temporaryPasswordExpiryInDays` | `number` | If an administrator sets a temporary password, choose how long before it expires. | 




### User Pool

**Connector Display Name**: User Pool

**Connector ID** - schema `connectorId`: skUserPool

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
|  | `customAttributes` | `array` |  | 




### Variable

**Connector Display Name**: Variable

**Connector ID** - schema `connectorId`: variablesConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### Venafi Account Service

**Connector Display Name**: Venafi Account Service

**Connector ID** - schema `connectorId`: connector-oai-venafi

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `authApiKey` | `` | The authentication key to the Venafi as a Service API for Account Service Operations |
| Base Path | `basePath` | `` | The base URL for contacting the API | 




### Vericlouds

**Connector Display Name**: Vericlouds

**Connector ID** - schema `connectorId`: connectorVericlouds

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| apiSecret | `apiSecret` | `` | The API secret assigned by VeriClouds to the customer. The secret is also used for decrypting sensitive data such as leaked passwords. It is important to never share the secret with any 3rd party. |
| apiKey | `apikey` | `` | The API key assigned by VeriClouds to the customer. | 




### Verosint

**Connector Display Name**: Verosint

**Connector ID** - schema `connectorId`: connector443id

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | This is the API key from your Verosint account. Remember, Your API KEY is like a serial number for your policy. If you want to utilize more than one policy, you can generate another API KEY and tailor that to a custom policy. | 




### Webhook

**Connector Display Name**: Webhook

**Connector ID** - schema `connectorId`: webhookConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Register URLs | `urls` | `string` | POST requests will be made to these registered url as selected later. | 




### WhatsApp for Business

**Connector Display Name**: WhatsApp for Business

**Connector ID** - schema `connectorId`: connectorWhatsAppBusiness

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Access Token | `accessToken` | `string` | WhatsApp Access Token |
| App Secret | `appSecret` | `` | WhatsApp App Secret for the application, it is used to verify the webhook signatures. |
| Redirect Webhook URI | `skWebhookUri` | `string` | Use this url as the Webhook URL in the Third Party Integration's configuration |
| Webhook Verify Token | `verifyToken` | `` | Meta webhook verify token |
| Version | `version` | `string` | WhatsApp Graph API Version | 




### WinMagic

**Connector Display Name**: WinMagic

**Connector ID** - schema `connectorId`: connectorWinmagic

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| OpenId Parameters | `openId` | `array` |  | 




### Zendesk

**Connector Display Name**: Zendesk

**Connector ID** - schema `connectorId`: connectorZendesk

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Zendesk API Token | `apiToken` | `` | An Active Zendesk API Token (admin center->Apps&Integrations->Zendesk API) |
| Email of User (username) | `emailUsername` | `` | Email used as 'username' for your Zendesk account |
| Subdomain | `subdomain` | `` | Your Zendesk subdomain (ex. {subdomain}.zendesk.com/api/v2/...) | 




### Zoop.one

**Connector Display Name**: Zoop.one

**Connector ID** - schema `connectorId`: zoopConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Zoop Agency ID | `agencyId` | `string` |  |
| Zoop API Key | `apiKey` | `string` |  |
| Zoop API URL | `apiUrl` | `string` |  | 




### Zscaler ZIA

**Connector Display Name**: Zscaler ZIA

**Connector ID** - schema `connectorId`: connectorZscaler

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Base Path | `basePath` | `` | basePath |
| Base URL | `baseURL` | `` | baseURL |
| Zscaler APIkey | `zscalerAPIkey` | `` | Zscaler APIkey |
| Zscaler Password | `zscalerPassword` | `` | Zscaler Domain Password |
| Zscaler Username | `zscalerUsername` | `` | Zscaler Domain Username | 




### iProov

**Connector Display Name**: iProov

**Connector ID** - schema `connectorId`: iproovConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Allow Landscape | `allowLandscape` | `boolean` |  |
| API Key | `apiKey` | `string` |  |
| Authentication Description | `authDescription` | `string` |  |
| Base URL | `baseUrl` | `string` |  |
| Loading Tint Color | `color1` | `string` | Ex. #000000 |
| Not Ready Tint Color | `color2` | `string` | Ex. #000000 |
| Ready Tint Color | `color3` | `string` | Ex. #000000 |
| Liveness Tint Color | `color4` | `string` | Ex. #000000 |
| Connector Name | `connectorName` | `string` |  |
| Custom Title | `customTitle` | `string` | Specify a custom title to be shown. Defaults to show an iProov-generated message. Set to empty string "" to hide the message entirely.  |
| Description | `description` | `string` |  |
| Credentials Details 1 | `details1` | `string` |  |
| Credentials Details 2 | `details2` | `string` |  |
| Enable Camera Selector | `enableCameraSelector` | `boolean` |  |
| Icon URL | `iconUrl` | `string` |  |
| Icon URL in PNG | `iconUrlPng` | `string` |  |
| CSS URL | `javascriptCSSUrl` | `string` |  |
| Javascript CDN URL | `javascriptCdnUrl` | `string` |  |
| Kiosk Mode | `kioskMode` | `boolean` |  |
| Logo | `logo` | `string` | You can use a custom logo by simply passing a relative link, absolute path or data URI to your logo. If you do not want a logo to show pass the logo attribute as null |
| Password | `password` | `string` |  |
| Secret | `secret` | `` |  |
| Show Countdown | `showCountdown` | `boolean` |  |
| Show Credentials Added On? | `showCredAddedOn` | `boolean` |  |
| Show Credentials Added through ? | `showCredAddedVia` | `boolean` |  |
| Start Screen Title | `startScreenTitle` | `string` |  |
| Title | `title` | `string` |  |
| Tooltip | `toolTip` | `string` |  |
| Username | `username` | `string` |  | 




### iovation

**Connector Display Name**: iovation

**Connector ID** - schema `connectorId`: iovationConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Server URL | `apiUrl` | `string` |  |
| iovation loader Javascript CDN URL | `javascriptCdnUrl` | `string` | iovation loader javascript CDN |
| Sub Key | `subKey` | `string` | This will be an iovation assigned value that tracks requests from your site. This is primarily used for debugging and troubleshooting purposes. |
| Subscriber Account | `subscriberAccount` | `string` |  |
| Subscriber ID | `subscriberId` | `string` |  |
| Subscriber Passcode | `subscriberPasscode` | `string` |  |
| Version | `version` | `string` | This is the version of the script to load. The value should either correspond to a specific version you wish to use, or one of the following aliases to get the latest version of the code: general5 - the latest stable version of the javascript, early5 - the latest available version of the javascript | 




### ipgeolocation.io

**Connector Display Name**: ipgeolocation.io

**Connector ID** - schema `connectorId`: connectorIPGeolocationio

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API key | `apiKey` | `string` | Developer subscription API key | 




### ipregistry

**Connector Display Name**: ipregistry

**Connector ID** - schema `connectorId`: connectorIPregistry

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| API Key | `apiKey` | `string` | API Key used to authenticate to the ipregistry.co API. | 




### ipstack

**Connector Display Name**: ipstack

**Connector ID** - schema `connectorId`: connectorIPStack

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Allow Insecure ipstack Connection? | `allowInsecureIPStackConnection` | `` | The Free IPStack Subscription Plan does not support HTTPS connections. For more information refer to https://ipstack.com/plan. |
| API key | `apiKey` | `string` | The ipstack API key to use the service | 




### neoEYED

**Connector Display Name**: neoEYED

**Connector ID** - schema `connectorId`: neoeyedConnector

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Application Key | `appKey` | `` | Unique key for the application |
| Javascript CDN URL | `javascriptCdnUrl` | `string` | URL of javascript CDN of neoEYED | 




### randomuser.me

**Connector Display Name**: randomuser.me

**Connector ID** - schema `connectorId`: connectorRandomUserMe

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----| 




### tru.ID

**Connector Display Name**: tru.ID

**Connector ID** - schema `connectorId`: connectorTruid

**Properties Table:** 



| Display Name | `name` | `value` Type | Description |
| ---- | ---- | ---- | ----|
| Custom Parameters | `customAuth` | `array` |  | 




