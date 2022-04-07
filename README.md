# iot-smart-retail

This project contains the backend for our project in IF4051 IoT System Development. Maintanined by 3 people:

- Matthew Kevin Amadeus - 13518035
- Samuel - 13518041
- Rakha Fadhilah - 13518097

### Tech Stack

Powered by Go 1.18. Notable libraries include:

- Fiber
- GORM

### Running the Project

#### Prerequisites

Go 1.18 is a must. Docker is optional but recommended to provision the database, cache, and broker. Alternatively, you can set it up on your own local machine.

#### Steps

1. Make a `.env` file, fill it in with correct details according to your system.
2. Run `docker-compose up` in the `docker/` directory.
3. Run `make build`.
4. Run the compiled executable.
