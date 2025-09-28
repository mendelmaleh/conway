# conway

[![Go Documentation](https://godocs.io/github.com/mendelmaleh/conway?status.svg)](https://godocs.io/github.com/mendelmaleh/conway)

Calendar conversion algorithms from Professor John A. Conway ([paper](files/conway-agus-slusky.pdf)), informed by David Bendory's [implementation](https://github.com/bendory/conway-hebrew-calendar).

> [!WARNING]
> This implementation is still new, use at your own risk.

## cmd/ical

Generate an [iCalendar file](files/hebrew.ics) with the Hebrew Holidays:

```sh
go run ./cmd/ical -year 5786 -diaspora > files/hebrew.ics
```

![An image of a Tishrei calendar](files/calendar.png "Tishrei calendar")
