CREATE TABLE Users (
  ID  int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  UserName nvarchar(255) NOT NULL,
  Parent int NULL
);
ALTER TABLE Users ADD FOREIGN KEY (Parent) REFERENCES Users (ID);

SELECT
  child.ID,
  child.UserName,
  COALESCE(parent.UserName, NULL) as 'ParentName'
FROM Users child
  LEFT JOIN Users parent ON (child.Parent = parent.ID);
  