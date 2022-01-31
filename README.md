Mongo:
sudo docker run -d -p 27017:27017 --name=mydb mongo:4.2
    d: the process was executing at background out of terminal
    -p 27017:27017 port container / port host
    --name dbname
    mongo:4.2 mongo version 
sudo docker exec -it mydatabase bash
    connection to container:
        -it: interactive form, with terminal control
        bash: shell of the conteiner
mongo


LOGIN HANDLER


curl -X POST http://localhost:9000/api/v1/users/login/ \
    -H 'Content-Type: application/json' \
    -d '{"username": "freischarler", "password": "123456"}'

returns:
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.Qi0Tc-jTChzascHaZhl0e6rRaCvAS6OJ8RLsI8Y-R78"
}

FOR QUERIE'S

curl http://localhost:9000/api/v1/users/ \
    -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.Qi0Tc-jTChzascHaZhl0e6rRaCvAS6OJ8RLsI8Y-R78'