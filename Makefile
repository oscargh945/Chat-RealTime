postgresinit:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=debug -e POSTGRES_PASSWORD=debug -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql -U debug

createdb:
	docker exec -it postgres15 createdb --username=debug --owner=debug go-chat

dropdb:
	docker exec -it postgres15 dropdb go-chat

.PHONY: postgresinit postgres createdb dropdb