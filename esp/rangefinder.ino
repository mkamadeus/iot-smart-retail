#include <WiFi.h>
#include <PubSubClient.h>

const int triggerPin = 13;
const int echoPin = 12;

const int checkInterval = 900000;

const char *SSID = "SSID";
const char *PWD = "12345678";
const char *SERVER = "192.168.111.111";

const char *MQTT_UNAME = "admin";
const char *MQTT_PWD = "admin";

const char *TOPIC = "/retail/item-stock";

char message[15];

WiFiClient espClient;
PubSubClient client(espClient);
int PORT = 1883;

void connectToWiFi() {
  Serial.print("Connecting to ");

  WiFi.begin(SSID, PWD);
  Serial.print(SSID);

  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
    Serial.println("");
    delay(500);
  }
  Serial.println("Connected.");
}

void initMQTT() {
  client.setServer(SERVER, PORT);
}

void reconnect() {
  Serial.println("Connecting to MQTT Broker...");
  while (!client.connected()) {
    Serial.println("Reconnecting to MQTT Broker...");
    String clientId = "ESP32Client";

    if (client.connect(clientId.c_str(), MQTT_UNAME, MQTT_PWD)) {
      Serial.println("Connected to MQTT Broker");
    } else {
      delay(2000); // retry connecting every 2 second
    }
  }
}

void setup() {
  // put your setup code here, to run once:
  pinMode(triggerPin, OUTPUT);
  pinMode(echoPin, INPUT);

  Serial.begin(115200);
  connectToWiFi();
  initMQTT();
}

void loop() {
  if (!client.connected()) {
    reconnect();
  }
  client.loop();
  
  unsigned int pulse, cm;
  digitalWrite(triggerPin, LOW);
  delayMicroseconds(2);
  digitalWrite(triggerPin, HIGH);
  delayMicroseconds(10);
  digitalWrite(triggerPin, LOW);
  
  pulse = pulseIn(echoPin, HIGH);
  cm = pulse / 58.0;

  if (cm > 10) {
    snprintf(message, 15, "Stok habis");
    client.publish(TOPIC, message);
  } else {
    snprintf(message, 15, "Stok ada");
    client.publish(TOPIC, message);
  }
}
