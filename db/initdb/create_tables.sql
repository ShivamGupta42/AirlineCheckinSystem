drop table if exists users
drop table if exists plane
drop table if exists seats
drop table if exists journey

create table users (
    user_id int primary key,
    user_name varchar
)

create table plane(
    plane_id int primary key
)

create table seats(
    seat_id int primary key
)

create table journey(
    id int primary key,
    plane_id int,
    seat_id int,
    user_id int,
    status varchar,
    CONSTRAINT fk_plane FOREIGN KEY(plane_id) REFERENCES plane(plane_id),
    CONSTRAINT fk_seats FOREIGN KEY(seat_id) REFERENCES seats(seat_id),
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(user_id)
)