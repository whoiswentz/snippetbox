CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE snippetbox;

CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';

create table snippets
(
    id      integer      not null primary key auto_increment,
    title   varchar(100) not null,
    content text         not null,
    created datetime     not null,
    expires datetime     not null
);

create index idx_snippets_created on snippets (created);

INSERT INTO snippets (title, content, created, expires)
VALUES ('An old silent pond',
        'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY));
INSERT INTO snippets (title, content, created, expires)
VALUES ('Over the wintry forest',
        'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY));
INSERT INTO snippets (title, content, created, expires)
VALUES ('First autumn morning',
        'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY));

create table users
(
    id              integer      not null primary key auto_increment,
    name            varchar(255) not null,
    email           varchar(255) not null,
    hashed_password char(60)     not null,
    created         datetime     not null,
    active          boolean      not null default true
);

alter table users
    add constraint users_uniq_email unique (email);