# ScrollFix - Backend
<img src="https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/scrollflix_logo.png" alt="alt text" width="200">


**Scrollflix** is a single page web-application that provides users curated movie recommendation based on how users interact with the landing page, which is the front page. It capture users' scroll behavior on movie posts and make movie suggestion based on the posts' view duration. This repository provides information and overview of back-end for the web application.In this repository you will all the coding from the backend and informtaion on how to run it. However, to get the web application running, front-end should be pulled too. For the application to run competely, both repository need to be pulled.


**Related Works**
- **ElasticSearch:** is a tool used in industry to collect data in both logged in users and site visitor.
- **Facebook:** uses elasticsearch for their News Feed algorithm. By analyzing data points such as likes and content interaction the algorithm can indirectly infer users' intention and provides personal recommendation to enhanced user experience
- **Netflix:** uses Elasticsearch to recommend content to millions of users by collects a variety of data point (viewing history, rating, scroll activity, navigation, to help users find relevant content)
- **Amplitude:** provides the API for developer to develop a product analytics dashboard to collect a user engagement data, thus they can provide data insights to stakeholders. To do this, we need to require a redesign the entire architecture from event-base system into a log-base system. What our team trying to do is to make an extension version of amplitude that is: you don't need to redesign your architecture into event base; if you know where to print the log to console.log, we can capture that and still able to use that data for content curation.


**Frontend** code can be found [here](https://github.com/mnguyen0226/movie-suggestions-frontend).

## Demo
Home Page
<br/>
<img src="https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/home_page.gif" alt="alt text">

Scroll Suggestion
<br/>
<img src="https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/movie_suggestion.gif" alt="alt text">

Movie Trend
<br/>
<img src="https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/movie_trend.gif" alt="alt text">

## Architecture Designs
<img src="https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/overall_arc.png" alt="alt text">


### Frontend
<img src="https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/frontend_arc.png" alt="alt text">

### Backend
<img src="https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/backend_arc.png" alt="alt text">


## How To Run
1. Clone the github repository
```sh
git clone  git@github.com:Harishgeth/movie-suggestions-api.git 
```
2. Install Docker Compose

3. Install Mongo Compass

    3.1 If you want to connect to our Mongodb (Mongodb starts up at http://localhost:27018. The default username/password is user/pass for non-m2 machines, but for m2 you just use noauth)

4. To Run Docker

 For windows and mac with M1 chip:
        

```sh
# Run Docker with rebuilding the API
make up

# Built and Run Docker
make build-and-up

# Disconnect Docker
make down
```

For mac with M2 chip:
        

```sh
# Run Docker without rebuilding the API
make up-m2

# Built and Run Docker
make build-and-up-m2

# Disconnect Docker
make down-m2
```

Please ensure the following ports are available, [if not kill applications running on ports.](https://stackoverflow.com/questions/11583562/how-to-kill-a-process-running-on-particular-port-in-linux) 

Ports required by applications to be brought up - 9200:Elasticsearch, 5601:Kibana, 3000-Golang, 27018 - MongoDB.


    

5. Pull the front-end code from [here.](https://github.com/mnguyen0226/movie-suggestions-frontend)

6. Follow the instructions in the frontend repository to run it, and it should automatically recognize the backend.

7. You can checkout the elasticsearch data at Kibana which runs in the [url.](http://localhost:5601/app/kibana). There are two index pattern of interest.

Operational Excellence Logs - filebeat*

Movie curation Index - movie*

Use the Discover tab of Analytics to create a data view with these index patterns described above to see the data :)
## Unit Test

## Tools
- **Frontends:** Vue.js, HTML, CSS, Javascript, Bootstrap, Mozilla InteractionObserver API.
- **Backends:** Filebeat, Kibana, ElasticSearch, Vector, Go, Python, MongoDB, Docker.
- **Development Tools:** VSCode, Postman, ElasticSearch, Docker.
- **Teamwork Tools:** Slack, Git.

## References
- [Presentation](https://github.com/mnguyen0226/movie-suggestions-front-end/blob/main/docs/presentation.pdf).
- [Report]().