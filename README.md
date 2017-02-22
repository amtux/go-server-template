# go-server-template
Basic template to quickly kick off a new project

**Make it go**
```bash
make clean #delete bin/ directory
make build #install dependencies and build
make run-dev #run with dev configuration
```

**Run straight-up**:
```bash
# build the binary
go build -o bin/server

# run in dev: log-level defaults to debug and logging to stdout
./bin/server

# run in production
./bin/server --log-level=warn --log-file="/var/log/server.log"
```

Uses:
- Default ["net/http"](https://golang.org/pkg/net/http/) server library
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) for mux
- [sirupsen/logrus](https://github.com/sirupsen/logrus) for logging
