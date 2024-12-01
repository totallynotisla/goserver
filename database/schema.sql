USE gerawana;

CREATE TABLE IF NOT EXISTS resetpassword (
  id varchar(191) NOT NULL,
  expires_at datetime(3) NOT NULL DEFAULT current_timestamp(3),
  userId varchar(191) NOT NULL,
  otp varchar(191) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY ResetPassword_userId_key (userId),
  UNIQUE KEY ResetPassword_otp_key (otp)
);

CREATE TABLE IF NOT EXISTS session (
  token varchar(191) NOT NULL,
  expires_at datetime(3) NOT NULL DEFAULT current_timestamp(3),
  userId varchar(191) NOT NULL,
  PRIMARY KEY (token),
  KEY Session_userId_fkey (userId)
);

CREATE TABLE IF NOT EXISTS shortlink (
  id varchar(191) NOT NULL,
  redirect varchar(191) NOT NULL,
  link varchar(191) NOT NULL,
  authorId varchar(191) NOT NULL,
  name varchar(191) NOT NULL,
  createdAt datetime(3) NOT NULL DEFAULT current_timestamp(3),
  PRIMARY KEY (id),
  UNIQUE KEY ShortLink_authorId_redirect_key (authorId,redirect),
  UNIQUE KEY ShortLink_link_key (link)
);

CREATE TABLE IF NOT EXISTS user (
  id varchar(191) NOT NULL,
  email varchar(191) NOT NULL,
  password varchar(191) NOT NULL,
  username varchar(191) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY User_username_key (username),
  UNIQUE KEY User_email_key (email)
);

CREATE TABLE IF NOT EXISTS visit (
  id varchar(191) NOT NULL,
  shortLinkId varchar(191) NOT NULL,
  visitedAt datetime(3) NOT NULL DEFAULT current_timestamp(3),
  PRIMARY KEY (id),
  KEY Visit_shortLinkId_fkey (shortLinkId)
);

ALTER TABLE resetpassword
  ADD CONSTRAINT ResetPassword_userId_fkey FOREIGN KEY (userId) REFERENCES user (id) ON UPDATE CASCADE;

ALTER TABLE session
  ADD CONSTRAINT Session_userId_fkey FOREIGN KEY (userId) REFERENCES user (id) ON UPDATE CASCADE;

ALTER TABLE shortlink
  ADD CONSTRAINT ShortLink_authorId_fkey FOREIGN KEY (authorId) REFERENCES user (id) ON UPDATE CASCADE;

ALTER TABLE visit
  ADD CONSTRAINT Visit_shortLinkId_fkey FOREIGN KEY (shortLinkId) REFERENCES shortlink (id) ON DELETE CASCADE ON UPDATE CASCADE;

