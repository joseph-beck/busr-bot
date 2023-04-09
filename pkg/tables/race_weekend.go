package tables

const race_weekend string = `
	create table busrtest.race_weekend(
	    id         int not null primary key,
	    track      int not null,
	    sprint     int null,
	    race       int not null,
	    qualifying int not null
	);
`
