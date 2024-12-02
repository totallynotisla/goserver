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

CREATE TABLE IF NOT EXISTS users (
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
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'resetpassword_userid_key') THEN
    ALTER TABLE resetpassword
      ADD CONSTRAINT ResetPassword_userId_key UNIQUE (userId);
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'resetpassword_otp_key') THEN
    ALTER TABLE resetpassword
      ADD CONSTRAINT ResetPassword_otp_key UNIQUE (otp);
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'resetpassword_userid_fkey') THEN
    ALTER TABLE resetpassword
      ADD CONSTRAINT ResetPassword_userId_fkey FOREIGN KEY (userId) REFERENCES users (id) ON UPDATE CASCADE;
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'session_userid_fkey') THEN
    ALTER TABLE session
      ADD CONSTRAINT Session_userId_fkey FOREIGN KEY (userId) REFERENCES users (id) ON UPDATE CASCADE;
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'shortlink_authorid_redirect_key') THEN
    ALTER TABLE shortlink
      ADD CONSTRAINT ShortLink_authorId_redirect_key UNIQUE (authorId, redirect);
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'shortlink_link_key') THEN
    ALTER TABLE shortlink
      ADD CONSTRAINT ShortLink_link_key UNIQUE (link);
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'shortlink_authorid_fkey') THEN
    ALTER TABLE shortlink
      ADD CONSTRAINT ShortLink_authorId_fkey FOREIGN KEY (authorId) REFERENCES users (id) ON UPDATE CASCADE;
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'user_username_key') THEN
    ALTER TABLE users
      ADD CONSTRAINT User_username_key UNIQUE (username);
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'user_email_key') THEN
    ALTER TABLE users
      ADD CONSTRAINT User_email_key UNIQUE (email);
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'visit_shortlinkid_fkey') THEN
    ALTER TABLE visit
      ADD CONSTRAINT Visit_shortLinkId_fkey FOREIGN KEY (shortLinkId) REFERENCES shortlink (id) ON DELETE CASCADE ON UPDATE CASCADE;
  END IF;
END $$;