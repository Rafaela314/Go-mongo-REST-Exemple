# Star Wars API

## Demo

To easily test the API, execute the `docker-compose.yml` file:

```bash

## notice the '--build', to rebuild the application container
sudo docker-compose -f  docker-compose.yaml up
```
This will compile the current source into a usable image, download the dependencies and build the necessary containers. Assuming everything goes smoothly, this command will create and run 2 containers:

1. The `mongo` container, that will run our database.
2. The `ms-starwars`, with the API itself.

To stop the containers, just hit `ctrl + c`. You can also use `docker-compose down` (inside the `demo` directory, of course) to stop and remove the containers. Be mindful that the database files themselves are in a separate volume, to keep you from accidentally losing data. 