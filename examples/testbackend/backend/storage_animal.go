package backend

import ()

func (storage *PostgreStorage) CreateAnimal(in *Animal) (*Animal, error) {
	stmt, err := storage.db.Prepare(`
		INSERT INTO animals(
			name,
			age,
			photo_url
		)
		VALUES ($1,$2,$3)
		RETURNING id;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		 in.Name,
		 in.Age,
		 in.PhotoUrl,
	).Scan(
		&(in.Id),
	)

	return in, err
}


func (storage *PostgreStorage) DeleteAnimal(id uint32) error {
	stmt, err := storage.db.Prepare("DELETE FROM animals WHERE id=$1;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}


func (storage *PostgreStorage) GetAnimal(id uint32) (*Animal, error) {
	stmt, err := storage.db.Prepare(`
		SELECT
			id,
			name,
			age,
			photo_url
		FROM animals
		WHERE id=$1;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := &Animal{}
	err = stmt.QueryRow(id).Scan(
		&ret.Id,
		&ret.Name,
		&ret.Age,
		&ret.PhotoUrl,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}


func (storage *PostgreStorage) UpdateAnimal(updated *Animal) (*Animal, error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
		UPDATE animals
		SET
			name=$1,
			age=$2,
			photo_url=$3
		WHERE
			id=$5
		RETURNING
			id,
			name,
			age,
			photo_url
		;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := &Animal{}
	err = stmt.QueryRow(
		updated.Name,
		updated.Age,
		updated.PhotoUrl,
	).Scan(
		&ret.Id,
		&ret.Name,
		&ret.Age,
		&ret.PhotoUrl,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return ret, err
}
