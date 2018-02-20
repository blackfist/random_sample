# random_sample
Take a list of newline separated values shuffles them.
Use the `-n` option to specify the number of returned values. Defaults to all

# example

```
pbpaste | cut -d" " -f2 | random_sample
```

```
pbpaste | random_sample -n 25
```