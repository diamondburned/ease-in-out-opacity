# ease-in-out-opacity

Semi-gradient generator for opacity values using the ease-in-out curve for a
more natural transition.

Usage:

```sh
# used for Dissent
$ go run . -max 0 -min 0.35 -start 42px -end 84px -steps 8
alpha(black, 0.35) 42px,
alpha(black, 0.35) 48px,
alpha(black, 0.32) 54px,
alpha(black, 0.24) 60px,
alpha(black, 0.11) 66px,
alpha(black, 0.03) 72px,
alpha(black, 0.00) 78px,
alpha(black, 0.00) 84px
```

Then, just use it as CSS:

```css
background-image: linear-gradient(
    to top,
    alpha(black, 0.35) 42px,
    alpha(black, 0.35) 48px,
    alpha(black, 0.32) 54px,
    alpha(black, 0.24) 60px,
    alpha(black, 0.11) 66px,
    alpha(black, 0.03) 72px,
    alpha(black, 0.00) 78px,
    alpha(black, 0.00) 84px
);
```

If you're not writing this for GTK, use the `-f` flag to change it to use
`rgba()` instead.
