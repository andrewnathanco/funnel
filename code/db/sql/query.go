package sql

const createBlackList string = `
create table if not exists black_list (
	answer_key string not null primary key
);`

const createGreenList string = `
create table if not exists green_list (
	answer_key string not null primary key
);`

const createYellowList string = `
create table if not exists yellow_list (
	answer_key string not null primary key
);`