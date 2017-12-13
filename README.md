# service-agenda

## maintainer
TangXuanzhao xuanzhaotang@gmail.com

## Tips on using
1. Create a server container with `docker run -d --name agendad agenda-golang`
2. Conntet to server with `docker run -it agenda-golang [command]`.
If the command is empty, it will start server in default. You can only use client container once because the process will quit after the command.
But you can choose to run with `docker run --entrypoint="" -it agenda-golang sh` to run with shell, in shell, you can use `agenda -h` to know more message about usage.
3. If you want to connent to other url in the client, please makesure to use `-e "SERVER_URL=value"` in docker, client will connect to the url you set on `SERVER_URL` environment.
4. Please configure two container(client and server) in one docker network, otherwise, they can not connect with eachother.
5. You can specify the server port with `PORT` enviroment, the default value is 80

## Sample
- Create a network named agenda
`docker network create agenda-golang`
- Create server container and connect it to the agenda network
`docker run -d --name agendad --network agenda agenda-golang`
- Create client container to connect to the server
`docker run -it --entrypoint="" --name agenda --network agenda agenda-golang sh`.
In shell, you can use `agenda -h` to konw how to use our program
