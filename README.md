# ForgeRock SaaS Software Engineer Coding Challenge

## Project Pre-Reqs

- Go
- GCC
- Make

## Building The Project

From the root of the project
"make build" - to build the services which will be in \<root\>/bin
"make test" - to run the tests
"docker build -t saas-interview-challenge1 ." - to build the project

## Running The Project
From the root of the project
"docker-compose up" - This will start a job service listening for web requests, a worker service subscribed to new job events from the job service, and a redis server as the pub\sub service. The job service will try to open port 8000 on localhost. If there is a conflict update the docker-compose.yml file

To try the service use Postman to submit new jobs
POST http://localhost:8000/job and change the body per the desired job as shown below

Simple "Sleep Worker"
    - Description: Takes number of seconds to sleep and after the sleep time expires it prints the provided message
    - Sample Request Body: {"Type":"Sleep", "Data":"{\"SleepTimeInSeconds\": 10, \"Message\": \"Hello World\"}"}

## Requirements / Concepts Discussed

Some requirements / concepts we would like to see (either in code or at least outlined in the README):

- "Are there any shortcomings of the code?"
    There are some limitations in its current condition but it is in a starting state so it could be iterated on and improved
    1. Running of workers implementation is just in process as the events are handled currently. The idea is that new Launcher implementations could be added. For example one that uses the docker api to launch a worker as a container that runs and exits. This would allow additional workers to be written in other technologies independent from the worker service.
    2. The example Launcher hard codes the knowledge of the worker. A real example should have a mechanism to allow workers and how to launch them to be registered either through a config file or an API
    3. I abstracted the redis publisher behind an API. The subscriber should be abstracted in a similar way. Then both the job and worker service should have which type of publisher\subscriber implementation to use set through configuration
    4. Events can be lost if the worker isn't up and listening when they are sent. There needs to be enough reliability in the pub/sub service chosen to allow for service restarts of itself and of the worker services
    5. There needs to be security on the services
    6. Logging is just simple prints. The code needs logging levels added
    7. This service is currently fire and hope for the best. There should be an additional repo for the job service where it could register the jobs and provide end user functionality to check job status. The workers could publish events for job status like completed or failed.
    8. In a real implementation, only 1 instance of a worker-service should process an event. Currently all instances will process the event
    9. More tests need to be written. Need to change the build to know when to run unit tests and when to run integration tests. Intgration tests need trigger setup of dependent services. Currently redis_test.go fails if a redis server isn't already running.
    

- "How might this project be scaled?"
    **Worker Service** - It depends on the requirements on what kind of workers need to be run and where the system could be hosted. If the workers tasks aren't too long and AWS is an option then using Lambda for the workers could be a good idea. Having the worker_service launch the jobs as containers in a kubernetes or swarm cluster could be another approach. Additionally the definition of work could be extended to include a platform variable. Then maybe certain jobs are launched via different channels where there might be a worker_service listening that can run a task on a GPU enabled host

    **Job Service** - Multiple instances of the service could be run behind a load balancer or as a service in something like Kubernetes or Cloud Foundry.

    **Pub / Sub Service** - This could be scaled by clustering the Redis service per its documentation. Besides scale, if event delivery needs to be guaranteed, then another clustered service may be best like Kafka


- How might one approach doing sequential versus parallel tasks?
    One approach for sequential could be to include a chain of execution as an optional field in the registration of the work. Workers could have a shared library that looks for their place in the chain once they are completed successfully and publishes the work as the next job type along with the payload that maintains any expected state. If working with larger amounts of data, then the data would have to be hosted on a shared resource.

    If the work for some reason needed to be isolated to the host that started the work, then a new worker type could be implemented that maintains the chain of work state and directly calls all of the workers as local functions. 


