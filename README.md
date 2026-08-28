# cs425-mp1

### Implementation goals:
- logger.go: produces machine.i.log file (i representing VM number)
- query.go: grep log files with query patterns of varying frequency
- unit tests prioritizing distributed logging first followed by query efficacy

### Setup/testing:
go mod init cs425/mp1 (one time call in terminal)

For manual testing you can run:

```
cd tests
go run .   // to populate machine.1.log file
// in a bash/linux terminal, run:
grep -n "Hello" machine.1.log
```

Build and run client and server files:
```
./build.sh
./bin/server  // default port 8080
./bin/client  // default port 8080

./bin/server -port 3001
./bin/client -port 3001
```


