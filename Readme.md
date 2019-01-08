In this example you will learn: 

    How to write a pod definition in Kubernetes. 
    
    How to expose the pod by writing a service
    
    Why writing a pod directly is a bad idea and how Replicaset solves it

    Drawbacks of Replicaset

    How Deployment solves issue with replicaset

Docker commands to build image of the application:

    docker build -t webserver-image:v1 .

    docker images

    docker run -d -p 80:80 webserver-image:v1

Make sure the main file is named index.html

**Steps:**

Have the game in folder
Create Dockerfile
Create image
Deploy the image in a pod
Write Deployment yaml
Deploy in gcloud

Steps to create docker image and push it to registry: 

    docker login

    docker build -t "game:v0" .

    docker images

    docker tag game:v0 emailtovamos/game-repo

    docker push emailtovamos/game-repo

Ensure you have VirtualBox running and the current-context of kubectl is set to be minikube:

    minikube start

    kubectl config get-contexts

    kubectl config use-context minikube

Run the yaml files: 

    kubectl apply -f pod.yaml

    kubectl apply -f service.yaml

    kubectl apply -f deployment1.yaml


Test: 

A test might be to check if setting a highscore gets reflected while you try to get the highscore. 
CircleCI can be included with this test. Also to deploy to Google Cloud.

Backend: 

Make the game harder by making shapes smaller if a user plays well. These decisions can be made in another service running on Go backend. There can be gRPC connections between this service, the highscore service and the BFF service.

