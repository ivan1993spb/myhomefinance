SELECT t1.document_guid, t1.unixtimestamp, t1.name, t1.amount, ROUND(SUM(t2.amount), 2) AS balance
FROM (
    SELECT document_guid, unixtimestamp, name, amount, description
    FROM inflow
    WHERE unixtimestamp BETWEEN 1 AND (strftime('%s', 'now'))

    UNION

    SELECT document_guid, unixtimestamp, name, -amount AS amount, description
    FROM outflow
    WHERE unixtimestamp BETWEEN 1 AND (strftime('%s', 'now'))
) AS t1, (
    SELECT document_guid, unixtimestamp, name, amount, description
    FROM inflow

    UNION

    SELECT document_guid, unixtimestamp, name, -amount AS amount, description
    FROM outflow
) AS t2
WHERE t2.unixtimestamp <= t1.unixtimestamp
GROUP BY t1.document_guid
ORDER BY t1.unixtimestamp DESC;
