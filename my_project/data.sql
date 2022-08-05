DROP TABLE IF EXISTS author CASCADE;
DROP TABLE IF EXISTS book CASCADE;
DROP TABLE IF EXISTS book_authors CASCADE;

CREATE TABLE public.author
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    born DATE,
    death DATE

);

CREATE TABLE public.book
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    date  INT,
    about VARCHAR(1000) NOT NULL,
    author_id UUID NOT NULL




);

CREATE TABLE public.book_authors
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id   UUID NOT NULL,
    author_id UUID NOT NULL,

    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book (id),
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author (id),
    CONSTRAINT book_author_unique UNIQUE (book_id, author_id)
);
DELETE FROM author WHERE name = 'Айзек Азимов';

INSERT INTO author (name, born, death) VALUES ('Айзек Азимов', '1920-01-02', '1992-04-06' );
INSERT INTO author (name, born, death) VALUES ('Джон Рональд Руэл Толкин', '1892-01-03', '1973-09-02');
INSERT INTO author (name, born, death) VALUES ('Джоан Роулинг', '1965-07-31', null)

    INSERT INTO book (name, date, about, author_id) VALUES ('Стальные Пещеры', '1954-01-01', 'Действие происходит в далёком будущем. Роботехники Хэн Фастольф и Родж Наменну Сартон занимались конструированием человекоподобных роботов. Сартоном был создан Р. Дэниел Оливо. Завязка сюжета — убийство Сартона в собственном доме с помощью бластера. Во избежание скандала, полиции пришлось выделить для расследования своего сотрудника Элайджа Бейли. Космониты выделяют ему в напарники Р. Дэниела Оливо', '8f24207e-8c14-4ed7-ba75-e263f7f70b9e' );
INSERT INTO book (name, date, about, author_id) VALUES ('Лжец!', '1941-05-01', 'Из-за ошибки, допущенной во время конструирования робота РБ-34, модель обретает телепатические способности. Работники U.S. Robots and Mechanical Men, Inc. исследуют робота, чтобы понять, где именно была допущена ошибка.', '8f24207e-8c14-4ed7-ba75-e263f7f70b9e');
INSERT INTO book (name, date, about, author_id) VALUES ('Хоббит, или Туда и обратно','1937-09-21', 'Состоятельный хоббит Бильбо Бэггинс ведёт размеренную жизнь, но однажды ему наносят визит волшебник Гэндальф и компания из тринадцати гномов. С тех пор спокойная жизнь у Бильбо заканчивается и начинается Приключение.', 'f8b594f6-8aa8-4724-8bfe-79a400d98479');
INSERT INTO book (name, date, about, author_id) VALUES ('Властелин колец', '1954-07-29', 'Сюжет «Властелин колец» привязан к событиям повести «Хоббит» и является её продолжением; действие начинается примерно через 60 лет после окончания событий «Хоббита»', 'f8b594f6-8aa8-4724-8bfe-79a400d98479');
INSERT INTO book (name, date, about, author_id) VALUES ('Гарри Поттер и философский камень', '1997-06-26', 'Гарри Поттер живет с дядей, тетей и двоюродным братом в пригороде Лондона. Он абсолютно уверен, что он самый обычный мальчик, однако вскоре он получает письмо из Школы Чародейства и Волшебства - Хогвартс', 'bed4359a-17bf-4e41-9c63-19a091b0b561');
INSERT INTO book (name, date, about, author_id) VALUES ('Гарри Поттер и Тайная комната', '1998-01-01', 'Продолжение истории о мальчике с шрамом в виде молнии. Три ученика факультета Гриффиндор снова нарушают школьные правила, сражаются с чудовищами и варят запрещенные зелья в туалете для девочек', 'bed4359a-17bf-4e41-9c63-19a091b0b561')

    INSERT INTO book_authors (book_id, author_id) VALUES ('89df023a-b935-45e8-a5fa-e26d7aa773c2', '8f24207e-8c14-4ed7-ba75-e263f7f70b9e');
INSERT INTO book_authors (book_id, author_id) VALUES ('e06f51f6-b37e-41da-a275-1867d906892f', '8f24207e-8c14-4ed7-ba75-e263f7f70b9e');
INSERT INTO book_authors (book_id, author_id) VALUES ('8cb34f31-0c76-4f7b-9179-2e851e0c0129', 'f8b594f6-8aa8-4724-8bfe-79a400d98479');
INSERT INTO book_authors (book_id, author_id) VALUES ('2e717892-e405-4928-967e-d06072d6601b', 'f8b594f6-8aa8-4724-8bfe-79a400d98479');
INSERT INTO book_authors (book_id, author_id) VALUES ('875b67a0-ddb0-402b-999d-3be362dfb4ff', 'bed4359a-17bf-4e41-9c63-19a091b0b561');
INSERT INTO book_authors (book_id, author_id) VALUES ('5ed0fa8c-d041-4d2c-a35f-2dbce61fa680', 'bed4359a-17bf-4e41-9c63-19a091b0b561')