# Bank Holidays

Package that tries to make it really easy to load and filter bank holidays.

That's it.

## Example Usage

```
srv, err := Build()
if err != nil {
    panic(err)
}

all, err := srv.All()
if err != nil {
    panic(err)
}

// filter down to US events, between someData and someOtherDate (inclusive)
opts := FilterOpts{
    Country: "US",
	ExcludeBefore: someDate,
    ExcludeBefore: someOtherDate,
}

filtered := srv.Filter(opts)
```
