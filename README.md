
```sh
docker build --no-cache -t cactus-bff .
docker run -it -d -v $(pwd):/go/src/app -p 3000:3000 --name cactus-bff cactus-bff
```