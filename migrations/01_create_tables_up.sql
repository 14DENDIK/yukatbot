CREATE TYPE step AS ENUM ('main');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    first_name VARCHAR(255) NULL,
    last_name VARCHAR(255) NULL,
    username VARCHAR(255) NULL,
    language_code VARCHAR(3) NOT NULL,
    current_step step NOT NULL DEFAULT 'main'
);

CREATE TABLE commands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE,
    text JSONB NULL
);


INSERT INTO commands (name, text) VALUES(
    '/start',
    '{"en":"Happy to see you here\nTo see more about this bot type /help command",
    "ru":"Рад виидеть вас здесь\nУзнать больше о возможностях этого бота можно набрав  команду /help",
    "uz":"Sizni k''orishdan hursandman\nBot haqida ko''proq ma''lumotga ega bo''lish uchun /help komandasini tering"}'
);

INSERT INTO commands (name, text) VALUES(
    '/help',
    '{"en":"List of all available commands:\n/start - starts bot\n/help - list of all available commands\n/settings - settings of this bot",
    "ru":"Список всех доступных команд\n/start - запуск бота\n/help - список всех доступных команд\n/settings - настройки этого бота",
    "uz":"Barcha komandalar jadvali\n/start - botni boshlash\n/help - barcha komandalar jadvali\n/settings - botni sozlamalari"}'
);

INSERT INTO commands (name, text) VALUES (
    '/settings',
    '{"en":"Change settings of this bot",
    "ru":"Изменить настройки этого бота",
    "uz":"Bot sozlamalarini o''zgartirish"}'
);