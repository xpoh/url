/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET TIME_ZONE = '+00:00' */;

CREATE DATABASE IF NOT EXISTS butty;
CREATE USER IF NOT EXISTS butty@'%' IDENTIFIED BY 'butty';
GRANT ALL PRIVILEGES ON butty.* TO butty@'%';

FLUSH PRIVILEGES;

USE butty;

create table urls
(
    id         int auto_increment,
    btUrl      varchar(1024) not null,
    url        varchar(1024) not null,
    createDate datetime      not null,
    constraint urls_pk
        primary key (id)
)
    comment 'Main table with butty links';

create unique index urls_id_uindex
    on urls (id);

create index urls_btUrl_index
    on urls (btUrl);

