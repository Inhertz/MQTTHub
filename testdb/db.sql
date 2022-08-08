create database TitDB;
go
use TitDB;
go
create table landholdings(id int primary key identity(1,1),
 landholding_name varchar(50) not null,
  landholding_address varchar(50),
   landholding_owner varchar(50));
go
create table greenhouse_sizes(id int primary key identity(1,1), area decimal(6,2) not null, unit varchar(50) not null);
go
create table greenhouses(id int primary key identity(1,1), designation varchar(50) not null, id_size int not null foreign key references greenhouse_sizes(id) on update cascade on delete cascade, id_landholding int not null foreign key references landholdings(id) on update cascade on delete cascade);
go
create table sensor_types(id int primary key identity(1,1), measurable varchar(50) not null, brand varchar(50) not null, model varchar(50) not null);
go
create table sensor_positions(id int primary key identity(1,1), x_relative decimal(6,2) not null, y_relative decimal(6,2) not null, unit varchar(50) not null);
go
create table sensors(id int primary key identity(1,1), device_add varchar(50) not null, id_position int not null foreign key references sensor_positions(id) on update cascade on delete cascade, id_sensor_type int not null foreign key references sensor_types(id) on update cascade on delete cascade, id_greenhouse int not null foreign key references greenhouses(id) on update cascade on delete cascade);
go
create table measurements(id int primary key identity(1,1), measurement_value decimal(5,3) not null, unit varchar(50) not null, measurement_date datetime not null, id_sensor int not null foreign key references sensors(id) on update cascade on delete cascade);
go
insert into landholdings(landholding_name, landholding_address, landholding_owner) values ('La Tormenta', 'Machachi', 'Juan Pérez');
go
insert into landholdings(landholding_name, landholding_address, landholding_owner) values ('Pinos Susurrantes', 'Cayambe', 'Marcela Muñoz');
go
insert into greenhouse_sizes(area, unit) values (100, 'metros cuadrados');
go
insert into greenhouse_sizes(area, unit) values (200, 'metros cuadrados');
go
insert into greenhouses(designation, id_size, id_landholding) values ('Bloque A', 1, 1);
go
insert into greenhouses(designation, id_size, id_landholding) values ('Bloque B', 2, 2);
go
insert into sensor_types(measurable, brand, model) values ('temperatura', 'Texas Instruments', 'DS18B20');
go
insert into sensor_positions(x_relative, y_relative, unit) values (50,50, 'metros');
go
insert into sensors(device_add, id_sensor_type, id_greenhouse, id_position) values ('ADD123', 1,1,1);
go
insert into measurements(measurement_value, unit, measurement_date, id_sensor) values (20, 'Celcius', '2022-06-06 00:00:00', 1);
go