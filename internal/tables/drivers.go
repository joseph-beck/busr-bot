package tables

const drivers string = `
	create table busrtest.drivers(
    	id         bigint not null primary key,
    	name       text   not null,
    	university text   not null,
    	constraint id
    	    unique (id)
	);
`