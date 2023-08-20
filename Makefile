swagger:
	bash .github/scripts/update_version.sh
	swag fmt
	swag init -d . -o ./backend/docs
	docker-compose up --build swagger

masterkey:
	bash .github/scripts/create_masterkey.sh
