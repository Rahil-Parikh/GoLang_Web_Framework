This is a project for login authentication with following funtionalities: 
	1. Encrypted JWT for login 
	2. Encrypted JWT for Refreshing the expired Login jwt
	3. Basic Request to testing authentication when logged in 
with the help of gin-ionic REST APIs.

All the Rest APIs for testing and refereces are in gin.postman_collection.json which can be imported in PostMan software

Dependencies :
	1. Redis-Server - sudo apt-get install redis-server
	2. Redis-Tools  - sudo apt install redis-tools
	3. Cross-platform Redis management GUI (RDM)
	
Steps after setting up the dependencies :
	1. redis-cli (default is started on localhost 6379)
	2. Run routes.go file 
	3. Start testing on localhost port 8080
