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


void indexHandler();
void statusHandler();
void onHandler();
void offHandler();
void notFoundHandler();
void setHandlers();

ESP8266WebServer server(80);

void indexHandler() {
    Serial.print("indexHandler called: ");
    server.send(200, "text/html", "Welcome :)");
    Serial.println("\"text/plain\", \"Hello world!\" returned");
}

void statusHandler() {
    Serial.print("statusHandler called: ");
    if (digitalRead(4)) {
        server.send(200, "application/json", "{\"status\": 1}");
        Serial.println("{\"status\": 1} returned");
    } else {
        server.send(200, "application/json", "{\"status\": 0}");
        Serial.println("{\"status\": 0} returned");
    }
}

void onHandler() {
    Serial.print("onHandler called: ");
    digitalWrite(4, HIGH);
    if (digitalRead(4)) {
        server.send(200, "application/json", "{\"action\": \"ok\"}");
        Serial.println("{\"action\": \"ok\"} returned");
    } else {
        server.send(200, "application/json", "{\"action\": \"no\"}");
        Serial.println("{\"action\": \"no\"} returned");
    }
}

void offHandler() {
    Serial.print("offHandler called: ");
    digitalWrite(4, LOW);
    if (!digitalRead(4)) {
        server.send(200, "application/json", "{\"action\": \"ok\"}");
        Serial.println("{\"action\": \"ok\"} returned");
    } else {
        server.send(200, "application/json", "{\"action\": \"no\"}");
        Serial.println("{\"action\": \"no\"} returned");
    }
}

void notFoundHandler() {
    Serial.println("notFoundHandler called");
    server.send(404, "text/plain", "404: Not found");
    Serial.println("\"text/plain\", \"404: Not found\" returned");

}

void setHandlers() {
    server.on("/", HTTP_GET, indexHandler);
    server.on("/status", HTTP_GET, statusHandler);
    server.on("/on", HTTP_GET, onHandler);
    server.on("/off", HTTP_GET, offHandler);
    server.onNotFound(notFoundHandler);
}