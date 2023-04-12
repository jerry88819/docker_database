CREATE ROLE jerry LOGIN PASSWORD 'jerry123';

CREATE DATABASE dc_database;
ALTER DATABASE dc_database OWNER TO jerry;
ALTER DATABASE dc_database SET timezone TO 'UTC';
REVOKE USAGE ON SCHEMA public FROM PUBLIC;
REVOKE CREATE ON SCHEMA public FROM PUBLIC;
GRANT USAGE ON SCHEMA public to jerry;
GRANT CREATE ON SCHEMA public to jerry;

\connect dc_database;
--  前面那個斜線不加會有錯

create table account_info (
   account VARCHAR (50) NOT NULL,
   password VARCHAR (50) NOT NULL
);

create table message_board_data (
   user_name VARCHAR (50) NOT NULL,
   message VARCHAR (50) NOT NULL,
   time timestamp with time zone not null default current_timestamp
);

ALTER TABLE account_info OWNER TO jerry;
ALTER TABLE message_board_data OWNER TO jerry;
-- GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLE account_info to dcard_user;
-- GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLE message_board_data to dcard_user;

\connect dc_database;

insert into account_info(account, password)
values('0412', '0412_aaa');

insert into account_info(account, password)
values('test', 'test_aaa');

insert into message_board_data(user_name, message)
values('0412', 'Today is a happy day!')