create database restoran;

use restoran;

create table menu (
	id int not null auto_increment,
    nama varchar(20),
    harga int,
    primary key (id)
);

create table meja (
	id int not null auto_increment,
    status varchar(5),
    primary key (id)
);

create table transaksi (
	id int not null auto_increment,
    meja_id int,
    tanggal datetime,
    notes varchar(30),
    primary key (id),
    foreign key (meja_id) references meja(id)
);

create table pesanan (
	id int not null auto_increment,
    transaksi_id int,
    menu_id int,
    qty int,
    primary key (id),
    foreign key (transaksi_id) references transaksi (id),
    foreign key (menu_id) references menu (id)
);

insert into menu (nama, harga) values
('Nasi goreng', 10000),
('Mie goreng', 15000),
('Teh manis dingin', 5000),
('Kopi hitam', 8000),
('Kopi susu', 10000);

select * from menu;

select * from transaksi;

select * from pesanan;

insert into meja (status) VALUES
('close'), ('close'),
('open'), ('open'),
('close'), ('open'),
('open'), ('open'),
('open'), ('open');

select * from meja;

select nama, qty, harga, qty*harga as total from pesanan ps join menu mn join transaksi tr
on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.meja_id = 3;

select sum(qty*harga) as grandTotal from pesanan ps join menu mn join transaksi tr
on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.meja_id = 3;