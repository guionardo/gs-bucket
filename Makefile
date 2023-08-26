swagger:
	bash .github/scripts/update_version.sh
	swag fmt
	swag init -d . -o ./backend/docs
	docker-compose up --build swagger

masterkey:
	bash .github/scripts/create_masterkey.sh

get_masterkey:
	fly ssh sftp get master.key ./master.key.fly
