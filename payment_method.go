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

// The payment_method field type represents information about the payment methods provided by the user. The value must be a nested object with the appropriate item subfields for the given payment method. Generally usedwith the create_order or transaction events. 
type PaymentMethod struct {

	// Values like - _cash, _check, _creditCard, _debitCard, _netBanking, _wallet, _cryptoCurrency, _electronicFundTransfer, _financing, _giftCard, _interac, _invoice, _moneyOrder, _masterPass, _points, _storeCredit, _thirdPartyProcessor, _voucher
	PaymentType string `json:"_paymentType,omitempty"`

	// The item unit price in numbers, in the base unit of the currency_code.e.g. \"2500\". In case of multiple payment methods in order it's useful.
	Amount string `json:"_amount,omitempty"`

	// The [ISO-4217](http://en.wikipedia.org/wiki/ISO_4217) currency code for the amount. e.g., USD, INR alternative currencies, like bitcoin or points systems. In case of multiple payment methods in order it's useful.
	CurrencyCode string `json:"_currencyCode,omitempty"`

	// fill value like bank name, gateway name, wallet name etc, e.g. payu, paypal, icici, paytm
	PaymentGateway string `json:"_paymentGateway,omitempty"`

	// Account name oif the user for that payment method
	AccountName string `json:"_accountName,omitempty"`

	// The first six digits of the credit card number. These numbers contain information about the card issuer, the geography and other card details.
	CardBin string `json:"_cardBin,omitempty"`

	// Response code from the AVS address verification system. Used in payments involving credit cards.
	AvsResponseCode string `json:"_avsResponseCode,omitempty"`

	// Response code from the credit card company indicating if the CVV number entered matches the number on record. Used in payments involving credit cards.
	CvvResponseCode string `json:"_cvvResponseCode,omitempty"`

	// The last four digits of the credit card number.
	CardLast4 string `json:"_cardLast4,omitempty"`

	// Expiry month of the card.
	CardExpiryMonth string `json:"_cardExpiryMonth,omitempty"`

	// Expiry year of the card.
	CardExpiryYear string `json:"_cardExpiryYear,omitempty"`
}
