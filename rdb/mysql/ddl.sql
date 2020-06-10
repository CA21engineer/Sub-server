INSERT INTO
  `icons` (`icon_id`, `icon_uri`, `is_original`)
VALUES
  (
    '1',
    'https://images-fe.ssl-images-amazon.com/images/I/411j1k1u9yL.png',
    false
  ),
  (
    '2',
    'https://www.youtube.com/yts/img/yt_1200-vflhSIVnY.png',
    false
  ),
  (
    '3',
    'https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRuX7izxLGFnXQ7k79lGWEew3njHyI2NCmkq3-y_RN3An1lS7cj&usqp=CAU',
    true
  );

INSERT INTO
  `subscriptions`(
    `subscription_id`,
    `icon_id`,
    `service_name`,
    `service_type`,
    `price`,
    `cycle`,
    `is_original`,
    `free_trial`
  )
VALUES
  ('1', '1', 'アマゾン', '3', 1080, 1, false, 1),
  ('2', '2', 'Youtube', '4', 1080, 1, false, 1),
  ('3', '3', '鬼滅の刃', '4', 1080, 1, true, 1);

INSERT INTO
  `user_subscriptions`(
    `user_subscription_id`,
    `subscription_id`,
    `user_id`,
    `price`,
    `cycle`
  )
VALUES
  ('1', '1', 'user_token1', 1080, 1),
  ('2', '2', 'user_token2', 1080, 1),
  ('3', '3', 'user_token2', 1500, 3);
