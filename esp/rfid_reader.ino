#include <SPI.h>
#include <MFRC522.h>
#include <WiFi.h>
#include <PubSubClient.h>

// constants
#define SDA_PIN 21
#define SCK_PIN 5
#define MISO_PIN 19
#define MOSI_PIN 23
#define RST_PIN 5

const char* MQTT_CLIENT_NAME = "esp32-rfid";
const char* MQTT_BROKER_IP = "192.168.2.122";
const char* MQTT_TAP_TOPIC = "tap";

const char* WIFI_SSID = "ASUS_GF";
const char* WIFI_PASSWORD = "melati25";

// clients
MFRC522 rfid(SDA_PIN, RST_PIN);
WiFiClient wifiClient;
PubSubClient mqttClient(wifiClient);

void setup() {
  Serial.begin(115200);
  setup_wifi();
  SPI.begin();
  rfid.PCD_Init();  
}

void loop() {
  if (!rfid.PICC_IsNewCardPresent()) return;
  
  String cardId = getCardId();
  if(cardId != "") {
    Serial.println(cardId);
    tapCard(cardId);
  }
}

// setup wifi
void setup_wifi() {
  Serial.print("Connecting to ");
  Serial.println(WIFI_SSID);
  
  WiFi.mode(WIFI_STA);
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);

  while(WiFi.status() != WL_CONNECTED) {
    delay(50);
    Serial.print(".");
  }
  Serial.println();
  Serial.print("Connected with address ");
  Serial.println(WiFi.localIP());
}

void setup_mqtt_client() {
  while(!mqttClient.connected()) {
    Serial.print("Connecting to MQTT broker at ");
    Serial.println(MQTT_BROKER_IP);
    if(mqttClient.connect(MQTT_CLIENT_NAME)) {
      Serial.println("Connected!");
    }
    else {
      Serial.println("Failed to connect, trying again in 5 seconds");
      delay(5000);
    }
  }
}

String getCardId() {
   if (!rfid.PICC_ReadCardSerial()) return "";
  
  // print UID in Serial Monitor in the hex format
  String result = "";
  for (int i = rfid.uid.size - 1; i >= 0; i--) {
    result += rfid.uid.uidByte[i];
  }
  
  rfid.PICC_HaltA(); // halt PICC
  rfid.PCD_StopCrypto1(); // stop encryption on PCD

  return result;
}

void tapCard(String cardId) {
  mqttClient.publish(MQTT_TAP_TOPIC, cardId.c_str());
}