# zip.poke

Pickle for [GNU poke][poke] 2.0 or later to read [ZIP files][zip].

Like [zipdetails], this tool allows you to read individual records in a ZIP
file.

The ZIP format is documented in <https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT>.

## Installation

Copy the `zip.pk` file to `~/.local/share/poke`.

```console
$ mkdir -p ~/.local/share/poke

$ wget -O ~/.local/share/poke/zip.pk https://github.com/ayosec/zip.pk/raw/master/zip.pk
```

## Usage

To load this pickle use `load zip` expression:

    $ poke

    (poke) load zip

If the module can't be loaded, ensure that  `~/.local/share/poke` is in the
value of [`load_path`][poke-modules].

To inspect functions and types defined in the pickle, you can use the
[`.info`][poke-info] command:

    (poke) .info functions zip

    (poke) .info types ZIP

### Example

The following example opens a new Poke session, load this pickle, open the
`a.zip` file, and lists its content:

    $ poke

    (poke) load zip

    (poke) .file a.zip

    (poke) zip_list_files
    ```



[poke-info]: https://www.jemarch.net/poke-2.3-manual/html_node/info-command.html
[poke-modules]: https://www.jemarch.net/poke-2.3-manual/html_node/Loading-pickles-as-Modules.html
[poke]: https://www.gnu.org/software/poke/
[zip]: https://en.wikipedia.org/wiki/ZIP_(file_format)
[zipdetails]: https://perldoc.perl.org/zipdetails
