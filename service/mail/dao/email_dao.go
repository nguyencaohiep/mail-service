package dao

import "mail-service/pkg/db"

type Emails struct {
	Emails []Email
}

type Email struct {
	Email string
}

func (repo *Emails) GetEmails() error {
	query := `SELECT email FROM public.mail;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var email = &Email{}
		err := rows.Scan(&email.Email)
		if err != nil {
			return err
		}

		repo.Emails = append(repo.Emails, *email)
	}

	return nil
}
