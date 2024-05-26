UPDATE todos
SET completed = TRUE, completed_at = NOW()
WHERE id = $1;