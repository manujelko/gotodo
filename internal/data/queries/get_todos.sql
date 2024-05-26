SELECT
    id,
    title,
    content,
    completed,
    created_at,
    completed_at
FROM todos
LIMIT $1
OFFSET $2