create table busrtest.sprint_result(
	id           int             not null,
	driver       bigint          not null primary key,
	position     int             null,
	minutes      int   default 0 null,
	seconds      int   default 0 null,
	milliseconds int   default 0 null,
	points       float default 0 null,
	constraint sprint_result_drivers_id_fk
	    foreign key (driver) references busrtest.drivers (id)
);
