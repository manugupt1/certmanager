# Welcome to CertManager!

## System requirements
- Docker
- Docker compose
- Bash
- Go

## Building
- Run the following executable `./build`  in the root folder
This will take care of building images and the volumes and finally migrating the database

NB: If you build again, it will drop the database, but won't remove the volumes.

## Running the app

`docker-compose up certmanager`

Run it on browser at

`http://localhost:3000`

NB: Right now, we run it in development mode as we are forcing SSL on production.
If the browser automatically redirects to `https`, this [stackoverflow](https://stackoverflow.com/questions/25277457/google-chrome-redirecting-localhost-to-https) post can be useful

## Running tests

`./test`

## Turning down the app to check persistence

`docker-compose down certmanager`

## API documentation

### Get all customer

		GET /customer/


### Create a customer

		POST /customer/

### Delete a customer

		DELETE /
		{
			"email":"example@example.com"
		}

### Get all, active, deactivated certificates for a given customer

		GET /customer/{cust_id}/certificates"

### Get the key as a blob for a given customer
		GET /customer/{cust_id}/certificate/{cert_id}/key"

		Allows download of only active certificate key

### Get the body of the certificate as a blob for a given customer
		GET /customer/{cust_id}/certificate/{cert_id}/body

		Allows download of only active certificate body

### Activate or deactivate a customer
		PATCH /{cust_id}/certificate/{cert_id}?active=true/false

NB: PATCH cannot be included as a payload in buffalo which is what we are using right now.


