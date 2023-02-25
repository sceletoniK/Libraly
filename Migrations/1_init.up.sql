begin;
create table genre(
    id serial primary key,
    name varchar(200) not null
);
create table book(
    id serial primary key,
    name varchar(200) not null,
    author varchar(200) not null,
    publisher varchar(200) not null
);
create table bookgenre(
    bookId integer references book (id),
    genreId integer references genre (id)
);
create table bookinstance(
    id serial primary key,
    originalId integer references book (id)
);
create table user(
    id serial primary key,
    login varchar(200) not null,
    password varchar(200) not null,
    admin boolean not null
);
create table bookselect(
    userId integer references user (id),
    bookId integer references book (id)
);
create table bookrent(
    userId integer references user (id),
    instanceBookId integer references book (id),
    start timestamp,
    end timestamp,
    deadline timestamp not null
);
create view topgenres as
    select 
        genre.name as name,
        count(bookrent.bookId) as rentCount
    from 
        genre, book, bookrent, bookinstance, bookgenre
    where 
        bookrent.bookId = bookinstance.id and 
        bookinstance.originalId = book.id and
        bookgenre.bookId = book.id and
        bookgenre.genreId = genre.id
    group by 
        genre.name
    order by 
        count(bookrent.bookId) desc limit 6;
create view bestseller as
    select 
        bookinstance.originalId as originalId,
        book.name as name,
        book.author as author,
        count(bookinstance.id) as rentCount
    from 
        bookrent, bookinstance, book
    where
        bookrent.start is not null and
        date_part ("month", bookrent.start) =  date_part ("month", CURRENT_TIMESTAMP) and
        bookrent.bookId = bookinstance.id and 
        book.id = bookinstance.originalId
    group by 
        bookinstance.originalId,
        book.name,
        book.author
    order by count(bookinstance.id) desc limit 1;
commit;