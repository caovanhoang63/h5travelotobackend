SET GLOBAL event_scheduler = ON;

CREATE EVENT IF NOT EXISTS update_expired_status
ON SCHEDULE EVERY 15 MINUTE
STARTS CURRENT_TIMESTAMP
DO
UPDATE payment_bookings
SET payment_status = 'expired'
WHERE (payment_status = 'not_started' OR payment_status = 'executing')
AND TIMESTAMPDIFF(MINUTE , created_at,NOW()) >= 15;

CREATE EVENT IF NOT EXISTS update_expired_status_booking
ON SCHEDULE EVERY 5 MINUTE
STARTS CURRENT_TIMESTAMP
DO
UPDATE bookings
SET state = 'expired'
WHERE (pay_in_hotel = 0 AND  state = 'pending')
AND TIMESTAMPDIFF(HOUR, created_at, NOW()   ) >= 1;


