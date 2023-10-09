
# Doppler-Go

Go SDK client for doppler.com API.


## Installation

```bash
  go get -u github.com/dilutedev/doppler
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DOPPLER_KEY`

Read more about keys [here](https://docs.doppler.com/reference/api#authentication)

## Documentation

[API Documentation](https://docs.doppler.com/reference/api)


## Features

- [x] Auth
- [x] Workplace
- [x] Workplace Roles
- [x] Activity logs
- [x] Projects
- [x] Project Roles
- [x] Project Members
- [x] Environments
- [x] Configs
- [x] Config logs
- [x] Trusted IPs
- [x] Secrets
- [x] Integrations
- [x] Secrets Sync
- [x] Dynamic Secrets
- [x] Service Tokens
- [x] Invites
- [x] Groups
- [x] Service Accounts
- [x] Audit
- [x] Share


## Usage/Examples

```golang
package main

import (
    "github.com/dilutedev/doppler"
    _ "github.com/joho/godotenv/autoload"
)

func main(){
    dp, err := doppler.NewFromEnv()
    if err != nil {
        panic(err)
    }

    projects, err := dp.ListProjects(1, nil)
    if err != nil {
        panic(err)
    }

    log.Println(projects)
}
```


## Authors

- [@dilutedev](https://www.github.com/dilutedev)
- [@sudodeo](https://www.github.com/sudodeo)

