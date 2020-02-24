#!/bin/bash
#

# docker run --rm -it  -p 9000:80  --name app-go -v $(pwd)/app:/go/src/app danisbagus/app-go
docker run --rm -p 3000:3000  --name app-go danisbagus/app-go

exec "$@"
