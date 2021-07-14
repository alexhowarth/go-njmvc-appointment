# Check for New Jersey MVC licence renewal appointments.

I need to snag an in person driving licence renewal appointment before my licence expires.

I wrote this, rather than spend all week [reloading the appointment page](https://telegov.njportal.com/njmvc/AppointmentWizard/11), in the hope that an appointment opens up.

Maybe someone else will find this useful.

## Installation

```bash
% go install github.com/alexhowarth/go-njmvc-appointment@latest
```

## Usage

For all locations:

```bash
% $GOPATH/bin/go-njmvc-appointment
Vineland          2021-07-15 10:00:00 -0400 EDT
Lawrenceville     2021-07-15 11:20:00 -0400 EDT
Salem             2021-07-20 08:40:00 -0400 EDT
North Cape May    2021-07-20 14:20:00 -0400 EDT
Egg Harbor Twp    2021-07-22 14:20:00 -0400 EDT
Thorofare         2021-07-27 08:40:00 -0400 EDT
Delanco           2021-07-29 09:40:00 -0400 EDT
Toms River        2021-08-03 09:40:00 -0400 EDT
Flemington        2021-08-04 12:20:00 -0400 EDT
Freehold          2021-08-04 12:40:00 -0400 EDT
Camden            2021-08-05 11:40:00 -0400 EDT
Eatontown         2021-08-06 09:20:00 -0400 EDT
Oakland           2021-08-09 09:20:00 -0400 EDT
Paterson          2021-08-10 08:40:00 -0400 EDT
Randolph          2021-08-10 13:20:00 -0400 EDT
Wayne             2021-08-11 10:40:00 -0400 EDT
South Plainfield  2021-08-11 13:40:00 -0400 EDT
Rahway            2021-08-13 10:20:00 -0400 EDT
Lodi              2021-08-17 11:40:00 -0400 EDT
Edison            2021-08-18 09:40:00 -0400 EDT
North Bergen      2021-08-18 10:40:00 -0400 EDT
Bayonne           2021-08-25 12:20:00 -0400 EDT
Newark            2021-08-27 13:40:00 -0400 EDT
```

Specify a comma-separated list of locations:

```bash
$GOPATH/bin/go-njmvc-appointment --location Bayonne,Newark,Rahway,Edison,"South Plainfield","North Bergen"
South Plainfield  2021-08-11 13:40:00 -0400 EDT
Rahway            2021-08-13 10:40:00 -0400 EDT
Edison            2021-08-18 10:40:00 -0400 EDT
North Bergen      2021-08-18 11:20:00 -0400 EDT
Bayonne           2021-08-25 12:40:00 -0400 EDT
Newark            2021-08-30 08:40:00 -0400 EDT
```