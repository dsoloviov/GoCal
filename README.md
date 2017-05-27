# GoCal

I hate creating new B-Day events in my Google Calendar. Every time I need to:

* choose the calendar (if it's not primary one)
* enter the name
* go to details and set yearly recurrence

This silly little project is aimed to help me with that :)

## Requirements

* Go 1.7.4
* [ishell](https://github.com/abiosoft/ishell)

## Prerequisites

* Obtain JSON file containing client ID and secret from [Google](https://developers.google.com/google-apps/calendar/quickstart/go) (Step 1).
* Rename the file to `client_secret.json` and move it to `.gocal` directory in your $HOME directory.

## Installation

Install dependencies:

```bash
go get github.com/abiosoft/ishell
go get -u google.golang.org/api/calendar/v3
go get -u golang.org/x/oauth2/...
```

Build the app:

```bash
go build
```

Run the app:

```bash
./gocal
Go to the following link in your browser then type the authorization code: 
https://accounts.google.com/o/oauth2/auth?somearguments
code >>> PASTE_YOUR_CODE_HERE

Connected!
```

Setting authorization code step will be skipped if you already have an access token in place.

## Usage

Choose calendar to be used for creation of events:

```bash
>>> choose
Which calendar to use?
  MyCalendar
❯ Birthday
Use calendar:  Birthday
```

Create event:

```bash
>>> add John Doe May-28
Event created: https://www.google.com/calendar/event?eid=EVENT_DETAILS
```

This will create event called "John Doe's birthday" scheduled for May-28 of current year. Calendar's default notification will be applied.

## Credits
Library | Use
------- | -----
[github.com/abiosoft/ishell](https://github.com/abiosoft/ishell) | command line interface

## TODO

* Get user's consent instead of forcing people to create Google Dev projects (I know it's insane)
* Batch creation (e.g. from the file)
* Be more flexible in user's input (e.g. accept only one word for name)
* Make events configurable (e.g. time, timezone, duration)