# go-ass4
## About
go-ass4 is only capable of running with three peers. No more and no less.
The critical section has been interpreted as having control of a plane. Thus only a single person has the control at on point in time.

## How to run
To run go-ass4 first start client 1:
```bash
go run main.go 1
```
Then start client 2:
```bash
go run main.go 2
```
Then start client 3:
```bash
go run main.go 3
```
This will start three peers. They will start to discuss internally which peer should get control of the plane.

## Good to know
The numbers used as args when starting "main" are interpreted as port numbers. Thus, in the system logs, you will not see the id that has been given but the port number instead. The port number is calculated by adding the id and 5000.
