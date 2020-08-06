Установка
------------

Склонировать репозиторий

```
git clone https://github.com/splinbl4/7Skills-test.git
```

Для того, чтобы поднять сервис локально, необходимо выполнить следующие команды:

```
make init
```

либо

```
docker build -t fibonacci-go .
```

```
docker run -p 8000:8000 fibonacci-go
```

Необходимые адреса.
------------
```
http://localhost:8000/?start=0&end=16
```