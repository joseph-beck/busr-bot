create table busrtest.race_result(
	id           int             not null,
	driver       bigint          not null primary key,
	position     int             null,
	minutes      int             null,
	seconds      int             null,
	milliseconds int             null,
	points       float default 0 null,
	constraint race_result_drivers_id_fk
	    foreign key (driver) references busrtest.drivers (id)
);