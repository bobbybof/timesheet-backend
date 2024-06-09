INSERT INTO users (name, rate)
VALUES
    ('Maraka Makaka', 12000);

INSERT INTO projects (name)
VALUES
    ('Web Develompont'),
    ('Desain UI'),
    ('Desain Logo'),
    ('Dokumentasi');

INSERT INTO activities (
    name,
    date_start,
    date_end,
    project_id,
    user_id
)
VALUES
    (
        'Wireframing untuk fitur/flow bidding',
        '2023-12-05 08:00:00',
        '2023-12-05 16:00:00',
        2,
        1
    ),
    (
        'Meeting',
        '2023-12-03 08:50:00',
        '2023-12-03 17:30:00',
        3,
        1
    ),
    (
        'Meeting',
        '2023-12-04 10:30:00',
        '2023-12-04 15:00:00',
        2,
        1
    );

