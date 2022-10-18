- api
	+ handlers: Handler Call API, inject service (usecase)
	+ middleware: Pre processing API: CORS
	+ presenter: include json response api
- entity
	+ entity of object, func New Object and Validate
	+ error message constant
- usecase
	+ interface.go: inclule interface of repository, and Usecase interface
	+ service: implement of Usecase interface
- infrastructure/repository
	+ implement interface of repository in folder usecase