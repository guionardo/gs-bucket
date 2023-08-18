swagger:
	swag fmt
	swag init -d . -o ./backend/docs
	docker-compose up --build swagger
