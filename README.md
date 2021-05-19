# Star Wars API

## Demo

To easily test the API, add a .env in the root of this projects with variables listed below and then execute the `docker-compose.yml` file:

```json

API_PORT=1234


PRODUCTION=false

SWAPI_URL=https://swapi.dev/api

MONGO_URL=mongodb://msstarwars.mongo:27017


```

```bash
sudo docker-compose -f  docker-compose.yaml up 
```
This will compile the current source into a usable image, download the dependencies and build the necessary containers. Assuming everything goes smoothly, this command will create and run 2 containers:

1. The `mongo` container, that will run our database.
2. The `ms-starwars`, with the API itself.

To stop the containers, just hit `ctrl + c`. You can also use `docker-compose down` (inside the `demo` directory, of course) to stop and remove the containers. Be mindful that the database files themselves are in a separate volume, to keep you from accidentally losing data. 

There is an available insomnia collection in the project that can be imported and used for integration test purposes.