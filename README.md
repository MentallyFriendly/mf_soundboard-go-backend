# MF Soundboard Backend

## Steps To Run locally

• git clone this repository `git clone git@github.com:MentallyFriendly/mf_soundboard-go-backend.git`  
• change directories (`cd`) to the new folder  
• make a `config` directory and within it a `config.go` file ie, `mkdir config && touch config/config.go`  
• Copy code from `config_example.go` into newly created `config.go`
• Rename package name to just `config` instead of `configexample`
• You can rename your database in the params string ie, dbname=new_app_database, but make sure to update the `POSTGRES_DB=mf_soundboard-v1` environment variable in the `docker-compose.yml` file also. If you leave the docker-compose file service as `db` then you can leave `host=db` in the params here too.  

---
You can now build your image running `docker-compose up --build`

If you haven't changed the port you can test your app is running by visiting `http://localhost:8080`. The database has has not been seeded yet.  

To seed the database first run `ctrl + c` to terminate the app, and `docker-compose down` to make sure the containers have stopped running.  
Then run `docker-compose run --service-ports -e SEED=true api` to seed the database.

You should now be able to see some dummy data at any of the GET endpoints in the `endpoints.txt` file.

If you stop the containers from running, and don't need to SEED the DB again you can run either the above command without passing the -e SEED flag ie, `docker-compose run --service-ports api`, or simply `docker-compose up`.

If you modify any of the structs and need to update the database schema run `docker-compose run --service-ports -e MIGRATE=true api`.
