migrateup:
	migrate -path db/migration -database "postgres://zeqnqsup:9Epyab6mf2Rq2NcQXq0iyRhtR6P9zZuf@cornelius.db.elephantsql.com/zeqnqsup" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://zeqnqsup:9Epyab6mf2Rq2NcQXq0iyRhtR6P9zZuf@cornelius.db.elephantsql.com/zeqnqsup" -verbose down