create table "tasks" (
	"id" SERIAL not null unique,
	"event" VARCHAR(256) not null default 'nothing',
	"requested_by" VARCHAR(256) not null default '<unknown>',
	"affects" TEXT,
	"deadline" TIMESTAMP not null default NOW()
);