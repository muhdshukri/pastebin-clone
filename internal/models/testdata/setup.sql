CREATETABLEsnippets(
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
);CREATEINDEX idx_snippets_created
  ON snippets(created);CREATETABLEusers(
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255)NOT NULL,
  email VARCHAR(255) NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created DATETIME NOT NULL
);ALTERTABLEusersADDCONSTRAINTusers_uc_emailUNIQUE(
    email
  );INSERT INTOusers(
    name,
    email,
    hashed_password,
    created
  )
VALUES(
  'Alice Jones',
  'alice@example.com''$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
  '2022-01-01 10:00:00'
);
