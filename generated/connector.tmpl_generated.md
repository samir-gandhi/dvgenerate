## DaVinci Connection Definitions

Below is a list of all available DaVinci Connections available for use in `davinci_connection` resource. 
If the `value` type of a property is not defined it must be inferred.


### (Demo) PingOne Authorize HTTP request and response

Connector ID (`connector_id` in the resource): `pingauthadapter`

*No properties*


Example:
```hcl
resource "davinci_connection" "pingauthadapter" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingauthadapter"
  name           = "My awesome pingauthadapter"
}
```


### 1Kosmos connector

Connector ID (`connector_id` in the resource): `connector1Kosmos`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connector1Kosmos" {
  environment_id = var.pingone_environment_id

  connector_id   = "connector1Kosmos"
  name           = "My awesome connector1Kosmos"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### AWS Lambda

Connector ID (`connector_id` in the resource): `connectorAWSLambda`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessKeyId` (*Type inferred from the provided value*): Access key to connect to AWS Environment. Console display name: "Access Key Id".
* `region` (string): AWS Region where the Lambda function is created. Console display name: "AWS Region".
* `secretAccessKey` (*Type inferred from the provided value*): Secret Key to access the AWS. Console display name: "AWS Secret Key".


Example:
```hcl
resource "davinci_connection" "connectorAWSLambda" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAWSLambda"
  name           = "My awesome connectorAWSLambda"

  property {
    name  = "accessKeyId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }

  property {
    name  = "secretAccessKey"
    value = # value here
  }
}
```


### AWS Login

Connector ID (`connector_id` in the resource): `awsIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "awsIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "awsIdpConnector"
  name           = "My awesome awsIdpConnector"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### AWS Secrets Manager

Connector ID (`connector_id` in the resource): `connectorAmazonAwsSecretsManager`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessKeyId` (*Type inferred from the provided value*): The AWS Access Key. Console display name: "AWS Access Key".
* `region` (string): The AWS Region. Console display name: "AWS Region".
* `secretAccessKey` (*Type inferred from the provided value*): The AWS Access Secret. Console display name: "AWS Access Secret".


Example:
```hcl
resource "davinci_connection" "connectorAmazonAwsSecretsManager" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAmazonAwsSecretsManager"
  name           = "My awesome connectorAmazonAwsSecretsManager"

  property {
    name  = "accessKeyId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }

  property {
    name  = "secretAccessKey"
    value = # value here
  }
}
```


### AbuseIPDB

Connector ID (`connector_id` in the resource): `connectorAbuseipdb`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key gathered from AbuseIPDB tenant. Console display name: "API Key".


Example:
```hcl
resource "davinci_connection" "connectorAbuseipdb" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAbuseipdb"
  name           = "My awesome connectorAbuseipdb"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### Acuant

Connector ID (`connector_id` in the resource): `connectorAcuant`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorAcuant" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAcuant"
  name           = "My awesome connectorAcuant"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Adobe Marketo

Connector ID (`connector_id` in the resource): `adobemarketoConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Your Adobe Marketo client ID. Console display name: "Client ID".
* `clientSecret` (string): Your Adobe Marketo client secret. Console display name: "Client Secret".
* `endpoint` (*Type inferred from the provided value*): The API endpoint for your Adobe Marketo instance, such as "abc123.mktorest.com/rest". Console display name: "API URL".


Example:
```hcl
resource "davinci_connection" "adobemarketoConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "adobemarketoConnector"
  name           = "My awesome adobemarketoConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "endpoint"
    value = # value here
  }
}
```


### Allthenticate

Connector ID (`connector_id` in the resource): `connectorAllthenticate`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorAllthenticate" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAllthenticate"
  name           = "My awesome connectorAllthenticate"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Amazon DynamoDB

Connector ID (`connector_id` in the resource): `connectorAmazonDynamoDB`

Properties (used in the `property` block in the resource as the `name` parameter):

* `awsAccessKey` (string): Your AWS Access Key. Console display name: "AWS Access Key".
* `awsAccessSecret` (string): Access Secret corresponding with Access Key found in Your Security Credentials. Console display name: "AWS Access Secret".
* `awsRegion` (string): The AWS Region you are using the connector for. Console display name: "AWS Region".


Example:
```hcl
resource "davinci_connection" "connectorAmazonDynamoDB" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAmazonDynamoDB"
  name           = "My awesome connectorAmazonDynamoDB"

  property {
    name  = "awsAccessKey"
    value = # value here
  }

  property {
    name  = "awsAccessSecret"
    value = # value here
  }

  property {
    name  = "awsRegion"
    value = # value here
  }
}
```


### Amazon Simple Email Service

Connector ID (`connector_id` in the resource): `amazonSimpleEmailConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `awsAccessKey` (string):  Console display name: "AWS Access Key".
* `awsAccessSecret` (string):  Console display name: "AWS Access Secret".
* `awsRegion` (string):  Console display name: "AWS Region".
* `from` (*Type inferred from the provided value*): The email address that the message appears to originate from, as registered with your AWS account, such as "support@mycompany.com". Console display name: "From (Default) *".


Example:
```hcl
resource "davinci_connection" "amazonSimpleEmailConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "amazonSimpleEmailConnector"
  name           = "My awesome amazonSimpleEmailConnector"

  property {
    name  = "awsAccessKey"
    value = # value here
  }

  property {
    name  = "awsAccessSecret"
    value = # value here
  }

  property {
    name  = "awsRegion"
    value = # value here
  }

  property {
    name  = "from"
    value = # value here
  }
}
```


### Annotation

Connector ID (`connector_id` in the resource): `annotationConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "annotationConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "annotationConnector"
  name           = "My awesome annotationConnector"
}
```


### Apple Login

Connector ID (`connector_id` in the resource): `appleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "appleConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "appleConnector"
  name           = "My awesome appleConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Argyle

Connector ID (`connector_id` in the resource): `argyleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string):  Console display name: "API Server URL".
* `clientId` (string):  Console display name: "Client ID".
* `clientSecret` (string):  Console display name: "Client Secret".
* `javascriptWebUrl` (*Type inferred from the provided value*): Argyle loader javascript web URL. Console display name: "Argyle Loader Javascript Web URL".
* `pluginKey` (*Type inferred from the provided value*):  Console display name: "Plugin Key".


Example:
```hcl
resource "davinci_connection" "argyleConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "argyleConnector"
  name           = "My awesome argyleConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "javascriptWebUrl"
    value = # value here
  }

  property {
    name  = "pluginKey"
    value = # value here
  }
}
```


### Asignio

Connector ID (`connector_id` in the resource): `connectorAsignio`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorAsignio" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAsignio"
  name           = "My awesome connectorAsignio"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### AuthID

Connector ID (`connector_id` in the resource): `connectorAuthid`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorAuthid" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAuthid"
  name           = "My awesome connectorAuthid"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### AuthenticID

Connector ID (`connector_id` in the resource): `authenticIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountAccessKey` (string): Your Account Access Key provided by AuthenticID . Console display name: "Account Access Key".
* `androidSDKLicenseKey` (string): License key is whitelisted for specific package name. Console display name: "Android SDK Licence Key".
* `apiUrl` (string): AuthenticID REST API URL for sandbox/production environments. Console display name: "REST API URL".
* `baseUrl` (string): AuthenticID API URL for sandbox/production environments. Console display name: "Base URL".
* `clientCertificate` (*Type inferred from the provided value*): Your Client Certificate provided by AuthenticID. Console display name: "Client Certificate".
* `clientKey` (*Type inferred from the provided value*): Your Client Key provided by AuthenticID. Console display name: "Client Key".
* `iOSSDKLicenseKey` (string): License key is whitelisted for specific bundle id. Console display name: "iOS SDK Licence Key".
* `passphrase` (*Type inferred from the provided value*): Your Certificate Passphrase provided by AuthenticID. Console display name: "Certificate Passphrase".
* `secretToken` (string): Your Secret Token provided by AuthenticID. Console display name: "Secret Token".


Example:
```hcl
resource "davinci_connection" "authenticIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "authenticIdConnector"
  name           = "My awesome authenticIdConnector"

  property {
    name  = "accountAccessKey"
    value = # value here
  }

  property {
    name  = "androidSDKLicenseKey"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "clientCertificate"
    value = # value here
  }

  property {
    name  = "clientKey"
    value = # value here
  }

  property {
    name  = "iOSSDKLicenseKey"
    value = # value here
  }

  property {
    name  = "passphrase"
    value = # value here
  }

  property {
    name  = "secretToken"
    value = # value here
  }
}
```


### Authomize Incident Connector

Connector ID (`connector_id` in the resource): `connectorAuthomize`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): The API Key from the Authomize API Tokens creation page. Console display name: "Authomize API Key".


Example:
```hcl
resource "davinci_connection" "connectorAuthomize" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorAuthomize"
  name           = "My awesome connectorAuthomize"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### Azure AD User Management

Connector ID (`connector_id` in the resource): `azureUserManagementConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string): The Microsoft API URL to target. For a custom value, select Use Custom API URL and enter a value in the Custom API URL field. Console display name: "API URL".
* `customApiUrl` (*Type inferred from the provided value*): The URL for the Microsoft Graph API, such as "https://graph.microsoft.com/v1.0". Console display name: "Custom API URL".
* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "azureUserManagementConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "azureUserManagementConnector"
  name           = "My awesome azureUserManagementConnector"

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "customApiUrl"
    value = # value here
  }

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Badge

Connector ID (`connector_id` in the resource): `connectorBadge`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorBadge" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorBadge"
  name           = "My awesome connectorBadge"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### BambooHR

Connector ID (`connector_id` in the resource): `bambooConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "API Key".
* `baseUrl` (string):  BambooHR Base URL. Console display name: "Base URL".
* `companySubDomain` (string):  Your BambooHR subdomain. Console display name: "Company Sub Domain".
* `flowId` (string): Select ID of the flow to execute when BambooHR sends a webhook. Console display name: "Flow ID".
* `skWebhookUri` (string): Use this url as the Webhook URL in the Third Party Integration's configuration. Console display name: "DaVinci Webhook URL".
* `webhookToken` (string): Create a webhook token and configure it in the bambooHR webhook url. Console display name: "Webhook Token".


Example:
```hcl
resource "davinci_connection" "bambooConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "bambooConnector"
  name           = "My awesome bambooConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "companySubDomain"
    value = # value here
  }

  property {
    name  = "flowId"
    value = # value here
  }

  property {
    name  = "skWebhookUri"
    value = # value here
  }

  property {
    name  = "webhookToken"
    value = # value here
  }
}
```


### Berbix

Connector ID (`connector_id` in the resource): `connectorBerbix`

Properties (used in the `property` block in the resource as the `name` parameter):

* `domainName` (*Type inferred from the provided value*): Provide Berbix domain name. Console display name: "Domain Name".
* `path` (*Type inferred from the provided value*): Provide path of the API. Console display name: "Path".
* `username` (string): Provide your Berbix user name. Console display name: "User Name".


Example:
```hcl
resource "davinci_connection" "connectorBerbix" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorBerbix"
  name           = "My awesome connectorBerbix"

  property {
    name  = "domainName"
    value = # value here
  }

  property {
    name  = "path"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### Beyond Identity

Connector ID (`connector_id` in the resource): `connectorBeyondIdentity`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "connectorBeyondIdentity" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorBeyondIdentity"
  name           = "My awesome connectorBeyondIdentity"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### BeyondTrust - Password Safe

Connector ID (`connector_id` in the resource): `connectorBTps`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key from your Password Safe environment. Console display name: "API Key".
* `apiUser` (*Type inferred from the provided value*): API User from your Password Safe environment. Console display name: "API User".
* `domain` (string): Domain of your Password Safe environment. Console display name: "PasswordSafe Hostname".


Example:
```hcl
resource "davinci_connection" "connectorBTps" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorBTps"
  name           = "My awesome connectorBTps"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiUser"
    value = # value here
  }

  property {
    name  = "domain"
    value = # value here
  }
}
```


### BeyondTrust - Privileged Remote Access

Connector ID (`connector_id` in the resource): `connectorBTpra`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (*Type inferred from the provided value*): PRA API Client ID. Console display name: "Client ID".
* `clientSecret` (string): PRA API Client Secret. Console display name: "Client Secret".
* `praAPIurl` (*Type inferred from the provided value*): URL of PRA Appliance. Console display name: "PRA Web API Address".


Example:
```hcl
resource "davinci_connection" "connectorBTpra" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorBTpra"
  name           = "My awesome connectorBTpra"

  property {
    name  = "clientID"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "praAPIurl"
    value = # value here
  }
}
```


### BeyondTrust - Remote Support

Connector ID (`connector_id` in the resource): `connectorBTrs`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (*Type inferred from the provided value*): RS API Client ID. Console display name: "Client ID".
* `clientSecret` (string): RS API Client Secret. Console display name: "Client Secret".
* `rsAPIurl` (*Type inferred from the provided value*): URL of RS Appliance. Console display name: "RS Web API Address".


Example:
```hcl
resource "davinci_connection" "connectorBTrs" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorBTrs"
  name           = "My awesome connectorBTrs"

  property {
    name  = "clientID"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "rsAPIurl"
    value = # value here
  }
}
```


### BioCatch

Connector ID (`connector_id` in the resource): `biocatchConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string):  Console display name: "API Server URL".
* `customerId` (string):  Console display name: "Customer ID".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `sdkToken` (string):  Console display name: "SDK Token".
* `truthApiKey` (string): Fraudulent/Genuine Session Reporting API Key. Console display name: "Truth-mapping API Key".
* `truthApiUrl` (string): Fraudulent/Genuine Session Reporting. Console display name: "Truth-mapping API URL".


Example:
```hcl
resource "davinci_connection" "biocatchConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "biocatchConnector"
  name           = "My awesome biocatchConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "customerId"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }

  property {
    name  = "sdkToken"
    value = # value here
  }

  property {
    name  = "truthApiKey"
    value = # value here
  }

  property {
    name  = "truthApiUrl"
    value = # value here
  }
}
```


### Bitbucket Login

Connector ID (`connector_id` in the resource): `bitbucketIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "bitbucketIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "bitbucketIdpConnector"
  name           = "My awesome bitbucketIdpConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### CASTLE

Connector ID (`connector_id` in the resource): `castleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiSecret` (*Type inferred from the provided value*): Your 32-character Castle API secret, such as “Olc…QBF”. Console display name: "API Secret".


Example:
```hcl
resource "davinci_connection" "castleConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "castleConnector"
  name           = "My awesome castleConnector"

  property {
    name  = "apiSecret"
    value = # value here
  }
}
```


### Challenge

Connector ID (`connector_id` in the resource): `challengeConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "challengeConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "challengeConnector"
  name           = "My awesome challengeConnector"
}
```


### Circle Access

Connector ID (`connector_id` in the resource): `connectorCircleAccess`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appKey` (*Type inferred from the provided value*): App Key. Console display name: "App Key".
* `customAuth` (array):  Console display name: "Custom Parameters".
* `loginUrl` (*Type inferred from the provided value*): The URL of your Circle Access login. Console display name: "Login Url".
* `readKey` (*Type inferred from the provided value*): Read Key. Console display name: "Read Key".
* `returnToUrl` (string): When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application. Console display name: "Application Return To URL".
* `writeKey` (*Type inferred from the provided value*): Write key. Console display name: "Write Key".


Example:
```hcl
resource "davinci_connection" "connectorCircleAccess" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorCircleAccess"
  name           = "My awesome connectorCircleAccess"

  property {
    name  = "appKey"
    value = # value here
  }

  property {
    name  = "customAuth"
    value = # value here
  }

  property {
    name  = "loginUrl"
    value = # value here
  }

  property {
    name  = "readKey"
    value = # value here
  }

  property {
    name  = "returnToUrl"
    value = # value here
  }

  property {
    name  = "writeKey"
    value = # value here
  }
}
```


### Clearbit

Connector ID (`connector_id` in the resource): `connectorClearbit`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Clearbit API Key. Console display name: "API Key".
* `riskApiVersion` (*Type inferred from the provided value*): Clearbit - Risk API Version. Console display name: "Risk API Version".
* `version` (string): Clearbit - Person API Version. Console display name: "Person API Version".


Example:
```hcl
resource "davinci_connection" "connectorClearbit" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorClearbit"
  name           = "My awesome connectorClearbit"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "riskApiVersion"
    value = # value here
  }

  property {
    name  = "version"
    value = # value here
  }
}
```


### Cloudflare

Connector ID (`connector_id` in the resource): `connectorCloudflare`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountId` (*Type inferred from the provided value*): Cloudflare Account ID. Console display name: "Account ID".
* `apiToken` (*Type inferred from the provided value*): Cloudflare API Token. Console display name: "API Token".


Example:
```hcl
resource "davinci_connection" "connectorCloudflare" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorCloudflare"
  name           = "My awesome connectorCloudflare"

  property {
    name  = "accountId"
    value = # value here
  }

  property {
    name  = "apiToken"
    value = # value here
  }
}
```


### Code Snippet

Connector ID (`connector_id` in the resource): `codeSnippetConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `code` (string): Follow example for code. Console display name: "Code Snippet".
* `inputSchema` (string): Follow example for JSON schema. Console display name: "Input Schema".
* `outputSchema` (string): Follow example for JSON schema. Console display name: "Output Schema".


Example:
```hcl
resource "davinci_connection" "codeSnippetConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "codeSnippetConnector"
  name           = "My awesome codeSnippetConnector"

  property {
    name  = "code"
    value = # value here
  }

  property {
    name  = "inputSchema"
    value = # value here
  }

  property {
    name  = "outputSchema"
    value = # value here
  }
}
```


### Comply Advantage

Connector ID (`connector_id` in the resource): `complyAdvatangeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key is the API key that you can retrieve from Comply Advantage Admin Portal. Console display name: "API Key".
* `baseUrl` (string): Comply Advantage API URL for sandbox/production environments. Console display name: "Base URL".


Example:
```hcl
resource "davinci_connection" "complyAdvatangeConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "complyAdvatangeConnector"
  name           = "My awesome complyAdvatangeConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }
}
```


### ConnectID

Connector ID (`connector_id` in the resource): `connectIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectIdConnector"
  name           = "My awesome connectIdConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Cookie

Connector ID (`connector_id` in the resource): `cookieConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `hmacSigningKey` (string): Base64 encoded 256 bit key. Console display name: "HMAC Signing Key".


Example:
```hcl
resource "davinci_connection" "cookieConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "cookieConnector"
  name           = "My awesome cookieConnector"

  property {
    name  = "hmacSigningKey"
    value = # value here
  }
}
```


### Credova

Connector ID (`connector_id` in the resource): `credovaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string): Base URL for Credova API. Console display name: "Base URL".
* `password` (string): Password for the Credova Developer Portal. Console display name: "Credova Password".
* `username` (string): Username for the Credova Developer Portal. Console display name: "Credova Username".


Example:
```hcl
resource "davinci_connection" "credovaConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "credovaConnector"
  name           = "My awesome credovaConnector"

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### CrowdStrike

Connector ID (`connector_id` in the resource): `crowdStrikeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (*Type inferred from the provided value*): The base URL of the CrowdStrike environment. Console display name: "CrowdStrike Base URL".
* `clientId` (string): The Client ID of the application in CrowdStrike. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret provided by CrowdStrike. Console display name: "Client Secret".


Example:
```hcl
resource "davinci_connection" "crowdStrikeConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "crowdStrikeConnector"
  name           = "My awesome crowdStrikeConnector"

  property {
    name  = "baseURL"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }
}
```


### Daon IDV

Connector ID (`connector_id` in the resource): `connectorDaonidv`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "connectorDaonidv" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorDaonidv"
  name           = "My awesome connectorDaonidv"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### Daon IdentityX

Connector ID (`connector_id` in the resource): `daonConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The protocol, host and base path to the IdX API. E.g. https://api.identityx-cloud.com/tenant1/IdentityXServices/rest/v1. Console display name: "API Base URL".
* `password` (string): The password of the user to authenticate API calls. Console display name: "Admin Password".
* `username` (string): The userId to authenticate API calls. Console display name: "Admin Username".


Example:
```hcl
resource "davinci_connection" "daonConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "daonConnector"
  name           = "My awesome daonConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### Data Zoo

Connector ID (`connector_id` in the resource): `dataZooConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `password` (string):  Console display name: "Data Zoo Password".
* `username` (string):  Console display name: "Data Zoo Username".


Example:
```hcl
resource "davinci_connection" "dataZooConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "dataZooConnector"
  name           = "My awesome dataZooConnector"

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### Datadog API

Connector ID (`connector_id` in the resource): `connector-oai-datadogapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authApiKey` (*Type inferred from the provided value*): The API key for an account that has access to the Datadog API. Console display name: "Authentication API Key".
* `authApplicationKey` (*Type inferred from the provided value*): The Application key for an account that has access to the Datadog API. Console display name: "Authentication Application Key".
* `basePath` (*Type inferred from the provided value*): The base URL for contacting the Datadog API, such as "https://api.us3.datadoghq.com". Console display name: "API URL".


Example:
```hcl
resource "davinci_connection" "connector-oai-datadogapi" {
  environment_id = var.pingone_environment_id

  connector_id   = "connector-oai-datadogapi"
  name           = "My awesome connector-oai-datadogapi"

  property {
    name  = "authApiKey"
    value = # value here
  }

  property {
    name  = "authApplicationKey"
    value = # value here
  }

  property {
    name  = "basePath"
    value = # value here
  }
}
```


### DeBounce

Connector ID (`connector_id` in the resource): `connectorDeBounce`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): A DeBounce API Key is physically a token/code of 13 random alphanumeric characters. If you need to create an API key, please log in to your DeBounce account and then navigate to the API section. Console display name: "API Key".


Example:
```hcl
resource "davinci_connection" "connectorDeBounce" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorDeBounce"
  name           = "My awesome connectorDeBounce"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### Device Policy

Connector ID (`connector_id` in the resource): `devicePolicyConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "devicePolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "devicePolicyConnector"
  name           = "My awesome devicePolicyConnector"
}
```


### DigiLocker

Connector ID (`connector_id` in the resource): `digilockerConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "digilockerConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "digilockerConnector"
  name           = "My awesome digilockerConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### Digidentity

Connector ID (`connector_id` in the resource): `digidentityConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "digidentityConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "digidentityConnector"
  name           = "My awesome digidentityConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### Duo

Connector ID (`connector_id` in the resource): `duoConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "duoConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "duoConnector"
  name           = "My awesome duoConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Entrust

Connector ID (`connector_id` in the resource): `entrustConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `applicationId` (string): The application ID for the Identity as a Service application. Console display name: "Application ID".
* `serviceDomain` (*Type inferred from the provided value*): The domain of the Entrust service. Format is '<customer>.<region>.trustedauth.com'. For example, 'mycompany.us.trustedauth.com'. Console display name: "Service Domain".


Example:
```hcl
resource "davinci_connection" "entrustConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "entrustConnector"
  name           = "My awesome entrustConnector"

  property {
    name  = "applicationId"
    value = # value here
  }

  property {
    name  = "serviceDomain"
    value = # value here
  }
}
```


### Equifax

Connector ID (`connector_id` in the resource): `equifaxConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string): Base URL for Equifax API. Console display name: "Base URL".
* `clientId` (string): When you Create a New App, Equifax will assign a Client ID per environment for the API Product. Console display name: "Client ID".
* `clientSecret` (string): When you Create a New App, Equifax will assign a Client Secret per environment for the API Product. Console display name: "Client Secret".
* `equifaxSoapApiEnvironment` (string): SOAP API WSDL Environment. Console display name: "SOAP API Environment".
* `memberNumber` (*Type inferred from the provided value*): Unique Identifier of Customer. Please contact Equifax Sales Representative during client onboarding for this value. Console display name: "Member Number".
* `password` (string): Password provided by Equifax for SOAP API. Console display name: "Password for SOAP API".
* `username` (string): Username provided by Equifax for SOAP API. Console display name: "Username for SOAP API".


Example:
```hcl
resource "davinci_connection" "equifaxConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "equifaxConnector"
  name           = "My awesome equifaxConnector"

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "equifaxSoapApiEnvironment"
    value = # value here
  }

  property {
    name  = "memberNumber"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### Error Message

Connector ID (`connector_id` in the resource): `errorConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "errorConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "errorConnector"
  name           = "My awesome errorConnector"
}
```


### Facebook Login

Connector ID (`connector_id` in the resource): `facebookIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "facebookIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "facebookIdpConnector"
  name           = "My awesome facebookIdpConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### Fingerprint JS

Connector ID (`connector_id` in the resource): `fingerprintjsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiToken` (*Type inferred from the provided value*):  Console display name: "Fingerprint Subscription API Token".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `token` (string):  Console display name: "Fingerprint Subscription Browser Token".


Example:
```hcl
resource "davinci_connection" "fingerprintjsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "fingerprintjsConnector"
  name           = "My awesome fingerprintjsConnector"

  property {
    name  = "apiToken"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }

  property {
    name  = "token"
    value = # value here
  }
}
```


### Finicity

Connector ID (`connector_id` in the resource): `finicityConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appKey` (*Type inferred from the provided value*): Finicity App Key from Finicity Developer Portal. Console display name: "Finicity App Key".
* `baseUrl` (string): Base URL for Finicity API. Console display name: "Base URL".
* `partnerId` (*Type inferred from the provided value*): The partner id you can obtain from your Finicity developer dashboard. Console display name: "Partner ID".
* `partnerSecret` (*Type inferred from the provided value*): Partner Secret from Finicity Developer Portal. Console display name: "Partner Secret".


Example:
```hcl
resource "davinci_connection" "finicityConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "finicityConnector"
  name           = "My awesome finicityConnector"

  property {
    name  = "appKey"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "partnerId"
    value = # value here
  }

  property {
    name  = "partnerSecret"
    value = # value here
  }
}
```


### Flow Analytics

Connector ID (`connector_id` in the resource): `analyticsConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "analyticsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "analyticsConnector"
  name           = "My awesome analyticsConnector"
}
```


### Flow Conductor

Connector ID (`connector_id` in the resource): `flowConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `enforcedSignedToken` (boolean):  Console display name: "Enforce Signed Token".
* `inputSchema` (string): Follow example for JSON schema. Console display name: "Input Schema".
* `pemPublicKey` (string): pem public key. Console display name: "Public Key".


Example:
```hcl
resource "davinci_connection" "flowConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "flowConnector"
  name           = "My awesome flowConnector"

  property {
    name  = "enforcedSignedToken"
    value = # value here
  }

  property {
    name  = "inputSchema"
    value = # value here
  }

  property {
    name  = "pemPublicKey"
    value = # value here
  }
}
```


### Freshdesk

Connector ID (`connector_id` in the resource): `connectorFreshdesk`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Make sure that the "APIkey:X" is Base64-encoded before pasting into the text field. Console display name: "Freshdesk API Key".
* `baseURL` (*Type inferred from the provided value*): The <tenant>.freshdesk.com URL or custom domain. Console display name: "Freshdesk Base URL (or Domain)".
* `version` (string): The current Freshdesk API Version. Console display name: "Freshdesk API Version".


Example:
```hcl
resource "davinci_connection" "connectorFreshdesk" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorFreshdesk"
  name           = "My awesome connectorFreshdesk"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "baseURL"
    value = # value here
  }

  property {
    name  = "version"
    value = # value here
  }
}
```


### Freshservice

Connector ID (`connector_id` in the resource): `connectorFreshservice`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Your Freshservice API key. Console display name: "API Key".
* `domain` (string): Your Freshservice domain. Example: https://domain.freshservice.com/. Console display name: "Domain".


Example:
```hcl
resource "davinci_connection" "connectorFreshservice" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorFreshservice"
  name           = "My awesome connectorFreshservice"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "domain"
    value = # value here
  }
}
```


### Functions

Connector ID (`connector_id` in the resource): `functionsConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "functionsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "functionsConnector"
  name           = "My awesome functionsConnector"
}
```


### GBG

Connector ID (`connector_id` in the resource): `gbgConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `password` (string):  Console display name: "GBG Password".
* `requestUrl` (string):  Console display name: "Request URL".
* `soapAction` (string): SOAP Action is a header required for the soap request. Console display name: "Soap Action URL".
* `username` (string):  Console display name: "GBG Username".


Example:
```hcl
resource "davinci_connection" "gbgConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "gbgConnector"
  name           = "My awesome gbgConnector"

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "requestUrl"
    value = # value here
  }

  property {
    name  = "soapAction"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### GitHub Login

Connector ID (`connector_id` in the resource): `githubIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "githubIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "githubIdpConnector"
  name           = "My awesome githubIdpConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### Google Analytics (Universal Analytics)

Connector ID (`connector_id` in the resource): `connectorGoogleanalyticsUA`

Properties (used in the `property` block in the resource as the `name` parameter):

* `trackingID` (*Type inferred from the provided value*): The tracking ID / web property ID. The format is UA-XXXX-Y. All collected data is associated by this ID. Console display name: "Tracking ID".
* `version` (string): The Protocol version. The current value is '1'. This will only change when there are changes made that are not backwards compatible. Console display name: "Version".


Example:
```hcl
resource "davinci_connection" "connectorGoogleanalyticsUA" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorGoogleanalyticsUA"
  name           = "My awesome connectorGoogleanalyticsUA"

  property {
    name  = "trackingID"
    value = # value here
  }

  property {
    name  = "version"
    value = # value here
  }
}
```


### Google Chrome Enterprise Device Trust

Connector ID (`connector_id` in the resource): `connectorGoogleChromeEnterprise`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorGoogleChromeEnterprise" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorGoogleChromeEnterprise"
  name           = "My awesome connectorGoogleChromeEnterprise"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Google Login

Connector ID (`connector_id` in the resource): `googleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "googleConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "googleConnector"
  name           = "My awesome googleConnector"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### Google Workspace Admin

Connector ID (`connector_id` in the resource): `googleWorkSpaceAdminConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `iss` (string): The email address associated with the Google Workspace service, such as "google-workspace-admin@xenon-set-123456.iam.gserviceaccount.com". Console display name: "Service Account Email Address".
* `privateKey` (*Type inferred from the provided value*): The private key associated with the public key that you added to the Google Workspace service. Console display name: "Private Key".
* `sub` (string): The administrator's email address. Console display name: "Admin Email Address".


Example:
```hcl
resource "davinci_connection" "googleWorkSpaceAdminConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "googleWorkSpaceAdminConnector"
  name           = "My awesome googleWorkSpaceAdminConnector"

  property {
    name  = "iss"
    value = # value here
  }

  property {
    name  = "privateKey"
    value = # value here
  }

  property {
    name  = "sub"
    value = # value here
  }
}
```


### HTTP

Connector ID (`connector_id` in the resource): `httpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `connectionId` (string):  Console display name: "Select an OpenID token management connection for signed HTTP responses.".
* `recaptchaSecretKey` (string): The Secret Key from reCAPTCHA Admin dashboard. Console display name: "reCAPTCHA v2 Secret Key".
* `recaptchaSiteKey` (string): The Site Key from reCAPTCHA Admin dashboard. Console display name: "reCAPTCHA v2 Site Key".


Example:
```hcl
resource "davinci_connection" "httpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "httpConnector"
  name           = "My awesome httpConnector"

  property {
    name  = "connectionId"
    value = # value here
  }

  property {
    name  = "recaptchaSecretKey"
    value = # value here
  }

  property {
    name  = "recaptchaSiteKey"
    value = # value here
  }
}
```


### HUMAN

Connector ID (`connector_id` in the resource): `humanCompromisedConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appId` (*Type inferred from the provided value*): App ID from your HUMAN Tenant. Console display name: "HUMAN App ID".
* `authToken` (*Type inferred from the provided value*): Auth Token from your HUMAN Tenant. Console display name: "HUMAN Auth Token".


Example:
```hcl
resource "davinci_connection" "humanCompromisedConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "humanCompromisedConnector"
  name           = "My awesome humanCompromisedConnector"

  property {
    name  = "appId"
    value = # value here
  }

  property {
    name  = "authToken"
    value = # value here
  }
}
```


### HUMAN

Connector ID (`connector_id` in the resource): `connectorHuman`

Properties (used in the `property` block in the resource as the `name` parameter):

* `humanAuthenticationToken` (*Type inferred from the provided value*): Bearer Token from HUMAN. Console display name: "HUMAN Authentication Token".
* `humanCustomerID` (*Type inferred from the provided value*): Customer ID from HUMAN. Console display name: "HUMAN Customer ID".
* `humanPolicyName` (*Type inferred from the provided value*): HUMAN mitigation policy name. Console display name: "HUMAN Policy Name".


Example:
```hcl
resource "davinci_connection" "connectorHuman" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorHuman"
  name           = "My awesome connectorHuman"

  property {
    name  = "humanAuthenticationToken"
    value = # value here
  }

  property {
    name  = "humanCustomerID"
    value = # value here
  }

  property {
    name  = "humanPolicyName"
    value = # value here
  }
}
```


### HYPR

Connector ID (`connector_id` in the resource): `hyprConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "hyprConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "hyprConnector"
  name           = "My awesome hyprConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### HYPR Adapt

Connector ID (`connector_id` in the resource): `connectorHyprAdapt`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessToken` (string): Access Token. Console display name: "HYPR Adapt Access Token".


Example:
```hcl
resource "davinci_connection" "connectorHyprAdapt" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorHyprAdapt"
  name           = "My awesome connectorHyprAdapt"

  property {
    name  = "accessToken"
    value = # value here
  }
}
```


### Have I Been Pwned

Connector ID (`connector_id` in the resource): `haveIBeenPwnedConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "Have I Been Pwned API Key".
* `apiUrl` (string):  Console display name: "API Server URL".
* `userAgent` (*Type inferred from the provided value*):  


Example:
```hcl
resource "davinci_connection" "haveIBeenPwnedConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "haveIBeenPwnedConnector"
  name           = "My awesome haveIBeenPwnedConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "userAgent"
    value = # value here
  }
}
```


### Hellō Connector

Connector ID (`connector_id` in the resource): `connectorHello`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorHello" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorHello"
  name           = "My awesome connectorHello"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Hubspot

Connector ID (`connector_id` in the resource): `connectorHubspot`

Properties (used in the `property` block in the resource as the `name` parameter):

* `bearerToken` (*Type inferred from the provided value*): Your unique API key. Console display name: "API Key".


Example:
```hcl
resource "davinci_connection" "connectorHubspot" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorHubspot"
  name           = "My awesome connectorHubspot"

  property {
    name  = "bearerToken"
    value = # value here
  }
}
```


### ID DataWeb

Connector ID (`connector_id` in the resource): `idDatawebConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "idDatawebConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "idDatawebConnector"
  name           = "My awesome idDatawebConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### ID R&D

Connector ID (`connector_id` in the resource): `idranddConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "API Key".
* `apiUrl` (string):  Console display name: "API Server URL".


Example:
```hcl
resource "davinci_connection" "idranddConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "idranddConnector"
  name           = "My awesome idranddConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }
}
```


### ID.me

Connector ID (`connector_id` in the resource): `idMeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "idMeConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "idMeConnector"
  name           = "My awesome idMeConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### ID.me - Identity Verification

Connector ID (`connector_id` in the resource): `connectorIdMeIdentity`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "connectorIdMeIdentity" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorIdMeIdentity"
  name           = "My awesome connectorIdMeIdentity"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### IDEMIA

Connector ID (`connector_id` in the resource): `idemiaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apikey` (*Type inferred from the provided value*):  Console display name: "API Key".
* `baseUrl` (string): Base Url for IDEMIA API. Can be found in the dashboard documents. Console display name: "IDEMIA API base URL".


Example:
```hcl
resource "davinci_connection" "idemiaConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "idemiaConnector"
  name           = "My awesome idemiaConnector"

  property {
    name  = "apikey"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }
}
```


### IDI Data

Connector ID (`connector_id` in the resource): `skPeopleIntelligenceConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authUrl` (string):  Console display name: "Authorization URL".
* `clientId` (string):  Console display name: "Client ID".
* `clientSecret` (string):  Console display name: "Client Secret".
* `dppa` (string):  Console display name: "DPPA".
* `glba` (string):  Console display name: "GLBA".
* `searchUrl` (string):  Console display name: "Search URL".


Example:
```hcl
resource "davinci_connection" "skPeopleIntelligenceConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "skPeopleIntelligenceConnector"
  name           = "My awesome skPeopleIntelligenceConnector"

  property {
    name  = "authUrl"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "dppa"
    value = # value here
  }

  property {
    name  = "glba"
    value = # value here
  }

  property {
    name  = "searchUrl"
    value = # value here
  }
}
```


### IDmelon

Connector ID (`connector_id` in the resource): `connectorIdmelon`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorIdmelon" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorIdmelon"
  name           = "My awesome connectorIdmelon"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### IDmission

Connector ID (`connector_id` in the resource): `idmissionConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `loginId` (string):  Console display name: "Sign On ID".
* `merchantId` (string):  Console display name: "Merchant ID".
* `password` (string):  Console display name: "Password".
* `productId` (string):  Console display name: "Product ID".
* `productName` (string):  Console display name: "Product Name".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `url` (string):  Console display name: "IDmission Server URL".


Example:
```hcl
resource "davinci_connection" "idmissionConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "idmissionConnector"
  name           = "My awesome idmissionConnector"

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "loginId"
    value = # value here
  }

  property {
    name  = "merchantId"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "productId"
    value = # value here
  }

  property {
    name  = "productName"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }

  property {
    name  = "url"
    value = # value here
  }
}
```


### Image

Connector ID (`connector_id` in the resource): `imageConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "imageConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "imageConnector"
  name           = "My awesome imageConnector"
}
```


### Incode

Connector ID (`connector_id` in the resource): `incodeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "incodeConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "incodeConnector"
  name           = "My awesome incodeConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Infinipoint

Connector ID (`connector_id` in the resource): `connectorInfinipoint`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorInfinipoint" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorInfinipoint"
  name           = "My awesome connectorInfinipoint"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Jamf

Connector ID (`connector_id` in the resource): `connectorJamf`

Properties (used in the `property` block in the resource as the `name` parameter):

* `jamfPassword` (*Type inferred from the provided value*): Enter Password for token. Console display name: "JAMF Password".
* `jamfUsername` (*Type inferred from the provided value*): Enter Username for token. Console display name: "JAMF Username".
* `serverName` (*Type inferred from the provided value*): Enter Server Name for Base URL. Console display name: "Server Name".


Example:
```hcl
resource "davinci_connection" "connectorJamf" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorJamf"
  name           = "My awesome connectorJamf"

  property {
    name  = "jamfPassword"
    value = # value here
  }

  property {
    name  = "jamfUsername"
    value = # value here
  }

  property {
    name  = "serverName"
    value = # value here
  }
}
```


### Jira

Connector ID (`connector_id` in the resource): `jiraConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): You may need to create a token from Jira with your credentials, if you haven't created one. Console display name: "Jira API token".
* `apiUrl` (string): Base URL of the Jira instance. Console display name: "Base Url".
* `email` (string): Email used for your Jira account. Console display name: "Email Address".


Example:
```hcl
resource "davinci_connection" "jiraConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "jiraConnector"
  name           = "My awesome jiraConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "email"
    value = # value here
  }
}
```


### Jira Service Desk

Connector ID (`connector_id` in the resource): `connectorJiraServiceDesk`

Properties (used in the `property` block in the resource as the `name` parameter):

* `JIRAServiceDeskAuth` (*Type inferred from the provided value*): Bearer Authorization Token for JIRA Service Desk. Console display name: "Bearer Authorization Token for JIRA Service Desk".
* `JIRAServiceDeskCreateData` (*Type inferred from the provided value*): Raw JSON body to create new JIRA service desk request. Example: {   "requestParticipants": ["qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d69abfa3980ce712caae"   ],   "serviceDeskId": "10",   "requestTypeId": "25",   "requestFieldValues": {     "summary": "Request JSD help via REST",     "description": "I need a new *mouse* for my Mac"   } }. Console display name: "Raw JSON for creating new JIRA service desk request".
* `JIRAServiceDeskURL` (*Type inferred from the provided value*): URL for JIRA Service Desk. Example: your-domain.atlassian.net. Console display name: "JIRA Service Desk URL".
* `JIRAServiceDeskUpdateData` (*Type inferred from the provided value*): Raw JSON body to update JIRA service desk request. Example: {"id": "1","additionalComment": {"body": "I have fixed the problem."}}. Console display name: "Raw JSON for updating JIRA service desk".
* `method` (*Type inferred from the provided value*): The HTTP Method. Console display name: "Method".


Example:
```hcl
resource "davinci_connection" "connectorJiraServiceDesk" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorJiraServiceDesk"
  name           = "My awesome connectorJiraServiceDesk"

  property {
    name  = "JIRAServiceDeskAuth"
    value = # value here
  }

  property {
    name  = "JIRAServiceDeskCreateData"
    value = # value here
  }

  property {
    name  = "JIRAServiceDeskURL"
    value = # value here
  }

  property {
    name  = "JIRAServiceDeskUpdateData"
    value = # value here
  }

  property {
    name  = "method"
    value = # value here
  }
}
```


### Jumio

Connector ID (`connector_id` in the resource): `jumioConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "API Key".
* `authDescription` (string):  Console display name: "Authentication Description".
* `authUrl` (string):  Console display name: "Base URL for Authentication".
* `authorizationTokenLifetime` (number): default: 1800 (30 minutes). maximum: 5184000 (60 days). Console display name: "Time Transaction URL Valid (seconds)".
* `baseColor` (string): Must be passed with bgColor. Console display name: "HEX Main Color".
* `bgColor` (string): Must be passed with baseColor. Console display name: "HEX Background Color.".
* `callbackUrl` (*Type inferred from the provided value*):  Console display name: "Callback URL".
* `clientSecret` (string):  Console display name: "API Secret".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `doNotShowInIframe` (boolean): If this is true, user will be redirected to the verification url and then redirected back when complete. Console display name: "Do not show in iFrame".
* `docVerificationUrl` (string):  Console display name: "Document Verification Url".
* `headerImageUrl` (string): Logo must be: landscape (16:9 or 4:3), min. height of 192 pixels, size 8-64 KB. Console display name: "Custom Header Logo URL".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `locale` (string): Renders content in the specified language. Console display name: "Locale".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```hcl
resource "davinci_connection" "jumioConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "jumioConnector"
  name           = "My awesome jumioConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "authUrl"
    value = # value here
  }

  property {
    name  = "authorizationTokenLifetime"
    value = # value here
  }

  property {
    name  = "baseColor"
    value = # value here
  }

  property {
    name  = "bgColor"
    value = # value here
  }

  property {
    name  = "callbackUrl"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "doNotShowInIframe"
    value = # value here
  }

  property {
    name  = "docVerificationUrl"
    value = # value here
  }

  property {
    name  = "headerImageUrl"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "locale"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }
}
```


### KBA

Connector ID (`connector_id` in the resource): `kbaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `formFieldsList` (array):  Console display name: "Fields List".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```hcl
resource "davinci_connection" "kbaConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "kbaConnector"
  name           = "My awesome kbaConnector"

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "formFieldsList"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }
}
```


### Kaizen Secure Voiz

Connector ID (`connector_id` in the resource): `kaizenVoizConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): example: http://<server_root>/ksvvoiceservice/rest/service. Console display name: "API Server URL".
* `applicationName` (string):  Console display name: "Application Name".
* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```hcl
resource "davinci_connection" "kaizenVoizConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "kaizenVoizConnector"
  name           = "My awesome kaizenVoizConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "applicationName"
    value = # value here
  }

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }
}
```


### Keyless

Connector ID (`connector_id` in the resource): `connectorKeyless`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorKeyless" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorKeyless"
  name           = "My awesome connectorKeyless"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Keyri QR Login

Connector ID (`connector_id` in the resource): `connectorKeyri`

*No properties*


Example:
```hcl
resource "davinci_connection" "connectorKeyri" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorKeyri"
  name           = "My awesome connectorKeyri"
}
```


### LDAP

Connector ID (`connector_id` in the resource): `pingOneLDAPConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne environment ID. Console display name: "Environment ID".
* `gatewayId` (*Type inferred from the provided value*): Your PingOne LDAP gateway ID. Console display name: "Gateway ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "pingOneLDAPConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneLDAPConnector"
  name           = "My awesome pingOneLDAPConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "gatewayId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### LexisNexis

Connector ID (`connector_id` in the resource): `lexisNexisConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `acasEndpoint` (*Type inferred from the provided value*): ACAS Endpoint. Console display name: "ACAS Endpoint".
* `accountId` (*Type inferred from the provided value*): Account ID provided by LexisNexis. Console display name: "Account ID".
* `acuantCameraScript` (*Type inferred from the provided value*): The URL for the Acuant camera script, such as "https://myhost.com/AcuantCamera.min.js". Console display name: "Camera Script".
* `acuantConfigurationScript` (*Type inferred from the provided value*): The URL for the Acuant configuration script, such as "https://myhost.com/configuration.js.". Console display name: "Configuration Script".
* `acuantJavascriptWebSdkScript` (*Type inferred from the provided value*): Tooltip: The URL for the Acuant JavaScript Web SDK, such as "https://myhost.com/AcuantJavascriptWebSdk.min.js". Console display name: "Javascript Web SDK Script".
* `acuantJavascriptWebSdkScriptSource` (*Type inferred from the provided value*): Select the version of the Acuant scripts that you want to use, or use your own copy of the scripts by selecting Use External Acuant Script URLs. Console display name: "Acuant JavaScript Web SDK Script Source".
* `acuantPassiveLivenessScript` (*Type inferred from the provided value*): The URL for the Acuant passive liveness script, such as "https://myhost.com/AcuantPassiveLiveness.min.js". Console display name: "Passive Liveness Script".
* `apiKey` (string): API Key provided by LexisNexis. Console display name: "API Key".
* `apiKey2` (*Type inferred from the provided value*): API Key provided by LexisNexis. Console display name: "API Key".
* `apiUrl` (string): The Base URL for Phone Finder, ID Verification and ThreatMetrix Capability for LexisNexis. Console display name: "API Base URL".
* `apiUrl2` (*Type inferred from the provided value*): The Base URL for OTP Verification, KBA and Document Verification for LexisNexis API. Console display name: "API Base URL".
* `apiUrl3` (*Type inferred from the provided value*): The Base URL for Emailage using LexisNexis. Console display name: "API Base URL".
* `apiUrl4` (*Type inferred from the provided value*): The Base URL for Emailage using LexisNexis. Remember to add the / in the end. Console display name: "API Base URL".
* `clientId` (string): Account SID for Emailage provided by LexisNexis. Console display name: "Client ID".
* `clientSecret` (string): OAuth Secret for Emailage provided by LexisNexis. Console display name: "Client Secret".
* `externalAcuantScriptURLsLabel` (*Type inferred from the provided value*): Provide the location of externally-hosted Acuant scripts. Note: The connector supports version 11.5.0 and later. Console display name: "External Acuant Script URLs".
* `javascriptCdnUrl` (string): This script is used for ThreatMetrix Profiling. Console display name: "Javascript CDN URL".
* `openCvScript` (*Type inferred from the provided value*): The URL for the Acuant Open CV script, such as "https://myhost.com/opencv.min.js". Console display name: "Open CV Script".
* `orgId` (*Type inferred from the provided value*): Organization ID provided by LexisNexis. Console display name: "Organization ID".
* `orgId2` (*Type inferred from the provided value*): Organization ID provided by LexisNexis. Console display name: "Organization ID".
* `password` (string): Account Password provided by LexisNexis for OTP, KBA or Document Verification. Console display name: "Password".
* `trueIdPassword` (*Type inferred from the provided value*): Account Password provided by LexisNexis for True ID. Console display name: "Password".
* `trueIdUsername` (*Type inferred from the provided value*): Account Username provided by LexisNexis for True ID. Console display name: "Username".
* `username` (string): Account Username provided by LexisNexis for OTP, KBA or Document Verification. Console display name: "Username".


Example:
```hcl
resource "davinci_connection" "lexisNexisConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "lexisNexisConnector"
  name           = "My awesome lexisNexisConnector"

  property {
    name  = "acasEndpoint"
    value = # value here
  }

  property {
    name  = "accountId"
    value = # value here
  }

  property {
    name  = "acuantCameraScript"
    value = # value here
  }

  property {
    name  = "acuantConfigurationScript"
    value = # value here
  }

  property {
    name  = "acuantJavascriptWebSdkScript"
    value = # value here
  }

  property {
    name  = "acuantJavascriptWebSdkScriptSource"
    value = # value here
  }

  property {
    name  = "acuantPassiveLivenessScript"
    value = # value here
  }

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiKey2"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "apiUrl2"
    value = # value here
  }

  property {
    name  = "apiUrl3"
    value = # value here
  }

  property {
    name  = "apiUrl4"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "externalAcuantScriptURLsLabel"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }

  property {
    name  = "openCvScript"
    value = # value here
  }

  property {
    name  = "orgId"
    value = # value here
  }

  property {
    name  = "orgId2"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "trueIdPassword"
    value = # value here
  }

  property {
    name  = "trueIdUsername"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### LinkedIn Login

Connector ID (`connector_id` in the resource): `linkedInConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "linkedInConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "linkedInConnector"
  name           = "My awesome linkedInConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### Location Policy

Connector ID (`connector_id` in the resource): `locationPolicyConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "locationPolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "locationPolicyConnector"
  name           = "My awesome locationPolicyConnector"
}
```


### MFA Container

Connector ID (`connector_id` in the resource): `mfaContainerConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "mfaContainerConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "mfaContainerConnector"
  name           = "My awesome mfaContainerConnector"
}
```


### Mailchimp

Connector ID (`connector_id` in the resource): `connectorMailchimp`

Properties (used in the `property` block in the resource as the `name` parameter):

* `transactionalApiKey` (*Type inferred from the provided value*): The Transactional API Key is used to send data to the transactional API. Console display name: "Transactional API Key".
* `transactionalApiVersion` (*Type inferred from the provided value*): Mailchimp - Transactional API Version. Console display name: "Transactional API Version".


Example:
```hcl
resource "davinci_connection" "connectorMailchimp" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorMailchimp"
  name           = "My awesome connectorMailchimp"

  property {
    name  = "transactionalApiKey"
    value = # value here
  }

  property {
    name  = "transactionalApiVersion"
    value = # value here
  }
}
```


### Mailgun

Connector ID (`connector_id` in the resource): `connectorMailgun`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Mailgun API Key. Console display name: "API Key".
* `apiVersion` (*Type inferred from the provided value*): Mailgun API Version. Console display name: "API Version".
* `mailgunDomain` (*Type inferred from the provided value*): Name of the desired domain (e.g. mail.mycompany.com). Console display name: "Domain".


Example:
```hcl
resource "davinci_connection" "connectorMailgun" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorMailgun"
  name           = "My awesome connectorMailgun"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiVersion"
    value = # value here
  }

  property {
    name  = "mailgunDomain"
    value = # value here
  }
}
```


### Melissa Global Address

Connector ID (`connector_id` in the resource): `melissaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): License Key is the API key that you can retrieve from Melissa Admin Portal. Console display name: "License Key".


Example:
```hcl
resource "davinci_connection" "melissaConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "melissaConnector"
  name           = "My awesome melissaConnector"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### Microsoft Intune

Connector ID (`connector_id` in the resource): `connectorMicrosoftIntune`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Client ID. Console display name: "Client ID".
* `clientSecret` (string): Client Secret. Console display name: "Client Secret".
* `domainName` (*Type inferred from the provided value*): Domain Name. Console display name: "Domain Name".
* `grantType` (string): Grant Type. Console display name: "Grant Type".
* `scope` (string): Scope. Console display name: "Scope".
* `tenant` (*Type inferred from the provided value*): Tenant. Console display name: "Tenant".


Example:
```hcl
resource "davinci_connection" "connectorMicrosoftIntune" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorMicrosoftIntune"
  name           = "My awesome connectorMicrosoftIntune"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "domainName"
    value = # value here
  }

  property {
    name  = "grantType"
    value = # value here
  }

  property {
    name  = "scope"
    value = # value here
  }

  property {
    name  = "tenant"
    value = # value here
  }
}
```


### Microsoft Login

Connector ID (`connector_id` in the resource): `microsoftIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "microsoftIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "microsoftIdpConnector"
  name           = "My awesome microsoftIdpConnector"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### Microsoft Teams

Connector ID (`connector_id` in the resource): `microsoftTeamsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "microsoftTeamsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "microsoftTeamsConnector"
  name           = "My awesome microsoftTeamsConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### NuData Security

Connector ID (`connector_id` in the resource): `nudataConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "nudataConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "nudataConnector"
  name           = "My awesome nudataConnector"
}
```


### Nuance

Connector ID (`connector_id` in the resource): `nuanceConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `configSetName` (*Type inferred from the provided value*): The Config Set Name for accessing Nuance API. Console display name: "Config Set Name".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `passphrase1` (*Type inferred from the provided value*): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase One".
* `passphrase2` (*Type inferred from the provided value*): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Two".
* `passphrase3` (*Type inferred from the provided value*): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Three".
* `passphrase4` (*Type inferred from the provided value*): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Four".
* `passphrase5` (*Type inferred from the provided value*): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Five".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```hcl
resource "davinci_connection" "nuanceConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "nuanceConnector"
  name           = "My awesome nuanceConnector"

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "configSetName"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "passphrase1"
    value = # value here
  }

  property {
    name  = "passphrase2"
    value = # value here
  }

  property {
    name  = "passphrase3"
    value = # value here
  }

  property {
    name  = "passphrase4"
    value = # value here
  }

  property {
    name  = "passphrase5"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }
}
```


### OIDC & OAuth IdP

Connector ID (`connector_id` in the resource): `genericConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "genericConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "genericConnector"
  name           = "My awesome genericConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### OPSWAT MetaAccess

Connector ID (`connector_id` in the resource): `connectorOpswat`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (*Type inferred from the provided value*): Oauth client key for authenticating API calls with MetaAccess. Console display name: "Oauth Client Key".
* `clientSecret` (string): Oauth client secret for authenticating API calls with MetaAccess. Console display name: "Oauth Client Secret".
* `crossDomainApiPort` (*Type inferred from the provided value*): MetaAccess Cross-Domain API integration port. Console display name: "Cross-Domain API Port".
* `maDomain` (*Type inferred from the provided value*): MetaAccess domain for your environment. Console display name: "MetaAccess Domain".


Example:
```hcl
resource "davinci_connection" "connectorOpswat" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorOpswat"
  name           = "My awesome connectorOpswat"

  property {
    name  = "clientID"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "crossDomainApiPort"
    value = # value here
  }

  property {
    name  = "maDomain"
    value = # value here
  }
}
```


### OneTrust

Connector ID (`connector_id` in the resource): `oneTrustConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Your OneTrust application client ID. Console display name: "Client ID".
* `clientSecret` (string): Your OneTrust application client secret. Console display name: "Client Secret".


Example:
```hcl
resource "davinci_connection" "oneTrustConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "oneTrustConnector"
  name           = "My awesome oneTrustConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }
}
```


### Onfido

Connector ID (`connector_id` in the resource): `onfidoConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `androidPackageName` (string): Your Android Application's Package Name. Console display name: "Android Application Package Name".
* `apiKey` (string):  Console display name: "API Key".
* `authDescription` (string):  Console display name: "Authentication Description".
* `baseUrl` (string):  Console display name: "Base URL".
* `connectorName` (string):  Console display name: "Connector Name".
* `customizeSteps` (boolean):  Console display name: "Customize Steps".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iOSBundleId` (string): Your iOS Application's Bundle ID. Console display name: "iOS Application Bundle ID".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `javascriptCSSUrl` (string):  Console display name: "CSS URL".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `language` (string):  Console display name: "Language".
* `referenceStepsList` (array):  
* `referrerUrl` (string):  Console display name: "Referrer URL".
* `retrieveReports` (boolean):  Console display name: "Retrieve Reports".
* `shouldCloseOnOverlayClick` (boolean):  Console display name: "Close on Overlay Click".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `stepsList` (boolean): The Proof of Address document capture is currently a BETA feature, and it cannot be used in conjunction with the document and face steps as part of a single SDK flow. Console display name: "ID Verification Steps".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `useLanguage` (boolean):  Console display name: "Customize Language".
* `useModal` (boolean):  Console display name: "Modal".
* `viewDescriptions` (string):  Console display name: "OnFido Description".
* `viewTitle` (string):  Console display name: "OnFido Title".


Example:
```hcl
resource "davinci_connection" "onfidoConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "onfidoConnector"
  name           = "My awesome onfidoConnector"

  property {
    name  = "androidPackageName"
    value = # value here
  }

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "customizeSteps"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iOSBundleId"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "javascriptCSSUrl"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }

  property {
    name  = "language"
    value = # value here
  }

  property {
    name  = "referenceStepsList"
    value = # value here
  }

  property {
    name  = "referrerUrl"
    value = # value here
  }

  property {
    name  = "retrieveReports"
    value = # value here
  }

  property {
    name  = "shouldCloseOnOverlayClick"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "stepsList"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }

  property {
    name  = "useLanguage"
    value = # value here
  }

  property {
    name  = "useModal"
    value = # value here
  }

  property {
    name  = "viewDescriptions"
    value = # value here
  }

  property {
    name  = "viewTitle"
    value = # value here
  }
}
```


### PaloAlto Prisma Connector

Connector ID (`connector_id` in the resource): `connectorPaloAltoPrisma`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (*Type inferred from the provided value*): Prisma Base URL. Console display name: "Prisma Base URL".
* `prismaPassword` (*Type inferred from the provided value*): Secret Key. Console display name: "Prisma - Secret Key".
* `prismaUsername` (*Type inferred from the provided value*): Access Key. Console display name: "Prisma - Access Key".


Example:
```hcl
resource "davinci_connection" "connectorPaloAltoPrisma" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorPaloAltoPrisma"
  name           = "My awesome connectorPaloAltoPrisma"

  property {
    name  = "baseURL"
    value = # value here
  }

  property {
    name  = "prismaPassword"
    value = # value here
  }

  property {
    name  = "prismaUsername"
    value = # value here
  }
}
```


### PingAccess Administration

Connector ID (`connector_id` in the resource): `connector-oai-pingaccessadministrativeapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authPassword` (*Type inferred from the provided value*): The password for an account that has access to the PingAccess administrative API. Console display name: "Authenticating Password".
* `authUsername` (*Type inferred from the provided value*): The username for an account that has access to the PingAccess administrative API. Console display name: "Authenticating Username".
* `basePath` (*Type inferred from the provided value*): The base URL for the PingAccess Administrative API, such as "https://localhost:9000/pa-admin-api/v3". Console display name: "API URL".
* `sslVerification` (*Type inferred from the provided value*): When enabled, DaVinci verifies the PingAccess SSL certificate and uses encrypted communication. Console display name: "Use SSL Verification".


Example:
```hcl
resource "davinci_connection" "connector-oai-pingaccessadministrativeapi" {
  environment_id = var.pingone_environment_id

  connector_id   = "connector-oai-pingaccessadministrativeapi"
  name           = "My awesome connector-oai-pingaccessadministrativeapi"

  property {
    name  = "authPassword"
    value = # value here
  }

  property {
    name  = "authUsername"
    value = # value here
  }

  property {
    name  = "basePath"
    value = # value here
  }

  property {
    name  = "sslVerification"
    value = # value here
  }
}
```


### PingFederate

Connector ID (`connector_id` in the resource): `pingFederateConnectorV2`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "pingFederateConnectorV2" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingFederateConnectorV2"
  name           = "My awesome pingFederateConnectorV2"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### PingFederate Administration

Connector ID (`connector_id` in the resource): `connector-oai-pfadminapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authPassword` (*Type inferred from the provided value*): The password for an account that has access to the PingFederate administrative API. Console display name: "Authenticating Password".
* `authUsername` (*Type inferred from the provided value*): The username for an account that has access to the PingFederate administrative API. Console display name: "Authenticating Username".
* `basePath` (*Type inferred from the provided value*): The base URL for the PingFederate administrative API, such as "https://8.8.4.4:9999/pf-admin-api/v1". Console display name: "API URL".
* `sslVerification` (*Type inferred from the provided value*): When enabled, DaVinci verifies the PingFederate SSL certificate and uses encrypted communication. Console display name: "Use SSL Verification".


Example:
```hcl
resource "davinci_connection" "connector-oai-pfadminapi" {
  environment_id = var.pingone_environment_id

  connector_id   = "connector-oai-pfadminapi"
  name           = "My awesome connector-oai-pfadminapi"

  property {
    name  = "authPassword"
    value = # value here
  }

  property {
    name  = "authUsername"
    value = # value here
  }

  property {
    name  = "basePath"
    value = # value here
  }

  property {
    name  = "sslVerification"
    value = # value here
  }
}
```


### PingID

Connector ID (`connector_id` in the resource): `pingIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "pingIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingIdConnector"
  name           = "My awesome pingIdConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### PingOne

Connector ID (`connector_id` in the resource): `pingOneSSOConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne environment ID. Console display name: "Environment ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "pingOneSSOConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneSSOConnector"
  name           = "My awesome pingOneSSOConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### PingOne Authentication

Connector ID (`connector_id` in the resource): `pingOneAuthenticationConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "pingOneAuthenticationConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneAuthenticationConnector"
  name           = "My awesome pingOneAuthenticationConnector"
}
```


### PingOne Authorize

Connector ID (`connector_id` in the resource): `pingOneAuthorizeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of the PingOne worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of the PingOne worker application. Console display name: "Client Secret".
* `endpointURL` (*Type inferred from the provided value*): The PingOne Authorize decision endpoint or ID to which the connector submits decision requests. Console display name: "Endpoint".


Example:
```hcl
resource "davinci_connection" "pingOneAuthorizeConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneAuthorizeConnector"
  name           = "My awesome pingOneAuthorizeConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "endpointURL"
    value = # value here
  }
}
```


### PingOne Credentials

Connector ID (`connector_id` in the resource): `pingOneCredentialsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `digitalWalletApplicationId` (*Type inferred from the provided value*): Identifier (UUID) associated with the credential digital wallet app. Console display name: "Digital Wallet Application ID".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `region` (string): The region your PingOne environment is in. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "pingOneCredentialsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneCredentialsConnector"
  name           = "My awesome pingOneCredentialsConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "digitalWalletApplicationId"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### PingOne Forms

Connector ID (`connector_id` in the resource): `pingOneFormsConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "pingOneFormsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneFormsConnector"
  name           = "My awesome pingOneFormsConnector"
}
```


### PingOne MFA

Connector ID (`connector_id` in the resource): `pingOneMfaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `policyId` (*Type inferred from the provided value*): The ID of your PingOne MFA device authentication policy. Console display name: "Policy ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "pingOneMfaConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneMfaConnector"
  name           = "My awesome pingOneMfaConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "policyId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### PingOne Notifications

Connector ID (`connector_id` in the resource): `notificationsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `notificationPolicyId` (*Type inferred from the provided value*): A unique identifier for the policy. Console display name: "Notification Policy ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "notificationsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "notificationsConnector"
  name           = "My awesome notificationsConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "notificationPolicyId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### PingOne Protect

Connector ID (`connector_id` in the resource): `pingOneRiskConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The id for your Application found in Ping's Dashboard. Console display name: "Client ID".
* `clientSecret` (string): Client Secret from your App in Ping's Dashboard. Console display name: "Client Secret".
* `envId` (string): Your Environment ID provided by Ping. Console display name: "Environment ID".
* `region` (string): The region your PingOne environment is in. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "pingOneRiskConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneRiskConnector"
  name           = "My awesome pingOneRiskConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### PingOne RADIUS Gateway

Connector ID (`connector_id` in the resource): `pingOneIntegrationsConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "pingOneIntegrationsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneIntegrationsConnector"
  name           = "My awesome pingOneIntegrationsConnector"
}
```


### PingOne Scope Consent

Connector ID (`connector_id` in the resource): `pingOneScopeConsentConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "pingOneScopeConsentConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneScopeConsentConnector"
  name           = "My awesome pingOneScopeConsentConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### PingOne Verify

Connector ID (`connector_id` in the resource): `pingOneVerifyConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `region` (string): The region your PingOne environment is in. Console display name: "Region".


Example:
```hcl
resource "davinci_connection" "pingOneVerifyConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "pingOneVerifyConnector"
  name           = "My awesome pingOneVerifyConnector"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "envId"
    value = # value here
  }

  property {
    name  = "region"
    value = # value here
  }
}
```


### Prove

Connector ID (`connector_id` in the resource): `payfoneConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appClientId` (string):  Console display name: "App Client ID".
* `baseUrl` (string):  Console display name: "Prove Base URL".
* `clientId` (string):  Console display name: "Client ID".
* `password` (string):  Console display name: "Password".
* `simulatorMode` (boolean):  Console display name: "Simulator Mode?".
* `simulatorPhoneNumber` (string):  Console display name: "Simulator Phone Number".
* `skCallbackBaseUrl` (string): Use this url as the callback base URL. Console display name: "Callback Base URL".
* `username` (string):  Console display name: "Username".


Example:
```hcl
resource "davinci_connection" "payfoneConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "payfoneConnector"
  name           = "My awesome payfoneConnector"

  property {
    name  = "appClientId"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "simulatorMode"
    value = # value here
  }

  property {
    name  = "simulatorPhoneNumber"
    value = # value here
  }

  property {
    name  = "skCallbackBaseUrl"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### Prove International

Connector ID (`connector_id` in the resource): `proveConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string):  Console display name: "Prove Base URL".
* `clientId` (string):  Console display name: "Prove Client ID".
* `grantType` (string):  Console display name: "Prove Grant Type".
* `password` (string):  Console display name: "Prove Password".
* `username` (string):  Console display name: "Prove Username".


Example:
```hcl
resource "davinci_connection" "proveConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "proveConnector"
  name           = "My awesome proveConnector"

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "grantType"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### RSA

Connector ID (`connector_id` in the resource): `rsaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessId` (*Type inferred from the provided value*): RSA Access ID from Administration API key file. Console display name: "Access ID".
* `accessKey` (*Type inferred from the provided value*): RSA Access Key from Administration API key file. Console display name: "Access Key".
* `baseUrl` (string): Base URL for RSA API that is provided in Administration API key file. Console display name: "Base URL".


Example:
```hcl
resource "davinci_connection" "rsaConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "rsaConnector"
  name           = "My awesome rsaConnector"

  property {
    name  = "accessId"
    value = # value here
  }

  property {
    name  = "accessKey"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }
}
```


### Red Violet

Connector ID (`connector_id` in the resource): `connectorIdiVERIFIED`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiSecret` (*Type inferred from the provided value*): Please enter your API secret that Red Violet has provided you. Console display name: "API Secret".
* `companyKey` (*Type inferred from the provided value*): Please enter the company key that Red Violet has assigned. Console display name: "Company Key".
* `idiEnv` (string): Please choose which coreIDENTITY environment you would like to query . Console display name: "Environment".
* `siteKey` (*Type inferred from the provided value*): Please enter your site key that Red Violet has provided you. Console display name: "Site Key".
* `uniqueUrl` (*Type inferred from the provided value*): Please enter your unique URL that Red Violet has provided you. Console display name: "Unique URL".


Example:
```hcl
resource "davinci_connection" "connectorIdiVERIFIED" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorIdiVERIFIED"
  name           = "My awesome connectorIdiVERIFIED"

  property {
    name  = "apiSecret"
    value = # value here
  }

  property {
    name  = "companyKey"
    value = # value here
  }

  property {
    name  = "idiEnv"
    value = # value here
  }

  property {
    name  = "siteKey"
    value = # value here
  }

  property {
    name  = "uniqueUrl"
    value = # value here
  }
}
```


### SAML

Connector ID (`connector_id` in the resource): `samlConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "samlConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "samlConnector"
  name           = "My awesome samlConnector"
}
```


### SAML IdP

Connector ID (`connector_id` in the resource): `samlIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `saml` (array):  Console display name: "SAML Parameters".


Example:
```hcl
resource "davinci_connection" "samlIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "samlIdpConnector"
  name           = "My awesome samlIdpConnector"

  property {
    name  = "saml"
    value = # value here
  }
}
```


### SEON

Connector ID (`connector_id` in the resource): `seonConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (*Type inferred from the provided value*): The API URL to target. Console display name: "API Base URL".
* `licenseKey` (*Type inferred from the provided value*): Your SEON license key. For help, see the SEON REST API documentation. Console display name: "License Key".


Example:
```hcl
resource "davinci_connection" "seonConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "seonConnector"
  name           = "My awesome seonConnector"

  property {
    name  = "baseURL"
    value = # value here
  }

  property {
    name  = "licenseKey"
    value = # value here
  }
}
```


### SMTP Client

Connector ID (`connector_id` in the resource): `smtpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `hostname` (string): Example: smtp-relay.gmail.com. Console display name: "SMTP Server/Host".
* `name` (string): Optional hostname of the client, used for identifying to the server, defaults to hostname of the machine. Console display name: "Client Name".
* `password` (string):  Console display name: "Password".
* `port` (number): Example: 25. Console display name: "SMTP Port".
* `secureFlag` (boolean):  Console display name: "Secure Flag?".
* `username` (string):  Console display name: "Username".


Example:
```hcl
resource "davinci_connection" "smtpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "smtpConnector"
  name           = "My awesome smtpConnector"

  property {
    name  = "hostname"
    value = # value here
  }

  property {
    name  = "name"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "port"
    value = # value here
  }

  property {
    name  = "secureFlag"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### SailPoint IdentityNow

Connector ID (`connector_id` in the resource): `connectorIdentityNow`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Client Id for your client found in IdentityNow's Dashboard. Console display name: "Client ID".
* `clientSecret` (string): Client Secret from your client in IdentityNow's Dashboard. Console display name: "Client Secret".
* `tenant` (*Type inferred from the provided value*): The org name is displayed within the Org Details section of the dashboard. Console display name: "IdentityNow Tenant".


Example:
```hcl
resource "davinci_connection" "connectorIdentityNow" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorIdentityNow"
  name           = "My awesome connectorIdentityNow"

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "tenant"
    value = # value here
  }
}
```


### Salesforce

Connector ID (`connector_id` in the resource): `salesforceConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `adminUsername` (*Type inferred from the provided value*): The username of your Salesforce administrator account. Console display name: "Username".
* `consumerKey` (*Type inferred from the provided value*): The consumer key shown on your Salesforce connected app. Console display name: "Consumer Key".
* `domainName` (*Type inferred from the provided value*): Your Salesforce domain name, such as "mycompany-dev-ed". Console display name: "Domain Name".
* `environment` (string): If the environment you specify in the Domain Name field is part of a sandbox organization, select Sandbox. Otherwise, select Production. Console display name: "Environment".
* `privateKey` (*Type inferred from the provided value*): The private key that corresponds to the X.509 certificate you added to your Salesforce connected app. Console display name: "Private Key".


Example:
```hcl
resource "davinci_connection" "salesforceConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "salesforceConnector"
  name           = "My awesome salesforceConnector"

  property {
    name  = "adminUsername"
    value = # value here
  }

  property {
    name  = "consumerKey"
    value = # value here
  }

  property {
    name  = "domainName"
    value = # value here
  }

  property {
    name  = "environment"
    value = # value here
  }

  property {
    name  = "privateKey"
    value = # value here
  }
}
```


### Salesforce Marketing Cloud (BETA)

Connector ID (`connector_id` in the resource): `connectorSalesforceMarketingCloud`

Properties (used in the `property` block in the resource as the `name` parameter):

* `SalesforceMarketingCloudURL` (*Type inferred from the provided value*): URL for Salesforce Marketing Cloud. Example: https://YOUR_SUBDOMAIN.rest.marketingcloudapis.com. Console display name: "Salesforce Marketing Cloud URL".
* `accountId` (*Type inferred from the provided value*): Account identifier, or MID, of the target business unit. Use to switch between business units. If you don’t specify account_id, the returned access token is in the context of the business unit that created the integration. Console display name: "Account ID".
* `clientId` (string): Client ID issued when you create the API integration in Installed Packages. Console display name: "Client ID".
* `clientSecret` (string): Client secret issued when you create the API integration in Installed Packages. Console display name: "Client Secret".
* `scope` (string): Space-separated list of data-access permissions for your application. Console display name: "Scope".


Example:
```hcl
resource "davinci_connection" "connectorSalesforceMarketingCloud" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSalesforceMarketingCloud"
  name           = "My awesome connectorSalesforceMarketingCloud"

  property {
    name  = "SalesforceMarketingCloudURL"
    value = # value here
  }

  property {
    name  = "accountId"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "scope"
    value = # value here
  }
}
```


### Saviynt Connector Flows

Connector ID (`connector_id` in the resource): `connectorSaviyntFlow`

Properties (used in the `property` block in the resource as the `name` parameter):

* `domainName` (*Type inferred from the provided value*): Provide your Saviynt domain name. Console display name: "Saviynt Domain Name".
* `path` (*Type inferred from the provided value*): Provide your Saviynt path name. Console display name: "Saviynt Path Name".
* `saviyntPassword` (*Type inferred from the provided value*): Provide your Saviynt password. Console display name: "Saviynt Password".
* `saviyntUserName` (*Type inferred from the provided value*): Provide your Saviynt user name. Console display name: "Saviynt User Name".


Example:
```hcl
resource "davinci_connection" "connectorSaviyntFlow" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSaviyntFlow"
  name           = "My awesome connectorSaviyntFlow"

  property {
    name  = "domainName"
    value = # value here
  }

  property {
    name  = "path"
    value = # value here
  }

  property {
    name  = "saviyntPassword"
    value = # value here
  }

  property {
    name  = "saviyntUserName"
    value = # value here
  }
}
```


### Screen

Connector ID (`connector_id` in the resource): `screenConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "screenConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "screenConnector"
  name           = "My awesome screenConnector"
}
```


### SecurID

Connector ID (`connector_id` in the resource): `securIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The URL of your SecurID authentication API, such as "https://company.auth.securid.com". Console display name: "SecurID Authentication API REST URL".
* `clientKey` (*Type inferred from the provided value*): Your SecurID authentication client key, such as "vowc450ahs6nry66vok0pvaizwnfr43ewsqcm7tz". Console display name: "Client Key".


Example:
```hcl
resource "davinci_connection" "securIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "securIdConnector"
  name           = "My awesome securIdConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "clientKey"
    value = # value here
  }
}
```


### Securonix

Connector ID (`connector_id` in the resource): `connectorSecuronix`

Properties (used in the `property` block in the resource as the `name` parameter):

* `domainName` (*Type inferred from the provided value*): Domain Name. Console display name: "Domain Name".
* `token` (string): Token for authentication. Console display name: "Token".


Example:
```hcl
resource "davinci_connection" "connectorSecuronix" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSecuronix"
  name           = "My awesome connectorSecuronix"

  property {
    name  = "domainName"
    value = # value here
  }

  property {
    name  = "token"
    value = # value here
  }
}
```


### Segment

Connector ID (`connector_id` in the resource): `connectorSegment`

Properties (used in the `property` block in the resource as the `name` parameter):

* `version` (string): Segment - HTTP Tracking API Version. Console display name: "HTTP Tracking API Version".
* `writeKey` (*Type inferred from the provided value*): The Write Key is used to send data to a specific workplace. Console display name: "Write Key".


Example:
```hcl
resource "davinci_connection" "connectorSegment" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSegment"
  name           = "My awesome connectorSegment"

  property {
    name  = "version"
    value = # value here
  }

  property {
    name  = "writeKey"
    value = # value here
  }
}
```


### SentiLink

Connector ID (`connector_id` in the resource): `sentilinkConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `account` (*Type inferred from the provided value*): Account ID of SentiLink. Console display name: "Account ID".
* `apiUrl` (string):  Console display name: "API URL".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `token` (string): Token ID for SentiLink account. Console display name: "Token ID".


Example:
```hcl
resource "davinci_connection" "sentilinkConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "sentilinkConnector"
  name           = "My awesome sentilinkConnector"

  property {
    name  = "account"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }

  property {
    name  = "token"
    value = # value here
  }
}
```


### ServiceNow

Connector ID (`connector_id` in the resource): `servicenowConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `adminUsername` (*Type inferred from the provided value*): Your ServiceNow administrator username. Console display name: "Username".
* `apiUrl` (string): The API URL to target, such as "https://mycompany.service-now.com". Console display name: "API URL".
* `password` (string): Your ServiceNow administrator password. Console display name: "Password".


Example:
```hcl
resource "davinci_connection" "servicenowConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "servicenowConnector"
  name           = "My awesome servicenowConnector"

  property {
    name  = "adminUsername"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }
}
```


### Shopify Connector

Connector ID (`connector_id` in the resource): `connectorShopify`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessToken` (string): Your store's unique Admin API Access Token that goes into the X-Shopify-Access-Token property. Required scopes when generating Admin API Access Token: 'read_customers' and 'write_customers'. Note any Custom Shopify API calls you intend to use with this connector via Make Custom API Call capability, will have to be added as well. Console display name: "Admin API Access Token".
* `apiVersion` (*Type inferred from the provided value*): The Shopify version name ( ex. 2022-04 ). Console display name: "API Version Name".
* `multipassSecret` (*Type inferred from the provided value*): Shopify Multipass Secret. Console display name: "Multipass Secret".
* `multipassStoreDomain` (*Type inferred from the provided value*): Shopify Multipass Store Domain (yourstorename.myshopify.com). Console display name: "Multipass Store Domain".
* `yourStoreName` (*Type inferred from the provided value*): The name of your store as Shopify identifies you ( first text that comes after HTTPS:// ). Console display name: "Store Name".


Example:
```hcl
resource "davinci_connection" "connectorShopify" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorShopify"
  name           = "My awesome connectorShopify"

  property {
    name  = "accessToken"
    value = # value here
  }

  property {
    name  = "apiVersion"
    value = # value here
  }

  property {
    name  = "multipassSecret"
    value = # value here
  }

  property {
    name  = "multipassStoreDomain"
    value = # value here
  }

  property {
    name  = "yourStoreName"
    value = # value here
  }
}
```


### Signicat

Connector ID (`connector_id` in the resource): `connectorSignicat`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorSignicat" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSignicat"
  name           = "My awesome connectorSignicat"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Singpass Login

Connector ID (`connector_id` in the resource): `singpassLoginConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "singpassLoginConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "singpassLoginConnector"
  name           = "My awesome singpassLoginConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Slack Login

Connector ID (`connector_id` in the resource): `slackConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "slackConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "slackConnector"
  name           = "My awesome slackConnector"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### Smarty Address Validator

Connector ID (`connector_id` in the resource): `connectorSmarty`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authId` (*Type inferred from the provided value*): Smarty Authentication ID (Found on 'API Keys' tab in Smarty tenant). Console display name: "Auth ID".
* `authToken` (*Type inferred from the provided value*): Smarty Authentication Token (Found on 'API Keys' tab in Smarty tenant). Console display name: "Auth Token".
* `license` (*Type inferred from the provided value*): Smarty License Value (Found on 'Subscriptions' tab in Smarty tenant). Console display name: "License".


Example:
```hcl
resource "davinci_connection" "connectorSmarty" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSmarty"
  name           = "My awesome connectorSmarty"

  property {
    name  = "authId"
    value = # value here
  }

  property {
    name  = "authToken"
    value = # value here
  }

  property {
    name  = "license"
    value = # value here
  }
}
```


### Socure

Connector ID (`connector_id` in the resource): `socureConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): ID+ Key is the API key that you can retrieve from Socure Admin Portal. Console display name: "ID+ Key".
* `baseUrl` (string): The Socure API URL to target. For a custom value, select Use Custom API URL and enter a value in the Custom API URL field. Console display name: "API URL".
* `customApiUrl` (*Type inferred from the provided value*): The URL for the Socure API, such as "https://example.socure.com". Console display name: "Custom API URL".
* `sdkKey` (*Type inferred from the provided value*): SDK Key that you can retrieve from Socure Admin Portal. Console display name: "SDK Key".
* `skWebhookUri` (string): Use this url as the Webhook URL in the Third Party Integration's configuration. Console display name: "Webhook URL".


Example:
```hcl
resource "davinci_connection" "socureConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "socureConnector"
  name           = "My awesome socureConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "customApiUrl"
    value = # value here
  }

  property {
    name  = "sdkKey"
    value = # value here
  }

  property {
    name  = "skWebhookUri"
    value = # value here
  }
}
```


### Splunk

Connector ID (`connector_id` in the resource): `splunkConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The Base API URL for Splunk. Console display name: "Base URL".
* `port` (number): API Server Port. Console display name: "Port".
* `token` (string): Splunk Token to make API requests. Console display name: "Token".


Example:
```hcl
resource "davinci_connection" "splunkConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "splunkConnector"
  name           = "My awesome splunkConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "port"
    value = # value here
  }

  property {
    name  = "token"
    value = # value here
  }
}
```


### Spotify

Connector ID (`connector_id` in the resource): `connectorSpotify`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (array):  Console display name: "Oauth2 Parameters".


Example:
```hcl
resource "davinci_connection" "connectorSpotify" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSpotify"
  name           = "My awesome connectorSpotify"

  property {
    name  = "oauth2"
    value = # value here
  }
}
```


### SpyCloud Enterprise Protection

Connector ID (`connector_id` in the resource): `connectorSpycloud`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Contact SpyCloud to acquire an Employee ATO Prevention API Key that will work with DaVinci. Console display name: "SpyCloud Employee ATO Prevention API Key".


Example:
```hcl
resource "davinci_connection" "connectorSpycloud" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSpycloud"
  name           = "My awesome connectorSpycloud"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### String

Connector ID (`connector_id` in the resource): `stringsConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "stringsConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "stringsConnector"
  name           = "My awesome stringsConnector"
}
```


### Svipe

Connector ID (`connector_id` in the resource): `connectorSvipe`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorSvipe" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorSvipe"
  name           = "My awesome connectorSvipe"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Symantec VIP

Connector ID (`connector_id` in the resource): `symc`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `pfxBase64` (string):  Console display name: "PFX File (Base64 encoded)".
* `pfxPassword` (string):  Console display name: "PFX Password".
* `pushLoginEnabled` (boolean):  Console display name: "Enable Push Sign On".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```hcl
resource "davinci_connection" "symc" {
  environment_id = var.pingone_environment_id

  connector_id   = "symc"
  name           = "My awesome symc"

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "pfxBase64"
    value = # value here
  }

  property {
    name  = "pfxPassword"
    value = # value here
  }

  property {
    name  = "pushLoginEnabled"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }
}
```


### TMT Analysis

Connector ID (`connector_id` in the resource): `tmtConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key for TMT Analysis. Console display name: "API Key".
* `apiSecret` (*Type inferred from the provided value*): API Secret for TMT Analysis. Console display name: "API Secret".
* `apiUrl` (string): The Base API URL for TMT Analysis. Console display name: "Base URL".


Example:
```hcl
resource "davinci_connection" "tmtConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "tmtConnector"
  name           = "My awesome tmtConnector"

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiSecret"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }
}
```


### Tableau

Connector ID (`connector_id` in the resource): `connectorTableau`

Properties (used in the `property` block in the resource as the `name` parameter):

* `addFlowPermissionsRequestBody` (*Type inferred from the provided value*): Add Flow Permissions Request Body in XML Format. Example: <tsRequest><task><flowRun><flow id="flow-id"/><flowRunSpec><flowParameterSpecs><flowParameterSpec parameterId="parameter-id" overrideValue= "overrideValue"/><flowParameterSpecs><flowRunSpec></flowRun></task></tsRequest>. Console display name: "Add Flow Permissions Request Body in XML Format.".
* `addUsertoSiteRequestBody` (*Type inferred from the provided value*): Add User to Site Request Body in XML Format. Example: <tsRequest><user name="user-name" siteRole="site-role" authSetting="auth-setting" /></tsRequest>. Console display name: "Add User to Site Request Body in XML Format.".
* `apiVersion` (*Type inferred from the provided value*): The version of the API to use, such as 3.16. Console display name: "api-version".
* `authId` (*Type inferred from the provided value*): The Tableau-Auth sent along with every request. Console display name: "auth-ID".
* `createScheduleBody` (*Type inferred from the provided value*): This should contain the entire XML. Eg: <tsRequest><schedule name="schedule-name"priority="schedule-priority"type="schedule-type"frequency="schedule-frequency"executionOrder="schedule-execution-order"><frequencyDetails start="start-time" end="end-time"><intervals><interval interval-expression /></intervals></frequencyDetails></schedule></tsRequest>. Console display name: "XML file format to be used for creating schedule".
* `datasourceId` (*Type inferred from the provided value*): The ID of the flow. Console display name: "datasource-id".
* `flowId` (string): The flow-id value for the flow you want to add permissions to. Console display name: "flow-id".
* `groupId` (*Type inferred from the provided value*): The ID of the group. Console display name: "group-id".
* `jobId` (*Type inferred from the provided value*): The ID of the job. Console display name: "job-id".
* `scheduleId` (*Type inferred from the provided value*): The ID of the schedule that you are associating with the data source. Console display name: "schedule-id".
* `serverUrl` (*Type inferred from the provided value*): The tableau server URL Example: https://www.tableau.com:8030. Console display name: "server-url".
* `siteId` (*Type inferred from the provided value*): The ID of the site that contains the view. Console display name: "site-id".
* `taskId` (*Type inferred from the provided value*): The ID of the extract refresh task. Console display name: "task-id".
* `updateScheduleRequestBody` (*Type inferred from the provided value*): This should contain the entire XML. Eg: <tsRequest><schedule name="hourly-schedule-1" priority="50" type="Extract" frequency="Hourly" executionOrder="Parallel"><frequencyDetails start="18:30:00" end="23:00:00"><intervals><interval hours="2" /></intervals></frequencyDetails></schedule></tsRequest>. Console display name: "XML file format to be used for updating schedule".
* `updateUserRequestBody` (*Type inferred from the provided value*): Update User Request Body in XML Format. <tsRequest><user fullName="new-full-name" email="new-email" password="new-password" siteRole="new-site-role" authSetting="new-auth-setting" /></tsRequest>. Console display name: "Update User Request Body in XML Format.".
* `userId` (string): The ID of the user to get/give information for. Console display name: "user-id".
* `workbookId` (*Type inferred from the provided value*): The ID of the workbook to add to the schedule. Console display name: "workbook-id".


Example:
```hcl
resource "davinci_connection" "connectorTableau" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorTableau"
  name           = "My awesome connectorTableau"

  property {
    name  = "addFlowPermissionsRequestBody"
    value = # value here
  }

  property {
    name  = "addUsertoSiteRequestBody"
    value = # value here
  }

  property {
    name  = "apiVersion"
    value = # value here
  }

  property {
    name  = "authId"
    value = # value here
  }

  property {
    name  = "createScheduleBody"
    value = # value here
  }

  property {
    name  = "datasourceId"
    value = # value here
  }

  property {
    name  = "flowId"
    value = # value here
  }

  property {
    name  = "groupId"
    value = # value here
  }

  property {
    name  = "jobId"
    value = # value here
  }

  property {
    name  = "scheduleId"
    value = # value here
  }

  property {
    name  = "serverUrl"
    value = # value here
  }

  property {
    name  = "siteId"
    value = # value here
  }

  property {
    name  = "taskId"
    value = # value here
  }

  property {
    name  = "updateScheduleRequestBody"
    value = # value here
  }

  property {
    name  = "updateUserRequestBody"
    value = # value here
  }

  property {
    name  = "userId"
    value = # value here
  }

  property {
    name  = "workbookId"
    value = # value here
  }
}
```


### Teleport

Connector ID (`connector_id` in the resource): `nodeConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "nodeConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "nodeConnector"
  name           = "My awesome nodeConnector"
}
```


### Telesign

Connector ID (`connector_id` in the resource): `telesignConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `password` (string):  Console display name: "Password".
* `providerName` (string):  Console display name: "Provider Name".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `username` (string):  Console display name: "Username".


Example:
```hcl
resource "davinci_connection" "telesignConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "telesignConnector"
  name           = "My awesome telesignConnector"

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "providerName"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### Token Management

Connector ID (`connector_id` in the resource): `skOpenIdConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "skOpenIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "skOpenIdConnector"
  name           = "My awesome skOpenIdConnector"
}
```


### TransUnion TLOxp

Connector ID (`connector_id` in the resource): `tutloxpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The URL for your TransUnion API. Unnecessary to change unless you're testing against a demo tenant. Console display name: "API URL".
* `dppaCode` (*Type inferred from the provided value*): The DPPA code that determines the level of data access in the API. Console display name: "DPPA Purpose Code".
* `glbCode` (*Type inferred from the provided value*): The GLB code that determines the level of data access in the API. Console display name: "GLB Purpose Code".
* `password` (string): The password for your API User. Console display name: "Password".
* `username` (string): The username for your API user. Console display name: "Username".


Example:
```hcl
resource "davinci_connection" "tutloxpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "tutloxpConnector"
  name           = "My awesome tutloxpConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "dppaCode"
    value = # value here
  }

  property {
    name  = "glbCode"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### TransUnion TruValidate

Connector ID (`connector_id` in the resource): `transunionConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The Base API URL for TransUnion. Console display name: "Base URL".
* `docVerificationPassword` (*Type inferred from the provided value*): Password for Document Verification, provided by TransUnion. Console display name: "Password".
* `docVerificationPublicKey` (*Type inferred from the provided value*): Public Key for Document Verification, provided by TransUnion. Console display name: "Public Key".
* `docVerificationSecret` (*Type inferred from the provided value*): Secret for Document Verification, provided by TransUnion. Console display name: "Secret".
* `docVerificationSiteId` (*Type inferred from the provided value*): Site ID for Document Verification, provided by TransUnion. Console display name: "Site ID".
* `docVerificationUsername` (*Type inferred from the provided value*): Username for Document Verification, provided by TransUnion. Console display name: "Username".
* `idVerificationPassword` (*Type inferred from the provided value*): Password for ID Verification, provided by TransUnion. Console display name: "Password".
* `idVerificationPublicKey` (*Type inferred from the provided value*): Public Key for ID Verification, provided by TransUnion. Console display name: "Public Key".
* `idVerificationSecret` (*Type inferred from the provided value*): Secret for ID Verification, provided by TransUnion. Console display name: "Secret".
* `idVerificationSiteId` (*Type inferred from the provided value*): Site ID for ID Verification, provided by TransUnion. Console display name: "Site ID".
* `idVerificationUsername` (*Type inferred from the provided value*): Username for ID Verification, provided by TransUnion. Console display name: "Username".
* `kbaPassword` (*Type inferred from the provided value*): Password for KBA, provided by TransUnion. Console display name: "Password".
* `kbaPublicKey` (*Type inferred from the provided value*): Public Key for KBA, provided by TransUnion. Console display name: "Public Key".
* `kbaSecret` (*Type inferred from the provided value*): Secret for KBA, provided by TransUnion. Console display name: "Secret".
* `kbaSiteId` (*Type inferred from the provided value*): Site ID for KBA, provided by TransUnion. Console display name: "Site ID".
* `kbaUsername` (*Type inferred from the provided value*): Username for KBA, provided by TransUnion. Console display name: "Username".
* `otpPassword` (*Type inferred from the provided value*): Password for otp Verification, provided by TransUnion. Console display name: "Password".
* `otpPublicKey` (*Type inferred from the provided value*): Public Key for otp Verification, provided by TransUnion. Console display name: "Public Key".
* `otpSecret` (*Type inferred from the provided value*): Secret for otp Verification, provided by TransUnion. Console display name: "Secret".
* `otpSiteId` (*Type inferred from the provided value*): Site ID for otp Verification, provided by TransUnion. Console display name: "Site ID".
* `otpUsername` (*Type inferred from the provided value*): Username for otp Verification, provided by TransUnion. Console display name: "Username".


Example:
```hcl
resource "davinci_connection" "transunionConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "transunionConnector"
  name           = "My awesome transunionConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "docVerificationPassword"
    value = # value here
  }

  property {
    name  = "docVerificationPublicKey"
    value = # value here
  }

  property {
    name  = "docVerificationSecret"
    value = # value here
  }

  property {
    name  = "docVerificationSiteId"
    value = # value here
  }

  property {
    name  = "docVerificationUsername"
    value = # value here
  }

  property {
    name  = "idVerificationPassword"
    value = # value here
  }

  property {
    name  = "idVerificationPublicKey"
    value = # value here
  }

  property {
    name  = "idVerificationSecret"
    value = # value here
  }

  property {
    name  = "idVerificationSiteId"
    value = # value here
  }

  property {
    name  = "idVerificationUsername"
    value = # value here
  }

  property {
    name  = "kbaPassword"
    value = # value here
  }

  property {
    name  = "kbaPublicKey"
    value = # value here
  }

  property {
    name  = "kbaSecret"
    value = # value here
  }

  property {
    name  = "kbaSiteId"
    value = # value here
  }

  property {
    name  = "kbaUsername"
    value = # value here
  }

  property {
    name  = "otpPassword"
    value = # value here
  }

  property {
    name  = "otpPublicKey"
    value = # value here
  }

  property {
    name  = "otpSecret"
    value = # value here
  }

  property {
    name  = "otpSiteId"
    value = # value here
  }

  property {
    name  = "otpUsername"
    value = # value here
  }
}
```


### Trulioo

Connector ID (`connector_id` in the resource): `connectorTrulioo`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (*Type inferred from the provided value*): Trulioo Client ID. Console display name: "Client ID".
* `clientSecret` (string): Trulioo Client Secret. Console display name: "Client Secret".


Example:
```hcl
resource "davinci_connection" "connectorTrulioo" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorTrulioo"
  name           = "My awesome connectorTrulioo"

  property {
    name  = "clientID"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }
}
```


### Twilio

Connector ID (`connector_id` in the resource): `twilioConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountSid` (*Type inferred from the provided value*):  Console display name: "Account Sid".
* `authDescription` (string):  Console display name: "Authentication Description".
* `authMessageTemplate` (string):  Console display name: "Text Message Template (Authentication)".
* `authToken` (*Type inferred from the provided value*):  Console display name: "Auth Token".
* `connectorName` (string):  Console display name: "Connector Name".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `registerMessageTemplate` (string):  Console display name: "Text Message Template (Registration)".
* `senderPhoneNumber` (*Type inferred from the provided value*):  Console display name: "Sender Phone Number".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```hcl
resource "davinci_connection" "twilioConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "twilioConnector"
  name           = "My awesome twilioConnector"

  property {
    name  = "accountSid"
    value = # value here
  }

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "authMessageTemplate"
    value = # value here
  }

  property {
    name  = "authToken"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "registerMessageTemplate"
    value = # value here
  }

  property {
    name  = "senderPhoneNumber"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }
}
```


### UnifyID

Connector ID (`connector_id` in the resource): `unifyIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountId` (*Type inferred from the provided value*):  Console display name: "Account ID".
* `apiKey` (string):  Console display name: "API Key".
* `connectorName` (string):  Console display name: "Connector Name".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `sdkToken` (string):  Console display name: "SDK Token".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```hcl
resource "davinci_connection" "unifyIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "unifyIdConnector"
  name           = "My awesome unifyIdConnector"

  property {
    name  = "accountId"
    value = # value here
  }

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "sdkToken"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }
}
```


### User Policy

Connector ID (`connector_id` in the resource): `userPolicyConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `passwordExpiryInDays` (number): Choose 0 for never expire. Console display name: "Expires in the specified number of days".
* `passwordExpiryNotification` (boolean):  Console display name: "Notify user before password expires".
* `passwordLengthMax` (number):  Console display name: "Maximum Password Length".
* `passwordLengthMin` (number):  Console display name: "Minimum Password Length".
* `passwordLockoutAttempts` (number):  Console display name: "Number of failed login attempts before account is locked".
* `passwordPreviousXPasswords` (number): Choose 0 if any previous passwords are allowed. This is not recommended. Console display name: "Number of unique user passwords associated with a user".
* `passwordRequireLowercase` (boolean): Should the password contain lowercase characters?. Console display name: "Require Lowercase Characters".
* `passwordRequireNumbers` (boolean): Should the password contain numbers?. Console display name: "Require Numbers".
* `passwordRequireSpecial` (boolean): Should the password contain special character?. Console display name: "Require Special Characters".
* `passwordRequireUppercase` (boolean): Should the password contain uppercase characters?. Console display name: "Require Uppercase Characters".
* `passwordSpacesOk` (boolean): Are spaces allowed in the password?. Console display name: "Spaces Accepted".
* `passwordsEnabled` (boolean):  Console display name: "Passwords Feature Enabled?".
* `temporaryPasswordExpiryInDays` (number): If an administrator sets a temporary password, choose how long before it expires. Console display name: "Temporary password expires in the specified number of days".


Example:
```hcl
resource "davinci_connection" "userPolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "userPolicyConnector"
  name           = "My awesome userPolicyConnector"

  property {
    name  = "passwordExpiryInDays"
    value = # value here
  }

  property {
    name  = "passwordExpiryNotification"
    value = # value here
  }

  property {
    name  = "passwordLengthMax"
    value = # value here
  }

  property {
    name  = "passwordLengthMin"
    value = # value here
  }

  property {
    name  = "passwordLockoutAttempts"
    value = # value here
  }

  property {
    name  = "passwordPreviousXPasswords"
    value = # value here
  }

  property {
    name  = "passwordRequireLowercase"
    value = # value here
  }

  property {
    name  = "passwordRequireNumbers"
    value = # value here
  }

  property {
    name  = "passwordRequireSpecial"
    value = # value here
  }

  property {
    name  = "passwordRequireUppercase"
    value = # value here
  }

  property {
    name  = "passwordSpacesOk"
    value = # value here
  }

  property {
    name  = "passwordsEnabled"
    value = # value here
  }

  property {
    name  = "temporaryPasswordExpiryInDays"
    value = # value here
  }
}
```


### User Pool

Connector ID (`connector_id` in the resource): `skUserPool`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAttributes` (array):  


Example:
```hcl
resource "davinci_connection" "skUserPool" {
  environment_id = var.pingone_environment_id

  connector_id   = "skUserPool"
  name           = "My awesome skUserPool"

  property {
    name  = "customAttributes"
    value = # value here
  }
}
```


### ValidSoft

Connector ID (`connector_id` in the resource): `connectorValidsoft`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorValidsoft" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorValidsoft"
  name           = "My awesome connectorValidsoft"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Variable

Connector ID (`connector_id` in the resource): `variablesConnector`

*No properties*


Example:
```hcl
resource "davinci_connection" "variablesConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "variablesConnector"
  name           = "My awesome variablesConnector"
}
```


### Vericlouds

Connector ID (`connector_id` in the resource): `connectorVericlouds`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiSecret` (*Type inferred from the provided value*): The API secret assigned by VeriClouds to the customer. The secret is also used for decrypting sensitive data such as leaked passwords. It is important to never share the secret with any 3rd party. Console display name: "apiSecret".
* `apikey` (*Type inferred from the provided value*): The API key assigned by VeriClouds to the customer. Console display name: "apiKey".


Example:
```hcl
resource "davinci_connection" "connectorVericlouds" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorVericlouds"
  name           = "My awesome connectorVericlouds"

  property {
    name  = "apiSecret"
    value = # value here
  }

  property {
    name  = "apikey"
    value = # value here
  }
}
```


### Veriff

Connector ID (`connector_id` in the resource): `veriffConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `access_token` (*Type inferred from the provided value*): The API Key provided by Veriff, such as "323aa031-b4af-4e12-b354-de0da91a2ab0". Console display name: "API Key".
* `baseUrl` (string): The API URL to target, such as “https://stationapi.veriff.com/”. Console display name: "Base URL".
* `password` (string): The Share Secret Key from Veriff to create HMAC signature, such as "20bf4sf0-fbg7-488c-b4f1-d9594lf707bk". Console display name: "Shared Secret Key".


Example:
```hcl
resource "davinci_connection" "veriffConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "veriffConnector"
  name           = "My awesome veriffConnector"

  property {
    name  = "access_token"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }
}
```


### Verosint

Connector ID (`connector_id` in the resource): `connector443id`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): This is the API key from your Verosint account. Remember, Your API KEY is like a serial number for your policy. If you want to utilize more than one policy, you can generate another API KEY and tailor that to a custom policy. Console display name: "API Key".


Example:
```hcl
resource "davinci_connection" "connector443id" {
  environment_id = var.pingone_environment_id

  connector_id   = "connector443id"
  name           = "My awesome connector443id"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### Webhook

Connector ID (`connector_id` in the resource): `webhookConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `urls` (string): POST requests will be made to these registered url as selected later. Console display name: "Register URLs".


Example:
```hcl
resource "davinci_connection" "webhookConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "webhookConnector"
  name           = "My awesome webhookConnector"

  property {
    name  = "urls"
    value = # value here
  }
}
```


### WhatsApp for Business

Connector ID (`connector_id` in the resource): `connectorWhatsAppBusiness`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessToken` (string): WhatsApp Access Token. Console display name: "Access Token".
* `appSecret` (*Type inferred from the provided value*): WhatsApp App Secret for the application, it is used to verify the webhook signatures. Console display name: "App Secret".
* `skWebhookUri` (string): Use this url as the Webhook URL in the Third Party Integration's configuration. Console display name: "Redirect Webhook URI".
* `verifyToken` (*Type inferred from the provided value*): Meta webhook verify token. Console display name: "Webhook Verify Token".
* `version` (string): WhatsApp Graph API Version. Console display name: "Version".


Example:
```hcl
resource "davinci_connection" "connectorWhatsAppBusiness" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorWhatsAppBusiness"
  name           = "My awesome connectorWhatsAppBusiness"

  property {
    name  = "accessToken"
    value = # value here
  }

  property {
    name  = "appSecret"
    value = # value here
  }

  property {
    name  = "skWebhookUri"
    value = # value here
  }

  property {
    name  = "verifyToken"
    value = # value here
  }

  property {
    name  = "version"
    value = # value here
  }
}
```


### WinMagic

Connector ID (`connector_id` in the resource): `connectorWinmagic`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (array):  Console display name: "OpenId Parameters".


Example:
```hcl
resource "davinci_connection" "connectorWinmagic" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorWinmagic"
  name           = "My awesome connectorWinmagic"

  property {
    name  = "openId"
    value = # value here
  }
}
```


### WireWheel

Connector ID (`connector_id` in the resource): `wireWheelConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (*Type inferred from the provided value*): The base API URL of the WireWheel environment. Console display name: "WireWheel Base API URL".
* `clientId` (string): Client ID from WireWheel Channel settings. Console display name: "Client ID".
* `clientSecret` (string): Client Secret from WireWheel Channel settings. Console display name: "Client Secret".
* `issuerId` (*Type inferred from the provided value*): Issuer URL from WireWheel Channel settings. Console display name: "Issuer URL".


Example:
```hcl
resource "davinci_connection" "wireWheelConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "wireWheelConnector"
  name           = "My awesome wireWheelConnector"

  property {
    name  = "baseURL"
    value = # value here
  }

  property {
    name  = "clientId"
    value = # value here
  }

  property {
    name  = "clientSecret"
    value = # value here
  }

  property {
    name  = "issuerId"
    value = # value here
  }
}
```


### X Login

Connector ID (`connector_id` in the resource): `twitterIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "twitterIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "twitterIdpConnector"
  name           = "My awesome twitterIdpConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Yoti

Connector ID (`connector_id` in the resource): `yotiConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "yotiConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "yotiConnector"
  name           = "My awesome yotiConnector"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```


### Zendesk

Connector ID (`connector_id` in the resource): `connectorZendesk`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiToken` (*Type inferred from the provided value*): An Active Zendesk API Token (admin center->Apps&Integrations->Zendesk API). Console display name: "Zendesk API Token".
* `emailUsername` (*Type inferred from the provided value*): Email used as 'username' for your Zendesk account. Console display name: "Email of User (username)".
* `subdomain` (*Type inferred from the provided value*): Your Zendesk subdomain (ex. {subdomain}.zendesk.com/api/v2/...). Console display name: "Subdomain".


Example:
```hcl
resource "davinci_connection" "connectorZendesk" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorZendesk"
  name           = "My awesome connectorZendesk"

  property {
    name  = "apiToken"
    value = # value here
  }

  property {
    name  = "emailUsername"
    value = # value here
  }

  property {
    name  = "subdomain"
    value = # value here
  }
}
```


### Zoop.one

Connector ID (`connector_id` in the resource): `zoopConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `agencyId` (string):  Console display name: "Zoop Agency ID".
* `apiKey` (string):  Console display name: "Zoop API Key".
* `apiUrl` (string):  Console display name: "Zoop API URL".


Example:
```hcl
resource "davinci_connection" "zoopConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "zoopConnector"
  name           = "My awesome zoopConnector"

  property {
    name  = "agencyId"
    value = # value here
  }

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "apiUrl"
    value = # value here
  }
}
```


### Zscaler ZIA

Connector ID (`connector_id` in the resource): `connectorZscaler`

Properties (used in the `property` block in the resource as the `name` parameter):

* `basePath` (*Type inferred from the provided value*): basePath. Console display name: "Base Path".
* `baseURL` (*Type inferred from the provided value*): baseURL. Console display name: "Base URL".
* `zscalerAPIkey` (*Type inferred from the provided value*): Zscaler APIkey. Console display name: "Zscaler APIkey".
* `zscalerPassword` (*Type inferred from the provided value*): Zscaler Domain Password. Console display name: "Zscaler Password".
* `zscalerUsername` (*Type inferred from the provided value*): Zscaler Domain Username. Console display name: "Zscaler Username".


Example:
```hcl
resource "davinci_connection" "connectorZscaler" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorZscaler"
  name           = "My awesome connectorZscaler"

  property {
    name  = "basePath"
    value = # value here
  }

  property {
    name  = "baseURL"
    value = # value here
  }

  property {
    name  = "zscalerAPIkey"
    value = # value here
  }

  property {
    name  = "zscalerPassword"
    value = # value here
  }

  property {
    name  = "zscalerUsername"
    value = # value here
  }
}
```


### iProov

Connector ID (`connector_id` in the resource): `iproovConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `allowLandscape` (boolean):  Console display name: "Allow Landscape".
* `apiKey` (string):  Console display name: "API Key".
* `authDescription` (string):  Console display name: "Authentication Description".
* `baseUrl` (string):  Console display name: "Base URL".
* `color1` (string): Ex. #000000. Console display name: "Loading Tint Color".
* `color2` (string): Ex. #000000. Console display name: "Not Ready Tint Color".
* `color3` (string): Ex. #000000. Console display name: "Ready Tint Color".
* `color4` (string): Ex. #000000. Console display name: "Liveness Tint Color".
* `connectorName` (string):  Console display name: "Connector Name".
* `customTitle` (string): Specify a custom title to be shown. Defaults to show an iProov-generated message. Set to empty string "" to hide the message entirely.  Console display name: "Custom Title".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `enableCameraSelector` (boolean):  Console display name: "Enable Camera Selector".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `javascriptCSSUrl` (string):  Console display name: "CSS URL".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `kioskMode` (boolean):  Console display name: "Kiosk Mode".
* `logo` (string): You can use a custom logo by simply passing a relative link, absolute path or data URI to your logo. If you do not want a logo to show pass the logo attribute as null. Console display name: "Logo".
* `password` (string):  Console display name: "Password".
* `secret` (*Type inferred from the provided value*):  Console display name: "Secret".
* `showCountdown` (boolean):  Console display name: "Show Countdown".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `startScreenTitle` (string):  Console display name: "Start Screen Title".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `username` (string):  Console display name: "Username".


Example:
```hcl
resource "davinci_connection" "iproovConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "iproovConnector"
  name           = "My awesome iproovConnector"

  property {
    name  = "allowLandscape"
    value = # value here
  }

  property {
    name  = "apiKey"
    value = # value here
  }

  property {
    name  = "authDescription"
    value = # value here
  }

  property {
    name  = "baseUrl"
    value = # value here
  }

  property {
    name  = "color1"
    value = # value here
  }

  property {
    name  = "color2"
    value = # value here
  }

  property {
    name  = "color3"
    value = # value here
  }

  property {
    name  = "color4"
    value = # value here
  }

  property {
    name  = "connectorName"
    value = # value here
  }

  property {
    name  = "customTitle"
    value = # value here
  }

  property {
    name  = "description"
    value = # value here
  }

  property {
    name  = "details1"
    value = # value here
  }

  property {
    name  = "details2"
    value = # value here
  }

  property {
    name  = "enableCameraSelector"
    value = # value here
  }

  property {
    name  = "iconUrl"
    value = # value here
  }

  property {
    name  = "iconUrlPng"
    value = # value here
  }

  property {
    name  = "javascriptCSSUrl"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }

  property {
    name  = "kioskMode"
    value = # value here
  }

  property {
    name  = "logo"
    value = # value here
  }

  property {
    name  = "password"
    value = # value here
  }

  property {
    name  = "secret"
    value = # value here
  }

  property {
    name  = "showCountdown"
    value = # value here
  }

  property {
    name  = "showCredAddedOn"
    value = # value here
  }

  property {
    name  = "showCredAddedVia"
    value = # value here
  }

  property {
    name  = "startScreenTitle"
    value = # value here
  }

  property {
    name  = "title"
    value = # value here
  }

  property {
    name  = "toolTip"
    value = # value here
  }

  property {
    name  = "username"
    value = # value here
  }
}
```


### iovation

Connector ID (`connector_id` in the resource): `iovationConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string):  Console display name: "API Server URL".
* `javascriptCdnUrl` (string): iovation loader javascript CDN. Console display name: "iovation loader Javascript CDN URL".
* `subKey` (string): This will be an iovation assigned value that tracks requests from your site. This is primarily used for debugging and troubleshooting purposes. Console display name: "Sub Key".
* `subscriberAccount` (string):  Console display name: "Subscriber Account".
* `subscriberId` (string):  Console display name: "Subscriber ID".
* `subscriberPasscode` (string):  Console display name: "Subscriber Passcode".
* `version` (string): This is the version of the script to load. The value should either correspond to a specific version you wish to use, or one of the following aliases to get the latest version of the code: general5 - the latest stable version of the javascript, early5 - the latest available version of the javascript. Console display name: "Version".


Example:
```hcl
resource "davinci_connection" "iovationConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "iovationConnector"
  name           = "My awesome iovationConnector"

  property {
    name  = "apiUrl"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }

  property {
    name  = "subKey"
    value = # value here
  }

  property {
    name  = "subscriberAccount"
    value = # value here
  }

  property {
    name  = "subscriberId"
    value = # value here
  }

  property {
    name  = "subscriberPasscode"
    value = # value here
  }

  property {
    name  = "version"
    value = # value here
  }
}
```


### ipgeolocation.io

Connector ID (`connector_id` in the resource): `connectorIPGeolocationio`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Developer subscription API key. Console display name: "API key".


Example:
```hcl
resource "davinci_connection" "connectorIPGeolocationio" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorIPGeolocationio"
  name           = "My awesome connectorIPGeolocationio"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### ipregistry

Connector ID (`connector_id` in the resource): `connectorIPregistry`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key used to authenticate to the ipregistry.co API. Console display name: "API Key".


Example:
```hcl
resource "davinci_connection" "connectorIPregistry" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorIPregistry"
  name           = "My awesome connectorIPregistry"

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### ipstack

Connector ID (`connector_id` in the resource): `connectorIPStack`

Properties (used in the `property` block in the resource as the `name` parameter):

* `allowInsecureIPStackConnection` (*Type inferred from the provided value*): The Free IPStack Subscription Plan does not support HTTPS connections. For more information refer to https://ipstack.com/plan. Console display name: "Allow Insecure ipstack Connection?".
* `apiKey` (string): The ipstack API key to use the service. Console display name: "API key".


Example:
```hcl
resource "davinci_connection" "connectorIPStack" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorIPStack"
  name           = "My awesome connectorIPStack"

  property {
    name  = "allowInsecureIPStackConnection"
    value = # value here
  }

  property {
    name  = "apiKey"
    value = # value here
  }
}
```


### neoEYED

Connector ID (`connector_id` in the resource): `neoeyedConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appKey` (*Type inferred from the provided value*): Unique key for the application. Console display name: "Application Key".
* `javascriptCdnUrl` (string): URL of javascript CDN of neoEYED. Console display name: "Javascript CDN URL".


Example:
```hcl
resource "davinci_connection" "neoeyedConnector" {
  environment_id = var.pingone_environment_id

  connector_id   = "neoeyedConnector"
  name           = "My awesome neoeyedConnector"

  property {
    name  = "appKey"
    value = # value here
  }

  property {
    name  = "javascriptCdnUrl"
    value = # value here
  }
}
```


### randomuser.me

Connector ID (`connector_id` in the resource): `connectorRandomUserMe`

*No properties*


Example:
```hcl
resource "davinci_connection" "connectorRandomUserMe" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorRandomUserMe"
  name           = "My awesome connectorRandomUserMe"
}
```


### tru.ID

Connector ID (`connector_id` in the resource): `connectorTruid`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (array):  Console display name: "Custom Parameters".


Example:
```hcl
resource "davinci_connection" "connectorTruid" {
  environment_id = var.pingone_environment_id

  connector_id   = "connectorTruid"
  name           = "My awesome connectorTruid"

  property {
    name  = "customAuth"
    value = # value here
  }
}
```
