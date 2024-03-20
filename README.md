# Golang Search-Engine

- This is a simple Full-Text-Search (FTS) made using golang.
- Made using reference from https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/

- We are going to search a part of the abstract of English Wikipedia. The latest dump is available at 
[dumps.wikimedia.org](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz). As of today, the file size after decompression is 913 MB. The XML file contains over 600K documents.

## Steps to run:
```bash


$ go build
$ ./Search-Engine
```

