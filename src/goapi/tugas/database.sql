create database people;

use people;

create table peoples (
	id int not null auto_increment,
	nama varchar(20),
	gender varchar(1),
    primary key (id)
);

insert into peoples (nama, gender) values
("Erwindo", "M"),
("Rifaldo", "M");

select * from peoples;