package database

import "github.com/google/uuid"

func (db *DB) GetSubscription(id uuid.UUID) (Subscription, error) {
	query := `
	SELECT id, service_name, price_rub, user_id, start_date, end_date
	FROM subscriptions
	WHERE id = $1;`

	var sub Subscription
	err := db.QueryRow(query, id).Scan(&sub.ID, &sub.ServiceName, &sub.PriceRub, &sub.UserID, &sub.StartDate, &sub.EndDate)
	if err != nil {
		return Subscription{}, err
	}

	return sub, nil
}

type SaveSubscriptionParams struct {
	ID          uuid.UUID
	ServiceName string
	PriceRub    int
	UserID      uuid.UUID
	StartDate   string
	EndDate     string
}

func (db *DB) SaveSubscription(params SaveSubscriptionParams) (Subscription, error) {
	query := `
		INSERT INTO subscriptions (id, service_name, price_rub, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, service_name, price_rub, user_id, start_date, end_date`

	var sub Subscription
	err := db.QueryRow(
		query,
		params.ID,
		params.ServiceName,
		params.PriceRub,
		params.UserID,
		params.StartDate,
		params.EndDate,
	).Scan(
		&sub.ID,
		&sub.ServiceName,
		&sub.PriceRub,
		&sub.UserID,
		&sub.StartDate,
		&sub.EndDate,
	)
	if err != nil {
		return Subscription{}, err
	}
	return sub, nil
}

func (db *DB) ChangeSubscription(id uuid.UUID, service_name string) (Subscription, error) {
	query := `
	UPDATE subscriptions
	SET service_name = $2
	WHERE id = $1
	RETURNING id, service_name, price_rub, user_id, start_date, end_date`

	var sub Subscription
	err := db.QueryRow(query, id, service_name).Scan(&sub.ID, &sub.ServiceName, &sub.PriceRub, &sub.UserID, &sub.StartDate, &sub.EndDate)
	if err != nil {
		return Subscription{}, err
	}

	return sub, nil
}

func (db *DB) DeleteSubscription(id uuid.UUID) error {
	query := `
	DELETE FROM subscriptions WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) TotalSumSubscriptions(startingDate string) (int, error) {
	query := `
    SELECT COALESCE(SUM(price_rub), 0)
    FROM subscriptions
    WHERE start_date = $1;`

	var sum int
	err := db.QueryRow(query, startingDate).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
