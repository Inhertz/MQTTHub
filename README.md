# **MQTTHub with Go!**

## __Description__

This is a part of my Bachelor of Engineering Final Proyect. The main goal is to recolect environmental variables generated by diferent sensors deployed on the field using IoT, specifically MQTT. The recolected data will be stored to be consulted by a Web App and analyzed in the future.

This program is the "hub" part of the proyect, therefore it recieves the incoming data, process it, and inserts it into the database.

## __Index__

1. Run with Docker (**Recomended**)
2. Run without Docker
3. Understanding the App

## __1. Run With Docker__

This is the easy way but it requires Docker! The repo provides a __"docker compose"__ file with the name: 

__"example-docker-compose.yaml"__

You can copy this file into the same directory as __"docker-compose.yaml"__. Using Linux distros:

```sh
cp ./example-docker-compose.yaml ./docker-compose.yaml
```

This will automatically run the containerized app and its dependencies with testing environmental variables.

In order to fully understand the program and customize it you should read the __"Understanding the App"__ section.

## __2. Run Without Docker__

This is the hard way! You will need installed Go in your environment, also there should be a database (check testdb/db.sql) running somewhere (SQL Server is required, but tweking the code you should be able to run with any SQL database), and a MQTT broker (I used Mosquitto). Once you have this requirements:

- Get the dependencies listed in the __"go.mod"__ file.
- Set your OS environment with this variables:

      DB_SERVER: "your_database_URI"
      DB_USER: "your_database_user"
      DB_PASSWORD: "your_database_user_password"
      DB_PORT: "your_database_port"
      DB_NAME: "your_database_name"
      
      MQTT_PROTOCOL: "your_broker_protocol_under_MQTT"
      MQTT_SERVER: "your_broker_URI"
      MQTT_PORT: "your_broker_port"
      
      HUB_SETTINGS_PATH: "./internal/adapters/mqtt/hub-settings.json"

- Run or build the go app with __"go run cmd/main.go"__ or __"go build cmd/main.go"__

## __3. Understanding the App__

The software is based on the Hexagonal Architecture, also known as "ports and adapters". This will allow the devs to separate and to "loosely couple" the domain logic in a "core" and the dependencies reliant code as a surrounding layer. Both will comunicate to each other through ports and adapters.

More specifically, this app defines three concentric layers:
- The core with the domain logic and the data models.
- The application layer with the ports and the api witch holds every use case of the app.
- The adapters that interact with the external dependencies and the core logic through the ports defined in the inner layer.

This division will help the devs to write "cleaner" and more testable code.

All of the adapters should be instantiated in the "main.go" file, this should also be "injected" (D.I.) to the application ports.

### __Paho MQTT__

The software uses a "Paho" client suscribed to all the sensor topics in order to recolect environmental data. Some MQTT client configuration can be achieved by modifying the __"hub-settings.json"__ file, but the broker URL and other sensitive information should be handled with environmental variables from the OS.

The whole project defines a topic tree that groups sensor topics in the sensor branch. For instance:

```
root/sensor/temperature/device_address
```

Each MQTT message should have a JSON payload. I.E:

```JSON
{"value": 20.34, "unit": "Celcius"}
```

The domain logic takes this unfinished model and the device address to complete it. This allows the app to recieve all kinds of sensor measurables within a single endpoint by suscribing to a wildcard "#" (check MQTT doc for more info).

### __GORM__

The database comms are managed by an adapter with GORM with SQL Server, (I know maybe PostgreSQL would have been a better choise but the B.Eng project required an SQL Server DB). This adapter will handle the connection and the queries generated by each use case in the application layer.

An example of the DB can be found inside the __"testdb"__ directory as a SQL script.

__Note:__ Following GORM's approach, the db adapter recieves data as references!

### __Docker__

Containerizing the app allows the entire project to scale better (as a hub, this program may become a bottleneck), also, a developer will find easier to perform tests this way.

The Dockerfile will create an image from Alpine Linux with Go (chosen because of its low size), source files will be copied inside the __"/app"__ directory and the compiled binary to the __"/usr/local/bin/"__. A simple entrypoint script is added just for waiting the DB initialization before launching the app.

An example of a docker compose file is included in the repo, this file contains the configuration of the app and its dependencies (also included in the testdb and testmqtt directories).