/*
 * Infiot API
 *
 * # Introduction  <b>Infiot</b> provides API endpoints for interacting with <b> Infiot Management Portal</b>, so that you can rapidly deploy IoT at scale anywhere with automation.  <b>Infiot's</b> Developer-friendly SDKs and APIs enable seamless integration.Leverage <b>Infiot SDK</b> and seamlessly integrate with additional services for early time-to-market.  The <b>Infiot API</b> is a powerful [REST API](https://en.wikipedia.org/wiki/Representational_state_transfer) that can be accessed by an [HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) client tools such as curl/wget, or HTTP libraries of most modern programming languages including Python, GoLang, Java.  The API provides abundance of features, listing some of the top rated ones, * Support to interact securely with our API Servers from a Client Web application (<b>API tokens should never be exposed outside</b>).  * The API responds with a well-formatted [JSON](http://www.json.org/) data.  * Support for built-in HTTP features, like HTTP authentication and HTTP verbs, which can be easily interpreted by any HTTP clients that are designed to comply with [HTTP RFC](https://tools.ietf.org/html/rfc2616).   If you have good knowledge with REST API, our reference guide will help serve you to get started.  To start using <b>Infiot</b> APIs, API tokens need to be created  If you have any questions you can reach out to us on [support@infiot.com](mailto:support@infiot.com)  # Endpoints  Our APIs can be accessed through HTTP requests to our API Servers. There are different API servers provisioned based on production stages. The user has to select the applicable API server from the given list of servers populated, and punch in the appropriate MSP tenant name field to get started:  ``` https://{tenant}.api.infiot.net ```  All our APIs are [versioned] (#versioning). If there are any changes to the API which either changes the response format or request parameter, the version would be incremented accordingly.  # Authentication  To authenticate your API request you will need to include your secret token. You can manage your API tokens in the Infiot MSP Portal.The user should have appropriate privileges in the tenant portal for the tokens option to be visible.  Tokens can be generated from `Tokens` navigation window.Your API tokens carry many privileges, so be sure to keep them secure. Do not share your secret API tokens in publicly accessible areas such as GitHub, client-side code, and so on.   All API requests must be made over secure HTTP [HTTPS](https://en.wikipedia.org/wiki/HTTPS). Calls made over plain HTTP or without authentication will fail.  Once the token is generated and ready to use, authorize the API token in the swagger hub, click on the `Authorize` button and in the pop-up, fill the API token and click on `Authorize` again.   # Request Methods  Our API endpoints use [HTTP request methods](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods) to specify the desired operation to be performed. The documentation below specifies request method supported by each endpoint and the resulting action with its appropriate CRUD workflow.  | Method Type   | CRUD   | Description   | |-  |-  |-  | | GET   | Read   | GET requests can be used to retrieve data (eg: get all tenant details from a Master MSP or MSP)   | | POST   | Create   | POST requests can be used to create a new record (eg: adding new tenants to the Master MSP or MSP)   | | PUT   | Update/Replace   | PUT requests can be used for updating an existing record (eg: updating the name or description of a MSP   | | DELETE   | Delete   | DELETE requests can be used to delete a record (like deleting a MSP from the Master MSP) |   # Response Codes  All API requests will respond with appropriate [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). Your API client should handle each response class differently. | Response Code    | Description | |-                 |-            | | 2XX              | These are successful responses and indicate that the API request returned the expected response | | 4XX              | These indicate that there was a problem with the request like a missing parameter or invalid values | |5XX               | These indicate server errors when the server is unreachable or is misconfigured. In this case, you should retry the API request after some delay |   # <a name=\"versioning\"></a>Versioning  All our APIs are versioned. Our current API version is `v1` and we are continuously working on improving it further and provide additional endpoints. If there are any changes to an API which either changes the response format or request parameter, we will increment the version.
 *
 * API version: v5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package netskopebwan
// TenantTypeInput : Tenant type of an already created tenant can't be modified, ignored in put request body
type TenantTypeInput string

// List of TenantTypeInput
const (
	MASTER_MSP_TenantTypeInput TenantTypeInput = "Master MSP"
	MSP_TenantTypeInput TenantTypeInput = "MSP"
	ORGANIZATION_TenantTypeInput TenantTypeInput = "Organization"
)
