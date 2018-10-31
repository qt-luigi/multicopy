# multicopy

multicopy copies the specified file after renaming the same name file in the specified directory.

## Usage

```sh
$ multicopy
[args] srcfile dstpath
```

## Example

```sh
$ multicopy test/file.txt test/

Before:

test/
|
+--file.txt (Contents:"src")
+--dir1/
|  +--file.txt (Contents:"dir1")
+--dir2/
   +--dir3/
      +--file.txt (Contents:"dir2/dir3")

After:

test/
|
+--file.txt (Contents:"src")
+--dir1/
|  +--file.txt (Contents:"src")
|  +--file.txt.YYYYMMDDHHNNSSZZZ (Contents:"dir1")
+--dir2/
   +--dir3/
      +--file.txt (Contents:"src")
      +--file.txt.YYYYMMDDHHNNSSZZZ (Contents:"dir2/dir3")
```

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)
