package sql

const createMoviesTable = `
    create table if not exists movies (
        id integer primary key,
        title text,
        original_title text,
        release_date date,
        overview text,
        vote_average real,
        vote_count integer,
        popularity real,
        adult boolean,
        video boolean,
        backdrop_path text,
        poster_path text
);
`

const createMetaTable = `
    create table if not exists funnel_meta (
        id int primary key,
        current_tmdb_page real
    );

    insert into funnel_meta(id, current_tmdb_page) values(1, 1)
    on conflict(id) do nothing;
`
const insertMovies = `
    insert into movies (id, title, original_title, release_date, overview, vote_average, vote_count, popularity, adult, video, backdrop_path, poster_path)
    values (:id, :title, :original_title, :release_date, :overview, :vote_average, :vote_count, :popularity, :adult, :video, :backdrop_path, :poster_path)
    on conflict(id) do nothing;
`

const updateMeta = `
    update funnel_meta
    set current_tmdb_page = :current_tmdb_page
    where id = :id;
`

// TODO: revisit this
const getRandomMovie = `
    select *
    from movies
    where id not in (select movie_key from ratings)
    order by random()
    limit 1;
`

const createSessionTable = `
    create table if not exists sessions (
        user_key text primary key,
        current_movie integer,
        session_status text,
        rating integer
    );
`

const getSessionForUser = `
    select * from sessions
    where user_key = ?
`
const saveSessionForUser = `
    insert into sessions (user_key, current_movie, session_status, rating)
    values (:user_key, :current_movie, :session_status, :rating)
    on conflict(user_key) do update set
        current_movie = excluded.current_movie,
        session_status = excluded.session_status,
        rating = excluded.rating;
`

const createRatingTable = `
    create table if not exists ratings (
        movie_key integer unique,
        rating integer
    );
`

const saveRating = `
    insert into ratings (movie_key, rating)
    values (:movie_key, :rating);
`
