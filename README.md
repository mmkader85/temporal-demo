# Demo repository to illustrate the use of temporal workflow and activity

### Steps to run the demo

Install temporal server
```shell
brew install temporal
```

Start the temporal server
```shell
temporal server start-dev
```

Run the worker first
```shell
go run worker/main.go
```

Then run the workflow
```shell
go run main.go
```

In a real-world scenario, the worker and the workflow would run in different processes.
