# go-ass4
## About
go-ass4 is only capable of running with three peers. No more and no less.
The critical section has been interpreted as having control of a plane. Thus only a single person has the control at on point in time.

## How to run
To run go-ass4 first start client 1:
```bash
go run main.go 1
```
Then start client 2 in another terminal:
```bash
go run main.go 2
```
Then start client 3 in another terminal:
```bash
go run main.go 3
```
This will start three peers. They will start to discuss internally which peer should get control of the plane.

## Good to know
### User Id / Port name
The numbers used as args when starting "main" are interpreted as port numbers. Thus, in the system logs, you will not see the id that has been given but the port number instead. The port number is calculated by adding the id and 5000.

### System log
When a peer request control of plane the reply it will get back will get in the form:
```bash
<time (HH:mm:ss.SSSSSS)> Got reply from id <id> : <request amount> : <is pilot>
```
When a peer is the pilot the log will be in the form:
```bash
<time (HH:mm:ss.SSSSSS)><id> is now pilot 	-----------------------
```
When they stop being the pilot the log will be in the form:
```bash
<time (HH:mm:ss.SSSSSS)><id> is not pilot 	-----------
```

