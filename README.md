# Installation

Clone the git repository with this command. Make sure you already installed [GIT](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) on your computer.

```bash
git clone https://github.com/ibrahimakbar31/comment-api-go
```

## Create and setup config.json file

Create **config.json** by copying template filename **config-example.json** located on root folder. **config.json** file must be saved at root folder.
You might change the **config.json** filename by updating ENV GOCONFIGFILENAME on **Dockerfile** if necessary.

Fill up each parameters. DB1 is main DB which used postgresql database.

**config.json** file have 4 environments: **development**, **test**, **staging** and **production**.

If **GOCUSTOMENV** environment is not set up on your computer, then it will run **development** environment as default environment

you might set **"Migration": true** if you wanna build database table for the first time. And set to **"Migration": false** when table already created.

Default application port is **8000**

## Testing

From root folder, go to **router** folder

```bash
cd router
```

and run this command to run testing

```bash
go test -v
```

## Docker

Make sure you already installed [Docker](https://docs.docker.com/get-docker/) on your computer.

You might setup docker image on **Dockerfile** if necessary. you might change your environment on **GOCUSTOMENV** parameter. Default Docker port is **8000**.

run this command to create docker image

```bash
docker build -t comment-api-go .
```

From above command we created docker image with name: **comment-api-go**.

run this command to mount docker image

```bash
docker run -p 8000:8000 -it comment-api-go
```

## API Documentation

API documentation can be found on:
[https://documenter.getpostman.com/view/3373119/T1DjkfTV](https://documenter.getpostman.com/view/3373119/T1DjkfTV)
