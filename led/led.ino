#define SECOND 1000

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
  // GPIO 5 = D1, PIR motion sensor pin
  pinMode(5, INPUT);
  //GPIO 4 = D2, LED pin
  pinMode(4, OUTPUT);
  Serial.begin(115200);
  delay(10);
  Serial.println('\n');

  // You can add multiple AP
  wifiMulti.addAP(_SSID, _PASS);
  //wifiMulti.addAP("ssid_from_AP_2", "your_password_for_AP_2");
  //wifiMulti.addAP("ssid_from_AP_3", "your_password_for_AP_3");

  Serial.println("Connecting ...");
  int i = 0;
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
  if (MDNS.begin("led")) {
    Serial.println("mDNS responder started");
  } else {
    Serial.println("Error setting up MDNS responder!");
  }

  // Sets path handlers to server object
  setHandlers();
  // Starts the ESP server
  server.begin();
  Serial.println("HTTP server started");
}

unsigned long lastDetect = millis();
void loop(void) {
  // If motion detected and last motion was 10 second or more time ago, sends info to main server
  if (digitalRead(5) && (millis() - lastDetect > SECOND * 10)) {
    // Updates last motion detection time
    lastDetect = millis();
    Serial.println("Motion detected");
    // Creates http object to send a post request to main server
    HTTPClient http;
    http.begin("http://192.168.2.145:8020/watch");
    http.addHeader("Content-Type", "application/x-www-form-urlencoded");
    String httpRequestData = "move=ok";
    int httpResponseCode = http.POST(httpRequestData);
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);
    http.end();
  }

  server.handleClient();
  MDNS.update();
}
