#include <ESP8266mDNS_Legacy.h>
#include <ESP8266mDNS.h>
#include <LEAmDNS_lwIPdefs.h>
#include <LEAmDNS_Priv.h>
#include <LEAmDNS.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WiFiMulti.h> 
#include <ESP8266mDNS.h>
#include <ESP8266WebServer.h>

// Includes server and handlers
#include "routes.hpp"
// Includes AP SSID and password
/*
secrets.hpp

#define _SSID "Your_SSID_here"
#define _PASS "Your_pass_here"

*/
#include "secrets.hpp"

ESP8266WiFiMulti wifiMulti;

void setup(void) {
  // GPIO 4 = D2, lamp pin
  pinMode(4, OUTPUT);
  Serial.begin(115200);
  delay(10);
  Serial.println('\n');

  // You can add multiple ap
  wifiMulti.addAP(_SSID, _PASS);
  //wifiMulti.addAP("ssid_from_AP_2", "your_password_for_AP_2");
  //wifiMulti.addAP("ssid_from_AP_3", "your_password_for_AP_3");

  Serial.println("Connecting ...");
  while (wifiMulti.run() != WL_CONNECTED) {
    delay(250);
    Serial.print('.');
  }
  Serial.println('\n');
  Serial.print("Connected to ");
  Serial.println(WiFi.SSID());
  Serial.print("IP address:\t");
  Serial.println(WiFi.localIP());

  // Starts multicast DNS 
  if (MDNS.begin("lamp")) {
    Serial.println("mDNS responder started");
  } else {
    Serial.println("An error occurred while setting up MDNS responder!");
  }

  // Sets path handlers to server object
  setHandlers();
  // Starts the ESP server
  server.begin();
  Serial.println("HTTP server started");
}

void loop(void) {
  server.handleClient();
  MDNS.update();
}

