package auth

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindOrCreateUser(userId string) (*User, error) {
	return s.repo.FindOrCreateUser(userId)
}

func (s *Service) GetUserById(userId string) (*User, error) {
	return s.repo.GetUserById(userId)
}


// func (s *Service) HandleMoviePurchase(req MoviePurchaseRequest) (*CampaignUser, error) {
// 	// 1. Resolve content credit
// 	content, err := s.repo.FindContentByName(req.ContentName)
// 	if err != nil {
// 		return nil, fmt.Errorf("unknown content %q: %w", req.ContentName, err)
// 	}

// 	// 2. Find or create the campaign user
// 	user, err := s.repo.FindOrCreateUser(req.SubID)
// 	if err != nil {
// 		return nil, fmt.Errorf("find or create user: %w", err)
// 	}

// 	// 3. Record the purchase
// 	purchase := &MoviePurchase{
// 		SubID:         req.SubID,
// 		ContentName:   req.ContentName,
// 		PointsAwarded: content.Credit,
// 	}
// 	if err := s.repo.CreateMoviePurchase(purchase); err != nil {
// 		return nil, fmt.Errorf("record purchase: %w", err)
// 	}

// 	// 4. Award points
// 	user, err = s.repo.UpdateUserPoints(req.SubID, content.Credit)
// 	if err != nil {
// 		return nil, fmt.Errorf("award points: %w", err)
// 	}

// 	return user, nil
// }