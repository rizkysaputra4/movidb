DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
     
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;

DROP TABLE IF EXISTS 
  user_information, 
  movie_information, 
  eps_information,
  movie_reviews,
  review_comment,
  movie_comment,
  user_movie_rating,
  user_short_info,
  user_eps_rating,
  user_social_activity,
  social_activity_type,
  like_count_review_comment,
  like_count_review_movie,
  like_count_movie_comment,
  movie_genre,
  genre_list,
  identifier,
  movie_tag,
  tag_list,
  person_in_movie,
  movie_person,
  role_list,
  movie_type,
  user_movie_activity,
  activity_type,
  movie_character,
  country_list,
  admin_info,
  admin_login,
  admin_information,
  activity_list,
  ref_table_info,
  activity_status,
  private_message,
  conversation,
  user_notification,
  notif_list;

CREATE TABLE "user_information" (
  "user_id" INT NOT NULL UNIQUE,
  "user_full_name" VARCHAR(80),
  "birthdate" DATE,
  "signup_date" DATE,
  "bio" VARCHAR(150),
  "fb_link" VARCHAR(80),
  "twitter_link" VARCHAR(40),
  "ig_link" VARCHAR(50),
  "sex" BOOLEAN,
  "contributor_points" INT DEFAULT 0,
  "monthly_contributor_points" INT DEFAULT 0
);

CREATE TABLE "movie_information" (
  "movie_id" SERIAL PRIMARY KEY NOT NULL,
  "movie_title" VARCHAR(255) NOT NULL,
  "movie_synopsis" VARCHAR NOT NULL,
  "release_date" DATE,
  "imdb_rating" FLOAT(2),
  "imdb_numb_vote" INT,
  "metacritics" FLOAT(2),
  "site_rating" FLOAT(2) DEFAULT 0,
  "site_numb_vote" INT DEFAULT 0,
  "poster_link" VARCHAR(255),
  "trailer_link" VARCHAR(255),
  "duration" VARCHAR(255),
  "awards" VARCHAR(255),
  "type_id" INT,
  "country_id" VARCHAR(2),
  "language" VARCHAR(255),
  "unique_link" VARCHAR(255) NOT NULL,
  "identifier_id" INT DEFAULT 0,
  "overall_rating" FLOAT(4),

  "popularity" INT DEFAULT 0,
  "daily_popularity" INT DEFAULT 0,
  "weekly_popularity" INT DEFAULT 0,
  "monthly_popularity" INT DEFAULT 0
);

CREATE TABLE "eps_information" (
  "eps_id" SERIAL PRIMARY KEY NOT NULL,
  "movie_id" INT NOT NULL,
  "episode_number" SMALLINT NOT NULL,
  "air_date" DATE NOT NULL,
  "season" SMALLINT NOT NULL,
  "synopsis" TEXT
);

CREATE TABLE "movie_reviews" (
  "review_id" SERIAL PRIMARY KEY NOT NULL,
  "movie_id" INT NOT NULL,
  "eps_id" INT DEFAULT 0,
  "user_id" INT NOT NULL,
  "text" TEXT NOT NULL,
  "story_rating" INT NOT NULL,
  "acting_rating" INT NOT NULL,
  "cinema_rating" INT NOT NULL,
  "music_rating" INT NOT NULL,
  "identifier_id" INT NOT NULL DEFAULT 0,
  "like_count" INT NOT NULL DEFAULT 0,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "review_comment" (
  "comment_id" SERIAL PRIMARY KEY NOT NULL,
  "identifier_id" INT NOT NULL,
  "review_id" INT NOT NULL,
  "reply_to" INT NOT NULL,
  "user_id" INT NOT NULL,
  "text" TEXT NOT NULL,
  "like" INT NOT NULL DEFAULT 0,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "movie_comment" (
  "comment_id" SERIAL PRIMARY KEY NOT NULL,
  "reply_to" INT NOT NULL,
  "movie_id" INT NOT NULL,
  "eps_id" INT DEFAULT 0,
  "text" TEXT NOT NULL,
  "like_count" INT NOT NULL DEFAULT 0,
  "identifier_id" INT NOT NULL DEFAULT 0,
  "user_id" INT NOT NULL,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "user_movie_rating" (
  "movie_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "rating" FLOAT(1) NOT NULL,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "user_short_info" (
  "user_id" SERIAL PRIMARY KEY NOT NULL,
  "user_name" VARCHAR(255) NOT NULL UNIQUE,
  "country_id" VARCHAR(2),
  "password" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL UNIQUE,
  "role" INT DEFAULT 21,
  "last_request" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "user_eps_rating" (
  "eps_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "rating" FLOAT(1) NOT NULL,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "user_social_activity" (
  "user_id" INT NOT NULL,
  "activity_type" INT NOT NULL,
  "activity_table_ref_id" INT
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
  "movie_id" INT NOT NULL,
  "eps_id" INT DEFAULT 0,
  "role_id" INT NOT NULL,
  "info" VARCHAR,
);

CREATE TABLE "movie_person" (
  "person_id" SERIAL PRIMARY KEY,
  "person_name" VARCHAR NOT NULL,
  "bio" TEXT,
  "birthdate" DATE,
  "role_id" INT NOT NULL,
  "picture_link" VARCHAR,

  "popularity" INT DEFAULT 0,
  "daily_popularity" INT DEFAULT 0,
  "weekly_popularity" INT DEFAULT 0,
  "monthly_popularity" INT DEFAULT 0
);

CREATE TABLE "role_list" (
  "role_id" SERIAL PRIMARY KEY,
  "role_name" VARCHAR NOT NULL
);

CREATE TABLE "movie_type" (
  "type_id" INT PRIMARY KEY,
  "type_name" VARCHAR NOT NULL
);

CREATE TABLE "user_movie_activity" (
  "user_id" INT NOT NULL,
  "activity_id" INT NOT NULL,
  "movie_id" INT NOT NULL,
  "TIMESTAMP" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "activity_type" (
  "activity_id" SERIAL PRIMARY KEY,
  "activity_type" VARCHAR NOT NULL
);

CREATE TABLE "movie_character" (
  "movie_id" INT NOT NULL,
  "person_id" INT NOT NULL,
  "eps_id" INT DEFAULT 0,
  "lead" INT DEFAULT 99,
  "character" VARCHAR,
  "info" VARCHAR,
  "picture_link" VARCHAR
);

CREATE TABLE "country_list" (
  "id" SMALLINT,
  "country_name" VARCHAR NOT NULL,
  "country_id" VARCHAR(2) PRIMARY KEY NOT NULL,
  "alpha_3" VARCHAR(3) NOT NULL
);

CREATE TABLE "admin_info" (
  "admin_id" SERIAL PRIMARY KEY,
  "admin_level" SMALLINT NOT NULL,
  "admin_full_name" VARCHAR NOT NULL,
  "last_request" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "admin_login" (
  "admin_id" SMALLINT NOT NULL,
  "admin_password" VARCHAR NOT NULL,
  "admin_email" VARCHAR NOT NULL
);

CREATE TABLE "admin_information" (
  "admin_id" SMALLINT NOT NULL,
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
  "status_id" INT NOT NULL DEFAULT 0,
  "admin_id" INT,
  "ref_table" INT NOT NULL,
  "ref_table_id" INT NOT NULL,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
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
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE "conversation" (
  "pm_id" INT NOT NULL,
  "user1" INT NOT NULL,
  "user2" INT NOT NULL
);

CREATE TABLE "user_notification" (
  "notif_id" INT NOT NULL,
  "user_id" INT NOT NULL,
  "user_2" INT,
  "info" VARCHAR,
  "table_id_reference" INT
);

CREATE TABLE "notif_list" (
  "notif_id" INT PRIMARY KEY NOT NULL,
  "notification_type" VARCHAR NOT NULL
);

ALTER TABLE "movie_reviews" ADD FOREIGN KEY ("eps_id") REFERENCES "eps_information" ("eps_id");

ALTER TABLE "person_in_movie" ADD FOREIGN KEY ("eps_id") REFERENCES "eps_information" ("eps_id");

ALTER TABLE "movie_comment" ADD FOREIGN KEY ("eps_id") REFERENCES "eps_information" ("eps_id");

ALTER TABLE "user_notification" ADD FOREIGN KEY ("notif_id") REFERENCES "notif_list" ("notif_id");

ALTER TABLE "user_notification" ADD FOREIGN KEY ("user_2") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "user_notification" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "person_in_movie" ADD FOREIGN KEY ("role_id") REFERENCES "role_list" ("role_id");

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

ALTER TABLE "user_movie_rating" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");

ALTER TABLE "user_eps_rating" ADD FOREIGN KEY ("eps_id") REFERENCES "eps_information" ("eps_id");

ALTER TABLE "user_eps_rating" ADD FOREIGN KEY ("user_id") REFERENCES "user_short_info" ("user_id");


INSERT INTO country_list (id, country_name, country_id, alpha_3) VALUES
(4,'Afghanistan','af','afg'),
(8,'Albania','al','alb'),
(12,'Algeria','dz','dza'),
(20,'Andorra','ad','and'),
(24,'Angola','ao','ago'),
(28,'Antigua and Barbuda','ag','atg'),
(32,'Argentina','ar','arg'),
(51,'Armenia','am','arm'),
(36,'Australia','au','aus'),
(40,'Austria','at','aut'),
(31,'Azerbaijan','az','aze'),
(44,'Bahamas','bs','bhs'),
(48,'Bahrain','bh','bhr'),
(50,'Bangladesh','bd','bgd'),
(52,'Barbados','bb','brb'),
(112,'Belarus','by','blr'),
(56,'Belgium','be','bel'),
(84,'Belize','bz','blz'),
(204,'Benin','bj','ben'),
(64,'Bhutan','bt','btn'),
(68,'Bolivia (Plurinational State of)','bo','bol'),
(70,'Bosnia and Herzegovina','ba','bih'),
(72,'Botswana','bw','bwa'),
(76,'Brazil','br','bra'),
(96,'Brunei Darussalam','bn','brn'),
(100,'Bulgaria','bg','bgr'),
(854,'Burkina Faso','bf','bfa'),
(108,'Burundi','bi','bdi'),
(132,'Cabo Verde','cv','cpv'),
(116,'Cambodia','kh','khm'),
(120,'Cameroon','cm','cmr'),
(124,'Canada','ca','can'),
(140,'Central African Republic','cf','caf'),
(148,'Chad','td','tcd'),
(152,'Chile','cl','chl'),
(156,'China','cn','chn'),
(170,'Colombia','co','col'),
(174,'Comoros','km','com'),
(178,'Congo','cg','cog'),
(180,'Congo, Democratic Republic of the','cd','cod'),
(188,'Costa Rica','cr','cri'),
(384,'Côte d''Ivoire','ci','civ'),
(191,'Croatia','hr','hrv'),
(192,'Cuba','cu','cub'),
(196,'Cyprus','cy','cyp'),
(203,'Czechia','cz','cze'),
(208,'Denmark','dk','dnk'),
(262,'Djibouti','dj','dji'),
(212,'Dominica','dm','dma'),
(214,'Dominican Republic','do','dom'),
(218,'Ecuador','ec','ecu'),
(818,'Egypt','eg','egy'),
(222,'El Salvador','sv','slv'),
(226,'Equatorial Guinea','gq','gnq'),
(232,'Eritrea','er','eri'),
(233,'Estonia','ee','est'),
(748,'Eswatini','sz','swz'),
(231,'Ethiopia','et','eth'),
(242,'Fiji','fj','fji'),
(246,'Finland','fi','fin'),
(250,'France','fr','fra'),
(266,'Gabon','ga','gab'),
(270,'Gambia','gm','gmb'),
(268,'Georgia','ge','geo'),
(276,'Germany','de','deu'),
(288,'Ghana','gh','gha'),
(300,'Greece','gr','grc'),
(308,'Grenada','gd','grd'),
(320,'Guatemala','gt','gtm'),
(324,'Guinea','gn','gin'),
(624,'Guinea-Bissau','gw','gnb'),
(328,'Guyana','gy','guy'),
(332,'Haiti','ht','hti'),
(340,'Honduras','hn','hnd'),
(348,'Hungary','hu','hun'),
(352,'Iceland','is','isl'),
(356,'India','in','ind'),
(360,'Indonesia','id','idn'),
(364,'Iran (Islamic Republic of)','ir','irn'),
(368,'Iraq','iq','irq'),
(372,'Ireland','ie','irl'),
(376,'Israel','il','isr'),
(380,'Italy','it','ita'),
(388,'Jamaica','jm','jam'),
(392,'Japan','jp','jpn'),
(400,'Jordan','jo','jor'),
(398,'Kazakhstan','kz','kaz'),
(404,'Kenya','ke','ken'),
(296,'Kiribati','ki','kir'),
(408,'Korea (Democratic People''s Republic of)','kp','prk'),
(410,'Korea, Republic of','kr','kor'),
(414,'Kuwait','kw','kwt'),
(417,'Kyrgyzstan','kg','kgz'),
(418,'Lao People''s Democratic Republic','la','lao'),
(428,'Latvia','lv','lva'),
(422,'Lebanon','lb','lbn'),
(426,'Lesotho','ls','lso'),
(430,'Liberia','lr','lbr'),
(434,'Libya','ly','lby'),
(438,'Liechtenstein','li','lie'),
(440,'Lithuania','lt','ltu'),
(442,'Luxembourg','lu','lux'),
(450,'Madagascar','mg','mdg'),
(454,'Malawi','mw','mwi'),
(458,'Malaysia','my','mys'),
(462,'Maldives','mv','mdv'),
(466,'Mali','ml','mli'),
(470,'Malta','mt','mlt'),
(584,'Marshall Islands','mh','mhl'),
(478,'Mauritania','mr','mrt'),
(480,'Mauritius','mu','mus'),
(484,'Mexico','mx','mex'),
(583,'Micronesia (Federated States of)','fm','fsm'),
(498,'Moldova, Republic of','md','mda'),
(492,'Monaco','mc','mco'),
(496,'Mongolia','mn','mng'),
(499,'Montenegro','me','mne'),
(504,'Morocco','ma','mar'),
(508,'Mozambique','mz','moz'),
(104,'Myanmar','mm','mmr'),
(516,'Namibia','na','nam'),
(520,'Nauru','nr','nru'),
(524,'Nepal','np','npl'),
(528,'Netherlands','nl','nld'),
(554,'New Zealand','nz','nzl'),
(558,'Nicaragua','ni','nic'),
(562,'Niger','ne','ner'),
(566,'Nigeria','ng','nga'),
(807,'North Macedonia','mk','mkd'),
(578,'Norway','no','nor'),
(512,'Oman','om','omn'),
(586,'Pakistan','pk','pak'),
(585,'Palau','pw','plw'),
(591,'Panama','pa','pan'),
(598,'Papua New Guinea','pg','png'),
(600,'Paraguay','py','pry'),
(604,'Peru','pe','per'),
(608,'Philippines','ph','phl'),
(616,'Poland','pl','pol'),
(620,'Portugal','pt','prt'),
(634,'Qatar','qa','qat'),
(642,'Romania','ro','rou'),
(643,'Russian Federation','ru','rus'),
(646,'Rwanda','rw','rwa'),
(659,'Saint Kitts and Nevis','kn','kna'),
(662,'Saint Lucia','lc','lca'),
(670,'Saint Vincent and the Grenadines','vc','vct'),
(882,'Samoa','ws','wsm'),
(674,'San Marino','sm','smr'),
(678,'Sao Tome and Principe','st','stp'),
(682,'Saudi Arabia','sa','sau'),
(686,'Senegal','sn','sen'),
(688,'Serbia','rs','srb'),
(690,'Seychelles','sc','syc'),
(694,'Sierra Leone','sl','sle'),
(702,'Singapore','sg','sgp'),
(703,'Slovakia','sk','svk'),
(705,'Slovenia','si','svn'),
(90,'Solomon Islands','sb','slb'),
(706,'Somalia','so','som'),
(710,'South Africa','za','zaf'),
(728,'South Sudan','ss','ssd'),
(724,'Spain','es','esp'),
(144,'Sri Lanka','lk','lka'),
(729,'Sudan','sd','sdn'),
(740,'Suriname','sr','sur'),
(752,'Sweden','se','swe'),
(756,'Switzerland','ch','che'),
(760,'Syrian Arab Republic','sy','syr'),
(762,'Tajikistan','tj','tjk'),
(834,'Tanzania, United Republic of','tz','tza'),
(764,'Thailand','th','tha'),
(626,'Timor-Leste','tl','tls'),
(768,'Togo','tg','tgo'),
(776,'Tonga','to','ton'),
(780,'Trinidad and Tobago','tt','tto'),
(788,'Tunisia','tn','tun'),
(792,'Turkey','tr','tur'),
(795,'Turkmenistan','tm','tkm'),
(798,'Tuvalu','tv','tuv'),
(800,'Uganda','ug','uga'),
(804,'Ukraine','ua','ukr'),
(784,'United Arab Emirates','ae','are'),
(826,'United Kingdom of Great Britain and Northern Ireland','gb','gbr'),
(840,'United States of America','us','usa'),
(858,'Uruguay','uy','ury'),
(860,'Uzbekistan','uz','uzb'),
(548,'Vanuatu','vu','vut'),
(862,'Venezuela (Bolivarian Republic of)','ve','ven'),
(704,'Viet Nam','vn','vnm'),
(887,'Yemen','ye','yem'),
(894,'Zambia','zm','zmb'),
(716,'Zimbabwe','zw','zwe');

INSERT INTO movie_type (type_id, type_name) VALUES
(1, 'Movie'),
(2, 'TV Series'),
(3, 'KDrama'),
(4, 'Dorama'),
(5, 'Anime');

INSERT INTO identifier (identifier_id, info) VALUES
(1, 'On Display'),
(2, 'On Review'),
(3, 'Need Review'),
(4, 'Hidden'),
(5, 'Reported'),
(6, 'Deleted');

INSERT INTO role_list (role_name) VALUES
('Actress'),
('Actor'),
('Director'),
('Screen Writer'),
('Producer'),
('Music'),
('Cinematography'),
('Editing'),
('Casting'),
('Production Designer'),
('Art Director'),
('Set Decoration'),
('Make Up'),
('Production Manager'),
('Assistant Director'),
('Art / PropMaker'),
('Sound Mixer'),
('Special Effect Technician'),
('Visual Effect'),
('Stunt'),
('Animator'),
('Costume Designer');