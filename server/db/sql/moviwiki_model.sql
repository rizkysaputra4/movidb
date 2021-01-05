CREATE TABLE "user_information" (
  "user_id" INT NOT NULL,
  "birthdate" DATE NOT NULL,
  "signup_date" DATE NOT NULL DEFAULT CURRENT_DATE,
  "bio" TEXT,
  "fb_link" VARCHAR(255),
  "twitter_link" VARCHAR(255),
  "ig_link" VARCHAR(255),
  "sex" BOOLEAN,
  "last_request" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "movie_information" (
  "movie_id" SERIAL PRIMARY KEY NOT NULL,
  "movie_title" VARCHAR(255) NOT NULL,
  "movie_synopsis" VARCHAR(255) NOT NULL,
  "release_date" DATE NOT NULL,
  "imdb_rating" FLOAT NOT NULL,
  "imdb_numb_vote" INT NOT NULL,
  "site_rating" FLOAT NOT NULL,
  "site_numb_vote" INT NOT NULL,
  "poster_link" VARCHAR(255) NOT NULL,
  "trailer_link" VARCHAR(255) NOT NULL,
  "duration" VARCHAR(255) NOT NULL,
  "awards" VARCHAR(255),
  "type_id" INT NOT NULL,
  "country_id" INT NOT NULL,
  "language" VARCHAR(255) NOT NULL,
  "unique_link" VARCHAR(255) NOT NULL,
  "identifier_id" INT DEFAULT '0'
);

CREATE TABLE "eps_information" (
  "eps_id" SERIAL PRIMARY KEY NOT NULL,
  "movie_id" INT NOT NULL,
  "episode_nmbr" INT NOT NULL,
  "year" INT NOT NULL,
  "season" INT NOT NULL,
  "synopsis" TEXT
);

CREATE TABLE "movie_reviews" (
  "review_id" SERIAL PRIMARY KEY NOT NULL,
  "movie_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "text" TEXT NOT NULL,
  "story_rating" INT NOT NULL,
  "acting_rating" INT NOT NULL,
  "cinema_rating" INT NOT NULL,
  "music_rating" INT NOT NULL,
  "identifier_id" INT NOT NULL DEFAULT '0',
  "like_count" INT NOT NULL DEFAULT '0',
  "timestamp" TIMESTAMP NOT NULL
);

CREATE TABLE "review_comment" (
  "comment_id" SERIAL PRIMARY KEY NOT NULL,
  "identifier_id" INT NOT NULL,
  "review_id" INT NOT NULL,
  "reply_to" INT NOT NULL,
  "user_id" INT NOT NULL,
  "text" TEXT NOT NULL,
  "like" INT NOT NULL DEFAULT '0',
  "timestamp" TIMESTAMP NOT NULL
);

CREATE TABLE "movie_comment" (
  "comment_id" SERIAL PRIMARY KEY NOT NULL,
  "reply_to" INT NOT NULL,
  "movie_id" INT NOT NULL,
  "text" TEXT NOT NULL,
  "like_count" INT NOT NULL DEFAULT '0',
  "identifier_id" INT NOT NULL DEFAULT '0',
  "user_id" INT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL
);

CREATE TABLE "user_login" (
  "user_id" INT NOT NULL,
  "user_name" VARCHAR(255) NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL
);

CREATE TABLE "user_movie_rating" (
  "movie_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "rating" FLOAT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL
);

CREATE TABLE "user_short_info" (
  "user_id" SERIAL PRIMARY KEY NOT NULL,
  "user_full_name" VARCHAR(255) NOT NULL,
  "country_id" INT
);

CREATE TABLE "user_eps_rating" (
  "eps_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "rating" FLOAT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL
);

CREATE TABLE "user_social_activity" (
  "user_id" INT NOT NULL,
  "activity_type" INT NOT NULL,
  "activity" INT NOT NULL
);

CREATE TABLE "social_activity_type" (
  "type_id" SERIAL PRIMARY KEY NOT NULL,
  "activity_type" VARCHAR NOT NULL
);

CREATE TABLE "like_count_review_comment" (
  "comment_id" INT NOT NULL,
  "user_id" INT NOT NULL
);

CREATE TABLE "like_count_review_movie" (
  "review_id" INT NOT NULL,
  "user_id" INT NOT NULL
);

CREATE TABLE "like_count_movie_comment" (
  "comment_id" INT NOT NULL,
  "user_id" INT NOT NULL
);

CREATE TABLE "movie_genre" (
  "movie_id" INT NOT NULL,
  "genre_id" INT NOT NULL
);

CREATE TABLE "genre_list" (
  "genre_id" SERIAL PRIMARY KEY NOT NULL,
  "genre_name" VARCHAR NOT NULL
);

CREATE TABLE "identifier" (
  "identifier_id" SERIAL PRIMARY KEY NOT NULL,
  "info" VARCHAR NOT NULL
);

CREATE TABLE "movie_tag" (
  "movie_id" INT NOT NULL,
  "tag_id" INT NOT NULL
);

CREATE TABLE "tag_list" (
  "tag_id" SERIAL PRIMARY KEY,
  "tag_name" VARCHAR NOT NULL
);

CREATE TABLE "person_in_movie" (
  "person_id" INT NOT NULL,
  "movie_id" INT NOT NULL
);

CREATE TABLE "movie_person" (
  "person_id" SERIAL PRIMARY KEY,
  "person_name" VARCHAR NOT NULL,
  "bio" TEXT NOT NULL,
  "birthdate" DATE,
  "role_id" INT NOT NULL,
  "picture_link" VARCHAR
);

CREATE TABLE "role_list" (
  "role_id" SERIAL PRIMARY KEY,
  "role_name" VARCHAR NOT NULL
);

CREATE TABLE "movie_type" (
  "type_id" SERIAL PRIMARY KEY,
  "type_name" VARCHAR NOT NULL
);

CREATE TABLE "user_movie_activity" (
  "user_id" INT NOT NULL,
  "activity_id" INT NOT NULL,
  "movie_id" INT NOT NULL,
  "TIMESTAMP" TIMESTAMP
);

CREATE TABLE "activity_type" (
  "activity_id" SERIAL PRIMARY KEY,
  "activity_type" VARCHAR NOT NULL
);

CREATE TABLE "movie_character" (
  "movie_id" INT NOT NULL,
  "person_id" INT NOT NULL,
  "picture_link" VARCHAR
);

CREATE TABLE "country_list" (
  "country_id" SERIAL PRIMARY KEY NOT NULL,
  "country_name" VARCHAR NOT NULL,
  "flag_link" VARCHAR NOT NULL
);

CREATE TABLE "admin_info" (
  "admin_id" SERIAL PRIMARY KEY,
  "admin_level" INT NOT NULL,
  "admin_full_name" VARCHAR NOT NULL,
  "last_request" TIMESTAMP NOT NULL
);

CREATE TABLE "admin_login" (
  "admin_id" INT NOT NULL,
  "admin_password" VARCHAR NOT NULL,
  "admin_email" VARCHAR NOT NULL
);

CREATE TABLE "admin_information" (
  "admin_id" INT NOT NULL,
  "birthdate" DATE NOT NULL,
  "added_date" TIMESTAMP NOT NULL,
  "fb_link" VARCHAR(255),
  "twitter_link" VARCHAR(255),
  "ig_link" VARCHAR(255),
  "sex" BOOLEAN
);

CREATE TABLE "activity_list" (
  "todo_id" SERIAL PRIMARY KEY NOT NULL,
  "info" TEXT NOT NULL,
  "status_id" INT NOT NULL DEFAULT '0',
  "admin_id" INT,
  "ref_table" INT NOT NULL,
  "ref_table_id" INT NOT NULL,
  "timestamp" TIMESTAMP
);

CREATE TABLE "ref_table_info" (
  "ref_id" SERIAL PRIMARY KEY,
  "tef_table" VARCHAR NOT NULL
);

CREATE TABLE "activity_status" (
  "status_id" SERIAL PRIMARY KEY,
  "info" VARCHAR NOT NULL
);

CREATE TABLE "private_message" (
  "pm_id" SERIAL PRIMARY KEY,
  "sender_id" INT NOT NULL,
  "reply_to" INT DEFAULT 0,
  "message" VARCHAR NOT NULL,
  "timestamp" TIMESTAMP NOT NULL
);

CREATE TABLE "conversation" (
  "pm_id" INT NOT NULL,
  "user1" INT NOT NULL,
  "user2" INT NOT NULL
);

ALTER TABLE "private_message" ADD FOREIGN KEY ("reply_to") REFERENCES "private_message" ("pm_id");

ALTER TABLE "conversation" ADD FOREIGN KEY ("pm_id") REFERENCES "private_message" ("pm_id");

ALTER TABLE "private_message" ADD FOREIGN KEY ("sender_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "conversation" ADD FOREIGN KEY ("user1") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "conversation" ADD FOREIGN KEY ("user2") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "movie_information" ADD FOREIGN KEY ("identifier_id") REFERENCES "identifier" ("identifier_id");

ALTER TABLE "activity_list" ADD FOREIGN KEY ("ref_table") REFERENCES "ref_table_info" ("ref_id");

ALTER TABLE "activity_list" ADD FOREIGN KEY ("status_id") REFERENCES "activity_status" ("status_id");

ALTER TABLE "admin_information" ADD FOREIGN KEY ("admin_id") REFERENCES "admin_info" ("admin_id");

ALTER TABLE "admin_login" ADD FOREIGN KEY ("admin_id") REFERENCES "admin_info" ("admin_id");

ALTER TABLE "eps_information" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "review_comment" ADD FOREIGN KEY ("review_id") REFERENCES "movie_reviews" ("review_id");

ALTER TABLE "review_comment" ADD FOREIGN KEY ("reply_to") REFERENCES "review_comment" ("comment_id");

ALTER TABLE "movie_reviews" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "movie_reviews" ADD FOREIGN KEY ("identifier_id") REFERENCES "identifier" ("identifier_id");

ALTER TABLE "movie_reviews" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "user_movie_rating" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "user_short_info" ADD FOREIGN KEY ("country_id") REFERENCES "country_list" ("country_id");

ALTER TABLE "movie_information" ADD FOREIGN KEY ("country_id") REFERENCES "country_list" ("country_id");

ALTER TABLE "movie_character" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "movie_character" ADD FOREIGN KEY ("person_id") REFERENCES "movie_person" ("person_id");

ALTER TABLE "user_movie_activity" ADD FOREIGN KEY ("activity_id") REFERENCES "activity_type" ("activity_id");

ALTER TABLE "user_movie_activity" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "movie_information" ADD FOREIGN KEY ("type_id") REFERENCES "movie_type" ("type_id");

ALTER TABLE "person_in_movie" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "movie_person" ADD FOREIGN KEY ("role_id") REFERENCES "role_list" ("role_id");

ALTER TABLE "person_in_movie" ADD FOREIGN KEY ("person_id") REFERENCES "movie_person" ("person_id");

ALTER TABLE "movie_tag" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "movie_tag" ADD FOREIGN KEY ("tag_id") REFERENCES "tag_list" ("tag_id");

ALTER TABLE "like_count_movie_comment" ADD FOREIGN KEY ("comment_id") REFERENCES "movie_comment" ("comment_id");

ALTER TABLE "like_count_movie_comment" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "movie_comment" ADD FOREIGN KEY ("identifier_id") REFERENCES "identifier" ("identifier_id");

ALTER TABLE "review_comment" ADD FOREIGN KEY ("identifier_id") REFERENCES "identifier" ("identifier_id");

ALTER TABLE "movie_genre" ADD FOREIGN KEY ("genre_id") REFERENCES "genre_list" ("genre_id");

ALTER TABLE "movie_genre" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "like_count_review_movie" ADD FOREIGN KEY ("review_id") REFERENCES "movie_reviews" ("review_id");

ALTER TABLE "like_count_review_movie" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "like_count_review_comment" ADD FOREIGN KEY ("comment_id") REFERENCES "review_comment" ("comment_id");

ALTER TABLE "like_count_review_comment" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "user_social_activity" ADD FOREIGN KEY ("activity_type") REFERENCES "social_activity_type" ("type_id");

ALTER TABLE "user_social_activity" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "user_information" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "review_comment" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "movie_comment" ADD FOREIGN KEY ("reply_to") REFERENCES "movie_comment" ("comment_id");

ALTER TABLE "movie_comment" ADD FOREIGN KEY ("movie_id") REFERENCES "movie_information" ("movie_id");

ALTER TABLE "movie_comment" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "user_login" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "user_movie_rating" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "user_eps_rating" ADD FOREIGN KEY ("eps_id") REFERENCES "eps_information" ("eps_id");

ALTER TABLE "user_eps_rating" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");
