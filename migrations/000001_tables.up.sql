
CREATE TABLE owner(
    fullname        VARCHAR(21)    NOT NULL,   
    password        VARCHAR(60)    NOT NULL,
    role            VARCHAR(5)     DEFAULT 'owner',
    phone_number    VARCHAR(20)    NOT NULL,       
    gmail           VARCHAR(64)    NOT NULL,
    telegram        VARCHAR(64)    NOT NULL,   
    github          VARCHAR(64)    NOT NULL,
    linked_in       VARCHAR(64)    NOT NULL,   
    leetcode        VARCHAR(64)    NOT NULL,
    about_me        TEXT
);

CREATE TABLE  categories(
    category_id     UUID             PRIMARY KEY,
    name            VARCHAR(128)     NOT NULL,
    created_at      TIMESTAMP        DEFAULT current_timestamp
);

CREATE TABLE sub_categories(
    sub_category_id UUID            PRIMARY KEY,
    name            VARCHAR(128)    NOT NULL,
    created_at      TIMESTAMP       DEFAULT current_timestamp,
    category_id     UUID            REFERENCES categories(category_id)    
);

CREATE TABLE articles(
    article_id         UUID             PRIMARY KEY,
    title              VARCHAR          NOT   NULL,
    content            TEXT             NOT   NULL,
    created_at         TIMESTAMP        DEFAULT current_timestamp,
    updated_at         TIMESTAMP        DEFAULT  NULL,
    deleted_at         TIMESTAMP        DEFAULT  NULL,
    category_id        UUID,
    sub_category_id    UUID,
    FOREIGN KEY (category_id)  REFERENCES categories(category_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (sub_category_id) REFERENCES sub_categories(sub_category_id) ON DELETE CASCADE ON UPDATE CASCADE

);

CREATE TABLE viewers(
    viewer_id       UUID  PRIMARY KEY,
    fullname        VARCHAR(64),
    username        VARCHAR(32)  UNIQUE   NOT NULL,
    gmail           VARCHAR(32)  UNIQUE   NOT NULL,
    password        VARCHAR(64)           NOT NULL
);

CREATE TABLE   comments(
    comment_id      UUID   PRIMARY KEY,
    content         TEXT   NOT  NULL,
    created_at      TIMESTAMP           DEFAULT current_timestamp,
    article_id      UUID    REFERENCES  articles(article_id),
    viewer_id       UUID    REFERENCES  viewers(viewer_id)
);

