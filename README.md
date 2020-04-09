# GoAPI

I didn't use any Go framework I just use here Gorilla toolkit and Gorm to build my application on to follow up with Framework pattern.

## To run the project
Create `.env` and run `docker-compose up`

## ENV
`POSTGRES_URL=postgres://postgres:yourpassword@postgres/postgres?sslmode=disable`

`PORT=`

## Routes

		Name("HomePage").
		Methods("GET").
		Path("/").
    
		Name("ServeSignupForm").
		Methods("GET").
		Path("/register").
	
		Name("RegisterAccount").
		Methods("POST").
		Path("/register").
	
		Name("ServeLoginForm").
		Methods("GET").
		Path("/login").

		Name("Login").
		Methods("POST").
		Path("/login").

		Name("Logout").
		Methods("GET").
		Path("/logout").
		
		Name("ServePhoneBookList").
		Methods("GET").
		Path("/phonebooks").
	
		Name("ServeNewPhoneBookForm").
		Methods("GET").
		Path("/phonebooks/new").

		Name("CreatePhoneBook").
		Methods("POST").
		Path("/phonebooks/new").

		Name("ServeUpdatePhoneBookForm").
		Methods("GET").
		Path("/phonebooks/{id:[0-9]+}/edit").

		Name("UpdatePhoneBook").
		Methods("POST").
		Path("/phonebooks/{id:[0-9]+}/edit").

		Name("DeletePhoneBook").
		Methods("POST").
		Path("/phonebooks/{id:[0-9]+}/delete").
