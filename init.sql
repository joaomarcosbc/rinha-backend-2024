CREATE TABLE IF NOT EXISTS clientes (
    id SERIAL PRIMARY KEY NOT NULL,
    nome VARCHAR(50) NOT NULL,
    limite INTEGER NOT NULL,
    saldo INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS transacoes (
    id SERIAL PRIMARY KEY NOT NULL,
    tipo CHAR(1) NOT NULL,
    descricao VARCHAR(10) NOT NULL,
    valor INTEGER NOT NULL,
    cliente_id INTEGER NOT NULL,
    realizada_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_cliente_id ON transacoes (cliente_id);

INSERT INTO clientes (nome, limite, saldo)
VALUES
    ('Joao', 100000, 0),
    ('Juba', 1000000, 0),
    ('Lau', 80000, 0),
    ('Yasmin', 10000000, 0),
    ('Lis', 500000, 0);