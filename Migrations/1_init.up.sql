begin;
create table genre (
    id serial primary key,
    name varchar(200) not null
);
create table book (
    id serial primary key,
    name varchar(200) not null,
    author varchar(200) not null,
    publisher varchar(200) not null
);
create table bookgenre (
    bookId integer references book (id),
    genreId integer references genre (id)
);
create table bookinstance (
    id serial primary key,
    originalId integer references book (id)
);
create table client (
    id serial primary key,
    login varchar(200) not null,
    password varchar(200) not null,
    admin boolean not null
);
create table bookselect (
    clientId integer references client (id),
    bookId integer references book (id)
);
create table bookrent (
    clientId integer references client (id),
    instanceBookId integer references book (id),
    requestDate timestamp not null,
    startRentDate timestamp,
    deadline timestamp not null
);
create table renthistory(
    clientId integer references client (id),
    bookId integer references book (id),
    startRentDate timestamp not null
);
create table sessions(
    sessionKey integer primary key,
    clientId integer references client (id), 
    deadline timestamp
);
create view topgenres as
    select 
        genre.name as name,
        count(bookrent.instanceBookId) as rentCount
    from 
        genre, book, bookrent, bookinstance, bookgenre
    where 
        bookrent.instanceBookId = bookinstance.id and 
        bookinstance.originalId = book.id and
        bookgenre.bookId = book.id and
        bookgenre.genreId = genre.id
    group by 
        genre.name
    order by 
        count(bookrent.instanceBookId) desc limit 6;
create view bestseller as
    select 
        bookinstance.originalId as originalId,
        book.name as name,
        book.author as author,
        count(bookinstance.id) as rentCount
    from 
        bookrent, bookinstance, book
    where
        bookrent.startRentDate is not null and
        extract(month from bookrent.startRentDate) =  extract(month from now()) and
        bookrent.instanceBookId = bookinstance.id and 
        book.id = bookinstance.originalId
    group by 
        bookinstance.originalId,
        book.name,
        book.author
    order by count(bookinstance.id) desc limit 1;

insert into genre (name) values
    ('Классика'),
    ('Русская'),
    ('Зарубежная'),
    ('Детская'),
    ('Детективы'),
    ('Фантастика');

commit;