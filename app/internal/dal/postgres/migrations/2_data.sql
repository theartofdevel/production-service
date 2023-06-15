-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

BEGIN;

INSERT INTO public.currency (name, symbol)
VALUES ('рубль', '₽');
INSERT INTO public.currency (name, symbol)
VALUES ('dollar', '$');


INSERT INTO public.category (name)
VALUES ('купоны');
INSERT INTO public.category (name)
VALUES ('цифровые билеты');

COMMIT;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

COMMIT;