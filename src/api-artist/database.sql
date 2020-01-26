create database artist;

use artist;

create table artist (
    id int not null auto_increment,
    nama varchar(10),
    debut date,
    category varchar(10),
    primary key(id)
);
