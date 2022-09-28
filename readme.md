# Chat application
- Component:
    + postgresql
    + golang(BE)
    + redis(pub|sub)
    + FE

- Before start chat application, settings database postgresql docker
+ docker run --name my-postgres -d -p 2345:5432 -e POSTGRES_PASSWORD=123321369 postgres 
Exec docker
+ docker exec -it my-postgres bash
Run command to create username and grant previleges 
+ create database chat_app;
+ create user trongpq with encrypted password '1234qwer';
+ grant all privileges on database trongpq to chat_app;

- Run redis-server
docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest