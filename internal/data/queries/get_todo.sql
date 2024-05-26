SELECT
    id,
    title,
    content,
    completed,
    created_at, 
    completed_at
FROM todos
WHERE id = $1;