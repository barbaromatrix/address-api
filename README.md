# address-api
Application responsible for return latitude and longitude based on IP address

## How to run locally
- Install [Go](https://go.dev/)
- Install docker [Rancher Desktop](https://ifood.atlassian.net/wiki/spaces/IL/pages/3049586786/Migrando+do+Docker+Desktop+para+o+Rancher+Desktop+no+Mac) / [Colima](https://ifood.atlassian.net/wiki/spaces/EN/pages/2971992107/Instala+o+do+Docker+no+MacBook+M1+e+Intel)

### Assemble Configuration File
This section is just to explain how to run configuration if necessary.

Run `./config.sh {deploymentName} {environment}`, this command will merge, chart values of deployment with override environment file and the secret file
The `deploymentName` is the name of the folder in k8s folder (./k8s/{deploymentName}), are available:
- address-api
To simplify in makefile we have the following commands:

| Command             | Environment                 |
|---------------------|-----------------------------|
| make config-local   | address-api local |
| make config-dev     | address-api dev   |
| make config-prod    | address-api prod  |

The `environment` are available:
- local
- dev
- prod

### Start Docker
```
make compose-infra-up
```

### Launch
```shell
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "cwd": "${workspaceFolder}",
            "program": "${cwd}/cmd/address-api/main.go"
        }
    ]
}
```

### Starting App No Debug
Once this is done, open another terminal and start the docker compose with dependencies using `docker compose up`, wait for the containers to start and run `make docker-create-topics` to create all the necessary topics
and use to start the app we can use:
| Command               | Environment           |
|-----------------------|-----------------------|
| make run-api-local    | address-api local |
| make run-api-dev      | address-api dev   |
| make run-api-prod     | address-api prod  |