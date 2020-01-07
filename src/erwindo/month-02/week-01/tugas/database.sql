CREATE DATABASE cashier;

USE cashier;

CREATE TABLE cashiers (
  id int(11) NOT NULL AUTO_INCREMENT,
  username varchar(20) DEFAULT NULL,
  password varchar(20) DEFAULT NULL,
  name varchar(20) DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE transactions (
  id int(11) NOT NULL AUTO_INCREMENT,
  cashier_id int(11) DEFAULT NULL,
  transaction_date timestamp,
  PRIMARY KEY (id),
  FOREIGN KEY (cashier_id) REFERENCES cashiers (id)
);

CREATE TABLE purchases (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(20) DEFAULT NULL,
  qty int(11) DEFAULT NULL,
  price int(11) DEFAULT NULL,
  transaction_id int(11) DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (transaction_id) REFERENCES transactions (id)
);

INSERT INTO cashiers (username, password, name) values
('kasir1', 'password', 'John Doe'),
('kasir2', 'password', 'Jane Doe');