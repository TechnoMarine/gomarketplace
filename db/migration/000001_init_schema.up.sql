CREATE TABLE "follows" (
                           "following_user_id" integer,
                           "followed_user_id" integer,
                           "created_at" timestamp not null default (now())
);

CREATE TABLE "users" (
                         "id" SERIAL PRIMARY KEY,
                         "username" varchar not null ,
                         "role" varchar not null ,
                         "created_at" timestamp not null default (now())
);

CREATE TABLE "posts" (
                         "id" SERIAL PRIMARY KEY,
                         "title" varchar not null ,
                         "body" text not null ,
                         "user_id" integer not null ,
                         "status" varchar not null,
                         "created_at" timestamp not null default (now())
);

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("following_user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id");
