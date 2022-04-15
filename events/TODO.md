[x] add `where not canceled` in events searches
[] add corruption layer for events searches
[] validate if event start date is valid for event update
[] update event usecase

###### after created...

- update event non-critical data
- add ticket option
- remove ticket option (if there is no ticket boughs of this option)
- update ticket option non-critical data
- remove ticket option lot
- update ticket option lot price and quantity

#### buy a ticket
- required info (account_id, event_id, ticket_option_id, lot_id, buyed_at)
- go to checkout
- when start checkout the ticket will be reserved until its payed or 20min timout or when user cancel the checkout
- collect credit card credentials
- go to stripe payments
- if accepted
  - create a ticket to the user (all ticket credentials, validation url, exp date...)

#### valide a ticket

- id
- owner-id -> buyer id
- event-id
- tiket-option-id
- ticket-option-lot-id
- event-organizer-id
- expiration-date
- status -> (not-used, used)

- frontend will generate a qr code to the validation url
- the organizer will access the validation url passing its account_id (to validate a route, the account id received must be the account id of the event organizer)
- once validated, the ticket wont be able to be validated again
