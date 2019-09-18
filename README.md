# Cactus BFF service

BFF service for the Cactus app.

```sh
docker build --no-cache -t cactus-bff .
docker run -it -d -v $(pwd):/go/src/app -p 3000:3000 --name cactus-bff cactus-bff
```

## License
MIT © [Jan Guzman](https://github.com/Krystian19)
