# Finance Server: A Project Both For Intellectual and Financial Profit
This project was build to learn to use the following tools
* Go
* Docker
* MakeFiles

## Details of Testing
For the purposes of testing, Docker containers were set up for these specific elements of the program
* The main server
* The postgres instance

When running tests, instead of running tests raw locally, docker containers are spun up for these elements so that each of these elements can be tested against each other, but but these parts can still run in an isolated fashion which would be more representative of how things would run in production where the main server and the postgres server might not be residing on the same machine

To avoid human error which can arise due to forgetting to rebuild the image after changes have been made to the source code, whenever any of the containers are spun up, the commands in the Makefile require that the images be rebuilt first before the containers are initialised. 

While it's true that even if there are no changes in the source code, this can cost some time while Docker is checking against cached files to see if any images actually need to be rebuilt, the decision was made that in the long run having this check would save time that would be spent debugging during testing due to the developer forgetting to rebuild images before running tests

