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
const char* MQTT_BROKER_IP = "192.168.100.1";
const char* MQTT_TAP_TOPIC = "enroll";

const char* MQTT_UNAME = "admin";
const char* MQTT_PWD = "admin";

const char* WIFI_SSID = "XXXX";
const char* WIFI_PASSWORD = "XXXX";


// clients
MFRC522 rfid(SDA_PIN, RST_PIN);
WiFiClient wifiClient;
PubSubClient mqttClient(wifiClient);
int PORT = 1883;

void setup() {
  Serial.begin(115200);
  setup_wifi();
  setup_mqtt_client();
  SPI.begin();
  rfid.PCD_Init();  
}

void loop() {
  if(!mqttClient.connected()) reconnect_mqtt_client();
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
  mqttClient.setServer(MQTT_BROKER_IP, PORT);
}

void reconnect_mqtt_client(){
  Serial.println("Connecting to MQTT Broker...");
  while (!mqttClient.connected()) {
    Serial.println("Reconnecting to MQTT Broker...");
    String clientId = "ESP32Client";

    if (mqttClient.connect(clientId.c_str(), MQTT_UNAME, MQTT_PWD)) {
      Serial.println("Connected to MQTT Broker");
    } else {
      delay(2000); // retry connecting every 2 second
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