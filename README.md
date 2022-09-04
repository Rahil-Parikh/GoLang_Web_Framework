[![](https://img.shields.io/badge/Made_with-Go-blue?style=for-the-badge&logo=Go)]( https://go.dev/ "GoLang")[![](https://img.shields.io/badge/GraphQL-black?style=for-the-badge&logo=GraphQL&&logoColor=pink)]( https://graphql.org/ "GraphQL")[![](https://img.shields.io/badge/Redis-red?style=for-the-badge&logo=Redis&&logoColor=white)]( https://redis.io/ "Redis")

<span><h2 align="center">GoLang Web Framework </h2></span>
Contains GraphQL golang app using [gqlgen](https://github.com/99designs/gqlgen) and RESTful golang app using [gin-gonic](https://github.com/gin-gonic/gin)

#### This is a project for login authentication with following funtionalities: ####
1. Encrypted JWT for login 
2. Encrypted JWT for Refreshing the expired Login jwt
3. Basic Request to testing authentication when logged in with the help of **gin-ionic** REST APIs in Auth_jwt folder and **gqlgen** in GraphQL_system folder.

All the Rest APIs for testing and refereces are in gin.postman_collection.json which can be imported in PostMan software

### Dependencies
- Redis-Server - sudo apt-get install redis-server
- Redis-Tools  - sudo apt install redis-tools
- Cross-platform Redis management GUI (RDM)
	
### Steps after setting up the dependencies :
1. redis-cli (default is started on localhost 6379)
2. Run routes.go file 
3. Start testing on localhost port 8080

Refer the [Presentation](https://github.com/Rahil-Parikh/GoLang_Web_Framework/files/9472823/Go-Lang.IA.pptx) for more details on the gin-gonic project
