package tables

const qualifying_result string = `
	create table busrtest.qualifying_result(
    	id           int             not null,
    	driver       bigint          not null primary key,
    	position     int             not null,
    	minutes      int   default 0 null,
    	seconds      int   default 0 null,
    	milliseconds int   default 0 null,
    	points       float default 0 null,
    	constraint driver
    	    unique (driver),
    	constraint driver
    	    foreign key (driver) references busrtest.drivers (id)
	);
`