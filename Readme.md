docker build -t webserver-image:v1 .

docker images

docker run -d -p 80:80 webserver-image:v1

Make sure the main file is names index.html