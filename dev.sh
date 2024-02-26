# /bin/bash

# wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run ./cmd/main.go & \
air & \
browser-sync start \
  --files './**/*.go, ./**/*.templ' \
  --ignore '*_templ.go' \
  --port 3001 \
  --reloadDelay '3000' \
  --proxy 'localhost:8080' \
  --middleware 'function(req, res, next) { \
    res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
    return next(); \
  }'

# nodemon
# yarn nodemon -e go,templ --ignore '*_templ.go' --exec 'templ generate && go run .'