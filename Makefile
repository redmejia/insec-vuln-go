build_up: build_user_service build_mail_service
	@echo "Start Service"
	@docker-compose up

down_services: clean_user_service clean_mail_service
	@echo "Stop Service"
	@docker-compose stop
	@docker-compose down

restart: down_services build_up
	@echo "Service Restart..."


build_user_service:
	@echo "Building Users Service..."
	@cd ./user-service && env GOOS=linux CGO_ENABLED=0 go build -o dist/user-service main.go 

build_mail_service:
	@echo "Building Mail Service..."
	@cd ./mail-service && env GOOS=linux CGO_ENABLED=0 go build -o dist/mail-service main.go 

# Removig Binari files
clean_user_service:
	@echo "Remove User Service..."
	@cd ./user-service && rm dist/user-service

clean_mail_service:
	@echo "Remove Mail Service..."
	@cd ./mail-service && rm dist/mail-service