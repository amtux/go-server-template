# go-server-template
Basic template to quickly kick off a new project

```bash
# build the binary
go build bin/server

# run in dev: log-level defaults to debug and logging to stdout
./bin/server

# run in production
./bin/server --log-level=warn --log-file="/var/log/server.log"
```

Uses:
- Default ["net/http"](https://golang.org/pkg/net/http/) server library
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) for mux
- [sirupsen/logrus](https://github.com/sirupsen/logrus) for logging
