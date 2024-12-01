CREATE SCHEMA IF NOT EXISTS gerawana;
SET search_path TO gerawana;

CREATE TABLE IF NOT EXISTS resetpassword (
  id VARCHAR(191) PRIMARY KEY,
  expires_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  userId VARCHAR(191) NOT NULL,
  otp VARCHAR(191) NOT NULL
);

CREATE TABLE IF NOT EXISTS session (
  token VARCHAR(191) PRIMARY KEY,
  expires_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  userId VARCHAR(191) NOT NULL
);

CREATE TABLE IF NOT EXISTS shortlink (
  id VARCHAR(191) PRIMARY KEY,
  redirect VARCHAR(191) NOT NULL,
  link VARCHAR(191) NOT NULL,
  authorId VARCHAR(191) NOT NULL,
  name VARCHAR(191) NOT NULL,
  createdAt TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user" (
  id VARCHAR(191) PRIMARY KEY,
  email VARCHAR(191) NOT NULL,
  password VARCHAR(191) NOT NULL,
  username VARCHAR(191) NOT NULL
);

CREATE TABLE IF NOT EXISTS visit (
  id VARCHAR(191) PRIMARY KEY,
  shortLinkId VARCHAR(191) NOT NULL,
  visitedAt TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add constraints after table declarations
ALTER TABLE resetpassword
  ADD CONSTRAINT ResetPassword_userId_key UNIQUE (userId),
  ADD CONSTRAINT ResetPassword_otp_key UNIQUE (otp),
  ADD CONSTRAINT ResetPassword_userId_fkey FOREIGN KEY (userId) REFERENCES "user" (id) ON UPDATE CASCADE;

ALTER TABLE session
  ADD CONSTRAINT Session_userId_fkey FOREIGN KEY (userId) REFERENCES "user" (id) ON UPDATE CASCADE;

ALTER TABLE shortlink
  ADD CONSTRAINT ShortLink_authorId_redirect_key UNIQUE (authorId, redirect),
  ADD CONSTRAINT ShortLink_link_key UNIQUE (link),
  ADD CONSTRAINT ShortLink_authorId_fkey FOREIGN KEY (authorId) REFERENCES "user" (id) ON UPDATE CASCADE;

ALTER TABLE "user"
  ADD CONSTRAINT User_username_key UNIQUE (username),
  ADD CONSTRAINT User_email_key UNIQUE (email);

ALTER TABLE visit
  ADD CONSTRAINT Visit_shortLinkId_fkey FOREIGN KEY (shortLinkId) REFERENCES shortlink (id) ON DELETE CASCADE ON UPDATE CASCADE;
