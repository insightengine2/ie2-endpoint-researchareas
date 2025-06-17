# IE2 Lambda Template

A template project for creating GO lambdas

## How to Use This Template

- Create a new repo based off this template
- Edit the src/main.go file
    - Add your custom logic to the function `EventHandler`

You should NOT need to change the `build.ps1` or `build.sh` scripts

This project assumes your lambda's primary entry point is the main function and assuming you don't change that everything should just work

## Usage PreRequisite(s)

- go v1.21.x
- Terraform

### Windows - Additional PreRequisites
- build-lambda-zip tool
    - `go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest`

## Compile

### Windows
- Open Powershell prompt in the root project
    - `./scripts/build.ps1`

### OSX or Linux
- Start a terminal or bash session in the project root
    - `./scripts/build.sh`

>
> You  may need to make the `build.sh` file executable. If so, issue the following in your terminal / bash session:
>
>   `chmod +x ./scripts/build.sh`
>

## Deployment

Deployments are handled automagically when your code is committed and merged with our project!

## REST Endpoint

This template project provides a file for defining your lambda function as a handler for a REST API endpoint in our system.

See the `ie2-config.yml` file for more information.
