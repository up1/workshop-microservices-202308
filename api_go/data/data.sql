create table message
(
    id  integer not null primary key,
    content varchar(100) not null
);

INSERT into message (id, content) values (1, 'Hello world from database');