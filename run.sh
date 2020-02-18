docker build . -t gin-app
docker run -d -p 8080:8080 gin-app