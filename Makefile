down-m2:
	docker-compose -f docker-compose-m2.yml down
up-m2:
	docker-compose -f docker-compose-m2.yml up -d
build-and-up-m2:
	docker-compose -f docker-compose-m2.yml up -d --build
up:
	docker-compose up -d
build-and-up:
	docker-compose up -d --build
down:
	docker-compose down
