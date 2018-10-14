docker build -t webserver-image:v1 .

docker images

docker run -d -p 80:80 webserver-image:v1

Make sure the main file is names index.html

Steps: 

Have the game
Create Dockerfile
Cretae image
Deploy the image in a pod
Write Deployment yaml
Deploy in gcloud

Steps to create docker image and push it to registry: 

    docker login

    docker build -t "game:v0" .

    docker images

    docker tag game:v0 emailtovamos/game-repo

    docker push emailtovamos/game-repo

kubectl deploy -f pod.yaml

kubectl deploy -f service.yaml