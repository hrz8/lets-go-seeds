# lets-go-bibit

## What's going on here?
- [DB Query](#01_Simple-Database-Query)
- [OMDB API](#02_OMDB-API)
- [Refactor](#03_Refactor)
- [Logic](#04_Logic)

## 01_Simple-Database-Query
```sql
# Create Users Table
CREATE TABLE Users (
  ID  int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  UserName nvarchar(255) NOT NULL,
  Parent int NULL
);
ALTER TABLE Users ADD FOREIGN KEY (Parent) REFERENCES Users (ID);

# Query for Select User and ParentName
SELECT
  child.ID,
  child.UserName,
  COALESCE(parent.UserName, NULL) as 'ParentName'
FROM Users child
  LEFT JOIN Users parent ON (child.Parent = parent.ID);
```

## 02_OMDB-API
> Movies Microservices Task

Reoisitory: https://github.com/hrz8/go-seeding-omdb

## 03_Refactor
> Refactor the function

```go
func FindFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if len(str) == 0 || indexFirstBracketFound < 0 {
		return ""
	}
	strRunes := []rune(str)
	wordsAfterFirstBracket := string(strRunes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}
	afterBracketRunes := []rune(wordsAfterFirstBracket)
	return strings.Split(string(afterBracketRunes[1:indexClosingBracketFound]), " ")[0]
}
```
### Play the Game

```bash
[lets-go-bibit]$ cd 03_Refactor/
[04_Logic]$ go test *.go
```

## 04_Logic

> Write a function to grouping list of string given based on anagram.

### Play the Game

```bash
[lets-go-bibit]$ cd 04_Logic/
[04_Logic]$ go run main.go
```
