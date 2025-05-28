CREATE SCHEMA IF NOT EXISTS blogs;

CREATE TABLE IF NOT EXISTS blogs.blogs (
    id           UUID        PRIMARY KEY,
    user_id      INT         NOT NULL,                            
    campaign_id  UUID        NOT NULL,                            
    content      TEXT        NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
ALTER TABLE blogs.blogs
    ADD CONSTRAINT fk_blog_user
        FOREIGN KEY (user_id)
        REFERENCES users.users (id);

ALTER TABLE blogs.blogs
    ADD CONSTRAINT fk_blog_campaign
        FOREIGN KEY (campaign_id)
        REFERENCES campaigns.campaigns (id);
CREATE INDEX IF NOT EXISTS idx_blog_campaign ON blogs.blogs (campaign_id);
CREATE INDEX IF NOT EXISTS idx_blog_user     ON blogs.blogs (user_id);
