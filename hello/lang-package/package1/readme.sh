go build -o=test -tags=dev .
# redis://127.0.0.1:6379/0

go build -o=test -tags=prod .
# redis://192.168.0.1:6379