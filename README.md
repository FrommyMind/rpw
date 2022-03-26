A example use of cobra.

1. mkdir rpw
    ```shell
    mkdir  rpw
    ```
2. init mod
   ```
    go mod init github.com/frommymind/rpw
   ```
3. use [cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md) init project.
    ```shell
   cobra-cli init
    ```
4. update `cmd/root.go` code.


rpw:

    ```shell
    rpw (Random Password) is used for randomly generator password

    Usage:
    rpw [flags]

    Flags:
    -h, --help         help for rpw
    -l, --length int   The length of the password (default 12)
    -e, --level int    The level of complexity
                            0 : Only use lower case alphabet [a-z] and numers [0-9]
                            1 : Use upper case alphabet [A-z],  plus level 0
                            2 : Use symbols [-], plus level 1
                            3 : Use symbols [@#$%+?], plus level 2
```
