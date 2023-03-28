![coffee gopher](https://i.imgur.com/gkyVDhE.png)

Note: This repo is part of a larger service. This is purely for Golang 
practice purposes.

# Interactable TODO List

## Backend API/Model Service

The backend portion of this project will contain two larger submodules.

1.) API

2.) Model

---
### API:
In charge of communication between frontend service and backend proceedures. Exposes endpoints for the frontend to Target.

This project will likely be mostly [CRUD](https://www.freecodecamp.org/news/crud-operations-explained/#:~:text=CRUD%20refers%20to%20the%20four,data%2C%20and%20delete%20the%20data.) operations, but we will likely get creative to push the limitations of Golang and see what we can do.

---

### Model:

In charge of the interaction between the data being brought in from the api and its interactions with the database.

This will likely involve heavy use of database queries. These operations should be triggered almost soley by the API.

---

### Workflow/Layers (frontmost to backmost services):

Frontend Application ([Next.js](https://nextjs.org/docs/getting-started)

API (Will be working with [Go Fiber](https://docs.gofiber.io/))

Model (Likely will be base [Go/Golang](https://go.dev/doc/))

Database ([SQL Server](https://learn.microsoft.com/en-us/sql/sql-server/?view=sql-server-ver16))

---

### Local Environment Setup:

In the current state, since there is no database associated with this project yet
the only thing to actually stand up is the API.

To do this navigate to the api folder from the root of the project folder and use the go run command.

```
cd api
go run main.go
```

After doing so, your terminal should return a bound connection to the service and the port on your local machine. this means the program successfully started and is accessible. 

![Fiber port bound](https://i.imgur.com/J29jJVB.png)

Your code will always map to 127.0.0.1 (aka, localhost) but the port is determined by the Fiber main code in the api/main.go declaration.

```
log.Fatal(app.listen(":3000"))
```

Our program should now be accessible at the following URLS

http://localhost:3000/

http://127.0.0.1:3000/

To terminate the program, click into the terminal that is running the service and use Ctrl+C. This will exit and make the program unreachable again.

---

### Other Useful Resources:

[git](https://git-scm.com/) - code version control software

[Postman](https://www.postman.com/) - program to run RESTful request against endpoints

["What is a REST API?"](https://www.youtube.com/watch?v=lsMQRaeKNDk&ab_channel=IBMTechnology)

["Learn Go Programming"](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=2863s&ab_channel=freeCodeCamp.org)

