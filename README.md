# MF Soundboard Backend

## Steps To Run locally

• git clone this repository `git clone git@github.com:MentallyFriendly/mf_soundboard-go-backend.git`  
• change directories `cd` to the new folder   

---

From the root of of the app you can now build your image running `docker-compose up --build -d` 

If you haven't changed the port you can test your app is running by visiting `http://localhost:8024`. The database has not been seeded yet.  

To seed the database, without stopping the running container, exec `docker-compose run -e SEED=true -d api`. Once complete run `docker ps` and note the containerID that has just seeded the DB (it will have no ports and have a _run prefix). Run `docker stop <containerID>` to stop the container.  

The DB has now been seeded. You can see dummy data at any of the GET endpoints in the `endpoints.txt` file.   

If you modify any of the structs and need to update the database schema run `docker-compose run -e MIGRATE=true -d api`, and then follow the steps above to remove the container again.  

In development run docker-compose without the `-d` flag as hot-reloading has been setup with this app which means errors and general app info will be printed out to the terminal.

