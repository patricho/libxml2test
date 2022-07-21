## Build and run Docker image

```shell
$ docker build -t libxml2test .
$ docker run --rm -it --entrypoint bash libxml2test
```

## Check installed packages

```shell
$ pkg-config --list-all

    shared-mime-info shared-mime-info - Freedesktop common MIME database
    libxcrypt        libxcrypt - Extended crypt library for DES, MD5, Blowfish and others
    icu-uc           icu-uc - International Components for Unicode: Common and Data libraries
    icu-i18n         icu-i18n - International Components for Unicode: Internationalization library
    libxml-2.0       libXML - libXML library version2.
    icu-io           icu-io - International Components for Unicode: Stream and I/O Library
    libcrypt         libxcrypt - Extended crypt library for DES, MD5, Blowfish and others
```

## Build

```shell
$ go build -x

    WORK=/tmp/go-build2796433898
    mkdir -p $WORK/b001/
    cat >$WORK/b001/importcfg.link << 'EOF' # internal
    ...
```

## Run

```shell
$ go run cmd/html/main.go

    2022/07/21 15:29:19 html
    2022/07/21 15:29:19 head
    2022/07/21 15:29:19 link
    2022/07/21 15:29:19 script
    ...

$ go run cmd/xml/main.go

    2022/07/21 15:56:51 feed
    2022/07/21 15:56:51 title
    2022/07/21 15:56:51 #text
    2022/07/21 15:56:51 id
    ...

$ go run cmd/xsd/main.go

    2022/07/21 16:01:03 error: Element 'foo': No matching global declaration available for the validation root.
```