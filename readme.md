# tmcac

## overview

i'm trying to get in to the rental car market, and would like to be able to do so in a way i like. Turo is the marketplace i will be using, due to an exclusivity clause, and because of that i'm building a tool to help better understand my best bet on a profitable car.

## cli

the cli will allow you to specify a series of details, and get back a table.

### range

```sh
$ tmcac range 30-45 'tucson, AZ' convertible
year | make |      model      | trim | daily |   kbb
2018   audi    a5 cabriolet     2.0t    117     32,000
```

this command allows you to specify a cost, area, and vehicle type to get back a list of the best choices for you.
