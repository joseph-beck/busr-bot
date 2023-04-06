package sql



const driversTable string = `
	create table drivers(
		id int primary key not null,
		name text not null,
		university text not null,
		wins int,
		poles int,
		podiums int,
		points float,
		starts int,
		avg_quali float,
		TODO: CHAMPIONSHIPS
	)
`

const seasonTable string = `
	create table season(
		id string primary key not null,
		name text not null,
		champion int foreign key,
		TODO: RACES
	)
`

const raceTable string = `
	create table race(
		id string primary key not null,
		name text not null,
		laps int not null,
		winner int foreign key
	)
`