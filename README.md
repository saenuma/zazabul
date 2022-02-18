# zazabul

A configuration file format

[![GoDoc](https://godoc.org/github.com/saenuma/zazabul?status.svg)](https://godoc.org/github.com/saenuma/zazabul)

## Sample Configuration
```
// email is used for communication
// email must follow the format for email ie username@host.ext
// email is compulsory
email: banker@banban.com

// region should be gotten from google cloud documentation
region: us-central1

// zone should be gotten from google cloud documentation
// zone usually derived from the regions and ending with -a or -b or -c
zone: us-central1-a

```

## Notes 
For grouping, use different files.

## License 

MIT
