# GCP-SELF-LEARN
The project mainly aims at implementing a simple login/register/logout feature with minimalistic data points . The project helps us/provides the basis to incorporate third party login/register .

# GCP Modules
The project is fully hosted in GCP In a public domain ,
## Service
  The GRPC client can be connected at 34.66.201.70:8080
  The service is auto scaling in terms of traffic upto 1000 instances
  The GCP Takes care of load balancing
  The Service has instnace monitoring that can be tied to on call
  The Instances for UI Has also been provided , but i couldn't do it on lack of time
### Databases
  The Database for PSQL has also been set up in GCP , with a back up taken at start of day .
  There is a pipiline to push the data to BQ for activity logs (Not set up due to lack of time)
  This ensures scaling and fault tolerance of GCP
### MemoryStore:
  The redis memory store is also set up in GCP ,But however turned off for costs as of now
  The redis also has a rdb back up
### Deployment:
  Cloud run has beeen set up to push to the deployments to the instance 
  Could not do side car and other deployment stuff due to lack of time
  The images are pushed to acr.gcr.io , and then are pulled into the instances
# Startuping service
Before you start up , make sure you have the following
  * local postgres (or GCP)
  * local redis (GCP)
  * the migration file under 0001_init_up.sql is run in the DB
  To start the service in a dockerized container just perform sh run.sh , in a non dockerized env , go run server/main.go
  The GRPC clients can be connected to via client/client.go
# To Do:
  *Add hashi-corp consul for config
  *Add logging for ELK kibana
  *Add 3rd party endpoints
  

