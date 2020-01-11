# Go API Aplikasi pemesanan restoran

## Using Postman

`http://localhost:8080/`

### See List of Table/Meja

Method: `GET`
URL: `http://localhost:8080/meja/list`

### See List of Menu

Method: `GET`
URL: `http://localhost:8080/menu/list`

### Open a table for order

Method: `PUT`
URL: `http://localhost:8080/meja/1`

**Param** is a number for mejaID

### Insert or Make a Order/Transaction

Method: `POST`
URL: `http://localhost:8080/transaksi`

**Body**

```
[
    {
        "MejaID" : 1,
        "Notes" : "Nasinya pedas dan teh tidak terlalu manis.",
        "Pesan" :
        [
            {
                "MenuID" : 1,
                "Qty": 2,
            },
            {
                "MenuID" : 2,
                "Qty": 2,
            }
        ] 
    }
]
```

### Create a Billing or Close a Table

Method: `PUT`
URL: `http://localhost:8080/transaksi/1`

**Param** is a number for Transaksi ID

## Database Query

```
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

insert into meja (status) VALUES
('close'), ('close'),
('open'), ('open'),
('close'), ('open'),
('open'), ('open'),
('open'), ('open');

select * from meja;
select * from menu;
select * from transaksi;
select * from pesanan;
```