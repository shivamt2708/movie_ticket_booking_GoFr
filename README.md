# movie_ticket_booking_GoFr
post man collection:
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://god.gw.postman.com/run-collection/29276299-0baea594-88ed-4f28-8cb6-2fc31d49d8a0?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D29276299-0baea594-88ed-4f28-8cb6-2fc31d49d8a0%26entityType%3Dcollection%26workspaceId%3D5614569b-b9a7-493f-8e16-3f436759bee4)

Backend Setup:
MYSQL Setup

Step 1:Run the docker image:

    docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=password -p 2001:3306 -d mysql:8.0.30

Step 2:create database,run this on path /zopsmart/gofr/examples/using-mysql:

    docker exec -i gofr-ssl-mysql mysql -u root -ppassword

Run
Step 2:Now run the example on path /zopsmart/gofr/examples/using-mysql by

go run main.go
