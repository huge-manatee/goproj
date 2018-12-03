## prereqs:

### node / npm
(refrence best practices for your enviroment)

### serverless framework
`npm install serverless -g`


## install:

cd into project directory

`npm install`

update `./config.js` to use the values specific for AWS enviroment: keys, bucketname, etc.

don't worry about the serverless URL yet, you'll get that after deploying the serverless stack

`sls deploy` 

update serverless_host in config to reflect the endpoint URL returned


## testing
`npm test`

exceutes local requests against the deployed serverless infrastructure
