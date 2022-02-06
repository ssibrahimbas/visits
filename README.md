<p align="center"><br><img src="https://avatars.githubusercontent.com/u/76786120?v=4" width="128" height="128" style="border-radius: 50px;" /></p>
<h3 align="center">Ssibrahimbas visits</h3>
<p align="center">
  A demo project that automatically restarts with a trio of docker, redis and go and transmits page visits.
</p>

### Usage

```shell

$ docker-compose up

```

 Step One:
>
> visit to this address: `http://localhost:8081`
>
> Your result should be: `0`


Step Two:

> visit to this address: `http://localhost:8081`
>
> Your result should be: `1`

Step Three:

> visit to this address: `http://localhost:8081/stop`
>
> This request stops the application.

Step Four: 

> visit to this address: `http://localhost:8081`
>
> Your result should be: `0`
> 
> The visit count reset because docker-compose restarts the app  when app kill.
>
> Look up docker-compose.yml:6