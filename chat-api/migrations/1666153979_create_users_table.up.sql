create table users
(
    id            int identity
        primary key,
    username      varchar(50),
    name          varchar(50),
    surname       varchar(50),
    email         varchar(100),
    phone         varchar(50),
    password_hash varchar(255),
    role          varchar(50)
)

INSERT INTO dbo.users (username, name, surname, email, phone, password_hash, role) 
VALUES ('dauren77', 'Dauren', 'Chert', 'email1@mail.ru', '87771234455', '6ec435c2d485b35b3b08949ba2999841', 'regular'),
       ('defaultse', 'Yernar', 'Master', 'email2@mail.ru', '87076665544', '93e8bcd480b515cca641cc52b0e1fa07', 'regular'),
       ('gachiman', 'Jotaro', 'Kujo', 'email3@mail.ru', '87001112233', '9746556cd57f32440074c588255fd365', 'regular'),
        