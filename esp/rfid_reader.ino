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
  
  if (rfid.PICC_ReadCardSerial()) { // NUID has been readed
    MFRC522::PICC_Type piccType = rfid.PICC_GetType(rfid.uid.sak);
    Serial.print("RFID/NFC Tag Type: ");
    Serial.println(rfid.PICC_GetTypeName(piccType));

    // print UID in Serial Monitor in the hex format
    Serial.print("UID:");
    for (int i = 0; i < rfid.uid.size; i++) {
      Serial.print(rfid.uid.uidByte[i] < 0x10 ? " 0" : " ");
      Serial.print(rfid.uid.uidByte[i], HEX);
    }
    Serial.println();

    rfid.PICC_HaltA(); // halt PICC
    rfid.PCD_StopCrypto1(); // stop encryption on PCD
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

void check_in(char* userId) {
  mqttClient.publish(MQTT_TAP_TOPIC, userId);
}
