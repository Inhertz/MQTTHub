use [master]
go
create database [TitDB];
go
use [TitDB];
go

create table [dbo].[landholdings]([id] int primary key identity(1,1),
 [name] varchar(200) not null,
  [address] varchar(200),
   [owner] varchar(200));
go
create table [dbo].[greenhouse_sizes]([id] int primary key identity(1,1), 
 [x_max] decimal(6,2) not null,
  [y_max] decimal(6,2) not null,
   [area] decimal(6,2) not null,
    [unit] varchar(200) not null);
go
create table [dbo].[greenhouses]([id] int primary key identity(1,1),
 [designation] varchar(200) not null,
  [id_size] int not null foreign key references [greenhouse_sizes]([id]) on update cascade on delete cascade,
   [id_landholding] int not null foreign key references [landholdings]([id]) on update cascade on delete cascade);
go
create table [dbo].[sensor_types]([id] int primary key identity(1,1),
 [measurable] varchar(200) not null,
  [brand] varchar(200) not null,
   [model] varchar(200) not null,
    [image_link] varchar(max));
go
create table [dbo].[sensor_positions]([id] int primary key identity(1,1),
 [x_relative] decimal(6,2) not null,
  [y_relative] decimal(6,2) not null,
   [unit] varchar(200) not null);
go
create table [dbo].[sensors]([id] int primary key identity(1,1),
 [device_add] varchar(200) not null,
  [id_position] int not null foreign key references [sensor_positions]([id]) on update cascade on delete cascade,
   [id_sensor_type] int not null foreign key references [sensor_types]([id]) on update cascade on delete cascade,
    [id_greenhouse] int not null foreign key references [greenhouses]([id]) on update cascade on delete cascade);
go
create table [dbo].[measurements]([id] int primary key identity(1,1),
 [value] decimal(5,3) not null,
  [unit] varchar(200) not null,
   [date] datetime not null,
    [id_sensor] int not null foreign key references [sensors]([id]) on update cascade on delete cascade);
go

create table [dbo].[users]([id] int primary key identity(1,1),
 [username] varchar(200) not null,
  [password_hash] binary(64) not null,
   [password_salt] varbinary(max) not null,
    [role] varchar(200) not null);
go

insert into [landholdings]([name], [address], [owner]) values ('La Tormenta', 'Machachi', 'Juan Pérez');
go
insert into [landholdings]([name], [address], [owner]) values ('Pinos Susurrantes', 'Cayambe', 'Marcela Muñoz');
go
insert into [greenhouse_sizes]([area], [unit], [x_max], [y_max]) values (100, 'metros cuadrados', 10, 10);
go
insert into [greenhouse_sizes]([area], [unit], [x_max], [y_max]) values (200, 'metros cuadrados', 20, 10);
go
insert into [greenhouses]([designation], [id_size], [id_landholding]) values ('Bloque A', 1, 1);
go
insert into [greenhouses]([designation], [id_size], [id_landholding]) values ('Bloque B', 2, 2);
go
insert into [sensor_types]([measurable], [brand], [model], [image_link]) values ('Temperatura de Suelo', 'Texas Instruments', 'DS18B20', 'https://i.pinimg.com/736x/61/24/df/6124df8ae6b5acfad551a1302d6d4092.jpg');
go
insert into [sensor_types]([measurable], [brand], [model], [image_link]) values ('Temperatura de Aire', 'Generico', 'DHT21' , 'https://www.kuongshun-ks.com/uploads/201810680/dht21-am2301-capacitive-digital44321655909.jpg');
go
insert into [sensor_types]([measurable], [brand], [model], [image_link]) values ('Humedad de Aire', 'Generico', 'DHT21', 'https://www.kuongshun-ks.com/uploads/201810680/dht21-am2301-capacitive-digital44321655909.jpg');
go
insert into [sensor_types]([measurable], [brand], [model], [image_link]) values ('Humedad de Suelo', 'Generico', 'Capacitive_Soil_Moisture_Sensor_1.2', 'https://www.beemong.com/wp-content/uploads/2020/07/FA3003-1-1.jpg');
go
insert into [sensor_positions]([x_relative], [y_relative], [unit]) values (5,5, 'metros');
go
insert into [sensors]([device_add], [id_sensor_type], [id_greenhouse], [id_position]) values ('ADD123', 1,1,1);
go
insert into [sensors]([device_add], [id_sensor_type], [id_greenhouse], [id_position]) values ('ADD127', 1,1,1);
go
insert into [sensors]([device_add], [id_sensor_type], [id_greenhouse], [id_position]) values ('ADD124', 2,1,1);
go
insert into [sensors]([device_add], [id_sensor_type], [id_greenhouse], [id_position]) values ('ADD125', 3,1,1);
go
insert into [sensors]([device_add], [id_sensor_type], [id_greenhouse], [id_position]) values ('ADD126', 4,1,1);
go
insert into [measurements]([value], [unit], [date], [id_sensor]) values (20, 'Celcius', '2021-06-06 00:00:00', 1);
go