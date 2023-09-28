-- name: CreateTransfer :one
insert into transfer (from_account_id, to_account_id, amount)
values ($1, $2, $3) returning *;

-- name: GetTransfer :one
select *
from transfer
where id = $1
    limit 1;

-- name: TransferList :many
select *
from transfer
where from_account_id = $1
   or to_account_id = $2
order by id
    limit $3 offset $4;