# lets-go-bibit

## What's going on here?
- [DB Query](#01_Simple-Database-Query)
- [OMDB API](#02_OMDB_API)
- [Refactor](#03_Refactor)
- [Logic](#04_Logic)
- [Algorithmic](#05_algorithmic)
- [Bonus: Parallel and Concurrency](#06_parallel-concurrency)
- [Bonus: Rounded Matrix Path](#07_rounding-path)

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

## 02_OMDB_API
> Movies Microservices Task

Reoisitory: https://github.com/hrz8/go-seeding-omdb

## 03_Refactor
> Refactor the function

```go
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}
```

## 04_Logic

> Write a function to grouping list of string given based on anagram.

### Play the Game

```bash
[lets-go-bibit]$ cd 04_Logic/
[04_Logic]$ go run main.go
```
