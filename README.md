## Minh's Note
- To Run:
  - root@pop-os:/home/mnguyen0226/Documents/school/graduate/spring_2023/cs5704/movie-suggestions-api# docker build -t minh/movie-suggestions-api .
  - root@pop-os:/home/mnguyen0226/Documents/school/graduate/spring_2023/cs5704/movie-suggestions-api# docker run -d -p 8080:8080 --env-file dev.env minh/movie-suggestions-api
  - Now, I don't have to go to root and just do docker-compose up -d on current directory.

# movie-suggestions-api

Requires Golang to run
Supposedly scraps data from IMDB, but serves as a boilerplate to structure backend
Move code to $GOPATH/src
Use ./build.sh to run it!

