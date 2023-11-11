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

const createFunnelSessions string = `
	create table if not exists funnel_sessions (
	user_key string not null primary key,
	session string not null
);`

const getMovieByDecadeQuery string = `
select answer_data,
(
	CAST(JSON_EXTRACT(answer_data, '$.vote_average') AS DECIMAL(10,2)) + 
	CAST(JSON_EXTRACT(answer_data, '$.vote_count') AS DECIMAL(10,2)) + 
	CAST(JSON_EXTRACT(answer_data, '$.popularity') AS DECIMAL(10,2)) 
) / 3 AS answer_average
from answers
where (cast(substr(json_extract(answer_data, '$.release_date'), 1, 4) as int) / 10) * 10 like ?
and answer_average > 500 and answer_average < 2000
order by random()
`
const getSessionByKeyQuery string = `
select session
from funnel_sessions 
where user_key = ?
`

const upsertBoardQuery string = `
    insert or replace into funnel_sessions (user_key, session)
    values (?, ?)
`
const getNumberPerDecadeQuery string = `
select count(answer_key),
(
	CAST(JSON_EXTRACT(answer_data, '$.vote_average') AS DECIMAL(10,2)) + 
	CAST(JSON_EXTRACT(answer_data, '$.vote_count') AS DECIMAL(10,2)) + 
	CAST(JSON_EXTRACT(answer_data, '$.popularity') AS DECIMAL(10,2)) 
) / 3 AS answer_average
from answers
where 
answer_key not in (
    select answer_key
    from red_list
) 
and 
answer_key not in (
    select answer_key
    from green_list
)
and
answer_key not in (
    select answer_key
    from yellow_list
)
and 
answer_key not in (
    select answer_key
    from yellow_list
)
and (cast(substr(json_extract(answer_data, '$.release_date'), 1, 4) as int) / 10) * 10 like ?
and answer_average > 500 and answer_average < 2000
`

const greenListMovie string = `
    insert or ignore into green_list (answer_key)
    values (?)
`

const yellowListMovie string = `
    insert or ignore into yellow_list (answer_key)
    values (?)
`

const blackListMovie string = `
    insert or ignore into black_list (answer_key)
    values (?)
`