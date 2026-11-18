debug:
	docker exec -d go-app-ecommerce dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient ./cmd/api/main.go
	@echo "✅ Debugger ready on :2345 "

start:
	docker-compose start

stop:
	docker-compose stop

rebuild:
	docker-compose up -d --build
