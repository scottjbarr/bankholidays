# Bank Holidays

Package that tries to make it really easy to load and filter bank holidays.

That's it.

## Example Usage

```
// create a service you can pass around
srv, err := Build()
if err != nil {
    panic(err)
}

// all of the holidays
all, err := srv.All()
if err != nil {
    panic(err)
}

// filter holidays down to US events, between someData and someOtherDate (inclusive)
opts := FilterOpts{
    Country: "US",
	ExcludeBefore: someDate,
    ExcludeBefore: someOtherDate,
}

filtered := srv.Filter(opts)
```

## More?

If you want to load your own data, you can do that with the `LoadFromReader(io.Reader)` function.

See `data/holidays.csv` for the format.

## Data source

The data came from a Grok chat.
