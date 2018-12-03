### Serverless solution to datastax programming homework to build RESTful interface to serve signedURLs for uploading and downloading assets to s3.

#### prepared by: https://github.com/tpinckard

-----
#### Problem statement: [Asset Uploader Assignment](Asset_uploader_assignment.pdf)

-----
## prereqs:

### node / npm
Refrence best practices for your enviroment. 

For OSX, I recommend using homebrew and installing by `brew install node`

### serverless framework
`npm install serverless -g`

-----
## install:

cd into project directory

`npm install`

update `./config.js` to use the values specific for AWS enviroment: keys, bucketname, etc.

don't worry about the serverless URL yet, you'll get that after deploying the serverless stack

Serverless stack will deploy to us-east-1 by deafult. 

It can be adjusted by modifying the provider region in the `serverless.yml` file.

`sls deploy` 

update serverless_host in config to reflect the endpoint URL returned

-----
## testing
`npm test`

exceutes local requests against the deployed serverless infrastructure
