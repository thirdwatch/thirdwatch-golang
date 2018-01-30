/* 
 * Thirdwatch API
 *
 * The first version of the Thirdwatch API is an exciting step forward towards making it easier for developers to pass data to Thirdwatch.  Once you've [registered your website/app](https://thirdwatch.ai) it's easy to start sending data to Thirdwatch.  All endpoints are only accessible via https and are located at `api.thirdwatch.ai`. For instance: you can send event at the moment by ```HTTP POST``` Request to the following URL with your API key in ```Header``` and ```JSON``` data in request body. ```   https://api.thirdwatch.ai/event/v1 ``` Every API request must contain ```API Key``` in header value ```X-THIRDWATCH-API-KEY```  Every event must contain your ```_userId``` (if this is not available, you can alternatively provide a ```_sessionId``` value also in ```_userId```).  JavaScript Fingerprinting module for capturing unique devices and tracking user interaction.  This script will identify unique devices with respect to the browser. For e.g., if chrome is opened in normal mode a unique device id is generated and this will be same if chrome is opened in incognito mode or reinstalled.  Paste the below script onto your webpage, just after the opening `<body>` tag. This script should be added to the page which is accessed externally by the user of your website. For e.g., If you want to track three different webpages then paste the below script onto each webpage, just after the opening `<body>` tag. This script should not be added onto internal tools or admin panels. ```   &lt;script id=\"thirdwatch\"     data-session-cookie-name=\"&lt;cookie_name&gt;\"     data-session-id-value=\"&lt;session_id&gt;\"     data-user-id=\"&lt;user_id&gt;\"     data-app-secret=\"&lt;app_secret&gt;\"     data-is-track-pageview=\"true\"&gt; (function() {         var loadDeviceJs = function() {         var element = document.createElement(\"script\");         element.async = 1;         element.src = \"https://cdn.thirdwatch.ai/tw.min.js\";         document.body.appendChild(element);         };         if (window.addEventListener) {              window.addEventListener(\"load\", loadDeviceJs, false);         } else if (window.attachEvent) {         window.attachEvent(\"onload\", loadDeviceJs);         }     })();   &lt;/script&gt; ``` * `data-session-cookie-name` -- The cookie name where you are saving the unique session id. We will pick the session id by reading its value from the cookie name. (Optional) * `data-session-id-value` -- In case you are not passing `data-session-cookie-name` than you can put session id directly in this parameter. In absence of both `data-session-cookie-name` and `data-session-id-value`, our system will generate a session Id. (Optional) * `data-user-id` -- Unique user id at your end. This can be email id or primary key in the database. In case of guest user, you can insert session Id here. * `data-app-secret` -- Unique App secret generated for you by Thirdwatch. * `data-is-track-pageview` -- If this is set to true, then the url on which this script is running will be sent to Thirdwatch, else the url will not be captured.   The Score API is use to get an up to date cutomer trust score after you have sent transaction event and order successful. This API will provide the riskiness score of the order with reasons. Some examples of when you may want to check the score are before:    - Before Shippement of a package   - Finalizing a money transfer   - Giving access to a prearranged vacation rental   - Sending voucher on mail  ```   https://api.thirdwatch.ai/neo/v1/score?api_key=<your api key>&order_id=<Order id> ```  According to Score you can decide to take action Approve or Reject. Orders with score more than 70 will consider as Riskey orders. We'll provide some reasons also in rules section.   ``` {   \"order_id\": \"OCT45671\",   \"user_id\": \"ajay_245\",   \"order_timestamp\": \"2017-05-09T09:40:45.717Z\",   \"score\": 82,   \"flag\": \"red\",     -\"reasons\": [     {         \"name\": \"_numOfFailedTransactions\",         \"display_name\": \"Number of failed transactions\",         \"flag\": \"green\",         \"value\": \"0\",         \"is_display\": true       },       {         \"name\": \"_accountAge\",         \"display_name\": \"Account age\",         \"flag\": \"red\",         \"value\": \"0\",         \"is_display\": true       },       {         \"name\": \"_numOfOrderSameIp\",         \"display_name\": \"Number of orders from same IP\",         \"flag\": \"red\",         \"value\": \"11\",         \"is_display\": true       }     ] } ``` 
 *
 * OpenAPI spec version: 0.0.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ai_thirdwatch

type ItemStatus struct {

	// The user's account ID according to your systems. Note that user IDs are case sensitive.
	UserId string `json:"_userId,omitempty"`

	// The user's current session ID, used to tie a user's action before and after login or account creation. Required if no user_id values is provided.
	SessionId string `json:"_sessionId,omitempty"`

	// The ID for the order that this chargeback is filed against. This field is not required if this chargeback was filed against a transaction with no _orderId.
	OrderId string `json:"_orderId,omitempty"`

	// The item's unique identifier according to your systems. Use the same ID that you would use to look up items on your website's database.
	ItemId string `json:"_itemId,omitempty"`

	// Indicates the high-level state of the order. e.g. _approved, _canceled, _held, _fulfilled, _returned, _rto
	ItemStatus string `json:"_itemStatus,omitempty"`

	// The reason for a cancellation. e.g. _paymentRisk, _abuse, _policy, _other
	Reason string `json:"_reason,omitempty"`

	// if _approved or _fulfilled than pass the shipping cost. e.g. \"50\"
	ShippingCost string `json:"_shippingCost,omitempty"`

	// if _approved or _fulfilled than pass the tracking number. e.g. \"55327470\"
	TrackingNumber string `json:"_trackingNumber,omitempty"`

	// if _approved or _fulfilled than pass the tracking url. e.g. \"http://fedex.com/track?q=abc123\"
	TrackingMethod string `json:"_trackingMethod,omitempty"`

	// The source of a decision. e.g. _automated, _manualReview\"
	Source string `json:"_source,omitempty"`

	// The analyst who made the decision, if manual.
	Analyst string `json:"_analyst,omitempty"`

	// Any additional information about this order status change.
	Description string `json:"_description,omitempty"`

	CustomInfo CustomInfo `json:"_customInfo,omitempty"`
}
