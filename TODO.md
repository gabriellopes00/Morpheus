user:
select ticket options and lots quantity
click on "buy now"

system
receive request with (user-id, event-id, ticket-opt-id, ticket-opt-lot-id)
validate if the event isn't sold-out

if event not sold-out ...
reduce 1 ticket of the available amount (reserve ticket for 10 min)
user go to checkout
await successful event from stripe ...
create a new ticket in `ticket's` service (... all required info)
