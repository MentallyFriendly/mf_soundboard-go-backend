# MF Soundboard Backend

## Steps To Run locally

• git clone this repository `git clone git@github.com:MentallyFriendly/mf_soundboard-go-backend.git`  
• change directories (`cd`) to the new folder   

---

From the root of of the app you can now build your image running `docker-compose up --build` 

If you haven't changed the port you can test your app is running by visiting `http://localhost:8080`. The database has not been seeded yet.  

To seed the database first run `ctrl + c` to terminate the app, then `docker-compose down` to make sure the containers have stopped running.  
Then run `docker-compose run --service-ports -e SEED=true api` to seed the database.  

You should now be able to see some dummy data at any of the GET endpoints in the `endpoints.txt` file.  

If you stop the containers from running, and don't need to SEED the DB again you can run either the above command without passing the -e SEED flag ie, `docker-compose run --service-ports api`, or simply `docker-compose up`.  

If you modify any of the structs and need to update the database schema run `docker-compose run --service-ports -e MIGRATE=true api`.  
