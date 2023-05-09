# pdfx
A simple command line tool used for Splitting and Merging Pdfs.

You can download the precompiled binary from [here](https://github.com/Raiyanx/pdfx/releases).

Note: This is just a self project using a [3rd party api](https://developer.ilovepdf.com/). The api limits the amount of file uploads per month. The project is not for everyday use.

## Merging

Merging can be done by the command

```
pdfx [-flags] merge [files]
```

## Splitting

Splitting can be done by the command

```
pdfx [-flags] split filename [Integer/Integer-Integer]
```

As an example, say we have a pdf called somepdf.pdf. We need to make three more pdfs from it containing pages 2-4, only page 3, and pages 4-7 respectively. Then we can use the command

```
pdfx split somepdf.pdf 2-4 3 4-7
```

## Flag

One can use the flag `-dn` to set the default name of a pdf.
