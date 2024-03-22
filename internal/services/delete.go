package services

func (s *Services) DeleteCategory(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteCategory(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteTag(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteTag(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteNewsTag(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteNewsTag(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteNewsCategory(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteNewsCategory(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteNews(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteNews(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}
