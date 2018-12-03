## prereqs:

### node / npm
(refrence best practices for your enviroment)

### serverless framework
`npm install serverless -g`


## install:

cd into project file
`npm install`

update ./config.js to use your values for AWS keys, bucketname, etc.
don't worry about the serverless URL yet, you'll get that after deploying the serverless stack
`sls deploy` 

update serverless_host in config to reflect the endpoint URL returned


## testing
`npm test`
