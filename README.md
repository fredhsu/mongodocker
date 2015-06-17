[fhsu@fred-arch hotaisle]$ docker build -t hotaisle .
[fhsu@fred-arch hotaisle]$ docker run -it --rm --name hotaislerun --link some-mongo:mongo hotaisle

