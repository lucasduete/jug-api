CREATE TABLE User(
  Email VARCHAR(25),
  Nome VARCHAR(50) NOT NULL,
  Username VARCHAR(10) NOT NULL,
  Senha VARCHAR(60) NOT NULL,
  CONSTRAINT User_PK_Email PRIMARY KEY (Email),
  CONSTRAINT User_Email_Valido CHECK (Email LIKE '%@%'),
  CONSTRAINT User_Username_Unique UNIQUE (Username)
)