Sums arrays of int64's from flat files.

There's a naive bufio one-at-time, a batch reader that uses unsafe to convert
the byte array into an int64 array, and a mmap'd version that maps the whole
thing into a giant int64 array using unsafe.

FWIW -- MAP_POPULATE was key to getting the mmap version to outperform the
batch reader using regular I/O.

Usage:

    $ go build write_out.go
    $ write_out 25000000 > 25m.dat

    $ go build read_mmap.go
    $ time ./read_mmap 25m.dat

