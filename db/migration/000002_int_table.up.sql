CREATE TABLE IF NOT EXISTS dotask.task (
	id_task bigserial NOT NULL,
 	title_task text NULL,
 	date_task date null,
 	start_task time null,
 	end_task time null,
    category varchar(50) null,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT task_pkey PRIMARY KEY (id_task)
);

CREATE TABLE IF NOT EXISTS dotask.users (
	id_user bigserial not null,
	nama varchar(75) null,
	username varchar(25) null,
	email varchar(50) null,
	telp varchar(25) null,
	tgl_lahir date null,
	alamat text null,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id_user)
);

CREATE TABLE IF NOT EXISTS dotask.friendlist (
	id_friendlist bigserial not null,
	user_id int8 not null,
	friend_id int8 not null,
	status int8 not null default(0),
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT friendlist_pkey PRIMARY KEY (id_friendlist)
);

CREATE TABLE IF NOT EXISTS dotask.events (
	id_event bigserial not null,
	title_event text null,
	date_event date null,
	start_event time null,
	end_event time null,
	location text null,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT event_pkey PRIMARY KEY (id_event)
);

CREATE TABLE IF NOT EXISTS dotask.friendlist_has_join_event (
	event_id int8 null,
	friendlist_id int8 null,
	status int8 null,
	created_at timestamp NULL,
	updated_at timestamp NULL
);