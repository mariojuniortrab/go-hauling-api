create table users
(
  ID varchar(255),
  name varchar(255),
  birth DATE,
  email varchar(255),
  active boolean DEFAULT 1,
  password  varchar(255),
  PRIMARY KEY (ID)
);


create table brands
(
  ID varchar(255),
  name varchar(255),
  PRIMARY KEY (ID)
);
