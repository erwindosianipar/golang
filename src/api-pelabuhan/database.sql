create database pelabuhan;

use pelabuhan;

create table kapal (
	id int not null auto_increment,
    kode varchar(5),
    muatan int,
    status varchar(10),
    is_delete tinyint default 0,
    primary key (id)
);

create table dock (
	id int not null auto_increment,
    kode varchar(5),
    status varchar(10),
    is_delete tinyint default 0,
    primary key (id)
);

create table transaksi (
	id int not null auto_increment,
    kapal_id int,
    dock_id int,
    tanggal_masuk date,
    tanggal_keluar date,
    is_delete tinyint default 0,
    primary key (id),
    foreign key (kapal_id) references kapal(id),
    foreign key (dock_id) references dock(id)
);

create table biaya (
	id int not null auto_increment,
    biaya_perhari int,
    biaya_muatan int,
    primary key (id)
);

select * from kapal;
select * from dock;
select * from transaksi;
select * from biaya;