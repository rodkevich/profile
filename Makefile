psql:
	psql -h localhost -p 5432 -W postgres -U postgres -E

run:
	docker-compose start